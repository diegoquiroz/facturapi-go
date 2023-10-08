# EventBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del evento | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha y hora de creación del evento | [optional] 
**Livemode** | Pointer to **bool** | Indica si el evento se generó en modo test (false) o en modo producción (true). | [optional] 
**Organization** | Pointer to **string** | ID de la organización a la que pertenece el evento | [optional] 

## Methods

### NewEventBase

`func NewEventBase() *EventBase`

NewEventBase instantiates a new EventBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBaseWithDefaults

`func NewEventBaseWithDefaults() *EventBase`

NewEventBaseWithDefaults instantiates a new EventBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EventBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EventBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EventBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EventBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLivemode

`func (o *EventBase) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *EventBase) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *EventBase) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *EventBase) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetOrganization

`func (o *EventBase) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EventBase) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EventBase) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *EventBase) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


