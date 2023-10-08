# BaseTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **float32** | Tasa del impuesto en fracción decimal. | 
**Base** | Pointer to **float32** | Base del impuesto. Valor: 100% del subtotal. | [optional] 
**Type** | Pointer to **string** | Tipo de impuesto. | [optional] [default to "IVA"]
**Factor** | Pointer to **string** | Tipo factor | [optional] [default to "Tasa"]
**Withholding** | Pointer to **bool** | Indica si se trata de un impuesto retenido (&#x60;true&#x60;), o un impuesto trasladado (&#x60;false&#x60;) | [optional] [default to false]

## Methods

### NewBaseTax

`func NewBaseTax(rate float32, ) *BaseTax`

NewBaseTax instantiates a new BaseTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTaxWithDefaults

`func NewBaseTaxWithDefaults() *BaseTax`

NewBaseTaxWithDefaults instantiates a new BaseTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *BaseTax) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *BaseTax) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *BaseTax) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetBase

`func (o *BaseTax) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *BaseTax) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *BaseTax) SetBase(v float32)`

SetBase sets Base field to given value.

### HasBase

`func (o *BaseTax) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetType

`func (o *BaseTax) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseTax) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseTax) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseTax) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFactor

`func (o *BaseTax) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *BaseTax) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *BaseTax) SetFactor(v string)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *BaseTax) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetWithholding

`func (o *BaseTax) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *BaseTax) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *BaseTax) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *BaseTax) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


