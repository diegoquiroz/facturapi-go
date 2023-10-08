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

// checks if the LineItemInputThirdParty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemInputThirdParty{}

// LineItemInputThirdParty struct for LineItemInputThirdParty
type LineItemInputThirdParty struct {
	// Nombre o razón social del tercero.
	LegalName string `json:"legal_name"`
	// RFC del tercero.
	TaxId string `json:"tax_id"`
	// Régimen fiscal del tercero.
	TaxSystem string `json:"tax_system"`
	// Código postal del tercero.
	Zip string `json:"zip"`
}

// NewLineItemInputThirdParty instantiates a new LineItemInputThirdParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemInputThirdParty(legalName string, taxId string, taxSystem string, zip string) *LineItemInputThirdParty {
	this := LineItemInputThirdParty{}
	this.LegalName = legalName
	this.TaxId = taxId
	this.TaxSystem = taxSystem
	this.Zip = zip
	return &this
}

// NewLineItemInputThirdPartyWithDefaults instantiates a new LineItemInputThirdParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemInputThirdPartyWithDefaults() *LineItemInputThirdParty {
	this := LineItemInputThirdParty{}
	return &this
}

// GetLegalName returns the LegalName field value
func (o *LineItemInputThirdParty) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *LineItemInputThirdParty) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *LineItemInputThirdParty) SetLegalName(v string) {
	o.LegalName = v
}

// GetTaxId returns the TaxId field value
func (o *LineItemInputThirdParty) GetTaxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value
// and a boolean to check if the value has been set.
func (o *LineItemInputThirdParty) GetTaxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxId, true
}

// SetTaxId sets field value
func (o *LineItemInputThirdParty) SetTaxId(v string) {
	o.TaxId = v
}

// GetTaxSystem returns the TaxSystem field value
func (o *LineItemInputThirdParty) GetTaxSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxSystem
}

// GetTaxSystemOk returns a tuple with the TaxSystem field value
// and a boolean to check if the value has been set.
func (o *LineItemInputThirdParty) GetTaxSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaxSystem, true
}

// SetTaxSystem sets field value
func (o *LineItemInputThirdParty) SetTaxSystem(v string) {
	o.TaxSystem = v
}

// GetZip returns the Zip field value
func (o *LineItemInputThirdParty) GetZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zip
}

// GetZipOk returns a tuple with the Zip field value
// and a boolean to check if the value has been set.
func (o *LineItemInputThirdParty) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zip, true
}

// SetZip sets field value
func (o *LineItemInputThirdParty) SetZip(v string) {
	o.Zip = v
}

func (o LineItemInputThirdParty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemInputThirdParty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["legal_name"] = o.LegalName
	toSerialize["tax_id"] = o.TaxId
	toSerialize["tax_system"] = o.TaxSystem
	toSerialize["zip"] = o.Zip
	return toSerialize, nil
}

type NullableLineItemInputThirdParty struct {
	value *LineItemInputThirdParty
	isSet bool
}

func (v NullableLineItemInputThirdParty) Get() *LineItemInputThirdParty {
	return v.value
}

func (v *NullableLineItemInputThirdParty) Set(val *LineItemInputThirdParty) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemInputThirdParty) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemInputThirdParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemInputThirdParty(val *LineItemInputThirdParty) *NullableLineItemInputThirdParty {
	return &NullableLineItemInputThirdParty{value: val, isSet: true}
}

func (v NullableLineItemInputThirdParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemInputThirdParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


