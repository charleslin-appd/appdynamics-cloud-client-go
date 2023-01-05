# {{classname}}

All URIs are relative to *http://cloudmonconnectionservice.cloudmon.svc.cluster.local:7778/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionsApi.md#CreateConnection) | **Post** /connections | Create a connection
[**DeleteConnection**](ConnectionsApi.md#DeleteConnection) | **Delete** /connections/{connectionId} | Delete connection
[**GetConnection**](ConnectionsApi.md#GetConnection) | **Get** /connections/{connectionId} | Get a connection
[**GetConnections**](ConnectionsApi.md#GetConnections) | **Get** /connections | Query connections
[**UpdateConnection**](ConnectionsApi.md#UpdateConnection) | **Patch** /connections/{connectionId} | Update a connection

# **CreateConnection**
> ConnectionResponse CreateConnection(ctx, body)
Create a connection

Create a connection for a specific provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionRequest**](ConnectionRequest.md)|  | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnection**
> DeleteConnection(ctx, connectionId)
Delete connection

Delete a connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnection**
> ConnectionResponse GetConnection(ctx, connectionId)
Get a connection

Get an individual connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection ID | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnections**
> Connections GetConnections(ctx, optional)
Query connections

GET all Connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConnectionsApiGetConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectionsApiGetConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter expression | 
 **sort** | **optional.String**| Sort by | 
 **order** | **optional.String**| Order by | 
 **max** | **optional.Int32**| Maximum number of results to return | [default to 10]
 **cursor** | **optional.String**| Cursor for the paginated requests | 

### Return type

[**Connections**](Connections.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnection**
> ConnectionResponse UpdateConnection(ctx, body, connectionId)
Update a connection

Update a connection by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionUpdate**](ConnectionUpdate.md)|  | 
  **connectionId** | [**string**](.md)| Connection ID | 

### Return type

[**ConnectionResponse**](ConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

