# CuustomerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto &#x60;customer&#x60; relacionado a la factura, en caso de no haber sido eliminado | [optional] 
**LegalName** | Pointer to **string** | Nombre Fiscal o Razón Social del cliente, *sin* incluir el régimen societario (ej.: S.A. de C.V.).  | [optional] 
**TaxId** | Pointer to **string** | RFC del cliente. | [optional] 
**Address** | Pointer to [**CuustomerInfoAddress**](CuustomerInfoAddress.md) |  | [optional] 

## Methods

### NewCuustomerInfo

`func NewCuustomerInfo() *CuustomerInfo`

NewCuustomerInfo instantiates a new CuustomerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCuustomerInfoWithDefaults

`func NewCuustomerInfoWithDefaults() *CuustomerInfo`

NewCuustomerInfoWithDefaults instantiates a new CuustomerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CuustomerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CuustomerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CuustomerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CuustomerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegalName

`func (o *CuustomerInfo) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CuustomerInfo) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CuustomerInfo) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CuustomerInfo) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxId

`func (o *CuustomerInfo) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CuustomerInfo) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CuustomerInfo) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CuustomerInfo) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetAddress

`func (o *CuustomerInfo) GetAddress() CuustomerInfoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CuustomerInfo) GetAddressOk() (*CuustomerInfoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CuustomerInfo) SetAddress(v CuustomerInfoAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CuustomerInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


