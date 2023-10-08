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

// checks if the InvoiceSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceSearchResult{}

// InvoiceSearchResult struct for InvoiceSearchResult
type InvoiceSearchResult struct {
	// Número de página de resultados
	Page *int32 `json:"page,omitempty"`
	// Número total de páginas de resultados
	TotalPages *int32 `json:"total_pages,omitempty"`
	// Número de elementos individuales en todas las páginas de resultados
	TotalResults *int32 `json:"total_results,omitempty"`
	Data []Invoice `json:"data,omitempty"`
}

// NewInvoiceSearchResult instantiates a new InvoiceSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceSearchResult() *InvoiceSearchResult {
	this := InvoiceSearchResult{}
	return &this
}

// NewInvoiceSearchResultWithDefaults instantiates a new InvoiceSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceSearchResultWithDefaults() *InvoiceSearchResult {
	this := InvoiceSearchResult{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *InvoiceSearchResult) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResult) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *InvoiceSearchResult) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *InvoiceSearchResult) SetPage(v int32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *InvoiceSearchResult) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResult) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *InvoiceSearchResult) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *InvoiceSearchResult) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *InvoiceSearchResult) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResult) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *InvoiceSearchResult) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *InvoiceSearchResult) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InvoiceSearchResult) GetData() []Invoice {
	if o == nil || IsNil(o.Data) {
		var ret []Invoice
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResult) GetDataOk() ([]Invoice, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InvoiceSearchResult) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Invoice and assigns it to the Data field.
func (o *InvoiceSearchResult) SetData(v []Invoice) {
	o.Data = v
}

func (o InvoiceSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["total_results"] = o.TotalResults
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInvoiceSearchResult struct {
	value *InvoiceSearchResult
	isSet bool
}

func (v NullableInvoiceSearchResult) Get() *InvoiceSearchResult {
	return v.value
}

func (v *NullableInvoiceSearchResult) Set(val *InvoiceSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceSearchResult(val *InvoiceSearchResult) *NullableInvoiceSearchResult {
	return &NullableInvoiceSearchResult{value: val, isSet: true}
}

func (v NullableInvoiceSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

