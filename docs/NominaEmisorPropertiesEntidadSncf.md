# NominaEmisorPropertiesEntidadSncf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrigenRecurso** | Pointer to **string** | Clave de origen de recurso.  - &#x60;“IP”&#x60;: Ingresos Propios - &#x60;“IF”&#x60;: Ingresos Federales - &#x60;“IM”&#x60;: Ingresos mixtos.  | [optional] 
**MontoRecursoPropio** | Pointer to **float32** | Inporte de recursos propios. Requerido cuando el origen del recurso es por ingresos mixtos.  | [optional] 

## Methods

### NewNominaEmisorPropertiesEntidadSncf

`func NewNominaEmisorPropertiesEntidadSncf() *NominaEmisorPropertiesEntidadSncf`

NewNominaEmisorPropertiesEntidadSncf instantiates a new NominaEmisorPropertiesEntidadSncf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaEmisorPropertiesEntidadSncfWithDefaults

`func NewNominaEmisorPropertiesEntidadSncfWithDefaults() *NominaEmisorPropertiesEntidadSncf`

NewNominaEmisorPropertiesEntidadSncfWithDefaults instantiates a new NominaEmisorPropertiesEntidadSncf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigenRecurso

`func (o *NominaEmisorPropertiesEntidadSncf) GetOrigenRecurso() string`

GetOrigenRecurso returns the OrigenRecurso field if non-nil, zero value otherwise.

### GetOrigenRecursoOk

`func (o *NominaEmisorPropertiesEntidadSncf) GetOrigenRecursoOk() (*string, bool)`

GetOrigenRecursoOk returns a tuple with the OrigenRecurso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigenRecurso

`func (o *NominaEmisorPropertiesEntidadSncf) SetOrigenRecurso(v string)`

SetOrigenRecurso sets OrigenRecurso field to given value.

### HasOrigenRecurso

`func (o *NominaEmisorPropertiesEntidadSncf) HasOrigenRecurso() bool`

HasOrigenRecurso returns a boolean if a field has been set.

### GetMontoRecursoPropio

`func (o *NominaEmisorPropertiesEntidadSncf) GetMontoRecursoPropio() float32`

GetMontoRecursoPropio returns the MontoRecursoPropio field if non-nil, zero value otherwise.

### GetMontoRecursoPropioOk

`func (o *NominaEmisorPropertiesEntidadSncf) GetMontoRecursoPropioOk() (*float32, bool)`

GetMontoRecursoPropioOk returns a tuple with the MontoRecursoPropio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMontoRecursoPropio

`func (o *NominaEmisorPropertiesEntidadSncf) SetMontoRecursoPropio(v float32)`

SetMontoRecursoPropio sets MontoRecursoPropio field to given value.

### HasMontoRecursoPropio

`func (o *NominaEmisorPropertiesEntidadSncf) HasMontoRecursoPropio() bool`

HasMontoRecursoPropio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


