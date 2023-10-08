# OrganizationLegal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Nombre comercial de la organización. | [optional] 
**LegalName** | Pointer to **string** | Nombre Fiscal o Razón Social de la organización, *sin* el régimen societario (ej.: S.A. de C.V.).  | [optional] 
**TaxSystem** | Pointer to **string** | Código de Régimen Fiscal, del [catálogo del SAT](#tipo-de-régimen). | [optional] 
**Website** | Pointer to **string** | Sitio web de la organización, que se utilizará al enviar la factura por correo electrónico. | [optional] 
**Phone** | Pointer to **string** | Teléfono de la organización, que aparecerá en el PDF de la factura. | [optional] 
**Address** | Pointer to [**OrganizationLegalAddress**](OrganizationLegalAddress.md) |  | [optional] 

## Methods

### NewOrganizationLegal

`func NewOrganizationLegal() *OrganizationLegal`

NewOrganizationLegal instantiates a new OrganizationLegal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLegalWithDefaults

`func NewOrganizationLegalWithDefaults() *OrganizationLegal`

NewOrganizationLegalWithDefaults instantiates a new OrganizationLegal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationLegal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationLegal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationLegal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationLegal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLegalName

`func (o *OrganizationLegal) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *OrganizationLegal) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *OrganizationLegal) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *OrganizationLegal) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetTaxSystem

`func (o *OrganizationLegal) GetTaxSystem() string`

GetTaxSystem returns the TaxSystem field if non-nil, zero value otherwise.

### GetTaxSystemOk

`func (o *OrganizationLegal) GetTaxSystemOk() (*string, bool)`

GetTaxSystemOk returns a tuple with the TaxSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSystem

`func (o *OrganizationLegal) SetTaxSystem(v string)`

SetTaxSystem sets TaxSystem field to given value.

### HasTaxSystem

`func (o *OrganizationLegal) HasTaxSystem() bool`

HasTaxSystem returns a boolean if a field has been set.

### GetWebsite

`func (o *OrganizationLegal) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *OrganizationLegal) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *OrganizationLegal) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *OrganizationLegal) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetPhone

`func (o *OrganizationLegal) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *OrganizationLegal) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *OrganizationLegal) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *OrganizationLegal) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *OrganizationLegal) GetAddress() OrganizationLegalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrganizationLegal) GetAddressOk() (*OrganizationLegalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrganizationLegal) SetAddress(v OrganizationLegalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrganizationLegal) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


