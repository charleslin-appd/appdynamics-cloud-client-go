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

type RoleDelegationConnectionResponseDetails struct {
	RoleName string `json:"roleName"`
	AccountId string `json:"accountId"`
	// AppDynamics AWS Account ID. Delegates a user to an Identity Access Management (IAM) role in AWS. The AWS IAM role provides AppDynamics access to resources.
	AppDynamicsAwsAccountId string `json:"appDynamicsAwsAccountId,omitempty"`
	// Returns an external ID for AWS role delegation connections
	ExternalId string `json:"externalId,omitempty"`
}
