# OrganizationCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasCertificates** | Pointer to **bool** | Indica si la organización ya tiene el Certificado de Sello Digital (CSD) cargado. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Fecha de la última actualización del certificado. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Fecha de expiración del certificado. | [optional] 

## Methods

### NewOrganizationCertificate

`func NewOrganizationCertificate() *OrganizationCertificate`

NewOrganizationCertificate instantiates a new OrganizationCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCertificateWithDefaults

`func NewOrganizationCertificateWithDefaults() *OrganizationCertificate`

NewOrganizationCertificateWithDefaults instantiates a new OrganizationCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasCertificates

`func (o *OrganizationCertificate) GetHasCertificates() bool`

GetHasCertificates returns the HasCertificates field if non-nil, zero value otherwise.

### GetHasCertificatesOk

`func (o *OrganizationCertificate) GetHasCertificatesOk() (*bool, bool)`

GetHasCertificatesOk returns a tuple with the HasCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCertificates

`func (o *OrganizationCertificate) SetHasCertificates(v bool)`

SetHasCertificates sets HasCertificates field to given value.

### HasHasCertificates

`func (o *OrganizationCertificate) HasHasCertificates() bool`

HasHasCertificates returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganizationCertificate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationCertificate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationCertificate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganizationCertificate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrganizationCertificate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrganizationCertificate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrganizationCertificate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrganizationCertificate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


