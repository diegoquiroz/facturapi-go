# OtroPago

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoOtroPago** | **string** | Clave del catálogo [Tipo de Otro Pago](#tipo-de-otro-pago). | 
**Clave** | **string** | Clave de otro pago de nómina propia de la contabilidad de cada patrón. | 
**Concepto** | Pointer to **string** | Descripción alternativa correspondiente a la clave utilizada. | [optional] 
**Importe** | **float32** | Importe por concepto de otro pago. | 
**SubsidioCausado** | Pointer to **float32** | Subsudio causado conforme a la tabla del subsidio para el empleo publicada en el Anexo 8 de la Resolución Miscelánea Fiscal vigente.  Este valor será insertado dentro del nodo &#x60;SubsidioAlEmpleo&#x60;, y es requerido cuando el valor de &#x60;tipo_otro_pago&#x60; es &#x60;\&quot;002\&quot;&#x60;.  | [optional] 
**CompensacionSaldosAFavor** | Pointer to [**NominaCompensacionInput**](NominaCompensacionInput.md) |  | [optional] 

## Methods

### NewOtroPago

`func NewOtroPago(tipoOtroPago string, clave string, importe float32, ) *OtroPago`

NewOtroPago instantiates a new OtroPago object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtroPagoWithDefaults

`func NewOtroPagoWithDefaults() *OtroPago`

NewOtroPagoWithDefaults instantiates a new OtroPago object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoOtroPago

`func (o *OtroPago) GetTipoOtroPago() string`

GetTipoOtroPago returns the TipoOtroPago field if non-nil, zero value otherwise.

### GetTipoOtroPagoOk

`func (o *OtroPago) GetTipoOtroPagoOk() (*string, bool)`

GetTipoOtroPagoOk returns a tuple with the TipoOtroPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoOtroPago

`func (o *OtroPago) SetTipoOtroPago(v string)`

SetTipoOtroPago sets TipoOtroPago field to given value.


### GetClave

`func (o *OtroPago) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *OtroPago) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *OtroPago) SetClave(v string)`

SetClave sets Clave field to given value.


### GetConcepto

`func (o *OtroPago) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *OtroPago) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *OtroPago) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *OtroPago) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetImporte

`func (o *OtroPago) GetImporte() float32`

GetImporte returns the Importe field if non-nil, zero value otherwise.

### GetImporteOk

`func (o *OtroPago) GetImporteOk() (*float32, bool)`

GetImporteOk returns a tuple with the Importe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporte

`func (o *OtroPago) SetImporte(v float32)`

SetImporte sets Importe field to given value.


### GetSubsidioCausado

`func (o *OtroPago) GetSubsidioCausado() float32`

GetSubsidioCausado returns the SubsidioCausado field if non-nil, zero value otherwise.

### GetSubsidioCausadoOk

`func (o *OtroPago) GetSubsidioCausadoOk() (*float32, bool)`

GetSubsidioCausadoOk returns a tuple with the SubsidioCausado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidioCausado

`func (o *OtroPago) SetSubsidioCausado(v float32)`

SetSubsidioCausado sets SubsidioCausado field to given value.

### HasSubsidioCausado

`func (o *OtroPago) HasSubsidioCausado() bool`

HasSubsidioCausado returns a boolean if a field has been set.

### GetCompensacionSaldosAFavor

`func (o *OtroPago) GetCompensacionSaldosAFavor() NominaCompensacionInput`

GetCompensacionSaldosAFavor returns the CompensacionSaldosAFavor field if non-nil, zero value otherwise.

### GetCompensacionSaldosAFavorOk

`func (o *OtroPago) GetCompensacionSaldosAFavorOk() (*NominaCompensacionInput, bool)`

GetCompensacionSaldosAFavorOk returns a tuple with the CompensacionSaldosAFavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensacionSaldosAFavor

`func (o *OtroPago) SetCompensacionSaldosAFavor(v NominaCompensacionInput)`

SetCompensacionSaldosAFavor sets CompensacionSaldosAFavor field to given value.

### HasCompensacionSaldosAFavor

`func (o *OtroPago) HasCompensacionSaldosAFavor() bool`

HasCompensacionSaldosAFavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


