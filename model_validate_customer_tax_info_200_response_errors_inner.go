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

// checks if the ValidateCustomerTaxInfo200ResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateCustomerTaxInfo200ResponseErrorsInner{}

// ValidateCustomerTaxInfo200ResponseErrorsInner struct for ValidateCustomerTaxInfo200ResponseErrorsInner
type ValidateCustomerTaxInfo200ResponseErrorsInner struct {
	// Nombre del campo que no coincide con los registros del SAT
	Path *string `json:"path,omitempty"`
	// Mensaje de error
	Message *string `json:"message,omitempty"`
}

// NewValidateCustomerTaxInfo200ResponseErrorsInner instantiates a new ValidateCustomerTaxInfo200ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateCustomerTaxInfo200ResponseErrorsInner() *ValidateCustomerTaxInfo200ResponseErrorsInner {
	this := ValidateCustomerTaxInfo200ResponseErrorsInner{}
	return &this
}

// NewValidateCustomerTaxInfo200ResponseErrorsInnerWithDefaults instantiates a new ValidateCustomerTaxInfo200ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateCustomerTaxInfo200ResponseErrorsInnerWithDefaults() *ValidateCustomerTaxInfo200ResponseErrorsInner {
	this := ValidateCustomerTaxInfo200ResponseErrorsInner{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) SetPath(v string) {
	o.Path = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidateCustomerTaxInfo200ResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

func (o ValidateCustomerTaxInfo200ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateCustomerTaxInfo200ResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableValidateCustomerTaxInfo200ResponseErrorsInner struct {
	value *ValidateCustomerTaxInfo200ResponseErrorsInner
	isSet bool
}

func (v NullableValidateCustomerTaxInfo200ResponseErrorsInner) Get() *ValidateCustomerTaxInfo200ResponseErrorsInner {
	return v.value
}

func (v *NullableValidateCustomerTaxInfo200ResponseErrorsInner) Set(val *ValidateCustomerTaxInfo200ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateCustomerTaxInfo200ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateCustomerTaxInfo200ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateCustomerTaxInfo200ResponseErrorsInner(val *ValidateCustomerTaxInfo200ResponseErrorsInner) *NullableValidateCustomerTaxInfo200ResponseErrorsInner {
	return &NullableValidateCustomerTaxInfo200ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableValidateCustomerTaxInfo200ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateCustomerTaxInfo200ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


