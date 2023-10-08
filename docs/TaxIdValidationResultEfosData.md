# TaxIdValidationResultEfosData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mensaje** | Pointer to **string** | Disponible sólo cuando el RFC no fue encontrado en la lista, lo cual es bueno.  | [optional] 
**FechaLista** | Pointer to **string** | Texto que indica la fecha de actualización de la lista. | [optional] 
**Detalles** | Pointer to [**[]TaxIdValidationResultEfosDataDetallesInner**](TaxIdValidationResultEfosDataDetallesInner.md) |  | [optional] 

## Methods

### NewTaxIdValidationResultEfosData

`func NewTaxIdValidationResultEfosData() *TaxIdValidationResultEfosData`

NewTaxIdValidationResultEfosData instantiates a new TaxIdValidationResultEfosData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIdValidationResultEfosDataWithDefaults

`func NewTaxIdValidationResultEfosDataWithDefaults() *TaxIdValidationResultEfosData`

NewTaxIdValidationResultEfosDataWithDefaults instantiates a new TaxIdValidationResultEfosData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMensaje

`func (o *TaxIdValidationResultEfosData) GetMensaje() string`

GetMensaje returns the Mensaje field if non-nil, zero value otherwise.

### GetMensajeOk

`func (o *TaxIdValidationResultEfosData) GetMensajeOk() (*string, bool)`

GetMensajeOk returns a tuple with the Mensaje field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMensaje

`func (o *TaxIdValidationResultEfosData) SetMensaje(v string)`

SetMensaje sets Mensaje field to given value.

### HasMensaje

`func (o *TaxIdValidationResultEfosData) HasMensaje() bool`

HasMensaje returns a boolean if a field has been set.

### GetFechaLista

`func (o *TaxIdValidationResultEfosData) GetFechaLista() string`

GetFechaLista returns the FechaLista field if non-nil, zero value otherwise.

### GetFechaListaOk

`func (o *TaxIdValidationResultEfosData) GetFechaListaOk() (*string, bool)`

GetFechaListaOk returns a tuple with the FechaLista field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaLista

`func (o *TaxIdValidationResultEfosData) SetFechaLista(v string)`

SetFechaLista sets FechaLista field to given value.

### HasFechaLista

`func (o *TaxIdValidationResultEfosData) HasFechaLista() bool`

HasFechaLista returns a boolean if a field has been set.

### GetDetalles

`func (o *TaxIdValidationResultEfosData) GetDetalles() []TaxIdValidationResultEfosDataDetallesInner`

GetDetalles returns the Detalles field if non-nil, zero value otherwise.

### GetDetallesOk

`func (o *TaxIdValidationResultEfosData) GetDetallesOk() (*[]TaxIdValidationResultEfosDataDetallesInner, bool)`

GetDetallesOk returns a tuple with the Detalles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetalles

`func (o *TaxIdValidationResultEfosData) SetDetalles(v []TaxIdValidationResultEfosDataDetallesInner)`

SetDetalles sets Detalles field to given value.

### HasDetalles

`func (o *TaxIdValidationResultEfosData) HasDetalles() bool`

HasDetalles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


