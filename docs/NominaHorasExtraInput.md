# NominaHorasExtraInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dias** | **int32** | Número de días en que el trabajador laboró horas extra adicionales a su jornada normal de trabajo. | 
**TipoHoras** | **string** | Clave del catálogo [Tipo de Horas](#tipo-de-Horas). | 
**HorasExtra** | **int32** | Número de horas extra trabajadas en el periodo. | 
**ImportePagado** | **float32** | Importe pagado por las horas extra. | 

## Methods

### NewNominaHorasExtraInput

`func NewNominaHorasExtraInput(dias int32, tipoHoras string, horasExtra int32, importePagado float32, ) *NominaHorasExtraInput`

NewNominaHorasExtraInput instantiates a new NominaHorasExtraInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaHorasExtraInputWithDefaults

`func NewNominaHorasExtraInputWithDefaults() *NominaHorasExtraInput`

NewNominaHorasExtraInputWithDefaults instantiates a new NominaHorasExtraInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDias

`func (o *NominaHorasExtraInput) GetDias() int32`

GetDias returns the Dias field if non-nil, zero value otherwise.

### GetDiasOk

`func (o *NominaHorasExtraInput) GetDiasOk() (*int32, bool)`

GetDiasOk returns a tuple with the Dias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDias

`func (o *NominaHorasExtraInput) SetDias(v int32)`

SetDias sets Dias field to given value.


### GetTipoHoras

`func (o *NominaHorasExtraInput) GetTipoHoras() string`

GetTipoHoras returns the TipoHoras field if non-nil, zero value otherwise.

### GetTipoHorasOk

`func (o *NominaHorasExtraInput) GetTipoHorasOk() (*string, bool)`

GetTipoHorasOk returns a tuple with the TipoHoras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoHoras

`func (o *NominaHorasExtraInput) SetTipoHoras(v string)`

SetTipoHoras sets TipoHoras field to given value.


### GetHorasExtra

`func (o *NominaHorasExtraInput) GetHorasExtra() int32`

GetHorasExtra returns the HorasExtra field if non-nil, zero value otherwise.

### GetHorasExtraOk

`func (o *NominaHorasExtraInput) GetHorasExtraOk() (*int32, bool)`

GetHorasExtraOk returns a tuple with the HorasExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorasExtra

`func (o *NominaHorasExtraInput) SetHorasExtra(v int32)`

SetHorasExtra sets HorasExtra field to given value.


### GetImportePagado

`func (o *NominaHorasExtraInput) GetImportePagado() float32`

GetImportePagado returns the ImportePagado field if non-nil, zero value otherwise.

### GetImportePagadoOk

`func (o *NominaHorasExtraInput) GetImportePagadoOk() (*float32, bool)`

GetImportePagadoOk returns a tuple with the ImportePagado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportePagado

`func (o *NominaHorasExtraInput) SetImportePagado(v float32)`

SetImportePagado sets ImportePagado field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


