# IepsTax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IepsMode** | Pointer to **string** | Indica la manera de cobrar el impuesto, y puede tener los valores:  &#x60;\&quot;sum_before_taxes\&quot;&#x60;: Aplica primero el IEPS al subtotal y usa el resultado como base del resto de impuestos en el producto.  &#x60;\&quot;break_down\&quot;&#x60;: Cobra y desgloza el IEPS al mismo nivel que el resto de los impuestos en el producto.  &#x60;\&quot;unit\&quot;&#x60;: Aplica el IEPS antes del precio unitario, y usa el precio unitario original como base para el resto de impuestos.  Consulta con tu contador qué caso aplica para tu giro de empresa y producto.  | [optional] [default to "sum_before_taxes"]

## Methods

### NewIepsTax

`func NewIepsTax() *IepsTax`

NewIepsTax instantiates a new IepsTax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIepsTaxWithDefaults

`func NewIepsTaxWithDefaults() *IepsTax`

NewIepsTaxWithDefaults instantiates a new IepsTax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIepsMode

`func (o *IepsTax) GetIepsMode() string`

GetIepsMode returns the IepsMode field if non-nil, zero value otherwise.

### GetIepsModeOk

`func (o *IepsTax) GetIepsModeOk() (*string, bool)`

GetIepsModeOk returns a tuple with the IepsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIepsMode

`func (o *IepsTax) SetIepsMode(v string)`

SetIepsMode sets IepsMode field to given value.

### HasIepsMode

`func (o *IepsTax) HasIepsMode() bool`

HasIepsMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


