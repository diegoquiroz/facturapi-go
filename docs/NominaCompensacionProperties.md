# NominaCompensacionProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaldoAFavor** | Pointer to **float32** | Monto por saldo a favor determinado por el patrón al trabajador en periodos o ejercicios anteriores. | [optional] 
**Ano** | Pointer to **int32** | Año en que se determinó el saldo a favor del trabajador. | [optional] 
**RemanenteSalFav** | Pointer to **float32** | Remanente del saldo a favor del trabajador. | [optional] 

## Methods

### NewNominaCompensacionProperties

`func NewNominaCompensacionProperties() *NominaCompensacionProperties`

NewNominaCompensacionProperties instantiates a new NominaCompensacionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaCompensacionPropertiesWithDefaults

`func NewNominaCompensacionPropertiesWithDefaults() *NominaCompensacionProperties`

NewNominaCompensacionPropertiesWithDefaults instantiates a new NominaCompensacionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaldoAFavor

`func (o *NominaCompensacionProperties) GetSaldoAFavor() float32`

GetSaldoAFavor returns the SaldoAFavor field if non-nil, zero value otherwise.

### GetSaldoAFavorOk

`func (o *NominaCompensacionProperties) GetSaldoAFavorOk() (*float32, bool)`

GetSaldoAFavorOk returns a tuple with the SaldoAFavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaldoAFavor

`func (o *NominaCompensacionProperties) SetSaldoAFavor(v float32)`

SetSaldoAFavor sets SaldoAFavor field to given value.

### HasSaldoAFavor

`func (o *NominaCompensacionProperties) HasSaldoAFavor() bool`

HasSaldoAFavor returns a boolean if a field has been set.

### GetAno

`func (o *NominaCompensacionProperties) GetAno() int32`

GetAno returns the Ano field if non-nil, zero value otherwise.

### GetAnoOk

`func (o *NominaCompensacionProperties) GetAnoOk() (*int32, bool)`

GetAnoOk returns a tuple with the Ano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAno

`func (o *NominaCompensacionProperties) SetAno(v int32)`

SetAno sets Ano field to given value.

### HasAno

`func (o *NominaCompensacionProperties) HasAno() bool`

HasAno returns a boolean if a field has been set.

### GetRemanenteSalFav

`func (o *NominaCompensacionProperties) GetRemanenteSalFav() float32`

GetRemanenteSalFav returns the RemanenteSalFav field if non-nil, zero value otherwise.

### GetRemanenteSalFavOk

`func (o *NominaCompensacionProperties) GetRemanenteSalFavOk() (*float32, bool)`

GetRemanenteSalFavOk returns a tuple with the RemanenteSalFav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanenteSalFav

`func (o *NominaCompensacionProperties) SetRemanenteSalFav(v float32)`

SetRemanenteSalFav sets RemanenteSalFav field to given value.

### HasRemanenteSalFav

`func (o *NominaCompensacionProperties) HasRemanenteSalFav() bool`

HasRemanenteSalFav returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


