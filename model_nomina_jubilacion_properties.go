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

// checks if the NominaJubilacionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaJubilacionProperties{}

// NominaJubilacionProperties Objeto con información detallada de pagos por jubilación, pensiones o haberes de retiro.
type NominaJubilacionProperties struct {
	// Monto total del pago entregado en una sola exhibición.
	TotalUnaExhibicion *float32 `json:"total_una_exhibicion,omitempty"`
	// Monto total del pago entregado en parcialidades.
	TotalParcialidad *float32 `json:"total_parcialidad,omitempty"`
	// Monto diario percibido por el trabajador cuando el pago se realiza en parcialidades.
	MontoDiario *float32 `json:"monto_diario,omitempty"`
	// Ingresos acumulables percibidos por el trabajador.
	IngresoAcumulable *float32 `json:"ingreso_acumulable,omitempty"`
	// Ingresos no acumulables percibidos por el trabajador.
	IngresoNoAcumulable *float32 `json:"ingreso_no_acumulable,omitempty"`
}

// NewNominaJubilacionProperties instantiates a new NominaJubilacionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaJubilacionProperties() *NominaJubilacionProperties {
	this := NominaJubilacionProperties{}
	return &this
}

// NewNominaJubilacionPropertiesWithDefaults instantiates a new NominaJubilacionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaJubilacionPropertiesWithDefaults() *NominaJubilacionProperties {
	this := NominaJubilacionProperties{}
	return &this
}

// GetTotalUnaExhibicion returns the TotalUnaExhibicion field value if set, zero value otherwise.
func (o *NominaJubilacionProperties) GetTotalUnaExhibicion() float32 {
	if o == nil || IsNil(o.TotalUnaExhibicion) {
		var ret float32
		return ret
	}
	return *o.TotalUnaExhibicion
}

// GetTotalUnaExhibicionOk returns a tuple with the TotalUnaExhibicion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaJubilacionProperties) GetTotalUnaExhibicionOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalUnaExhibicion) {
		return nil, false
	}
	return o.TotalUnaExhibicion, true
}

// HasTotalUnaExhibicion returns a boolean if a field has been set.
func (o *NominaJubilacionProperties) HasTotalUnaExhibicion() bool {
	if o != nil && !IsNil(o.TotalUnaExhibicion) {
		return true
	}

	return false
}

// SetTotalUnaExhibicion gets a reference to the given float32 and assigns it to the TotalUnaExhibicion field.
func (o *NominaJubilacionProperties) SetTotalUnaExhibicion(v float32) {
	o.TotalUnaExhibicion = &v
}

// GetTotalParcialidad returns the TotalParcialidad field value if set, zero value otherwise.
func (o *NominaJubilacionProperties) GetTotalParcialidad() float32 {
	if o == nil || IsNil(o.TotalParcialidad) {
		var ret float32
		return ret
	}
	return *o.TotalParcialidad
}

// GetTotalParcialidadOk returns a tuple with the TotalParcialidad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaJubilacionProperties) GetTotalParcialidadOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalParcialidad) {
		return nil, false
	}
	return o.TotalParcialidad, true
}

// HasTotalParcialidad returns a boolean if a field has been set.
func (o *NominaJubilacionProperties) HasTotalParcialidad() bool {
	if o != nil && !IsNil(o.TotalParcialidad) {
		return true
	}

	return false
}

// SetTotalParcialidad gets a reference to the given float32 and assigns it to the TotalParcialidad field.
func (o *NominaJubilacionProperties) SetTotalParcialidad(v float32) {
	o.TotalParcialidad = &v
}

// GetMontoDiario returns the MontoDiario field value if set, zero value otherwise.
func (o *NominaJubilacionProperties) GetMontoDiario() float32 {
	if o == nil || IsNil(o.MontoDiario) {
		var ret float32
		return ret
	}
	return *o.MontoDiario
}

// GetMontoDiarioOk returns a tuple with the MontoDiario field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaJubilacionProperties) GetMontoDiarioOk() (*float32, bool) {
	if o == nil || IsNil(o.MontoDiario) {
		return nil, false
	}
	return o.MontoDiario, true
}

// HasMontoDiario returns a boolean if a field has been set.
func (o *NominaJubilacionProperties) HasMontoDiario() bool {
	if o != nil && !IsNil(o.MontoDiario) {
		return true
	}

	return false
}

// SetMontoDiario gets a reference to the given float32 and assigns it to the MontoDiario field.
func (o *NominaJubilacionProperties) SetMontoDiario(v float32) {
	o.MontoDiario = &v
}

// GetIngresoAcumulable returns the IngresoAcumulable field value if set, zero value otherwise.
func (o *NominaJubilacionProperties) GetIngresoAcumulable() float32 {
	if o == nil || IsNil(o.IngresoAcumulable) {
		var ret float32
		return ret
	}
	return *o.IngresoAcumulable
}

// GetIngresoAcumulableOk returns a tuple with the IngresoAcumulable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaJubilacionProperties) GetIngresoAcumulableOk() (*float32, bool) {
	if o == nil || IsNil(o.IngresoAcumulable) {
		return nil, false
	}
	return o.IngresoAcumulable, true
}

// HasIngresoAcumulable returns a boolean if a field has been set.
func (o *NominaJubilacionProperties) HasIngresoAcumulable() bool {
	if o != nil && !IsNil(o.IngresoAcumulable) {
		return true
	}

	return false
}

// SetIngresoAcumulable gets a reference to the given float32 and assigns it to the IngresoAcumulable field.
func (o *NominaJubilacionProperties) SetIngresoAcumulable(v float32) {
	o.IngresoAcumulable = &v
}

// GetIngresoNoAcumulable returns the IngresoNoAcumulable field value if set, zero value otherwise.
func (o *NominaJubilacionProperties) GetIngresoNoAcumulable() float32 {
	if o == nil || IsNil(o.IngresoNoAcumulable) {
		var ret float32
		return ret
	}
	return *o.IngresoNoAcumulable
}

// GetIngresoNoAcumulableOk returns a tuple with the IngresoNoAcumulable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaJubilacionProperties) GetIngresoNoAcumulableOk() (*float32, bool) {
	if o == nil || IsNil(o.IngresoNoAcumulable) {
		return nil, false
	}
	return o.IngresoNoAcumulable, true
}

// HasIngresoNoAcumulable returns a boolean if a field has been set.
func (o *NominaJubilacionProperties) HasIngresoNoAcumulable() bool {
	if o != nil && !IsNil(o.IngresoNoAcumulable) {
		return true
	}

	return false
}

// SetIngresoNoAcumulable gets a reference to the given float32 and assigns it to the IngresoNoAcumulable field.
func (o *NominaJubilacionProperties) SetIngresoNoAcumulable(v float32) {
	o.IngresoNoAcumulable = &v
}

func (o NominaJubilacionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaJubilacionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalUnaExhibicion) {
		toSerialize["total_una_exhibicion"] = o.TotalUnaExhibicion
	}
	if !IsNil(o.TotalParcialidad) {
		toSerialize["total_parcialidad"] = o.TotalParcialidad
	}
	if !IsNil(o.MontoDiario) {
		toSerialize["monto_diario"] = o.MontoDiario
	}
	if !IsNil(o.IngresoAcumulable) {
		toSerialize["ingreso_acumulable"] = o.IngresoAcumulable
	}
	if !IsNil(o.IngresoNoAcumulable) {
		toSerialize["ingreso_no_acumulable"] = o.IngresoNoAcumulable
	}
	return toSerialize, nil
}

type NullableNominaJubilacionProperties struct {
	value *NominaJubilacionProperties
	isSet bool
}

func (v NullableNominaJubilacionProperties) Get() *NominaJubilacionProperties {
	return v.value
}

func (v *NullableNominaJubilacionProperties) Set(val *NominaJubilacionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaJubilacionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaJubilacionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaJubilacionProperties(val *NominaJubilacionProperties) *NullableNominaJubilacionProperties {
	return &NullableNominaJubilacionProperties{value: val, isSet: true}
}

func (v NullableNominaJubilacionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaJubilacionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


