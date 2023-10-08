# LineItemTrasladoInputProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Descripción del bien o servicio como aparecerá en la factura. | 
**ProductKey** | Pointer to **string** | Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys). | [optional] 
**UnitKey** | Pointer to **string** | Clave de unidad de medida, del catálogo del SAT. El valor por default &#x60;\&quot;H87\&quot;&#x60; (elemento) es la clave para representar una pieza o unidad de venta (lápiz, cuaderno, televisión, etc). Si la unidad de tu producto es kilogramos, litros, horas u otra unidad, te proporcionamos una manera conveniente de encontrar la clave utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys).  | [optional] [default to "H87"]
**UnitName** | Pointer to **string** | Palabra que representa la unidad de medida de tu producto. Debe estar relacionada con la clave de unidad &#x60;unit_key&#x60;. | [optional] [default to "Elemento"]
**Sku** | Pointer to **string** | Identificador de uso interno designado por la empresa. Puede tener cualquier valor. | [optional] 

## Methods

### NewLineItemTrasladoInputProduct

`func NewLineItemTrasladoInputProduct(description string, ) *LineItemTrasladoInputProduct`

NewLineItemTrasladoInputProduct instantiates a new LineItemTrasladoInputProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemTrasladoInputProductWithDefaults

`func NewLineItemTrasladoInputProductWithDefaults() *LineItemTrasladoInputProduct`

NewLineItemTrasladoInputProductWithDefaults instantiates a new LineItemTrasladoInputProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LineItemTrasladoInputProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LineItemTrasladoInputProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LineItemTrasladoInputProduct) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductKey

`func (o *LineItemTrasladoInputProduct) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *LineItemTrasladoInputProduct) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *LineItemTrasladoInputProduct) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *LineItemTrasladoInputProduct) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetUnitKey

`func (o *LineItemTrasladoInputProduct) GetUnitKey() string`

GetUnitKey returns the UnitKey field if non-nil, zero value otherwise.

### GetUnitKeyOk

`func (o *LineItemTrasladoInputProduct) GetUnitKeyOk() (*string, bool)`

GetUnitKeyOk returns a tuple with the UnitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitKey

`func (o *LineItemTrasladoInputProduct) SetUnitKey(v string)`

SetUnitKey sets UnitKey field to given value.

### HasUnitKey

`func (o *LineItemTrasladoInputProduct) HasUnitKey() bool`

HasUnitKey returns a boolean if a field has been set.

### GetUnitName

`func (o *LineItemTrasladoInputProduct) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *LineItemTrasladoInputProduct) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *LineItemTrasladoInputProduct) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *LineItemTrasladoInputProduct) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetSku

`func (o *LineItemTrasladoInputProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *LineItemTrasladoInputProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *LineItemTrasladoInputProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *LineItemTrasladoInputProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


