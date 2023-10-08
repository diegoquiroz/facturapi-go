# CustomerPropertiesAllOfAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | Pointer to **string** | Nombre de la calle | [optional] 
**Exterior** | Pointer to **string** | Número exterior. | [optional] 
**Interior** | Pointer to **string** | Número interior. | [optional] 
**Neighborhood** | Pointer to **string** | Colonia | [optional] 
**City** | Pointer to **string** | Ciudad | [optional] 
**Municipality** | Pointer to **string** | Municipio o delegación | [optional] 
**Zip** | Pointer to **string** | Código postal | [optional] 
**State** | Pointer to **string** | Si el país es México (\&quot;MEX\&quot;), contiene el nombre del Estado o Entidad Federativa. Para extranjeros contiene el código de Estado de acuerdo al estándar [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2), que puedes consultar en nuestro [Catálogo de Estados](https://dashboard.facturapi.io/catalogs/state). | [optional] 
**Country** | Pointer to **string** | Código de país acorde al estándar [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3), del [Catálogo de Países](https://dashboard.facturapi.io/catalogs/country). | [optional] [default to "MEX"]

## Methods

### NewCustomerPropertiesAllOfAddress

`func NewCustomerPropertiesAllOfAddress() *CustomerPropertiesAllOfAddress`

NewCustomerPropertiesAllOfAddress instantiates a new CustomerPropertiesAllOfAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPropertiesAllOfAddressWithDefaults

`func NewCustomerPropertiesAllOfAddressWithDefaults() *CustomerPropertiesAllOfAddress`

NewCustomerPropertiesAllOfAddressWithDefaults instantiates a new CustomerPropertiesAllOfAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *CustomerPropertiesAllOfAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *CustomerPropertiesAllOfAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *CustomerPropertiesAllOfAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *CustomerPropertiesAllOfAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetExterior

`func (o *CustomerPropertiesAllOfAddress) GetExterior() string`

GetExterior returns the Exterior field if non-nil, zero value otherwise.

### GetExteriorOk

`func (o *CustomerPropertiesAllOfAddress) GetExteriorOk() (*string, bool)`

GetExteriorOk returns a tuple with the Exterior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterior

`func (o *CustomerPropertiesAllOfAddress) SetExterior(v string)`

SetExterior sets Exterior field to given value.

### HasExterior

`func (o *CustomerPropertiesAllOfAddress) HasExterior() bool`

HasExterior returns a boolean if a field has been set.

### GetInterior

`func (o *CustomerPropertiesAllOfAddress) GetInterior() string`

GetInterior returns the Interior field if non-nil, zero value otherwise.

### GetInteriorOk

`func (o *CustomerPropertiesAllOfAddress) GetInteriorOk() (*string, bool)`

GetInteriorOk returns a tuple with the Interior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterior

`func (o *CustomerPropertiesAllOfAddress) SetInterior(v string)`

SetInterior sets Interior field to given value.

### HasInterior

`func (o *CustomerPropertiesAllOfAddress) HasInterior() bool`

HasInterior returns a boolean if a field has been set.

### GetNeighborhood

`func (o *CustomerPropertiesAllOfAddress) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *CustomerPropertiesAllOfAddress) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *CustomerPropertiesAllOfAddress) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *CustomerPropertiesAllOfAddress) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetCity

`func (o *CustomerPropertiesAllOfAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerPropertiesAllOfAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerPropertiesAllOfAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerPropertiesAllOfAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetMunicipality

`func (o *CustomerPropertiesAllOfAddress) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *CustomerPropertiesAllOfAddress) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *CustomerPropertiesAllOfAddress) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *CustomerPropertiesAllOfAddress) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### GetZip

`func (o *CustomerPropertiesAllOfAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *CustomerPropertiesAllOfAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *CustomerPropertiesAllOfAddress) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *CustomerPropertiesAllOfAddress) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetState

`func (o *CustomerPropertiesAllOfAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerPropertiesAllOfAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerPropertiesAllOfAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerPropertiesAllOfAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerPropertiesAllOfAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerPropertiesAllOfAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerPropertiesAllOfAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerPropertiesAllOfAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


