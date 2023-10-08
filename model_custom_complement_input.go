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

// checks if the CustomComplementInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomComplementInput{}

// CustomComplementInput struct for CustomComplementInput
type CustomComplementInput struct {
	// Tipo de complemento.
	Type string `json:"type"`
	// Código XML de tu complemento tal cual como quieres que se inserte en el XML. Debe contener sólamente un nodo XML raíz.
	Data string `json:"data"`
}

// NewCustomComplementInput instantiates a new CustomComplementInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomComplementInput(type_ string, data string) *CustomComplementInput {
	this := CustomComplementInput{}
	this.Type = type_
	this.Data = data
	return &this
}

// NewCustomComplementInputWithDefaults instantiates a new CustomComplementInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomComplementInputWithDefaults() *CustomComplementInput {
	this := CustomComplementInput{}
	return &this
}

// GetType returns the Type field value
func (o *CustomComplementInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomComplementInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomComplementInput) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *CustomComplementInput) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomComplementInput) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomComplementInput) SetData(v string) {
	o.Data = v
}

func (o CustomComplementInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomComplementInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCustomComplementInput struct {
	value *CustomComplementInput
	isSet bool
}

func (v NullableCustomComplementInput) Get() *CustomComplementInput {
	return v.value
}

func (v *NullableCustomComplementInput) Set(val *CustomComplementInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomComplementInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomComplementInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomComplementInput(val *CustomComplementInput) *NullableCustomComplementInput {
	return &NullableCustomComplementInput{value: val, isSet: true}
}

func (v NullableCustomComplementInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomComplementInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


