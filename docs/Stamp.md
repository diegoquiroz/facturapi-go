# Stamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to **string** | Sello digital del comprobante fiscal. | [optional] 
**Date** | Pointer to **time.Time** | Fecha de timbrado en formato ISO8601 (UTC String). | [optional] 
**SatCertNumber** | Pointer to **string** | Número de serie del certificado del SAT usado para timbrar. | [optional] 
**SatSignature** | Pointer to **string** | Sello digital del timbre fiscal digital. | [optional] 

## Methods

### NewStamp

`func NewStamp() *Stamp`

NewStamp instantiates a new Stamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStampWithDefaults

`func NewStampWithDefaults() *Stamp`

NewStampWithDefaults instantiates a new Stamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *Stamp) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Stamp) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Stamp) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *Stamp) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetDate

`func (o *Stamp) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Stamp) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Stamp) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Stamp) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSatCertNumber

`func (o *Stamp) GetSatCertNumber() string`

GetSatCertNumber returns the SatCertNumber field if non-nil, zero value otherwise.

### GetSatCertNumberOk

`func (o *Stamp) GetSatCertNumberOk() (*string, bool)`

GetSatCertNumberOk returns a tuple with the SatCertNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatCertNumber

`func (o *Stamp) SetSatCertNumber(v string)`

SetSatCertNumber sets SatCertNumber field to given value.

### HasSatCertNumber

`func (o *Stamp) HasSatCertNumber() bool`

HasSatCertNumber returns a boolean if a field has been set.

### GetSatSignature

`func (o *Stamp) GetSatSignature() string`

GetSatSignature returns the SatSignature field if non-nil, zero value otherwise.

### GetSatSignatureOk

`func (o *Stamp) GetSatSignatureOk() (*string, bool)`

GetSatSignatureOk returns a tuple with the SatSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatSignature

`func (o *Stamp) SetSatSignature(v string)`

SetSatSignature sets SatSignature field to given value.

### HasSatSignature

`func (o *Stamp) HasSatSignature() bool`

HasSatSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


