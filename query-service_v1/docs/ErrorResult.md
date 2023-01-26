# ErrorResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | URI to the documentation of the error. | [default to null]
**ErrorCode** | **string** | The code of a error. | [optional] [default to null]
**Timestamp** | [**time.Time**](time.Time.md) | Time and date the error occurred. | [default to null]
**TraceId** | **string** | Id to correlate events and logs across dependent services for a single query. | [default to null]
**Title** | **string** | Brief description of the error. | [default to null]
**Detail** | **string** | Usually a more detailed description of the error. | [default to null]
**TenantId** | **string** | The id of a tenant for which the query was executed. | [default to null]
**Query** | **string** | The input query. | [default to null]
**ErrorDetails** | [***[]AnyOfErrorDetailsArrayItems**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

