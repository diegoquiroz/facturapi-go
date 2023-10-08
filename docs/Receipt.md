# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha de registro | [optional] 
**Livemode** | Pointer to **bool** | Si el valor es &#x60;true&#x60;, indica que el objeto fue creado en ambiente Live; o si es &#x60;false&#x60;, en ambiente Test. | [optional] 
**Date** | Pointer to **time.Time** | Fecha de emisión del recibo. Por defecto se utiliza la fecha actual. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Fecha de expiración en formato ISO8601 (UTC String). Es la fecha límite para que el cliente pueda facturar su recibo en el portal de autofactura. Se calcula automáticamente a partir de las configuraciones de recibo de la organización.  | [optional] 
**Status** | Pointer to **string** | Estado actual del recibo. | [optional] 
**SelfInvoiceUrl** | Pointer to **string** | Dirección URL para realizar autofactura. Incluye el &#x60;key&#x60; del recibo. Puedes usarla para generar un botón o un QR de facturación para tus clientes.  | [optional] 
**Total** | Pointer to **float32** | Monto total de la operación | [optional] 
**Invoice** | Pointer to **string** | ID de la factura asociada, en caso de estar facturado. | [optional] 
**Key** | Pointer to **string** | Autogenerado. Identificador único alfanumérico corto, útil para acceder a la autofactura desde tu micrositio en factura.space | [optional] 
**Items** | Pointer to [**[]LineItem**](LineItem.md) |  | [optional] 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar este recibo con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. | [optional] 
**PaymentForm** | Pointer to **string** | Código que representa la forma de pago, según el [catálogo del SAT](#forma-de-pago). | [optional] 
**FolioNumber** | Pointer to **int32** | Autoincremental. Número de folio del recibo para control interno y sin validez fiscal. | [optional] 
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] 
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] 
**Branch** | Pointer to **string** | Nombre de la sucursal donde se expidió el recibo. | [optional] 

## Methods

### NewReceipt

`func NewReceipt() *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Receipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Receipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Receipt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Receipt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Receipt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Receipt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Receipt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Receipt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLivemode

`func (o *Receipt) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Receipt) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Receipt) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Receipt) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetDate

`func (o *Receipt) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Receipt) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Receipt) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Receipt) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Receipt) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Receipt) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Receipt) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Receipt) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetStatus

`func (o *Receipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Receipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Receipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Receipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSelfInvoiceUrl

`func (o *Receipt) GetSelfInvoiceUrl() string`

GetSelfInvoiceUrl returns the SelfInvoiceUrl field if non-nil, zero value otherwise.

### GetSelfInvoiceUrlOk

`func (o *Receipt) GetSelfInvoiceUrlOk() (*string, bool)`

GetSelfInvoiceUrlOk returns a tuple with the SelfInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfInvoiceUrl

`func (o *Receipt) SetSelfInvoiceUrl(v string)`

SetSelfInvoiceUrl sets SelfInvoiceUrl field to given value.

### HasSelfInvoiceUrl

`func (o *Receipt) HasSelfInvoiceUrl() bool`

HasSelfInvoiceUrl returns a boolean if a field has been set.

### GetTotal

`func (o *Receipt) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Receipt) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Receipt) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Receipt) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetInvoice

`func (o *Receipt) GetInvoice() string`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Receipt) GetInvoiceOk() (*string, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Receipt) SetInvoice(v string)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *Receipt) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetKey

`func (o *Receipt) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Receipt) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Receipt) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Receipt) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetItems

`func (o *Receipt) GetItems() []LineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Receipt) GetItemsOk() (*[]LineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Receipt) SetItems(v []LineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Receipt) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetExternalId

`func (o *Receipt) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Receipt) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Receipt) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Receipt) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *Receipt) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Receipt) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Receipt) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Receipt) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetPaymentForm

`func (o *Receipt) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *Receipt) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *Receipt) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *Receipt) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.

### GetFolioNumber

`func (o *Receipt) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *Receipt) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *Receipt) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *Receipt) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetCurrency

`func (o *Receipt) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Receipt) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Receipt) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Receipt) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *Receipt) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Receipt) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Receipt) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Receipt) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetBranch

`func (o *Receipt) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Receipt) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Receipt) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Receipt) HasBranch() bool`

HasBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


