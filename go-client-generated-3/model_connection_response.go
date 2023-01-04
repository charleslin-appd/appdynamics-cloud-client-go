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

type ConnectionResponse struct {
	// Configuration ID assigned to the connection
	ConfigurationId string `json:"configurationId,omitempty"`
	Details *ConnectionResponseDetails `json:"details"`
}
