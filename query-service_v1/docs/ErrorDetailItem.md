# ErrorDetailItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A short description on error cause. | [default to null]
**FixSuggestion** | **string** | A hint to resolve the error. | [optional] [default to null]
**FixPossibilities** | [**[]AnyOfErrorDetailItemFixPossibilitiesItems**](.md) | A list of fix possibilities to resolve the error. | [default to null]
**ErrorType** | **string** | The type of the error. | [default to null]
**ErrorFrom** | **string** | The start position of the error in format &#x27;lineNum:charIdx&#x27;. | [default to null]
**ErrorTo** | **string** | The end position of the error in format &#x27;lineNum:charIdx&#x27;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

