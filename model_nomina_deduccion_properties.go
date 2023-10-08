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

// checks if the NominaDeduccionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaDeduccionProperties{}

// NominaDeduccionProperties struct for NominaDeduccionProperties
type NominaDeduccionProperties struct {
	// Clave del catálogo [Tipo de deducción](#tipo-de-deducción).
	TipoDeduccion *string `json:"tipo_deduccion,omitempty"`
	// Concepto de la deducción. Si no se envía, se utilizará la descripción del catálogo del tipo de deducción.
	Concepto *string `json:"concepto,omitempty"`
	// Clave de control interno que asigna el patrón a cada deducción (descuento) de nómina propia de su contabilidad.
	Clave *string `json:"clave,omitempty"`
	// Importe del concepto de deducción.
	Importe *float32 `json:"importe,omitempty"`
}

// NewNominaDeduccionProperties instantiates a new NominaDeduccionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaDeduccionProperties() *NominaDeduccionProperties {
	this := NominaDeduccionProperties{}
	return &this
}

// NewNominaDeduccionPropertiesWithDefaults instantiates a new NominaDeduccionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaDeduccionPropertiesWithDefaults() *NominaDeduccionProperties {
	this := NominaDeduccionProperties{}
	return &this
}

// GetTipoDeduccion returns the TipoDeduccion field value if set, zero value otherwise.
func (o *NominaDeduccionProperties) GetTipoDeduccion() string {
	if o == nil || IsNil(o.TipoDeduccion) {
		var ret string
		return ret
	}
	return *o.TipoDeduccion
}

// GetTipoDeduccionOk returns a tuple with the TipoDeduccion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaDeduccionProperties) GetTipoDeduccionOk() (*string, bool) {
	if o == nil || IsNil(o.TipoDeduccion) {
		return nil, false
	}
	return o.TipoDeduccion, true
}

// HasTipoDeduccion returns a boolean if a field has been set.
func (o *NominaDeduccionProperties) HasTipoDeduccion() bool {
	if o != nil && !IsNil(o.TipoDeduccion) {
		return true
	}

	return false
}

// SetTipoDeduccion gets a reference to the given string and assigns it to the TipoDeduccion field.
func (o *NominaDeduccionProperties) SetTipoDeduccion(v string) {
	o.TipoDeduccion = &v
}

// GetConcepto returns the Concepto field value if set, zero value otherwise.
func (o *NominaDeduccionProperties) GetConcepto() string {
	if o == nil || IsNil(o.Concepto) {
		var ret string
		return ret
	}
	return *o.Concepto
}

// GetConceptoOk returns a tuple with the Concepto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaDeduccionProperties) GetConceptoOk() (*string, bool) {
	if o == nil || IsNil(o.Concepto) {
		return nil, false
	}
	return o.Concepto, true
}

// HasConcepto returns a boolean if a field has been set.
func (o *NominaDeduccionProperties) HasConcepto() bool {
	if o != nil && !IsNil(o.Concepto) {
		return true
	}

	return false
}

// SetConcepto gets a reference to the given string and assigns it to the Concepto field.
func (o *NominaDeduccionProperties) SetConcepto(v string) {
	o.Concepto = &v
}

// GetClave returns the Clave field value if set, zero value otherwise.
func (o *NominaDeduccionProperties) GetClave() string {
	if o == nil || IsNil(o.Clave) {
		var ret string
		return ret
	}
	return *o.Clave
}

// GetClaveOk returns a tuple with the Clave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaDeduccionProperties) GetClaveOk() (*string, bool) {
	if o == nil || IsNil(o.Clave) {
		return nil, false
	}
	return o.Clave, true
}

// HasClave returns a boolean if a field has been set.
func (o *NominaDeduccionProperties) HasClave() bool {
	if o != nil && !IsNil(o.Clave) {
		return true
	}

	return false
}

// SetClave gets a reference to the given string and assigns it to the Clave field.
func (o *NominaDeduccionProperties) SetClave(v string) {
	o.Clave = &v
}

// GetImporte returns the Importe field value if set, zero value otherwise.
func (o *NominaDeduccionProperties) GetImporte() float32 {
	if o == nil || IsNil(o.Importe) {
		var ret float32
		return ret
	}
	return *o.Importe
}

// GetImporteOk returns a tuple with the Importe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaDeduccionProperties) GetImporteOk() (*float32, bool) {
	if o == nil || IsNil(o.Importe) {
		return nil, false
	}
	return o.Importe, true
}

// HasImporte returns a boolean if a field has been set.
func (o *NominaDeduccionProperties) HasImporte() bool {
	if o != nil && !IsNil(o.Importe) {
		return true
	}

	return false
}

// SetImporte gets a reference to the given float32 and assigns it to the Importe field.
func (o *NominaDeduccionProperties) SetImporte(v float32) {
	o.Importe = &v
}

func (o NominaDeduccionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaDeduccionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TipoDeduccion) {
		toSerialize["tipo_deduccion"] = o.TipoDeduccion
	}
	if !IsNil(o.Concepto) {
		toSerialize["concepto"] = o.Concepto
	}
	if !IsNil(o.Clave) {
		toSerialize["clave"] = o.Clave
	}
	if !IsNil(o.Importe) {
		toSerialize["importe"] = o.Importe
	}
	return toSerialize, nil
}

type NullableNominaDeduccionProperties struct {
	value *NominaDeduccionProperties
	isSet bool
}

func (v NullableNominaDeduccionProperties) Get() *NominaDeduccionProperties {
	return v.value
}

func (v *NullableNominaDeduccionProperties) Set(val *NominaDeduccionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaDeduccionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaDeduccionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaDeduccionProperties(val *NominaDeduccionProperties) *NullableNominaDeduccionProperties {
	return &NullableNominaDeduccionProperties{value: val, isSet: true}
}

func (v NullableNominaDeduccionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaDeduccionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


