/*
 * Amplify Unified Catalog APIs
 *
 * APIs for Amplify Unified Catalog
 *
 * API version: 1.43.0
 * Contact: support@axway.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package unifiedcatalog
// CatalogDefinitionRevision Defines the properties specs for a specific definition
type CatalogDefinitionRevision struct {
	Value int32 `json:"value,omitempty"`
	Revision []CatalogDefinitionPropertySpec `json:"revision,omitempty"`
	Subscription []CatalogDefinitionPropertySpec `json:"subscription,omitempty"`
	CatalogItem []CatalogDefinitionPropertySpec `json:"catalogItem,omitempty"`
}