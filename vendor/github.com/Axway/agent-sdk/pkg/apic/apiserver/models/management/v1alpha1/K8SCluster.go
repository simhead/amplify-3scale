/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_K8SClusterGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "K8SCluster",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	K8SClusterScope = ""

	K8SClusterResourceName = "k8sclusters"
)

func K8SClusterGVK() apiv1.GroupVersionKind {
	return _K8SClusterGVK
}

func init() {
	apiv1.RegisterGVK(_K8SClusterGVK, K8SClusterScope, K8SClusterResourceName)
}

// K8SCluster Resource
type K8SCluster struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec K8SClusterSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a K8SCluster
func (res *K8SCluster) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &K8SClusterSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = K8SCluster{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// K8SClusterFromInstanceArray converts a []*ResourceInstance to a []*K8SCluster
func K8SClusterFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*K8SCluster, error) {
	newArray := make([]*K8SCluster, 0)
	for _, item := range fromArray {
		res := &K8SCluster{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*K8SCluster, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a K8SCluster to a ResourceInstance
func (res *K8SCluster) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = K8SClusterGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
