# NominaReceptorInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Curp** | **string** | CURP del trabajador. | 
**NumSeguridadSocial** | Pointer to **string** | Número de seguridad social. | [optional] 
**FechaInicioRelLaboral** | Pointer to **string** | Fecha de inicio de la relación laboral entre el empleador y el empleado. | [optional] 
**Antiguedad** | Pointer to [**NominaReceptorDirectPropertiesAntiguedad**](NominaReceptorDirectPropertiesAntiguedad.md) |  | [optional] [default to true]
**TipoContrato** | **string** | Clave del catálogo del SAT [Tipo de Contrato](#tipo-de-contrato). | 
**Sindicalizado** | Pointer to **bool** | Indica si el trabajador está asociado a un sindicato. | [optional] [default to false]
**TipoJornada** | Pointer to **string** | Clave del catálogo del SAT [Tipo de Jornada](#tipo-de-jornada). | [optional] 
**TipoRegimen** | **string** | Clave del catálogo del SAT [Tipo de Régimen](#tipo-de-régimen). | 
**NumEmpleado** | **string** | Número interno de empleado, asignado por el empleador. | 
**Departamento** | Pointer to **string** | Nombre del departamento o área a la que pertenece el trabajador. | [optional] 
**Puesto** | Pointer to **string** | Nombre del puesto asignado al empleado o el nombre de la actividad que realiza. | [optional] 
**RiesgoPuesto** | Pointer to **string** | Clave del catálogo del SAT [Riesgo del Puesto](#riesgo-del-puesto). | [optional] 
**PeriodicidadPago** | **string** | Clave del catálogo del SAT [Periodicidad de Pago](#periodicidad-del-pago). | 
**Banco** | Pointer to **string** | Clave del banco de acuerdo al catálogo del SAT “Bancos” que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/bank). | [optional] 
**CuentaBancaria** | Pointer to **string** | Número de cuenta bancaria (11 caracteres) o número de teléfono celular (10 caracteres) o número de tarjeta (15 ó 16 caracteres) o la CLABE (18 caracteres) o número de monedero electrónico donde se realiza el depósito de nómina.  | [optional] 
**SalarioBaseCotApor** | Pointer to **float32** | Importe de la retribución en efectivo por cuota diaria, gratificaciones, percepciones, alimentación, habitación, primas, comisiones, prestaciones en especie, etc. | [optional] 
**SalarioDiarioIntegrado** | Pointer to **float32** | Salario que se integra con los pagos hechos en efectivo por cuota diaria, gratificaciones, percepciones, habitación, primas, comisiones, prestaciones en especie y cualquier otra cantidad o prestación que se entregue al trabajador por su trabajo. | [optional] 
**ClaveEntFed** | **string** | Clave de la entidad federativa en donde el trabajador prestó sus servicios al empleador, que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/state). | 
**SubContratacion** | Pointer to [**[]NominaReceptorNestedInputSubContratacionInner**](NominaReceptorNestedInputSubContratacionInner.md) |  | [optional] 

## Methods

### NewNominaReceptorInput

`func NewNominaReceptorInput(curp string, tipoContrato string, tipoRegimen string, numEmpleado string, periodicidadPago string, claveEntFed string, ) *NominaReceptorInput`

NewNominaReceptorInput instantiates a new NominaReceptorInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaReceptorInputWithDefaults

`func NewNominaReceptorInputWithDefaults() *NominaReceptorInput`

NewNominaReceptorInputWithDefaults instantiates a new NominaReceptorInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurp

`func (o *NominaReceptorInput) GetCurp() string`

GetCurp returns the Curp field if non-nil, zero value otherwise.

### GetCurpOk

`func (o *NominaReceptorInput) GetCurpOk() (*string, bool)`

GetCurpOk returns a tuple with the Curp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurp

`func (o *NominaReceptorInput) SetCurp(v string)`

SetCurp sets Curp field to given value.


### GetNumSeguridadSocial

`func (o *NominaReceptorInput) GetNumSeguridadSocial() string`

GetNumSeguridadSocial returns the NumSeguridadSocial field if non-nil, zero value otherwise.

### GetNumSeguridadSocialOk

`func (o *NominaReceptorInput) GetNumSeguridadSocialOk() (*string, bool)`

GetNumSeguridadSocialOk returns a tuple with the NumSeguridadSocial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSeguridadSocial

`func (o *NominaReceptorInput) SetNumSeguridadSocial(v string)`

SetNumSeguridadSocial sets NumSeguridadSocial field to given value.

### HasNumSeguridadSocial

`func (o *NominaReceptorInput) HasNumSeguridadSocial() bool`

HasNumSeguridadSocial returns a boolean if a field has been set.

### GetFechaInicioRelLaboral

`func (o *NominaReceptorInput) GetFechaInicioRelLaboral() string`

GetFechaInicioRelLaboral returns the FechaInicioRelLaboral field if non-nil, zero value otherwise.

### GetFechaInicioRelLaboralOk

`func (o *NominaReceptorInput) GetFechaInicioRelLaboralOk() (*string, bool)`

GetFechaInicioRelLaboralOk returns a tuple with the FechaInicioRelLaboral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaInicioRelLaboral

`func (o *NominaReceptorInput) SetFechaInicioRelLaboral(v string)`

SetFechaInicioRelLaboral sets FechaInicioRelLaboral field to given value.

### HasFechaInicioRelLaboral

`func (o *NominaReceptorInput) HasFechaInicioRelLaboral() bool`

HasFechaInicioRelLaboral returns a boolean if a field has been set.

### GetAntiguedad

`func (o *NominaReceptorInput) GetAntiguedad() NominaReceptorDirectPropertiesAntiguedad`

GetAntiguedad returns the Antiguedad field if non-nil, zero value otherwise.

### GetAntiguedadOk

`func (o *NominaReceptorInput) GetAntiguedadOk() (*NominaReceptorDirectPropertiesAntiguedad, bool)`

GetAntiguedadOk returns a tuple with the Antiguedad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiguedad

`func (o *NominaReceptorInput) SetAntiguedad(v NominaReceptorDirectPropertiesAntiguedad)`

SetAntiguedad sets Antiguedad field to given value.

### HasAntiguedad

`func (o *NominaReceptorInput) HasAntiguedad() bool`

HasAntiguedad returns a boolean if a field has been set.

### GetTipoContrato

`func (o *NominaReceptorInput) GetTipoContrato() string`

GetTipoContrato returns the TipoContrato field if non-nil, zero value otherwise.

### GetTipoContratoOk

`func (o *NominaReceptorInput) GetTipoContratoOk() (*string, bool)`

GetTipoContratoOk returns a tuple with the TipoContrato field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoContrato

`func (o *NominaReceptorInput) SetTipoContrato(v string)`

SetTipoContrato sets TipoContrato field to given value.


### GetSindicalizado

`func (o *NominaReceptorInput) GetSindicalizado() bool`

GetSindicalizado returns the Sindicalizado field if non-nil, zero value otherwise.

### GetSindicalizadoOk

`func (o *NominaReceptorInput) GetSindicalizadoOk() (*bool, bool)`

GetSindicalizadoOk returns a tuple with the Sindicalizado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSindicalizado

`func (o *NominaReceptorInput) SetSindicalizado(v bool)`

SetSindicalizado sets Sindicalizado field to given value.

### HasSindicalizado

`func (o *NominaReceptorInput) HasSindicalizado() bool`

HasSindicalizado returns a boolean if a field has been set.

### GetTipoJornada

`func (o *NominaReceptorInput) GetTipoJornada() string`

GetTipoJornada returns the TipoJornada field if non-nil, zero value otherwise.

### GetTipoJornadaOk

`func (o *NominaReceptorInput) GetTipoJornadaOk() (*string, bool)`

GetTipoJornadaOk returns a tuple with the TipoJornada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoJornada

`func (o *NominaReceptorInput) SetTipoJornada(v string)`

SetTipoJornada sets TipoJornada field to given value.

### HasTipoJornada

`func (o *NominaReceptorInput) HasTipoJornada() bool`

HasTipoJornada returns a boolean if a field has been set.

### GetTipoRegimen

`func (o *NominaReceptorInput) GetTipoRegimen() string`

GetTipoRegimen returns the TipoRegimen field if non-nil, zero value otherwise.

### GetTipoRegimenOk

`func (o *NominaReceptorInput) GetTipoRegimenOk() (*string, bool)`

GetTipoRegimenOk returns a tuple with the TipoRegimen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoRegimen

`func (o *NominaReceptorInput) SetTipoRegimen(v string)`

SetTipoRegimen sets TipoRegimen field to given value.


### GetNumEmpleado

`func (o *NominaReceptorInput) GetNumEmpleado() string`

GetNumEmpleado returns the NumEmpleado field if non-nil, zero value otherwise.

### GetNumEmpleadoOk

`func (o *NominaReceptorInput) GetNumEmpleadoOk() (*string, bool)`

GetNumEmpleadoOk returns a tuple with the NumEmpleado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEmpleado

`func (o *NominaReceptorInput) SetNumEmpleado(v string)`

SetNumEmpleado sets NumEmpleado field to given value.


### GetDepartamento

`func (o *NominaReceptorInput) GetDepartamento() string`

GetDepartamento returns the Departamento field if non-nil, zero value otherwise.

### GetDepartamentoOk

`func (o *NominaReceptorInput) GetDepartamentoOk() (*string, bool)`

GetDepartamentoOk returns a tuple with the Departamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartamento

`func (o *NominaReceptorInput) SetDepartamento(v string)`

SetDepartamento sets Departamento field to given value.

### HasDepartamento

`func (o *NominaReceptorInput) HasDepartamento() bool`

HasDepartamento returns a boolean if a field has been set.

### GetPuesto

`func (o *NominaReceptorInput) GetPuesto() string`

GetPuesto returns the Puesto field if non-nil, zero value otherwise.

### GetPuestoOk

`func (o *NominaReceptorInput) GetPuestoOk() (*string, bool)`

GetPuestoOk returns a tuple with the Puesto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuesto

`func (o *NominaReceptorInput) SetPuesto(v string)`

SetPuesto sets Puesto field to given value.

### HasPuesto

`func (o *NominaReceptorInput) HasPuesto() bool`

HasPuesto returns a boolean if a field has been set.

### GetRiesgoPuesto

`func (o *NominaReceptorInput) GetRiesgoPuesto() string`

GetRiesgoPuesto returns the RiesgoPuesto field if non-nil, zero value otherwise.

### GetRiesgoPuestoOk

`func (o *NominaReceptorInput) GetRiesgoPuestoOk() (*string, bool)`

GetRiesgoPuestoOk returns a tuple with the RiesgoPuesto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiesgoPuesto

`func (o *NominaReceptorInput) SetRiesgoPuesto(v string)`

SetRiesgoPuesto sets RiesgoPuesto field to given value.

### HasRiesgoPuesto

`func (o *NominaReceptorInput) HasRiesgoPuesto() bool`

HasRiesgoPuesto returns a boolean if a field has been set.

### GetPeriodicidadPago

`func (o *NominaReceptorInput) GetPeriodicidadPago() string`

GetPeriodicidadPago returns the PeriodicidadPago field if non-nil, zero value otherwise.

### GetPeriodicidadPagoOk

`func (o *NominaReceptorInput) GetPeriodicidadPagoOk() (*string, bool)`

GetPeriodicidadPagoOk returns a tuple with the PeriodicidadPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicidadPago

`func (o *NominaReceptorInput) SetPeriodicidadPago(v string)`

SetPeriodicidadPago sets PeriodicidadPago field to given value.


### GetBanco

`func (o *NominaReceptorInput) GetBanco() string`

GetBanco returns the Banco field if non-nil, zero value otherwise.

### GetBancoOk

`func (o *NominaReceptorInput) GetBancoOk() (*string, bool)`

GetBancoOk returns a tuple with the Banco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanco

`func (o *NominaReceptorInput) SetBanco(v string)`

SetBanco sets Banco field to given value.

### HasBanco

`func (o *NominaReceptorInput) HasBanco() bool`

HasBanco returns a boolean if a field has been set.

### GetCuentaBancaria

`func (o *NominaReceptorInput) GetCuentaBancaria() string`

GetCuentaBancaria returns the CuentaBancaria field if non-nil, zero value otherwise.

### GetCuentaBancariaOk

`func (o *NominaReceptorInput) GetCuentaBancariaOk() (*string, bool)`

GetCuentaBancariaOk returns a tuple with the CuentaBancaria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuentaBancaria

`func (o *NominaReceptorInput) SetCuentaBancaria(v string)`

SetCuentaBancaria sets CuentaBancaria field to given value.

### HasCuentaBancaria

`func (o *NominaReceptorInput) HasCuentaBancaria() bool`

HasCuentaBancaria returns a boolean if a field has been set.

### GetSalarioBaseCotApor

`func (o *NominaReceptorInput) GetSalarioBaseCotApor() float32`

GetSalarioBaseCotApor returns the SalarioBaseCotApor field if non-nil, zero value otherwise.

### GetSalarioBaseCotAporOk

`func (o *NominaReceptorInput) GetSalarioBaseCotAporOk() (*float32, bool)`

GetSalarioBaseCotAporOk returns a tuple with the SalarioBaseCotApor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalarioBaseCotApor

`func (o *NominaReceptorInput) SetSalarioBaseCotApor(v float32)`

SetSalarioBaseCotApor sets SalarioBaseCotApor field to given value.

### HasSalarioBaseCotApor

`func (o *NominaReceptorInput) HasSalarioBaseCotApor() bool`

HasSalarioBaseCotApor returns a boolean if a field has been set.

### GetSalarioDiarioIntegrado

`func (o *NominaReceptorInput) GetSalarioDiarioIntegrado() float32`

GetSalarioDiarioIntegrado returns the SalarioDiarioIntegrado field if non-nil, zero value otherwise.

### GetSalarioDiarioIntegradoOk

`func (o *NominaReceptorInput) GetSalarioDiarioIntegradoOk() (*float32, bool)`

GetSalarioDiarioIntegradoOk returns a tuple with the SalarioDiarioIntegrado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalarioDiarioIntegrado

`func (o *NominaReceptorInput) SetSalarioDiarioIntegrado(v float32)`

SetSalarioDiarioIntegrado sets SalarioDiarioIntegrado field to given value.

### HasSalarioDiarioIntegrado

`func (o *NominaReceptorInput) HasSalarioDiarioIntegrado() bool`

HasSalarioDiarioIntegrado returns a boolean if a field has been set.

### GetClaveEntFed

`func (o *NominaReceptorInput) GetClaveEntFed() string`

GetClaveEntFed returns the ClaveEntFed field if non-nil, zero value otherwise.

### GetClaveEntFedOk

`func (o *NominaReceptorInput) GetClaveEntFedOk() (*string, bool)`

GetClaveEntFedOk returns a tuple with the ClaveEntFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaveEntFed

`func (o *NominaReceptorInput) SetClaveEntFed(v string)`

SetClaveEntFed sets ClaveEntFed field to given value.


### GetSubContratacion

`func (o *NominaReceptorInput) GetSubContratacion() []NominaReceptorNestedInputSubContratacionInner`

GetSubContratacion returns the SubContratacion field if non-nil, zero value otherwise.

### GetSubContratacionOk

`func (o *NominaReceptorInput) GetSubContratacionOk() (*[]NominaReceptorNestedInputSubContratacionInner, bool)`

GetSubContratacionOk returns a tuple with the SubContratacion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubContratacion

`func (o *NominaReceptorInput) SetSubContratacion(v []NominaReceptorNestedInputSubContratacionInner)`

SetSubContratacion sets SubContratacion field to given value.

### HasSubContratacion

`func (o *NominaReceptorInput) HasSubContratacion() bool`

HasSubContratacion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


