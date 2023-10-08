# Parts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Descripción del producto o servicio. | [optional] 
**ProductKey** | Pointer to **string** | Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra herramienta de búsqueda de claves. | [optional] 
**Quantity** | Pointer to **float32** | Cantidad | [optional] 
**Sku** | Pointer to **string** | Identificador de uso interno designado por la empresa. Puede tener cualquier valor. | [optional] 
**UnitPrice** | Pointer to **float32** | Precio unitario | [optional] 
**UnitName** | Pointer to **string** | Nombre de la unidad de medida que expresa la cantidad. | [optional] 
**CustomsKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewParts

`func NewParts() *Parts`

NewParts instantiates a new Parts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartsWithDefaults

`func NewPartsWithDefaults() *Parts`

NewPartsWithDefaults instantiates a new Parts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Parts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Parts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Parts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Parts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProductKey

`func (o *Parts) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *Parts) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *Parts) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *Parts) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetQuantity

`func (o *Parts) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Parts) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Parts) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Parts) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *Parts) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Parts) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Parts) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Parts) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitPrice

`func (o *Parts) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Parts) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Parts) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *Parts) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnitName

`func (o *Parts) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *Parts) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *Parts) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *Parts) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetCustomsKeys

`func (o *Parts) GetCustomsKeys() []string`

GetCustomsKeys returns the CustomsKeys field if non-nil, zero value otherwise.

### GetCustomsKeysOk

`func (o *Parts) GetCustomsKeysOk() (*[]string, bool)`

GetCustomsKeysOk returns a tuple with the CustomsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsKeys

`func (o *Parts) SetCustomsKeys(v []string)`

SetCustomsKeys sets CustomsKeys field to given value.

### HasCustomsKeys

`func (o *Parts) HasCustomsKeys() bool`

HasCustomsKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


