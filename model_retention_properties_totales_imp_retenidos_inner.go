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

// checks if the RetentionPropertiesTotalesImpRetenidosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionPropertiesTotalesImpRetenidosInner{}

// RetentionPropertiesTotalesImpRetenidosInner struct for RetentionPropertiesTotalesImpRetenidosInner
type RetentionPropertiesTotalesImpRetenidosInner struct {
	// Base del impuesto retenido.
	Base *float32 `json:"base,omitempty"`
	// Clave del tipo de impuesto retenido, del catálogo del SAT.
	Impuesto *string `json:"impuesto,omitempty"`
	// Importe del impuesto retenido
	Monto *float32 `json:"monto,omitempty"`
	// - `01`: Pago definitivo IVA - `02`: Pago definitivo IEPS - `03`: Pago definitivo ISR Plataformas - `04`: Pago provisional ISR 
	TipoPagoRet *string `json:"tipo_pago_ret,omitempty"`
}

// NewRetentionPropertiesTotalesImpRetenidosInner instantiates a new RetentionPropertiesTotalesImpRetenidosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionPropertiesTotalesImpRetenidosInner() *RetentionPropertiesTotalesImpRetenidosInner {
	this := RetentionPropertiesTotalesImpRetenidosInner{}
	return &this
}

// NewRetentionPropertiesTotalesImpRetenidosInnerWithDefaults instantiates a new RetentionPropertiesTotalesImpRetenidosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionPropertiesTotalesImpRetenidosInnerWithDefaults() *RetentionPropertiesTotalesImpRetenidosInner {
	this := RetentionPropertiesTotalesImpRetenidosInner{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetBase() float32 {
	if o == nil || IsNil(o.Base) {
		var ret float32
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given float32 and assigns it to the Base field.
func (o *RetentionPropertiesTotalesImpRetenidosInner) SetBase(v float32) {
	o.Base = &v
}

// GetImpuesto returns the Impuesto field value if set, zero value otherwise.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetImpuesto() string {
	if o == nil || IsNil(o.Impuesto) {
		var ret string
		return ret
	}
	return *o.Impuesto
}

// GetImpuestoOk returns a tuple with the Impuesto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetImpuestoOk() (*string, bool) {
	if o == nil || IsNil(o.Impuesto) {
		return nil, false
	}
	return o.Impuesto, true
}

// HasImpuesto returns a boolean if a field has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) HasImpuesto() bool {
	if o != nil && !IsNil(o.Impuesto) {
		return true
	}

	return false
}

// SetImpuesto gets a reference to the given string and assigns it to the Impuesto field.
func (o *RetentionPropertiesTotalesImpRetenidosInner) SetImpuesto(v string) {
	o.Impuesto = &v
}

// GetMonto returns the Monto field value if set, zero value otherwise.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetMonto() float32 {
	if o == nil || IsNil(o.Monto) {
		var ret float32
		return ret
	}
	return *o.Monto
}

// GetMontoOk returns a tuple with the Monto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetMontoOk() (*float32, bool) {
	if o == nil || IsNil(o.Monto) {
		return nil, false
	}
	return o.Monto, true
}

// HasMonto returns a boolean if a field has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) HasMonto() bool {
	if o != nil && !IsNil(o.Monto) {
		return true
	}

	return false
}

// SetMonto gets a reference to the given float32 and assigns it to the Monto field.
func (o *RetentionPropertiesTotalesImpRetenidosInner) SetMonto(v float32) {
	o.Monto = &v
}

// GetTipoPagoRet returns the TipoPagoRet field value if set, zero value otherwise.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetTipoPagoRet() string {
	if o == nil || IsNil(o.TipoPagoRet) {
		var ret string
		return ret
	}
	return *o.TipoPagoRet
}

// GetTipoPagoRetOk returns a tuple with the TipoPagoRet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) GetTipoPagoRetOk() (*string, bool) {
	if o == nil || IsNil(o.TipoPagoRet) {
		return nil, false
	}
	return o.TipoPagoRet, true
}

// HasTipoPagoRet returns a boolean if a field has been set.
func (o *RetentionPropertiesTotalesImpRetenidosInner) HasTipoPagoRet() bool {
	if o != nil && !IsNil(o.TipoPagoRet) {
		return true
	}

	return false
}

// SetTipoPagoRet gets a reference to the given string and assigns it to the TipoPagoRet field.
func (o *RetentionPropertiesTotalesImpRetenidosInner) SetTipoPagoRet(v string) {
	o.TipoPagoRet = &v
}

func (o RetentionPropertiesTotalesImpRetenidosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionPropertiesTotalesImpRetenidosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.Impuesto) {
		toSerialize["impuesto"] = o.Impuesto
	}
	if !IsNil(o.Monto) {
		toSerialize["monto"] = o.Monto
	}
	if !IsNil(o.TipoPagoRet) {
		toSerialize["tipo_pago_ret"] = o.TipoPagoRet
	}
	return toSerialize, nil
}

type NullableRetentionPropertiesTotalesImpRetenidosInner struct {
	value *RetentionPropertiesTotalesImpRetenidosInner
	isSet bool
}

func (v NullableRetentionPropertiesTotalesImpRetenidosInner) Get() *RetentionPropertiesTotalesImpRetenidosInner {
	return v.value
}

func (v *NullableRetentionPropertiesTotalesImpRetenidosInner) Set(val *RetentionPropertiesTotalesImpRetenidosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionPropertiesTotalesImpRetenidosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionPropertiesTotalesImpRetenidosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionPropertiesTotalesImpRetenidosInner(val *RetentionPropertiesTotalesImpRetenidosInner) *NullableRetentionPropertiesTotalesImpRetenidosInner {
	return &NullableRetentionPropertiesTotalesImpRetenidosInner{value: val, isSet: true}
}

func (v NullableRetentionPropertiesTotalesImpRetenidosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionPropertiesTotalesImpRetenidosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


