/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_SpecDiscoveryGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "SpecDiscovery",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	SpecDiscoveryScope = "K8SCluster"

	SpecDiscoveryResourceName = "specdiscoveries"
)

func SpecDiscoveryGVK() apiv1.GroupVersionKind {
	return _SpecDiscoveryGVK
}

func init() {
	apiv1.RegisterGVK(_SpecDiscoveryGVK, SpecDiscoveryScope, SpecDiscoveryResourceName)
}

// SpecDiscovery Resource
type SpecDiscovery struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec SpecDiscoverySpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a SpecDiscovery
func (res *SpecDiscovery) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &SpecDiscoverySpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = SpecDiscovery{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// SpecDiscoveryFromInstanceArray converts a []*ResourceInstance to a []*SpecDiscovery
func SpecDiscoveryFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*SpecDiscovery, error) {
	newArray := make([]*SpecDiscovery, 0)
	for _, item := range fromArray {
		res := &SpecDiscovery{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*SpecDiscovery, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a SpecDiscovery to a ResourceInstance
func (res *SpecDiscovery) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = SpecDiscoveryGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}