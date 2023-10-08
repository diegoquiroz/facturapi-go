# CustomerProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | Pointer to **string** | Nombre Fiscal o Razón Social del cliente. *sin* el régimen societario (ej.: S.A. de C.V.).  | [optional] 
**TaxId** | Pointer to **string** | En clientes de México contiene el RFC del cliente. Para extranjeros es opcional y representa el número de registro de identificacón tributaria, es decir, el equivalente al RFC en el país del cliente. | [optional] 
**TaxSystem** | Pointer to **string** | Requerido para clientes nacionales. Clave del régimen fiscal del cliente. | [optional] 
**Email** | Pointer to **string** | Dirección de correo electrónico al cual enviar las facturas generadas. | [optional] 
**Phone** | Pointer to **string** | Teléfono del cliente. | [optional] 
**Address** | Pointer to [**CustomerPropertiesAllOfAddress**](CustomerPropertiesAllOfAddress.md) |  | [optional] 

## Methods

### NewCustomerProperties

`func NewCustomerProperties() *CustomerProperties`

NewCustomerProperties instantiates a new CustomerProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPropertiesWithDefaults

`func NewCustomerPropertiesWithDefaults() *CustomerProperties`

NewCustomerPropertiesWithDefaults instantiates a new CustomerProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *CustomerProperties) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CustomerProperties) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CustomerProperties) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CustomerProperties) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxId

`func (o *CustomerProperties) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CustomerProperties) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CustomerProperties) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CustomerProperties) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxSystem

`func (o *CustomerProperties) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *CustomerProperties) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *CustomerProperties) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.

### HasTaxSystem

`func (o *CustomerProperties) HasTaxSystem() bool`

HasTaxSystem returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerProperties) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerProperties) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerProperties) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerProperties) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerProperties) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerProperties) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerProperties) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *CustomerProperties) GetAddress() CustomerPropertiesAllOfAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CustomerProperties) GetAddressOk() (*CustomerPropertiesAllOfAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CustomerProperties) SetAddress(v CustomerPropertiesAllOfAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CustomerProperties) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


