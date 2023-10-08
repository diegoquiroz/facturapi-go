# NominaComplementDataNestedProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emisor** | Pointer to [**NominaEmisorProperties**](NominaEmisorProperties.md) |  | [optional] 
**Receptor** | Pointer to [**NominaReceptorProperties**](NominaReceptorProperties.md) |  | [optional] 
**Percepciones** | Pointer to [**NominaPercepcionesProperties**](NominaPercepcionesProperties.md) |  | [optional] 
**Deducciones** | Pointer to [**[]NominaDeduccionProperties**](NominaDeduccionProperties.md) |  | [optional] 
**OtrosPagos** | Pointer to [**[]NominaComplementDataNestedPropertiesOtrosPagosInner**](NominaComplementDataNestedPropertiesOtrosPagosInner.md) |  | [optional] 
**Incapacidades** | Pointer to [**[]NominaIncapacidadProperties**](NominaIncapacidadProperties.md) |  | [optional] 

## Methods

### NewNominaComplementDataNestedProperties

`func NewNominaComplementDataNestedProperties() *NominaComplementDataNestedProperties`

NewNominaComplementDataNestedProperties instantiates a new NominaComplementDataNestedProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataNestedPropertiesWithDefaults

`func NewNominaComplementDataNestedPropertiesWithDefaults() *NominaComplementDataNestedProperties`

NewNominaComplementDataNestedPropertiesWithDefaults instantiates a new NominaComplementDataNestedProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmisor

`func (o *NominaComplementDataNestedProperties) GetEmisor() NominaEmisorProperties`

GetEmisor returns the Emisor field if non-nil, zero value otherwise.

### GetEmisorOk

`func (o *NominaComplementDataNestedProperties) GetEmisorOk() (*NominaEmisorProperties, bool)`

GetEmisorOk returns a tuple with the Emisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisor

`func (o *NominaComplementDataNestedProperties) SetEmisor(v NominaEmisorProperties)`

SetEmisor sets Emisor field to given value.

### HasEmisor

`func (o *NominaComplementDataNestedProperties) HasEmisor() bool`

HasEmisor returns a boolean if a field has been set.

### GetReceptor

`func (o *NominaComplementDataNestedProperties) GetReceptor() NominaReceptorProperties`

GetReceptor returns the Receptor field if non-nil, zero value otherwise.

### GetReceptorOk

`func (o *NominaComplementDataNestedProperties) GetReceptorOk() (*NominaReceptorProperties, bool)`

GetReceptorOk returns a tuple with the Receptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptor

`func (o *NominaComplementDataNestedProperties) SetReceptor(v NominaReceptorProperties)`

SetReceptor sets Receptor field to given value.

### HasReceptor

`func (o *NominaComplementDataNestedProperties) HasReceptor() bool`

HasReceptor returns a boolean if a field has been set.

### GetPercepciones

`func (o *NominaComplementDataNestedProperties) GetPercepciones() NominaPercepcionesProperties`

GetPercepciones returns the Percepciones field if non-nil, zero value otherwise.

### GetPercepcionesOk

`func (o *NominaComplementDataNestedProperties) GetPercepcionesOk() (*NominaPercepcionesProperties, bool)`

GetPercepcionesOk returns a tuple with the Percepciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercepciones

`func (o *NominaComplementDataNestedProperties) SetPercepciones(v NominaPercepcionesProperties)`

SetPercepciones sets Percepciones field to given value.

### HasPercepciones

`func (o *NominaComplementDataNestedProperties) HasPercepciones() bool`

HasPercepciones returns a boolean if a field has been set.

### GetDeducciones

`func (o *NominaComplementDataNestedProperties) GetDeducciones() []NominaDeduccionProperties`

GetDeducciones returns the Deducciones field if non-nil, zero value otherwise.

### GetDeduccionesOk

`func (o *NominaComplementDataNestedProperties) GetDeduccionesOk() (*[]NominaDeduccionProperties, bool)`

GetDeduccionesOk returns a tuple with the Deducciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeducciones

`func (o *NominaComplementDataNestedProperties) SetDeducciones(v []NominaDeduccionProperties)`

SetDeducciones sets Deducciones field to given value.

### HasDeducciones

`func (o *NominaComplementDataNestedProperties) HasDeducciones() bool`

HasDeducciones returns a boolean if a field has been set.

### GetOtrosPagos

`func (o *NominaComplementDataNestedProperties) GetOtrosPagos() []NominaComplementDataNestedPropertiesOtrosPagosInner`

GetOtrosPagos returns the OtrosPagos field if non-nil, zero value otherwise.

### GetOtrosPagosOk

`func (o *NominaComplementDataNestedProperties) GetOtrosPagosOk() (*[]NominaComplementDataNestedPropertiesOtrosPagosInner, bool)`

GetOtrosPagosOk returns a tuple with the OtrosPagos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtrosPagos

`func (o *NominaComplementDataNestedProperties) SetOtrosPagos(v []NominaComplementDataNestedPropertiesOtrosPagosInner)`

SetOtrosPagos sets OtrosPagos field to given value.

### HasOtrosPagos

`func (o *NominaComplementDataNestedProperties) HasOtrosPagos() bool`

HasOtrosPagos returns a boolean if a field has been set.

### GetIncapacidades

`func (o *NominaComplementDataNestedProperties) GetIncapacidades() []NominaIncapacidadProperties`

GetIncapacidades returns the Incapacidades field if non-nil, zero value otherwise.

### GetIncapacidadesOk

`func (o *NominaComplementDataNestedProperties) GetIncapacidadesOk() (*[]NominaIncapacidadProperties, bool)`

GetIncapacidadesOk returns a tuple with the Incapacidades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncapacidades

`func (o *NominaComplementDataNestedProperties) SetIncapacidades(v []NominaIncapacidadProperties)`

SetIncapacidades sets Incapacidades field to given value.

### HasIncapacidades

`func (o *NominaComplementDataNestedProperties) HasIncapacidades() bool`

HasIncapacidades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


