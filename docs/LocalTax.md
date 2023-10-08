# LocalTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | **float32** | Tasa del impuesto en fracción decimal. | 
**Base** | Pointer to **float32** | Base del impuesto. Valor: 100% del subtotal. | [optional] 
**Type** | **string** | Nombre del impuesto. Texto libre. | 
**Withholding** | Pointer to **bool** | Indica si se trata de un impuesto retenido (&#x60;true&#x60;), o un impuesto trasladado (&#x60;false&#x60;) | [optional] [default to false]

## Methods

### NewLocalTax

`func NewLocalTax(rate float32, type_ string, ) *LocalTax`

NewLocalTax instantiates a new LocalTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalTaxWithDefaults

`func NewLocalTaxWithDefaults() *LocalTax`

NewLocalTaxWithDefaults instantiates a new LocalTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *LocalTax) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *LocalTax) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *LocalTax) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetBase

`func (o *LocalTax) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *LocalTax) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *LocalTax) SetBase(v float32)`

SetBase sets Base field to given value.

### HasBase

`func (o *LocalTax) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetType

`func (o *LocalTax) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalTax) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalTax) SetType(v string)`

SetType sets Type field to given value.


### GetWithholding

`func (o *LocalTax) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *LocalTax) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *LocalTax) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *LocalTax) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


