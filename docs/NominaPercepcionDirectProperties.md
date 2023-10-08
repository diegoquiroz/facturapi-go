# NominaPercepcionDirectProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoPercepcion** | Pointer to **string** | Clave del catálogo [Tipo de percepción](#tipo-de-percepcion). | [optional] 
**Concepto** | Pointer to **string** | Concepto de la percepción. Si no se envía, se utilizará la descripción del catálogo del tipo de percepción. | [optional] 
**Clave** | Pointer to **string** | Clave de control interno que asigna el patrón a cada percepción de nómina propia de su contabilidad. | [optional] 
**ImporteGravado** | Pointer to **float32** | Importe gravado por el concepto indicado en el tipo de percepción. | [optional] 
**ImporteExento** | Pointer to **float32** | Importe exento por el concepto indicado en el tipo de percepción. | [optional] 

## Methods

### NewNominaPercepcionDirectProperties

`func NewNominaPercepcionDirectProperties() *NominaPercepcionDirectProperties`

NewNominaPercepcionDirectProperties instantiates a new NominaPercepcionDirectProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaPercepcionDirectPropertiesWithDefaults

`func NewNominaPercepcionDirectPropertiesWithDefaults() *NominaPercepcionDirectProperties`

NewNominaPercepcionDirectPropertiesWithDefaults instantiates a new NominaPercepcionDirectProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoPercepcion

`func (o *NominaPercepcionDirectProperties) GetTipoPercepcion() string`

GetTipoPercepcion returns the TipoPercepcion field if non-nil, zero value otherwise.

### GetTipoPercepcionOk

`func (o *NominaPercepcionDirectProperties) GetTipoPercepcionOk() (*string, bool)`

GetTipoPercepcionOk returns a tuple with the TipoPercepcion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoPercepcion

`func (o *NominaPercepcionDirectProperties) SetTipoPercepcion(v string)`

SetTipoPercepcion sets TipoPercepcion field to given value.

### HasTipoPercepcion

`func (o *NominaPercepcionDirectProperties) HasTipoPercepcion() bool`

HasTipoPercepcion returns a boolean if a field has been set.

### GetConcepto

`func (o *NominaPercepcionDirectProperties) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *NominaPercepcionDirectProperties) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *NominaPercepcionDirectProperties) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *NominaPercepcionDirectProperties) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetClave

`func (o *NominaPercepcionDirectProperties) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *NominaPercepcionDirectProperties) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *NominaPercepcionDirectProperties) SetClave(v string)`

SetClave sets Clave field to given value.

### HasClave

`func (o *NominaPercepcionDirectProperties) HasClave() bool`

HasClave returns a boolean if a field has been set.

### GetImporteGravado

`func (o *NominaPercepcionDirectProperties) GetImporteGravado() float32`

GetImporteGravado returns the ImporteGravado field if non-nil, zero value otherwise.

### GetImporteGravadoOk

`func (o *NominaPercepcionDirectProperties) GetImporteGravadoOk() (*float32, bool)`

GetImporteGravadoOk returns a tuple with the ImporteGravado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporteGravado

`func (o *NominaPercepcionDirectProperties) SetImporteGravado(v float32)`

SetImporteGravado sets ImporteGravado field to given value.

### HasImporteGravado

`func (o *NominaPercepcionDirectProperties) HasImporteGravado() bool`

HasImporteGravado returns a boolean if a field has been set.

### GetImporteExento

`func (o *NominaPercepcionDirectProperties) GetImporteExento() float32`

GetImporteExento returns the ImporteExento field if non-nil, zero value otherwise.

### GetImporteExentoOk

`func (o *NominaPercepcionDirectProperties) GetImporteExentoOk() (*float32, bool)`

GetImporteExentoOk returns a tuple with the ImporteExento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporteExento

`func (o *NominaPercepcionDirectProperties) SetImporteExento(v float32)`

SetImporteExento sets ImporteExento field to given value.

### HasImporteExento

`func (o *NominaPercepcionDirectProperties) HasImporteExento() bool`

HasImporteExento returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


