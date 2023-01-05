# {{classname}}

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfiguration**](ConfigurationsApi.md#CreateConfiguration) | **Post** /configurations | Create a configuration
[**DeleteConfiguration**](ConfigurationsApi.md#DeleteConfiguration) | **Delete** /configurations/{configurationId} | Delete configuration
[**GetConfiguration**](ConfigurationsApi.md#GetConfiguration) | **Get** /configurations/{configurationId} | Get a configuration
[**GetConfigurations**](ConfigurationsApi.md#GetConfigurations) | **Get** /configurations | Query configurations
[**UpdateConfiguration**](ConfigurationsApi.md#UpdateConfiguration) | **Patch** /configurations/{configurationId} | Update a configuration

# **CreateConfiguration**
> ConfigurationDetail CreateConfiguration(ctx, body)
Create a configuration

Create a configuration for a specific provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Configuration**](Configuration.md)|  | 

### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfiguration**
> DeleteConfiguration(ctx, configurationId)
Delete configuration

Delete a configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurationId** | [**string**](.md)| Configuration ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfiguration**
> ConfigurationDetail GetConfiguration(ctx, configurationId)
Get a configuration

Get a configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurationId** | [**string**](.md)| Configuration ID | 

### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigurations**
> Configurations GetConfigurations(ctx, optional)
Query configurations

Get all configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationsApiGetConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationsApiGetConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter expression | 
 **sort** | **optional.String**| Sort by | 
 **order** | **optional.String**| Order by | 
 **max** | **optional.Int32**| Maximum number of results to return | [default to 10]
 **cursor** | **optional.String**| Cursor for the paginated requests | 

### Return type

[**Configurations**](Configurations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConfiguration**
> ConfigurationDetail UpdateConfiguration(ctx, body, configurationId)
Update a configuration

Update a configuration by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigurationUpdate**](ConfigurationUpdate.md)|  | 
  **configurationId** | [**string**](.md)| Configuration ID | 

### Return type

[**ConfigurationDetail**](ConfigurationDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

