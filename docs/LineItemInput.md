# LineItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | Cantidad de unidades incluídas del mismo concepto. | [optional] [default to 1]
**Discount** | Pointer to **float32** | Monto total de descuento aplicado a este concepto. | [optional] [default to 0]
**Product** | [**LineItemInputProduct**](LineItemInputProduct.md) |  | 
**Parts** | Pointer to [**[]PartInput**](PartInput.md) |  | [optional] 
**CustomsKeys** | Pointer to **[]string** |  | [optional] 
**Complement** | Pointer to **string** | Código XML de tu complemento. | [optional] 
**ThirdParty** | Pointer to [**LineItemInputThirdParty**](LineItemInputThirdParty.md) |  | [optional] 
**PropertyTaxAccount** | Pointer to **string** | Número de cuenta para el impuesto predial. | [optional] 

## Methods

### NewLineItemInput

`func NewLineItemInput(product LineItemInputProduct, ) *LineItemInput`

NewLineItemInput instantiates a new LineItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemInputWithDefaults

`func NewLineItemInputWithDefaults() *LineItemInput`

NewLineItemInputWithDefaults instantiates a new LineItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *LineItemInput) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItemInput) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItemInput) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDiscount

`func (o *LineItemInput) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *LineItemInput) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *LineItemInput) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *LineItemInput) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetProduct

`func (o *LineItemInput) GetProduct() LineItemInputProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItemInput) GetProductOk() (*LineItemInputProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItemInput) SetProduct(v LineItemInputProduct)`

SetProduct sets Product field to given value.


### GetParts

`func (o *LineItemInput) GetParts() []PartInput`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *LineItemInput) GetPartsOk() (*[]PartInput, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *LineItemInput) SetParts(v []PartInput)`

SetParts sets Parts field to given value.

### HasParts

`func (o *LineItemInput) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetCustomsKeys

`func (o *LineItemInput) GetCustomsKeys() []string`

GetCustomsKeys returns the CustomsKeys field if non-nil, zero value otherwise.

### GetCustomsKeysOk

`func (o *LineItemInput) GetCustomsKeysOk() (*[]string, bool)`

GetCustomsKeysOk returns a tuple with the CustomsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsKeys

`func (o *LineItemInput) SetCustomsKeys(v []string)`

SetCustomsKeys sets CustomsKeys field to given value.

### HasCustomsKeys

`func (o *LineItemInput) HasCustomsKeys() bool`

HasCustomsKeys returns a boolean if a field has been set.

### GetComplement

`func (o *LineItemInput) GetComplement() string`

GetComplement returns the Complement field if non-nil, zero value otherwise.

### GetComplementOk

`func (o *LineItemInput) GetComplementOk() (*string, bool)`

GetComplementOk returns a tuple with the Complement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplement

`func (o *LineItemInput) SetComplement(v string)`

SetComplement sets Complement field to given value.

### HasComplement

`func (o *LineItemInput) HasComplement() bool`

HasComplement returns a boolean if a field has been set.

### GetThirdParty

`func (o *LineItemInput) GetThirdParty() LineItemInputThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *LineItemInput) GetThirdPartyOk() (*LineItemInputThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *LineItemInput) SetThirdParty(v LineItemInputThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *LineItemInput) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.

### GetPropertyTaxAccount

`func (o *LineItemInput) GetPropertyTaxAccount() string`

GetPropertyTaxAccount returns the PropertyTaxAccount field if non-nil, zero value otherwise.

### GetPropertyTaxAccountOk

`func (o *LineItemInput) GetPropertyTaxAccountOk() (*string, bool)`

GetPropertyTaxAccountOk returns a tuple with the PropertyTaxAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyTaxAccount

`func (o *LineItemInput) SetPropertyTaxAccount(v string)`

SetPropertyTaxAccount sets PropertyTaxAccount field to given value.

### HasPropertyTaxAccount

`func (o *LineItemInput) HasPropertyTaxAccount() bool`

HasPropertyTaxAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


