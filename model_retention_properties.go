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
	"time"
)

// checks if the RetentionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionProperties{}

// RetentionProperties struct for RetentionProperties
type RetentionProperties struct {
	// Clave de la retención o información de pagos de acuerdo al catálogo del SAT.
	CveRetenc *string `json:"cve_retenc,omitempty"`
	// Fecha de expedición del comprobante en formato ISO8601 (UTC String).
	FechaExp *time.Time `json:"fecha_exp,omitempty"`
	// Si la clave de la retención es “25” (Otro tipo de retenciones), este campo se usa para registrar la descripción de la retención.
	DescRetenc *string `json:"desc_retenc,omitempty"`
	// Identificador alfanumérico para control interno de la empresa y sin relevancia fiscal.
	FolioInt *string `json:"folio_int,omitempty"`
	Periodo *RetentionPropertiesPeriodo `json:"periodo,omitempty"`
	Totales *RetentionPropertiesTotales `json:"totales,omitempty"`
	// Identificador opcional que puedes usar para relacionar esta retención con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único.
	ExternalId *string `json:"external_id,omitempty"`
	// Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta. 
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	Complements []string `json:"complements,omitempty"`
	// En caso de que necesites incluir más información en el PDF, este campo te permite insertar código HTML con tu propio contenido.
	PdfCustomSection *string `json:"pdf_custom_section,omitempty"`
	// Código XML con la Addenda que se necesite agregar a la factura.
	Addenda *string `json:"addenda,omitempty"`
	Namespaces []NamespaceProperties `json:"namespaces,omitempty"`
}

// NewRetentionProperties instantiates a new RetentionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionProperties() *RetentionProperties {
	this := RetentionProperties{}
	return &this
}

// NewRetentionPropertiesWithDefaults instantiates a new RetentionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionPropertiesWithDefaults() *RetentionProperties {
	this := RetentionProperties{}
	return &this
}

// GetCveRetenc returns the CveRetenc field value if set, zero value otherwise.
func (o *RetentionProperties) GetCveRetenc() string {
	if o == nil || IsNil(o.CveRetenc) {
		var ret string
		return ret
	}
	return *o.CveRetenc
}

// GetCveRetencOk returns a tuple with the CveRetenc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetCveRetencOk() (*string, bool) {
	if o == nil || IsNil(o.CveRetenc) {
		return nil, false
	}
	return o.CveRetenc, true
}

// HasCveRetenc returns a boolean if a field has been set.
func (o *RetentionProperties) HasCveRetenc() bool {
	if o != nil && !IsNil(o.CveRetenc) {
		return true
	}

	return false
}

// SetCveRetenc gets a reference to the given string and assigns it to the CveRetenc field.
func (o *RetentionProperties) SetCveRetenc(v string) {
	o.CveRetenc = &v
}

// GetFechaExp returns the FechaExp field value if set, zero value otherwise.
func (o *RetentionProperties) GetFechaExp() time.Time {
	if o == nil || IsNil(o.FechaExp) {
		var ret time.Time
		return ret
	}
	return *o.FechaExp
}

// GetFechaExpOk returns a tuple with the FechaExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetFechaExpOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FechaExp) {
		return nil, false
	}
	return o.FechaExp, true
}

// HasFechaExp returns a boolean if a field has been set.
func (o *RetentionProperties) HasFechaExp() bool {
	if o != nil && !IsNil(o.FechaExp) {
		return true
	}

	return false
}

// SetFechaExp gets a reference to the given time.Time and assigns it to the FechaExp field.
func (o *RetentionProperties) SetFechaExp(v time.Time) {
	o.FechaExp = &v
}

// GetDescRetenc returns the DescRetenc field value if set, zero value otherwise.
func (o *RetentionProperties) GetDescRetenc() string {
	if o == nil || IsNil(o.DescRetenc) {
		var ret string
		return ret
	}
	return *o.DescRetenc
}

// GetDescRetencOk returns a tuple with the DescRetenc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetDescRetencOk() (*string, bool) {
	if o == nil || IsNil(o.DescRetenc) {
		return nil, false
	}
	return o.DescRetenc, true
}

// HasDescRetenc returns a boolean if a field has been set.
func (o *RetentionProperties) HasDescRetenc() bool {
	if o != nil && !IsNil(o.DescRetenc) {
		return true
	}

	return false
}

// SetDescRetenc gets a reference to the given string and assigns it to the DescRetenc field.
func (o *RetentionProperties) SetDescRetenc(v string) {
	o.DescRetenc = &v
}

// GetFolioInt returns the FolioInt field value if set, zero value otherwise.
func (o *RetentionProperties) GetFolioInt() string {
	if o == nil || IsNil(o.FolioInt) {
		var ret string
		return ret
	}
	return *o.FolioInt
}

// GetFolioIntOk returns a tuple with the FolioInt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetFolioIntOk() (*string, bool) {
	if o == nil || IsNil(o.FolioInt) {
		return nil, false
	}
	return o.FolioInt, true
}

// HasFolioInt returns a boolean if a field has been set.
func (o *RetentionProperties) HasFolioInt() bool {
	if o != nil && !IsNil(o.FolioInt) {
		return true
	}

	return false
}

// SetFolioInt gets a reference to the given string and assigns it to the FolioInt field.
func (o *RetentionProperties) SetFolioInt(v string) {
	o.FolioInt = &v
}

// GetPeriodo returns the Periodo field value if set, zero value otherwise.
func (o *RetentionProperties) GetPeriodo() RetentionPropertiesPeriodo {
	if o == nil || IsNil(o.Periodo) {
		var ret RetentionPropertiesPeriodo
		return ret
	}
	return *o.Periodo
}

// GetPeriodoOk returns a tuple with the Periodo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetPeriodoOk() (*RetentionPropertiesPeriodo, bool) {
	if o == nil || IsNil(o.Periodo) {
		return nil, false
	}
	return o.Periodo, true
}

// HasPeriodo returns a boolean if a field has been set.
func (o *RetentionProperties) HasPeriodo() bool {
	if o != nil && !IsNil(o.Periodo) {
		return true
	}

	return false
}

// SetPeriodo gets a reference to the given RetentionPropertiesPeriodo and assigns it to the Periodo field.
func (o *RetentionProperties) SetPeriodo(v RetentionPropertiesPeriodo) {
	o.Periodo = &v
}

// GetTotales returns the Totales field value if set, zero value otherwise.
func (o *RetentionProperties) GetTotales() RetentionPropertiesTotales {
	if o == nil || IsNil(o.Totales) {
		var ret RetentionPropertiesTotales
		return ret
	}
	return *o.Totales
}

// GetTotalesOk returns a tuple with the Totales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetTotalesOk() (*RetentionPropertiesTotales, bool) {
	if o == nil || IsNil(o.Totales) {
		return nil, false
	}
	return o.Totales, true
}

// HasTotales returns a boolean if a field has been set.
func (o *RetentionProperties) HasTotales() bool {
	if o != nil && !IsNil(o.Totales) {
		return true
	}

	return false
}

// SetTotales gets a reference to the given RetentionPropertiesTotales and assigns it to the Totales field.
func (o *RetentionProperties) SetTotales(v RetentionPropertiesTotales) {
	o.Totales = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *RetentionProperties) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *RetentionProperties) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *RetentionProperties) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *RetentionProperties) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *RetentionProperties) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *RetentionProperties) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetComplements returns the Complements field value if set, zero value otherwise.
func (o *RetentionProperties) GetComplements() []string {
	if o == nil || IsNil(o.Complements) {
		var ret []string
		return ret
	}
	return o.Complements
}

// GetComplementsOk returns a tuple with the Complements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetComplementsOk() ([]string, bool) {
	if o == nil || IsNil(o.Complements) {
		return nil, false
	}
	return o.Complements, true
}

// HasComplements returns a boolean if a field has been set.
func (o *RetentionProperties) HasComplements() bool {
	if o != nil && !IsNil(o.Complements) {
		return true
	}

	return false
}

// SetComplements gets a reference to the given []string and assigns it to the Complements field.
func (o *RetentionProperties) SetComplements(v []string) {
	o.Complements = v
}

// GetPdfCustomSection returns the PdfCustomSection field value if set, zero value otherwise.
func (o *RetentionProperties) GetPdfCustomSection() string {
	if o == nil || IsNil(o.PdfCustomSection) {
		var ret string
		return ret
	}
	return *o.PdfCustomSection
}

// GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetPdfCustomSectionOk() (*string, bool) {
	if o == nil || IsNil(o.PdfCustomSection) {
		return nil, false
	}
	return o.PdfCustomSection, true
}

// HasPdfCustomSection returns a boolean if a field has been set.
func (o *RetentionProperties) HasPdfCustomSection() bool {
	if o != nil && !IsNil(o.PdfCustomSection) {
		return true
	}

	return false
}

// SetPdfCustomSection gets a reference to the given string and assigns it to the PdfCustomSection field.
func (o *RetentionProperties) SetPdfCustomSection(v string) {
	o.PdfCustomSection = &v
}

// GetAddenda returns the Addenda field value if set, zero value otherwise.
func (o *RetentionProperties) GetAddenda() string {
	if o == nil || IsNil(o.Addenda) {
		var ret string
		return ret
	}
	return *o.Addenda
}

// GetAddendaOk returns a tuple with the Addenda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetAddendaOk() (*string, bool) {
	if o == nil || IsNil(o.Addenda) {
		return nil, false
	}
	return o.Addenda, true
}

// HasAddenda returns a boolean if a field has been set.
func (o *RetentionProperties) HasAddenda() bool {
	if o != nil && !IsNil(o.Addenda) {
		return true
	}

	return false
}

// SetAddenda gets a reference to the given string and assigns it to the Addenda field.
func (o *RetentionProperties) SetAddenda(v string) {
	o.Addenda = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *RetentionProperties) GetNamespaces() []NamespaceProperties {
	if o == nil || IsNil(o.Namespaces) {
		var ret []NamespaceProperties
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionProperties) GetNamespacesOk() ([]NamespaceProperties, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *RetentionProperties) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []NamespaceProperties and assigns it to the Namespaces field.
func (o *RetentionProperties) SetNamespaces(v []NamespaceProperties) {
	o.Namespaces = v
}

func (o RetentionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CveRetenc) {
		toSerialize["cve_retenc"] = o.CveRetenc
	}
	if !IsNil(o.FechaExp) {
		toSerialize["fecha_exp"] = o.FechaExp
	}
	if !IsNil(o.DescRetenc) {
		toSerialize["desc_retenc"] = o.DescRetenc
	}
	if !IsNil(o.FolioInt) {
		toSerialize["folio_int"] = o.FolioInt
	}
	if !IsNil(o.Periodo) {
		toSerialize["periodo"] = o.Periodo
	}
	if !IsNil(o.Totales) {
		toSerialize["totales"] = o.Totales
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if !IsNil(o.Complements) {
		toSerialize["complements"] = o.Complements
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

type NullableRetentionProperties struct {
	value *RetentionProperties
	isSet bool
}

func (v NullableRetentionProperties) Get() *RetentionProperties {
	return v.value
}

func (v *NullableRetentionProperties) Set(val *RetentionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionProperties(val *RetentionProperties) *NullableRetentionProperties {
	return &NullableRetentionProperties{value: val, isSet: true}
}

func (v NullableRetentionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


