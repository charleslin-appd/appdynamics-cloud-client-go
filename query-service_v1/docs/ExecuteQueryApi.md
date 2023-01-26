# {{classname}}

All URIs are relative to *https://{tenant-name}.observe.appdynamics.com/monitoring/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteQuery**](ExecuteQueryApi.md#ExecuteQuery) | **Post** /query/execute | Execute a query to fetch observation data.

# **ExecuteQuery**
> []AnyOfQueryResponseArrayBodyItems ExecuteQuery(ctx, optional)
Execute a query to fetch observation data.

Execute a query to fetch observation data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecuteQueryApiExecuteQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecuteQueryApiExecuteQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QueryRequestBody**](QueryRequestBody.md)| A simple JSON object containing a query string. | 

### Return type

[**[]AnyOfQueryResponseArrayBodyItems**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/x-ndjson, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

