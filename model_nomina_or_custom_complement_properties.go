/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi-go

import (
	"encoding/json"
)

// checks if the NominaOrCustomComplementProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaOrCustomComplementProperties{}

// NominaOrCustomComplementProperties struct for NominaOrCustomComplementProperties
type NominaOrCustomComplementProperties struct {
	// Tipo de complemento.
	Type *string `json:"type,omitempty"`
}

// NewNominaOrCustomComplementProperties instantiates a new NominaOrCustomComplementProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaOrCustomComplementProperties() *NominaOrCustomComplementProperties {
	this := NominaOrCustomComplementProperties{}
	return &this
}

// NewNominaOrCustomComplementPropertiesWithDefaults instantiates a new NominaOrCustomComplementProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaOrCustomComplementPropertiesWithDefaults() *NominaOrCustomComplementProperties {
	this := NominaOrCustomComplementProperties{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NominaOrCustomComplementProperties) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaOrCustomComplementProperties) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NominaOrCustomComplementProperties) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NominaOrCustomComplementProperties) SetType(v string) {
	o.Type = &v
}

func (o NominaOrCustomComplementProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaOrCustomComplementProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNominaOrCustomComplementProperties struct {
	value *NominaOrCustomComplementProperties
	isSet bool
}

func (v NullableNominaOrCustomComplementProperties) Get() *NominaOrCustomComplementProperties {
	return v.value
}

func (v *NullableNominaOrCustomComplementProperties) Set(val *NominaOrCustomComplementProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaOrCustomComplementProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaOrCustomComplementProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaOrCustomComplementProperties(val *NominaOrCustomComplementProperties) *NullableNominaOrCustomComplementProperties {
	return &NullableNominaOrCustomComplementProperties{value: val, isSet: true}
}

func (v NullableNominaOrCustomComplementProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaOrCustomComplementProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


