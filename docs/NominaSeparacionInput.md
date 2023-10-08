# NominaSeparacionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPagado** | **float32** | Monto total pagado por concepto de separación o indemnización. | 
**NumAnosServicio** | **int32** | Años de servicio que laboró el trabajador, redondeado al entero inmediato superior. | 
**UltimoSueldoMensOrd** | **float32** | Último sueldo mensual ordinario percibido por el trabajador. | 
**IngresoAcumulable** | **float32** | Monto por ingresos acumulables. | 
**IngresoNoAcumulable** | **float32** | Monto por ingresos no acumulables. | 

## Methods

### NewNominaSeparacionInput

`func NewNominaSeparacionInput(totalPagado float32, numAnosServicio int32, ultimoSueldoMensOrd float32, ingresoAcumulable float32, ingresoNoAcumulable float32, ) *NominaSeparacionInput`

NewNominaSeparacionInput instantiates a new NominaSeparacionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaSeparacionInputWithDefaults

`func NewNominaSeparacionInputWithDefaults() *NominaSeparacionInput`

NewNominaSeparacionInputWithDefaults instantiates a new NominaSeparacionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPagado

`func (o *NominaSeparacionInput) GetTotalPagado() float32`

GetTotalPagado returns the TotalPagado field if non-nil, zero value otherwise.

### GetTotalPagadoOk

`func (o *NominaSeparacionInput) GetTotalPagadoOk() (*float32, bool)`

GetTotalPagadoOk returns a tuple with the TotalPagado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPagado

`func (o *NominaSeparacionInput) SetTotalPagado(v float32)`

SetTotalPagado sets TotalPagado field to given value.


### GetNumAnosServicio

`func (o *NominaSeparacionInput) GetNumAnosServicio() int32`

GetNumAnosServicio returns the NumAnosServicio field if non-nil, zero value otherwise.

### GetNumAnosServicioOk

`func (o *NominaSeparacionInput) GetNumAnosServicioOk() (*int32, bool)`

GetNumAnosServicioOk returns a tuple with the NumAnosServicio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAnosServicio

`func (o *NominaSeparacionInput) SetNumAnosServicio(v int32)`

SetNumAnosServicio sets NumAnosServicio field to given value.


### GetUltimoSueldoMensOrd

`func (o *NominaSeparacionInput) GetUltimoSueldoMensOrd() float32`

GetUltimoSueldoMensOrd returns the UltimoSueldoMensOrd field if non-nil, zero value otherwise.

### GetUltimoSueldoMensOrdOk

`func (o *NominaSeparacionInput) GetUltimoSueldoMensOrdOk() (*float32, bool)`

GetUltimoSueldoMensOrdOk returns a tuple with the UltimoSueldoMensOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltimoSueldoMensOrd

`func (o *NominaSeparacionInput) SetUltimoSueldoMensOrd(v float32)`

SetUltimoSueldoMensOrd sets UltimoSueldoMensOrd field to given value.


### GetIngresoAcumulable

`func (o *NominaSeparacionInput) GetIngresoAcumulable() float32`

GetIngresoAcumulable returns the IngresoAcumulable field if non-nil, zero value otherwise.

### GetIngresoAcumulableOk

`func (o *NominaSeparacionInput) GetIngresoAcumulableOk() (*float32, bool)`

GetIngresoAcumulableOk returns a tuple with the IngresoAcumulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresoAcumulable

`func (o *NominaSeparacionInput) SetIngresoAcumulable(v float32)`

SetIngresoAcumulable sets IngresoAcumulable field to given value.


### GetIngresoNoAcumulable

`func (o *NominaSeparacionInput) GetIngresoNoAcumulable() float32`

GetIngresoNoAcumulable returns the IngresoNoAcumulable field if non-nil, zero value otherwise.

### GetIngresoNoAcumulableOk

`func (o *NominaSeparacionInput) GetIngresoNoAcumulableOk() (*float32, bool)`

GetIngresoNoAcumulableOk returns a tuple with the IngresoNoAcumulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresoNoAcumulable

`func (o *NominaSeparacionInput) SetIngresoNoAcumulable(v float32)`

SetIngresoNoAcumulable sets IngresoNoAcumulable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


