# NominaPercepcionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoPercepcion** | **string** | Clave del catálogo [Tipo de percepción](#tipo-de-percepcion). | 
**Concepto** | Pointer to **string** | Concepto de la percepción. Si no se envía, se utilizará la descripción del catálogo del tipo de percepción. | [optional] 
**Clave** | **string** | Clave de control interno que asigna el patrón a cada percepción de nómina propia de su contabilidad. | 
**ImporteGravado** | **float32** | Importe gravado por el concepto indicado en el tipo de percepción. | 
**ImporteExento** | **float32** | Importe exento por el concepto indicado en el tipo de percepción. | 
**AccionesOTitulos** | Pointer to [**NominaAccionesInput**](NominaAccionesInput.md) |  | [optional] 
**HorasExtra** | Pointer to [**[]NominaHorasExtraInput**](NominaHorasExtraInput.md) |  | [optional] 

## Methods

### NewNominaPercepcionInput

`func NewNominaPercepcionInput(tipoPercepcion string, clave string, importeGravado float32, importeExento float32, ) *NominaPercepcionInput`

NewNominaPercepcionInput instantiates a new NominaPercepcionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaPercepcionInputWithDefaults

`func NewNominaPercepcionInputWithDefaults() *NominaPercepcionInput`

NewNominaPercepcionInputWithDefaults instantiates a new NominaPercepcionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoPercepcion

`func (o *NominaPercepcionInput) GetTipoPercepcion() string`

GetTipoPercepcion returns the TipoPercepcion field if non-nil, zero value otherwise.

### GetTipoPercepcionOk

`func (o *NominaPercepcionInput) GetTipoPercepcionOk() (*string, bool)`

GetTipoPercepcionOk returns a tuple with the TipoPercepcion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoPercepcion

`func (o *NominaPercepcionInput) SetTipoPercepcion(v string)`

SetTipoPercepcion sets TipoPercepcion field to given value.


### GetConcepto

`func (o *NominaPercepcionInput) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *NominaPercepcionInput) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *NominaPercepcionInput) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *NominaPercepcionInput) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetClave

`func (o *NominaPercepcionInput) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *NominaPercepcionInput) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *NominaPercepcionInput) SetClave(v string)`

SetClave sets Clave field to given value.


### GetImporteGravado

`func (o *NominaPercepcionInput) GetImporteGravado() float32`

GetImporteGravado returns the ImporteGravado field if non-nil, zero value otherwise.

### GetImporteGravadoOk

`func (o *NominaPercepcionInput) GetImporteGravadoOk() (*float32, bool)`

GetImporteGravadoOk returns a tuple with the ImporteGravado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporteGravado

`func (o *NominaPercepcionInput) SetImporteGravado(v float32)`

SetImporteGravado sets ImporteGravado field to given value.


### GetImporteExento

`func (o *NominaPercepcionInput) GetImporteExento() float32`

GetImporteExento returns the ImporteExento field if non-nil, zero value otherwise.

### GetImporteExentoOk

`func (o *NominaPercepcionInput) GetImporteExentoOk() (*float32, bool)`

GetImporteExentoOk returns a tuple with the ImporteExento field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporteExento

`func (o *NominaPercepcionInput) SetImporteExento(v float32)`

SetImporteExento sets ImporteExento field to given value.


### GetAccionesOTitulos

`func (o *NominaPercepcionInput) GetAccionesOTitulos() NominaAccionesInput`

GetAccionesOTitulos returns the AccionesOTitulos field if non-nil, zero value otherwise.

### GetAccionesOTitulosOk

`func (o *NominaPercepcionInput) GetAccionesOTitulosOk() (*NominaAccionesInput, bool)`

GetAccionesOTitulosOk returns a tuple with the AccionesOTitulos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccionesOTitulos

`func (o *NominaPercepcionInput) SetAccionesOTitulos(v NominaAccionesInput)`

SetAccionesOTitulos sets AccionesOTitulos field to given value.

### HasAccionesOTitulos

`func (o *NominaPercepcionInput) HasAccionesOTitulos() bool`

HasAccionesOTitulos returns a boolean if a field has been set.

### GetHorasExtra

`func (o *NominaPercepcionInput) GetHorasExtra() []NominaHorasExtraInput`

GetHorasExtra returns the HorasExtra field if non-nil, zero value otherwise.

### GetHorasExtraOk

`func (o *NominaPercepcionInput) GetHorasExtraOk() (*[]NominaHorasExtraInput, bool)`

GetHorasExtraOk returns a tuple with the HorasExtra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorasExtra

`func (o *NominaPercepcionInput) SetHorasExtra(v []NominaHorasExtraInput)`

SetHorasExtra sets HorasExtra field to given value.

### HasHorasExtra

`func (o *NominaPercepcionInput) HasHorasExtra() bool`

HasHorasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


