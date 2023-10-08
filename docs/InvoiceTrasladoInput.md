# InvoiceTrasladoInput

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
**Items** | [**[]LineItemTrasladoInput**](LineItemTrasladoInput.md) |  | 
**Complements** | Pointer to [**[]CustomComplementInput**](CustomComplementInput.md) |  | [optional] 
**Use** | Pointer to **string** | Código de Uso CFDI según el catálogo del SAT. Puedes ver los códigos en [esta tabla](#uso-cfdi), o utilizar las constantes incluídas en nuestras librerías.  | [optional] [default to "G01"]
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] [default to "MXN"]
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos (MXN) que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;.  | [optional] [default to 1]
**RelatedDocuments** | Pointer to [**[]RelatedDocumentInput**](RelatedDocumentInput.md) |  | [optional] 

## Methods

### NewInvoiceTrasladoInput

`func NewInvoiceTrasladoInput(type_ string, items []LineItemTrasladoInput, ) *InvoiceTrasladoInput`

NewInvoiceTrasladoInput instantiates a new InvoiceTrasladoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceTrasladoInputWithDefaults

`func NewInvoiceTrasladoInputWithDefaults() *InvoiceTrasladoInput`

NewInvoiceTrasladoInputWithDefaults instantiates a new InvoiceTrasladoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *InvoiceTrasladoInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceTrasladoInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceTrasladoInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InvoiceTrasladoInput) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetDate

`func (o *InvoiceTrasladoInput) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceTrasladoInput) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceTrasladoInput) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceTrasladoInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAddress

`func (o *InvoiceTrasladoInput) GetAddress() InvoiceCommonInputPropertiesAllOfAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceTrasladoInput) GetAddressOk() (*InvoiceCommonInputPropertiesAllOfAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceTrasladoInput) SetAddress(v InvoiceCommonInputPropertiesAllOfAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceTrasladoInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetExternalId

`func (o *InvoiceTrasladoInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InvoiceTrasladoInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InvoiceTrasladoInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InvoiceTrasladoInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *InvoiceTrasladoInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *InvoiceTrasladoInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *InvoiceTrasladoInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *InvoiceTrasladoInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetFolioNumber

`func (o *InvoiceTrasladoInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *InvoiceTrasladoInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *InvoiceTrasladoInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *InvoiceTrasladoInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *InvoiceTrasladoInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InvoiceTrasladoInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InvoiceTrasladoInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InvoiceTrasladoInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *InvoiceTrasladoInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *InvoiceTrasladoInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *InvoiceTrasladoInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *InvoiceTrasladoInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *InvoiceTrasladoInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *InvoiceTrasladoInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *InvoiceTrasladoInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *InvoiceTrasladoInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *InvoiceTrasladoInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InvoiceTrasladoInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InvoiceTrasladoInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InvoiceTrasladoInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetType

`func (o *InvoiceTrasladoInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceTrasladoInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceTrasladoInput) SetType(v string)`

SetType sets Type field to given value.


### GetItems

`func (o *InvoiceTrasladoInput) GetItems() []LineItemTrasladoInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceTrasladoInput) GetItemsOk() (*[]LineItemTrasladoInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceTrasladoInput) SetItems(v []LineItemTrasladoInput)`

SetItems sets Items field to given value.


### GetComplements

`func (o *InvoiceTrasladoInput) GetComplements() []CustomComplementInput`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *InvoiceTrasladoInput) GetComplementsOk() (*[]CustomComplementInput, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *InvoiceTrasladoInput) SetComplements(v []CustomComplementInput)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *InvoiceTrasladoInput) HasComplements() bool`

HasComplements returns a boolean if a field has been set.

### GetUse

`func (o *InvoiceTrasladoInput) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *InvoiceTrasladoInput) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *InvoiceTrasladoInput) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *InvoiceTrasladoInput) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceTrasladoInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceTrasladoInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceTrasladoInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceTrasladoInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *InvoiceTrasladoInput) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InvoiceTrasladoInput) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InvoiceTrasladoInput) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InvoiceTrasladoInput) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetRelatedDocuments

`func (o *InvoiceTrasladoInput) GetRelatedDocuments() []RelatedDocumentInput`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *InvoiceTrasladoInput) GetRelatedDocumentsOk() (*[]RelatedDocumentInput, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *InvoiceTrasladoInput) SetRelatedDocuments(v []RelatedDocumentInput)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *InvoiceTrasladoInput) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


