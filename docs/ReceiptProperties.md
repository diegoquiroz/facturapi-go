# ReceiptProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | Fecha de emisión del recibo. | [optional] 
**PaymentForm** | Pointer to **string** | Código que representa la forma de pago, según el [catálogo del SAT](#forma-de-pago). | [optional] 
**FolioNumber** | Pointer to **int32** | Autoincremental. Número de folio del recibo para control interno y sin validez fiscal. | [optional] 
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] 
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] 
**Branch** | Pointer to **string** | Nombre de la sucursal donde se expidió el recibo. | [optional] 
**ExternalId** | Pointer to **string** | Identificador que puedes usar para relacionar este recibo con tus registros para después buscar por este número. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Fecha de expiración en formato ISO8601 (UTC String). Es la fecha límite para que el cliente pueda facturar su recibo en el portal de autofactura. Se calcula automáticamente a partir de las configuraciones de recibo de la organización.  | [optional] 
**Status** | Pointer to **string** | Estado actual del recibo. | [optional] 
**SelfInvoiceUrl** | Pointer to **string** | Dirección URL para realizar autofactura. Incluye el &#x60;key&#x60; del recibo. Puedes usarla para generar un botón o un QR de facturación para tus clientes.  | [optional] 
**Total** | Pointer to **float32** | Monto total de la operación | [optional] 
**Invoice** | Pointer to **string** | ID de la factura asociada, en caso de estar facturado. | [optional] 
**Key** | Pointer to **string** | Autogenerado. Identificador único alfanumérico corto, útil para acceder a la autofactura desde tu micrositio en factura.space | [optional] 
**Items** | Pointer to [**[]LineItem**](LineItem.md) |  | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. | [optional] 

## Methods

### NewReceiptProperties

`func NewReceiptProperties() *ReceiptProperties`

NewReceiptProperties instantiates a new ReceiptProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptPropertiesWithDefaults

`func NewReceiptPropertiesWithDefaults() *ReceiptProperties`

NewReceiptPropertiesWithDefaults instantiates a new ReceiptProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ReceiptProperties) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ReceiptProperties) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ReceiptProperties) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ReceiptProperties) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentForm

`func (o *ReceiptProperties) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *ReceiptProperties) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *ReceiptProperties) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *ReceiptProperties) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.

### GetFolioNumber

`func (o *ReceiptProperties) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *ReceiptProperties) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *ReceiptProperties) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *ReceiptProperties) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *ReceiptProperties) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ReceiptProperties) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ReceiptProperties) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ReceiptProperties) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *ReceiptProperties) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReceiptProperties) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReceiptProperties) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReceiptProperties) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBranch

`func (o *ReceiptProperties) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReceiptProperties) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReceiptProperties) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ReceiptProperties) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetExternalId

`func (o *ReceiptProperties) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ReceiptProperties) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ReceiptProperties) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ReceiptProperties) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ReceiptProperties) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ReceiptProperties) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ReceiptProperties) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ReceiptProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatus

`func (o *ReceiptProperties) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReceiptProperties) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReceiptProperties) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReceiptProperties) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSelfInvoiceUrl

`func (o *ReceiptProperties) GetSelfInvoiceUrl() string`

GetSelfInvoiceUrl returns the SelfInvoiceUrl field if non-nil, zero value otherwise.

### GetSelfInvoiceUrlOk

`func (o *ReceiptProperties) GetSelfInvoiceUrlOk() (*string, bool)`

GetSelfInvoiceUrlOk returns a tuple with the SelfInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfInvoiceUrl

`func (o *ReceiptProperties) SetSelfInvoiceUrl(v string)`

SetSelfInvoiceUrl sets SelfInvoiceUrl field to given value.

### HasSelfInvoiceUrl

`func (o *ReceiptProperties) HasSelfInvoiceUrl() bool`

HasSelfInvoiceUrl returns a boolean if a field has been set.

### GetTotal

`func (o *ReceiptProperties) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReceiptProperties) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReceiptProperties) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ReceiptProperties) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetInvoice

`func (o *ReceiptProperties) GetInvoice() string`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *ReceiptProperties) GetInvoiceOk() (*string, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *ReceiptProperties) SetInvoice(v string)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *ReceiptProperties) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetKey

`func (o *ReceiptProperties) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReceiptProperties) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReceiptProperties) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ReceiptProperties) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetItems

`func (o *ReceiptProperties) GetItems() []LineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ReceiptProperties) GetItemsOk() (*[]LineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ReceiptProperties) SetItems(v []LineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ReceiptProperties) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *ReceiptProperties) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ReceiptProperties) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ReceiptProperties) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ReceiptProperties) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


