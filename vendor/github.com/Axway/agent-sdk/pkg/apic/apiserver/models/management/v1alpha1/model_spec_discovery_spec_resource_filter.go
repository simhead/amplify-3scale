/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// SpecDiscoverySpecResourceFilter Filter pod resources by name or labels.
type SpecDiscoverySpecResourceFilter struct {
	Names       []string          `json:"names,omitempty"`
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}