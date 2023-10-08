# OrganizationLegalInputAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | **string** | Nombre de la calle | 
**Exterior** | **string** | Número exterior. | 
**Interior** | Pointer to **string** | Número interior. | [optional] 
**Neighborhood** | Pointer to **string** | Colonia | [optional] 
**City** | Pointer to **string** | Ciudad | [optional] 
**Municipality** | Pointer to **string** | Municipio o delegación | [optional] 
**Zip** | **string** | Código postal | 
**State** | Pointer to **string** | Nombre del Estado o Entidad Federativa. | [optional] 

## Methods

### NewOrganizationLegalInputAddress

`func NewOrganizationLegalInputAddress(street string, exterior string, zip string, ) *OrganizationLegalInputAddress`

NewOrganizationLegalInputAddress instantiates a new OrganizationLegalInputAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationLegalInputAddressWithDefaults

`func NewOrganizationLegalInputAddressWithDefaults() *OrganizationLegalInputAddress`

NewOrganizationLegalInputAddressWithDefaults instantiates a new OrganizationLegalInputAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *OrganizationLegalInputAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *OrganizationLegalInputAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *OrganizationLegalInputAddress) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetExterior

`func (o *OrganizationLegalInputAddress) GetExterior() string`

GetExterior returns the Exterior field if non-nil, zero value otherwise.

### GetExteriorOk

`func (o *OrganizationLegalInputAddress) GetExteriorOk() (*string, bool)`

GetExteriorOk returns a tuple with the Exterior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterior

`func (o *OrganizationLegalInputAddress) SetExterior(v string)`

SetExterior sets Exterior field to given value.


### GetInterior

`func (o *OrganizationLegalInputAddress) GetInterior() string`

GetInterior returns the Interior field if non-nil, zero value otherwise.

### GetInteriorOk

`func (o *OrganizationLegalInputAddress) GetInteriorOk() (*string, bool)`

GetInteriorOk returns a tuple with the Interior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterior

`func (o *OrganizationLegalInputAddress) SetInterior(v string)`

SetInterior sets Interior field to given value.

### HasInterior

`func (o *OrganizationLegalInputAddress) HasInterior() bool`

HasInterior returns a boolean if a field has been set.

### GetNeighborhood

`func (o *OrganizationLegalInputAddress) GetNeighborhood() string`

GetNeighborhood returns the Neighborhood field if non-nil, zero value otherwise.

### GetNeighborhoodOk

`func (o *OrganizationLegalInputAddress) GetNeighborhoodOk() (*string, bool)`

GetNeighborhoodOk returns a tuple with the Neighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborhood

`func (o *OrganizationLegalInputAddress) SetNeighborhood(v string)`

SetNeighborhood sets Neighborhood field to given value.

### HasNeighborhood

`func (o *OrganizationLegalInputAddress) HasNeighborhood() bool`

HasNeighborhood returns a boolean if a field has been set.

### GetCity

`func (o *OrganizationLegalInputAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrganizationLegalInputAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrganizationLegalInputAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *OrganizationLegalInputAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetMunicipality

`func (o *OrganizationLegalInputAddress) GetMunicipality() string`

GetMunicipality returns the Municipality field if non-nil, zero value otherwise.

### GetMunicipalityOk

`func (o *OrganizationLegalInputAddress) GetMunicipalityOk() (*string, bool)`

GetMunicipalityOk returns a tuple with the Municipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMunicipality

`func (o *OrganizationLegalInputAddress) SetMunicipality(v string)`

SetMunicipality sets Municipality field to given value.

### HasMunicipality

`func (o *OrganizationLegalInputAddress) HasMunicipality() bool`

HasMunicipality returns a boolean if a field has been set.

### GetZip

`func (o *OrganizationLegalInputAddress) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *OrganizationLegalInputAddress) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *OrganizationLegalInputAddress) SetZip(v string)`

SetZip sets Zip field to given value.


### GetState

`func (o *OrganizationLegalInputAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrganizationLegalInputAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrganizationLegalInputAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrganizationLegalInputAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


