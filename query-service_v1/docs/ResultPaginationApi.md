# {{classname}}

All URIs are relative to *https://{tenant-name}.observe.appdynamics.com/monitoring/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResultPagination**](ResultPaginationApi.md#ResultPagination) | **Get** /query/continue | Fetch the next page of results.

# **ResultPagination**
> []AnyOfQueryResponseArrayBodyItems ResultPagination(ctx, cursor)
Fetch the next page of results.

Fetch the next page of results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cursor** | **string**| An opaque string which will allow the retrieval of the next page of results. The cursor is obtained from a previous query response. | 

### Return type

[**[]AnyOfQueryResponseArrayBodyItems**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/x-ndjson, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

