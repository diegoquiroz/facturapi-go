# NominaDeduccionProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoDeduccion** | Pointer to **string** | Clave del catálogo [Tipo de deducción](#tipo-de-deducción). | [optional] 
**Concepto** | Pointer to **string** | Concepto de la deducción. Si no se envía, se utilizará la descripción del catálogo del tipo de deducción. | [optional] 
**Clave** | Pointer to **string** | Clave de control interno que asigna el patrón a cada deducción (descuento) de nómina propia de su contabilidad. | [optional] 
**Importe** | Pointer to **float32** | Importe del concepto de deducción. | [optional] 

## Methods

### NewNominaDeduccionProperties

`func NewNominaDeduccionProperties() *NominaDeduccionProperties`

NewNominaDeduccionProperties instantiates a new NominaDeduccionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaDeduccionPropertiesWithDefaults

`func NewNominaDeduccionPropertiesWithDefaults() *NominaDeduccionProperties`

NewNominaDeduccionPropertiesWithDefaults instantiates a new NominaDeduccionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoDeduccion

`func (o *NominaDeduccionProperties) GetTipoDeduccion() string`

GetTipoDeduccion returns the TipoDeduccion field if non-nil, zero value otherwise.

### GetTipoDeduccionOk

`func (o *NominaDeduccionProperties) GetTipoDeduccionOk() (*string, bool)`

GetTipoDeduccionOk returns a tuple with the TipoDeduccion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoDeduccion

`func (o *NominaDeduccionProperties) SetTipoDeduccion(v string)`

SetTipoDeduccion sets TipoDeduccion field to given value.

### HasTipoDeduccion

`func (o *NominaDeduccionProperties) HasTipoDeduccion() bool`

HasTipoDeduccion returns a boolean if a field has been set.

### GetConcepto

`func (o *NominaDeduccionProperties) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *NominaDeduccionProperties) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *NominaDeduccionProperties) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *NominaDeduccionProperties) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetClave

`func (o *NominaDeduccionProperties) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *NominaDeduccionProperties) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *NominaDeduccionProperties) SetClave(v string)`

SetClave sets Clave field to given value.

### HasClave

`func (o *NominaDeduccionProperties) HasClave() bool`

HasClave returns a boolean if a field has been set.

### GetImporte

`func (o *NominaDeduccionProperties) GetImporte() float32`

GetImporte returns the Importe field if non-nil, zero value otherwise.

### GetImporteOk

`func (o *NominaDeduccionProperties) GetImporteOk() (*float32, bool)`

GetImporteOk returns a tuple with the Importe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporte

`func (o *NominaDeduccionProperties) SetImporte(v float32)`

SetImporte sets Importe field to given value.

### HasImporte

`func (o *NominaDeduccionProperties) HasImporte() bool`

HasImporte returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


