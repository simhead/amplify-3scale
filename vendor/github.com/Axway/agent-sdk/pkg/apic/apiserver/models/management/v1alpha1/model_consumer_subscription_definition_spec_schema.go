/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ConsumerSubscriptionDefinitionSpecSchema Subscription definition properties to be used by the consumers.
type ConsumerSubscriptionDefinitionSpecSchema struct {
	// Defines the subscription schema properties as key/value pairs.
	Properties []ConsumerSubscriptionDefinitionSpecSchemaProperties `json:"properties,omitempty"`
}
