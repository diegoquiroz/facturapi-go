# CustomerCommonProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | Pointer to **string** | Nombre Fiscal o Razón Social del cliente. *sin* el régimen societario (ej.: S.A. de C.V.).  | [optional] 
**TaxId** | Pointer to **string** | En clientes de México contiene el RFC del cliente. Para extranjeros es opcional y representa el número de registro de identificacón tributaria, es decir, el equivalente al RFC en el país del cliente. | [optional] 
**TaxSystem** | Pointer to **string** | Requerido para clientes nacionales. Clave del régimen fiscal del cliente. | [optional] 
**Email** | Pointer to **string** | Dirección de correo electrónico al cual enviar las facturas generadas. | [optional] 
**Phone** | Pointer to **string** | Teléfono del cliente. | [optional] 

## Methods

### NewCustomerCommonProperties

`func NewCustomerCommonProperties() *CustomerCommonProperties`

NewCustomerCommonProperties instantiates a new CustomerCommonProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCommonPropertiesWithDefaults

`func NewCustomerCommonPropertiesWithDefaults() *CustomerCommonProperties`

NewCustomerCommonPropertiesWithDefaults instantiates a new CustomerCommonProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *CustomerCommonProperties) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CustomerCommonProperties) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CustomerCommonProperties) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CustomerCommonProperties) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxId

`func (o *CustomerCommonProperties) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerCommonProperties) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerCommonProperties) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerCommonProperties) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxSystem

`func (o *CustomerCommonProperties) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *CustomerCommonProperties) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *CustomerCommonProperties) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.

### HasTaxSystem

`func (o *CustomerCommonProperties) HasTaxSystem() bool`

HasTaxSystem returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerCommonProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerCommonProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerCommonProperties) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerCommonProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerCommonProperties) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerCommonProperties) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerCommonProperties) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerCommonProperties) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


