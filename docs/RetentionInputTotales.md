# RetentionInputTotales

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MontoTotOperacion** | **float32** | Monto total de la operación, con precisión de hasta 6 decimales. | 
**MontoTotGrav** | Pointer to **float32** | Monto total gravado. | [optional] 
**MontoTotExent** | **float32** | Monto total exento. | 
**MontoTotRet** | Pointer to **float32** | Suma de los montos de impuestos retenidos. | [optional] 
**ImpRetenidos** | [**[]RetentionInputTotalesImpRetenidosInner**](RetentionInputTotalesImpRetenidosInner.md) |  | 

## Methods

### NewRetentionInputTotales

`func NewRetentionInputTotales(montoTotOperacion float32, montoTotExent float32, impRetenidos []RetentionInputTotalesImpRetenidosInner, ) *RetentionInputTotales`

NewRetentionInputTotales instantiates a new RetentionInputTotales object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionInputTotalesWithDefaults

`func NewRetentionInputTotalesWithDefaults() *RetentionInputTotales`

NewRetentionInputTotalesWithDefaults instantiates a new RetentionInputTotales object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMontoTotOperacion

`func (o *RetentionInputTotales) GetMontoTotOperacion() float32`

GetMontoTotOperacion returns the MontoTotOperacion field if non-nil, zero value otherwise.

### GetMontoTotOperacionOk

`func (o *RetentionInputTotales) GetMontoTotOperacionOk() (*float32, bool)`

GetMontoTotOperacionOk returns a tuple with the MontoTotOperacion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoTotOperacion

`func (o *RetentionInputTotales) SetMontoTotOperacion(v float32)`

SetMontoTotOperacion sets MontoTotOperacion field to given value.


### GetMontoTotGrav

`func (o *RetentionInputTotales) GetMontoTotGrav() float32`

GetMontoTotGrav returns the MontoTotGrav field if non-nil, zero value otherwise.

### GetMontoTotGravOk

`func (o *RetentionInputTotales) GetMontoTotGravOk() (*float32, bool)`

GetMontoTotGravOk returns a tuple with the MontoTotGrav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoTotGrav

`func (o *RetentionInputTotales) SetMontoTotGrav(v float32)`

SetMontoTotGrav sets MontoTotGrav field to given value.

### HasMontoTotGrav

`func (o *RetentionInputTotales) HasMontoTotGrav() bool`

HasMontoTotGrav returns a boolean if a field has been set.

### GetMontoTotExent

`func (o *RetentionInputTotales) GetMontoTotExent() float32`

GetMontoTotExent returns the MontoTotExent field if non-nil, zero value otherwise.

### GetMontoTotExentOk

`func (o *RetentionInputTotales) GetMontoTotExentOk() (*float32, bool)`

GetMontoTotExentOk returns a tuple with the MontoTotExent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoTotExent

`func (o *RetentionInputTotales) SetMontoTotExent(v float32)`

SetMontoTotExent sets MontoTotExent field to given value.


### GetMontoTotRet

`func (o *RetentionInputTotales) GetMontoTotRet() float32`

GetMontoTotRet returns the MontoTotRet field if non-nil, zero value otherwise.

### GetMontoTotRetOk

`func (o *RetentionInputTotales) GetMontoTotRetOk() (*float32, bool)`

GetMontoTotRetOk returns a tuple with the MontoTotRet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoTotRet

`func (o *RetentionInputTotales) SetMontoTotRet(v float32)`

SetMontoTotRet sets MontoTotRet field to given value.

### HasMontoTotRet

`func (o *RetentionInputTotales) HasMontoTotRet() bool`

HasMontoTotRet returns a boolean if a field has been set.

### GetImpRetenidos

`func (o *RetentionInputTotales) GetImpRetenidos() []RetentionInputTotalesImpRetenidosInner`

GetImpRetenidos returns the ImpRetenidos field if non-nil, zero value otherwise.

### GetImpRetenidosOk

`func (o *RetentionInputTotales) GetImpRetenidosOk() (*[]RetentionInputTotalesImpRetenidosInner, bool)`

GetImpRetenidosOk returns a tuple with the ImpRetenidos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpRetenidos

`func (o *RetentionInputTotales) SetImpRetenidos(v []RetentionInputTotalesImpRetenidosInner)`

SetImpRetenidos sets ImpRetenidos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


