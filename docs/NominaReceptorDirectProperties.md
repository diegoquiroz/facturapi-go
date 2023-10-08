# NominaReceptorDirectProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Curp** | Pointer to **string** | CURP del trabajador. | [optional] 
**NumSeguridadSocial** | Pointer to **string** | Número de seguridad social. | [optional] 
**FechaInicioRelLaboral** | Pointer to **string** | Fecha de inicio de la relación laboral entre el empleador y el empleado. | [optional] 
**Antiguedad** | Pointer to [**NominaReceptorDirectPropertiesAntiguedad**](NominaReceptorDirectPropertiesAntiguedad.md) |  | [optional] [default to true]
**TipoContrato** | Pointer to **string** | Clave del catálogo del SAT [Tipo de Contrato](#tipo-de-contrato). | [optional] 
**Sindicalizado** | Pointer to **bool** | Indica si el trabajador está asociado a un sindicato. | [optional] [default to false]
**TipoJornada** | Pointer to **string** | Clave del catálogo del SAT [Tipo de Jornada](#tipo-de-jornada). | [optional] 
**TipoRegimen** | Pointer to **string** | Clave del catálogo del SAT [Tipo de Régimen](#tipo-de-régimen). | [optional] 
**NumEmpleado** | Pointer to **string** | Número interno de empleado, asignado por el empleador. | [optional] 
**Departamento** | Pointer to **string** | Nombre del departamento o área a la que pertenece el trabajador. | [optional] 
**Puesto** | Pointer to **string** | Nombre del puesto asignado al empleado o el nombre de la actividad que realiza. | [optional] 
**RiesgoPuesto** | Pointer to **string** | Clave del catálogo del SAT [Riesgo del Puesto](#riesgo-del-puesto). | [optional] 
**PeriodicidadPago** | Pointer to **string** | Clave del catálogo del SAT [Periodicidad de Pago](#periodicidad-del-pago). | [optional] 
**Banco** | Pointer to **string** | Clave del banco de acuerdo al catálogo del SAT “Bancos” que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/bank). | [optional] 
**CuentaBancaria** | Pointer to **string** | Número de cuenta bancaria (11 caracteres) o número de teléfono celular (10 caracteres) o número de tarjeta (15 ó 16 caracteres) o la CLABE (18 caracteres) o número de monedero electrónico donde se realiza el depósito de nómina.  | [optional] 
**SalarioBaseCotApor** | Pointer to **float32** | Importe de la retribución en efectivo por cuota diaria, gratificaciones, percepciones, alimentación, habitación, primas, comisiones, prestaciones en especie, etc. | [optional] 
**SalarioDiarioIntegrado** | Pointer to **float32** | Salario que se integra con los pagos hechos en efectivo por cuota diaria, gratificaciones, percepciones, habitación, primas, comisiones, prestaciones en especie y cualquier otra cantidad o prestación que se entregue al trabajador por su trabajo. | [optional] 
**ClaveEntFed** | Pointer to **string** | Clave de la entidad federativa en donde el trabajador prestó sus servicios al empleador, que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/state). | [optional] 

## Methods

### NewNominaReceptorDirectProperties

`func NewNominaReceptorDirectProperties() *NominaReceptorDirectProperties`

NewNominaReceptorDirectProperties instantiates a new NominaReceptorDirectProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaReceptorDirectPropertiesWithDefaults

`func NewNominaReceptorDirectPropertiesWithDefaults() *NominaReceptorDirectProperties`

NewNominaReceptorDirectPropertiesWithDefaults instantiates a new NominaReceptorDirectProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurp

`func (o *NominaReceptorDirectProperties) GetCurp() string`

GetCurp returns the Curp field if non-nil, zero value otherwise.

### GetCurpOk

`func (o *NominaReceptorDirectProperties) GetCurpOk() (*string, bool)`

GetCurpOk returns a tuple with the Curp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurp

`func (o *NominaReceptorDirectProperties) SetCurp(v string)`

SetCurp sets Curp field to given value.

### HasCurp

`func (o *NominaReceptorDirectProperties) HasCurp() bool`

HasCurp returns a boolean if a field has been set.

### GetNumSeguridadSocial

`func (o *NominaReceptorDirectProperties) GetNumSeguridadSocial() string`

GetNumSeguridadSocial returns the NumSeguridadSocial field if non-nil, zero value otherwise.

### GetNumSeguridadSocialOk

`func (o *NominaReceptorDirectProperties) GetNumSeguridadSocialOk() (*string, bool)`

GetNumSeguridadSocialOk returns a tuple with the NumSeguridadSocial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSeguridadSocial

`func (o *NominaReceptorDirectProperties) SetNumSeguridadSocial(v string)`

SetNumSeguridadSocial sets NumSeguridadSocial field to given value.

### HasNumSeguridadSocial

`func (o *NominaReceptorDirectProperties) HasNumSeguridadSocial() bool`

HasNumSeguridadSocial returns a boolean if a field has been set.

### GetFechaInicioRelLaboral

`func (o *NominaReceptorDirectProperties) GetFechaInicioRelLaboral() string`

GetFechaInicioRelLaboral returns the FechaInicioRelLaboral field if non-nil, zero value otherwise.

### GetFechaInicioRelLaboralOk

`func (o *NominaReceptorDirectProperties) GetFechaInicioRelLaboralOk() (*string, bool)`

GetFechaInicioRelLaboralOk returns a tuple with the FechaInicioRelLaboral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaInicioRelLaboral

`func (o *NominaReceptorDirectProperties) SetFechaInicioRelLaboral(v string)`

SetFechaInicioRelLaboral sets FechaInicioRelLaboral field to given value.

### HasFechaInicioRelLaboral

`func (o *NominaReceptorDirectProperties) HasFechaInicioRelLaboral() bool`

HasFechaInicioRelLaboral returns a boolean if a field has been set.

### GetAntiguedad

`func (o *NominaReceptorDirectProperties) GetAntiguedad() NominaReceptorDirectPropertiesAntiguedad`

GetAntiguedad returns the Antiguedad field if non-nil, zero value otherwise.

### GetAntiguedadOk

`func (o *NominaReceptorDirectProperties) GetAntiguedadOk() (*NominaReceptorDirectPropertiesAntiguedad, bool)`

GetAntiguedadOk returns a tuple with the Antiguedad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiguedad

`func (o *NominaReceptorDirectProperties) SetAntiguedad(v NominaReceptorDirectPropertiesAntiguedad)`

SetAntiguedad sets Antiguedad field to given value.

### HasAntiguedad

`func (o *NominaReceptorDirectProperties) HasAntiguedad() bool`

HasAntiguedad returns a boolean if a field has been set.

### GetTipoContrato

`func (o *NominaReceptorDirectProperties) GetTipoContrato() string`

GetTipoContrato returns the TipoContrato field if non-nil, zero value otherwise.

### GetTipoContratoOk

`func (o *NominaReceptorDirectProperties) GetTipoContratoOk() (*string, bool)`

GetTipoContratoOk returns a tuple with the TipoContrato field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoContrato

`func (o *NominaReceptorDirectProperties) SetTipoContrato(v string)`

SetTipoContrato sets TipoContrato field to given value.

### HasTipoContrato

`func (o *NominaReceptorDirectProperties) HasTipoContrato() bool`

HasTipoContrato returns a boolean if a field has been set.

### GetSindicalizado

`func (o *NominaReceptorDirectProperties) GetSindicalizado() bool`

GetSindicalizado returns the Sindicalizado field if non-nil, zero value otherwise.

### GetSindicalizadoOk

`func (o *NominaReceptorDirectProperties) GetSindicalizadoOk() (*bool, bool)`

GetSindicalizadoOk returns a tuple with the Sindicalizado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSindicalizado

`func (o *NominaReceptorDirectProperties) SetSindicalizado(v bool)`

SetSindicalizado sets Sindicalizado field to given value.

### HasSindicalizado

`func (o *NominaReceptorDirectProperties) HasSindicalizado() bool`

HasSindicalizado returns a boolean if a field has been set.

### GetTipoJornada

`func (o *NominaReceptorDirectProperties) GetTipoJornada() string`

GetTipoJornada returns the TipoJornada field if non-nil, zero value otherwise.

### GetTipoJornadaOk

`func (o *NominaReceptorDirectProperties) GetTipoJornadaOk() (*string, bool)`

GetTipoJornadaOk returns a tuple with the TipoJornada field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoJornada

`func (o *NominaReceptorDirectProperties) SetTipoJornada(v string)`

SetTipoJornada sets TipoJornada field to given value.

### HasTipoJornada

`func (o *NominaReceptorDirectProperties) HasTipoJornada() bool`

HasTipoJornada returns a boolean if a field has been set.

### GetTipoRegimen

`func (o *NominaReceptorDirectProperties) GetTipoRegimen() string`

GetTipoRegimen returns the TipoRegimen field if non-nil, zero value otherwise.

### GetTipoRegimenOk

`func (o *NominaReceptorDirectProperties) GetTipoRegimenOk() (*string, bool)`

GetTipoRegimenOk returns a tuple with the TipoRegimen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoRegimen

`func (o *NominaReceptorDirectProperties) SetTipoRegimen(v string)`

SetTipoRegimen sets TipoRegimen field to given value.

### HasTipoRegimen

`func (o *NominaReceptorDirectProperties) HasTipoRegimen() bool`

HasTipoRegimen returns a boolean if a field has been set.

### GetNumEmpleado

`func (o *NominaReceptorDirectProperties) GetNumEmpleado() string`

GetNumEmpleado returns the NumEmpleado field if non-nil, zero value otherwise.

### GetNumEmpleadoOk

`func (o *NominaReceptorDirectProperties) GetNumEmpleadoOk() (*string, bool)`

GetNumEmpleadoOk returns a tuple with the NumEmpleado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumEmpleado

`func (o *NominaReceptorDirectProperties) SetNumEmpleado(v string)`

SetNumEmpleado sets NumEmpleado field to given value.

### HasNumEmpleado

`func (o *NominaReceptorDirectProperties) HasNumEmpleado() bool`

HasNumEmpleado returns a boolean if a field has been set.

### GetDepartamento

`func (o *NominaReceptorDirectProperties) GetDepartamento() string`

GetDepartamento returns the Departamento field if non-nil, zero value otherwise.

### GetDepartamentoOk

`func (o *NominaReceptorDirectProperties) GetDepartamentoOk() (*string, bool)`

GetDepartamentoOk returns a tuple with the Departamento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartamento

`func (o *NominaReceptorDirectProperties) SetDepartamento(v string)`

SetDepartamento sets Departamento field to given value.

### HasDepartamento

`func (o *NominaReceptorDirectProperties) HasDepartamento() bool`

HasDepartamento returns a boolean if a field has been set.

### GetPuesto

`func (o *NominaReceptorDirectProperties) GetPuesto() string`

GetPuesto returns the Puesto field if non-nil, zero value otherwise.

### GetPuestoOk

`func (o *NominaReceptorDirectProperties) GetPuestoOk() (*string, bool)`

GetPuestoOk returns a tuple with the Puesto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuesto

`func (o *NominaReceptorDirectProperties) SetPuesto(v string)`

SetPuesto sets Puesto field to given value.

### HasPuesto

`func (o *NominaReceptorDirectProperties) HasPuesto() bool`

HasPuesto returns a boolean if a field has been set.

### GetRiesgoPuesto

`func (o *NominaReceptorDirectProperties) GetRiesgoPuesto() string`

GetRiesgoPuesto returns the RiesgoPuesto field if non-nil, zero value otherwise.

### GetRiesgoPuestoOk

`func (o *NominaReceptorDirectProperties) GetRiesgoPuestoOk() (*string, bool)`

GetRiesgoPuestoOk returns a tuple with the RiesgoPuesto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiesgoPuesto

`func (o *NominaReceptorDirectProperties) SetRiesgoPuesto(v string)`

SetRiesgoPuesto sets RiesgoPuesto field to given value.

### HasRiesgoPuesto

`func (o *NominaReceptorDirectProperties) HasRiesgoPuesto() bool`

HasRiesgoPuesto returns a boolean if a field has been set.

### GetPeriodicidadPago

`func (o *NominaReceptorDirectProperties) GetPeriodicidadPago() string`

GetPeriodicidadPago returns the PeriodicidadPago field if non-nil, zero value otherwise.

### GetPeriodicidadPagoOk

`func (o *NominaReceptorDirectProperties) GetPeriodicidadPagoOk() (*string, bool)`

GetPeriodicidadPagoOk returns a tuple with the PeriodicidadPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicidadPago

`func (o *NominaReceptorDirectProperties) SetPeriodicidadPago(v string)`

SetPeriodicidadPago sets PeriodicidadPago field to given value.

### HasPeriodicidadPago

`func (o *NominaReceptorDirectProperties) HasPeriodicidadPago() bool`

HasPeriodicidadPago returns a boolean if a field has been set.

### GetBanco

`func (o *NominaReceptorDirectProperties) GetBanco() string`

GetBanco returns the Banco field if non-nil, zero value otherwise.

### GetBancoOk

`func (o *NominaReceptorDirectProperties) GetBancoOk() (*string, bool)`

GetBancoOk returns a tuple with the Banco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanco

`func (o *NominaReceptorDirectProperties) SetBanco(v string)`

SetBanco sets Banco field to given value.

### HasBanco

`func (o *NominaReceptorDirectProperties) HasBanco() bool`

HasBanco returns a boolean if a field has been set.

### GetCuentaBancaria

`func (o *NominaReceptorDirectProperties) GetCuentaBancaria() string`

GetCuentaBancaria returns the CuentaBancaria field if non-nil, zero value otherwise.

### GetCuentaBancariaOk

`func (o *NominaReceptorDirectProperties) GetCuentaBancariaOk() (*string, bool)`

GetCuentaBancariaOk returns a tuple with the CuentaBancaria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuentaBancaria

`func (o *NominaReceptorDirectProperties) SetCuentaBancaria(v string)`

SetCuentaBancaria sets CuentaBancaria field to given value.

### HasCuentaBancaria

`func (o *NominaReceptorDirectProperties) HasCuentaBancaria() bool`

HasCuentaBancaria returns a boolean if a field has been set.

### GetSalarioBaseCotApor

`func (o *NominaReceptorDirectProperties) GetSalarioBaseCotApor() float32`

GetSalarioBaseCotApor returns the SalarioBaseCotApor field if non-nil, zero value otherwise.

### GetSalarioBaseCotAporOk

`func (o *NominaReceptorDirectProperties) GetSalarioBaseCotAporOk() (*float32, bool)`

GetSalarioBaseCotAporOk returns a tuple with the SalarioBaseCotApor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalarioBaseCotApor

`func (o *NominaReceptorDirectProperties) SetSalarioBaseCotApor(v float32)`

SetSalarioBaseCotApor sets SalarioBaseCotApor field to given value.

### HasSalarioBaseCotApor

`func (o *NominaReceptorDirectProperties) HasSalarioBaseCotApor() bool`

HasSalarioBaseCotApor returns a boolean if a field has been set.

### GetSalarioDiarioIntegrado

`func (o *NominaReceptorDirectProperties) GetSalarioDiarioIntegrado() float32`

GetSalarioDiarioIntegrado returns the SalarioDiarioIntegrado field if non-nil, zero value otherwise.

### GetSalarioDiarioIntegradoOk

`func (o *NominaReceptorDirectProperties) GetSalarioDiarioIntegradoOk() (*float32, bool)`

GetSalarioDiarioIntegradoOk returns a tuple with the SalarioDiarioIntegrado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalarioDiarioIntegrado

`func (o *NominaReceptorDirectProperties) SetSalarioDiarioIntegrado(v float32)`

SetSalarioDiarioIntegrado sets SalarioDiarioIntegrado field to given value.

### HasSalarioDiarioIntegrado

`func (o *NominaReceptorDirectProperties) HasSalarioDiarioIntegrado() bool`

HasSalarioDiarioIntegrado returns a boolean if a field has been set.

### GetClaveEntFed

`func (o *NominaReceptorDirectProperties) GetClaveEntFed() string`

GetClaveEntFed returns the ClaveEntFed field if non-nil, zero value otherwise.

### GetClaveEntFedOk

`func (o *NominaReceptorDirectProperties) GetClaveEntFedOk() (*string, bool)`

GetClaveEntFedOk returns a tuple with the ClaveEntFed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaveEntFed

`func (o *NominaReceptorDirectProperties) SetClaveEntFed(v string)`

SetClaveEntFed sets ClaveEntFed field to given value.

### HasClaveEntFed

`func (o *NominaReceptorDirectProperties) HasClaveEntFed() bool`

HasClaveEntFed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


