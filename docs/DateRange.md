# DateRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gt** | Pointer to **time.Time** | Límite inferior exclusivo del rango de fechas a solicitar. | [optional] 
**Gte** | Pointer to **time.Time** | Límite inferior inclusivo del rango de fechas a solicitar. | [optional] 
**Lt** | Pointer to **time.Time** | Límite superior exclusivo del rango de fechas a solicitar. | [optional] 
**Lte** | Pointer to **time.Time** | Límite superior inclusivo del rango de fechas a solicitar. | [optional] 

## Methods

### NewDateRange

`func NewDateRange() *DateRange`

NewDateRange instantiates a new DateRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateRangeWithDefaults

`func NewDateRangeWithDefaults() *DateRange`

NewDateRangeWithDefaults instantiates a new DateRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGt

`func (o *DateRange) GetGt() time.Time`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *DateRange) GetGtOk() (*time.Time, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *DateRange) SetGt(v time.Time)`

SetGt sets Gt field to given value.

### HasGt

`func (o *DateRange) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *DateRange) GetGte() time.Time`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *DateRange) GetGteOk() (*time.Time, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *DateRange) SetGte(v time.Time)`

SetGte sets Gte field to given value.

### HasGte

`func (o *DateRange) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *DateRange) GetLt() time.Time`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *DateRange) GetLtOk() (*time.Time, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *DateRange) SetLt(v time.Time)`

SetLt sets Lt field to given value.

### HasLt

`func (o *DateRange) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *DateRange) GetLte() time.Time`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *DateRange) GetLteOk() (*time.Time, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *DateRange) SetLte(v time.Time)`

SetLte sets Lte field to given value.

### HasLte

`func (o *DateRange) HasLte() bool`

HasLte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


