# OrganizationSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Número de página de resultados | [optional] 
**TotalPages** | Pointer to **int32** | Número total de páginas de resultados | [optional] 
**TotalResults** | Pointer to **int32** | Número de elementos individuales en todas las páginas de resultados | [optional] 
**Data** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 

## Methods

### NewOrganizationSearchResult

`func NewOrganizationSearchResult() *OrganizationSearchResult`

NewOrganizationSearchResult instantiates a new OrganizationSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSearchResultWithDefaults

`func NewOrganizationSearchResultWithDefaults() *OrganizationSearchResult`

NewOrganizationSearchResultWithDefaults instantiates a new OrganizationSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *OrganizationSearchResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrganizationSearchResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrganizationSearchResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *OrganizationSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *OrganizationSearchResult) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *OrganizationSearchResult) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *OrganizationSearchResult) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *OrganizationSearchResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *OrganizationSearchResult) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *OrganizationSearchResult) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *OrganizationSearchResult) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *OrganizationSearchResult) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetData

`func (o *OrganizationSearchResult) GetData() []Organization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationSearchResult) GetDataOk() (*[]Organization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationSearchResult) SetData(v []Organization)`

SetData sets Data field to given value.

### HasData

`func (o *OrganizationSearchResult) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


