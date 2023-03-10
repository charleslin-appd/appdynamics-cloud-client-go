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

type ConnectionUpdate struct {
	// Updated name for the connection
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
	// Configuration ID assigned to the connection
	ConfigurationId string `json:"configurationId,omitempty"`
	// Set the state of the connection
	State string `json:"state,omitempty"`
	Details *OneOfConnectionUpdateDetails `json:"details,omitempty"`
}
