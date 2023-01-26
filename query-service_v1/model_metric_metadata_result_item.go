/*
 * AppDynamics Cloud Query Service API
 *
 * An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package querysvc

// Metadata for a metric data chunk.
type MetricMetadataResultItem struct {
	// Actual granularity of metric data points. The value is the length of the time interval of a single data point.
	GranularitySeconds float64 `json:"granularitySeconds,omitempty"`
}