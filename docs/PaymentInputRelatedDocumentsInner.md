# PaymentInputRelatedDocumentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Folio fiscal ó UUID del comprobante de ingreso relacionado. | 
**Amount** | **float32** | Cantidad del pago correspondiente al comprobante relacionado, usando el método de pago indicado en este elemento del arreglo de pagos. Este valor debe ser expresado en la moneda definida en &#x60;related_documents[].currency&#x60;.  | 
**Taxes** | [**[]PaymentInputRelatedDocumentsInnerTaxesInner**](PaymentInputRelatedDocumentsInnerTaxesInner.md) |  | 
**Taxability** | Pointer to **string** | Código que representa si el bien o servicio es objeto de impuesto o no. Este atributo corresponde al campo \&quot;ObjetoImp\&quot; en el CFDI.  - &#x60;01&#x60;: No objeto de impuesto. - &#x60;02&#x60;: Sí objeto de impuesto. - &#x60;03&#x60;: Sí objeto de impuesto, pero no obligado a desglose. - &#x60;04&#x60;: Sí objeto de impuesto, y no causa impuesto.  | [optional] [default to "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "]
**Installment** | **int32** | Número de parcialidad del pago. | 
**LastBalance** | **float32** | Cantidad que estaba pendiente por pagar antes de recibir este pago. Este valor se expresa en la moneda definida en &#x60;payments[].related[].currency&#x60;. | 
**Currency** | Pointer to **string** | Si la moneda utilizada en la factura relacionada no es moneda nacional (MXN), debe especificarse su valor acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] [default to "MXN"]
**Exchange** | Pointer to **float32** | Obligatorio cuando la moneda del documento relacionado es distinta a la moneda de pago. Tipo de cambio entre las dos monedas al momento del pago. Ejemplo: La factura de iingreso relacionada se registra en USD, mientras que el pago actual se realiza en MXN, este atributo debería registrarse como &#x60;0.45&#x60; (USD/MXN).  | [optional] 
**FolioNumber** | Pointer to **int32** | Opcionalmente se puede incluir el número de folio del documento relacionado. | [optional] 
**Series** | Pointer to **string** | Opcionalmente se puede incluir la serie del documento relacionado. | [optional] 

## Methods

### NewPaymentInputRelatedDocumentsInner

`func NewPaymentInputRelatedDocumentsInner(uuid string, amount float32, taxes []PaymentInputRelatedDocumentsInnerTaxesInner, installment int32, lastBalance float32, ) *PaymentInputRelatedDocumentsInner`

NewPaymentInputRelatedDocumentsInner instantiates a new PaymentInputRelatedDocumentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInputRelatedDocumentsInnerWithDefaults

`func NewPaymentInputRelatedDocumentsInnerWithDefaults() *PaymentInputRelatedDocumentsInner`

NewPaymentInputRelatedDocumentsInnerWithDefaults instantiates a new PaymentInputRelatedDocumentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PaymentInputRelatedDocumentsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PaymentInputRelatedDocumentsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PaymentInputRelatedDocumentsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetAmount

`func (o *PaymentInputRelatedDocumentsInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInputRelatedDocumentsInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInputRelatedDocumentsInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetTaxes

`func (o *PaymentInputRelatedDocumentsInner) GetTaxes() []PaymentInputRelatedDocumentsInnerTaxesInner`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *PaymentInputRelatedDocumentsInner) GetTaxesOk() (*[]PaymentInputRelatedDocumentsInnerTaxesInner, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *PaymentInputRelatedDocumentsInner) SetTaxes(v []PaymentInputRelatedDocumentsInnerTaxesInner)`

SetTaxes sets Taxes field to given value.


### GetTaxability

`func (o *PaymentInputRelatedDocumentsInner) GetTaxability() string`

GetTaxability returns the Taxability field if non-nil, zero value otherwise.

### GetTaxabilityOk

`func (o *PaymentInputRelatedDocumentsInner) GetTaxabilityOk() (*string, bool)`

GetTaxabilityOk returns a tuple with the Taxability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxability

`func (o *PaymentInputRelatedDocumentsInner) SetTaxability(v string)`

SetTaxability sets Taxability field to given value.

### HasTaxability

`func (o *PaymentInputRelatedDocumentsInner) HasTaxability() bool`

HasTaxability returns a boolean if a field has been set.

### GetInstallment

`func (o *PaymentInputRelatedDocumentsInner) GetInstallment() int32`

GetInstallment returns the Installment field if non-nil, zero value otherwise.

### GetInstallmentOk

`func (o *PaymentInputRelatedDocumentsInner) GetInstallmentOk() (*int32, bool)`

GetInstallmentOk returns a tuple with the Installment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallment

`func (o *PaymentInputRelatedDocumentsInner) SetInstallment(v int32)`

SetInstallment sets Installment field to given value.


### GetLastBalance

`func (o *PaymentInputRelatedDocumentsInner) GetLastBalance() float32`

GetLastBalance returns the LastBalance field if non-nil, zero value otherwise.

### GetLastBalanceOk

`func (o *PaymentInputRelatedDocumentsInner) GetLastBalanceOk() (*float32, bool)`

GetLastBalanceOk returns a tuple with the LastBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBalance

`func (o *PaymentInputRelatedDocumentsInner) SetLastBalance(v float32)`

SetLastBalance sets LastBalance field to given value.


### GetCurrency

`func (o *PaymentInputRelatedDocumentsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentInputRelatedDocumentsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentInputRelatedDocumentsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentInputRelatedDocumentsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *PaymentInputRelatedDocumentsInner) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *PaymentInputRelatedDocumentsInner) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *PaymentInputRelatedDocumentsInner) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *PaymentInputRelatedDocumentsInner) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetFolioNumber

`func (o *PaymentInputRelatedDocumentsInner) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *PaymentInputRelatedDocumentsInner) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *PaymentInputRelatedDocumentsInner) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *PaymentInputRelatedDocumentsInner) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *PaymentInputRelatedDocumentsInner) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *PaymentInputRelatedDocumentsInner) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *PaymentInputRelatedDocumentsInner) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *PaymentInputRelatedDocumentsInner) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


