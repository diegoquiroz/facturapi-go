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

// checks if the RetentionSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetentionSearchResult{}

// RetentionSearchResult struct for RetentionSearchResult
type RetentionSearchResult struct {
	// Número de página de resultados
	Page *int32 `json:"page,omitempty"`
	// Número total de páginas de resultados
	TotalPages *int32 `json:"total_pages,omitempty"`
	// Número de elementos individuales en todas las páginas de resultados
	TotalResults *int32 `json:"total_results,omitempty"`
	Data []Retention `json:"data,omitempty"`
}

// NewRetentionSearchResult instantiates a new RetentionSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetentionSearchResult() *RetentionSearchResult {
	this := RetentionSearchResult{}
	return &this
}

// NewRetentionSearchResultWithDefaults instantiates a new RetentionSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetentionSearchResultWithDefaults() *RetentionSearchResult {
	this := RetentionSearchResult{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *RetentionSearchResult) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchResult) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *RetentionSearchResult) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *RetentionSearchResult) SetPage(v int32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *RetentionSearchResult) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchResult) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *RetentionSearchResult) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *RetentionSearchResult) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *RetentionSearchResult) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchResult) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *RetentionSearchResult) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *RetentionSearchResult) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RetentionSearchResult) GetData() []Retention {
	if o == nil || IsNil(o.Data) {
		var ret []Retention
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearchResult) GetDataOk() ([]Retention, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RetentionSearchResult) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Retention and assigns it to the Data field.
func (o *RetentionSearchResult) SetData(v []Retention) {
	o.Data = v
}

func (o RetentionSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetentionSearchResult) ToMap() (map[string]interface{}, error) {
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

type NullableRetentionSearchResult struct {
	value *RetentionSearchResult
	isSet bool
}

func (v NullableRetentionSearchResult) Get() *RetentionSearchResult {
	return v.value
}

func (v *NullableRetentionSearchResult) Set(val *RetentionSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRetentionSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRetentionSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetentionSearchResult(val *RetentionSearchResult) *NullableRetentionSearchResult {
	return &NullableRetentionSearchResult{value: val, isSet: true}
}

func (v NullableRetentionSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetentionSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


