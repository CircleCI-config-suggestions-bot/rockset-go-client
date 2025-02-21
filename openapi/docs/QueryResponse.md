# QueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]string** | List of collections referenced in the query. | [optional] 
**ColumnFields** | Pointer to [**[]QueryFieldType**](QueryFieldType.md) | Meta information about each column in the result set. Not populated in &#x60;SELECT *&#x60; queries. | [optional] 
**LastOffset** | Pointer to **string** | If this was a write query, this is the log offset the query was written to. | [optional] 
**Pagination** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 
**QueryErrors** | Pointer to [**[]QueryError**](QueryError.md) | Errors encountered while executing the query. | [optional] 
**QueryId** | Pointer to **string** | Unique ID for this query. | [optional] 
**QueryLambdaPath** | Pointer to **string** | The full path of the executed query lambda. Includes version information. | [optional] 
**Results** | Pointer to **[]map[string]interface{}** | Results from the query. | [optional] 
**ResultsTotalDocCount** | Pointer to **int64** | Number of results generated by the query. | [optional] 
**Stats** | Pointer to [**QueryResponseStats**](QueryResponseStats.md) |  | [optional] 
**Warnings** | Pointer to **[]string** | Warnings generated by the query. Only populated if &#x60;generate_warnings&#x60; is specified in the query request. | [optional] 

## Methods

### NewQueryResponse

`func NewQueryResponse() *QueryResponse`

NewQueryResponse instantiates a new QueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryResponseWithDefaults

`func NewQueryResponseWithDefaults() *QueryResponse`

NewQueryResponseWithDefaults instantiates a new QueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *QueryResponse) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *QueryResponse) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *QueryResponse) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *QueryResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetColumnFields

`func (o *QueryResponse) GetColumnFields() []QueryFieldType`

GetColumnFields returns the ColumnFields field if non-nil, zero value otherwise.

### GetColumnFieldsOk

`func (o *QueryResponse) GetColumnFieldsOk() (*[]QueryFieldType, bool)`

GetColumnFieldsOk returns a tuple with the ColumnFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnFields

`func (o *QueryResponse) SetColumnFields(v []QueryFieldType)`

SetColumnFields sets ColumnFields field to given value.

### HasColumnFields

`func (o *QueryResponse) HasColumnFields() bool`

HasColumnFields returns a boolean if a field has been set.

### GetLastOffset

`func (o *QueryResponse) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *QueryResponse) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *QueryResponse) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *QueryResponse) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.

### GetPagination

`func (o *QueryResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *QueryResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *QueryResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *QueryResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetQueryErrors

`func (o *QueryResponse) GetQueryErrors() []QueryError`

GetQueryErrors returns the QueryErrors field if non-nil, zero value otherwise.

### GetQueryErrorsOk

`func (o *QueryResponse) GetQueryErrorsOk() (*[]QueryError, bool)`

GetQueryErrorsOk returns a tuple with the QueryErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryErrors

`func (o *QueryResponse) SetQueryErrors(v []QueryError)`

SetQueryErrors sets QueryErrors field to given value.

### HasQueryErrors

`func (o *QueryResponse) HasQueryErrors() bool`

HasQueryErrors returns a boolean if a field has been set.

### GetQueryId

`func (o *QueryResponse) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *QueryResponse) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *QueryResponse) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *QueryResponse) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetQueryLambdaPath

`func (o *QueryResponse) GetQueryLambdaPath() string`

GetQueryLambdaPath returns the QueryLambdaPath field if non-nil, zero value otherwise.

### GetQueryLambdaPathOk

`func (o *QueryResponse) GetQueryLambdaPathOk() (*string, bool)`

GetQueryLambdaPathOk returns a tuple with the QueryLambdaPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryLambdaPath

`func (o *QueryResponse) SetQueryLambdaPath(v string)`

SetQueryLambdaPath sets QueryLambdaPath field to given value.

### HasQueryLambdaPath

`func (o *QueryResponse) HasQueryLambdaPath() bool`

HasQueryLambdaPath returns a boolean if a field has been set.

### GetResults

`func (o *QueryResponse) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryResponse) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryResponse) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *QueryResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsTotalDocCount

`func (o *QueryResponse) GetResultsTotalDocCount() int64`

GetResultsTotalDocCount returns the ResultsTotalDocCount field if non-nil, zero value otherwise.

### GetResultsTotalDocCountOk

`func (o *QueryResponse) GetResultsTotalDocCountOk() (*int64, bool)`

GetResultsTotalDocCountOk returns a tuple with the ResultsTotalDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTotalDocCount

`func (o *QueryResponse) SetResultsTotalDocCount(v int64)`

SetResultsTotalDocCount sets ResultsTotalDocCount field to given value.

### HasResultsTotalDocCount

`func (o *QueryResponse) HasResultsTotalDocCount() bool`

HasResultsTotalDocCount returns a boolean if a field has been set.

### GetStats

`func (o *QueryResponse) GetStats() QueryResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *QueryResponse) GetStatsOk() (*QueryResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *QueryResponse) SetStats(v QueryResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *QueryResponse) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetWarnings

`func (o *QueryResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *QueryResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *QueryResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *QueryResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


