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

// checks if the RetentionInputPeriodo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionInputPeriodo{}

// RetentionInputPeriodo Información sobre el periodo de la retención.
type RetentionInputPeriodo struct {
	// Mes inicial del periodo de la retención.
	MesIni int32 `json:"mes_ini"`
	// Mes final del periodo de la retención.
	MesFin int32 `json:"mes_fin"`
	// Año o ejercicio fiscal en que se realizó la retención.
	Ejerc int32 `json:"ejerc"`
}

// NewRetentionInputPeriodo instantiates a new RetentionInputPeriodo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionInputPeriodo(mesIni int32, mesFin int32, ejerc int32) *RetentionInputPeriodo {
	this := RetentionInputPeriodo{}
	this.MesIni = mesIni
	this.MesFin = mesFin
	this.Ejerc = ejerc
	return &this
}

// NewRetentionInputPeriodoWithDefaults instantiates a new RetentionInputPeriodo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionInputPeriodoWithDefaults() *RetentionInputPeriodo {
	this := RetentionInputPeriodo{}
	return &this
}

// GetMesIni returns the MesIni field value
func (o *RetentionInputPeriodo) GetMesIni() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MesIni
}

// GetMesIniOk returns a tuple with the MesIni field value
// and a boolean to check if the value has been set.
func (o *RetentionInputPeriodo) GetMesIniOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MesIni, true
}

// SetMesIni sets field value
func (o *RetentionInputPeriodo) SetMesIni(v int32) {
	o.MesIni = v
}

// GetMesFin returns the MesFin field value
func (o *RetentionInputPeriodo) GetMesFin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MesFin
}

// GetMesFinOk returns a tuple with the MesFin field value
// and a boolean to check if the value has been set.
func (o *RetentionInputPeriodo) GetMesFinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MesFin, true
}

// SetMesFin sets field value
func (o *RetentionInputPeriodo) SetMesFin(v int32) {
	o.MesFin = v
}

// GetEjerc returns the Ejerc field value
func (o *RetentionInputPeriodo) GetEjerc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ejerc
}

// GetEjercOk returns a tuple with the Ejerc field value
// and a boolean to check if the value has been set.
func (o *RetentionInputPeriodo) GetEjercOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ejerc, true
}

// SetEjerc sets field value
func (o *RetentionInputPeriodo) SetEjerc(v int32) {
	o.Ejerc = v
}

func (o RetentionInputPeriodo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionInputPeriodo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mes_ini"] = o.MesIni
	toSerialize["mes_fin"] = o.MesFin
	toSerialize["ejerc"] = o.Ejerc
	return toSerialize, nil
}

type NullableRetentionInputPeriodo struct {
	value *RetentionInputPeriodo
	isSet bool
}

func (v NullableRetentionInputPeriodo) Get() *RetentionInputPeriodo {
	return v.value
}

func (v *NullableRetentionInputPeriodo) Set(val *RetentionInputPeriodo) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionInputPeriodo) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionInputPeriodo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionInputPeriodo(val *RetentionInputPeriodo) *NullableRetentionInputPeriodo {
	return &NullableRetentionInputPeriodo{value: val, isSet: true}
}

func (v NullableRetentionInputPeriodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionInputPeriodo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


