# NominaJubilacionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUnaExhibicion** | Pointer to **float32** | Monto total del pago entregado en una sola exhibición. | [optional] 
**TotalParcialidad** | Pointer to **float32** | Monto total del pago entregado en parcialidades. | [optional] 
**MontoDiario** | Pointer to **float32** | Monto diario percibido por el trabajador cuando el pago se realiza en parcialidades. | [optional] 
**IngresoAcumulable** | **float32** | Ingresos acumulables percibidos por el trabajador. | 
**IngresoNoAcumulable** | **float32** | Ingresos no acumulables percibidos por el trabajador. | 

## Methods

### NewNominaJubilacionInput

`func NewNominaJubilacionInput(ingresoAcumulable float32, ingresoNoAcumulable float32, ) *NominaJubilacionInput`

NewNominaJubilacionInput instantiates a new NominaJubilacionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaJubilacionInputWithDefaults

`func NewNominaJubilacionInputWithDefaults() *NominaJubilacionInput`

NewNominaJubilacionInputWithDefaults instantiates a new NominaJubilacionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUnaExhibicion

`func (o *NominaJubilacionInput) GetTotalUnaExhibicion() float32`

GetTotalUnaExhibicion returns the TotalUnaExhibicion field if non-nil, zero value otherwise.

### GetTotalUnaExhibicionOk

`func (o *NominaJubilacionInput) GetTotalUnaExhibicionOk() (*float32, bool)`

GetTotalUnaExhibicionOk returns a tuple with the TotalUnaExhibicion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUnaExhibicion

`func (o *NominaJubilacionInput) SetTotalUnaExhibicion(v float32)`

SetTotalUnaExhibicion sets TotalUnaExhibicion field to given value.

### HasTotalUnaExhibicion

`func (o *NominaJubilacionInput) HasTotalUnaExhibicion() bool`

HasTotalUnaExhibicion returns a boolean if a field has been set.

### GetTotalParcialidad

`func (o *NominaJubilacionInput) GetTotalParcialidad() float32`

GetTotalParcialidad returns the TotalParcialidad field if non-nil, zero value otherwise.

### GetTotalParcialidadOk

`func (o *NominaJubilacionInput) GetTotalParcialidadOk() (*float32, bool)`

GetTotalParcialidadOk returns a tuple with the TotalParcialidad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParcialidad

`func (o *NominaJubilacionInput) SetTotalParcialidad(v float32)`

SetTotalParcialidad sets TotalParcialidad field to given value.

### HasTotalParcialidad

`func (o *NominaJubilacionInput) HasTotalParcialidad() bool`

HasTotalParcialidad returns a boolean if a field has been set.

### GetMontoDiario

`func (o *NominaJubilacionInput) GetMontoDiario() float32`

GetMontoDiario returns the MontoDiario field if non-nil, zero value otherwise.

### GetMontoDiarioOk

`func (o *NominaJubilacionInput) GetMontoDiarioOk() (*float32, bool)`

GetMontoDiarioOk returns a tuple with the MontoDiario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoDiario

`func (o *NominaJubilacionInput) SetMontoDiario(v float32)`

SetMontoDiario sets MontoDiario field to given value.

### HasMontoDiario

`func (o *NominaJubilacionInput) HasMontoDiario() bool`

HasMontoDiario returns a boolean if a field has been set.

### GetIngresoAcumulable

`func (o *NominaJubilacionInput) GetIngresoAcumulable() float32`

GetIngresoAcumulable returns the IngresoAcumulable field if non-nil, zero value otherwise.

### GetIngresoAcumulableOk

`func (o *NominaJubilacionInput) GetIngresoAcumulableOk() (*float32, bool)`

GetIngresoAcumulableOk returns a tuple with the IngresoAcumulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresoAcumulable

`func (o *NominaJubilacionInput) SetIngresoAcumulable(v float32)`

SetIngresoAcumulable sets IngresoAcumulable field to given value.


### GetIngresoNoAcumulable

`func (o *NominaJubilacionInput) GetIngresoNoAcumulable() float32`

GetIngresoNoAcumulable returns the IngresoNoAcumulable field if non-nil, zero value otherwise.

### GetIngresoNoAcumulableOk

`func (o *NominaJubilacionInput) GetIngresoNoAcumulableOk() (*float32, bool)`

GetIngresoNoAcumulableOk returns a tuple with the IngresoNoAcumulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresoNoAcumulable

`func (o *NominaJubilacionInput) SetIngresoNoAcumulable(v float32)`

SetIngresoNoAcumulable sets IngresoNoAcumulable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


