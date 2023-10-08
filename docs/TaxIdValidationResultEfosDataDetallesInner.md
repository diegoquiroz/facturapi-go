# TaxIdValidationResultEfosDataDetallesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rfc** | Pointer to **string** | El RFC consultado, a manera de confirmación. | [optional] 
**RazonSocial** | Pointer to **string** | Razón social del contribuyente. | [optional] 
**SituacionContribuyente** | Pointer to **string** | Texto que indica la situación actual. Consulta [esta tabla](#situación-del-contribuyente) para ver el detalle de los distintos valores.  | [optional] 
**NumFechaPresuncion** | Pointer to **string** | Texto con identificador y fecha del reporte de presunción. | [optional] 
**PubFechaSatPresuntos** | Pointer to **string** | Fecha de publicación de presunción. | [optional] 
**NumGlobalPresuncion** | Pointer to **string** | Texto con identificador y fecha de publicación en el listado global de presunción. | [optional] 
**PubFechaDofPresuntos** | Pointer to **string** | Fecha de publicación en el Diario Oficial de la Federación (DOF). | [optional] 
**PubSatDefinitivos** | Pointer to **string** | Identificador de la publicación de estado “Definitivo”. | [optional] 
**PubDofDefinitivos** | Pointer to **string** | Fecha de la publicación de estado “Definitivo” en el DOF. | [optional] 
**NumFechaSentFav** | Pointer to **string** | Texto con identificador y fecha de sentencia favorable. | [optional] 
**PubSatSentFav** | Pointer to **string** | Fecha de sentencia favorable | [optional] 

## Methods

### NewTaxIdValidationResultEfosDataDetallesInner

`func NewTaxIdValidationResultEfosDataDetallesInner() *TaxIdValidationResultEfosDataDetallesInner`

NewTaxIdValidationResultEfosDataDetallesInner instantiates a new TaxIdValidationResultEfosDataDetallesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIdValidationResultEfosDataDetallesInnerWithDefaults

`func NewTaxIdValidationResultEfosDataDetallesInnerWithDefaults() *TaxIdValidationResultEfosDataDetallesInner`

NewTaxIdValidationResultEfosDataDetallesInnerWithDefaults instantiates a new TaxIdValidationResultEfosDataDetallesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRfc

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetRfc() string`

GetRfc returns the Rfc field if non-nil, zero value otherwise.

### GetRfcOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetRfcOk() (*string, bool)`

GetRfcOk returns a tuple with the Rfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfc

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetRfc(v string)`

SetRfc sets Rfc field to given value.

### HasRfc

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasRfc() bool`

HasRfc returns a boolean if a field has been set.

### GetRazonSocial

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetRazonSocial() string`

GetRazonSocial returns the RazonSocial field if non-nil, zero value otherwise.

### GetRazonSocialOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetRazonSocialOk() (*string, bool)`

GetRazonSocialOk returns a tuple with the RazonSocial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazonSocial

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetRazonSocial(v string)`

SetRazonSocial sets RazonSocial field to given value.

### HasRazonSocial

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasRazonSocial() bool`

HasRazonSocial returns a boolean if a field has been set.

### GetSituacionContribuyente

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetSituacionContribuyente() string`

GetSituacionContribuyente returns the SituacionContribuyente field if non-nil, zero value otherwise.

### GetSituacionContribuyenteOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetSituacionContribuyenteOk() (*string, bool)`

GetSituacionContribuyenteOk returns a tuple with the SituacionContribuyente field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSituacionContribuyente

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetSituacionContribuyente(v string)`

SetSituacionContribuyente sets SituacionContribuyente field to given value.

### HasSituacionContribuyente

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasSituacionContribuyente() bool`

HasSituacionContribuyente returns a boolean if a field has been set.

### GetNumFechaPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumFechaPresuncion() string`

GetNumFechaPresuncion returns the NumFechaPresuncion field if non-nil, zero value otherwise.

### GetNumFechaPresuncionOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumFechaPresuncionOk() (*string, bool)`

GetNumFechaPresuncionOk returns a tuple with the NumFechaPresuncion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFechaPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetNumFechaPresuncion(v string)`

SetNumFechaPresuncion sets NumFechaPresuncion field to given value.

### HasNumFechaPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasNumFechaPresuncion() bool`

HasNumFechaPresuncion returns a boolean if a field has been set.

### GetPubFechaSatPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubFechaSatPresuntos() string`

GetPubFechaSatPresuntos returns the PubFechaSatPresuntos field if non-nil, zero value otherwise.

### GetPubFechaSatPresuntosOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubFechaSatPresuntosOk() (*string, bool)`

GetPubFechaSatPresuntosOk returns a tuple with the PubFechaSatPresuntos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubFechaSatPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetPubFechaSatPresuntos(v string)`

SetPubFechaSatPresuntos sets PubFechaSatPresuntos field to given value.

### HasPubFechaSatPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasPubFechaSatPresuntos() bool`

HasPubFechaSatPresuntos returns a boolean if a field has been set.

### GetNumGlobalPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumGlobalPresuncion() string`

GetNumGlobalPresuncion returns the NumGlobalPresuncion field if non-nil, zero value otherwise.

### GetNumGlobalPresuncionOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumGlobalPresuncionOk() (*string, bool)`

GetNumGlobalPresuncionOk returns a tuple with the NumGlobalPresuncion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGlobalPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetNumGlobalPresuncion(v string)`

SetNumGlobalPresuncion sets NumGlobalPresuncion field to given value.

### HasNumGlobalPresuncion

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasNumGlobalPresuncion() bool`

HasNumGlobalPresuncion returns a boolean if a field has been set.

### GetPubFechaDofPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubFechaDofPresuntos() string`

GetPubFechaDofPresuntos returns the PubFechaDofPresuntos field if non-nil, zero value otherwise.

### GetPubFechaDofPresuntosOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubFechaDofPresuntosOk() (*string, bool)`

GetPubFechaDofPresuntosOk returns a tuple with the PubFechaDofPresuntos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubFechaDofPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetPubFechaDofPresuntos(v string)`

SetPubFechaDofPresuntos sets PubFechaDofPresuntos field to given value.

### HasPubFechaDofPresuntos

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasPubFechaDofPresuntos() bool`

HasPubFechaDofPresuntos returns a boolean if a field has been set.

### GetPubSatDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubSatDefinitivos() string`

GetPubSatDefinitivos returns the PubSatDefinitivos field if non-nil, zero value otherwise.

### GetPubSatDefinitivosOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubSatDefinitivosOk() (*string, bool)`

GetPubSatDefinitivosOk returns a tuple with the PubSatDefinitivos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubSatDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetPubSatDefinitivos(v string)`

SetPubSatDefinitivos sets PubSatDefinitivos field to given value.

### HasPubSatDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasPubSatDefinitivos() bool`

HasPubSatDefinitivos returns a boolean if a field has been set.

### GetPubDofDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubDofDefinitivos() string`

GetPubDofDefinitivos returns the PubDofDefinitivos field if non-nil, zero value otherwise.

### GetPubDofDefinitivosOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubDofDefinitivosOk() (*string, bool)`

GetPubDofDefinitivosOk returns a tuple with the PubDofDefinitivos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubDofDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetPubDofDefinitivos(v string)`

SetPubDofDefinitivos sets PubDofDefinitivos field to given value.

### HasPubDofDefinitivos

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasPubDofDefinitivos() bool`

HasPubDofDefinitivos returns a boolean if a field has been set.

### GetNumFechaSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumFechaSentFav() string`

GetNumFechaSentFav returns the NumFechaSentFav field if non-nil, zero value otherwise.

### GetNumFechaSentFavOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetNumFechaSentFavOk() (*string, bool)`

GetNumFechaSentFavOk returns a tuple with the NumFechaSentFav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFechaSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetNumFechaSentFav(v string)`

SetNumFechaSentFav sets NumFechaSentFav field to given value.

### HasNumFechaSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasNumFechaSentFav() bool`

HasNumFechaSentFav returns a boolean if a field has been set.

### GetPubSatSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubSatSentFav() string`

GetPubSatSentFav returns the PubSatSentFav field if non-nil, zero value otherwise.

### GetPubSatSentFavOk

`func (o *TaxIdValidationResultEfosDataDetallesInner) GetPubSatSentFavOk() (*string, bool)`

GetPubSatSentFavOk returns a tuple with the PubSatSentFav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubSatSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) SetPubSatSentFav(v string)`

SetPubSatSentFav sets PubSatSentFav field to given value.

### HasPubSatSentFav

`func (o *TaxIdValidationResultEfosDataDetallesInner) HasPubSatSentFav() bool`

HasPubSatSentFav returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


