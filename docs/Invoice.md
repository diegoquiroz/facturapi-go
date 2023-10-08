# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha de registro | [optional] 
**Livemode** | Pointer to **bool** | Si el valor es &#x60;true&#x60;, indica que el objeto fue creado en ambiente Live; o si es &#x60;false&#x60;, en ambiente Test. | [optional] 
**Status** | Pointer to **string** | Estado actual de la factura.  | [optional] 
**CancellationStatus** | Pointer to **string** | Estado actual de la solicitud de cancelación, en caso de haberla realizado. Puedes leer más a detalle en la sección de [Cancelar Factura](#operation/deleteInvoice)).  | [optional] 
**VerificationUrl** | Pointer to **string** | Dirección URL para verificar el estado del CFDI en el portal del SAT. Este link es el mismo que aparece en el código QR, en el PDF de la factura. | [optional] 
**Date** | Pointer to **time.Time** | Fecha de expedición del comprobante en formato ISO8601 (UTC String). | [optional] [default to "now"]
**Address** | Pointer to [**InvoicePropertiesAddress**](InvoicePropertiesAddress.md) |  | [optional] 
**Type** | Pointer to **string** | Tipo de comprobante. Puede tener los valores &#x60;\&quot;I\&quot;&#x60;: Ingreso, &#x60;\&quot;P\&quot;&#x60;: Pago, &#x60;\&quot;E\&quot;&#x60;: Egreso, &#x60;\&quot;N\&quot;&#x60;: Nómina, &#x60;\&quot;T\&quot;&#x60;: Traslado.  | [optional] 
**Customer** | Pointer to [**CuustomerInfo**](CuustomerInfo.md) |  | [optional] 
**Total** | Pointer to **float32** | Monto total facturado. | [optional] 
**Uuid** | Pointer to **string** | Folio fiscal de la factura, asignado por el SAT. | [optional] 
**FolioNumber** | Pointer to **int32** | Número de folio autoincremental para control interno y sin validez fiscal. | [optional] 
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. En el PDF se imprime junto al //www.facturapi.io/img/logo.svg | [optional] 
**ExternalId** | Pointer to **string** | Identificador que puedes usar para relacionar esta factura con tus registros para después buscar por este número. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. | [optional] 
**PaymentForm** | Pointer to **string** | Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). | [optional] 
**Items** | Pointer to [**[]LineItem**](LineItem.md) |  | [optional] 
**RelatedDocuments** | Pointer to [**[]RelatedDocument**](RelatedDocument.md) |  | [optional] 
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] 
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] 
**Complements** | Pointer to [**[]NominaOrCustomComplementProperties**](NominaOrCustomComplementProperties.md) |  | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite insertar código HTML con tu propio contenido. | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]NamespaceProperties**](NamespaceProperties.md) |  | [optional] 
**Stamp** | Pointer to [**Stamp**](Stamp.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Invoice) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLivemode

`func (o *Invoice) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Invoice) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Invoice) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Invoice) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCancellationStatus

`func (o *Invoice) GetCancellationStatus() string`

GetCancellationStatus returns the CancellationStatus field if non-nil, zero value otherwise.

### GetCancellationStatusOk

`func (o *Invoice) GetCancellationStatusOk() (*string, bool)`

GetCancellationStatusOk returns a tuple with the CancellationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationStatus

`func (o *Invoice) SetCancellationStatus(v string)`

SetCancellationStatus sets CancellationStatus field to given value.

### HasCancellationStatus

`func (o *Invoice) HasCancellationStatus() bool`

HasCancellationStatus returns a boolean if a field has been set.

### GetVerificationUrl

`func (o *Invoice) GetVerificationUrl() string`

GetVerificationUrl returns the VerificationUrl field if non-nil, zero value otherwise.

### GetVerificationUrlOk

`func (o *Invoice) GetVerificationUrlOk() (*string, bool)`

GetVerificationUrlOk returns a tuple with the VerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUrl

`func (o *Invoice) SetVerificationUrl(v string)`

SetVerificationUrl sets VerificationUrl field to given value.

### HasVerificationUrl

`func (o *Invoice) HasVerificationUrl() bool`

HasVerificationUrl returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAddress

`func (o *Invoice) GetAddress() InvoicePropertiesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Invoice) GetAddressOk() (*InvoicePropertiesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Invoice) SetAddress(v InvoicePropertiesAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Invoice) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetType

`func (o *Invoice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Invoice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCustomer

`func (o *Invoice) GetCustomer() CuustomerInfo`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Invoice) GetCustomerOk() (*CuustomerInfo, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Invoice) SetCustomer(v CuustomerInfo)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Invoice) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUuid

`func (o *Invoice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Invoice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Invoice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Invoice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetFolioNumber

`func (o *Invoice) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *Invoice) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *Invoice) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *Invoice) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *Invoice) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Invoice) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Invoice) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Invoice) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetExternalId

`func (o *Invoice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Invoice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Invoice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Invoice) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *Invoice) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Invoice) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Invoice) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Invoice) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetPaymentForm

`func (o *Invoice) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *Invoice) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *Invoice) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.

### HasPaymentForm

`func (o *Invoice) HasPaymentForm() bool`

HasPaymentForm returns a boolean if a field has been set.

### GetItems

`func (o *Invoice) GetItems() []LineItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Invoice) GetItemsOk() (*[]LineItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Invoice) SetItems(v []LineItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *Invoice) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRelatedDocuments

`func (o *Invoice) GetRelatedDocuments() []RelatedDocument`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *Invoice) GetRelatedDocumentsOk() (*[]RelatedDocument, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *Invoice) SetRelatedDocuments(v []RelatedDocument)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *Invoice) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *Invoice) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Invoice) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Invoice) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Invoice) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetComplements

`func (o *Invoice) GetComplements() []NominaOrCustomComplementProperties`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *Invoice) GetComplementsOk() (*[]NominaOrCustomComplementProperties, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *Invoice) SetComplements(v []NominaOrCustomComplementProperties)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *Invoice) HasComplements() bool`

HasComplements returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *Invoice) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *Invoice) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *Invoice) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *Invoice) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *Invoice) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *Invoice) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *Invoice) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *Invoice) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *Invoice) GetNamespaces() []NamespaceProperties`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *Invoice) GetNamespacesOk() (*[]NamespaceProperties, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *Invoice) SetNamespaces(v []NamespaceProperties)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *Invoice) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetStamp

`func (o *Invoice) GetStamp() Stamp`

GetStamp returns the Stamp field if non-nil, zero value otherwise.

### GetStampOk

`func (o *Invoice) GetStampOk() (*Stamp, bool)`

GetStampOk returns a tuple with the Stamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamp

`func (o *Invoice) SetStamp(v Stamp)`

SetStamp sets Stamp field to given value.

### HasStamp

`func (o *Invoice) HasStamp() bool`

HasStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


