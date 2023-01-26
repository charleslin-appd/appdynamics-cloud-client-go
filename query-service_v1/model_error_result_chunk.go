/*
 * AppDynamics Cloud Query Service API
 *
 * An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package querysvc

// A signal that an error occurred during query execution.
type ErrorResultChunk struct {
	// Attribute identifying a response chunk as holding information of an error.
	Type_ string `json:"type,omitempty"`
	Error_ *ErrorResultChunkError `json:"error,omitempty"`
}