/*
 * AppDynamics Cloud Connections API
 *
 * Enables you to manage cloud connections
 *
 * API version: 1.0.0
 * Contact: support@appdynamics.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ConfigurationUpdate struct {
	// Name of the configuration
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
	Details *AnyOfConfigurationUpdateDetails `json:"details,omitempty"`
}
