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

// checks if the ValidateCustomerTaxInfo200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateCustomerTaxInfo200Response{}

// ValidateCustomerTaxInfo200Response struct for ValidateCustomerTaxInfo200Response
type ValidateCustomerTaxInfo200Response struct {
	// Indica si la información fiscal del cliente coincide con los registros del SAT
	IsValid *bool `json:"is_valid,omitempty"`
	Errors []ValidateCustomerTaxInfo200ResponseErrorsInner `json:"errors,omitempty"`
}

// NewValidateCustomerTaxInfo200Response instantiates a new ValidateCustomerTaxInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateCustomerTaxInfo200Response() *ValidateCustomerTaxInfo200Response {
	this := ValidateCustomerTaxInfo200Response{}
	return &this
}

// NewValidateCustomerTaxInfo200ResponseWithDefaults instantiates a new ValidateCustomerTaxInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateCustomerTaxInfo200ResponseWithDefaults() *ValidateCustomerTaxInfo200Response {
	this := ValidateCustomerTaxInfo200Response{}
	return &this
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *ValidateCustomerTaxInfo200Response) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateCustomerTaxInfo200Response) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *ValidateCustomerTaxInfo200Response) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *ValidateCustomerTaxInfo200Response) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidateCustomerTaxInfo200Response) GetErrors() []ValidateCustomerTaxInfo200ResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []ValidateCustomerTaxInfo200ResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateCustomerTaxInfo200Response) GetErrorsOk() ([]ValidateCustomerTaxInfo200ResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidateCustomerTaxInfo200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ValidateCustomerTaxInfo200ResponseErrorsInner and assigns it to the Errors field.
func (o *ValidateCustomerTaxInfo200Response) SetErrors(v []ValidateCustomerTaxInfo200ResponseErrorsInner) {
	o.Errors = v
}

func (o ValidateCustomerTaxInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateCustomerTaxInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsValid) {
		toSerialize["is_valid"] = o.IsValid
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableValidateCustomerTaxInfo200Response struct {
	value *ValidateCustomerTaxInfo200Response
	isSet bool
}

func (v NullableValidateCustomerTaxInfo200Response) Get() *ValidateCustomerTaxInfo200Response {
	return v.value
}

func (v *NullableValidateCustomerTaxInfo200Response) Set(val *ValidateCustomerTaxInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateCustomerTaxInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateCustomerTaxInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateCustomerTaxInfo200Response(val *ValidateCustomerTaxInfo200Response) *NullableValidateCustomerTaxInfo200Response {
	return &NullableValidateCustomerTaxInfo200Response{value: val, isSet: true}
}

func (v NullableValidateCustomerTaxInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateCustomerTaxInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


