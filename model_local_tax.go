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

// checks if the LocalTax type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalTax{}

// LocalTax struct for LocalTax
type LocalTax struct {
	// Tasa del impuesto en fracción decimal.
	Rate float32 `json:"rate"`
	// Base del impuesto. Valor: 100% del subtotal.
	Base *float32 `json:"base,omitempty"`
	// Nombre del impuesto. Texto libre.
	Type string `json:"type"`
	// Indica si se trata de un impuesto retenido (`true`), o un impuesto trasladado (`false`)
	Withholding *bool `json:"withholding,omitempty"`
}

// NewLocalTax instantiates a new LocalTax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalTax(rate float32, type_ string) *LocalTax {
	this := LocalTax{}
	this.Rate = rate
	this.Type = type_
	var withholding bool = false
	this.Withholding = &withholding
	return &this
}

// NewLocalTaxWithDefaults instantiates a new LocalTax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalTaxWithDefaults() *LocalTax {
	this := LocalTax{}
	var withholding bool = false
	this.Withholding = &withholding
	return &this
}

// GetRate returns the Rate field value
func (o *LocalTax) GetRate() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *LocalTax) GetRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *LocalTax) SetRate(v float32) {
	o.Rate = v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *LocalTax) GetBase() float32 {
	if o == nil || IsNil(o.Base) {
		var ret float32
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalTax) GetBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *LocalTax) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given float32 and assigns it to the Base field.
func (o *LocalTax) SetBase(v float32) {
	o.Base = &v
}

// GetType returns the Type field value
func (o *LocalTax) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LocalTax) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LocalTax) SetType(v string) {
	o.Type = v
}

// GetWithholding returns the Withholding field value if set, zero value otherwise.
func (o *LocalTax) GetWithholding() bool {
	if o == nil || IsNil(o.Withholding) {
		var ret bool
		return ret
	}
	return *o.Withholding
}

// GetWithholdingOk returns a tuple with the Withholding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalTax) GetWithholdingOk() (*bool, bool) {
	if o == nil || IsNil(o.Withholding) {
		return nil, false
	}
	return o.Withholding, true
}

// HasWithholding returns a boolean if a field has been set.
func (o *LocalTax) HasWithholding() bool {
	if o != nil && !IsNil(o.Withholding) {
		return true
	}

	return false
}

// SetWithholding gets a reference to the given bool and assigns it to the Withholding field.
func (o *LocalTax) SetWithholding(v bool) {
	o.Withholding = &v
}

func (o LocalTax) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalTax) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rate"] = o.Rate
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Withholding) {
		toSerialize["withholding"] = o.Withholding
	}
	return toSerialize, nil
}

type NullableLocalTax struct {
	value *LocalTax
	isSet bool
}

func (v NullableLocalTax) Get() *LocalTax {
	return v.value
}

func (v *NullableLocalTax) Set(val *LocalTax) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalTax) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalTax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalTax(val *LocalTax) *NullableLocalTax {
	return &NullableLocalTax{value: val, isSet: true}
}

func (v NullableLocalTax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalTax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


