/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceHookSpec struct for ResourceHookSpec
type ResourceHookSpec struct {
	Triggers []ResourceHookSpecTriggers `json:"triggers"`
	// List of Webhook kind resource names that dictates what webhooks will be invoked on matching triggers.
	Webhooks []string `json:"webhooks"`
}