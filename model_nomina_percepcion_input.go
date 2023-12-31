/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi

import (
	"encoding/json"
)

// checks if the NominaPercepcionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaPercepcionInput{}

// NominaPercepcionInput struct for NominaPercepcionInput
type NominaPercepcionInput struct {
	// Clave del catálogo [Tipo de percepción](#tipo-de-percepcion).
	TipoPercepcion string `json:"tipo_percepcion"`
	// Concepto de la percepción. Si no se envía, se utilizará la descripción del catálogo del tipo de percepción.
	Concepto *string `json:"concepto,omitempty"`
	// Clave de control interno que asigna el patrón a cada percepción de nómina propia de su contabilidad.
	Clave string `json:"clave"`
	// Importe gravado por el concepto indicado en el tipo de percepción.
	ImporteGravado float32 `json:"importe_gravado"`
	// Importe exento por el concepto indicado en el tipo de percepción.
	ImporteExento float32 `json:"importe_exento"`
	AccionesOTitulos *NominaAccionesInput `json:"acciones_o_titulos,omitempty"`
	HorasExtra []NominaHorasExtraInput `json:"horas_extra,omitempty"`
}

// NewNominaPercepcionInput instantiates a new NominaPercepcionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaPercepcionInput(tipoPercepcion string, clave string, importeGravado float32, importeExento float32) *NominaPercepcionInput {
	this := NominaPercepcionInput{}
	this.TipoPercepcion = tipoPercepcion
	this.Clave = clave
	this.ImporteGravado = importeGravado
	this.ImporteExento = importeExento
	return &this
}

// NewNominaPercepcionInputWithDefaults instantiates a new NominaPercepcionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaPercepcionInputWithDefaults() *NominaPercepcionInput {
	this := NominaPercepcionInput{}
	return &this
}

// GetTipoPercepcion returns the TipoPercepcion field value
func (o *NominaPercepcionInput) GetTipoPercepcion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TipoPercepcion
}

// GetTipoPercepcionOk returns a tuple with the TipoPercepcion field value
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetTipoPercepcionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TipoPercepcion, true
}

// SetTipoPercepcion sets field value
func (o *NominaPercepcionInput) SetTipoPercepcion(v string) {
	o.TipoPercepcion = v
}

// GetConcepto returns the Concepto field value if set, zero value otherwise.
func (o *NominaPercepcionInput) GetConcepto() string {
	if o == nil || IsNil(o.Concepto) {
		var ret string
		return ret
	}
	return *o.Concepto
}

// GetConceptoOk returns a tuple with the Concepto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetConceptoOk() (*string, bool) {
	if o == nil || IsNil(o.Concepto) {
		return nil, false
	}
	return o.Concepto, true
}

// HasConcepto returns a boolean if a field has been set.
func (o *NominaPercepcionInput) HasConcepto() bool {
	if o != nil && !IsNil(o.Concepto) {
		return true
	}

	return false
}

// SetConcepto gets a reference to the given string and assigns it to the Concepto field.
func (o *NominaPercepcionInput) SetConcepto(v string) {
	o.Concepto = &v
}

// GetClave returns the Clave field value
func (o *NominaPercepcionInput) GetClave() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Clave
}

// GetClaveOk returns a tuple with the Clave field value
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetClaveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clave, true
}

// SetClave sets field value
func (o *NominaPercepcionInput) SetClave(v string) {
	o.Clave = v
}

// GetImporteGravado returns the ImporteGravado field value
func (o *NominaPercepcionInput) GetImporteGravado() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ImporteGravado
}

// GetImporteGravadoOk returns a tuple with the ImporteGravado field value
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetImporteGravadoOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImporteGravado, true
}

// SetImporteGravado sets field value
func (o *NominaPercepcionInput) SetImporteGravado(v float32) {
	o.ImporteGravado = v
}

// GetImporteExento returns the ImporteExento field value
func (o *NominaPercepcionInput) GetImporteExento() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ImporteExento
}

// GetImporteExentoOk returns a tuple with the ImporteExento field value
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetImporteExentoOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImporteExento, true
}

// SetImporteExento sets field value
func (o *NominaPercepcionInput) SetImporteExento(v float32) {
	o.ImporteExento = v
}

// GetAccionesOTitulos returns the AccionesOTitulos field value if set, zero value otherwise.
func (o *NominaPercepcionInput) GetAccionesOTitulos() NominaAccionesInput {
	if o == nil || IsNil(o.AccionesOTitulos) {
		var ret NominaAccionesInput
		return ret
	}
	return *o.AccionesOTitulos
}

// GetAccionesOTitulosOk returns a tuple with the AccionesOTitulos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetAccionesOTitulosOk() (*NominaAccionesInput, bool) {
	if o == nil || IsNil(o.AccionesOTitulos) {
		return nil, false
	}
	return o.AccionesOTitulos, true
}

// HasAccionesOTitulos returns a boolean if a field has been set.
func (o *NominaPercepcionInput) HasAccionesOTitulos() bool {
	if o != nil && !IsNil(o.AccionesOTitulos) {
		return true
	}

	return false
}

// SetAccionesOTitulos gets a reference to the given NominaAccionesInput and assigns it to the AccionesOTitulos field.
func (o *NominaPercepcionInput) SetAccionesOTitulos(v NominaAccionesInput) {
	o.AccionesOTitulos = &v
}

// GetHorasExtra returns the HorasExtra field value if set, zero value otherwise.
func (o *NominaPercepcionInput) GetHorasExtra() []NominaHorasExtraInput {
	if o == nil || IsNil(o.HorasExtra) {
		var ret []NominaHorasExtraInput
		return ret
	}
	return o.HorasExtra
}

// GetHorasExtraOk returns a tuple with the HorasExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionInput) GetHorasExtraOk() ([]NominaHorasExtraInput, bool) {
	if o == nil || IsNil(o.HorasExtra) {
		return nil, false
	}
	return o.HorasExtra, true
}

// HasHorasExtra returns a boolean if a field has been set.
func (o *NominaPercepcionInput) HasHorasExtra() bool {
	if o != nil && !IsNil(o.HorasExtra) {
		return true
	}

	return false
}

// SetHorasExtra gets a reference to the given []NominaHorasExtraInput and assigns it to the HorasExtra field.
func (o *NominaPercepcionInput) SetHorasExtra(v []NominaHorasExtraInput) {
	o.HorasExtra = v
}

func (o NominaPercepcionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaPercepcionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tipo_percepcion"] = o.TipoPercepcion
	if !IsNil(o.Concepto) {
		toSerialize["concepto"] = o.Concepto
	}
	toSerialize["clave"] = o.Clave
	toSerialize["importe_gravado"] = o.ImporteGravado
	toSerialize["importe_exento"] = o.ImporteExento
	if !IsNil(o.AccionesOTitulos) {
		toSerialize["acciones_o_titulos"] = o.AccionesOTitulos
	}
	if !IsNil(o.HorasExtra) {
		toSerialize["horas_extra"] = o.HorasExtra
	}
	return toSerialize, nil
}

type NullableNominaPercepcionInput struct {
	value *NominaPercepcionInput
	isSet bool
}

func (v NullableNominaPercepcionInput) Get() *NominaPercepcionInput {
	return v.value
}

func (v *NullableNominaPercepcionInput) Set(val *NominaPercepcionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaPercepcionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaPercepcionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaPercepcionInput(val *NominaPercepcionInput) *NullableNominaPercepcionInput {
	return &NullableNominaPercepcionInput{value: val, isSet: true}
}

func (v NullableNominaPercepcionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaPercepcionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


