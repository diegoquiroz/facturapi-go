# InvoiceEgresoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**InvoiceCommonInputPropertiesAllOfCustomer**](InvoiceCommonInputPropertiesAllOfCustomer.md) |  | [optional] 
**Date** | Pointer to **time.Time** | Fecha de expedición del comprobante en formato ISO8601 (UTC String). No puede ser anterior a 72 horas en el pasado, ni posterior al presente. | [optional] [default to "now"]
**Address** | Pointer to [**InvoiceCommonInputPropertiesAllOfAddress**](InvoiceCommonInputPropertiesAllOfAddress.md) |  | [optional] 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar esta factura con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta.  | [optional] 
**FolioNumber** | Pointer to **int32** | Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización. | [optional] [default to autoincremental]
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: &#x60;h1&#x60;, &#x60;h2&#x60;, &#x60;h3&#x60;, &#x60;h4&#x60;, &#x60;h5&#x60;, &#x60;h6&#x60;, &#x60;div&#x60;, &#x60;p&#x60;, &#x60;span&#x60;, &#x60;small&#x60;, &#x60;br&#x60;, &#x60;b&#x60;, &#x60;i&#x60;, &#x60;ul&#x60;, &#x60;ol&#x60;, &#x60;li&#x60;, &#x60;strong&#x60;, &#x60;table&#x60;, &#x60;thead&#x60;, &#x60;tbody&#x60;, &#x60;tfoot&#x60;, &#x60;tr&#x60;, &#x60;th&#x60; y &#x60;td&#x60;. No se permiten atributos ni estilos.  | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]InvoiceableCommonInputNamespacesInner**](InvoiceableCommonInputNamespacesInner.md) |  | [optional] 
**Type** | **string** |  | 
**PaymentForm** | **string** | Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). | 
**RelatedDocuments** | Pointer to [**[]RelatedDocumentInput**](RelatedDocumentInput.md) |  | [optional] 
**Items** | [**[]LineItemEgresoInput**](LineItemEgresoInput.md) |  | 
**Use** | Pointer to **string** | Código de Uso CFDI según el catálogo del SAT. Puedes ver los códigos en [esta tabla](#uso-cfdi), o utilizar las constantes incluídas en nuestras librerías. | [optional] [default to "G01"]
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] [default to "MXN"]
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos (MXN) que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] [default to 1]
**Complements** | Pointer to [**[]CustomComplementInput**](CustomComplementInput.md) |  | [optional] 

## Methods

### NewInvoiceEgresoInput

`func NewInvoiceEgresoInput(type_ string, paymentForm string, items []LineItemEgresoInput, ) *InvoiceEgresoInput`

NewInvoiceEgresoInput instantiates a new InvoiceEgresoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceEgresoInputWithDefaults

`func NewInvoiceEgresoInputWithDefaults() *InvoiceEgresoInput`

NewInvoiceEgresoInputWithDefaults instantiates a new InvoiceEgresoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *InvoiceEgresoInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceEgresoInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceEgresoInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InvoiceEgresoInput) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetDate

`func (o *InvoiceEgresoInput) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceEgresoInput) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceEgresoInput) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceEgresoInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAddress

`func (o *InvoiceEgresoInput) GetAddress() InvoiceCommonInputPropertiesAllOfAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceEgresoInput) GetAddressOk() (*InvoiceCommonInputPropertiesAllOfAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceEgresoInput) SetAddress(v InvoiceCommonInputPropertiesAllOfAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceEgresoInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetExternalId

`func (o *InvoiceEgresoInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InvoiceEgresoInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InvoiceEgresoInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InvoiceEgresoInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *InvoiceEgresoInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *InvoiceEgresoInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *InvoiceEgresoInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *InvoiceEgresoInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetFolioNumber

`func (o *InvoiceEgresoInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *InvoiceEgresoInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *InvoiceEgresoInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *InvoiceEgresoInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *InvoiceEgresoInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InvoiceEgresoInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InvoiceEgresoInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InvoiceEgresoInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *InvoiceEgresoInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *InvoiceEgresoInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *InvoiceEgresoInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *InvoiceEgresoInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *InvoiceEgresoInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *InvoiceEgresoInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *InvoiceEgresoInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *InvoiceEgresoInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *InvoiceEgresoInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InvoiceEgresoInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InvoiceEgresoInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InvoiceEgresoInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetType

`func (o *InvoiceEgresoInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceEgresoInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceEgresoInput) SetType(v string)`

SetType sets Type field to given value.


### GetPaymentForm

`func (o *InvoiceEgresoInput) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *InvoiceEgresoInput) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *InvoiceEgresoInput) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.


### GetRelatedDocuments

`func (o *InvoiceEgresoInput) GetRelatedDocuments() []RelatedDocumentInput`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *InvoiceEgresoInput) GetRelatedDocumentsOk() (*[]RelatedDocumentInput, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *InvoiceEgresoInput) SetRelatedDocuments(v []RelatedDocumentInput)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *InvoiceEgresoInput) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetItems

`func (o *InvoiceEgresoInput) GetItems() []LineItemEgresoInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceEgresoInput) GetItemsOk() (*[]LineItemEgresoInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceEgresoInput) SetItems(v []LineItemEgresoInput)`

SetItems sets Items field to given value.


### GetUse

`func (o *InvoiceEgresoInput) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *InvoiceEgresoInput) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *InvoiceEgresoInput) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *InvoiceEgresoInput) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceEgresoInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceEgresoInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceEgresoInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceEgresoInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *InvoiceEgresoInput) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InvoiceEgresoInput) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InvoiceEgresoInput) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InvoiceEgresoInput) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetComplements

`func (o *InvoiceEgresoInput) GetComplements() []CustomComplementInput`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *InvoiceEgresoInput) GetComplementsOk() (*[]CustomComplementInput, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *InvoiceEgresoInput) SetComplements(v []CustomComplementInput)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *InvoiceEgresoInput) HasComplements() bool`

HasComplements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


