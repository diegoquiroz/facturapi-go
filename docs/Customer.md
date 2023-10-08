# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha de registro | [optional] 
**Livemode** | Pointer to **bool** | Si el valor es &#x60;true&#x60;, indica que el objeto fue creado en ambiente Live; o si es &#x60;false&#x60;, en ambiente Test. | [optional] 
**LegalName** | Pointer to **string** | Nombre Fiscal o Razón Social del cliente. *sin* el régimen societario (ej.: S.A. de C.V.).  | [optional] 
**TaxId** | Pointer to **string** | En clientes de México contiene el RFC del cliente. Para extranjeros es opcional y representa el número de registro de identificacón tributaria, es decir, el equivalente al RFC en el país del cliente. | [optional] 
**TaxSystem** | Pointer to **string** | Requerido para clientes nacionales. Clave del régimen fiscal del cliente. | [optional] 
**Email** | Pointer to **string** | Dirección de correo electrónico al cual enviar las facturas generadas. | [optional] 
**Phone** | Pointer to **string** | Teléfono del cliente. | [optional] 
**Address** | Pointer to [**CustomerPropertiesAllOfAddress**](CustomerPropertiesAllOfAddress.md) |  | [optional] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Customer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Customer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Customer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Customer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLivemode

`func (o *Customer) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Customer) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Customer) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Customer) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetLegalName

`func (o *Customer) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *Customer) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *Customer) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *Customer) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxId

`func (o *Customer) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Customer) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Customer) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Customer) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxSystem

`func (o *Customer) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *Customer) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *Customer) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.

### HasTaxSystem

`func (o *Customer) HasTaxSystem() bool`

HasTaxSystem returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Customer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *Customer) GetAddress() CustomerPropertiesAllOfAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Customer) GetAddressOk() (*CustomerPropertiesAllOfAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Customer) SetAddress(v CustomerPropertiesAllOfAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Customer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


