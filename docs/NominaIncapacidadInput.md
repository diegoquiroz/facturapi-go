# NominaIncapacidadInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiasIncapacidad** | **int32** | Número de días enteros que el trabajador se incapacitó en el periodo. | 
**TipoIncapacidad** | **string** | Clave del catálogo [Tipo de Incapacidad](#tipo-de-incapacidad). | 
**ImporteMonetario** | Pointer to **float32** | Monto del importe monetario de la incapacidad. | [optional] 

## Methods

### NewNominaIncapacidadInput

`func NewNominaIncapacidadInput(diasIncapacidad int32, tipoIncapacidad string, ) *NominaIncapacidadInput`

NewNominaIncapacidadInput instantiates a new NominaIncapacidadInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaIncapacidadInputWithDefaults

`func NewNominaIncapacidadInputWithDefaults() *NominaIncapacidadInput`

NewNominaIncapacidadInputWithDefaults instantiates a new NominaIncapacidadInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiasIncapacidad

`func (o *NominaIncapacidadInput) GetDiasIncapacidad() int32`

GetDiasIncapacidad returns the DiasIncapacidad field if non-nil, zero value otherwise.

### GetDiasIncapacidadOk

`func (o *NominaIncapacidadInput) GetDiasIncapacidadOk() (*int32, bool)`

GetDiasIncapacidadOk returns a tuple with the DiasIncapacidad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiasIncapacidad

`func (o *NominaIncapacidadInput) SetDiasIncapacidad(v int32)`

SetDiasIncapacidad sets DiasIncapacidad field to given value.


### GetTipoIncapacidad

`func (o *NominaIncapacidadInput) GetTipoIncapacidad() string`

GetTipoIncapacidad returns the TipoIncapacidad field if non-nil, zero value otherwise.

### GetTipoIncapacidadOk

`func (o *NominaIncapacidadInput) GetTipoIncapacidadOk() (*string, bool)`

GetTipoIncapacidadOk returns a tuple with the TipoIncapacidad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoIncapacidad

`func (o *NominaIncapacidadInput) SetTipoIncapacidad(v string)`

SetTipoIncapacidad sets TipoIncapacidad field to given value.


### GetImporteMonetario

`func (o *NominaIncapacidadInput) GetImporteMonetario() float32`

GetImporteMonetario returns the ImporteMonetario field if non-nil, zero value otherwise.

### GetImporteMonetarioOk

`func (o *NominaIncapacidadInput) GetImporteMonetarioOk() (*float32, bool)`

GetImporteMonetarioOk returns a tuple with the ImporteMonetario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporteMonetario

`func (o *NominaIncapacidadInput) SetImporteMonetario(v float32)`

SetImporteMonetario sets ImporteMonetario field to given value.

### HasImporteMonetario

`func (o *NominaIncapacidadInput) HasImporteMonetario() bool`

HasImporteMonetario returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


