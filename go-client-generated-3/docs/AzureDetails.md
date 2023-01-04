# AzureDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client IDs, also known as Application IDs, are long-term credentials for an Azure user, or account root user. The Client ID is one of three properties needed to authenticate to Azure, the other two being Client Secret and Tenant (Directory) ID | [default to null]
**ClientSecret** | **string** | A Client Secret allows an Azure application to provide its identity when requesting an access token. The Client Secret is one of three properties needed to authenticate to Azure, the other two being Client ID (Application ID) and Tenant (Directory) ID | [default to null]
**TenantId** | **string** | The Azure AD Tenant (Directory) IDis one of three properties needed to authenticate to Azure. The other two are Client Secret and Client ID (Application ID). | [default to null]
**SubscriptionId** | **string** | Specify a GUID Subscription ID to monitor. If monitoring all subscriptions, do not specify a Subscription ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

