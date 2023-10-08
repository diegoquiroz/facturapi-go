# LineItemProduct

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
**Id** | Pointer to **string** | ID del producto base. Sólo presente si se utilizó como base un objeto &#x60;Product&#x60; guardado previamente. | [optional] 

## Methods

### NewLineItemProduct

`func NewLineItemProduct() *LineItemProduct`

NewLineItemProduct instantiates a new LineItemProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemProductWithDefaults

`func NewLineItemProductWithDefaults() *LineItemProduct`

NewLineItemProductWithDefaults instantiates a new LineItemProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LineItemProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LineItemProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductKey

`func (o *LineItemProduct) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *LineItemProduct) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *LineItemProduct) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *LineItemProduct) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetPrice

`func (o *LineItemProduct) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *LineItemProduct) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *LineItemProduct) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *LineItemProduct) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxIncluded

`func (o *LineItemProduct) GetTaxIncluded() bool`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *LineItemProduct) GetTaxIncludedOk() (*bool, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *LineItemProduct) SetTaxIncluded(v bool)`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *LineItemProduct) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### GetTaxability

`func (o *LineItemProduct) GetTaxability() string`

GetTaxability returns the Taxability field if non-nil, zero value otherwise.

### GetTaxabilityOk

`func (o *LineItemProduct) GetTaxabilityOk() (*string, bool)`

GetTaxabilityOk returns a tuple with the Taxability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxability

`func (o *LineItemProduct) SetTaxability(v string)`

SetTaxability sets Taxability field to given value.

### HasTaxability

`func (o *LineItemProduct) HasTaxability() bool`

HasTaxability returns a boolean if a field has been set.

### GetTaxes

`func (o *LineItemProduct) GetTaxes() []BaseTax`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *LineItemProduct) GetTaxesOk() (*[]BaseTax, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *LineItemProduct) SetTaxes(v []BaseTax)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *LineItemProduct) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetLocalTaxes

`func (o *LineItemProduct) GetLocalTaxes() []LocalTax`

GetLocalTaxes returns the LocalTaxes field if non-nil, zero value otherwise.

### GetLocalTaxesOk

`func (o *LineItemProduct) GetLocalTaxesOk() (*[]LocalTax, bool)`

GetLocalTaxesOk returns a tuple with the LocalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTaxes

`func (o *LineItemProduct) SetLocalTaxes(v []LocalTax)`

SetLocalTaxes sets LocalTaxes field to given value.

### HasLocalTaxes

`func (o *LineItemProduct) HasLocalTaxes() bool`

HasLocalTaxes returns a boolean if a field has been set.

### GetUnitKey

`func (o *LineItemProduct) GetUnitKey() string`

GetUnitKey returns the UnitKey field if non-nil, zero value otherwise.

### GetUnitKeyOk

`func (o *LineItemProduct) GetUnitKeyOk() (*string, bool)`

GetUnitKeyOk returns a tuple with the UnitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitKey

`func (o *LineItemProduct) SetUnitKey(v string)`

SetUnitKey sets UnitKey field to given value.

### HasUnitKey

`func (o *LineItemProduct) HasUnitKey() bool`

HasUnitKey returns a boolean if a field has been set.

### GetUnitName

`func (o *LineItemProduct) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *LineItemProduct) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *LineItemProduct) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *LineItemProduct) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetSku

`func (o *LineItemProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItemProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItemProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItemProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetId

`func (o *LineItemProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineItemProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineItemProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LineItemProduct) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


