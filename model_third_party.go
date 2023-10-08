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

// checks if the ThirdParty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdParty{}

// ThirdParty Objeto con información del contribuyente tercero, a cuenta del que se realiza la operación.  Corresponde al campo \"ACuentaTerceros\" en el CFDI. 
type ThirdParty struct {
	// Nombre o razón social del tercero.
	LegalName *string `json:"legal_name,omitempty"`
	// RFC del tercero.
	TaxId *string `json:"tax_id,omitempty"`
	// Régimen fiscal del tercero.
	TaxSystem *string `json:"tax_system,omitempty"`
	// Código postal del tercero.
	Zip *string `json:"zip,omitempty"`
}

// NewThirdParty instantiates a new ThirdParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdParty() *ThirdParty {
	this := ThirdParty{}
	return &this
}

// NewThirdPartyWithDefaults instantiates a new ThirdParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyWithDefaults() *ThirdParty {
	this := ThirdParty{}
	return &this
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *ThirdParty) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdParty) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *ThirdParty) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *ThirdParty) SetLegalName(v string) {
	o.LegalName = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *ThirdParty) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdParty) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *ThirdParty) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *ThirdParty) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxSystem returns the TaxSystem field value if set, zero value otherwise.
func (o *ThirdParty) GetTaxSystem() string {
	if o == nil || IsNil(o.TaxSystem) {
		var ret string
		return ret
	}
	return *o.TaxSystem
}

// GetTaxSystemOk returns a tuple with the TaxSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdParty) GetTaxSystemOk() (*string, bool) {
	if o == nil || IsNil(o.TaxSystem) {
		return nil, false
	}
	return o.TaxSystem, true
}

// HasTaxSystem returns a boolean if a field has been set.
func (o *ThirdParty) HasTaxSystem() bool {
	if o != nil && !IsNil(o.TaxSystem) {
		return true
	}

	return false
}

// SetTaxSystem gets a reference to the given string and assigns it to the TaxSystem field.
func (o *ThirdParty) SetTaxSystem(v string) {
	o.TaxSystem = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *ThirdParty) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdParty) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *ThirdParty) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *ThirdParty) SetZip(v string) {
	o.Zip = &v
}

func (o ThirdParty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdParty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalName) {
		toSerialize["legal_name"] = o.LegalName
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.TaxSystem) {
		toSerialize["tax_system"] = o.TaxSystem
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableThirdParty struct {
	value *ThirdParty
	isSet bool
}

func (v NullableThirdParty) Get() *ThirdParty {
	return v.value
}

func (v *NullableThirdParty) Set(val *ThirdParty) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdParty) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdParty(val *ThirdParty) *NullableThirdParty {
	return &NullableThirdParty{value: val, isSet: true}
}

func (v NullableThirdParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


