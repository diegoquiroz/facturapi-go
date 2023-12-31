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

// checks if the CheckGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckGet200Response{}

// CheckGet200Response a
type CheckGet200Response struct {
	Ok *bool `json:"ok,omitempty"`
}

// NewCheckGet200Response instantiates a new CheckGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckGet200Response() *CheckGet200Response {
	this := CheckGet200Response{}
	return &this
}

// NewCheckGet200ResponseWithDefaults instantiates a new CheckGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckGet200ResponseWithDefaults() *CheckGet200Response {
	this := CheckGet200Response{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *CheckGet200Response) GetOk() bool {
	if o == nil || IsNil(o.Ok) {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckGet200Response) GetOkOk() (*bool, bool) {
	if o == nil || IsNil(o.Ok) {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *CheckGet200Response) HasOk() bool {
	if o != nil && !IsNil(o.Ok) {
		return true
	}

	return false
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *CheckGet200Response) SetOk(v bool) {
	o.Ok = &v
}

func (o CheckGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ok) {
		toSerialize["ok"] = o.Ok
	}
	return toSerialize, nil
}

type NullableCheckGet200Response struct {
	value *CheckGet200Response
	isSet bool
}

func (v NullableCheckGet200Response) Get() *CheckGet200Response {
	return v.value
}

func (v *NullableCheckGet200Response) Set(val *CheckGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckGet200Response(val *CheckGet200Response) *NullableCheckGet200Response {
	return &NullableCheckGet200Response{value: val, isSet: true}
}

func (v NullableCheckGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


