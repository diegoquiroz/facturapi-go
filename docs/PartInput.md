# PartInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Descripción del producto o servicio. | 
**ProductKey** | **string** | Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra herramienta de búsqueda de claves. | 
**Quantity** | Pointer to **float32** | Cantidad | [optional] 
**Sku** | Pointer to **string** | Identificador de uso interno designado por la empresa. Puede tener cualquier valor. | [optional] 
**UnitPrice** | Pointer to **float32** | Precio unitario | [optional] 
**UnitName** | Pointer to **string** | Nombre de la unidad de medida que expresa la cantidad. | [optional] 
**CustomsKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPartInput

`func NewPartInput(description string, productKey string, ) *PartInput`

NewPartInput instantiates a new PartInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartInputWithDefaults

`func NewPartInputWithDefaults() *PartInput`

NewPartInputWithDefaults instantiates a new PartInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PartInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartInput) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductKey

`func (o *PartInput) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *PartInput) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *PartInput) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.


### GetQuantity

`func (o *PartInput) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PartInput) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PartInput) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PartInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSku

`func (o *PartInput) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PartInput) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PartInput) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *PartInput) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetUnitPrice

`func (o *PartInput) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PartInput) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PartInput) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *PartInput) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnitName

`func (o *PartInput) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *PartInput) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *PartInput) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *PartInput) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetCustomsKeys

`func (o *PartInput) GetCustomsKeys() []string`

GetCustomsKeys returns the CustomsKeys field if non-nil, zero value otherwise.

### GetCustomsKeysOk

`func (o *PartInput) GetCustomsKeysOk() (*[]string, bool)`

GetCustomsKeysOk returns a tuple with the CustomsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsKeys

`func (o *PartInput) SetCustomsKeys(v []string)`

SetCustomsKeys sets CustomsKeys field to given value.

### HasCustomsKeys

`func (o *PartInput) HasCustomsKeys() bool`

HasCustomsKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


