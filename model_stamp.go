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
	"time"
)

// checks if the Stamp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Stamp{}

// Stamp Información sobre el timbre fiscal digital agregado por el PAC.
type Stamp struct {
	// Sello digital del comprobante fiscal.
	Signature *string `json:"signature,omitempty"`
	// Fecha de timbrado en formato ISO8601 (UTC String).
	Date *time.Time `json:"date,omitempty"`
	// Número de serie del certificado del SAT usado para timbrar.
	SatCertNumber *string `json:"sat_cert_number,omitempty"`
	// Sello digital del timbre fiscal digital.
	SatSignature *string `json:"sat_signature,omitempty"`
}

// NewStamp instantiates a new Stamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStamp() *Stamp {
	this := Stamp{}
	return &this
}

// NewStampWithDefaults instantiates a new Stamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStampWithDefaults() *Stamp {
	this := Stamp{}
	return &this
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *Stamp) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stamp) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *Stamp) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *Stamp) SetSignature(v string) {
	o.Signature = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Stamp) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stamp) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Stamp) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *Stamp) SetDate(v time.Time) {
	o.Date = &v
}

// GetSatCertNumber returns the SatCertNumber field value if set, zero value otherwise.
func (o *Stamp) GetSatCertNumber() string {
	if o == nil || IsNil(o.SatCertNumber) {
		var ret string
		return ret
	}
	return *o.SatCertNumber
}

// GetSatCertNumberOk returns a tuple with the SatCertNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stamp) GetSatCertNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SatCertNumber) {
		return nil, false
	}
	return o.SatCertNumber, true
}

// HasSatCertNumber returns a boolean if a field has been set.
func (o *Stamp) HasSatCertNumber() bool {
	if o != nil && !IsNil(o.SatCertNumber) {
		return true
	}

	return false
}

// SetSatCertNumber gets a reference to the given string and assigns it to the SatCertNumber field.
func (o *Stamp) SetSatCertNumber(v string) {
	o.SatCertNumber = &v
}

// GetSatSignature returns the SatSignature field value if set, zero value otherwise.
func (o *Stamp) GetSatSignature() string {
	if o == nil || IsNil(o.SatSignature) {
		var ret string
		return ret
	}
	return *o.SatSignature
}

// GetSatSignatureOk returns a tuple with the SatSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stamp) GetSatSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.SatSignature) {
		return nil, false
	}
	return o.SatSignature, true
}

// HasSatSignature returns a boolean if a field has been set.
func (o *Stamp) HasSatSignature() bool {
	if o != nil && !IsNil(o.SatSignature) {
		return true
	}

	return false
}

// SetSatSignature gets a reference to the given string and assigns it to the SatSignature field.
func (o *Stamp) SetSatSignature(v string) {
	o.SatSignature = &v
}

func (o Stamp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Stamp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.SatCertNumber) {
		toSerialize["sat_cert_number"] = o.SatCertNumber
	}
	if !IsNil(o.SatSignature) {
		toSerialize["sat_signature"] = o.SatSignature
	}
	return toSerialize, nil
}

type NullableStamp struct {
	value *Stamp
	isSet bool
}

func (v NullableStamp) Get() *Stamp {
	return v.value
}

func (v *NullableStamp) Set(val *Stamp) {
	v.value = val
	v.isSet = true
}

func (v NullableStamp) IsSet() bool {
	return v.isSet
}

func (v *NullableStamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStamp(val *Stamp) *NullableStamp {
	return &NullableStamp{value: val, isSet: true}
}

func (v NullableStamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


