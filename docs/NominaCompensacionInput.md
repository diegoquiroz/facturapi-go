# NominaCompensacionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaldoAFavor** | **float32** | Monto por saldo a favor determinado por el patrón al trabajador en periodos o ejercicios anteriores. | 
**Ano** | **int32** | Año en que se determinó el saldo a favor del trabajador. | 
**RemanenteSalFav** | **float32** | Remanente del saldo a favor del trabajador. | 

## Methods

### NewNominaCompensacionInput

`func NewNominaCompensacionInput(saldoAFavor float32, ano int32, remanenteSalFav float32, ) *NominaCompensacionInput`

NewNominaCompensacionInput instantiates a new NominaCompensacionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaCompensacionInputWithDefaults

`func NewNominaCompensacionInputWithDefaults() *NominaCompensacionInput`

NewNominaCompensacionInputWithDefaults instantiates a new NominaCompensacionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaldoAFavor

`func (o *NominaCompensacionInput) GetSaldoAFavor() float32`

GetSaldoAFavor returns the SaldoAFavor field if non-nil, zero value otherwise.

### GetSaldoAFavorOk

`func (o *NominaCompensacionInput) GetSaldoAFavorOk() (*float32, bool)`

GetSaldoAFavorOk returns a tuple with the SaldoAFavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaldoAFavor

`func (o *NominaCompensacionInput) SetSaldoAFavor(v float32)`

SetSaldoAFavor sets SaldoAFavor field to given value.


### GetAno

`func (o *NominaCompensacionInput) GetAno() int32`

GetAno returns the Ano field if non-nil, zero value otherwise.

### GetAnoOk

`func (o *NominaCompensacionInput) GetAnoOk() (*int32, bool)`

GetAnoOk returns a tuple with the Ano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAno

`func (o *NominaCompensacionInput) SetAno(v int32)`

SetAno sets Ano field to given value.


### GetRemanenteSalFav

`func (o *NominaCompensacionInput) GetRemanenteSalFav() float32`

GetRemanenteSalFav returns the RemanenteSalFav field if non-nil, zero value otherwise.

### GetRemanenteSalFavOk

`func (o *NominaCompensacionInput) GetRemanenteSalFavOk() (*float32, bool)`

GetRemanenteSalFavOk returns a tuple with the RemanenteSalFav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanenteSalFav

`func (o *NominaCompensacionInput) SetRemanenteSalFav(v float32)`

SetRemanenteSalFav sets RemanenteSalFav field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


