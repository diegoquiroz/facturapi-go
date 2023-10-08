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

// checks if the NominaSubContratacionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaSubContratacionProperties{}

// NominaSubContratacionProperties struct for NominaSubContratacionProperties
type NominaSubContratacionProperties struct {
	// RFC de la persona o empresa que subcontrata, es decir, de la persona o empresa en donde el trabajador prestó directamente sus servicios.
	RfcLabora *string `json:"rfc_labora,omitempty"`
	// Porcentaje de tiempo en que el trabajador prestó sus servicios a la persona o empresa que lo subcontrató.
	PorcentajeTiempo *float32 `json:"porcentaje_tiempo,omitempty"`
}

// NewNominaSubContratacionProperties instantiates a new NominaSubContratacionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaSubContratacionProperties() *NominaSubContratacionProperties {
	this := NominaSubContratacionProperties{}
	return &this
}

// NewNominaSubContratacionPropertiesWithDefaults instantiates a new NominaSubContratacionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaSubContratacionPropertiesWithDefaults() *NominaSubContratacionProperties {
	this := NominaSubContratacionProperties{}
	return &this
}

// GetRfcLabora returns the RfcLabora field value if set, zero value otherwise.
func (o *NominaSubContratacionProperties) GetRfcLabora() string {
	if o == nil || IsNil(o.RfcLabora) {
		var ret string
		return ret
	}
	return *o.RfcLabora
}

// GetRfcLaboraOk returns a tuple with the RfcLabora field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaSubContratacionProperties) GetRfcLaboraOk() (*string, bool) {
	if o == nil || IsNil(o.RfcLabora) {
		return nil, false
	}
	return o.RfcLabora, true
}

// HasRfcLabora returns a boolean if a field has been set.
func (o *NominaSubContratacionProperties) HasRfcLabora() bool {
	if o != nil && !IsNil(o.RfcLabora) {
		return true
	}

	return false
}

// SetRfcLabora gets a reference to the given string and assigns it to the RfcLabora field.
func (o *NominaSubContratacionProperties) SetRfcLabora(v string) {
	o.RfcLabora = &v
}

// GetPorcentajeTiempo returns the PorcentajeTiempo field value if set, zero value otherwise.
func (o *NominaSubContratacionProperties) GetPorcentajeTiempo() float32 {
	if o == nil || IsNil(o.PorcentajeTiempo) {
		var ret float32
		return ret
	}
	return *o.PorcentajeTiempo
}

// GetPorcentajeTiempoOk returns a tuple with the PorcentajeTiempo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaSubContratacionProperties) GetPorcentajeTiempoOk() (*float32, bool) {
	if o == nil || IsNil(o.PorcentajeTiempo) {
		return nil, false
	}
	return o.PorcentajeTiempo, true
}

// HasPorcentajeTiempo returns a boolean if a field has been set.
func (o *NominaSubContratacionProperties) HasPorcentajeTiempo() bool {
	if o != nil && !IsNil(o.PorcentajeTiempo) {
		return true
	}

	return false
}

// SetPorcentajeTiempo gets a reference to the given float32 and assigns it to the PorcentajeTiempo field.
func (o *NominaSubContratacionProperties) SetPorcentajeTiempo(v float32) {
	o.PorcentajeTiempo = &v
}

func (o NominaSubContratacionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaSubContratacionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RfcLabora) {
		toSerialize["rfc_labora"] = o.RfcLabora
	}
	if !IsNil(o.PorcentajeTiempo) {
		toSerialize["porcentaje_tiempo"] = o.PorcentajeTiempo
	}
	return toSerialize, nil
}

type NullableNominaSubContratacionProperties struct {
	value *NominaSubContratacionProperties
	isSet bool
}

func (v NullableNominaSubContratacionProperties) Get() *NominaSubContratacionProperties {
	return v.value
}

func (v *NullableNominaSubContratacionProperties) Set(val *NominaSubContratacionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaSubContratacionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaSubContratacionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaSubContratacionProperties(val *NominaSubContratacionProperties) *NullableNominaSubContratacionProperties {
	return &NullableNominaSubContratacionProperties{value: val, isSet: true}
}

func (v NullableNominaSubContratacionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaSubContratacionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


