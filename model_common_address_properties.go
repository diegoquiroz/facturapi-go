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
)

// checks if the CommonAddressProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAddressProperties{}

// CommonAddressProperties struct for CommonAddressProperties
type CommonAddressProperties struct {
	// Nombre de la calle
	Street *string `json:"street,omitempty"`
	// Número exterior.
	Exterior *string `json:"exterior,omitempty"`
	// Número interior.
	Interior *string `json:"interior,omitempty"`
	// Colonia
	Neighborhood *string `json:"neighborhood,omitempty"`
	// Ciudad
	City *string `json:"city,omitempty"`
	// Municipio o delegación
	Municipality *string `json:"municipality,omitempty"`
	// Código postal
	Zip *string `json:"zip,omitempty"`
}

// NewCommonAddressProperties instantiates a new CommonAddressProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAddressProperties() *CommonAddressProperties {
	this := CommonAddressProperties{}
	return &this
}

// NewCommonAddressPropertiesWithDefaults instantiates a new CommonAddressProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAddressPropertiesWithDefaults() *CommonAddressProperties {
	this := CommonAddressProperties{}
	return &this
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *CommonAddressProperties) SetStreet(v string) {
	o.Street = &v
}

// GetExterior returns the Exterior field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetExterior() string {
	if o == nil || IsNil(o.Exterior) {
		var ret string
		return ret
	}
	return *o.Exterior
}

// GetExteriorOk returns a tuple with the Exterior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetExteriorOk() (*string, bool) {
	if o == nil || IsNil(o.Exterior) {
		return nil, false
	}
	return o.Exterior, true
}

// HasExterior returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasExterior() bool {
	if o != nil && !IsNil(o.Exterior) {
		return true
	}

	return false
}

// SetExterior gets a reference to the given string and assigns it to the Exterior field.
func (o *CommonAddressProperties) SetExterior(v string) {
	o.Exterior = &v
}

// GetInterior returns the Interior field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetInterior() string {
	if o == nil || IsNil(o.Interior) {
		var ret string
		return ret
	}
	return *o.Interior
}

// GetInteriorOk returns a tuple with the Interior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetInteriorOk() (*string, bool) {
	if o == nil || IsNil(o.Interior) {
		return nil, false
	}
	return o.Interior, true
}

// HasInterior returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasInterior() bool {
	if o != nil && !IsNil(o.Interior) {
		return true
	}

	return false
}

// SetInterior gets a reference to the given string and assigns it to the Interior field.
func (o *CommonAddressProperties) SetInterior(v string) {
	o.Interior = &v
}

// GetNeighborhood returns the Neighborhood field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetNeighborhood() string {
	if o == nil || IsNil(o.Neighborhood) {
		var ret string
		return ret
	}
	return *o.Neighborhood
}

// GetNeighborhoodOk returns a tuple with the Neighborhood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetNeighborhoodOk() (*string, bool) {
	if o == nil || IsNil(o.Neighborhood) {
		return nil, false
	}
	return o.Neighborhood, true
}

// HasNeighborhood returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasNeighborhood() bool {
	if o != nil && !IsNil(o.Neighborhood) {
		return true
	}

	return false
}

// SetNeighborhood gets a reference to the given string and assigns it to the Neighborhood field.
func (o *CommonAddressProperties) SetNeighborhood(v string) {
	o.Neighborhood = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *CommonAddressProperties) SetCity(v string) {
	o.City = &v
}

// GetMunicipality returns the Municipality field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetMunicipality() string {
	if o == nil || IsNil(o.Municipality) {
		var ret string
		return ret
	}
	return *o.Municipality
}

// GetMunicipalityOk returns a tuple with the Municipality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetMunicipalityOk() (*string, bool) {
	if o == nil || IsNil(o.Municipality) {
		return nil, false
	}
	return o.Municipality, true
}

// HasMunicipality returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasMunicipality() bool {
	if o != nil && !IsNil(o.Municipality) {
		return true
	}

	return false
}

// SetMunicipality gets a reference to the given string and assigns it to the Municipality field.
func (o *CommonAddressProperties) SetMunicipality(v string) {
	o.Municipality = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *CommonAddressProperties) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAddressProperties) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *CommonAddressProperties) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *CommonAddressProperties) SetZip(v string) {
	o.Zip = &v
}

func (o CommonAddressProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAddressProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.Exterior) {
		toSerialize["exterior"] = o.Exterior
	}
	if !IsNil(o.Interior) {
		toSerialize["interior"] = o.Interior
	}
	if !IsNil(o.Neighborhood) {
		toSerialize["neighborhood"] = o.Neighborhood
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Municipality) {
		toSerialize["municipality"] = o.Municipality
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableCommonAddressProperties struct {
	value *CommonAddressProperties
	isSet bool
}

func (v NullableCommonAddressProperties) Get() *CommonAddressProperties {
	return v.value
}

func (v *NullableCommonAddressProperties) Set(val *CommonAddressProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAddressProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAddressProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAddressProperties(val *CommonAddressProperties) *NullableCommonAddressProperties {
	return &NullableCommonAddressProperties{value: val, isSet: true}
}

func (v NullableCommonAddressProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAddressProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


