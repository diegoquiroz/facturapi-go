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

// checks if the NominaComplementDataNestedPropertiesOtrosPagosInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaComplementDataNestedPropertiesOtrosPagosInner{}

// NominaComplementDataNestedPropertiesOtrosPagosInner struct for NominaComplementDataNestedPropertiesOtrosPagosInner
type NominaComplementDataNestedPropertiesOtrosPagosInner struct {
	// Clave del catálogo [Tipo de Otro Pago](#tipo-de-otro-pago).
	TipoOtroPago *string `json:"tipo_otro_pago,omitempty"`
	// Clave de otro pago de nómina propia de la contabilidad de cada patrón.
	Clave *string `json:"clave,omitempty"`
	// Descripción alternativa correspondiente a la clave utilizada.
	Concepto *string `json:"concepto,omitempty"`
	// Importe por concepto de otro pago.
	Importe *float32 `json:"importe,omitempty"`
	// Subsudio causado conforme a la tabla del subsidio para el empleo publicada en el Anexo 8 de la Resolución Miscelánea Fiscal vigente.  Este valor será insertado dentro del nodo `SubsidioAlEmpleo`, y es requerido cuando el valor de `tipo_otro_pago` es `\"002\"`. 
	SubsidioCausado *float32 `json:"subsidio_causado,omitempty"`
	CompensacionSaldosAFavor *NominaCompensacionProperties `json:"compensacion_saldos_a_favor,omitempty"`
}

// NewNominaComplementDataNestedPropertiesOtrosPagosInner instantiates a new NominaComplementDataNestedPropertiesOtrosPagosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaComplementDataNestedPropertiesOtrosPagosInner() *NominaComplementDataNestedPropertiesOtrosPagosInner {
	this := NominaComplementDataNestedPropertiesOtrosPagosInner{}
	return &this
}

// NewNominaComplementDataNestedPropertiesOtrosPagosInnerWithDefaults instantiates a new NominaComplementDataNestedPropertiesOtrosPagosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaComplementDataNestedPropertiesOtrosPagosInnerWithDefaults() *NominaComplementDataNestedPropertiesOtrosPagosInner {
	this := NominaComplementDataNestedPropertiesOtrosPagosInner{}
	return &this
}

// GetTipoOtroPago returns the TipoOtroPago field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetTipoOtroPago() string {
	if o == nil || IsNil(o.TipoOtroPago) {
		var ret string
		return ret
	}
	return *o.TipoOtroPago
}

// GetTipoOtroPagoOk returns a tuple with the TipoOtroPago field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetTipoOtroPagoOk() (*string, bool) {
	if o == nil || IsNil(o.TipoOtroPago) {
		return nil, false
	}
	return o.TipoOtroPago, true
}

// HasTipoOtroPago returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasTipoOtroPago() bool {
	if o != nil && !IsNil(o.TipoOtroPago) {
		return true
	}

	return false
}

// SetTipoOtroPago gets a reference to the given string and assigns it to the TipoOtroPago field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetTipoOtroPago(v string) {
	o.TipoOtroPago = &v
}

// GetClave returns the Clave field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetClave() string {
	if o == nil || IsNil(o.Clave) {
		var ret string
		return ret
	}
	return *o.Clave
}

// GetClaveOk returns a tuple with the Clave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetClaveOk() (*string, bool) {
	if o == nil || IsNil(o.Clave) {
		return nil, false
	}
	return o.Clave, true
}

// HasClave returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasClave() bool {
	if o != nil && !IsNil(o.Clave) {
		return true
	}

	return false
}

// SetClave gets a reference to the given string and assigns it to the Clave field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetClave(v string) {
	o.Clave = &v
}

// GetConcepto returns the Concepto field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetConcepto() string {
	if o == nil || IsNil(o.Concepto) {
		var ret string
		return ret
	}
	return *o.Concepto
}

// GetConceptoOk returns a tuple with the Concepto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetConceptoOk() (*string, bool) {
	if o == nil || IsNil(o.Concepto) {
		return nil, false
	}
	return o.Concepto, true
}

// HasConcepto returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasConcepto() bool {
	if o != nil && !IsNil(o.Concepto) {
		return true
	}

	return false
}

// SetConcepto gets a reference to the given string and assigns it to the Concepto field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetConcepto(v string) {
	o.Concepto = &v
}

// GetImporte returns the Importe field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetImporte() float32 {
	if o == nil || IsNil(o.Importe) {
		var ret float32
		return ret
	}
	return *o.Importe
}

// GetImporteOk returns a tuple with the Importe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetImporteOk() (*float32, bool) {
	if o == nil || IsNil(o.Importe) {
		return nil, false
	}
	return o.Importe, true
}

// HasImporte returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasImporte() bool {
	if o != nil && !IsNil(o.Importe) {
		return true
	}

	return false
}

// SetImporte gets a reference to the given float32 and assigns it to the Importe field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetImporte(v float32) {
	o.Importe = &v
}

// GetSubsidioCausado returns the SubsidioCausado field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetSubsidioCausado() float32 {
	if o == nil || IsNil(o.SubsidioCausado) {
		var ret float32
		return ret
	}
	return *o.SubsidioCausado
}

// GetSubsidioCausadoOk returns a tuple with the SubsidioCausado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetSubsidioCausadoOk() (*float32, bool) {
	if o == nil || IsNil(o.SubsidioCausado) {
		return nil, false
	}
	return o.SubsidioCausado, true
}

// HasSubsidioCausado returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasSubsidioCausado() bool {
	if o != nil && !IsNil(o.SubsidioCausado) {
		return true
	}

	return false
}

// SetSubsidioCausado gets a reference to the given float32 and assigns it to the SubsidioCausado field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetSubsidioCausado(v float32) {
	o.SubsidioCausado = &v
}

// GetCompensacionSaldosAFavor returns the CompensacionSaldosAFavor field value if set, zero value otherwise.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetCompensacionSaldosAFavor() NominaCompensacionProperties {
	if o == nil || IsNil(o.CompensacionSaldosAFavor) {
		var ret NominaCompensacionProperties
		return ret
	}
	return *o.CompensacionSaldosAFavor
}

// GetCompensacionSaldosAFavorOk returns a tuple with the CompensacionSaldosAFavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetCompensacionSaldosAFavorOk() (*NominaCompensacionProperties, bool) {
	if o == nil || IsNil(o.CompensacionSaldosAFavor) {
		return nil, false
	}
	return o.CompensacionSaldosAFavor, true
}

// HasCompensacionSaldosAFavor returns a boolean if a field has been set.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasCompensacionSaldosAFavor() bool {
	if o != nil && !IsNil(o.CompensacionSaldosAFavor) {
		return true
	}

	return false
}

// SetCompensacionSaldosAFavor gets a reference to the given NominaCompensacionProperties and assigns it to the CompensacionSaldosAFavor field.
func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetCompensacionSaldosAFavor(v NominaCompensacionProperties) {
	o.CompensacionSaldosAFavor = &v
}

func (o NominaComplementDataNestedPropertiesOtrosPagosInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaComplementDataNestedPropertiesOtrosPagosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TipoOtroPago) {
		toSerialize["tipo_otro_pago"] = o.TipoOtroPago
	}
	if !IsNil(o.Clave) {
		toSerialize["clave"] = o.Clave
	}
	if !IsNil(o.Concepto) {
		toSerialize["concepto"] = o.Concepto
	}
	if !IsNil(o.Importe) {
		toSerialize["importe"] = o.Importe
	}
	if !IsNil(o.SubsidioCausado) {
		toSerialize["subsidio_causado"] = o.SubsidioCausado
	}
	if !IsNil(o.CompensacionSaldosAFavor) {
		toSerialize["compensacion_saldos_a_favor"] = o.CompensacionSaldosAFavor
	}
	return toSerialize, nil
}

type NullableNominaComplementDataNestedPropertiesOtrosPagosInner struct {
	value *NominaComplementDataNestedPropertiesOtrosPagosInner
	isSet bool
}

func (v NullableNominaComplementDataNestedPropertiesOtrosPagosInner) Get() *NominaComplementDataNestedPropertiesOtrosPagosInner {
	return v.value
}

func (v *NullableNominaComplementDataNestedPropertiesOtrosPagosInner) Set(val *NominaComplementDataNestedPropertiesOtrosPagosInner) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaComplementDataNestedPropertiesOtrosPagosInner) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaComplementDataNestedPropertiesOtrosPagosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaComplementDataNestedPropertiesOtrosPagosInner(val *NominaComplementDataNestedPropertiesOtrosPagosInner) *NullableNominaComplementDataNestedPropertiesOtrosPagosInner {
	return &NullableNominaComplementDataNestedPropertiesOtrosPagosInner{value: val, isSet: true}
}

func (v NullableNominaComplementDataNestedPropertiesOtrosPagosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaComplementDataNestedPropertiesOtrosPagosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

