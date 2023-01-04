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

type BaseConfigurationDetails struct {
	// Geographic locations used to fetch metrics
	Regions []string `json:"regions,omitempty"`
	Polling *ScheduleInterval `json:"polling,omitempty"`
	ImportTags *ImportTagConfiguration `json:"importTags,omitempty"`
	TagFilter string `json:"tagFilter,omitempty"`
	// services for which we will fetch metrics
	Services []Service `json:"services,omitempty"`
}