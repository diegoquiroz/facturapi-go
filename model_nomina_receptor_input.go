/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NominaReceptorInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaReceptorInput{}

// NominaReceptorInput Información del trabajador.
type NominaReceptorInput struct {
	// CURP del trabajador.
	Curp string `json:"curp"`
	// Número de seguridad social.
	NumSeguridadSocial *string `json:"num_seguridad_social,omitempty"`
	// Fecha de inicio de la relación laboral entre el empleador y el empleado.
	FechaInicioRelLaboral *string `json:"fecha_inicio_rel_laboral,omitempty"`
	Antiguedad *NominaReceptorDirectPropertiesAntiguedad `json:"antiguedad,omitempty"`
	// Clave del catálogo del SAT [Tipo de Contrato](#tipo-de-contrato).
	TipoContrato string `json:"tipo_contrato"`
	// Indica si el trabajador está asociado a un sindicato.
	Sindicalizado *bool `json:"sindicalizado,omitempty"`
	// Clave del catálogo del SAT [Tipo de Jornada](#tipo-de-jornada).
	TipoJornada *string `json:"tipo_jornada,omitempty"`
	// Clave del catálogo del SAT [Tipo de Régimen](#tipo-de-régimen).
	TipoRegimen string `json:"tipo_regimen"`
	// Número interno de empleado, asignado por el empleador.
	NumEmpleado string `json:"num_empleado"`
	// Nombre del departamento o área a la que pertenece el trabajador.
	Departamento *string `json:"departamento,omitempty"`
	// Nombre del puesto asignado al empleado o el nombre de la actividad que realiza.
	Puesto *string `json:"puesto,omitempty"`
	// Clave del catálogo del SAT [Riesgo del Puesto](#riesgo-del-puesto).
	RiesgoPuesto *string `json:"riesgo_puesto,omitempty"`
	// Clave del catálogo del SAT [Periodicidad de Pago](#periodicidad-del-pago).
	PeriodicidadPago string `json:"periodicidad_pago"`
	// Clave del banco de acuerdo al catálogo del SAT “Bancos” que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/bank).
	Banco *string `json:"banco,omitempty"`
	// Número de cuenta bancaria (11 caracteres) o número de teléfono celular (10 caracteres) o número de tarjeta (15 ó 16 caracteres) o la CLABE (18 caracteres) o número de monedero electrónico donde se realiza el depósito de nómina. 
	CuentaBancaria *string `json:"cuenta_bancaria,omitempty"`
	// Importe de la retribución en efectivo por cuota diaria, gratificaciones, percepciones, alimentación, habitación, primas, comisiones, prestaciones en especie, etc.
	SalarioBaseCotApor *float32 `json:"salario_base_cot_apor,omitempty"`
	// Salario que se integra con los pagos hechos en efectivo por cuota diaria, gratificaciones, percepciones, habitación, primas, comisiones, prestaciones en especie y cualquier otra cantidad o prestación que se entregue al trabajador por su trabajo.
	SalarioDiarioIntegrado *float32 `json:"salario_diario_integrado,omitempty"`
	// Clave de la entidad federativa en donde el trabajador prestó sus servicios al empleador, que puedes consultar utilizando nuestra [herramienta de búsqueda](https://dashboard.facturapi.io/catalogs/state).
	ClaveEntFed string `json:"clave_ent_fed"`
	SubContratacion []NominaReceptorNestedInputSubContratacionInner `json:"sub_contratacion,omitempty"`
}

// NewNominaReceptorInput instantiates a new NominaReceptorInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaReceptorInput(curp string, tipoContrato string, tipoRegimen string, numEmpleado string, periodicidadPago string, claveEntFed string) *NominaReceptorInput {
	this := NominaReceptorInput{}
	this.Curp = curp
	var antiguedad NominaReceptorDirectPropertiesAntiguedad = true
	this.Antiguedad = &antiguedad
	this.TipoContrato = tipoContrato
	var sindicalizado bool = false
	this.Sindicalizado = &sindicalizado
	this.TipoRegimen = tipoRegimen
	this.NumEmpleado = numEmpleado
	this.PeriodicidadPago = periodicidadPago
	this.ClaveEntFed = claveEntFed
	return &this
}

// NewNominaReceptorInputWithDefaults instantiates a new NominaReceptorInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaReceptorInputWithDefaults() *NominaReceptorInput {
	this := NominaReceptorInput{}
	var antiguedad NominaReceptorDirectPropertiesAntiguedad = true
	this.Antiguedad = &antiguedad
	var sindicalizado bool = false
	this.Sindicalizado = &sindicalizado
	return &this
}

// GetCurp returns the Curp field value
func (o *NominaReceptorInput) GetCurp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Curp
}

// GetCurpOk returns a tuple with the Curp field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetCurpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Curp, true
}

// SetCurp sets field value
func (o *NominaReceptorInput) SetCurp(v string) {
	o.Curp = v
}

// GetNumSeguridadSocial returns the NumSeguridadSocial field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetNumSeguridadSocial() string {
	if o == nil || IsNil(o.NumSeguridadSocial) {
		var ret string
		return ret
	}
	return *o.NumSeguridadSocial
}

// GetNumSeguridadSocialOk returns a tuple with the NumSeguridadSocial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetNumSeguridadSocialOk() (*string, bool) {
	if o == nil || IsNil(o.NumSeguridadSocial) {
		return nil, false
	}
	return o.NumSeguridadSocial, true
}

// HasNumSeguridadSocial returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasNumSeguridadSocial() bool {
	if o != nil && !IsNil(o.NumSeguridadSocial) {
		return true
	}

	return false
}

// SetNumSeguridadSocial gets a reference to the given string and assigns it to the NumSeguridadSocial field.
func (o *NominaReceptorInput) SetNumSeguridadSocial(v string) {
	o.NumSeguridadSocial = &v
}

// GetFechaInicioRelLaboral returns the FechaInicioRelLaboral field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetFechaInicioRelLaboral() string {
	if o == nil || IsNil(o.FechaInicioRelLaboral) {
		var ret string
		return ret
	}
	return *o.FechaInicioRelLaboral
}

// GetFechaInicioRelLaboralOk returns a tuple with the FechaInicioRelLaboral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetFechaInicioRelLaboralOk() (*string, bool) {
	if o == nil || IsNil(o.FechaInicioRelLaboral) {
		return nil, false
	}
	return o.FechaInicioRelLaboral, true
}

// HasFechaInicioRelLaboral returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasFechaInicioRelLaboral() bool {
	if o != nil && !IsNil(o.FechaInicioRelLaboral) {
		return true
	}

	return false
}

// SetFechaInicioRelLaboral gets a reference to the given string and assigns it to the FechaInicioRelLaboral field.
func (o *NominaReceptorInput) SetFechaInicioRelLaboral(v string) {
	o.FechaInicioRelLaboral = &v
}

// GetAntiguedad returns the Antiguedad field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetAntiguedad() NominaReceptorDirectPropertiesAntiguedad {
	if o == nil || IsNil(o.Antiguedad) {
		var ret NominaReceptorDirectPropertiesAntiguedad
		return ret
	}
	return *o.Antiguedad
}

// GetAntiguedadOk returns a tuple with the Antiguedad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetAntiguedadOk() (*NominaReceptorDirectPropertiesAntiguedad, bool) {
	if o == nil || IsNil(o.Antiguedad) {
		return nil, false
	}
	return o.Antiguedad, true
}

// HasAntiguedad returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasAntiguedad() bool {
	if o != nil && !IsNil(o.Antiguedad) {
		return true
	}

	return false
}

// SetAntiguedad gets a reference to the given NominaReceptorDirectPropertiesAntiguedad and assigns it to the Antiguedad field.
func (o *NominaReceptorInput) SetAntiguedad(v NominaReceptorDirectPropertiesAntiguedad) {
	o.Antiguedad = &v
}

// GetTipoContrato returns the TipoContrato field value
func (o *NominaReceptorInput) GetTipoContrato() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TipoContrato
}

// GetTipoContratoOk returns a tuple with the TipoContrato field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetTipoContratoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TipoContrato, true
}

// SetTipoContrato sets field value
func (o *NominaReceptorInput) SetTipoContrato(v string) {
	o.TipoContrato = v
}

// GetSindicalizado returns the Sindicalizado field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetSindicalizado() bool {
	if o == nil || IsNil(o.Sindicalizado) {
		var ret bool
		return ret
	}
	return *o.Sindicalizado
}

// GetSindicalizadoOk returns a tuple with the Sindicalizado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetSindicalizadoOk() (*bool, bool) {
	if o == nil || IsNil(o.Sindicalizado) {
		return nil, false
	}
	return o.Sindicalizado, true
}

// HasSindicalizado returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasSindicalizado() bool {
	if o != nil && !IsNil(o.Sindicalizado) {
		return true
	}

	return false
}

// SetSindicalizado gets a reference to the given bool and assigns it to the Sindicalizado field.
func (o *NominaReceptorInput) SetSindicalizado(v bool) {
	o.Sindicalizado = &v
}

// GetTipoJornada returns the TipoJornada field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetTipoJornada() string {
	if o == nil || IsNil(o.TipoJornada) {
		var ret string
		return ret
	}
	return *o.TipoJornada
}

// GetTipoJornadaOk returns a tuple with the TipoJornada field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetTipoJornadaOk() (*string, bool) {
	if o == nil || IsNil(o.TipoJornada) {
		return nil, false
	}
	return o.TipoJornada, true
}

// HasTipoJornada returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasTipoJornada() bool {
	if o != nil && !IsNil(o.TipoJornada) {
		return true
	}

	return false
}

// SetTipoJornada gets a reference to the given string and assigns it to the TipoJornada field.
func (o *NominaReceptorInput) SetTipoJornada(v string) {
	o.TipoJornada = &v
}

// GetTipoRegimen returns the TipoRegimen field value
func (o *NominaReceptorInput) GetTipoRegimen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TipoRegimen
}

// GetTipoRegimenOk returns a tuple with the TipoRegimen field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetTipoRegimenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TipoRegimen, true
}

// SetTipoRegimen sets field value
func (o *NominaReceptorInput) SetTipoRegimen(v string) {
	o.TipoRegimen = v
}

// GetNumEmpleado returns the NumEmpleado field value
func (o *NominaReceptorInput) GetNumEmpleado() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NumEmpleado
}

// GetNumEmpleadoOk returns a tuple with the NumEmpleado field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetNumEmpleadoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumEmpleado, true
}

// SetNumEmpleado sets field value
func (o *NominaReceptorInput) SetNumEmpleado(v string) {
	o.NumEmpleado = v
}

// GetDepartamento returns the Departamento field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetDepartamento() string {
	if o == nil || IsNil(o.Departamento) {
		var ret string
		return ret
	}
	return *o.Departamento
}

// GetDepartamentoOk returns a tuple with the Departamento field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetDepartamentoOk() (*string, bool) {
	if o == nil || IsNil(o.Departamento) {
		return nil, false
	}
	return o.Departamento, true
}

// HasDepartamento returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasDepartamento() bool {
	if o != nil && !IsNil(o.Departamento) {
		return true
	}

	return false
}

// SetDepartamento gets a reference to the given string and assigns it to the Departamento field.
func (o *NominaReceptorInput) SetDepartamento(v string) {
	o.Departamento = &v
}

// GetPuesto returns the Puesto field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetPuesto() string {
	if o == nil || IsNil(o.Puesto) {
		var ret string
		return ret
	}
	return *o.Puesto
}

// GetPuestoOk returns a tuple with the Puesto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetPuestoOk() (*string, bool) {
	if o == nil || IsNil(o.Puesto) {
		return nil, false
	}
	return o.Puesto, true
}

// HasPuesto returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasPuesto() bool {
	if o != nil && !IsNil(o.Puesto) {
		return true
	}

	return false
}

// SetPuesto gets a reference to the given string and assigns it to the Puesto field.
func (o *NominaReceptorInput) SetPuesto(v string) {
	o.Puesto = &v
}

// GetRiesgoPuesto returns the RiesgoPuesto field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetRiesgoPuesto() string {
	if o == nil || IsNil(o.RiesgoPuesto) {
		var ret string
		return ret
	}
	return *o.RiesgoPuesto
}

// GetRiesgoPuestoOk returns a tuple with the RiesgoPuesto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetRiesgoPuestoOk() (*string, bool) {
	if o == nil || IsNil(o.RiesgoPuesto) {
		return nil, false
	}
	return o.RiesgoPuesto, true
}

// HasRiesgoPuesto returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasRiesgoPuesto() bool {
	if o != nil && !IsNil(o.RiesgoPuesto) {
		return true
	}

	return false
}

// SetRiesgoPuesto gets a reference to the given string and assigns it to the RiesgoPuesto field.
func (o *NominaReceptorInput) SetRiesgoPuesto(v string) {
	o.RiesgoPuesto = &v
}

// GetPeriodicidadPago returns the PeriodicidadPago field value
func (o *NominaReceptorInput) GetPeriodicidadPago() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodicidadPago
}

// GetPeriodicidadPagoOk returns a tuple with the PeriodicidadPago field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetPeriodicidadPagoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodicidadPago, true
}

// SetPeriodicidadPago sets field value
func (o *NominaReceptorInput) SetPeriodicidadPago(v string) {
	o.PeriodicidadPago = v
}

// GetBanco returns the Banco field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetBanco() string {
	if o == nil || IsNil(o.Banco) {
		var ret string
		return ret
	}
	return *o.Banco
}

// GetBancoOk returns a tuple with the Banco field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetBancoOk() (*string, bool) {
	if o == nil || IsNil(o.Banco) {
		return nil, false
	}
	return o.Banco, true
}

// HasBanco returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasBanco() bool {
	if o != nil && !IsNil(o.Banco) {
		return true
	}

	return false
}

// SetBanco gets a reference to the given string and assigns it to the Banco field.
func (o *NominaReceptorInput) SetBanco(v string) {
	o.Banco = &v
}

// GetCuentaBancaria returns the CuentaBancaria field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetCuentaBancaria() string {
	if o == nil || IsNil(o.CuentaBancaria) {
		var ret string
		return ret
	}
	return *o.CuentaBancaria
}

// GetCuentaBancariaOk returns a tuple with the CuentaBancaria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetCuentaBancariaOk() (*string, bool) {
	if o == nil || IsNil(o.CuentaBancaria) {
		return nil, false
	}
	return o.CuentaBancaria, true
}

// HasCuentaBancaria returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasCuentaBancaria() bool {
	if o != nil && !IsNil(o.CuentaBancaria) {
		return true
	}

	return false
}

// SetCuentaBancaria gets a reference to the given string and assigns it to the CuentaBancaria field.
func (o *NominaReceptorInput) SetCuentaBancaria(v string) {
	o.CuentaBancaria = &v
}

// GetSalarioBaseCotApor returns the SalarioBaseCotApor field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetSalarioBaseCotApor() float32 {
	if o == nil || IsNil(o.SalarioBaseCotApor) {
		var ret float32
		return ret
	}
	return *o.SalarioBaseCotApor
}

// GetSalarioBaseCotAporOk returns a tuple with the SalarioBaseCotApor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetSalarioBaseCotAporOk() (*float32, bool) {
	if o == nil || IsNil(o.SalarioBaseCotApor) {
		return nil, false
	}
	return o.SalarioBaseCotApor, true
}

// HasSalarioBaseCotApor returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasSalarioBaseCotApor() bool {
	if o != nil && !IsNil(o.SalarioBaseCotApor) {
		return true
	}

	return false
}

// SetSalarioBaseCotApor gets a reference to the given float32 and assigns it to the SalarioBaseCotApor field.
func (o *NominaReceptorInput) SetSalarioBaseCotApor(v float32) {
	o.SalarioBaseCotApor = &v
}

// GetSalarioDiarioIntegrado returns the SalarioDiarioIntegrado field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetSalarioDiarioIntegrado() float32 {
	if o == nil || IsNil(o.SalarioDiarioIntegrado) {
		var ret float32
		return ret
	}
	return *o.SalarioDiarioIntegrado
}

// GetSalarioDiarioIntegradoOk returns a tuple with the SalarioDiarioIntegrado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetSalarioDiarioIntegradoOk() (*float32, bool) {
	if o == nil || IsNil(o.SalarioDiarioIntegrado) {
		return nil, false
	}
	return o.SalarioDiarioIntegrado, true
}

// HasSalarioDiarioIntegrado returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasSalarioDiarioIntegrado() bool {
	if o != nil && !IsNil(o.SalarioDiarioIntegrado) {
		return true
	}

	return false
}

// SetSalarioDiarioIntegrado gets a reference to the given float32 and assigns it to the SalarioDiarioIntegrado field.
func (o *NominaReceptorInput) SetSalarioDiarioIntegrado(v float32) {
	o.SalarioDiarioIntegrado = &v
}

// GetClaveEntFed returns the ClaveEntFed field value
func (o *NominaReceptorInput) GetClaveEntFed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaveEntFed
}

// GetClaveEntFedOk returns a tuple with the ClaveEntFed field value
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetClaveEntFedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaveEntFed, true
}

// SetClaveEntFed sets field value
func (o *NominaReceptorInput) SetClaveEntFed(v string) {
	o.ClaveEntFed = v
}

// GetSubContratacion returns the SubContratacion field value if set, zero value otherwise.
func (o *NominaReceptorInput) GetSubContratacion() []NominaReceptorNestedInputSubContratacionInner {
	if o == nil || IsNil(o.SubContratacion) {
		var ret []NominaReceptorNestedInputSubContratacionInner
		return ret
	}
	return o.SubContratacion
}

// GetSubContratacionOk returns a tuple with the SubContratacion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorInput) GetSubContratacionOk() ([]NominaReceptorNestedInputSubContratacionInner, bool) {
	if o == nil || IsNil(o.SubContratacion) {
		return nil, false
	}
	return o.SubContratacion, true
}

// HasSubContratacion returns a boolean if a field has been set.
func (o *NominaReceptorInput) HasSubContratacion() bool {
	if o != nil && !IsNil(o.SubContratacion) {
		return true
	}

	return false
}

// SetSubContratacion gets a reference to the given []NominaReceptorNestedInputSubContratacionInner and assigns it to the SubContratacion field.
func (o *NominaReceptorInput) SetSubContratacion(v []NominaReceptorNestedInputSubContratacionInner) {
	o.SubContratacion = v
}

func (o NominaReceptorInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaReceptorInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["curp"] = o.Curp
	if !IsNil(o.NumSeguridadSocial) {
		toSerialize["num_seguridad_social"] = o.NumSeguridadSocial
	}
	if !IsNil(o.FechaInicioRelLaboral) {
		toSerialize["fecha_inicio_rel_laboral"] = o.FechaInicioRelLaboral
	}
	if !IsNil(o.Antiguedad) {
		toSerialize["antiguedad"] = o.Antiguedad
	}
	toSerialize["tipo_contrato"] = o.TipoContrato
	if !IsNil(o.Sindicalizado) {
		toSerialize["sindicalizado"] = o.Sindicalizado
	}
	if !IsNil(o.TipoJornada) {
		toSerialize["tipo_jornada"] = o.TipoJornada
	}
	toSerialize["tipo_regimen"] = o.TipoRegimen
	toSerialize["num_empleado"] = o.NumEmpleado
	if !IsNil(o.Departamento) {
		toSerialize["departamento"] = o.Departamento
	}
	if !IsNil(o.Puesto) {
		toSerialize["puesto"] = o.Puesto
	}
	if !IsNil(o.RiesgoPuesto) {
		toSerialize["riesgo_puesto"] = o.RiesgoPuesto
	}
	toSerialize["periodicidad_pago"] = o.PeriodicidadPago
	if !IsNil(o.Banco) {
		toSerialize["banco"] = o.Banco
	}
	if !IsNil(o.CuentaBancaria) {
		toSerialize["cuenta_bancaria"] = o.CuentaBancaria
	}
	if !IsNil(o.SalarioBaseCotApor) {
		toSerialize["salario_base_cot_apor"] = o.SalarioBaseCotApor
	}
	if !IsNil(o.SalarioDiarioIntegrado) {
		toSerialize["salario_diario_integrado"] = o.SalarioDiarioIntegrado
	}
	toSerialize["clave_ent_fed"] = o.ClaveEntFed
	if !IsNil(o.SubContratacion) {
		toSerialize["sub_contratacion"] = o.SubContratacion
	}
	return toSerialize, nil
}

type NullableNominaReceptorInput struct {
	value *NominaReceptorInput
	isSet bool
}

func (v NullableNominaReceptorInput) Get() *NominaReceptorInput {
	return v.value
}

func (v *NullableNominaReceptorInput) Set(val *NominaReceptorInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaReceptorInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaReceptorInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaReceptorInput(val *NominaReceptorInput) *NullableNominaReceptorInput {
	return &NullableNominaReceptorInput{value: val, isSet: true}
}

func (v NullableNominaReceptorInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaReceptorInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


