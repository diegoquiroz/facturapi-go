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

// checks if the OrganizationLegal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLegal{}

// OrganizationLegal Datos fiscales de la empresa.
type OrganizationLegal struct {
	// Nombre comercial de la organización.
	Name *string `json:"name,omitempty"`
	// Nombre Fiscal o Razón Social de la organización, *sin* el régimen societario (ej.: S.A. de C.V.). 
	LegalName *string `json:"legal_name,omitempty"`
	// Código de Régimen Fiscal, del [catálogo del SAT](#tipo-de-régimen).
	TaxSystem *string `json:"tax_system,omitempty"`
	// Sitio web de la organización, que se utilizará al enviar la factura por correo electrónico.
	Website *string `json:"website,omitempty"`
	// Teléfono de la organización, que aparecerá en el PDF de la factura.
	Phone *string `json:"phone,omitempty"`
	Address *OrganizationLegalAddress `json:"address,omitempty"`
}

// NewOrganizationLegal instantiates a new OrganizationLegal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLegal() *OrganizationLegal {
	this := OrganizationLegal{}
	return &this
}

// NewOrganizationLegalWithDefaults instantiates a new OrganizationLegal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLegalWithDefaults() *OrganizationLegal {
	this := OrganizationLegal{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationLegal) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationLegal) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationLegal) SetName(v string) {
	o.Name = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *OrganizationLegal) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *OrganizationLegal) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *OrganizationLegal) SetLegalName(v string) {
	o.LegalName = &v
}

// GetTaxSystem returns the TaxSystem field value if set, zero value otherwise.
func (o *OrganizationLegal) GetTaxSystem() string {
	if o == nil || IsNil(o.TaxSystem) {
		var ret string
		return ret
	}
	return *o.TaxSystem
}

// GetTaxSystemOk returns a tuple with the TaxSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetTaxSystemOk() (*string, bool) {
	if o == nil || IsNil(o.TaxSystem) {
		return nil, false
	}
	return o.TaxSystem, true
}

// HasTaxSystem returns a boolean if a field has been set.
func (o *OrganizationLegal) HasTaxSystem() bool {
	if o != nil && !IsNil(o.TaxSystem) {
		return true
	}

	return false
}

// SetTaxSystem gets a reference to the given string and assigns it to the TaxSystem field.
func (o *OrganizationLegal) SetTaxSystem(v string) {
	o.TaxSystem = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *OrganizationLegal) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *OrganizationLegal) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *OrganizationLegal) SetWebsite(v string) {
	o.Website = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrganizationLegal) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrganizationLegal) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrganizationLegal) SetPhone(v string) {
	o.Phone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationLegal) GetAddress() OrganizationLegalAddress {
	if o == nil || IsNil(o.Address) {
		var ret OrganizationLegalAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLegal) GetAddressOk() (*OrganizationLegalAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationLegal) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given OrganizationLegalAddress and assigns it to the Address field.
func (o *OrganizationLegal) SetAddress(v OrganizationLegalAddress) {
	o.Address = &v
}

func (o OrganizationLegal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLegal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LegalName) {
		toSerialize["legal_name"] = o.LegalName
	}
	if !IsNil(o.TaxSystem) {
		toSerialize["tax_system"] = o.TaxSystem
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableOrganizationLegal struct {
	value *OrganizationLegal
	isSet bool
}

func (v NullableOrganizationLegal) Get() *OrganizationLegal {
	return v.value
}

func (v *NullableOrganizationLegal) Set(val *OrganizationLegal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLegal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLegal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLegal(val *OrganizationLegal) *NullableOrganizationLegal {
	return &NullableOrganizationLegal{value: val, isSet: true}
}

func (v NullableOrganizationLegal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLegal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

