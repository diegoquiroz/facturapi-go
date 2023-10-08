# CustomComplementProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Tipo de complemento. | [optional] 
**Data** | Pointer to **string** | Código XML de tu complemento tal cual como quieres que se inserte en el XML. Debe contener sólamente un nodo XML raíz. | [optional] 

## Methods

### NewCustomComplementProperties

`func NewCustomComplementProperties() *CustomComplementProperties`

NewCustomComplementProperties instantiates a new CustomComplementProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComplementPropertiesWithDefaults

`func NewCustomComplementPropertiesWithDefaults() *CustomComplementProperties`

NewCustomComplementPropertiesWithDefaults instantiates a new CustomComplementProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomComplementProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomComplementProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomComplementProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomComplementProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *CustomComplementProperties) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComplementProperties) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComplementProperties) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CustomComplementProperties) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


