# LineItemEgresoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | Cantidad de unidades incluídas del mismo concepto. | [optional] [default to 1]
**Discount** | Pointer to **float32** | Monto total de descuento aplicado a este concepto. | [optional] [default to 0]
**Product** | [**LineItemEgresoInputProduct**](LineItemEgresoInputProduct.md) |  | 
**Parts** | Pointer to [**[]PartInput**](PartInput.md) |  | [optional] 
**CustomsKeys** | Pointer to **[]string** |  | [optional] 
**Complement** | Pointer to **string** | Código XML de tu complemento. | [optional] 
**ThirdParty** | Pointer to [**LineItemInputThirdParty**](LineItemInputThirdParty.md) |  | [optional] 

## Methods

### NewLineItemEgresoInput

`func NewLineItemEgresoInput(product LineItemEgresoInputProduct, ) *LineItemEgresoInput`

NewLineItemEgresoInput instantiates a new LineItemEgresoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemEgresoInputWithDefaults

`func NewLineItemEgresoInputWithDefaults() *LineItemEgresoInput`

NewLineItemEgresoInputWithDefaults instantiates a new LineItemEgresoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *LineItemEgresoInput) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItemEgresoInput) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItemEgresoInput) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItemEgresoInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDiscount

`func (o *LineItemEgresoInput) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *LineItemEgresoInput) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *LineItemEgresoInput) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *LineItemEgresoInput) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetProduct

`func (o *LineItemEgresoInput) GetProduct() LineItemEgresoInputProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItemEgresoInput) GetProductOk() (*LineItemEgresoInputProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItemEgresoInput) SetProduct(v LineItemEgresoInputProduct)`

SetProduct sets Product field to given value.


### GetParts

`func (o *LineItemEgresoInput) GetParts() []PartInput`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *LineItemEgresoInput) GetPartsOk() (*[]PartInput, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *LineItemEgresoInput) SetParts(v []PartInput)`

SetParts sets Parts field to given value.

### HasParts

`func (o *LineItemEgresoInput) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetCustomsKeys

`func (o *LineItemEgresoInput) GetCustomsKeys() []string`

GetCustomsKeys returns the CustomsKeys field if non-nil, zero value otherwise.

### GetCustomsKeysOk

`func (o *LineItemEgresoInput) GetCustomsKeysOk() (*[]string, bool)`

GetCustomsKeysOk returns a tuple with the CustomsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsKeys

`func (o *LineItemEgresoInput) SetCustomsKeys(v []string)`

SetCustomsKeys sets CustomsKeys field to given value.

### HasCustomsKeys

`func (o *LineItemEgresoInput) HasCustomsKeys() bool`

HasCustomsKeys returns a boolean if a field has been set.

### GetComplement

`func (o *LineItemEgresoInput) GetComplement() string`

GetComplement returns the Complement field if non-nil, zero value otherwise.

### GetComplementOk

`func (o *LineItemEgresoInput) GetComplementOk() (*string, bool)`

GetComplementOk returns a tuple with the Complement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplement

`func (o *LineItemEgresoInput) SetComplement(v string)`

SetComplement sets Complement field to given value.

### HasComplement

`func (o *LineItemEgresoInput) HasComplement() bool`

HasComplement returns a boolean if a field has been set.

### GetThirdParty

`func (o *LineItemEgresoInput) GetThirdParty() LineItemInputThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *LineItemEgresoInput) GetThirdPartyOk() (*LineItemInputThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *LineItemEgresoInput) SetThirdParty(v LineItemInputThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *LineItemEgresoInput) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


