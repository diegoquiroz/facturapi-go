# NominaDeduccionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoDeduccion** | **string** | Clave del catálogo [Tipo de deducción](#tipo-de-deducción). | 
**Concepto** | Pointer to **string** | Concepto de la deducción. Si no se envía, se utilizará la descripción del catálogo del tipo de deducción. | [optional] 
**Clave** | **string** | Clave de control interno que asigna el patrón a cada deducción (descuento) de nómina propia de su contabilidad. | 
**Importe** | **float32** | Importe del concepto de deducción. | 

## Methods

### NewNominaDeduccionInput

`func NewNominaDeduccionInput(tipoDeduccion string, clave string, importe float32, ) *NominaDeduccionInput`

NewNominaDeduccionInput instantiates a new NominaDeduccionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaDeduccionInputWithDefaults

`func NewNominaDeduccionInputWithDefaults() *NominaDeduccionInput`

NewNominaDeduccionInputWithDefaults instantiates a new NominaDeduccionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoDeduccion

`func (o *NominaDeduccionInput) GetTipoDeduccion() string`

GetTipoDeduccion returns the TipoDeduccion field if non-nil, zero value otherwise.

### GetTipoDeduccionOk

`func (o *NominaDeduccionInput) GetTipoDeduccionOk() (*string, bool)`

GetTipoDeduccionOk returns a tuple with the TipoDeduccion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoDeduccion

`func (o *NominaDeduccionInput) SetTipoDeduccion(v string)`

SetTipoDeduccion sets TipoDeduccion field to given value.


### GetConcepto

`func (o *NominaDeduccionInput) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *NominaDeduccionInput) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *NominaDeduccionInput) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *NominaDeduccionInput) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetClave

`func (o *NominaDeduccionInput) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *NominaDeduccionInput) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *NominaDeduccionInput) SetClave(v string)`

SetClave sets Clave field to given value.


### GetImporte

`func (o *NominaDeduccionInput) GetImporte() float32`

GetImporte returns the Importe field if non-nil, zero value otherwise.

### GetImporteOk

`func (o *NominaDeduccionInput) GetImporteOk() (*float32, bool)`

GetImporteOk returns a tuple with the Importe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporte

`func (o *NominaDeduccionInput) SetImporte(v float32)`

SetImporte sets Importe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


