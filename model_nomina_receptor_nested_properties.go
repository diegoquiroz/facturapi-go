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

// checks if the NominaReceptorNestedProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaReceptorNestedProperties{}

// NominaReceptorNestedProperties struct for NominaReceptorNestedProperties
type NominaReceptorNestedProperties struct {
	SubContratacion []NominaSubContratacionProperties `json:"sub_contratacion,omitempty"`
}

// NewNominaReceptorNestedProperties instantiates a new NominaReceptorNestedProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaReceptorNestedProperties() *NominaReceptorNestedProperties {
	this := NominaReceptorNestedProperties{}
	return &this
}

// NewNominaReceptorNestedPropertiesWithDefaults instantiates a new NominaReceptorNestedProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaReceptorNestedPropertiesWithDefaults() *NominaReceptorNestedProperties {
	this := NominaReceptorNestedProperties{}
	return &this
}

// GetSubContratacion returns the SubContratacion field value if set, zero value otherwise.
func (o *NominaReceptorNestedProperties) GetSubContratacion() []NominaSubContratacionProperties {
	if o == nil || IsNil(o.SubContratacion) {
		var ret []NominaSubContratacionProperties
		return ret
	}
	return o.SubContratacion
}

// GetSubContratacionOk returns a tuple with the SubContratacion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaReceptorNestedProperties) GetSubContratacionOk() ([]NominaSubContratacionProperties, bool) {
	if o == nil || IsNil(o.SubContratacion) {
		return nil, false
	}
	return o.SubContratacion, true
}

// HasSubContratacion returns a boolean if a field has been set.
func (o *NominaReceptorNestedProperties) HasSubContratacion() bool {
	if o != nil && !IsNil(o.SubContratacion) {
		return true
	}

	return false
}

// SetSubContratacion gets a reference to the given []NominaSubContratacionProperties and assigns it to the SubContratacion field.
func (o *NominaReceptorNestedProperties) SetSubContratacion(v []NominaSubContratacionProperties) {
	o.SubContratacion = v
}

func (o NominaReceptorNestedProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaReceptorNestedProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubContratacion) {
		toSerialize["sub_contratacion"] = o.SubContratacion
	}
	return toSerialize, nil
}

type NullableNominaReceptorNestedProperties struct {
	value *NominaReceptorNestedProperties
	isSet bool
}

func (v NullableNominaReceptorNestedProperties) Get() *NominaReceptorNestedProperties {
	return v.value
}

func (v *NullableNominaReceptorNestedProperties) Set(val *NominaReceptorNestedProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaReceptorNestedProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaReceptorNestedProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaReceptorNestedProperties(val *NominaReceptorNestedProperties) *NullableNominaReceptorNestedProperties {
	return &NullableNominaReceptorNestedProperties{value: val, isSet: true}
}

func (v NullableNominaReceptorNestedProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaReceptorNestedProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

