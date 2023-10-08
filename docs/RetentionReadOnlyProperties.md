# RetentionReadOnlyProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Estado actual de la retención.  | [optional] 
**VerificationUrl** | Pointer to **string** | Dirección URL para verificar el estado de la retención en el portal del SAT. Este link es el mismo que aparece en el código QR, en el PDF de la retención. | [optional] 
**Type** | Pointer to **string** | Tipo de comprobante. | [optional] 
**Uuid** | Pointer to **string** | Folio fiscal de la retención, asignado por el SAT. | [optional] 
**Stamp** | Pointer to [**Stamp**](Stamp.md) |  | [optional] 
**Customer** | Pointer to [**CuustomerInfo**](CuustomerInfo.md) |  | [optional] 

## Methods

### NewRetentionReadOnlyProperties

`func NewRetentionReadOnlyProperties() *RetentionReadOnlyProperties`

NewRetentionReadOnlyProperties instantiates a new RetentionReadOnlyProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionReadOnlyPropertiesWithDefaults

`func NewRetentionReadOnlyPropertiesWithDefaults() *RetentionReadOnlyProperties`

NewRetentionReadOnlyPropertiesWithDefaults instantiates a new RetentionReadOnlyProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RetentionReadOnlyProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RetentionReadOnlyProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RetentionReadOnlyProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RetentionReadOnlyProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationUrl

`func (o *RetentionReadOnlyProperties) GetVerificationUrl() string`

GetVerificationUrl returns the VerificationUrl field if non-nil, zero value otherwise.

### GetVerificationUrlOk

`func (o *RetentionReadOnlyProperties) GetVerificationUrlOk() (*string, bool)`

GetVerificationUrlOk returns a tuple with the VerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUrl

`func (o *RetentionReadOnlyProperties) SetVerificationUrl(v string)`

SetVerificationUrl sets VerificationUrl field to given value.

### HasVerificationUrl

`func (o *RetentionReadOnlyProperties) HasVerificationUrl() bool`

HasVerificationUrl returns a boolean if a field has been set.

### GetType

`func (o *RetentionReadOnlyProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RetentionReadOnlyProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RetentionReadOnlyProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RetentionReadOnlyProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RetentionReadOnlyProperties) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RetentionReadOnlyProperties) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RetentionReadOnlyProperties) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RetentionReadOnlyProperties) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStamp

`func (o *RetentionReadOnlyProperties) GetStamp() Stamp`

GetStamp returns the Stamp field if non-nil, zero value otherwise.

### GetStampOk

`func (o *RetentionReadOnlyProperties) GetStampOk() (*Stamp, bool)`

GetStampOk returns a tuple with the Stamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamp

`func (o *RetentionReadOnlyProperties) SetStamp(v Stamp)`

SetStamp sets Stamp field to given value.

### HasStamp

`func (o *RetentionReadOnlyProperties) HasStamp() bool`

HasStamp returns a boolean if a field has been set.

### GetCustomer

`func (o *RetentionReadOnlyProperties) GetCustomer() CuustomerInfo`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *RetentionReadOnlyProperties) GetCustomerOk() (*CuustomerInfo, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *RetentionReadOnlyProperties) SetCustomer(v CuustomerInfo)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *RetentionReadOnlyProperties) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


