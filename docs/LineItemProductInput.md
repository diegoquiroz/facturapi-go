# LineItemProductInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Descripción del bien o servicio como aparecerá en la factura. | [optional] 
**ProductKey** | Pointer to **string** | Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys). | [optional] 
**Price** | Pointer to **float32** | Precio por unidad del bien o servicio. Este valor representará el precio con IVA incluído o sin él, dependiendo del valor de &#x60;tax_included&#x60;. | [optional] 
**TaxIncluded** | Pointer to **bool** | - &#x60;true&#x60;: Indica que todos los impuestos aplicables están incluídos en el precio (atributo price) y se desglosarán automáticamente al emitir la factura. - &#x60;false&#x60;: Indica que el atributo price no incluye impuestos, por lo que aquellos impuestos a aplicar se sumarán en el precio final.  | [optional] [default to true]
**Taxability** | Pointer to **string** | Código que representa si el bien o servicio es objeto de impuesto o no. Este atributo corresponde al campo \&quot;ObjetoImp\&quot; en el CFDI.  - &#x60;01&#x60;: No objeto de impuesto. - &#x60;02&#x60;: Sí objeto de impuesto. - &#x60;03&#x60;: Sí objeto de impuesto, pero no obligado a desglose. - &#x60;04&#x60;: Sí objeto de impuesto, y no causa impuesto.  | [optional] [default to "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "]
**Taxes** | Pointer to [**[]BaseTax**](BaseTax.md) |  | [optional] 
**LocalTaxes** | Pointer to [**[]LocalTax**](LocalTax.md) |  | [optional] 
**UnitKey** | Pointer to **string** | Clave de unidad de medida, del catálogo del SAT. El valor por default &#x60;\&quot;H87\&quot;&#x60; (elemento) es la clave para representar una pieza o unidad de venta (lápiz, cuaderno, televisión, etc). Si la unidad de tu producto es kilogramos, litros, horas u otra unidad, puedes encontrar la clave utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys).  | [optional] [default to "H87"]
**UnitName** | Pointer to **string** | Palabra que representa la unidad de medida de tu producto. Debe estar relacionada con la clave de unidad &#x60;unit_key&#x60;. | [optional] [default to "Elemento"]
**Sku** | Pointer to **string** | Identificador de uso interno designado por la empresa. Puede tener cualquier valor. | [optional] 

## Methods

### NewLineItemProductInput

`func NewLineItemProductInput() *LineItemProductInput`

NewLineItemProductInput instantiates a new LineItemProductInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemProductInputWithDefaults

`func NewLineItemProductInputWithDefaults() *LineItemProductInput`

NewLineItemProductInputWithDefaults instantiates a new LineItemProductInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LineItemProductInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemProductInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemProductInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItemProductInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductKey

`func (o *LineItemProductInput) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *LineItemProductInput) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *LineItemProductInput) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *LineItemProductInput) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetPrice

`func (o *LineItemProductInput) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LineItemProductInput) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LineItemProductInput) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LineItemProductInput) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *LineItemProductInput) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *LineItemProductInput) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *LineItemProductInput) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *LineItemProductInput) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetTaxability

`func (o *LineItemProductInput) GetTaxability() string`

GetTaxability returns the Taxability field if non-nil, zero value otherwise.

### GetTaxabilityOk

`func (o *LineItemProductInput) GetTaxabilityOk() (*string, bool)`

GetTaxabilityOk returns a tuple with the Taxability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxability

`func (o *LineItemProductInput) SetTaxability(v string)`

SetTaxability sets Taxability field to given value.

### HasTaxability

`func (o *LineItemProductInput) HasTaxability() bool`

HasTaxability returns a boolean if a field has been set.

### GetTaxes

`func (o *LineItemProductInput) GetTaxes() []BaseTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *LineItemProductInput) GetTaxesOk() (*[]BaseTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *LineItemProductInput) SetTaxes(v []BaseTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *LineItemProductInput) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetLocalTaxes

`func (o *LineItemProductInput) GetLocalTaxes() []LocalTax`

GetLocalTaxes returns the LocalTaxes field if non-nil, zero value otherwise.

### GetLocalTaxesOk

`func (o *LineItemProductInput) GetLocalTaxesOk() (*[]LocalTax, bool)`

GetLocalTaxesOk returns a tuple with the LocalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTaxes

`func (o *LineItemProductInput) SetLocalTaxes(v []LocalTax)`

SetLocalTaxes sets LocalTaxes field to given value.

### HasLocalTaxes

`func (o *LineItemProductInput) HasLocalTaxes() bool`

HasLocalTaxes returns a boolean if a field has been set.

### GetUnitKey

`func (o *LineItemProductInput) GetUnitKey() string`

GetUnitKey returns the UnitKey field if non-nil, zero value otherwise.

### GetUnitKeyOk

`func (o *LineItemProductInput) GetUnitKeyOk() (*string, bool)`

GetUnitKeyOk returns a tuple with the UnitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitKey

`func (o *LineItemProductInput) SetUnitKey(v string)`

SetUnitKey sets UnitKey field to given value.

### HasUnitKey

`func (o *LineItemProductInput) HasUnitKey() bool`

HasUnitKey returns a boolean if a field has been set.

### GetUnitName

`func (o *LineItemProductInput) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *LineItemProductInput) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *LineItemProductInput) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *LineItemProductInput) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetSku

`func (o *LineItemProductInput) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItemProductInput) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItemProductInput) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItemProductInput) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


