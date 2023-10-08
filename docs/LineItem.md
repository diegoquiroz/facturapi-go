# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | Cantidad de unidades incluídas del mismo concepto. | [optional] 
**Discount** | Pointer to **float32** | Monto total de descuento aplicado a este concepto. | [optional] 
**Product** | Pointer to [**LineItemProduct**](LineItemProduct.md) | Objeto con información del producto o servicio facturado. | [optional] 
**Parts** | Pointer to [**Parts**](Parts.md) | Objeto con información de las partes de la factura. | [optional] 

## Methods

### NewLineItem

`func NewLineItem() *LineItem`

NewLineItem instantiates a new LineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *LineItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDiscount

`func (o *LineItem) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *LineItem) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *LineItem) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *LineItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetProduct

`func (o *LineItem) GetProduct() LineItemProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItem) GetProductOk() (*LineItemProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItem) SetProduct(v LineItemProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LineItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetParts

`func (o *LineItem) GetParts() Parts`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *LineItem) GetPartsOk() (*Parts, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *LineItem) SetParts(v Parts)`

SetParts sets Parts field to given value.

### HasParts

`func (o *LineItem) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


