# Go API client for querysvc

An API providing access to observation data using an AppDynamics domain-specific language.  The language is declarative, read-only (it does not allow for data modification) and specific to the AppD MELT model. It presents MELT data (metrics, events, logs, traces) and State in the scope of monitored topology.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./querysvc"
```

## Documentation for API Endpoints

All URIs are relative to *https://{tenant-name}.observe.appdynamics.com/monitoring/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ExecuteQueryApi* | [**ExecuteQuery**](docs/ExecuteQueryApi.md#executequery) | **Post** /query/execute | Execute a query to fetch observation data.
*ResultPaginationApi* | [**ResultPagination**](docs/ResultPaginationApi.md#resultpagination) | **Get** /query/continue | Fetch the next page of results.

## Documentation For Models

 - [AnyOfDataResultItemItems](docs/AnyOfDataResultItemItems.md)
 - [AnyOfErrorDetailItemFixPossibilitiesItems](docs/AnyOfErrorDetailItemFixPossibilitiesItems.md)
 - [AnyOfErrorDetailsArrayItems](docs/AnyOfErrorDetailsArrayItems.md)
 - [AnyOfModelFieldsItems](docs/AnyOfModelFieldsItems.md)
 - [AnyOfQueryResponseArrayBodyItems](docs/AnyOfQueryResponseArrayBodyItems.md)
 - [ComplexModelResultItem](docs/ComplexModelResultItem.md)
 - [DataResultChunk](docs/DataResultChunk.md)
 - [DatasetReference](docs/DatasetReference.md)
 - [ErrorDetailItem](docs/ErrorDetailItem.md)
 - [ErrorResult](docs/ErrorResult.md)
 - [ErrorResultChunk](docs/ErrorResultChunk.md)
 - [ErrorResultChunkError](docs/ErrorResultChunkError.md)
 - [EventMetadataResultItem](docs/EventMetadataResultItem.md)
 - [HeartbeatResultChunk](docs/HeartbeatResultChunk.md)
 - [Hints](docs/Hints.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [MainMetadataResultItem](docs/MainMetadataResultItem.md)
 - [MetadataResultItem](docs/MetadataResultItem.md)
 - [MetricMetadataResultItem](docs/MetricMetadataResultItem.md)
 - [Model](docs/Model.md)
 - [ModelReference](docs/ModelReference.md)
 - [ModelResultChunk](docs/ModelResultChunk.md)
 - [ModelResultItem](docs/ModelResultItem.md)
 - [PaginationReference](docs/PaginationReference.md)
 - [PaginationReferenceNext](docs/PaginationReferenceNext.md)
 - [QueryRequestBody](docs/QueryRequestBody.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


