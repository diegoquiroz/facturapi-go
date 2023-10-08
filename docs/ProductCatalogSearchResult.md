# ProductCatalogSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Número de página de resultados | [optional] 
**TotalPages** | Pointer to **int32** | Número total de páginas de resultados | [optional] 
**TotalResults** | Pointer to **int32** | Número de elementos individuales en todas las páginas de resultados | [optional] 
**Data** | Pointer to [**[]ProductCatalogResult**](ProductCatalogResult.md) |  | [optional] 

## Methods

### NewProductCatalogSearchResult

`func NewProductCatalogSearchResult() *ProductCatalogSearchResult`

NewProductCatalogSearchResult instantiates a new ProductCatalogSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCatalogSearchResultWithDefaults

`func NewProductCatalogSearchResultWithDefaults() *ProductCatalogSearchResult`

NewProductCatalogSearchResultWithDefaults instantiates a new ProductCatalogSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ProductCatalogSearchResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProductCatalogSearchResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProductCatalogSearchResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProductCatalogSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *ProductCatalogSearchResult) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ProductCatalogSearchResult) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ProductCatalogSearchResult) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ProductCatalogSearchResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *ProductCatalogSearchResult) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ProductCatalogSearchResult) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ProductCatalogSearchResult) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *ProductCatalogSearchResult) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetData

`func (o *ProductCatalogSearchResult) GetData() []ProductCatalogResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductCatalogSearchResult) GetDataOk() (*[]ProductCatalogResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductCatalogSearchResult) SetData(v []ProductCatalogResult)`

SetData sets Data field to given value.

### HasData

`func (o *ProductCatalogSearchResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


