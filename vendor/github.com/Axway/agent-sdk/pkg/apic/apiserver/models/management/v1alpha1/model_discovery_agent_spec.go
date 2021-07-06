/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// DiscoveryAgentSpec struct for DiscoveryAgentSpec
type DiscoveryAgentSpec struct {
	// The dataplane type that this agent connects to
	DataplaneType string                    `json:"dataplaneType"`
	Config        DiscoveryAgentSpecConfig  `json:"config"`
	Logging       DiscoveryAgentSpecLogging `json:"logging,omitempty"`
}
