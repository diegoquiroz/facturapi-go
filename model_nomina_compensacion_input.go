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

// checks if the NominaCompensacionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaCompensacionInput{}

// NominaCompensacionInput struct for NominaCompensacionInput
type NominaCompensacionInput struct {
	// Monto por saldo a favor determinado por el patrón al trabajador en periodos o ejercicios anteriores.
	SaldoAFavor float32 `json:"saldo_a_favor"`
	// Año en que se determinó el saldo a favor del trabajador.
	Ano int32 `json:"ano"`
	// Remanente del saldo a favor del trabajador.
	RemanenteSalFav float32 `json:"remanente_sal_fav"`
}

// NewNominaCompensacionInput instantiates a new NominaCompensacionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaCompensacionInput(saldoAFavor float32, ano int32, remanenteSalFav float32) *NominaCompensacionInput {
	this := NominaCompensacionInput{}
	this.SaldoAFavor = saldoAFavor
	this.Ano = ano
	this.RemanenteSalFav = remanenteSalFav
	return &this
}

// NewNominaCompensacionInputWithDefaults instantiates a new NominaCompensacionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaCompensacionInputWithDefaults() *NominaCompensacionInput {
	this := NominaCompensacionInput{}
	return &this
}

// GetSaldoAFavor returns the SaldoAFavor field value
func (o *NominaCompensacionInput) GetSaldoAFavor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SaldoAFavor
}

// GetSaldoAFavorOk returns a tuple with the SaldoAFavor field value
// and a boolean to check if the value has been set.
func (o *NominaCompensacionInput) GetSaldoAFavorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaldoAFavor, true
}

// SetSaldoAFavor sets field value
func (o *NominaCompensacionInput) SetSaldoAFavor(v float32) {
	o.SaldoAFavor = v
}

// GetAno returns the Ano field value
func (o *NominaCompensacionInput) GetAno() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ano
}

// GetAnoOk returns a tuple with the Ano field value
// and a boolean to check if the value has been set.
func (o *NominaCompensacionInput) GetAnoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ano, true
}

// SetAno sets field value
func (o *NominaCompensacionInput) SetAno(v int32) {
	o.Ano = v
}

// GetRemanenteSalFav returns the RemanenteSalFav field value
func (o *NominaCompensacionInput) GetRemanenteSalFav() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemanenteSalFav
}

// GetRemanenteSalFavOk returns a tuple with the RemanenteSalFav field value
// and a boolean to check if the value has been set.
func (o *NominaCompensacionInput) GetRemanenteSalFavOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemanenteSalFav, true
}

// SetRemanenteSalFav sets field value
func (o *NominaCompensacionInput) SetRemanenteSalFav(v float32) {
	o.RemanenteSalFav = v
}

func (o NominaCompensacionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaCompensacionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["saldo_a_favor"] = o.SaldoAFavor
	toSerialize["ano"] = o.Ano
	toSerialize["remanente_sal_fav"] = o.RemanenteSalFav
	return toSerialize, nil
}

type NullableNominaCompensacionInput struct {
	value *NominaCompensacionInput
	isSet bool
}

func (v NullableNominaCompensacionInput) Get() *NominaCompensacionInput {
	return v.value
}

func (v *NullableNominaCompensacionInput) Set(val *NominaCompensacionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaCompensacionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaCompensacionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaCompensacionInput(val *NominaCompensacionInput) *NullableNominaCompensacionInput {
	return &NullableNominaCompensacionInput{value: val, isSet: true}
}

func (v NullableNominaCompensacionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaCompensacionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


