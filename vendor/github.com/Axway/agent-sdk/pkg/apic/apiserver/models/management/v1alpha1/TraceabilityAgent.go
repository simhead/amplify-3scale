/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_TraceabilityAgentGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "TraceabilityAgent",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	TraceabilityAgentScope = "Environment"

	TraceabilityAgentResourceName = "traceabilityagents"
)

func TraceabilityAgentGVK() apiv1.GroupVersionKind {
	return _TraceabilityAgentGVK
}

func init() {
	apiv1.RegisterGVK(_TraceabilityAgentGVK, TraceabilityAgentScope, TraceabilityAgentResourceName)
}

// TraceabilityAgent Resource
type TraceabilityAgent struct {
	apiv1.ResourceMeta

	// GENERATE: The following code has been modified after code generation
	// 	Owner struct{} `json:"owner"`
	Owner *struct{} `json:"owner,omitempty"`

	Spec TraceabilityAgentSpec `json:"spec"`

	Status TraceabilityAgentStatus `json:"status"`
}

// FromInstance converts a ResourceInstance to a TraceabilityAgent
func (res *TraceabilityAgent) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &TraceabilityAgentSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = TraceabilityAgent{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// TraceabilityAgentFromInstanceArray converts a []*ResourceInstance to a []*TraceabilityAgent
func TraceabilityAgentFromInstanceArray(fromArray []*apiv1.ResourceInstance) ([]*TraceabilityAgent, error) {
	newArray := make([]*TraceabilityAgent, 0)
	for _, item := range fromArray {
		res := &TraceabilityAgent{}
		err := res.FromInstance(item)
		if err != nil {
			return make([]*TraceabilityAgent, 0), err
		}
		newArray = append(newArray, res)
	}

	return newArray, nil
}

// AsInstance converts a TraceabilityAgent to a ResourceInstance
func (res *TraceabilityAgent) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = TraceabilityAgentGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}