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

type AwsAccessKeyDetails struct {
	AccessType *ConnectionAccessType `json:"accessType"`
	// AWS Access keys are long-term credentials for an AWS IAM user, or account root user. The access key ID is one of two access keys needed to authenticate to AWS. The other is a secret access key. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.
	AccessKeyId string `json:"accessKeyId"`
	// The secret access key is one of two access keys needed to authenticate to AWS. The other is an access key ID. The secret access key is only available once, when you create it. Download the generated secret access key and save in a secure location. If the secret access key is lost or deleted, you must create a new one. You need access keys to make programmatic calls using the AWS CLI, AWS Tools, or PowerShell.
	SecretAccessKey string `json:"secretAccessKey"`
}