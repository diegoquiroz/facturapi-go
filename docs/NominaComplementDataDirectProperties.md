# NominaComplementDataDirectProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TipoNomina** | Pointer to **string** | Tipo de nómina. - &#x60;“O”&#x60; (Ordinaria): Cuando corresponde a un pago que se realiza de manera habitual, como sueldos. - &#x60;“E”&#x60; (Extraordinaria): Para pagos fuera de lo habitual, como liquidaciones, aguinaldos o bonos.  | [optional] [default to "O"]
**FechaPago** | Pointer to **string** | Fecha de pago de la nómina al trabajador. | [optional] [default to "now"]
**FechaInicialPago** | Pointer to **string** | Fecha inicial del periodo de pago. | [optional] 
**FechaFinalPago** | Pointer to **string** | Fecha final del periodo de pago. | [optional] 
**NumDiasPagados** | Pointer to **float32** | Número de días pagados. Puede ser entero o fracción. | [optional] 

## Methods

### NewNominaComplementDataDirectProperties

`func NewNominaComplementDataDirectProperties() *NominaComplementDataDirectProperties`

NewNominaComplementDataDirectProperties instantiates a new NominaComplementDataDirectProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaComplementDataDirectPropertiesWithDefaults

`func NewNominaComplementDataDirectPropertiesWithDefaults() *NominaComplementDataDirectProperties`

NewNominaComplementDataDirectPropertiesWithDefaults instantiates a new NominaComplementDataDirectProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTipoNomina

`func (o *NominaComplementDataDirectProperties) GetTipoNomina() string`

GetTipoNomina returns the TipoNomina field if non-nil, zero value otherwise.

### GetTipoNominaOk

`func (o *NominaComplementDataDirectProperties) GetTipoNominaOk() (*string, bool)`

GetTipoNominaOk returns a tuple with the TipoNomina field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoNomina

`func (o *NominaComplementDataDirectProperties) SetTipoNomina(v string)`

SetTipoNomina sets TipoNomina field to given value.

### HasTipoNomina

`func (o *NominaComplementDataDirectProperties) HasTipoNomina() bool`

HasTipoNomina returns a boolean if a field has been set.

### GetFechaPago

`func (o *NominaComplementDataDirectProperties) GetFechaPago() string`

GetFechaPago returns the FechaPago field if non-nil, zero value otherwise.

### GetFechaPagoOk

`func (o *NominaComplementDataDirectProperties) GetFechaPagoOk() (*string, bool)`

GetFechaPagoOk returns a tuple with the FechaPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaPago

`func (o *NominaComplementDataDirectProperties) SetFechaPago(v string)`

SetFechaPago sets FechaPago field to given value.

### HasFechaPago

`func (o *NominaComplementDataDirectProperties) HasFechaPago() bool`

HasFechaPago returns a boolean if a field has been set.

### GetFechaInicialPago

`func (o *NominaComplementDataDirectProperties) GetFechaInicialPago() string`

GetFechaInicialPago returns the FechaInicialPago field if non-nil, zero value otherwise.

### GetFechaInicialPagoOk

`func (o *NominaComplementDataDirectProperties) GetFechaInicialPagoOk() (*string, bool)`

GetFechaInicialPagoOk returns a tuple with the FechaInicialPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaInicialPago

`func (o *NominaComplementDataDirectProperties) SetFechaInicialPago(v string)`

SetFechaInicialPago sets FechaInicialPago field to given value.

### HasFechaInicialPago

`func (o *NominaComplementDataDirectProperties) HasFechaInicialPago() bool`

HasFechaInicialPago returns a boolean if a field has been set.

### GetFechaFinalPago

`func (o *NominaComplementDataDirectProperties) GetFechaFinalPago() string`

GetFechaFinalPago returns the FechaFinalPago field if non-nil, zero value otherwise.

### GetFechaFinalPagoOk

`func (o *NominaComplementDataDirectProperties) GetFechaFinalPagoOk() (*string, bool)`

GetFechaFinalPagoOk returns a tuple with the FechaFinalPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaFinalPago

`func (o *NominaComplementDataDirectProperties) SetFechaFinalPago(v string)`

SetFechaFinalPago sets FechaFinalPago field to given value.

### HasFechaFinalPago

`func (o *NominaComplementDataDirectProperties) HasFechaFinalPago() bool`

HasFechaFinalPago returns a boolean if a field has been set.

### GetNumDiasPagados

`func (o *NominaComplementDataDirectProperties) GetNumDiasPagados() float32`

GetNumDiasPagados returns the NumDiasPagados field if non-nil, zero value otherwise.

### GetNumDiasPagadosOk

`func (o *NominaComplementDataDirectProperties) GetNumDiasPagadosOk() (*float32, bool)`

GetNumDiasPagadosOk returns a tuple with the NumDiasPagados field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDiasPagados

`func (o *NominaComplementDataDirectProperties) SetNumDiasPagados(v float32)`

SetNumDiasPagados sets NumDiasPagados field to given value.

### HasNumDiasPagados

`func (o *NominaComplementDataDirectProperties) HasNumDiasPagados() bool`

HasNumDiasPagados returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


