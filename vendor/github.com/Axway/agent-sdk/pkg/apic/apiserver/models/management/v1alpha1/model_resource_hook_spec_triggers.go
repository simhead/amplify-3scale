/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceHookSpecTriggers struct for ResourceHookSpecTriggers
type ResourceHookSpecTriggers struct {
	// Value for the group of the resource. Use \"*\" for any.
	Group string                `json:"group"`
	Scope ResourceHookSpecScope `json:"scope,omitempty"`
	// Value for the Kind of the resource. Use \"*\" for any.
	Kind string `json:"kind"`
	// Name of the resource. Use \"*\" for any.
	Name string   `json:"name"`
	Type []string `json:"type"`
}
