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

// checks if the InvoiceableCommonInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceableCommonInput{}

// InvoiceableCommonInput struct for InvoiceableCommonInput
type InvoiceableCommonInput struct {
	// Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización.
	FolioNumber *int32 `json:"folio_number,omitempty"`
	// Serie. Caracteres designados por la empresa para control interno y sin validez fiscal.
	Series *string `json:"series,omitempty"`
	// En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: `h1`, `h2`, `h3`, `h4`, `h5`, `h6`, `div`, `p`, `span`, `small`, `br`, `b`, `i`, `ul`, `ol`, `li`, `strong`, `table`, `thead`, `tbody`, `tfoot`, `tr`, `th` y `td`. No se permiten atributos ni estilos. 
	PdfCustomSection *string `json:"pdf_custom_section,omitempty"`
	// Código XML con la Addenda que se necesite agregar a la factura.
	Addenda *string `json:"addenda,omitempty"`
	Namespaces []InvoiceableCommonInputNamespacesInner `json:"namespaces,omitempty"`
}

// NewInvoiceableCommonInput instantiates a new InvoiceableCommonInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceableCommonInput() *InvoiceableCommonInput {
	this := InvoiceableCommonInput{}
	var folioNumber int32 = autoincremental
	this.FolioNumber = &folioNumber
	return &this
}

// NewInvoiceableCommonInputWithDefaults instantiates a new InvoiceableCommonInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceableCommonInputWithDefaults() *InvoiceableCommonInput {
	this := InvoiceableCommonInput{}
	var folioNumber int32 = autoincremental
	this.FolioNumber = &folioNumber
	return &this
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *InvoiceableCommonInput) GetFolioNumber() int32 {
	if o == nil || IsNil(o.FolioNumber) {
		var ret int32
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInput) GetFolioNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioNumber) {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *InvoiceableCommonInput) HasFolioNumber() bool {
	if o != nil && !IsNil(o.FolioNumber) {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given int32 and assigns it to the FolioNumber field.
func (o *InvoiceableCommonInput) SetFolioNumber(v int32) {
	o.FolioNumber = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *InvoiceableCommonInput) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInput) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *InvoiceableCommonInput) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *InvoiceableCommonInput) SetSeries(v string) {
	o.Series = &v
}

// GetPdfCustomSection returns the PdfCustomSection field value if set, zero value otherwise.
func (o *InvoiceableCommonInput) GetPdfCustomSection() string {
	if o == nil || IsNil(o.PdfCustomSection) {
		var ret string
		return ret
	}
	return *o.PdfCustomSection
}

// GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInput) GetPdfCustomSectionOk() (*string, bool) {
	if o == nil || IsNil(o.PdfCustomSection) {
		return nil, false
	}
	return o.PdfCustomSection, true
}

// HasPdfCustomSection returns a boolean if a field has been set.
func (o *InvoiceableCommonInput) HasPdfCustomSection() bool {
	if o != nil && !IsNil(o.PdfCustomSection) {
		return true
	}

	return false
}

// SetPdfCustomSection gets a reference to the given string and assigns it to the PdfCustomSection field.
func (o *InvoiceableCommonInput) SetPdfCustomSection(v string) {
	o.PdfCustomSection = &v
}

// GetAddenda returns the Addenda field value if set, zero value otherwise.
func (o *InvoiceableCommonInput) GetAddenda() string {
	if o == nil || IsNil(o.Addenda) {
		var ret string
		return ret
	}
	return *o.Addenda
}

// GetAddendaOk returns a tuple with the Addenda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInput) GetAddendaOk() (*string, bool) {
	if o == nil || IsNil(o.Addenda) {
		return nil, false
	}
	return o.Addenda, true
}

// HasAddenda returns a boolean if a field has been set.
func (o *InvoiceableCommonInput) HasAddenda() bool {
	if o != nil && !IsNil(o.Addenda) {
		return true
	}

	return false
}

// SetAddenda gets a reference to the given string and assigns it to the Addenda field.
func (o *InvoiceableCommonInput) SetAddenda(v string) {
	o.Addenda = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *InvoiceableCommonInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner {
	if o == nil || IsNil(o.Namespaces) {
		var ret []InvoiceableCommonInputNamespacesInner
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceableCommonInput) GetNamespacesOk() ([]InvoiceableCommonInputNamespacesInner, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *InvoiceableCommonInput) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []InvoiceableCommonInputNamespacesInner and assigns it to the Namespaces field.
func (o *InvoiceableCommonInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner) {
	o.Namespaces = v
}

func (o InvoiceableCommonInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceableCommonInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FolioNumber) {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.PdfCustomSection) {
		toSerialize["pdf_custom_section"] = o.PdfCustomSection
	}
	if !IsNil(o.Addenda) {
		toSerialize["addenda"] = o.Addenda
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	return toSerialize, nil
}

type NullableInvoiceableCommonInput struct {
	value *InvoiceableCommonInput
	isSet bool
}

func (v NullableInvoiceableCommonInput) Get() *InvoiceableCommonInput {
	return v.value
}

func (v *NullableInvoiceableCommonInput) Set(val *InvoiceableCommonInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceableCommonInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceableCommonInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceableCommonInput(val *InvoiceableCommonInput) *NullableInvoiceableCommonInput {
	return &NullableInvoiceableCommonInput{value: val, isSet: true}
}

func (v NullableInvoiceableCommonInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceableCommonInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


