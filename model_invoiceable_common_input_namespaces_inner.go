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

// checks if the InvoiceableCommonInputNamespacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceableCommonInputNamespacesInner{}

// InvoiceableCommonInputNamespacesInner struct for InvoiceableCommonInputNamespacesInner
type InvoiceableCommonInputNamespacesInner struct {
	// Prefijo o nombre del namespace.
	Prefix *string `json:"prefix,omitempty"`
	// Dirección URL asociada al namespace.
	Uri *string `json:"uri,omitempty"`
	// Dirección URL del esquema de validación XSD.
	SchemaLocation *string `json:"schema_location,omitempty"`
}

// NewInvoiceableCommonInputNamespacesInner instantiates a new InvoiceableCommonInputNamespacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceableCommonInputNamespacesInner() *InvoiceableCommonInputNamespacesInner {
	this := InvoiceableCommonInputNamespacesInner{}
	return &this
}

// NewInvoiceableCommonInputNamespacesInnerWithDefaults instantiates a new InvoiceableCommonInputNamespacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceableCommonInputNamespacesInnerWithDefaults() *InvoiceableCommonInputNamespacesInner {
	this := InvoiceableCommonInputNamespacesInner{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InvoiceableCommonInputNamespacesInner) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInputNamespacesInner) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InvoiceableCommonInputNamespacesInner) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InvoiceableCommonInputNamespacesInner) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InvoiceableCommonInputNamespacesInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInputNamespacesInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InvoiceableCommonInputNamespacesInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InvoiceableCommonInputNamespacesInner) SetUri(v string) {
	o.Uri = &v
}

// GetSchemaLocation returns the SchemaLocation field value if set, zero value otherwise.
func (o *InvoiceableCommonInputNamespacesInner) GetSchemaLocation() string {
	if o == nil || IsNil(o.SchemaLocation) {
		var ret string
		return ret
	}
	return *o.SchemaLocation
}

// GetSchemaLocationOk returns a tuple with the SchemaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInputNamespacesInner) GetSchemaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaLocation) {
		return nil, false
	}
	return o.SchemaLocation, true
}

// HasSchemaLocation returns a boolean if a field has been set.
func (o *InvoiceableCommonInputNamespacesInner) HasSchemaLocation() bool {
	if o != nil && !IsNil(o.SchemaLocation) {
		return true
	}

	return false
}

// SetSchemaLocation gets a reference to the given string and assigns it to the SchemaLocation field.
func (o *InvoiceableCommonInputNamespacesInner) SetSchemaLocation(v string) {
	o.SchemaLocation = &v
}

func (o InvoiceableCommonInputNamespacesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceableCommonInputNamespacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.SchemaLocation) {
		toSerialize["schema_location"] = o.SchemaLocation
	}
	return toSerialize, nil
}

type NullableInvoiceableCommonInputNamespacesInner struct {
	value *InvoiceableCommonInputNamespacesInner
	isSet bool
}

func (v NullableInvoiceableCommonInputNamespacesInner) Get() *InvoiceableCommonInputNamespacesInner {
	return v.value
}

func (v *NullableInvoiceableCommonInputNamespacesInner) Set(val *InvoiceableCommonInputNamespacesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceableCommonInputNamespacesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceableCommonInputNamespacesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceableCommonInputNamespacesInner(val *InvoiceableCommonInputNamespacesInner) *NullableInvoiceableCommonInputNamespacesInner {
	return &NullableInvoiceableCommonInputNamespacesInner{value: val, isSet: true}
}

func (v NullableInvoiceableCommonInputNamespacesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceableCommonInputNamespacesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


