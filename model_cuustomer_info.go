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

// checks if the CuustomerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CuustomerInfo{}

// CuustomerInfo Objeto con información parcial del cliente receptor del comprobante. Para obtener el objeto `Customer` completo, deberás consultarlo con el método de [Obtener Cliente]('#/operation/getCustomer').
type CuustomerInfo struct {
	// ID del objeto `customer` relacionado a la factura, en caso de no haber sido eliminado
	Id *string `json:"id,omitempty"`
	// Nombre Fiscal o Razón Social del cliente, *sin* incluir el régimen societario (ej.: S.A. de C.V.). 
	LegalName *string `json:"legal_name,omitempty"`
	// RFC del cliente.
	TaxId *string `json:"tax_id,omitempty"`
	Address *CuustomerInfoAddress `json:"address,omitempty"`
}

// NewCuustomerInfo instantiates a new CuustomerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCuustomerInfo() *CuustomerInfo {
	this := CuustomerInfo{}
	return &this
}

// NewCuustomerInfoWithDefaults instantiates a new CuustomerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCuustomerInfoWithDefaults() *CuustomerInfo {
	this := CuustomerInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CuustomerInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CuustomerInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CuustomerInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CuustomerInfo) SetId(v string) {
	o.Id = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *CuustomerInfo) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CuustomerInfo) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *CuustomerInfo) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *CuustomerInfo) SetLegalName(v string) {
	o.LegalName = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *CuustomerInfo) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CuustomerInfo) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *CuustomerInfo) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *CuustomerInfo) SetTaxId(v string) {
	o.TaxId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CuustomerInfo) GetAddress() CuustomerInfoAddress {
	if o == nil || IsNil(o.Address) {
		var ret CuustomerInfoAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CuustomerInfo) GetAddressOk() (*CuustomerInfoAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CuustomerInfo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CuustomerInfoAddress and assigns it to the Address field.
func (o *CuustomerInfo) SetAddress(v CuustomerInfoAddress) {
	o.Address = &v
}

func (o CuustomerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CuustomerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LegalName) {
		toSerialize["legal_name"] = o.LegalName
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableCuustomerInfo struct {
	value *CuustomerInfo
	isSet bool
}

func (v NullableCuustomerInfo) Get() *CuustomerInfo {
	return v.value
}

func (v *NullableCuustomerInfo) Set(val *CuustomerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCuustomerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCuustomerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCuustomerInfo(val *CuustomerInfo) *NullableCuustomerInfo {
	return &NullableCuustomerInfo{value: val, isSet: true}
}

func (v NullableCuustomerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCuustomerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


