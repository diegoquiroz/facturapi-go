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

// checks if the NominaEmisorProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaEmisorProperties{}

// NominaEmisorProperties Información del emisor, en caso de ser requerida.
type NominaEmisorProperties struct {
	// Requerido cuando el empleador es persona física. CURP del empleador.
	Curp *string `json:"curp,omitempty"`
	// Clave de registro patronal asignada por la institución de seguridad social al patrón.
	RegistroPatronal *string `json:"registro_patronal,omitempty"`
	// RFC de la persona que fungió como patrón. Se usa cuando el pago se realiza a través de un tercero.
	RfcPatronOrigen *string `json:"rfc_patron_origen,omitempty"`
	EntidadSncf *NominaEmisorPropertiesEntidadSncf `json:"entidad_sncf,omitempty"`
}

// NewNominaEmisorProperties instantiates a new NominaEmisorProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaEmisorProperties() *NominaEmisorProperties {
	this := NominaEmisorProperties{}
	return &this
}

// NewNominaEmisorPropertiesWithDefaults instantiates a new NominaEmisorProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaEmisorPropertiesWithDefaults() *NominaEmisorProperties {
	this := NominaEmisorProperties{}
	return &this
}

// GetCurp returns the Curp field value if set, zero value otherwise.
func (o *NominaEmisorProperties) GetCurp() string {
	if o == nil || IsNil(o.Curp) {
		var ret string
		return ret
	}
	return *o.Curp
}

// GetCurpOk returns a tuple with the Curp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaEmisorProperties) GetCurpOk() (*string, bool) {
	if o == nil || IsNil(o.Curp) {
		return nil, false
	}
	return o.Curp, true
}

// HasCurp returns a boolean if a field has been set.
func (o *NominaEmisorProperties) HasCurp() bool {
	if o != nil && !IsNil(o.Curp) {
		return true
	}

	return false
}

// SetCurp gets a reference to the given string and assigns it to the Curp field.
func (o *NominaEmisorProperties) SetCurp(v string) {
	o.Curp = &v
}

// GetRegistroPatronal returns the RegistroPatronal field value if set, zero value otherwise.
func (o *NominaEmisorProperties) GetRegistroPatronal() string {
	if o == nil || IsNil(o.RegistroPatronal) {
		var ret string
		return ret
	}
	return *o.RegistroPatronal
}

// GetRegistroPatronalOk returns a tuple with the RegistroPatronal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaEmisorProperties) GetRegistroPatronalOk() (*string, bool) {
	if o == nil || IsNil(o.RegistroPatronal) {
		return nil, false
	}
	return o.RegistroPatronal, true
}

// HasRegistroPatronal returns a boolean if a field has been set.
func (o *NominaEmisorProperties) HasRegistroPatronal() bool {
	if o != nil && !IsNil(o.RegistroPatronal) {
		return true
	}

	return false
}

// SetRegistroPatronal gets a reference to the given string and assigns it to the RegistroPatronal field.
func (o *NominaEmisorProperties) SetRegistroPatronal(v string) {
	o.RegistroPatronal = &v
}

// GetRfcPatronOrigen returns the RfcPatronOrigen field value if set, zero value otherwise.
func (o *NominaEmisorProperties) GetRfcPatronOrigen() string {
	if o == nil || IsNil(o.RfcPatronOrigen) {
		var ret string
		return ret
	}
	return *o.RfcPatronOrigen
}

// GetRfcPatronOrigenOk returns a tuple with the RfcPatronOrigen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaEmisorProperties) GetRfcPatronOrigenOk() (*string, bool) {
	if o == nil || IsNil(o.RfcPatronOrigen) {
		return nil, false
	}
	return o.RfcPatronOrigen, true
}

// HasRfcPatronOrigen returns a boolean if a field has been set.
func (o *NominaEmisorProperties) HasRfcPatronOrigen() bool {
	if o != nil && !IsNil(o.RfcPatronOrigen) {
		return true
	}

	return false
}

// SetRfcPatronOrigen gets a reference to the given string and assigns it to the RfcPatronOrigen field.
func (o *NominaEmisorProperties) SetRfcPatronOrigen(v string) {
	o.RfcPatronOrigen = &v
}

// GetEntidadSncf returns the EntidadSncf field value if set, zero value otherwise.
func (o *NominaEmisorProperties) GetEntidadSncf() NominaEmisorPropertiesEntidadSncf {
	if o == nil || IsNil(o.EntidadSncf) {
		var ret NominaEmisorPropertiesEntidadSncf
		return ret
	}
	return *o.EntidadSncf
}

// GetEntidadSncfOk returns a tuple with the EntidadSncf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaEmisorProperties) GetEntidadSncfOk() (*NominaEmisorPropertiesEntidadSncf, bool) {
	if o == nil || IsNil(o.EntidadSncf) {
		return nil, false
	}
	return o.EntidadSncf, true
}

// HasEntidadSncf returns a boolean if a field has been set.
func (o *NominaEmisorProperties) HasEntidadSncf() bool {
	if o != nil && !IsNil(o.EntidadSncf) {
		return true
	}

	return false
}

// SetEntidadSncf gets a reference to the given NominaEmisorPropertiesEntidadSncf and assigns it to the EntidadSncf field.
func (o *NominaEmisorProperties) SetEntidadSncf(v NominaEmisorPropertiesEntidadSncf) {
	o.EntidadSncf = &v
}

func (o NominaEmisorProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaEmisorProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Curp) {
		toSerialize["curp"] = o.Curp
	}
	if !IsNil(o.RegistroPatronal) {
		toSerialize["registro_patronal"] = o.RegistroPatronal
	}
	if !IsNil(o.RfcPatronOrigen) {
		toSerialize["rfc_patron_origen"] = o.RfcPatronOrigen
	}
	if !IsNil(o.EntidadSncf) {
		toSerialize["entidad_sncf"] = o.EntidadSncf
	}
	return toSerialize, nil
}

type NullableNominaEmisorProperties struct {
	value *NominaEmisorProperties
	isSet bool
}

func (v NullableNominaEmisorProperties) Get() *NominaEmisorProperties {
	return v.value
}

func (v *NullableNominaEmisorProperties) Set(val *NominaEmisorProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaEmisorProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaEmisorProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaEmisorProperties(val *NominaEmisorProperties) *NullableNominaEmisorProperties {
	return &NullableNominaEmisorProperties{value: val, isSet: true}
}

func (v NullableNominaEmisorProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaEmisorProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

