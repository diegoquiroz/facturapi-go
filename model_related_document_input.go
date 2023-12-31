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

// checks if the RelatedDocumentInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedDocumentInput{}

// RelatedDocumentInput struct for RelatedDocumentInput
type RelatedDocumentInput struct {
	// Clave de relación del catálogo del SAT (que puedes consultar en [esta tabla](#relacion-entre-facturas)). Es requerido cuando se envíe el parámetro `related`.
	Relationship string `json:"relationship"`
	Documents []string `json:"documents,omitempty"`
}

// NewRelatedDocumentInput instantiates a new RelatedDocumentInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedDocumentInput(relationship string) *RelatedDocumentInput {
	this := RelatedDocumentInput{}
	this.Relationship = relationship
	return &this
}

// NewRelatedDocumentInputWithDefaults instantiates a new RelatedDocumentInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedDocumentInputWithDefaults() *RelatedDocumentInput {
	this := RelatedDocumentInput{}
	return &this
}

// GetRelationship returns the Relationship field value
func (o *RelatedDocumentInput) GetRelationship() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value
// and a boolean to check if the value has been set.
func (o *RelatedDocumentInput) GetRelationshipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationship, true
}

// SetRelationship sets field value
func (o *RelatedDocumentInput) SetRelationship(v string) {
	o.Relationship = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *RelatedDocumentInput) GetDocuments() []string {
	if o == nil || IsNil(o.Documents) {
		var ret []string
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedDocumentInput) GetDocumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *RelatedDocumentInput) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []string and assigns it to the Documents field.
func (o *RelatedDocumentInput) SetDocuments(v []string) {
	o.Documents = v
}

func (o RelatedDocumentInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedDocumentInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["relationship"] = o.Relationship
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	return toSerialize, nil
}

type NullableRelatedDocumentInput struct {
	value *RelatedDocumentInput
	isSet bool
}

func (v NullableRelatedDocumentInput) Get() *RelatedDocumentInput {
	return v.value
}

func (v *NullableRelatedDocumentInput) Set(val *RelatedDocumentInput) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedDocumentInput) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedDocumentInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedDocumentInput(val *RelatedDocumentInput) *NullableRelatedDocumentInput {
	return &NullableRelatedDocumentInput{value: val, isSet: true}
}

func (v NullableRelatedDocumentInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedDocumentInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


