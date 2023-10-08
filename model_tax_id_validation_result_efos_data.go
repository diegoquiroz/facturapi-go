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

// checks if the TaxIdValidationResultEfosData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxIdValidationResultEfosData{}

// TaxIdValidationResultEfosData Objeto con el resultado de la búqueda ante el SAT. Toda la información contenida en este objeto proviene del SAT. 
type TaxIdValidationResultEfosData struct {
	// Disponible sólo cuando el RFC no fue encontrado en la lista, lo cual es bueno. 
	Mensaje *string `json:"mensaje,omitempty"`
	// Texto que indica la fecha de actualización de la lista.
	FechaLista *string `json:"fechaLista,omitempty"`
	Detalles []TaxIdValidationResultEfosDataDetallesInner `json:"detalles,omitempty"`
}

// NewTaxIdValidationResultEfosData instantiates a new TaxIdValidationResultEfosData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxIdValidationResultEfosData() *TaxIdValidationResultEfosData {
	this := TaxIdValidationResultEfosData{}
	return &this
}

// NewTaxIdValidationResultEfosDataWithDefaults instantiates a new TaxIdValidationResultEfosData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxIdValidationResultEfosDataWithDefaults() *TaxIdValidationResultEfosData {
	this := TaxIdValidationResultEfosData{}
	return &this
}

// GetMensaje returns the Mensaje field value if set, zero value otherwise.
func (o *TaxIdValidationResultEfosData) GetMensaje() string {
	if o == nil || IsNil(o.Mensaje) {
		var ret string
		return ret
	}
	return *o.Mensaje
}

// GetMensajeOk returns a tuple with the Mensaje field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxIdValidationResultEfosData) GetMensajeOk() (*string, bool) {
	if o == nil || IsNil(o.Mensaje) {
		return nil, false
	}
	return o.Mensaje, true
}

// HasMensaje returns a boolean if a field has been set.
func (o *TaxIdValidationResultEfosData) HasMensaje() bool {
	if o != nil && !IsNil(o.Mensaje) {
		return true
	}

	return false
}

// SetMensaje gets a reference to the given string and assigns it to the Mensaje field.
func (o *TaxIdValidationResultEfosData) SetMensaje(v string) {
	o.Mensaje = &v
}

// GetFechaLista returns the FechaLista field value if set, zero value otherwise.
func (o *TaxIdValidationResultEfosData) GetFechaLista() string {
	if o == nil || IsNil(o.FechaLista) {
		var ret string
		return ret
	}
	return *o.FechaLista
}

// GetFechaListaOk returns a tuple with the FechaLista field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxIdValidationResultEfosData) GetFechaListaOk() (*string, bool) {
	if o == nil || IsNil(o.FechaLista) {
		return nil, false
	}
	return o.FechaLista, true
}

// HasFechaLista returns a boolean if a field has been set.
func (o *TaxIdValidationResultEfosData) HasFechaLista() bool {
	if o != nil && !IsNil(o.FechaLista) {
		return true
	}

	return false
}

// SetFechaLista gets a reference to the given string and assigns it to the FechaLista field.
func (o *TaxIdValidationResultEfosData) SetFechaLista(v string) {
	o.FechaLista = &v
}

// GetDetalles returns the Detalles field value if set, zero value otherwise.
func (o *TaxIdValidationResultEfosData) GetDetalles() []TaxIdValidationResultEfosDataDetallesInner {
	if o == nil || IsNil(o.Detalles) {
		var ret []TaxIdValidationResultEfosDataDetallesInner
		return ret
	}
	return o.Detalles
}

// GetDetallesOk returns a tuple with the Detalles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxIdValidationResultEfosData) GetDetallesOk() ([]TaxIdValidationResultEfosDataDetallesInner, bool) {
	if o == nil || IsNil(o.Detalles) {
		return nil, false
	}
	return o.Detalles, true
}

// HasDetalles returns a boolean if a field has been set.
func (o *TaxIdValidationResultEfosData) HasDetalles() bool {
	if o != nil && !IsNil(o.Detalles) {
		return true
	}

	return false
}

// SetDetalles gets a reference to the given []TaxIdValidationResultEfosDataDetallesInner and assigns it to the Detalles field.
func (o *TaxIdValidationResultEfosData) SetDetalles(v []TaxIdValidationResultEfosDataDetallesInner) {
	o.Detalles = v
}

func (o TaxIdValidationResultEfosData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxIdValidationResultEfosData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mensaje) {
		toSerialize["mensaje"] = o.Mensaje
	}
	if !IsNil(o.FechaLista) {
		toSerialize["fechaLista"] = o.FechaLista
	}
	if !IsNil(o.Detalles) {
		toSerialize["detalles"] = o.Detalles
	}
	return toSerialize, nil
}

type NullableTaxIdValidationResultEfosData struct {
	value *TaxIdValidationResultEfosData
	isSet bool
}

func (v NullableTaxIdValidationResultEfosData) Get() *TaxIdValidationResultEfosData {
	return v.value
}

func (v *NullableTaxIdValidationResultEfosData) Set(val *TaxIdValidationResultEfosData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxIdValidationResultEfosData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxIdValidationResultEfosData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxIdValidationResultEfosData(val *TaxIdValidationResultEfosData) *NullableTaxIdValidationResultEfosData {
	return &NullableTaxIdValidationResultEfosData{value: val, isSet: true}
}

func (v NullableTaxIdValidationResultEfosData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxIdValidationResultEfosData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


