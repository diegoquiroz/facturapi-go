# SendInvoiceByEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Dirección de correo eletrónico | [optional] 

## Methods

### NewSendInvoiceByEmailRequest

`func NewSendInvoiceByEmailRequest() *SendInvoiceByEmailRequest`

NewSendInvoiceByEmailRequest instantiates a new SendInvoiceByEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInvoiceByEmailRequestWithDefaults

`func NewSendInvoiceByEmailRequestWithDefaults() *SendInvoiceByEmailRequest`

NewSendInvoiceByEmailRequestWithDefaults instantiates a new SendInvoiceByEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendInvoiceByEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendInvoiceByEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendInvoiceByEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SendInvoiceByEmailRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SendInvoiceByEmailRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SendInvoiceByEmailRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


