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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	// ID del objeto
	Id *string `json:"id,omitempty"`
	// Fecha de registro
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Indica si la organización tiene información necesaria para facturar en ambiente Live.
	IsProductionReady *bool `json:"is_production_ready,omitempty"`
	PendingSteps []OrganizationPendingStepsInner `json:"pending_steps,omitempty"`
	Legal *OrganizationLegal `json:"legal,omitempty"`
	Customization *OrganizationCustomization `json:"customization,omitempty"`
	Certificate *OrganizationCertificate `json:"certificate,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Organization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Organization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Organization) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetIsProductionReady returns the IsProductionReady field value if set, zero value otherwise.
func (o *Organization) GetIsProductionReady() bool {
	if o == nil || IsNil(o.IsProductionReady) {
		var ret bool
		return ret
	}
	return *o.IsProductionReady
}

// GetIsProductionReadyOk returns a tuple with the IsProductionReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetIsProductionReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProductionReady) {
		return nil, false
	}
	return o.IsProductionReady, true
}

// HasIsProductionReady returns a boolean if a field has been set.
func (o *Organization) HasIsProductionReady() bool {
	if o != nil && !IsNil(o.IsProductionReady) {
		return true
	}

	return false
}

// SetIsProductionReady gets a reference to the given bool and assigns it to the IsProductionReady field.
func (o *Organization) SetIsProductionReady(v bool) {
	o.IsProductionReady = &v
}

// GetPendingSteps returns the PendingSteps field value if set, zero value otherwise.
func (o *Organization) GetPendingSteps() []OrganizationPendingStepsInner {
	if o == nil || IsNil(o.PendingSteps) {
		var ret []OrganizationPendingStepsInner
		return ret
	}
	return o.PendingSteps
}

// GetPendingStepsOk returns a tuple with the PendingSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPendingStepsOk() ([]OrganizationPendingStepsInner, bool) {
	if o == nil || IsNil(o.PendingSteps) {
		return nil, false
	}
	return o.PendingSteps, true
}

// HasPendingSteps returns a boolean if a field has been set.
func (o *Organization) HasPendingSteps() bool {
	if o != nil && !IsNil(o.PendingSteps) {
		return true
	}

	return false
}

// SetPendingSteps gets a reference to the given []OrganizationPendingStepsInner and assigns it to the PendingSteps field.
func (o *Organization) SetPendingSteps(v []OrganizationPendingStepsInner) {
	o.PendingSteps = v
}

// GetLegal returns the Legal field value if set, zero value otherwise.
func (o *Organization) GetLegal() OrganizationLegal {
	if o == nil || IsNil(o.Legal) {
		var ret OrganizationLegal
		return ret
	}
	return *o.Legal
}

// GetLegalOk returns a tuple with the Legal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetLegalOk() (*OrganizationLegal, bool) {
	if o == nil || IsNil(o.Legal) {
		return nil, false
	}
	return o.Legal, true
}

// HasLegal returns a boolean if a field has been set.
func (o *Organization) HasLegal() bool {
	if o != nil && !IsNil(o.Legal) {
		return true
	}

	return false
}

// SetLegal gets a reference to the given OrganizationLegal and assigns it to the Legal field.
func (o *Organization) SetLegal(v OrganizationLegal) {
	o.Legal = &v
}

// GetCustomization returns the Customization field value if set, zero value otherwise.
func (o *Organization) GetCustomization() OrganizationCustomization {
	if o == nil || IsNil(o.Customization) {
		var ret OrganizationCustomization
		return ret
	}
	return *o.Customization
}

// GetCustomizationOk returns a tuple with the Customization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCustomizationOk() (*OrganizationCustomization, bool) {
	if o == nil || IsNil(o.Customization) {
		return nil, false
	}
	return o.Customization, true
}

// HasCustomization returns a boolean if a field has been set.
func (o *Organization) HasCustomization() bool {
	if o != nil && !IsNil(o.Customization) {
		return true
	}

	return false
}

// SetCustomization gets a reference to the given OrganizationCustomization and assigns it to the Customization field.
func (o *Organization) SetCustomization(v OrganizationCustomization) {
	o.Customization = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *Organization) GetCertificate() OrganizationCertificate {
	if o == nil || IsNil(o.Certificate) {
		var ret OrganizationCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCertificateOk() (*OrganizationCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *Organization) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given OrganizationCertificate and assigns it to the Certificate field.
func (o *Organization) SetCertificate(v OrganizationCertificate) {
	o.Certificate = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.IsProductionReady) {
		toSerialize["is_production_ready"] = o.IsProductionReady
	}
	if !IsNil(o.PendingSteps) {
		toSerialize["pending_steps"] = o.PendingSteps
	}
	if !IsNil(o.Legal) {
		toSerialize["legal"] = o.Legal
	}
	if !IsNil(o.Customization) {
		toSerialize["customization"] = o.Customization
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


