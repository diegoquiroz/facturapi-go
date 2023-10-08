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

// checks if the NominaAccionesProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaAccionesProperties{}

// NominaAccionesProperties Objeto para expresar ingresos por acciones o títulos valor que representan bienes. Es requerido cuando existan ingresos por sueldos derivados de adquisición de acciones o títulos.
type NominaAccionesProperties struct {
	// Valor de mercado de las Acciones o Títulos valor al ejercer la opción.
	ValorMercado *float32 `json:"valor_mercado,omitempty"`
	// Precio establecido al otorgarse la opción de ingresos en acciones o títulos valor.
	PrecioAlOtorgarse *float32 `json:"precio_al_otorgarse,omitempty"`
}

// NewNominaAccionesProperties instantiates a new NominaAccionesProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaAccionesProperties() *NominaAccionesProperties {
	this := NominaAccionesProperties{}
	return &this
}

// NewNominaAccionesPropertiesWithDefaults instantiates a new NominaAccionesProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaAccionesPropertiesWithDefaults() *NominaAccionesProperties {
	this := NominaAccionesProperties{}
	return &this
}

// GetValorMercado returns the ValorMercado field value if set, zero value otherwise.
func (o *NominaAccionesProperties) GetValorMercado() float32 {
	if o == nil || IsNil(o.ValorMercado) {
		var ret float32
		return ret
	}
	return *o.ValorMercado
}

// GetValorMercadoOk returns a tuple with the ValorMercado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaAccionesProperties) GetValorMercadoOk() (*float32, bool) {
	if o == nil || IsNil(o.ValorMercado) {
		return nil, false
	}
	return o.ValorMercado, true
}

// HasValorMercado returns a boolean if a field has been set.
func (o *NominaAccionesProperties) HasValorMercado() bool {
	if o != nil && !IsNil(o.ValorMercado) {
		return true
	}

	return false
}

// SetValorMercado gets a reference to the given float32 and assigns it to the ValorMercado field.
func (o *NominaAccionesProperties) SetValorMercado(v float32) {
	o.ValorMercado = &v
}

// GetPrecioAlOtorgarse returns the PrecioAlOtorgarse field value if set, zero value otherwise.
func (o *NominaAccionesProperties) GetPrecioAlOtorgarse() float32 {
	if o == nil || IsNil(o.PrecioAlOtorgarse) {
		var ret float32
		return ret
	}
	return *o.PrecioAlOtorgarse
}

// GetPrecioAlOtorgarseOk returns a tuple with the PrecioAlOtorgarse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaAccionesProperties) GetPrecioAlOtorgarseOk() (*float32, bool) {
	if o == nil || IsNil(o.PrecioAlOtorgarse) {
		return nil, false
	}
	return o.PrecioAlOtorgarse, true
}

// HasPrecioAlOtorgarse returns a boolean if a field has been set.
func (o *NominaAccionesProperties) HasPrecioAlOtorgarse() bool {
	if o != nil && !IsNil(o.PrecioAlOtorgarse) {
		return true
	}

	return false
}

// SetPrecioAlOtorgarse gets a reference to the given float32 and assigns it to the PrecioAlOtorgarse field.
func (o *NominaAccionesProperties) SetPrecioAlOtorgarse(v float32) {
	o.PrecioAlOtorgarse = &v
}

func (o NominaAccionesProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaAccionesProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValorMercado) {
		toSerialize["valor_mercado"] = o.ValorMercado
	}
	if !IsNil(o.PrecioAlOtorgarse) {
		toSerialize["precio_al_otorgarse"] = o.PrecioAlOtorgarse
	}
	return toSerialize, nil
}

type NullableNominaAccionesProperties struct {
	value *NominaAccionesProperties
	isSet bool
}

func (v NullableNominaAccionesProperties) Get() *NominaAccionesProperties {
	return v.value
}

func (v *NullableNominaAccionesProperties) Set(val *NominaAccionesProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaAccionesProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaAccionesProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaAccionesProperties(val *NominaAccionesProperties) *NullableNominaAccionesProperties {
	return &NullableNominaAccionesProperties{value: val, isSet: true}
}

func (v NullableNominaAccionesProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaAccionesProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


