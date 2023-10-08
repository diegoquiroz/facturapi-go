# LineItemProductEgresoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Resumen de la operación en una sola descripción. Deben mencionarse cada uno de los productos que contempla el descuento, devolución o bonificación aplicada y que contienen las facturas relacionadas. Si el egreso está basado en un pocentaje (como al aplicar un 30% de descuento), dicho porcentaje debe incluirse en la descripción junto al nombre del producto que corresponda. | [optional] 
**ProductKey** | Pointer to **string** | Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys). | [optional] [default to "84111506"]
**Price** | Pointer to **float32** | Suma total de la cantidad devuelta, descontada o bonificada. | [optional] 
**TaxIncluded** | Pointer to **bool** | - &#x60;true&#x60;: Indica que todos los impuestos aplicables están incluídos en el precio (atributo price) y se desglosarán automáticamente al emitir la factura. - &#x60;false&#x60;: Indica que el atributo price no incluye impuestos, por lo que aquellos impuestos a aplicar se sumarán en el precio final.  | [optional] [default to true]
**Taxability** | Pointer to **string** | Código que representa si el bien o servicio es objeto de impuesto o no. Este atributo corresponde al campo \&quot;ObjetoImp\&quot; en el CFDI.  - &#x60;01&#x60;: No objeto de impuesto. - &#x60;02&#x60;: Sí objeto de impuesto. - &#x60;03&#x60;: Sí objeto de impuesto, pero no obligado a desglose. - &#x60;04&#x60;: Sí objeto de impuesto, y no causa impuesto.  | [optional] [default to "02"]
**Taxes** | Pointer to [**[]BaseTax**](BaseTax.md) |  | [optional] 
**LocalTaxes** | Pointer to [**[]LocalTax**](LocalTax.md) |  | [optional] 
**UnitKey** | Pointer to **string** | Clave de unidad de medida, del catálogo del SAT. Puedes encontrar la clave utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys).  | [optional] [default to "ACT"]
**UnitName** | Pointer to **string** | Palabra que representa la unidad de medida de tu producto. Debe estar relacionada con la clave de unidad &#x60;unit_key&#x60;. | [optional] [default to "Actividad"]

## Methods

### NewLineItemProductEgresoInput

`func NewLineItemProductEgresoInput() *LineItemProductEgresoInput`

NewLineItemProductEgresoInput instantiates a new LineItemProductEgresoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemProductEgresoInputWithDefaults

`func NewLineItemProductEgresoInputWithDefaults() *LineItemProductEgresoInput`

NewLineItemProductEgresoInputWithDefaults instantiates a new LineItemProductEgresoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LineItemProductEgresoInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemProductEgresoInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemProductEgresoInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItemProductEgresoInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductKey

`func (o *LineItemProductEgresoInput) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *LineItemProductEgresoInput) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *LineItemProductEgresoInput) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *LineItemProductEgresoInput) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetPrice

`func (o *LineItemProductEgresoInput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LineItemProductEgresoInput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LineItemProductEgresoInput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LineItemProductEgresoInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *LineItemProductEgresoInput) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *LineItemProductEgresoInput) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *LineItemProductEgresoInput) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *LineItemProductEgresoInput) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetTaxability

`func (o *LineItemProductEgresoInput) GetTaxability() string`

GetTaxability returns the Taxability field if non-nil, zero value otherwise.

### GetTaxabilityOk

`func (o *LineItemProductEgresoInput) GetTaxabilityOk() (*string, bool)`

GetTaxabilityOk returns a tuple with the Taxability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxability

`func (o *LineItemProductEgresoInput) SetTaxability(v string)`

SetTaxability sets Taxability field to given value.

### HasTaxability

`func (o *LineItemProductEgresoInput) HasTaxability() bool`

HasTaxability returns a boolean if a field has been set.

### GetTaxes

`func (o *LineItemProductEgresoInput) GetTaxes() []BaseTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *LineItemProductEgresoInput) GetTaxesOk() (*[]BaseTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *LineItemProductEgresoInput) SetTaxes(v []BaseTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *LineItemProductEgresoInput) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetLocalTaxes

`func (o *LineItemProductEgresoInput) GetLocalTaxes() []LocalTax`

GetLocalTaxes returns the LocalTaxes field if non-nil, zero value otherwise.

### GetLocalTaxesOk

`func (o *LineItemProductEgresoInput) GetLocalTaxesOk() (*[]LocalTax, bool)`

GetLocalTaxesOk returns a tuple with the LocalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTaxes

`func (o *LineItemProductEgresoInput) SetLocalTaxes(v []LocalTax)`

SetLocalTaxes sets LocalTaxes field to given value.

### HasLocalTaxes

`func (o *LineItemProductEgresoInput) HasLocalTaxes() bool`

HasLocalTaxes returns a boolean if a field has been set.

### GetUnitKey

`func (o *LineItemProductEgresoInput) GetUnitKey() string`

GetUnitKey returns the UnitKey field if non-nil, zero value otherwise.

### GetUnitKeyOk

`func (o *LineItemProductEgresoInput) GetUnitKeyOk() (*string, bool)`

GetUnitKeyOk returns a tuple with the UnitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitKey

`func (o *LineItemProductEgresoInput) SetUnitKey(v string)`

SetUnitKey sets UnitKey field to given value.

### HasUnitKey

`func (o *LineItemProductEgresoInput) HasUnitKey() bool`

HasUnitKey returns a boolean if a field has been set.

### GetUnitName

`func (o *LineItemProductEgresoInput) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *LineItemProductEgresoInput) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *LineItemProductEgresoInput) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *LineItemProductEgresoInput) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


