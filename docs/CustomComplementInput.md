# CustomComplementInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Tipo de complemento. | 
**Data** | **string** | Código XML de tu complemento tal cual como quieres que se inserte en el XML. Debe contener sólamente un nodo XML raíz. | 

## Methods

### NewCustomComplementInput

`func NewCustomComplementInput(type_ string, data string, ) *CustomComplementInput`

NewCustomComplementInput instantiates a new CustomComplementInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComplementInputWithDefaults

`func NewCustomComplementInputWithDefaults() *CustomComplementInput`

NewCustomComplementInputWithDefaults instantiates a new CustomComplementInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomComplementInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomComplementInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomComplementInput) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *CustomComplementInput) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComplementInput) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComplementInput) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


