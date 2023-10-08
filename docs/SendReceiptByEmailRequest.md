# SendReceiptByEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | Dirección de correo eletrónico | 

## Methods

### NewSendReceiptByEmailRequest

`func NewSendReceiptByEmailRequest(email NullableString, ) *SendReceiptByEmailRequest`

NewSendReceiptByEmailRequest instantiates a new SendReceiptByEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendReceiptByEmailRequestWithDefaults

`func NewSendReceiptByEmailRequestWithDefaults() *SendReceiptByEmailRequest`

NewSendReceiptByEmailRequestWithDefaults instantiates a new SendReceiptByEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendReceiptByEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendReceiptByEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendReceiptByEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *SendReceiptByEmailRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SendReceiptByEmailRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


