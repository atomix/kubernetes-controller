// Copyright 2019-present Open Networking Foundation.
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
	v2beta1 "github.com/atomix/atomix-controller/pkg/apis/core/v2beta1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// PrimitiveReconciler is a reconciler for Primitive resources
type PrimitiveReconciler struct {
	client client.Client
	scheme *runtime.Scheme
	config *rest.Config
	kind   schema.GroupVersionKind
}

// Reconcile reads that state of the cluster for a primitive object and makes changes based on the pod's annotations
func (r *PrimitiveReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	log.Infof("Reconciling %s '%s'", r.kind.Kind, request.NamespacedName)

	object := &unstructured.Unstructured{}
	object.SetGroupVersionKind(r.kind)
	err := r.client.Get(context.TODO(), request.NamespacedName, object)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			log.Errorf("Reconciling %s '%s' failed", r.kind.Kind, request.NamespacedName, err)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	primitive := &v2beta1.Primitive{}
	err = r.client.Get(context.TODO(), request.NamespacedName, primitive)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			log.Errorf("Reconciling %s '%s' failed", r.kind.Kind, request.NamespacedName, err)
			return reconcile.Result{}, err
		}

		log.Infof("Creating Primitive '%s'", request.NamespacedName)

		var reference corev1.ObjectReference
		if name, ok, err := unstructured.NestedString(object.UnstructuredContent(), "spec", "store", "name"); ok {
			reference.Name = name
		} else if err != nil {
			log.Errorf("Reconciling %s '%s' failed", r.kind.Kind, request.NamespacedName, err)
			return reconcile.Result{}, err
		}
		if namespace, ok, err := unstructured.NestedString(object.UnstructuredContent(), "spec", "store", "namespace"); ok {
			reference.Namespace = namespace
		} else if err != nil {
			log.Errorf("Reconciling %s '%s' failed", r.kind.Kind, request.NamespacedName, err)
			return reconcile.Result{}, err
		}

		primitive = &v2beta1.Primitive{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:   request.Namespace,
				Name:        request.Name,
				Labels:      object.GetLabels(),
				Annotations: object.GetAnnotations(),
			},
			Spec: v2beta1.PrimitiveSpec{
				Type:  r.kind.Kind,
				Store: reference,
			},
		}
		if err := controllerutil.SetControllerReference(object, primitive, r.scheme); err != nil {
			log.Errorf("Creating Primitive '%s' failed", request.NamespacedName, err)
			return reconcile.Result{}, err
		}
		if err := r.client.Create(context.TODO(), primitive); err != nil {
			if !k8serrors.IsAlreadyExists(err) {
				log.Errorf("Creating Primitive '%s' failed", request.NamespacedName, err)
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}
	return reconcile.Result{}, nil
}

var _ reconcile.Reconciler = &PrimitiveReconciler{}
