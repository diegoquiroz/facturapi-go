# RetentionInputTotalesImpRetenidosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseRet** | Pointer to **float32** | Base del impuesto retenido. | [optional] 
**Impuesto** | Pointer to **string** | Clave del tipo de impuesto retenido, del catálogo del SAT. | [optional] 
**MontoRet** | **float32** | Importe del impuesto retenido | 
**TipoPagoRet** | **string** | - &#x60;01&#x60;: Pago definitivo IVA - &#x60;02&#x60;: Pago definitivo IEPS - &#x60;03&#x60;: Pago definitivo ISR Plataformas - &#x60;04&#x60;: Pago provisional ISR  | 

## Methods

### NewRetentionInputTotalesImpRetenidosInner

`func NewRetentionInputTotalesImpRetenidosInner(montoRet float32, tipoPagoRet string, ) *RetentionInputTotalesImpRetenidosInner`

NewRetentionInputTotalesImpRetenidosInner instantiates a new RetentionInputTotalesImpRetenidosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionInputTotalesImpRetenidosInnerWithDefaults

`func NewRetentionInputTotalesImpRetenidosInnerWithDefaults() *RetentionInputTotalesImpRetenidosInner`

NewRetentionInputTotalesImpRetenidosInnerWithDefaults instantiates a new RetentionInputTotalesImpRetenidosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseRet

`func (o *RetentionInputTotalesImpRetenidosInner) GetBaseRet() float32`

GetBaseRet returns the BaseRet field if non-nil, zero value otherwise.

### GetBaseRetOk

`func (o *RetentionInputTotalesImpRetenidosInner) GetBaseRetOk() (*float32, bool)`

GetBaseRetOk returns a tuple with the BaseRet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRet

`func (o *RetentionInputTotalesImpRetenidosInner) SetBaseRet(v float32)`

SetBaseRet sets BaseRet field to given value.

### HasBaseRet

`func (o *RetentionInputTotalesImpRetenidosInner) HasBaseRet() bool`

HasBaseRet returns a boolean if a field has been set.

### GetImpuesto

`func (o *RetentionInputTotalesImpRetenidosInner) GetImpuesto() string`

GetImpuesto returns the Impuesto field if non-nil, zero value otherwise.

### GetImpuestoOk

`func (o *RetentionInputTotalesImpRetenidosInner) GetImpuestoOk() (*string, bool)`

GetImpuestoOk returns a tuple with the Impuesto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpuesto

`func (o *RetentionInputTotalesImpRetenidosInner) SetImpuesto(v string)`

SetImpuesto sets Impuesto field to given value.

### HasImpuesto

`func (o *RetentionInputTotalesImpRetenidosInner) HasImpuesto() bool`

HasImpuesto returns a boolean if a field has been set.

### GetMontoRet

`func (o *RetentionInputTotalesImpRetenidosInner) GetMontoRet() float32`

GetMontoRet returns the MontoRet field if non-nil, zero value otherwise.

### GetMontoRetOk

`func (o *RetentionInputTotalesImpRetenidosInner) GetMontoRetOk() (*float32, bool)`

GetMontoRetOk returns a tuple with the MontoRet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoRet

`func (o *RetentionInputTotalesImpRetenidosInner) SetMontoRet(v float32)`

SetMontoRet sets MontoRet field to given value.


### GetTipoPagoRet

`func (o *RetentionInputTotalesImpRetenidosInner) GetTipoPagoRet() string`

GetTipoPagoRet returns the TipoPagoRet field if non-nil, zero value otherwise.

### GetTipoPagoRetOk

`func (o *RetentionInputTotalesImpRetenidosInner) GetTipoPagoRetOk() (*string, bool)`

GetTipoPagoRetOk returns a tuple with the TipoPagoRet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoPagoRet

`func (o *RetentionInputTotalesImpRetenidosInner) SetTipoPagoRet(v string)`

SetTipoPagoRet sets TipoPagoRet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


