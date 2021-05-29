// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2beta1

import (
	"context"
	"fmt"
	"github.com/atomix/atomix-controller/pkg/apis/core/v2beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"strconv"
	"strings"
)

const (
	driverTypeEnv      = "ATOMIX_DRIVER_TYPE"
	driverNodeEnv      = "ATOMIX_DRIVER_NODE"
	driverNamespaceEnv = "ATOMIX_DRIVER_NAMESPACE"
	driverNameEnv      = "ATOMIX_DRIVER_NAME"
)

const (
	driverInjectPath = "/inject-drivers"
)

const (
	storageProfileAnnotation = "storage.atomix.io/profile"
)

func addDriverController(mgr manager.Manager) error {
	mgr.GetWebhookServer().Register(driverInjectPath, &webhook.Admission{
		Handler: &DriverInjector{
			client: mgr.GetClient(),
			scheme: mgr.GetScheme(),
		},
	})
	return nil
}

// DriverInjector is a mutating webhook that injects driver containers into pods
type DriverInjector struct {
	client  client.Client
	scheme  *runtime.Scheme
	decoder *admission.Decoder
}

// InjectDecoder :
func (i *DriverInjector) InjectDecoder(decoder *admission.Decoder) error {
	i.decoder = decoder
	return nil
}

// Handle :
func (i *DriverInjector) Handle(ctx context.Context, request admission.Request) admission.Response {
	namespacedName := types.NamespacedName{
		Namespace: request.Namespace,
		Name:      request.Name,
	}
	log.Infof("Received admission request for Pod '%s'", namespacedName)

	// Decode the pod
	pod := &corev1.Pod{}
	if err := i.decoder.Decode(request, pod); err != nil {
		log.Errorf("Could not decode Pod '%s'", namespacedName, err)
		return admission.Errored(http.StatusBadRequest, err)
	}

	pluginList := &v2beta1.StoragePluginList{}
	if err := i.client.List(ctx, pluginList); err != nil {
		log.Error("Could not list StoragePlugins", err)
		return admission.Errored(http.StatusInternalServerError, err)
	}

	allPlugins := make(map[string]v2beta1.StoragePlugin)
	for _, plugin := range pluginList.Items {
		allPlugins[plugin.Name] = plugin
	}

	injectPlugins := make(map[string]v2beta1.StoragePlugin)
	profileName, ok := pod.Annotations[storageProfileAnnotation]
	if ok {
		profile := &v2beta1.StorageProfile{}
		profileNamespacedName := types.NamespacedName{
			Namespace: request.Namespace,
			Name:      profileName,
		}
		if err := i.client.Get(ctx, profileNamespacedName, profile); err != nil {
			if errors.IsNotFound(err) {
				return admission.Denied(fmt.Sprintf("StorageProfile %s not found", profileNamespacedName))
			}
			return admission.Errored(http.StatusInternalServerError, err)
		}

		for _, driver := range profile.Spec.Drivers {
			plugin := v2beta1.StoragePlugin{}
			pluginNamespacedName := types.NamespacedName{
				Name: driver,
			}
			if err := i.client.Get(ctx, pluginNamespacedName, &plugin); err != nil {
				if errors.IsNotFound(err) {
					return admission.Denied(fmt.Sprintf("StoragePlugin %s not found", profileNamespacedName))
				}
				return admission.Errored(http.StatusInternalServerError, err)
			}
			injectPlugins[plugin.Name] = plugin
		}
	}

	for annotation, value := range pod.Annotations {
		parts := strings.Split(annotation, "/")
		if len(parts) != 2 {
			continue
		}

		domain, path := parts[0], parts[1]
		if path != "inject" {
			continue
		}

		if value != fmt.Sprint(true) {
			continue
		}

		if plugin, ok := allPlugins[domain]; ok {
			injectPlugins[domain] = plugin
		}
	}

	for _, plugin := range injectPlugins {
		for _, version := range plugin.Spec.Versions {
			driverQualifiedName := fmt.Sprintf("%s.%s", version.Name, plugin.Name)
			pluginName := plugin.Name[len(plugin.Name)-len(plugin.Spec.Group):]
			statusAnnotation := fmt.Sprintf("%s/status", driverQualifiedName)
			statusValue := pod.Annotations[statusAnnotation]
			if statusValue == injectedStatus {
				continue
			}

			port := 5680
			for a, v := range pod.Annotations {
				parts := strings.Split(a, "/")
				if len(parts) != 2 {
					continue
				}

				domain, path := parts[0], parts[1]
				if path != "port" {
					continue
				}

				_, ok := injectPlugins[domain]
				if !ok {
					continue
				}

				i, err := strconv.Atoi(v)
				if err != nil {
					log.Errorf("Could not decode port annotation '%s'", a, err)
					return admission.Errored(http.StatusInternalServerError, err)
				} else if i > port {
					port = i + 1
				}
			}

			container := corev1.Container{
				Name:            strings.ReplaceAll(fmt.Sprintf("driver-%s-%s", pluginName, version.Name), ".", "-"),
				Image:           version.Driver.Image,
				ImagePullPolicy: corev1.PullIfNotPresent,
				Args: []string{
					fmt.Sprintf(":%d", port),
				},
				Ports: []corev1.ContainerPort{
					{
						Name:          "driver",
						ContainerPort: int32(port),
					},
				},
				Env: []corev1.EnvVar{
					{
						Name:  driverTypeEnv,
						Value: driverQualifiedName,
					},
					{
						Name:  driverNamespaceEnv,
						Value: namespacedName.Namespace,
					},
					{
						Name:  driverNameEnv,
						Value: namespacedName.Name,
					},
					{
						Name: driverNodeEnv,
						ValueFrom: &corev1.EnvVarSource{
							FieldRef: &corev1.ObjectFieldSelector{
								FieldPath: "spec.nodeName",
							},
						},
					},
				},
			}
			pod.Spec.Containers = append(pod.Spec.Containers, container)

			pod.Annotations[statusAnnotation] = injectedStatus
			portAnnotation := fmt.Sprintf("%s/port", driverQualifiedName)
			pod.Annotations[portAnnotation] = fmt.Sprint(port)

			// Marshal the pod and return a patch response
			marshaledPod, err := json.Marshal(pod)
			if err != nil {
				log.Errorf("Driver injection failed for Pod '%s'", namespacedName, err)
				return admission.Errored(http.StatusInternalServerError, err)
			}
			return admission.PatchResponseFromRaw(request.Object.Raw, marshaledPod)
		}
	}
	return admission.Allowed("No drivers to inject")
}

var _ admission.Handler = &DriverInjector{}
