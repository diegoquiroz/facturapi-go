# OrganizationCustomizationPdfExtra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codes** | Pointer to **bool** | Mostrar códigos de catálogos del SAT junto a sus descripciones. Ejemplo: “KGM Kilogramo”.  | [optional] [default to true]
**ProductKey** | Pointer to **bool** | Mostrar la clave de producto-servicio.  | [optional] [default to true]
**RoundUnitPrice** | Pointer to **bool** | Redondear el precio unitario en el PDF a 2 decimales, pero conservar los 6 decimales en el XML.  | [optional] [default to false]
**TaxBreakdown** | Pointer to **bool** | Mostrar el desglose de impuestos en el PDF. Si se desactiva, sólo se mostratán los impuestos en los totales, pero no en el detalle de cada concepto.  | [optional] [default to true]

## Methods

### NewOrganizationCustomizationPdfExtra

`func NewOrganizationCustomizationPdfExtra() *OrganizationCustomizationPdfExtra`

NewOrganizationCustomizationPdfExtra instantiates a new OrganizationCustomizationPdfExtra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCustomizationPdfExtraWithDefaults

`func NewOrganizationCustomizationPdfExtraWithDefaults() *OrganizationCustomizationPdfExtra`

NewOrganizationCustomizationPdfExtraWithDefaults instantiates a new OrganizationCustomizationPdfExtra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodes

`func (o *OrganizationCustomizationPdfExtra) GetCodes() bool`

GetCodes returns the Codes field if non-nil, zero value otherwise.

### GetCodesOk

`func (o *OrganizationCustomizationPdfExtra) GetCodesOk() (*bool, bool)`

GetCodesOk returns a tuple with the Codes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodes

`func (o *OrganizationCustomizationPdfExtra) SetCodes(v bool)`

SetCodes sets Codes field to given value.

### HasCodes

`func (o *OrganizationCustomizationPdfExtra) HasCodes() bool`

HasCodes returns a boolean if a field has been set.

### GetProductKey

`func (o *OrganizationCustomizationPdfExtra) GetProductKey() bool`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *OrganizationCustomizationPdfExtra) GetProductKeyOk() (*bool, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *OrganizationCustomizationPdfExtra) SetProductKey(v bool)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *OrganizationCustomizationPdfExtra) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetRoundUnitPrice

`func (o *OrganizationCustomizationPdfExtra) GetRoundUnitPrice() bool`

GetRoundUnitPrice returns the RoundUnitPrice field if non-nil, zero value otherwise.

### GetRoundUnitPriceOk

`func (o *OrganizationCustomizationPdfExtra) GetRoundUnitPriceOk() (*bool, bool)`

GetRoundUnitPriceOk returns a tuple with the RoundUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundUnitPrice

`func (o *OrganizationCustomizationPdfExtra) SetRoundUnitPrice(v bool)`

SetRoundUnitPrice sets RoundUnitPrice field to given value.

### HasRoundUnitPrice

`func (o *OrganizationCustomizationPdfExtra) HasRoundUnitPrice() bool`

HasRoundUnitPrice returns a boolean if a field has been set.

### GetTaxBreakdown

`func (o *OrganizationCustomizationPdfExtra) GetTaxBreakdown() bool`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *OrganizationCustomizationPdfExtra) GetTaxBreakdownOk() (*bool, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *OrganizationCustomizationPdfExtra) SetTaxBreakdown(v bool)`

SetTaxBreakdown sets TaxBreakdown field to given value.

### HasTaxBreakdown

`func (o *OrganizationCustomizationPdfExtra) HasTaxBreakdown() bool`

HasTaxBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


