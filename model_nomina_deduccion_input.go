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

// checks if the NominaDeduccionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaDeduccionInput{}

// NominaDeduccionInput struct for NominaDeduccionInput
type NominaDeduccionInput struct {
	// Clave del catálogo [Tipo de deducción](#tipo-de-deducción).
	TipoDeduccion string `json:"tipo_deduccion"`
	// Concepto de la deducción. Si no se envía, se utilizará la descripción del catálogo del tipo de deducción.
	Concepto *string `json:"concepto,omitempty"`
	// Clave de control interno que asigna el patrón a cada deducción (descuento) de nómina propia de su contabilidad.
	Clave string `json:"clave"`
	// Importe del concepto de deducción.
	Importe float32 `json:"importe"`
}

// NewNominaDeduccionInput instantiates a new NominaDeduccionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaDeduccionInput(tipoDeduccion string, clave string, importe float32) *NominaDeduccionInput {
	this := NominaDeduccionInput{}
	this.TipoDeduccion = tipoDeduccion
	this.Clave = clave
	this.Importe = importe
	return &this
}

// NewNominaDeduccionInputWithDefaults instantiates a new NominaDeduccionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaDeduccionInputWithDefaults() *NominaDeduccionInput {
	this := NominaDeduccionInput{}
	return &this
}

// GetTipoDeduccion returns the TipoDeduccion field value
func (o *NominaDeduccionInput) GetTipoDeduccion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TipoDeduccion
}

// GetTipoDeduccionOk returns a tuple with the TipoDeduccion field value
// and a boolean to check if the value has been set.
func (o *NominaDeduccionInput) GetTipoDeduccionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TipoDeduccion, true
}

// SetTipoDeduccion sets field value
func (o *NominaDeduccionInput) SetTipoDeduccion(v string) {
	o.TipoDeduccion = v
}

// GetConcepto returns the Concepto field value if set, zero value otherwise.
func (o *NominaDeduccionInput) GetConcepto() string {
	if o == nil || IsNil(o.Concepto) {
		var ret string
		return ret
	}
	return *o.Concepto
}

// GetConceptoOk returns a tuple with the Concepto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaDeduccionInput) GetConceptoOk() (*string, bool) {
	if o == nil || IsNil(o.Concepto) {
		return nil, false
	}
	return o.Concepto, true
}

// HasConcepto returns a boolean if a field has been set.
func (o *NominaDeduccionInput) HasConcepto() bool {
	if o != nil && !IsNil(o.Concepto) {
		return true
	}

	return false
}

// SetConcepto gets a reference to the given string and assigns it to the Concepto field.
func (o *NominaDeduccionInput) SetConcepto(v string) {
	o.Concepto = &v
}

// GetClave returns the Clave field value
func (o *NominaDeduccionInput) GetClave() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Clave
}

// GetClaveOk returns a tuple with the Clave field value
// and a boolean to check if the value has been set.
func (o *NominaDeduccionInput) GetClaveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clave, true
}

// SetClave sets field value
func (o *NominaDeduccionInput) SetClave(v string) {
	o.Clave = v
}

// GetImporte returns the Importe field value
func (o *NominaDeduccionInput) GetImporte() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Importe
}

// GetImporteOk returns a tuple with the Importe field value
// and a boolean to check if the value has been set.
func (o *NominaDeduccionInput) GetImporteOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Importe, true
}

// SetImporte sets field value
func (o *NominaDeduccionInput) SetImporte(v float32) {
	o.Importe = v
}

func (o NominaDeduccionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaDeduccionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tipo_deduccion"] = o.TipoDeduccion
	if !IsNil(o.Concepto) {
		toSerialize["concepto"] = o.Concepto
	}
	toSerialize["clave"] = o.Clave
	toSerialize["importe"] = o.Importe
	return toSerialize, nil
}

type NullableNominaDeduccionInput struct {
	value *NominaDeduccionInput
	isSet bool
}

func (v NullableNominaDeduccionInput) Get() *NominaDeduccionInput {
	return v.value
}

func (v *NullableNominaDeduccionInput) Set(val *NominaDeduccionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaDeduccionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaDeduccionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaDeduccionInput(val *NominaDeduccionInput) *NullableNominaDeduccionInput {
	return &NullableNominaDeduccionInput{value: val, isSet: true}
}

func (v NullableNominaDeduccionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaDeduccionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


