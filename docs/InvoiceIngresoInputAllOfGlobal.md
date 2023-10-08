# InvoiceIngresoInputAllOfGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Periodicity** | **string** | Periodicidad que abarca la factura global.  - &#x60;day&#x60;: Diario - &#x60;week&#x60;: Semanal - &#x60;fortnight&#x60;: Quincenal - &#x60;month&#x60;: Mensual - &#x60;two_months&#x60;: Bimestral  | 
**Months** | **string** | Clave que representa el mes o bimestre de la factura. Consulta los posibles valores en el [catálogo de Meses y Bimestres](#meses-y-bimestres).  | 
**Year** | **int32** | Año de la factura. | 

## Methods

### NewInvoiceIngresoInputAllOfGlobal

`func NewInvoiceIngresoInputAllOfGlobal(periodicity string, months string, year int32, ) *InvoiceIngresoInputAllOfGlobal`

NewInvoiceIngresoInputAllOfGlobal instantiates a new InvoiceIngresoInputAllOfGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIngresoInputAllOfGlobalWithDefaults

`func NewInvoiceIngresoInputAllOfGlobalWithDefaults() *InvoiceIngresoInputAllOfGlobal`

NewInvoiceIngresoInputAllOfGlobalWithDefaults instantiates a new InvoiceIngresoInputAllOfGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodicity

`func (o *InvoiceIngresoInputAllOfGlobal) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *InvoiceIngresoInputAllOfGlobal) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *InvoiceIngresoInputAllOfGlobal) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.


### GetMonths

`func (o *InvoiceIngresoInputAllOfGlobal) GetMonths() string`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *InvoiceIngresoInputAllOfGlobal) GetMonthsOk() (*string, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *InvoiceIngresoInputAllOfGlobal) SetMonths(v string)`

SetMonths sets Months field to given value.


### GetYear

`func (o *InvoiceIngresoInputAllOfGlobal) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *InvoiceIngresoInputAllOfGlobal) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *InvoiceIngresoInputAllOfGlobal) SetYear(v int32)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


