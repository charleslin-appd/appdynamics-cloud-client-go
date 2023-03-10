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

// A representation of an error that occurred during query execution. It contains at least brief a description of the error and contextual information of the corresponding request.
type ErrorResult struct {
	// URI to the documentation of the error.
	Type_ string `json:"type"`
	// The code of a error.
	ErrorCode string `json:"errorCode,omitempty"`
	// Time and date the error occurred.
	Timestamp time.Time `json:"timestamp"`
	// Id to correlate events and logs across dependent services for a single query.
	TraceId string `json:"traceId"`
	// Brief description of the error.
	Title string `json:"title"`
	// Usually a more detailed description of the error.
	Detail string `json:"detail"`
	// The id of a tenant for which the query was executed.
	TenantId string `json:"tenantId"`
	// The input query.
	Query string `json:"query"`
	ErrorDetails *[]AnyOfErrorDetailsArrayItems `json:"errorDetails,omitempty"`
}
