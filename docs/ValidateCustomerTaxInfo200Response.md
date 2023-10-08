# ValidateCustomerTaxInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | Pointer to **bool** | Indica si la información fiscal del cliente coincide con los registros del SAT | [optional] 
**Errors** | Pointer to [**[]ValidateCustomerTaxInfo200ResponseErrorsInner**](ValidateCustomerTaxInfo200ResponseErrorsInner.md) |  | [optional] 

## Methods

### NewValidateCustomerTaxInfo200Response

`func NewValidateCustomerTaxInfo200Response() *ValidateCustomerTaxInfo200Response`

NewValidateCustomerTaxInfo200Response instantiates a new ValidateCustomerTaxInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateCustomerTaxInfo200ResponseWithDefaults

`func NewValidateCustomerTaxInfo200ResponseWithDefaults() *ValidateCustomerTaxInfo200Response`

NewValidateCustomerTaxInfo200ResponseWithDefaults instantiates a new ValidateCustomerTaxInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsValid

`func (o *ValidateCustomerTaxInfo200Response) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ValidateCustomerTaxInfo200Response) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ValidateCustomerTaxInfo200Response) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *ValidateCustomerTaxInfo200Response) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetErrors

`func (o *ValidateCustomerTaxInfo200Response) GetErrors() []ValidateCustomerTaxInfo200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidateCustomerTaxInfo200Response) GetErrorsOk() (*[]ValidateCustomerTaxInfo200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidateCustomerTaxInfo200Response) SetErrors(v []ValidateCustomerTaxInfo200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidateCustomerTaxInfo200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


