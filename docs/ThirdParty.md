# ThirdParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | Pointer to **string** | Nombre o razón social del tercero. | [optional] 
**TaxId** | Pointer to **string** | RFC del tercero. | [optional] 
**TaxSystem** | Pointer to **string** | Régimen fiscal del tercero. | [optional] 
**Zip** | Pointer to **string** | Código postal del tercero. | [optional] 

## Methods

### NewThirdParty

`func NewThirdParty() *ThirdParty`

NewThirdParty instantiates a new ThirdParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyWithDefaults

`func NewThirdPartyWithDefaults() *ThirdParty`

NewThirdPartyWithDefaults instantiates a new ThirdParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *ThirdParty) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *ThirdParty) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *ThirdParty) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *ThirdParty) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxId

`func (o *ThirdParty) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ThirdParty) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ThirdParty) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ThirdParty) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxSystem

`func (o *ThirdParty) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *ThirdParty) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *ThirdParty) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.

### HasTaxSystem

`func (o *ThirdParty) HasTaxSystem() bool`

HasTaxSystem returns a boolean if a field has been set.

### GetZip

`func (o *ThirdParty) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ThirdParty) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ThirdParty) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ThirdParty) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


