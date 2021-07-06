/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_APIServiceInstanceGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "APIServiceInstance",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	APIServiceInstanceScope = "Environment"

	APIServiceInstanceResourceName = "apiserviceinstances"
)

func APIServiceInstanceGVK() apiv1.GroupVersionKind {
	return _APIServiceInstanceGVK
}

func init() {
	apiv1.RegisterGVK(_APIServiceInstanceGVK, APIServiceInstanceScope, APIServiceInstanceResourceName)
}

// APIServiceInstance Resource
type APIServiceInstance struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec ApiServiceInstanceSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a APIServiceInstance
func (res *APIServiceInstance) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ApiServiceInstanceSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = APIServiceInstance{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// APIServiceInstanceFromInstanceArray converts a []*ResourceInstance to a []*APIServiceInstance
func APIServiceInstanceFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*APIServiceInstance, error) {
	newArray := make([]*APIServiceInstance, 0)
	for _, item := range fromArray {
		res := &APIServiceInstance{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*APIServiceInstance, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a APIServiceInstance to a ResourceInstance
func (res *APIServiceInstance) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = APIServiceInstanceGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
