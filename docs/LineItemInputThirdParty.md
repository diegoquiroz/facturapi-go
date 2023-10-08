# LineItemInputThirdParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | **string** | Nombre o razón social del tercero. | 
**TaxId** | **string** | RFC del tercero. | 
**TaxSystem** | **string** | Régimen fiscal del tercero. | 
**Zip** | **string** | Código postal del tercero. | 

## Methods

### NewLineItemInputThirdParty

`func NewLineItemInputThirdParty(legalName string, taxId string, taxSystem string, zip string, ) *LineItemInputThirdParty`

NewLineItemInputThirdParty instantiates a new LineItemInputThirdParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemInputThirdPartyWithDefaults

`func NewLineItemInputThirdPartyWithDefaults() *LineItemInputThirdParty`

NewLineItemInputThirdPartyWithDefaults instantiates a new LineItemInputThirdParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *LineItemInputThirdParty) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *LineItemInputThirdParty) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *LineItemInputThirdParty) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetTaxId

`func (o *LineItemInputThirdParty) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *LineItemInputThirdParty) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *LineItemInputThirdParty) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.


### GetTaxSystem

`func (o *LineItemInputThirdParty) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *LineItemInputThirdParty) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *LineItemInputThirdParty) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.


### GetZip

`func (o *LineItemInputThirdParty) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *LineItemInputThirdParty) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *LineItemInputThirdParty) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


