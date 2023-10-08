# NominaComplementDataProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoNomina** | Pointer to **string** | Tipo de nómina. - &#x60;“O”&#x60; (Ordinaria): Cuando corresponde a un pago que se realiza de manera habitual, como sueldos. - &#x60;“E”&#x60; (Extraordinaria): Para pagos fuera de lo habitual, como liquidaciones, aguinaldos o bonos.  | [optional] [default to "O"]
**FechaPago** | Pointer to **string** | Fecha de pago de la nómina al trabajador. | [optional] [default to "now"]
**FechaInicialPago** | Pointer to **string** | Fecha inicial del periodo de pago. | [optional] 
**FechaFinalPago** | Pointer to **string** | Fecha final del periodo de pago. | [optional] 
**NumDiasPagados** | Pointer to **float32** | Número de días pagados. Puede ser entero o fracción. | [optional] 
**Emisor** | Pointer to [**NominaEmisorProperties**](NominaEmisorProperties.md) |  | [optional] 
**Receptor** | Pointer to [**NominaReceptorProperties**](NominaReceptorProperties.md) |  | [optional] 
**Percepciones** | Pointer to [**NominaPercepcionesProperties**](NominaPercepcionesProperties.md) |  | [optional] 
**Deducciones** | Pointer to [**[]NominaDeduccionProperties**](NominaDeduccionProperties.md) |  | [optional] 
**OtrosPagos** | Pointer to [**[]NominaComplementDataNestedPropertiesOtrosPagosInner**](NominaComplementDataNestedPropertiesOtrosPagosInner.md) |  | [optional] 
**Incapacidades** | Pointer to [**[]NominaIncapacidadProperties**](NominaIncapacidadProperties.md) |  | [optional] 

## Methods

### NewNominaComplementDataProperties

`func NewNominaComplementDataProperties() *NominaComplementDataProperties`

NewNominaComplementDataProperties instantiates a new NominaComplementDataProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataPropertiesWithDefaults

`func NewNominaComplementDataPropertiesWithDefaults() *NominaComplementDataProperties`

NewNominaComplementDataPropertiesWithDefaults instantiates a new NominaComplementDataProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoNomina

`func (o *NominaComplementDataProperties) GetTipoNomina() string`

GetTipoNomina returns the TipoNomina field if non-nil, zero value otherwise.

### GetTipoNominaOk

`func (o *NominaComplementDataProperties) GetTipoNominaOk() (*string, bool)`

GetTipoNominaOk returns a tuple with the TipoNomina field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoNomina

`func (o *NominaComplementDataProperties) SetTipoNomina(v string)`

SetTipoNomina sets TipoNomina field to given value.

### HasTipoNomina

`func (o *NominaComplementDataProperties) HasTipoNomina() bool`

HasTipoNomina returns a boolean if a field has been set.

### GetFechaPago

`func (o *NominaComplementDataProperties) GetFechaPago() string`

GetFechaPago returns the FechaPago field if non-nil, zero value otherwise.

### GetFechaPagoOk

`func (o *NominaComplementDataProperties) GetFechaPagoOk() (*string, bool)`

GetFechaPagoOk returns a tuple with the FechaPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaPago

`func (o *NominaComplementDataProperties) SetFechaPago(v string)`

SetFechaPago sets FechaPago field to given value.

### HasFechaPago

`func (o *NominaComplementDataProperties) HasFechaPago() bool`

HasFechaPago returns a boolean if a field has been set.

### GetFechaInicialPago

`func (o *NominaComplementDataProperties) GetFechaInicialPago() string`

GetFechaInicialPago returns the FechaInicialPago field if non-nil, zero value otherwise.

### GetFechaInicialPagoOk

`func (o *NominaComplementDataProperties) GetFechaInicialPagoOk() (*string, bool)`

GetFechaInicialPagoOk returns a tuple with the FechaInicialPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaInicialPago

`func (o *NominaComplementDataProperties) SetFechaInicialPago(v string)`

SetFechaInicialPago sets FechaInicialPago field to given value.

### HasFechaInicialPago

`func (o *NominaComplementDataProperties) HasFechaInicialPago() bool`

HasFechaInicialPago returns a boolean if a field has been set.

### GetFechaFinalPago

`func (o *NominaComplementDataProperties) GetFechaFinalPago() string`

GetFechaFinalPago returns the FechaFinalPago field if non-nil, zero value otherwise.

### GetFechaFinalPagoOk

`func (o *NominaComplementDataProperties) GetFechaFinalPagoOk() (*string, bool)`

GetFechaFinalPagoOk returns a tuple with the FechaFinalPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaFinalPago

`func (o *NominaComplementDataProperties) SetFechaFinalPago(v string)`

SetFechaFinalPago sets FechaFinalPago field to given value.

### HasFechaFinalPago

`func (o *NominaComplementDataProperties) HasFechaFinalPago() bool`

HasFechaFinalPago returns a boolean if a field has been set.

### GetNumDiasPagados

`func (o *NominaComplementDataProperties) GetNumDiasPagados() float32`

GetNumDiasPagados returns the NumDiasPagados field if non-nil, zero value otherwise.

### GetNumDiasPagadosOk

`func (o *NominaComplementDataProperties) GetNumDiasPagadosOk() (*float32, bool)`

GetNumDiasPagadosOk returns a tuple with the NumDiasPagados field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDiasPagados

`func (o *NominaComplementDataProperties) SetNumDiasPagados(v float32)`

SetNumDiasPagados sets NumDiasPagados field to given value.

### HasNumDiasPagados

`func (o *NominaComplementDataProperties) HasNumDiasPagados() bool`

HasNumDiasPagados returns a boolean if a field has been set.

### GetEmisor

`func (o *NominaComplementDataProperties) GetEmisor() NominaEmisorProperties`

GetEmisor returns the Emisor field if non-nil, zero value otherwise.

### GetEmisorOk

`func (o *NominaComplementDataProperties) GetEmisorOk() (*NominaEmisorProperties, bool)`

GetEmisorOk returns a tuple with the Emisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisor

`func (o *NominaComplementDataProperties) SetEmisor(v NominaEmisorProperties)`

SetEmisor sets Emisor field to given value.

### HasEmisor

`func (o *NominaComplementDataProperties) HasEmisor() bool`

HasEmisor returns a boolean if a field has been set.

### GetReceptor

`func (o *NominaComplementDataProperties) GetReceptor() NominaReceptorProperties`

GetReceptor returns the Receptor field if non-nil, zero value otherwise.

### GetReceptorOk

`func (o *NominaComplementDataProperties) GetReceptorOk() (*NominaReceptorProperties, bool)`

GetReceptorOk returns a tuple with the Receptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptor

`func (o *NominaComplementDataProperties) SetReceptor(v NominaReceptorProperties)`

SetReceptor sets Receptor field to given value.

### HasReceptor

`func (o *NominaComplementDataProperties) HasReceptor() bool`

HasReceptor returns a boolean if a field has been set.

### GetPercepciones

`func (o *NominaComplementDataProperties) GetPercepciones() NominaPercepcionesProperties`

GetPercepciones returns the Percepciones field if non-nil, zero value otherwise.

### GetPercepcionesOk

`func (o *NominaComplementDataProperties) GetPercepcionesOk() (*NominaPercepcionesProperties, bool)`

GetPercepcionesOk returns a tuple with the Percepciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercepciones

`func (o *NominaComplementDataProperties) SetPercepciones(v NominaPercepcionesProperties)`

SetPercepciones sets Percepciones field to given value.

### HasPercepciones

`func (o *NominaComplementDataProperties) HasPercepciones() bool`

HasPercepciones returns a boolean if a field has been set.

### GetDeducciones

`func (o *NominaComplementDataProperties) GetDeducciones() []NominaDeduccionProperties`

GetDeducciones returns the Deducciones field if non-nil, zero value otherwise.

### GetDeduccionesOk

`func (o *NominaComplementDataProperties) GetDeduccionesOk() (*[]NominaDeduccionProperties, bool)`

GetDeduccionesOk returns a tuple with the Deducciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeducciones

`func (o *NominaComplementDataProperties) SetDeducciones(v []NominaDeduccionProperties)`

SetDeducciones sets Deducciones field to given value.

### HasDeducciones

`func (o *NominaComplementDataProperties) HasDeducciones() bool`

HasDeducciones returns a boolean if a field has been set.

### GetOtrosPagos

`func (o *NominaComplementDataProperties) GetOtrosPagos() []NominaComplementDataNestedPropertiesOtrosPagosInner`

GetOtrosPagos returns the OtrosPagos field if non-nil, zero value otherwise.

### GetOtrosPagosOk

`func (o *NominaComplementDataProperties) GetOtrosPagosOk() (*[]NominaComplementDataNestedPropertiesOtrosPagosInner, bool)`

GetOtrosPagosOk returns a tuple with the OtrosPagos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtrosPagos

`func (o *NominaComplementDataProperties) SetOtrosPagos(v []NominaComplementDataNestedPropertiesOtrosPagosInner)`

SetOtrosPagos sets OtrosPagos field to given value.

### HasOtrosPagos

`func (o *NominaComplementDataProperties) HasOtrosPagos() bool`

HasOtrosPagos returns a boolean if a field has been set.

### GetIncapacidades

`func (o *NominaComplementDataProperties) GetIncapacidades() []NominaIncapacidadProperties`

GetIncapacidades returns the Incapacidades field if non-nil, zero value otherwise.

### GetIncapacidadesOk

`func (o *NominaComplementDataProperties) GetIncapacidadesOk() (*[]NominaIncapacidadProperties, bool)`

GetIncapacidadesOk returns a tuple with the Incapacidades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncapacidades

`func (o *NominaComplementDataProperties) SetIncapacidades(v []NominaIncapacidadProperties)`

SetIncapacidades sets Incapacidades field to given value.

### HasIncapacidades

`func (o *NominaComplementDataProperties) HasIncapacidades() bool`

HasIncapacidades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


