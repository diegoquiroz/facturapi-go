# OrganizationLegalInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Nombre comercial de la organización. | 
**LegalName** | **string** | Nombre Fiscal o Razón Social de la organización, *sin* el régimen societario (ej.: S.A. de C.V.).  | 
**TaxSystem** | **string** | Código del Régimen Fiscal, del [catálogo del SAT](#régimen-fiscal). | 
**Website** | Pointer to **string** | Sitio web de la organización, que aparecerá en el PDF y correos de facturas y recibos. | [optional] 
**SupportEmail** | Pointer to **string** | Dirección de correo electrónico para aclaraciones. Aparecerá en el PDF y correos de facturas y recibos. | [optional] 
**Phone** | Pointer to **string** | Teléfono de la organización, que aparecerá en el PDF y correos de facturas y recibos. | [optional] 
**Address** | [**OrganizationLegalInputAddress**](OrganizationLegalInputAddress.md) |  | 

## Methods

### NewOrganizationLegalInput

`func NewOrganizationLegalInput(name string, legalName string, taxSystem string, address OrganizationLegalInputAddress, ) *OrganizationLegalInput`

NewOrganizationLegalInput instantiates a new OrganizationLegalInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLegalInputWithDefaults

`func NewOrganizationLegalInputWithDefaults() *OrganizationLegalInput`

NewOrganizationLegalInputWithDefaults instantiates a new OrganizationLegalInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationLegalInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationLegalInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationLegalInput) SetName(v string)`

SetName sets Name field to given value.


### GetLegalName

`func (o *OrganizationLegalInput) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OrganizationLegalInput) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OrganizationLegalInput) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetTaxSystem

`func (o *OrganizationLegalInput) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *OrganizationLegalInput) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *OrganizationLegalInput) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.


### GetWebsite

`func (o *OrganizationLegalInput) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *OrganizationLegalInput) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *OrganizationLegalInput) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *OrganizationLegalInput) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetSupportEmail

`func (o *OrganizationLegalInput) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *OrganizationLegalInput) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *OrganizationLegalInput) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *OrganizationLegalInput) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetPhone

`func (o *OrganizationLegalInput) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationLegalInput) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationLegalInput) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationLegalInput) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *OrganizationLegalInput) GetAddress() OrganizationLegalInputAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationLegalInput) GetAddressOk() (*OrganizationLegalInputAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationLegalInput) SetAddress(v OrganizationLegalInputAddress)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


