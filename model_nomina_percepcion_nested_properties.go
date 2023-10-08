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

// checks if the NominaPercepcionNestedProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaPercepcionNestedProperties{}

// NominaPercepcionNestedProperties struct for NominaPercepcionNestedProperties
type NominaPercepcionNestedProperties struct {
	AccionesOTitulos *NominaAccionesProperties `json:"acciones_o_titulos,omitempty"`
	HorasExtra []NominaHorasExtraProperties `json:"horas_extra,omitempty"`
}

// NewNominaPercepcionNestedProperties instantiates a new NominaPercepcionNestedProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaPercepcionNestedProperties() *NominaPercepcionNestedProperties {
	this := NominaPercepcionNestedProperties{}
	return &this
}

// NewNominaPercepcionNestedPropertiesWithDefaults instantiates a new NominaPercepcionNestedProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaPercepcionNestedPropertiesWithDefaults() *NominaPercepcionNestedProperties {
	this := NominaPercepcionNestedProperties{}
	return &this
}

// GetAccionesOTitulos returns the AccionesOTitulos field value if set, zero value otherwise.
func (o *NominaPercepcionNestedProperties) GetAccionesOTitulos() NominaAccionesProperties {
	if o == nil || IsNil(o.AccionesOTitulos) {
		var ret NominaAccionesProperties
		return ret
	}
	return *o.AccionesOTitulos
}

// GetAccionesOTitulosOk returns a tuple with the AccionesOTitulos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionNestedProperties) GetAccionesOTitulosOk() (*NominaAccionesProperties, bool) {
	if o == nil || IsNil(o.AccionesOTitulos) {
		return nil, false
	}
	return o.AccionesOTitulos, true
}

// HasAccionesOTitulos returns a boolean if a field has been set.
func (o *NominaPercepcionNestedProperties) HasAccionesOTitulos() bool {
	if o != nil && !IsNil(o.AccionesOTitulos) {
		return true
	}

	return false
}

// SetAccionesOTitulos gets a reference to the given NominaAccionesProperties and assigns it to the AccionesOTitulos field.
func (o *NominaPercepcionNestedProperties) SetAccionesOTitulos(v NominaAccionesProperties) {
	o.AccionesOTitulos = &v
}

// GetHorasExtra returns the HorasExtra field value if set, zero value otherwise.
func (o *NominaPercepcionNestedProperties) GetHorasExtra() []NominaHorasExtraProperties {
	if o == nil || IsNil(o.HorasExtra) {
		var ret []NominaHorasExtraProperties
		return ret
	}
	return o.HorasExtra
}

// GetHorasExtraOk returns a tuple with the HorasExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionNestedProperties) GetHorasExtraOk() ([]NominaHorasExtraProperties, bool) {
	if o == nil || IsNil(o.HorasExtra) {
		return nil, false
	}
	return o.HorasExtra, true
}

// HasHorasExtra returns a boolean if a field has been set.
func (o *NominaPercepcionNestedProperties) HasHorasExtra() bool {
	if o != nil && !IsNil(o.HorasExtra) {
		return true
	}

	return false
}

// SetHorasExtra gets a reference to the given []NominaHorasExtraProperties and assigns it to the HorasExtra field.
func (o *NominaPercepcionNestedProperties) SetHorasExtra(v []NominaHorasExtraProperties) {
	o.HorasExtra = v
}

func (o NominaPercepcionNestedProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaPercepcionNestedProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccionesOTitulos) {
		toSerialize["acciones_o_titulos"] = o.AccionesOTitulos
	}
	if !IsNil(o.HorasExtra) {
		toSerialize["horas_extra"] = o.HorasExtra
	}
	return toSerialize, nil
}

type NullableNominaPercepcionNestedProperties struct {
	value *NominaPercepcionNestedProperties
	isSet bool
}

func (v NullableNominaPercepcionNestedProperties) Get() *NominaPercepcionNestedProperties {
	return v.value
}

func (v *NullableNominaPercepcionNestedProperties) Set(val *NominaPercepcionNestedProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaPercepcionNestedProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaPercepcionNestedProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaPercepcionNestedProperties(val *NominaPercepcionNestedProperties) *NullableNominaPercepcionNestedProperties {
	return &NullableNominaPercepcionNestedProperties{value: val, isSet: true}
}

func (v NullableNominaPercepcionNestedProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaPercepcionNestedProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


