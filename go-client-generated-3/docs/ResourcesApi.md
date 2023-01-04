# {{classname}}

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRegions**](ResourcesApi.md#GetRegions) | **Get** /regions | Get all supported hosting regions
[**GetServices**](ResourcesApi.md#GetServices) | **Get** /services | Get all supported hosting services

# **GetRegions**
> ConfigurationResourceList GetRegions(ctx, type_)
Get all supported hosting regions

Get all supported hosting regions for a given cloud provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | [**ProviderType**](.md)| Cloud hosting provider type | 

### Return type

[**ConfigurationResourceList**](ConfigurationResourceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServices**
> ConfigurationResourceList GetServices(ctx, type_)
Get all supported hosting services

Get all supported hosting services for a given provider type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | [**ProviderType**](.md)| Hosting provider type | 

### Return type

[**ConfigurationResourceList**](ConfigurationResourceList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

