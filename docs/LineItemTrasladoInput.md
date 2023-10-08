# LineItemTrasladoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | Cantidad de unidades incluídas del mismo concepto. | [optional] [default to 1]
**Product** | [**LineItemTrasladoInputProduct**](LineItemTrasladoInputProduct.md) |  | 
**CustomsKeys** | Pointer to **[]string** |  | [optional] 
**Complement** | Pointer to **string** | Código XML de tu complemento. | [optional] 
**Parts** | Pointer to [**[]PartInput**](PartInput.md) |  | [optional] 
**ThirdParty** | Pointer to [**LineItemTrasladoInputThirdParty**](LineItemTrasladoInputThirdParty.md) |  | [optional] 

## Methods

### NewLineItemTrasladoInput

`func NewLineItemTrasladoInput(product LineItemTrasladoInputProduct, ) *LineItemTrasladoInput`

NewLineItemTrasladoInput instantiates a new LineItemTrasladoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemTrasladoInputWithDefaults

`func NewLineItemTrasladoInputWithDefaults() *LineItemTrasladoInput`

NewLineItemTrasladoInputWithDefaults instantiates a new LineItemTrasladoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *LineItemTrasladoInput) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *LineItemTrasladoInput) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *LineItemTrasladoInput) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *LineItemTrasladoInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetProduct

`func (o *LineItemTrasladoInput) GetProduct() LineItemTrasladoInputProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LineItemTrasladoInput) GetProductOk() (*LineItemTrasladoInputProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LineItemTrasladoInput) SetProduct(v LineItemTrasladoInputProduct)`

SetProduct sets Product field to given value.


### GetCustomsKeys

`func (o *LineItemTrasladoInput) GetCustomsKeys() []string`

GetCustomsKeys returns the CustomsKeys field if non-nil, zero value otherwise.

### GetCustomsKeysOk

`func (o *LineItemTrasladoInput) GetCustomsKeysOk() (*[]string, bool)`

GetCustomsKeysOk returns a tuple with the CustomsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsKeys

`func (o *LineItemTrasladoInput) SetCustomsKeys(v []string)`

SetCustomsKeys sets CustomsKeys field to given value.

### HasCustomsKeys

`func (o *LineItemTrasladoInput) HasCustomsKeys() bool`

HasCustomsKeys returns a boolean if a field has been set.

### GetComplement

`func (o *LineItemTrasladoInput) GetComplement() string`

GetComplement returns the Complement field if non-nil, zero value otherwise.

### GetComplementOk

`func (o *LineItemTrasladoInput) GetComplementOk() (*string, bool)`

GetComplementOk returns a tuple with the Complement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplement

`func (o *LineItemTrasladoInput) SetComplement(v string)`

SetComplement sets Complement field to given value.

### HasComplement

`func (o *LineItemTrasladoInput) HasComplement() bool`

HasComplement returns a boolean if a field has been set.

### GetParts

`func (o *LineItemTrasladoInput) GetParts() []PartInput`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *LineItemTrasladoInput) GetPartsOk() (*[]PartInput, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *LineItemTrasladoInput) SetParts(v []PartInput)`

SetParts sets Parts field to given value.

### HasParts

`func (o *LineItemTrasladoInput) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetThirdParty

`func (o *LineItemTrasladoInput) GetThirdParty() LineItemTrasladoInputThirdParty`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *LineItemTrasladoInput) GetThirdPartyOk() (*LineItemTrasladoInputThirdParty, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *LineItemTrasladoInput) SetThirdParty(v LineItemTrasladoInputThirdParty)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *LineItemTrasladoInput) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


