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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// PrimitiveSpec is the base specification for all primitives
type PrimitiveSpec struct {
	Store PrimitiveStore `json:"store,omitempty"`
}

// PrimitiveStore is a primitive store configuration
type PrimitiveStore struct {
	corev1.ObjectReference `json:",inline"`
	Config                 map[string]runtime.RawExtension `json:"config,omitempty"`
}
