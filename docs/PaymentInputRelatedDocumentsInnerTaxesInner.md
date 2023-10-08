# PaymentInputRelatedDocumentsInnerTaxesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | **float32** | Base utilizada para el cálculo del impuestos.  | 
**Type** | **string** | Tipo de impuesto.  | 
**Rate** | **float32** | Tasa o cuota del impuesto  | 
**Factor** | Pointer to **string** | Tipo factor. | [optional] [default to "Tasa"]
**Withholding** | Pointer to **bool** | Indica si el impuesto es una retención (&#x60;true&#x60;) o un traslado (&#x60;false&#x60;). | [optional] [default to false]

## Methods

### NewPaymentInputRelatedDocumentsInnerTaxesInner

`func NewPaymentInputRelatedDocumentsInnerTaxesInner(base float32, type_ string, rate float32, ) *PaymentInputRelatedDocumentsInnerTaxesInner`

NewPaymentInputRelatedDocumentsInnerTaxesInner instantiates a new PaymentInputRelatedDocumentsInnerTaxesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInputRelatedDocumentsInnerTaxesInnerWithDefaults

`func NewPaymentInputRelatedDocumentsInnerTaxesInnerWithDefaults() *PaymentInputRelatedDocumentsInnerTaxesInner`

NewPaymentInputRelatedDocumentsInnerTaxesInnerWithDefaults instantiates a new PaymentInputRelatedDocumentsInnerTaxesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetBase() float32`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetBaseOk() (*float32, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) SetBase(v float32)`

SetBase sets Base field to given value.


### GetType

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) SetType(v string)`

SetType sets Type field to given value.


### GetRate

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) SetRate(v float32)`

SetRate sets Rate field to given value.


### GetFactor

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) SetFactor(v string)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetWithholding

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetWithholding() bool`

GetWithholding returns the Withholding field if non-nil, zero value otherwise.

### GetWithholdingOk

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) GetWithholdingOk() (*bool, bool)`

GetWithholdingOk returns a tuple with the Withholding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithholding

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) SetWithholding(v bool)`

SetWithholding sets Withholding field to given value.

### HasWithholding

`func (o *PaymentInputRelatedDocumentsInnerTaxesInner) HasWithholding() bool`

HasWithholding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


