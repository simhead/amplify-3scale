/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_AssetMappingTemplateGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "AssetMappingTemplate",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	AssetMappingTemplateScope = "Environment"

	AssetMappingTemplateResourceName = "assetmappingtemplates"
)

func AssetMappingTemplateGVK() apiv1.GroupVersionKind {
	return _AssetMappingTemplateGVK
}

func init() {
	apiv1.RegisterGVK(_AssetMappingTemplateGVK, AssetMappingTemplateScope, AssetMappingTemplateResourceName)
}

// AssetMappingTemplate Resource
type AssetMappingTemplate struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec AssetMappingTemplateSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a AssetMappingTemplate
func (res *AssetMappingTemplate) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &AssetMappingTemplateSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = AssetMappingTemplate{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AssetMappingTemplateFromInstanceArray converts a []*ResourceInstance to a []*AssetMappingTemplate
func AssetMappingTemplateFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*AssetMappingTemplate, error) {
	newArray := make([]*AssetMappingTemplate, 0)
	for _, item := range fromArray {
		res := &AssetMappingTemplate{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*AssetMappingTemplate, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a AssetMappingTemplate to a ResourceInstance
func (res *AssetMappingTemplate) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = AssetMappingTemplateGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
