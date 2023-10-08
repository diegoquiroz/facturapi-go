# NominaComplementDataNestedInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emisor** | Pointer to [**NominaEmisorProperties**](NominaEmisorProperties.md) |  | [optional] 
**Receptor** | Pointer to [**NominaReceptorInput**](NominaReceptorInput.md) |  | [optional] 
**Percepciones** | Pointer to [**NominaPercepcionesInput**](NominaPercepcionesInput.md) |  | [optional] 
**Deducciones** | Pointer to [**[]NominaDeduccionInput**](NominaDeduccionInput.md) |  | [optional] 
**OtrosPagos** | Pointer to [**[]OtroPago**](OtroPago.md) |  | [optional] 
**Incapacidades** | Pointer to [**[]NominaIncapacidadInput**](NominaIncapacidadInput.md) |  | [optional] 

## Methods

### NewNominaComplementDataNestedInput

`func NewNominaComplementDataNestedInput() *NominaComplementDataNestedInput`

NewNominaComplementDataNestedInput instantiates a new NominaComplementDataNestedInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataNestedInputWithDefaults

`func NewNominaComplementDataNestedInputWithDefaults() *NominaComplementDataNestedInput`

NewNominaComplementDataNestedInputWithDefaults instantiates a new NominaComplementDataNestedInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmisor

`func (o *NominaComplementDataNestedInput) GetEmisor() NominaEmisorProperties`

GetEmisor returns the Emisor field if non-nil, zero value otherwise.

### GetEmisorOk

`func (o *NominaComplementDataNestedInput) GetEmisorOk() (*NominaEmisorProperties, bool)`

GetEmisorOk returns a tuple with the Emisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisor

`func (o *NominaComplementDataNestedInput) SetEmisor(v NominaEmisorProperties)`

SetEmisor sets Emisor field to given value.

### HasEmisor

`func (o *NominaComplementDataNestedInput) HasEmisor() bool`

HasEmisor returns a boolean if a field has been set.

### GetReceptor

`func (o *NominaComplementDataNestedInput) GetReceptor() NominaReceptorInput`

GetReceptor returns the Receptor field if non-nil, zero value otherwise.

### GetReceptorOk

`func (o *NominaComplementDataNestedInput) GetReceptorOk() (*NominaReceptorInput, bool)`

GetReceptorOk returns a tuple with the Receptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptor

`func (o *NominaComplementDataNestedInput) SetReceptor(v NominaReceptorInput)`

SetReceptor sets Receptor field to given value.

### HasReceptor

`func (o *NominaComplementDataNestedInput) HasReceptor() bool`

HasReceptor returns a boolean if a field has been set.

### GetPercepciones

`func (o *NominaComplementDataNestedInput) GetPercepciones() NominaPercepcionesInput`

GetPercepciones returns the Percepciones field if non-nil, zero value otherwise.

### GetPercepcionesOk

`func (o *NominaComplementDataNestedInput) GetPercepcionesOk() (*NominaPercepcionesInput, bool)`

GetPercepcionesOk returns a tuple with the Percepciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercepciones

`func (o *NominaComplementDataNestedInput) SetPercepciones(v NominaPercepcionesInput)`

SetPercepciones sets Percepciones field to given value.

### HasPercepciones

`func (o *NominaComplementDataNestedInput) HasPercepciones() bool`

HasPercepciones returns a boolean if a field has been set.

### GetDeducciones

`func (o *NominaComplementDataNestedInput) GetDeducciones() []NominaDeduccionInput`

GetDeducciones returns the Deducciones field if non-nil, zero value otherwise.

### GetDeduccionesOk

`func (o *NominaComplementDataNestedInput) GetDeduccionesOk() (*[]NominaDeduccionInput, bool)`

GetDeduccionesOk returns a tuple with the Deducciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeducciones

`func (o *NominaComplementDataNestedInput) SetDeducciones(v []NominaDeduccionInput)`

SetDeducciones sets Deducciones field to given value.

### HasDeducciones

`func (o *NominaComplementDataNestedInput) HasDeducciones() bool`

HasDeducciones returns a boolean if a field has been set.

### GetOtrosPagos

`func (o *NominaComplementDataNestedInput) GetOtrosPagos() []OtroPago`

GetOtrosPagos returns the OtrosPagos field if non-nil, zero value otherwise.

### GetOtrosPagosOk

`func (o *NominaComplementDataNestedInput) GetOtrosPagosOk() (*[]OtroPago, bool)`

GetOtrosPagosOk returns a tuple with the OtrosPagos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtrosPagos

`func (o *NominaComplementDataNestedInput) SetOtrosPagos(v []OtroPago)`

SetOtrosPagos sets OtrosPagos field to given value.

### HasOtrosPagos

`func (o *NominaComplementDataNestedInput) HasOtrosPagos() bool`

HasOtrosPagos returns a boolean if a field has been set.

### GetIncapacidades

`func (o *NominaComplementDataNestedInput) GetIncapacidades() []NominaIncapacidadInput`

GetIncapacidades returns the Incapacidades field if non-nil, zero value otherwise.

### GetIncapacidadesOk

`func (o *NominaComplementDataNestedInput) GetIncapacidadesOk() (*[]NominaIncapacidadInput, bool)`

GetIncapacidadesOk returns a tuple with the Incapacidades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncapacidades

`func (o *NominaComplementDataNestedInput) SetIncapacidades(v []NominaIncapacidadInput)`

SetIncapacidades sets Incapacidades field to given value.

### HasIncapacidades

`func (o *NominaComplementDataNestedInput) HasIncapacidades() bool`

HasIncapacidades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


