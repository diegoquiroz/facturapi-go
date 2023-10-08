# TaxIdValidationResultEfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** | Indica si el RFC tiene algún asunto relacionado con esta lista. &#x60;true&#x60;: El RFC no está en la lista de EFOS o su situación fue apelada y resultó favorable. &#x60;false&#x60;: El RFC está registrado como “Presunto” o “Definitivo” en la lista de EFOS.  | [optional] 
**Data** | Pointer to [**TaxIdValidationResultEfosData**](TaxIdValidationResultEfosData.md) |  | [optional] 

## Methods

### NewTaxIdValidationResultEfos

`func NewTaxIdValidationResultEfos() *TaxIdValidationResultEfos`

NewTaxIdValidationResultEfos instantiates a new TaxIdValidationResultEfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIdValidationResultEfosWithDefaults

`func NewTaxIdValidationResultEfosWithDefaults() *TaxIdValidationResultEfos`

NewTaxIdValidationResultEfosWithDefaults instantiates a new TaxIdValidationResultEfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *TaxIdValidationResultEfos) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *TaxIdValidationResultEfos) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *TaxIdValidationResultEfos) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *TaxIdValidationResultEfos) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetData

`func (o *TaxIdValidationResultEfos) GetData() TaxIdValidationResultEfosData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaxIdValidationResultEfos) GetDataOk() (*TaxIdValidationResultEfosData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaxIdValidationResultEfos) SetData(v TaxIdValidationResultEfosData)`

SetData sets Data field to given value.

### HasData

`func (o *TaxIdValidationResultEfos) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


