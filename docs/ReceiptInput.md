# ReceiptInput

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
**Items** | [**[]LineItemInput**](LineItemInput.md) |  | 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta.  | [optional] 

## Methods

### NewReceiptInput

`func NewReceiptInput(items []LineItemInput, ) *ReceiptInput`

NewReceiptInput instantiates a new ReceiptInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptInputWithDefaults

`func NewReceiptInputWithDefaults() *ReceiptInput`

NewReceiptInputWithDefaults instantiates a new ReceiptInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ReceiptInput) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ReceiptInput) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ReceiptInput) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ReceiptInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentForm

`func (o *ReceiptInput) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *ReceiptInput) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *ReceiptInput) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *ReceiptInput) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.

### GetFolioNumber

`func (o *ReceiptInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *ReceiptInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *ReceiptInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *ReceiptInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *ReceiptInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReceiptInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *ReceiptInput) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReceiptInput) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReceiptInput) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReceiptInput) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBranch

`func (o *ReceiptInput) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReceiptInput) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReceiptInput) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ReceiptInput) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetExternalId

`func (o *ReceiptInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReceiptInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReceiptInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReceiptInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetItems

`func (o *ReceiptInput) GetItems() []LineItemInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReceiptInput) GetItemsOk() (*[]LineItemInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReceiptInput) SetItems(v []LineItemInput)`

SetItems sets Items field to given value.


### GetIdempotencyKey

`func (o *ReceiptInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ReceiptInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ReceiptInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ReceiptInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


