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

type Description struct {
	// Name of the connection or configuration
	DisplayName string `json:"displayName"`
	// Description for this connection or configuration
	Description string `json:"description,omitempty"`
}
