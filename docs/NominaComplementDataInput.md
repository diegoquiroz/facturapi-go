# NominaComplementDataInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoNomina** | Pointer to **string** | Tipo de nómina. - &#x60;“O”&#x60; (Ordinaria): Cuando corresponde a un pago que se realiza de manera habitual, como sueldos. - &#x60;“E”&#x60; (Extraordinaria): Para pagos fuera de lo habitual, como liquidaciones, aguinaldos o bonos.  | [optional] [default to "O"]
**FechaPago** | Pointer to **string** | Fecha de pago de la nómina al trabajador. | [optional] [default to "now"]
**FechaInicialPago** | **string** | Fecha inicial del periodo de pago. | 
**FechaFinalPago** | **string** | Fecha final del periodo de pago. | 
**NumDiasPagados** | **float32** | Número de días pagados. Puede ser entero o fracción. | 
**Emisor** | Pointer to [**NominaEmisorProperties**](NominaEmisorProperties.md) |  | [optional] 
**Receptor** | [**NominaReceptorInput**](NominaReceptorInput.md) |  | 
**Percepciones** | [**NominaPercepcionesInput**](NominaPercepcionesInput.md) |  | 
**Deducciones** | Pointer to [**[]NominaDeduccionInput**](NominaDeduccionInput.md) |  | [optional] 
**OtrosPagos** | Pointer to [**[]OtroPago**](OtroPago.md) |  | [optional] 
**Incapacidades** | Pointer to [**[]NominaIncapacidadInput**](NominaIncapacidadInput.md) |  | [optional] 

## Methods

### NewNominaComplementDataInput

`func NewNominaComplementDataInput(fechaInicialPago string, fechaFinalPago string, numDiasPagados float32, receptor NominaReceptorInput, percepciones NominaPercepcionesInput, ) *NominaComplementDataInput`

NewNominaComplementDataInput instantiates a new NominaComplementDataInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataInputWithDefaults

`func NewNominaComplementDataInputWithDefaults() *NominaComplementDataInput`

NewNominaComplementDataInputWithDefaults instantiates a new NominaComplementDataInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoNomina

`func (o *NominaComplementDataInput) GetTipoNomina() string`

GetTipoNomina returns the TipoNomina field if non-nil, zero value otherwise.

### GetTipoNominaOk

`func (o *NominaComplementDataInput) GetTipoNominaOk() (*string, bool)`

GetTipoNominaOk returns a tuple with the TipoNomina field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoNomina

`func (o *NominaComplementDataInput) SetTipoNomina(v string)`

SetTipoNomina sets TipoNomina field to given value.

### HasTipoNomina

`func (o *NominaComplementDataInput) HasTipoNomina() bool`

HasTipoNomina returns a boolean if a field has been set.

### GetFechaPago

`func (o *NominaComplementDataInput) GetFechaPago() string`

GetFechaPago returns the FechaPago field if non-nil, zero value otherwise.

### GetFechaPagoOk

`func (o *NominaComplementDataInput) GetFechaPagoOk() (*string, bool)`

GetFechaPagoOk returns a tuple with the FechaPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaPago

`func (o *NominaComplementDataInput) SetFechaPago(v string)`

SetFechaPago sets FechaPago field to given value.

### HasFechaPago

`func (o *NominaComplementDataInput) HasFechaPago() bool`

HasFechaPago returns a boolean if a field has been set.

### GetFechaInicialPago

`func (o *NominaComplementDataInput) GetFechaInicialPago() string`

GetFechaInicialPago returns the FechaInicialPago field if non-nil, zero value otherwise.

### GetFechaInicialPagoOk

`func (o *NominaComplementDataInput) GetFechaInicialPagoOk() (*string, bool)`

GetFechaInicialPagoOk returns a tuple with the FechaInicialPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaInicialPago

`func (o *NominaComplementDataInput) SetFechaInicialPago(v string)`

SetFechaInicialPago sets FechaInicialPago field to given value.


### GetFechaFinalPago

`func (o *NominaComplementDataInput) GetFechaFinalPago() string`

GetFechaFinalPago returns the FechaFinalPago field if non-nil, zero value otherwise.

### GetFechaFinalPagoOk

`func (o *NominaComplementDataInput) GetFechaFinalPagoOk() (*string, bool)`

GetFechaFinalPagoOk returns a tuple with the FechaFinalPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaFinalPago

`func (o *NominaComplementDataInput) SetFechaFinalPago(v string)`

SetFechaFinalPago sets FechaFinalPago field to given value.


### GetNumDiasPagados

`func (o *NominaComplementDataInput) GetNumDiasPagados() float32`

GetNumDiasPagados returns the NumDiasPagados field if non-nil, zero value otherwise.

### GetNumDiasPagadosOk

`func (o *NominaComplementDataInput) GetNumDiasPagadosOk() (*float32, bool)`

GetNumDiasPagadosOk returns a tuple with the NumDiasPagados field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDiasPagados

`func (o *NominaComplementDataInput) SetNumDiasPagados(v float32)`

SetNumDiasPagados sets NumDiasPagados field to given value.


### GetEmisor

`func (o *NominaComplementDataInput) GetEmisor() NominaEmisorProperties`

GetEmisor returns the Emisor field if non-nil, zero value otherwise.

### GetEmisorOk

`func (o *NominaComplementDataInput) GetEmisorOk() (*NominaEmisorProperties, bool)`

GetEmisorOk returns a tuple with the Emisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmisor

`func (o *NominaComplementDataInput) SetEmisor(v NominaEmisorProperties)`

SetEmisor sets Emisor field to given value.

### HasEmisor

`func (o *NominaComplementDataInput) HasEmisor() bool`

HasEmisor returns a boolean if a field has been set.

### GetReceptor

`func (o *NominaComplementDataInput) GetReceptor() NominaReceptorInput`

GetReceptor returns the Receptor field if non-nil, zero value otherwise.

### GetReceptorOk

`func (o *NominaComplementDataInput) GetReceptorOk() (*NominaReceptorInput, bool)`

GetReceptorOk returns a tuple with the Receptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptor

`func (o *NominaComplementDataInput) SetReceptor(v NominaReceptorInput)`

SetReceptor sets Receptor field to given value.


### GetPercepciones

`func (o *NominaComplementDataInput) GetPercepciones() NominaPercepcionesInput`

GetPercepciones returns the Percepciones field if non-nil, zero value otherwise.

### GetPercepcionesOk

`func (o *NominaComplementDataInput) GetPercepcionesOk() (*NominaPercepcionesInput, bool)`

GetPercepcionesOk returns a tuple with the Percepciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercepciones

`func (o *NominaComplementDataInput) SetPercepciones(v NominaPercepcionesInput)`

SetPercepciones sets Percepciones field to given value.


### GetDeducciones

`func (o *NominaComplementDataInput) GetDeducciones() []NominaDeduccionInput`

GetDeducciones returns the Deducciones field if non-nil, zero value otherwise.

### GetDeduccionesOk

`func (o *NominaComplementDataInput) GetDeduccionesOk() (*[]NominaDeduccionInput, bool)`

GetDeduccionesOk returns a tuple with the Deducciones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeducciones

`func (o *NominaComplementDataInput) SetDeducciones(v []NominaDeduccionInput)`

SetDeducciones sets Deducciones field to given value.

### HasDeducciones

`func (o *NominaComplementDataInput) HasDeducciones() bool`

HasDeducciones returns a boolean if a field has been set.

### GetOtrosPagos

`func (o *NominaComplementDataInput) GetOtrosPagos() []OtroPago`

GetOtrosPagos returns the OtrosPagos field if non-nil, zero value otherwise.

### GetOtrosPagosOk

`func (o *NominaComplementDataInput) GetOtrosPagosOk() (*[]OtroPago, bool)`

GetOtrosPagosOk returns a tuple with the OtrosPagos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtrosPagos

`func (o *NominaComplementDataInput) SetOtrosPagos(v []OtroPago)`

SetOtrosPagos sets OtrosPagos field to given value.

### HasOtrosPagos

`func (o *NominaComplementDataInput) HasOtrosPagos() bool`

HasOtrosPagos returns a boolean if a field has been set.

### GetIncapacidades

`func (o *NominaComplementDataInput) GetIncapacidades() []NominaIncapacidadInput`

GetIncapacidades returns the Incapacidades field if non-nil, zero value otherwise.

### GetIncapacidadesOk

`func (o *NominaComplementDataInput) GetIncapacidadesOk() (*[]NominaIncapacidadInput, bool)`

GetIncapacidadesOk returns a tuple with the Incapacidades field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncapacidades

`func (o *NominaComplementDataInput) SetIncapacidades(v []NominaIncapacidadInput)`

SetIncapacidades sets Incapacidades field to given value.

### HasIncapacidades

`func (o *NominaComplementDataInput) HasIncapacidades() bool`

HasIncapacidades returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


