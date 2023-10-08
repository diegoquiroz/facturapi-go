# GlobalInvoiceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Fecha inicial de los recibos que se incluirán en la factura global. Por default, este valor es el inicio del último periodo (día, semana, quincena o mes), según el valor de \&quot;Periodicidad\&quot; (&#x60;periodicity&#x60;) en la configuración de recibos de tu organización.  | [optional] [default to "Inicio del último periodo"]
**To** | Pointer to **string** | Fecha final de los recibos que se incluirán en la factura global. Por default, este valor es el fin del último periodo (día, semana, quincena o mes), según el valor de \&quot;Periodicidad\&quot; (&#x60;periodicity&#x60;) en la configuración de recibos de tu organización.  | [optional] [default to "Fin del último periodo"]
**Periodicity** | **string** | Periodicidad que corresponde al rango de fechas utilizado. Si se omite, se tomará la configuración de recibos de la organización.  | [default to "Propiedad `periodicity` de la configuración de recibos de la organización."]
**Months** | Pointer to **string** | Clave que representa el mes o bimestre de la factura. Consulta los posibles valores en el [catálogo de Meses y Bimestres](#meses-y-bimestres).  | [optional] [default to "Mes contenido en el rango de fechas utilizado."]
**FolioNumber** | Pointer to **int32** | Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización.  | [optional] [default to autoincremental]
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. | [optional] 
**Date** | Pointer to **string** | Fecha de emisión de la factura.  | [optional] [default to "Valor del atributo `to`"]
**PaymentForm** | Pointer to **string** | description: Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). Si se incluye, los recibos se agruparán y se crearán la factura global por la forma de pago.  | [optional] 

## Methods

### NewGlobalInvoiceInput

`func NewGlobalInvoiceInput(periodicity string, ) *GlobalInvoiceInput`

NewGlobalInvoiceInput instantiates a new GlobalInvoiceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalInvoiceInputWithDefaults

`func NewGlobalInvoiceInputWithDefaults() *GlobalInvoiceInput`

NewGlobalInvoiceInputWithDefaults instantiates a new GlobalInvoiceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *GlobalInvoiceInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GlobalInvoiceInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GlobalInvoiceInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GlobalInvoiceInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GlobalInvoiceInput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GlobalInvoiceInput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GlobalInvoiceInput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GlobalInvoiceInput) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetPeriodicity

`func (o *GlobalInvoiceInput) GetPeriodicity() string`

GetPeriodicity returns the Periodicity field if non-nil, zero value otherwise.

### GetPeriodicityOk

`func (o *GlobalInvoiceInput) GetPeriodicityOk() (*string, bool)`

GetPeriodicityOk returns a tuple with the Periodicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicity

`func (o *GlobalInvoiceInput) SetPeriodicity(v string)`

SetPeriodicity sets Periodicity field to given value.


### GetMonths

`func (o *GlobalInvoiceInput) GetMonths() string`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *GlobalInvoiceInput) GetMonthsOk() (*string, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *GlobalInvoiceInput) SetMonths(v string)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *GlobalInvoiceInput) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetFolioNumber

`func (o *GlobalInvoiceInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *GlobalInvoiceInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *GlobalInvoiceInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *GlobalInvoiceInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *GlobalInvoiceInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *GlobalInvoiceInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *GlobalInvoiceInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *GlobalInvoiceInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetDate

`func (o *GlobalInvoiceInput) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GlobalInvoiceInput) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GlobalInvoiceInput) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GlobalInvoiceInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentForm

`func (o *GlobalInvoiceInput) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *GlobalInvoiceInput) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *GlobalInvoiceInput) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *GlobalInvoiceInput) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


