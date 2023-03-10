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

type BaseConfigurationResource struct {
	// unique identifier for this resource
	Id string `json:"id"`
	// display name for this resource
	DisplayName string `json:"displayName"`
	// a description of this resource
	Description string `json:"description,omitempty"`
}
