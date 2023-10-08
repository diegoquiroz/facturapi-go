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

// checks if the SendReceiptByEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendReceiptByEmailRequest{}

// SendReceiptByEmailRequest struct for SendReceiptByEmailRequest
type SendReceiptByEmailRequest struct {
	// Dirección de correo eletrónico
	Email NullableString `json:"email"`
}

// NewSendReceiptByEmailRequest instantiates a new SendReceiptByEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendReceiptByEmailRequest(email NullableString) *SendReceiptByEmailRequest {
	this := SendReceiptByEmailRequest{}
	this.Email = email
	return &this
}

// NewSendReceiptByEmailRequestWithDefaults instantiates a new SendReceiptByEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendReceiptByEmailRequestWithDefaults() *SendReceiptByEmailRequest {
	this := SendReceiptByEmailRequest{}
	return &this
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SendReceiptByEmailRequest) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendReceiptByEmailRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *SendReceiptByEmailRequest) SetEmail(v string) {
	o.Email.Set(&v)
}

func (o SendReceiptByEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendReceiptByEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email.Get()
	return toSerialize, nil
}

type NullableSendReceiptByEmailRequest struct {
	value *SendReceiptByEmailRequest
	isSet bool
}

func (v NullableSendReceiptByEmailRequest) Get() *SendReceiptByEmailRequest {
	return v.value
}

func (v *NullableSendReceiptByEmailRequest) Set(val *SendReceiptByEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSendReceiptByEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSendReceiptByEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendReceiptByEmailRequest(val *SendReceiptByEmailRequest) *NullableSendReceiptByEmailRequest {
	return &NullableSendReceiptByEmailRequest{value: val, isSet: true}
}

func (v NullableSendReceiptByEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendReceiptByEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

