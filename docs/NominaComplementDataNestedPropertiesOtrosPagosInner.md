# NominaComplementDataNestedPropertiesOtrosPagosInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoOtroPago** | Pointer to **string** | Clave del catálogo [Tipo de Otro Pago](#tipo-de-otro-pago). | [optional] 
**Clave** | Pointer to **string** | Clave de otro pago de nómina propia de la contabilidad de cada patrón. | [optional] 
**Concepto** | Pointer to **string** | Descripción alternativa correspondiente a la clave utilizada. | [optional] 
**Importe** | Pointer to **float32** | Importe por concepto de otro pago. | [optional] 
**SubsidioCausado** | Pointer to **float32** | Subsudio causado conforme a la tabla del subsidio para el empleo publicada en el Anexo 8 de la Resolución Miscelánea Fiscal vigente.  Este valor será insertado dentro del nodo &#x60;SubsidioAlEmpleo&#x60;, y es requerido cuando el valor de &#x60;tipo_otro_pago&#x60; es &#x60;\&quot;002\&quot;&#x60;.  | [optional] 
**CompensacionSaldosAFavor** | Pointer to [**NominaCompensacionProperties**](NominaCompensacionProperties.md) |  | [optional] 

## Methods

### NewNominaComplementDataNestedPropertiesOtrosPagosInner

`func NewNominaComplementDataNestedPropertiesOtrosPagosInner() *NominaComplementDataNestedPropertiesOtrosPagosInner`

NewNominaComplementDataNestedPropertiesOtrosPagosInner instantiates a new NominaComplementDataNestedPropertiesOtrosPagosInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataNestedPropertiesOtrosPagosInnerWithDefaults

`func NewNominaComplementDataNestedPropertiesOtrosPagosInnerWithDefaults() *NominaComplementDataNestedPropertiesOtrosPagosInner`

NewNominaComplementDataNestedPropertiesOtrosPagosInnerWithDefaults instantiates a new NominaComplementDataNestedPropertiesOtrosPagosInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoOtroPago

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetTipoOtroPago() string`

GetTipoOtroPago returns the TipoOtroPago field if non-nil, zero value otherwise.

### GetTipoOtroPagoOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetTipoOtroPagoOk() (*string, bool)`

GetTipoOtroPagoOk returns a tuple with the TipoOtroPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoOtroPago

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetTipoOtroPago(v string)`

SetTipoOtroPago sets TipoOtroPago field to given value.

### HasTipoOtroPago

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasTipoOtroPago() bool`

HasTipoOtroPago returns a boolean if a field has been set.

### GetClave

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetClave() string`

GetClave returns the Clave field if non-nil, zero value otherwise.

### GetClaveOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetClaveOk() (*string, bool)`

GetClaveOk returns a tuple with the Clave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClave

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetClave(v string)`

SetClave sets Clave field to given value.

### HasClave

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasClave() bool`

HasClave returns a boolean if a field has been set.

### GetConcepto

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetConcepto() string`

GetConcepto returns the Concepto field if non-nil, zero value otherwise.

### GetConceptoOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetConceptoOk() (*string, bool)`

GetConceptoOk returns a tuple with the Concepto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcepto

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetConcepto(v string)`

SetConcepto sets Concepto field to given value.

### HasConcepto

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasConcepto() bool`

HasConcepto returns a boolean if a field has been set.

### GetImporte

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetImporte() float32`

GetImporte returns the Importe field if non-nil, zero value otherwise.

### GetImporteOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetImporteOk() (*float32, bool)`

GetImporteOk returns a tuple with the Importe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporte

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetImporte(v float32)`

SetImporte sets Importe field to given value.

### HasImporte

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasImporte() bool`

HasImporte returns a boolean if a field has been set.

### GetSubsidioCausado

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetSubsidioCausado() float32`

GetSubsidioCausado returns the SubsidioCausado field if non-nil, zero value otherwise.

### GetSubsidioCausadoOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetSubsidioCausadoOk() (*float32, bool)`

GetSubsidioCausadoOk returns a tuple with the SubsidioCausado field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidioCausado

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetSubsidioCausado(v float32)`

SetSubsidioCausado sets SubsidioCausado field to given value.

### HasSubsidioCausado

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasSubsidioCausado() bool`

HasSubsidioCausado returns a boolean if a field has been set.

### GetCompensacionSaldosAFavor

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetCompensacionSaldosAFavor() NominaCompensacionProperties`

GetCompensacionSaldosAFavor returns the CompensacionSaldosAFavor field if non-nil, zero value otherwise.

### GetCompensacionSaldosAFavorOk

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) GetCompensacionSaldosAFavorOk() (*NominaCompensacionProperties, bool)`

GetCompensacionSaldosAFavorOk returns a tuple with the CompensacionSaldosAFavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensacionSaldosAFavor

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) SetCompensacionSaldosAFavor(v NominaCompensacionProperties)`

SetCompensacionSaldosAFavor sets CompensacionSaldosAFavor field to given value.

### HasCompensacionSaldosAFavor

`func (o *NominaComplementDataNestedPropertiesOtrosPagosInner) HasCompensacionSaldosAFavor() bool`

HasCompensacionSaldosAFavor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


