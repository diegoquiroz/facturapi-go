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
	"time"
)

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Customer{}

// Customer struct for Customer
type Customer struct {
	// ID del objeto
	Id *string `json:"id,omitempty"`
	// Fecha de registro
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Si el valor es `true`, indica que el objeto fue creado en ambiente Live; o si es `false`, en ambiente Test.
	Livemode *bool `json:"livemode,omitempty"`
	// Nombre Fiscal o Razón Social del cliente. *sin* el régimen societario (ej.: S.A. de C.V.). 
	LegalName *string `json:"legal_name,omitempty"`
	// En clientes de México contiene el RFC del cliente. Para extranjeros es opcional y representa el número de registro de identificacón tributaria, es decir, el equivalente al RFC en el país del cliente.
	TaxId *string `json:"tax_id,omitempty"`
	// Requerido para clientes nacionales. Clave del régimen fiscal del cliente.
	TaxSystem *string `json:"tax_system,omitempty"`
	// Dirección de correo electrónico al cual enviar las facturas generadas.
	Email *string `json:"email,omitempty"`
	// Teléfono del cliente.
	Phone *string `json:"phone,omitempty"`
	Address *CustomerPropertiesAllOfAddress `json:"address,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer() *Customer {
	this := Customer{}
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Customer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Customer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Customer) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Customer) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Customer) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Customer) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *Customer) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *Customer) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *Customer) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *Customer) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *Customer) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *Customer) SetLegalName(v string) {
	o.LegalName = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *Customer) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *Customer) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *Customer) SetTaxId(v string) {
	o.TaxId = &v
}

// GetTaxSystem returns the TaxSystem field value if set, zero value otherwise.
func (o *Customer) GetTaxSystem() string {
	if o == nil || IsNil(o.TaxSystem) {
		var ret string
		return ret
	}
	return *o.TaxSystem
}

// GetTaxSystemOk returns a tuple with the TaxSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetTaxSystemOk() (*string, bool) {
	if o == nil || IsNil(o.TaxSystem) {
		return nil, false
	}
	return o.TaxSystem, true
}

// HasTaxSystem returns a boolean if a field has been set.
func (o *Customer) HasTaxSystem() bool {
	if o != nil && !IsNil(o.TaxSystem) {
		return true
	}

	return false
}

// SetTaxSystem gets a reference to the given string and assigns it to the TaxSystem field.
func (o *Customer) SetTaxSystem(v string) {
	o.TaxSystem = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Customer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Customer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Customer) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Customer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Customer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Customer) SetPhone(v string) {
	o.Phone = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Customer) GetAddress() CustomerPropertiesAllOfAddress {
	if o == nil || IsNil(o.Address) {
		var ret CustomerPropertiesAllOfAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetAddressOk() (*CustomerPropertiesAllOfAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Customer) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CustomerPropertiesAllOfAddress and assigns it to the Address field.
func (o *Customer) SetAddress(v CustomerPropertiesAllOfAddress) {
	o.Address = &v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.LegalName) {
		toSerialize["legal_name"] = o.LegalName
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.TaxSystem) {
		toSerialize["tax_system"] = o.TaxSystem
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


