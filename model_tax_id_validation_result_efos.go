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

// checks if the TaxIdValidationResultEfos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxIdValidationResultEfos{}

// TaxIdValidationResultEfos Resultado de la validación en la lista de Empresas que Facturan Operaciones Simuladas del SAT. 
type TaxIdValidationResultEfos struct {
	// Indica si el RFC tiene algún asunto relacionado con esta lista. `true`: El RFC no está en la lista de EFOS o su situación fue apelada y resultó favorable. `false`: El RFC está registrado como “Presunto” o “Definitivo” en la lista de EFOS. 
	IsValid *bool `json:"is_valid,omitempty"`
	Data *TaxIdValidationResultEfosData `json:"data,omitempty"`
}

// NewTaxIdValidationResultEfos instantiates a new TaxIdValidationResultEfos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxIdValidationResultEfos() *TaxIdValidationResultEfos {
	this := TaxIdValidationResultEfos{}
	return &this
}

// NewTaxIdValidationResultEfosWithDefaults instantiates a new TaxIdValidationResultEfos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxIdValidationResultEfosWithDefaults() *TaxIdValidationResultEfos {
	this := TaxIdValidationResultEfos{}
	return &this
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *TaxIdValidationResultEfos) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxIdValidationResultEfos) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *TaxIdValidationResultEfos) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *TaxIdValidationResultEfos) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TaxIdValidationResultEfos) GetData() TaxIdValidationResultEfosData {
	if o == nil || IsNil(o.Data) {
		var ret TaxIdValidationResultEfosData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxIdValidationResultEfos) GetDataOk() (*TaxIdValidationResultEfosData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TaxIdValidationResultEfos) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TaxIdValidationResultEfosData and assigns it to the Data field.
func (o *TaxIdValidationResultEfos) SetData(v TaxIdValidationResultEfosData) {
	o.Data = &v
}

func (o TaxIdValidationResultEfos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxIdValidationResultEfos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsValid) {
		toSerialize["is_valid"] = o.IsValid
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableTaxIdValidationResultEfos struct {
	value *TaxIdValidationResultEfos
	isSet bool
}

func (v NullableTaxIdValidationResultEfos) Get() *TaxIdValidationResultEfos {
	return v.value
}

func (v *NullableTaxIdValidationResultEfos) Set(val *TaxIdValidationResultEfos) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxIdValidationResultEfos) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxIdValidationResultEfos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxIdValidationResultEfos(val *TaxIdValidationResultEfos) *NullableTaxIdValidationResultEfos {
	return &NullableTaxIdValidationResultEfos{value: val, isSet: true}
}

func (v NullableTaxIdValidationResultEfos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxIdValidationResultEfos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


