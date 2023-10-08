# UnitCatalogResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Clave del catálogo | [optional] 
**Description** | Pointer to **string** | Descripción | [optional] 
**Score** | Pointer to **float32** | Número del 0 al 1 que representa el nivel de coincidencia del resultado con respecto a la consulta de búsqueda.  | [optional] 

## Methods

### NewUnitCatalogResult

`func NewUnitCatalogResult() *UnitCatalogResult`

NewUnitCatalogResult instantiates a new UnitCatalogResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitCatalogResultWithDefaults

`func NewUnitCatalogResultWithDefaults() *UnitCatalogResult`

NewUnitCatalogResultWithDefaults instantiates a new UnitCatalogResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *UnitCatalogResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UnitCatalogResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UnitCatalogResult) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UnitCatalogResult) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *UnitCatalogResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnitCatalogResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnitCatalogResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnitCatalogResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScore

`func (o *UnitCatalogResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UnitCatalogResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UnitCatalogResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *UnitCatalogResult) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


