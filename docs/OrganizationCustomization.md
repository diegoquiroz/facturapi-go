# OrganizationCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasLogo** | Pointer to **bool** | Indica si la organización ya tiene un logotipo cargado. | [optional] 
**Color** | Pointer to **string** | Color distintivo de la marca en representación Hexadecimal RGB de 6 caracteres. | [optional] 
**NextFolioNumber** | Pointer to **int32** | Número de folio que se asignará a la siguiente factura en ambiente Live (y que se incrementará automáticamente por cada nueva factura). | [optional] 
**NextFolioNumberTest** | Pointer to **int32** | Número de folio que se asignará a la siguiente factura en ambiente Test (y que se incrementará automáticamente por cada nueva factura). | [optional] 
**PdfExtra** | Pointer to [**OrganizationCustomizationPdfExtra**](OrganizationCustomizationPdfExtra.md) |  | [optional] 

## Methods

### NewOrganizationCustomization

`func NewOrganizationCustomization() *OrganizationCustomization`

NewOrganizationCustomization instantiates a new OrganizationCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomizationWithDefaults

`func NewOrganizationCustomizationWithDefaults() *OrganizationCustomization`

NewOrganizationCustomizationWithDefaults instantiates a new OrganizationCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasLogo

`func (o *OrganizationCustomization) GetHasLogo() bool`

GetHasLogo returns the HasLogo field if non-nil, zero value otherwise.

### GetHasLogoOk

`func (o *OrganizationCustomization) GetHasLogoOk() (*bool, bool)`

GetHasLogoOk returns a tuple with the HasLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLogo

`func (o *OrganizationCustomization) SetHasLogo(v bool)`

SetHasLogo sets HasLogo field to given value.

### HasHasLogo

`func (o *OrganizationCustomization) HasHasLogo() bool`

HasHasLogo returns a boolean if a field has been set.

### GetColor

`func (o *OrganizationCustomization) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *OrganizationCustomization) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *OrganizationCustomization) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *OrganizationCustomization) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetNextFolioNumber

`func (o *OrganizationCustomization) GetNextFolioNumber() int32`

GetNextFolioNumber returns the NextFolioNumber field if non-nil, zero value otherwise.

### GetNextFolioNumberOk

`func (o *OrganizationCustomization) GetNextFolioNumberOk() (*int32, bool)`

GetNextFolioNumberOk returns a tuple with the NextFolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFolioNumber

`func (o *OrganizationCustomization) SetNextFolioNumber(v int32)`

SetNextFolioNumber sets NextFolioNumber field to given value.

### HasNextFolioNumber

`func (o *OrganizationCustomization) HasNextFolioNumber() bool`

HasNextFolioNumber returns a boolean if a field has been set.

### GetNextFolioNumberTest

`func (o *OrganizationCustomization) GetNextFolioNumberTest() int32`

GetNextFolioNumberTest returns the NextFolioNumberTest field if non-nil, zero value otherwise.

### GetNextFolioNumberTestOk

`func (o *OrganizationCustomization) GetNextFolioNumberTestOk() (*int32, bool)`

GetNextFolioNumberTestOk returns a tuple with the NextFolioNumberTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFolioNumberTest

`func (o *OrganizationCustomization) SetNextFolioNumberTest(v int32)`

SetNextFolioNumberTest sets NextFolioNumberTest field to given value.

### HasNextFolioNumberTest

`func (o *OrganizationCustomization) HasNextFolioNumberTest() bool`

HasNextFolioNumberTest returns a boolean if a field has been set.

### GetPdfExtra

`func (o *OrganizationCustomization) GetPdfExtra() OrganizationCustomizationPdfExtra`

GetPdfExtra returns the PdfExtra field if non-nil, zero value otherwise.

### GetPdfExtraOk

`func (o *OrganizationCustomization) GetPdfExtraOk() (*OrganizationCustomizationPdfExtra, bool)`

GetPdfExtraOk returns a tuple with the PdfExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfExtra

`func (o *OrganizationCustomization) SetPdfExtra(v OrganizationCustomizationPdfExtra)`

SetPdfExtra sets PdfExtra field to given value.

### HasPdfExtra

`func (o *OrganizationCustomization) HasPdfExtra() bool`

HasPdfExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


