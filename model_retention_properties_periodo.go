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

// checks if the RetentionPropertiesPeriodo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionPropertiesPeriodo{}

// RetentionPropertiesPeriodo Información sobre el periodo de la retención.
type RetentionPropertiesPeriodo struct {
	// Mes inicial del periodo de la retención.
	MesIni *int32 `json:"mes_ini,omitempty"`
	// Mes final del periodo de la retención.
	MesFin *int32 `json:"mes_fin,omitempty"`
	// Año o ejercicio fiscal en que se realizó la retención.
	Ejerc *int32 `json:"ejerc,omitempty"`
}

// NewRetentionPropertiesPeriodo instantiates a new RetentionPropertiesPeriodo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionPropertiesPeriodo() *RetentionPropertiesPeriodo {
	this := RetentionPropertiesPeriodo{}
	return &this
}

// NewRetentionPropertiesPeriodoWithDefaults instantiates a new RetentionPropertiesPeriodo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionPropertiesPeriodoWithDefaults() *RetentionPropertiesPeriodo {
	this := RetentionPropertiesPeriodo{}
	return &this
}

// GetMesIni returns the MesIni field value if set, zero value otherwise.
func (o *RetentionPropertiesPeriodo) GetMesIni() int32 {
	if o == nil || IsNil(o.MesIni) {
		var ret int32
		return ret
	}
	return *o.MesIni
}

// GetMesIniOk returns a tuple with the MesIni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesPeriodo) GetMesIniOk() (*int32, bool) {
	if o == nil || IsNil(o.MesIni) {
		return nil, false
	}
	return o.MesIni, true
}

// HasMesIni returns a boolean if a field has been set.
func (o *RetentionPropertiesPeriodo) HasMesIni() bool {
	if o != nil && !IsNil(o.MesIni) {
		return true
	}

	return false
}

// SetMesIni gets a reference to the given int32 and assigns it to the MesIni field.
func (o *RetentionPropertiesPeriodo) SetMesIni(v int32) {
	o.MesIni = &v
}

// GetMesFin returns the MesFin field value if set, zero value otherwise.
func (o *RetentionPropertiesPeriodo) GetMesFin() int32 {
	if o == nil || IsNil(o.MesFin) {
		var ret int32
		return ret
	}
	return *o.MesFin
}

// GetMesFinOk returns a tuple with the MesFin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesPeriodo) GetMesFinOk() (*int32, bool) {
	if o == nil || IsNil(o.MesFin) {
		return nil, false
	}
	return o.MesFin, true
}

// HasMesFin returns a boolean if a field has been set.
func (o *RetentionPropertiesPeriodo) HasMesFin() bool {
	if o != nil && !IsNil(o.MesFin) {
		return true
	}

	return false
}

// SetMesFin gets a reference to the given int32 and assigns it to the MesFin field.
func (o *RetentionPropertiesPeriodo) SetMesFin(v int32) {
	o.MesFin = &v
}

// GetEjerc returns the Ejerc field value if set, zero value otherwise.
func (o *RetentionPropertiesPeriodo) GetEjerc() int32 {
	if o == nil || IsNil(o.Ejerc) {
		var ret int32
		return ret
	}
	return *o.Ejerc
}

// GetEjercOk returns a tuple with the Ejerc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesPeriodo) GetEjercOk() (*int32, bool) {
	if o == nil || IsNil(o.Ejerc) {
		return nil, false
	}
	return o.Ejerc, true
}

// HasEjerc returns a boolean if a field has been set.
func (o *RetentionPropertiesPeriodo) HasEjerc() bool {
	if o != nil && !IsNil(o.Ejerc) {
		return true
	}

	return false
}

// SetEjerc gets a reference to the given int32 and assigns it to the Ejerc field.
func (o *RetentionPropertiesPeriodo) SetEjerc(v int32) {
	o.Ejerc = &v
}

func (o RetentionPropertiesPeriodo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionPropertiesPeriodo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MesIni) {
		toSerialize["mes_ini"] = o.MesIni
	}
	if !IsNil(o.MesFin) {
		toSerialize["mes_fin"] = o.MesFin
	}
	if !IsNil(o.Ejerc) {
		toSerialize["ejerc"] = o.Ejerc
	}
	return toSerialize, nil
}

type NullableRetentionPropertiesPeriodo struct {
	value *RetentionPropertiesPeriodo
	isSet bool
}

func (v NullableRetentionPropertiesPeriodo) Get() *RetentionPropertiesPeriodo {
	return v.value
}

func (v *NullableRetentionPropertiesPeriodo) Set(val *RetentionPropertiesPeriodo) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionPropertiesPeriodo) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionPropertiesPeriodo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionPropertiesPeriodo(val *RetentionPropertiesPeriodo) *NullableRetentionPropertiesPeriodo {
	return &NullableRetentionPropertiesPeriodo{value: val, isSet: true}
}

func (v NullableRetentionPropertiesPeriodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionPropertiesPeriodo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

