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

// checks if the NominaPercepcionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NominaPercepcionProperties{}

// NominaPercepcionProperties struct for NominaPercepcionProperties
type NominaPercepcionProperties struct {
	// Clave del catálogo [Tipo de percepción](#tipo-de-percepcion).
	TipoPercepcion *string `json:"tipo_percepcion,omitempty"`
	// Concepto de la percepción. Si no se envía, se utilizará la descripción del catálogo del tipo de percepción.
	Concepto *string `json:"concepto,omitempty"`
	// Clave de control interno que asigna el patrón a cada percepción de nómina propia de su contabilidad.
	Clave *string `json:"clave,omitempty"`
	// Importe gravado por el concepto indicado en el tipo de percepción.
	ImporteGravado *float32 `json:"importe_gravado,omitempty"`
	// Importe exento por el concepto indicado en el tipo de percepción.
	ImporteExento *float32 `json:"importe_exento,omitempty"`
	AccionesOTitulos *NominaAccionesProperties `json:"acciones_o_titulos,omitempty"`
	HorasExtra []NominaHorasExtraProperties `json:"horas_extra,omitempty"`
}

// NewNominaPercepcionProperties instantiates a new NominaPercepcionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNominaPercepcionProperties() *NominaPercepcionProperties {
	this := NominaPercepcionProperties{}
	return &this
}

// NewNominaPercepcionPropertiesWithDefaults instantiates a new NominaPercepcionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNominaPercepcionPropertiesWithDefaults() *NominaPercepcionProperties {
	this := NominaPercepcionProperties{}
	return &this
}

// GetTipoPercepcion returns the TipoPercepcion field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetTipoPercepcion() string {
	if o == nil || IsNil(o.TipoPercepcion) {
		var ret string
		return ret
	}
	return *o.TipoPercepcion
}

// GetTipoPercepcionOk returns a tuple with the TipoPercepcion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetTipoPercepcionOk() (*string, bool) {
	if o == nil || IsNil(o.TipoPercepcion) {
		return nil, false
	}
	return o.TipoPercepcion, true
}

// HasTipoPercepcion returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasTipoPercepcion() bool {
	if o != nil && !IsNil(o.TipoPercepcion) {
		return true
	}

	return false
}

// SetTipoPercepcion gets a reference to the given string and assigns it to the TipoPercepcion field.
func (o *NominaPercepcionProperties) SetTipoPercepcion(v string) {
	o.TipoPercepcion = &v
}

// GetConcepto returns the Concepto field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetConcepto() string {
	if o == nil || IsNil(o.Concepto) {
		var ret string
		return ret
	}
	return *o.Concepto
}

// GetConceptoOk returns a tuple with the Concepto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetConceptoOk() (*string, bool) {
	if o == nil || IsNil(o.Concepto) {
		return nil, false
	}
	return o.Concepto, true
}

// HasConcepto returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasConcepto() bool {
	if o != nil && !IsNil(o.Concepto) {
		return true
	}

	return false
}

// SetConcepto gets a reference to the given string and assigns it to the Concepto field.
func (o *NominaPercepcionProperties) SetConcepto(v string) {
	o.Concepto = &v
}

// GetClave returns the Clave field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetClave() string {
	if o == nil || IsNil(o.Clave) {
		var ret string
		return ret
	}
	return *o.Clave
}

// GetClaveOk returns a tuple with the Clave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetClaveOk() (*string, bool) {
	if o == nil || IsNil(o.Clave) {
		return nil, false
	}
	return o.Clave, true
}

// HasClave returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasClave() bool {
	if o != nil && !IsNil(o.Clave) {
		return true
	}

	return false
}

// SetClave gets a reference to the given string and assigns it to the Clave field.
func (o *NominaPercepcionProperties) SetClave(v string) {
	o.Clave = &v
}

// GetImporteGravado returns the ImporteGravado field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetImporteGravado() float32 {
	if o == nil || IsNil(o.ImporteGravado) {
		var ret float32
		return ret
	}
	return *o.ImporteGravado
}

// GetImporteGravadoOk returns a tuple with the ImporteGravado field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetImporteGravadoOk() (*float32, bool) {
	if o == nil || IsNil(o.ImporteGravado) {
		return nil, false
	}
	return o.ImporteGravado, true
}

// HasImporteGravado returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasImporteGravado() bool {
	if o != nil && !IsNil(o.ImporteGravado) {
		return true
	}

	return false
}

// SetImporteGravado gets a reference to the given float32 and assigns it to the ImporteGravado field.
func (o *NominaPercepcionProperties) SetImporteGravado(v float32) {
	o.ImporteGravado = &v
}

// GetImporteExento returns the ImporteExento field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetImporteExento() float32 {
	if o == nil || IsNil(o.ImporteExento) {
		var ret float32
		return ret
	}
	return *o.ImporteExento
}

// GetImporteExentoOk returns a tuple with the ImporteExento field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetImporteExentoOk() (*float32, bool) {
	if o == nil || IsNil(o.ImporteExento) {
		return nil, false
	}
	return o.ImporteExento, true
}

// HasImporteExento returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasImporteExento() bool {
	if o != nil && !IsNil(o.ImporteExento) {
		return true
	}

	return false
}

// SetImporteExento gets a reference to the given float32 and assigns it to the ImporteExento field.
func (o *NominaPercepcionProperties) SetImporteExento(v float32) {
	o.ImporteExento = &v
}

// GetAccionesOTitulos returns the AccionesOTitulos field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetAccionesOTitulos() NominaAccionesProperties {
	if o == nil || IsNil(o.AccionesOTitulos) {
		var ret NominaAccionesProperties
		return ret
	}
	return *o.AccionesOTitulos
}

// GetAccionesOTitulosOk returns a tuple with the AccionesOTitulos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetAccionesOTitulosOk() (*NominaAccionesProperties, bool) {
	if o == nil || IsNil(o.AccionesOTitulos) {
		return nil, false
	}
	return o.AccionesOTitulos, true
}

// HasAccionesOTitulos returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasAccionesOTitulos() bool {
	if o != nil && !IsNil(o.AccionesOTitulos) {
		return true
	}

	return false
}

// SetAccionesOTitulos gets a reference to the given NominaAccionesProperties and assigns it to the AccionesOTitulos field.
func (o *NominaPercepcionProperties) SetAccionesOTitulos(v NominaAccionesProperties) {
	o.AccionesOTitulos = &v
}

// GetHorasExtra returns the HorasExtra field value if set, zero value otherwise.
func (o *NominaPercepcionProperties) GetHorasExtra() []NominaHorasExtraProperties {
	if o == nil || IsNil(o.HorasExtra) {
		var ret []NominaHorasExtraProperties
		return ret
	}
	return o.HorasExtra
}

// GetHorasExtraOk returns a tuple with the HorasExtra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NominaPercepcionProperties) GetHorasExtraOk() ([]NominaHorasExtraProperties, bool) {
	if o == nil || IsNil(o.HorasExtra) {
		return nil, false
	}
	return o.HorasExtra, true
}

// HasHorasExtra returns a boolean if a field has been set.
func (o *NominaPercepcionProperties) HasHorasExtra() bool {
	if o != nil && !IsNil(o.HorasExtra) {
		return true
	}

	return false
}

// SetHorasExtra gets a reference to the given []NominaHorasExtraProperties and assigns it to the HorasExtra field.
func (o *NominaPercepcionProperties) SetHorasExtra(v []NominaHorasExtraProperties) {
	o.HorasExtra = v
}

func (o NominaPercepcionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NominaPercepcionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TipoPercepcion) {
		toSerialize["tipo_percepcion"] = o.TipoPercepcion
	}
	if !IsNil(o.Concepto) {
		toSerialize["concepto"] = o.Concepto
	}
	if !IsNil(o.Clave) {
		toSerialize["clave"] = o.Clave
	}
	if !IsNil(o.ImporteGravado) {
		toSerialize["importe_gravado"] = o.ImporteGravado
	}
	if !IsNil(o.ImporteExento) {
		toSerialize["importe_exento"] = o.ImporteExento
	}
	if !IsNil(o.AccionesOTitulos) {
		toSerialize["acciones_o_titulos"] = o.AccionesOTitulos
	}
	if !IsNil(o.HorasExtra) {
		toSerialize["horas_extra"] = o.HorasExtra
	}
	return toSerialize, nil
}

type NullableNominaPercepcionProperties struct {
	value *NominaPercepcionProperties
	isSet bool
}

func (v NullableNominaPercepcionProperties) Get() *NominaPercepcionProperties {
	return v.value
}

func (v *NullableNominaPercepcionProperties) Set(val *NominaPercepcionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNominaPercepcionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNominaPercepcionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNominaPercepcionProperties(val *NominaPercepcionProperties) *NullableNominaPercepcionProperties {
	return &NullableNominaPercepcionProperties{value: val, isSet: true}
}

func (v NullableNominaPercepcionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNominaPercepcionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


