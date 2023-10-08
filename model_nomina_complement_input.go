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

// checks if the NominaComplementInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaComplementInput{}

// NominaComplementInput struct for NominaComplementInput
type NominaComplementInput struct {
	NominaOrCustomComplementInput
	Data *NominaComplementDataInput `json:"data,omitempty"`
}

// NewNominaComplementInput instantiates a new NominaComplementInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaComplementInput(type_ string) *NominaComplementInput {
	this := NominaComplementInput{}
	this.Type = type_
	return &this
}

// NewNominaComplementInputWithDefaults instantiates a new NominaComplementInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaComplementInputWithDefaults() *NominaComplementInput {
	this := NominaComplementInput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *NominaComplementInput) GetData() NominaComplementDataInput {
	if o == nil || IsNil(o.Data) {
		var ret NominaComplementDataInput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementInput) GetDataOk() (*NominaComplementDataInput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *NominaComplementInput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given NominaComplementDataInput and assigns it to the Data field.
func (o *NominaComplementInput) SetData(v NominaComplementDataInput) {
	o.Data = &v
}

func (o NominaComplementInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaComplementInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedNominaOrCustomComplementInput, errNominaOrCustomComplementInput := json.Marshal(o.NominaOrCustomComplementInput)
	if errNominaOrCustomComplementInput != nil {
		return map[string]interface{}{}, errNominaOrCustomComplementInput
	}
	errNominaOrCustomComplementInput = json.Unmarshal([]byte(serializedNominaOrCustomComplementInput), &toSerialize)
	if errNominaOrCustomComplementInput != nil {
		return map[string]interface{}{}, errNominaOrCustomComplementInput
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableNominaComplementInput struct {
	value *NominaComplementInput
	isSet bool
}

func (v NullableNominaComplementInput) Get() *NominaComplementInput {
	return v.value
}

func (v *NullableNominaComplementInput) Set(val *NominaComplementInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaComplementInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaComplementInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaComplementInput(val *NominaComplementInput) *NullableNominaComplementInput {
	return &NullableNominaComplementInput{value: val, isSet: true}
}

func (v NullableNominaComplementInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaComplementInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


