# NominaHorasExtraProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dias** | Pointer to **int32** | Número de días en que el trabajador laboró horas extra adicionales a su jornada normal de trabajo. | [optional] 
**TipoHoras** | Pointer to **string** | Clave del catálogo [Tipo de Horas](#tipo-de-Horas). | [optional] 
**HorasExtra** | Pointer to **int32** | Número de horas extra trabajadas en el periodo. | [optional] 
**ImportePagado** | Pointer to **float32** | Importe pagado por las horas extra. | [optional] 

## Methods

### NewNominaHorasExtraProperties

`func NewNominaHorasExtraProperties() *NominaHorasExtraProperties`

NewNominaHorasExtraProperties instantiates a new NominaHorasExtraProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaHorasExtraPropertiesWithDefaults

`func NewNominaHorasExtraPropertiesWithDefaults() *NominaHorasExtraProperties`

NewNominaHorasExtraPropertiesWithDefaults instantiates a new NominaHorasExtraProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDias

`func (o *NominaHorasExtraProperties) GetDias() int32`

GetDias returns the Dias field if non-nil, zero value otherwise.

### GetDiasOk

`func (o *NominaHorasExtraProperties) GetDiasOk() (*int32, bool)`

GetDiasOk returns a tuple with the Dias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDias

`func (o *NominaHorasExtraProperties) SetDias(v int32)`

SetDias sets Dias field to given value.

### HasDias

`func (o *NominaHorasExtraProperties) HasDias() bool`

HasDias returns a boolean if a field has been set.

### GetTipoHoras

`func (o *NominaHorasExtraProperties) GetTipoHoras() string`

GetTipoHoras returns the TipoHoras field if non-nil, zero value otherwise.

### GetTipoHorasOk

`func (o *NominaHorasExtraProperties) GetTipoHorasOk() (*string, bool)`

GetTipoHorasOk returns a tuple with the TipoHoras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoHoras

`func (o *NominaHorasExtraProperties) SetTipoHoras(v string)`

SetTipoHoras sets TipoHoras field to given value.

### HasTipoHoras

`func (o *NominaHorasExtraProperties) HasTipoHoras() bool`

HasTipoHoras returns a boolean if a field has been set.

### GetHorasExtra

`func (o *NominaHorasExtraProperties) GetHorasExtra() int32`

GetHorasExtra returns the HorasExtra field if non-nil, zero value otherwise.

### GetHorasExtraOk

`func (o *NominaHorasExtraProperties) GetHorasExtraOk() (*int32, bool)`

GetHorasExtraOk returns a tuple with the HorasExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorasExtra

`func (o *NominaHorasExtraProperties) SetHorasExtra(v int32)`

SetHorasExtra sets HorasExtra field to given value.

### HasHorasExtra

`func (o *NominaHorasExtraProperties) HasHorasExtra() bool`

HasHorasExtra returns a boolean if a field has been set.

### GetImportePagado

`func (o *NominaHorasExtraProperties) GetImportePagado() float32`

GetImportePagado returns the ImportePagado field if non-nil, zero value otherwise.

### GetImportePagadoOk

`func (o *NominaHorasExtraProperties) GetImportePagadoOk() (*float32, bool)`

GetImportePagadoOk returns a tuple with the ImportePagado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportePagado

`func (o *NominaHorasExtraProperties) SetImportePagado(v float32)`

SetImportePagado sets ImportePagado field to given value.

### HasImportePagado

`func (o *NominaHorasExtraProperties) HasImportePagado() bool`

HasImportePagado returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


