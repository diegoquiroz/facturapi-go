# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha de registro | [optional] 
**IsProductionReady** | Pointer to **bool** | Indica si la organización tiene información necesaria para facturar en ambiente Live. | [optional] 
**PendingSteps** | Pointer to [**[]OrganizationPendingStepsInner**](OrganizationPendingStepsInner.md) |  | [optional] 
**Legal** | Pointer to [**OrganizationLegal**](OrganizationLegal.md) |  | [optional] 
**Customization** | Pointer to [**OrganizationCustomization**](OrganizationCustomization.md) |  | [optional] 
**Certificate** | Pointer to [**OrganizationCertificate**](OrganizationCertificate.md) |  | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsProductionReady

`func (o *Organization) GetIsProductionReady() bool`

GetIsProductionReady returns the IsProductionReady field if non-nil, zero value otherwise.

### GetIsProductionReadyOk

`func (o *Organization) GetIsProductionReadyOk() (*bool, bool)`

GetIsProductionReadyOk returns a tuple with the IsProductionReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProductionReady

`func (o *Organization) SetIsProductionReady(v bool)`

SetIsProductionReady sets IsProductionReady field to given value.

### HasIsProductionReady

`func (o *Organization) HasIsProductionReady() bool`

HasIsProductionReady returns a boolean if a field has been set.

### GetPendingSteps

`func (o *Organization) GetPendingSteps() []OrganizationPendingStepsInner`

GetPendingSteps returns the PendingSteps field if non-nil, zero value otherwise.

### GetPendingStepsOk

`func (o *Organization) GetPendingStepsOk() (*[]OrganizationPendingStepsInner, bool)`

GetPendingStepsOk returns a tuple with the PendingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingSteps

`func (o *Organization) SetPendingSteps(v []OrganizationPendingStepsInner)`

SetPendingSteps sets PendingSteps field to given value.

### HasPendingSteps

`func (o *Organization) HasPendingSteps() bool`

HasPendingSteps returns a boolean if a field has been set.

### GetLegal

`func (o *Organization) GetLegal() OrganizationLegal`

GetLegal returns the Legal field if non-nil, zero value otherwise.

### GetLegalOk

`func (o *Organization) GetLegalOk() (*OrganizationLegal, bool)`

GetLegalOk returns a tuple with the Legal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegal

`func (o *Organization) SetLegal(v OrganizationLegal)`

SetLegal sets Legal field to given value.

### HasLegal

`func (o *Organization) HasLegal() bool`

HasLegal returns a boolean if a field has been set.

### GetCustomization

`func (o *Organization) GetCustomization() OrganizationCustomization`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *Organization) GetCustomizationOk() (*OrganizationCustomization, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *Organization) SetCustomization(v OrganizationCustomization)`

SetCustomization sets Customization field to given value.

### HasCustomization

`func (o *Organization) HasCustomization() bool`

HasCustomization returns a boolean if a field has been set.

### GetCertificate

`func (o *Organization) GetCertificate() OrganizationCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Organization) GetCertificateOk() (*OrganizationCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Organization) SetCertificate(v OrganizationCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Organization) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


