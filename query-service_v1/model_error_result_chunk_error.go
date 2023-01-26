/*
 * AppDynamics Cloud Query Service API
 *
 * An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package querysvc
import (
	"time"
)

type ErrorResultChunkError struct {
	// URI to the documentation of the error.
	Type_ string `json:"type,omitempty"`
	// The code of the error.
	ErrorCode string `json:"errorCode,omitempty"`
	// Time and date the `ErrorResultChunk` was created.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// It correlates events and logs messages across dependent services for single business transaction.
	TraceId string `json:"traceId,omitempty"`
	// Brief description of the error.
	Title string `json:"title,omitempty"`
	// Usually more detailed description of the error.
	Detail string `json:"detail,omitempty"`
	// The input query.
	Query string `json:"query,omitempty"`
}