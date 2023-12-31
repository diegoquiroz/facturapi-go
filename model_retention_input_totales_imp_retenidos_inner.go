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

// checks if the RetentionInputTotalesImpRetenidosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionInputTotalesImpRetenidosInner{}

// RetentionInputTotalesImpRetenidosInner struct for RetentionInputTotalesImpRetenidosInner
type RetentionInputTotalesImpRetenidosInner struct {
	// Base del impuesto retenido.
	BaseRet *float32 `json:"base_ret,omitempty"`
	// Clave del tipo de impuesto retenido, del catálogo del SAT.
	Impuesto *string `json:"impuesto,omitempty"`
	// Importe del impuesto retenido
	MontoRet float32 `json:"monto_ret"`
	// - `01`: Pago definitivo IVA - `02`: Pago definitivo IEPS - `03`: Pago definitivo ISR Plataformas - `04`: Pago provisional ISR 
	TipoPagoRet string `json:"tipo_pago_ret"`
}

// NewRetentionInputTotalesImpRetenidosInner instantiates a new RetentionInputTotalesImpRetenidosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionInputTotalesImpRetenidosInner(montoRet float32, tipoPagoRet string) *RetentionInputTotalesImpRetenidosInner {
	this := RetentionInputTotalesImpRetenidosInner{}
	this.MontoRet = montoRet
	this.TipoPagoRet = tipoPagoRet
	return &this
}

// NewRetentionInputTotalesImpRetenidosInnerWithDefaults instantiates a new RetentionInputTotalesImpRetenidosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionInputTotalesImpRetenidosInnerWithDefaults() *RetentionInputTotalesImpRetenidosInner {
	this := RetentionInputTotalesImpRetenidosInner{}
	return &this
}

// GetBaseRet returns the BaseRet field value if set, zero value otherwise.
func (o *RetentionInputTotalesImpRetenidosInner) GetBaseRet() float32 {
	if o == nil || IsNil(o.BaseRet) {
		var ret float32
		return ret
	}
	return *o.BaseRet
}

// GetBaseRetOk returns a tuple with the BaseRet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionInputTotalesImpRetenidosInner) GetBaseRetOk() (*float32, bool) {
	if o == nil || IsNil(o.BaseRet) {
		return nil, false
	}
	return o.BaseRet, true
}

// HasBaseRet returns a boolean if a field has been set.
func (o *RetentionInputTotalesImpRetenidosInner) HasBaseRet() bool {
	if o != nil && !IsNil(o.BaseRet) {
		return true
	}

	return false
}

// SetBaseRet gets a reference to the given float32 and assigns it to the BaseRet field.
func (o *RetentionInputTotalesImpRetenidosInner) SetBaseRet(v float32) {
	o.BaseRet = &v
}

// GetImpuesto returns the Impuesto field value if set, zero value otherwise.
func (o *RetentionInputTotalesImpRetenidosInner) GetImpuesto() string {
	if o == nil || IsNil(o.Impuesto) {
		var ret string
		return ret
	}
	return *o.Impuesto
}

// GetImpuestoOk returns a tuple with the Impuesto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionInputTotalesImpRetenidosInner) GetImpuestoOk() (*string, bool) {
	if o == nil || IsNil(o.Impuesto) {
		return nil, false
	}
	return o.Impuesto, true
}

// HasImpuesto returns a boolean if a field has been set.
func (o *RetentionInputTotalesImpRetenidosInner) HasImpuesto() bool {
	if o != nil && !IsNil(o.Impuesto) {
		return true
	}

	return false
}

// SetImpuesto gets a reference to the given string and assigns it to the Impuesto field.
func (o *RetentionInputTotalesImpRetenidosInner) SetImpuesto(v string) {
	o.Impuesto = &v
}

// GetMontoRet returns the MontoRet field value
func (o *RetentionInputTotalesImpRetenidosInner) GetMontoRet() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MontoRet
}

// GetMontoRetOk returns a tuple with the MontoRet field value
// and a boolean to check if the value has been set.
func (o *RetentionInputTotalesImpRetenidosInner) GetMontoRetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MontoRet, true
}

// SetMontoRet sets field value
func (o *RetentionInputTotalesImpRetenidosInner) SetMontoRet(v float32) {
	o.MontoRet = v
}

// GetTipoPagoRet returns the TipoPagoRet field value
func (o *RetentionInputTotalesImpRetenidosInner) GetTipoPagoRet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TipoPagoRet
}

// GetTipoPagoRetOk returns a tuple with the TipoPagoRet field value
// and a boolean to check if the value has been set.
func (o *RetentionInputTotalesImpRetenidosInner) GetTipoPagoRetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TipoPagoRet, true
}

// SetTipoPagoRet sets field value
func (o *RetentionInputTotalesImpRetenidosInner) SetTipoPagoRet(v string) {
	o.TipoPagoRet = v
}

func (o RetentionInputTotalesImpRetenidosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionInputTotalesImpRetenidosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseRet) {
		toSerialize["base_ret"] = o.BaseRet
	}
	if !IsNil(o.Impuesto) {
		toSerialize["impuesto"] = o.Impuesto
	}
	toSerialize["monto_ret"] = o.MontoRet
	toSerialize["tipo_pago_ret"] = o.TipoPagoRet
	return toSerialize, nil
}

type NullableRetentionInputTotalesImpRetenidosInner struct {
	value *RetentionInputTotalesImpRetenidosInner
	isSet bool
}

func (v NullableRetentionInputTotalesImpRetenidosInner) Get() *RetentionInputTotalesImpRetenidosInner {
	return v.value
}

func (v *NullableRetentionInputTotalesImpRetenidosInner) Set(val *RetentionInputTotalesImpRetenidosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionInputTotalesImpRetenidosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionInputTotalesImpRetenidosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionInputTotalesImpRetenidosInner(val *RetentionInputTotalesImpRetenidosInner) *NullableRetentionInputTotalesImpRetenidosInner {
	return &NullableRetentionInputTotalesImpRetenidosInner{value: val, isSet: true}
}

func (v NullableRetentionInputTotalesImpRetenidosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionInputTotalesImpRetenidosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


