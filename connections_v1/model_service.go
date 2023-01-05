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

type Service struct {
	Name string `json:"name"`
	ImportTags *ImportTagConfiguration `json:"importTags,omitempty"`
	TagFilter string `json:"tagFilter,omitempty"`
	Polling *ScheduleInterval `json:"polling,omitempty"`
}