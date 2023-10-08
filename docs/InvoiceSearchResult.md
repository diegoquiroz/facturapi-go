# InvoiceSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Número de página de resultados | [optional] 
**TotalPages** | Pointer to **int32** | Número total de páginas de resultados | [optional] 
**TotalResults** | Pointer to **int32** | Número de elementos individuales en todas las páginas de resultados | [optional] 
**Data** | Pointer to [**[]Invoice**](Invoice.md) |  | [optional] 

## Methods

### NewInvoiceSearchResult

`func NewInvoiceSearchResult() *InvoiceSearchResult`

NewInvoiceSearchResult instantiates a new InvoiceSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSearchResultWithDefaults

`func NewInvoiceSearchResultWithDefaults() *InvoiceSearchResult`

NewInvoiceSearchResultWithDefaults instantiates a new InvoiceSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *InvoiceSearchResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InvoiceSearchResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InvoiceSearchResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *InvoiceSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *InvoiceSearchResult) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *InvoiceSearchResult) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *InvoiceSearchResult) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *InvoiceSearchResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *InvoiceSearchResult) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *InvoiceSearchResult) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *InvoiceSearchResult) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *InvoiceSearchResult) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetData

`func (o *InvoiceSearchResult) GetData() []Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoiceSearchResult) GetDataOk() (*[]Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoiceSearchResult) SetData(v []Invoice)`

SetData sets Data field to given value.

### HasData

`func (o *InvoiceSearchResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


