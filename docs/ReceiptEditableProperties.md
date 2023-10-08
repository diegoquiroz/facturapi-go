# ReceiptEditableProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | Fecha de emisión del recibo. Por defecto se utiliza la fecha actual. | [optional] 
**PaymentForm** | Pointer to **string** | Código que representa la forma de pago, según el [catálogo del SAT](#forma-de-pago). | [optional] 
**FolioNumber** | Pointer to **int32** | Autoincremental. Número de folio del recibo para control interno y sin validez fiscal. | [optional] 
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] 
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] 
**Branch** | Pointer to **string** | Nombre de la sucursal donde se expidió el recibo. | [optional] 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar este recibo con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 

## Methods

### NewReceiptEditableProperties

`func NewReceiptEditableProperties() *ReceiptEditableProperties`

NewReceiptEditableProperties instantiates a new ReceiptEditableProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptEditablePropertiesWithDefaults

`func NewReceiptEditablePropertiesWithDefaults() *ReceiptEditableProperties`

NewReceiptEditablePropertiesWithDefaults instantiates a new ReceiptEditableProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ReceiptEditableProperties) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ReceiptEditableProperties) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ReceiptEditableProperties) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ReceiptEditableProperties) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentForm

`func (o *ReceiptEditableProperties) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *ReceiptEditableProperties) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *ReceiptEditableProperties) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *ReceiptEditableProperties) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.

### GetFolioNumber

`func (o *ReceiptEditableProperties) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *ReceiptEditableProperties) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *ReceiptEditableProperties) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *ReceiptEditableProperties) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *ReceiptEditableProperties) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptEditableProperties) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptEditableProperties) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReceiptEditableProperties) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *ReceiptEditableProperties) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReceiptEditableProperties) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReceiptEditableProperties) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReceiptEditableProperties) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBranch

`func (o *ReceiptEditableProperties) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReceiptEditableProperties) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReceiptEditableProperties) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ReceiptEditableProperties) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetExternalId

`func (o *ReceiptEditableProperties) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReceiptEditableProperties) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReceiptEditableProperties) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReceiptEditableProperties) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


