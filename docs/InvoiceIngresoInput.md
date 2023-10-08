# InvoiceIngresoInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [**InvoiceCommonInputPropertiesAllOfCustomer**](InvoiceCommonInputPropertiesAllOfCustomer.md) |  | 
**Date** | Pointer to **time.Time** | Fecha de expedición del comprobante en formato ISO8601 (UTC String). No puede ser anterior a 72 horas en el pasado, ni posterior al presente. | [optional] [default to "now"]
**Address** | Pointer to [**InvoiceCommonInputPropertiesAllOfAddress**](InvoiceCommonInputPropertiesAllOfAddress.md) |  | [optional] 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar esta factura con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta.  | [optional] 
**FolioNumber** | Pointer to **int32** | Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización. | [optional] [default to autoincremental]
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: &#x60;h1&#x60;, &#x60;h2&#x60;, &#x60;h3&#x60;, &#x60;h4&#x60;, &#x60;h5&#x60;, &#x60;h6&#x60;, &#x60;div&#x60;, &#x60;p&#x60;, &#x60;span&#x60;, &#x60;small&#x60;, &#x60;br&#x60;, &#x60;b&#x60;, &#x60;i&#x60;, &#x60;ul&#x60;, &#x60;ol&#x60;, &#x60;li&#x60;, &#x60;strong&#x60;, &#x60;table&#x60;, &#x60;thead&#x60;, &#x60;tbody&#x60;, &#x60;tfoot&#x60;, &#x60;tr&#x60;, &#x60;th&#x60; y &#x60;td&#x60;. No se permiten atributos ni estilos.  | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]InvoiceableCommonInputNamespacesInner**](InvoiceableCommonInputNamespacesInner.md) |  | [optional] 
**Type** | Pointer to **string** | Tipo de comprobante. El valor default es &#x60;“I”&#x60; (Ingreso). | [optional] [default to "I"]
**Items** | [**[]LineItemInput**](LineItemInput.md) |  | 
**PaymentForm** | **string** | Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago). | 
**PaymentMethod** | Pointer to **string** | Código del método de pago según el catálogo del SAT.  - &#x60;PUE&#x60;: Pago en Una sola Exhibición - &#x60;PPD&#x60;: Pago en Parcialidades o Diferido  | [optional] [default to "PUE"]
**Use** | Pointer to **string** | Código de Uso CFDI según el catálogo del SAT. Puedes ver los códigos en [esta tabla](#uso-cfdi), o utilizar las constantes incluídas en nuestras librerías.  Para factura global debe ingresarse la clave &#x60;S01&#x60;&#x60;.  | [optional] [default to "G01"]
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] [default to "MXN"]
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos (MXN) que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;.  | [optional] [default to 1]
**Conditions** | Pointer to **string** | Condiciones de pago | [optional] 
**RelatedDocuments** | Pointer to [**[]RelatedDocumentInput**](RelatedDocumentInput.md) |  | [optional] 
**Global** | Pointer to [**InvoiceIngresoInputAllOfGlobal**](InvoiceIngresoInputAllOfGlobal.md) |  | [optional] 
**Export** | Pointer to **string** | Indica si el comprobante ampara una operación de exportación.  - &#x60;01&#x60;: No aplica - &#x60;02&#x60;: Definitiva con clave A1 - &#x60;03&#x60;: Temporal - &#x60;04&#x60;: Definitiva con clave distinta a A1 o cuando no existe enajenación en términos del CFF  | [optional] [default to "01"]
**Complements** | Pointer to [**[]CustomComplementInput**](CustomComplementInput.md) |  | [optional] 

## Methods

### NewInvoiceIngresoInput

`func NewInvoiceIngresoInput(customer InvoiceCommonInputPropertiesAllOfCustomer, items []LineItemInput, paymentForm string, ) *InvoiceIngresoInput`

NewInvoiceIngresoInput instantiates a new InvoiceIngresoInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIngresoInputWithDefaults

`func NewInvoiceIngresoInputWithDefaults() *InvoiceIngresoInput`

NewInvoiceIngresoInputWithDefaults instantiates a new InvoiceIngresoInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *InvoiceIngresoInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceIngresoInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceIngresoInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer)`

SetCustomer sets Customer field to given value.


### GetDate

`func (o *InvoiceIngresoInput) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceIngresoInput) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceIngresoInput) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceIngresoInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAddress

`func (o *InvoiceIngresoInput) GetAddress() InvoiceCommonInputPropertiesAllOfAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InvoiceIngresoInput) GetAddressOk() (*InvoiceCommonInputPropertiesAllOfAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InvoiceIngresoInput) SetAddress(v InvoiceCommonInputPropertiesAllOfAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InvoiceIngresoInput) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetExternalId

`func (o *InvoiceIngresoInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InvoiceIngresoInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InvoiceIngresoInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InvoiceIngresoInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *InvoiceIngresoInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *InvoiceIngresoInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *InvoiceIngresoInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *InvoiceIngresoInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetFolioNumber

`func (o *InvoiceIngresoInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *InvoiceIngresoInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *InvoiceIngresoInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *InvoiceIngresoInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *InvoiceIngresoInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InvoiceIngresoInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InvoiceIngresoInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InvoiceIngresoInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *InvoiceIngresoInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *InvoiceIngresoInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *InvoiceIngresoInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *InvoiceIngresoInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *InvoiceIngresoInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *InvoiceIngresoInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *InvoiceIngresoInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *InvoiceIngresoInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *InvoiceIngresoInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InvoiceIngresoInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InvoiceIngresoInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InvoiceIngresoInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetType

`func (o *InvoiceIngresoInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceIngresoInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceIngresoInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceIngresoInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *InvoiceIngresoInput) GetItems() []LineItemInput`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InvoiceIngresoInput) GetItemsOk() (*[]LineItemInput, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InvoiceIngresoInput) SetItems(v []LineItemInput)`

SetItems sets Items field to given value.


### GetPaymentForm

`func (o *InvoiceIngresoInput) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *InvoiceIngresoInput) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *InvoiceIngresoInput) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.


### GetPaymentMethod

`func (o *InvoiceIngresoInput) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InvoiceIngresoInput) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InvoiceIngresoInput) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *InvoiceIngresoInput) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetUse

`func (o *InvoiceIngresoInput) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *InvoiceIngresoInput) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *InvoiceIngresoInput) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *InvoiceIngresoInput) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceIngresoInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceIngresoInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceIngresoInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceIngresoInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *InvoiceIngresoInput) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InvoiceIngresoInput) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InvoiceIngresoInput) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InvoiceIngresoInput) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetConditions

`func (o *InvoiceIngresoInput) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InvoiceIngresoInput) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InvoiceIngresoInput) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InvoiceIngresoInput) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetRelatedDocuments

`func (o *InvoiceIngresoInput) GetRelatedDocuments() []RelatedDocumentInput`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *InvoiceIngresoInput) GetRelatedDocumentsOk() (*[]RelatedDocumentInput, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *InvoiceIngresoInput) SetRelatedDocuments(v []RelatedDocumentInput)`

SetRelatedDocuments sets RelatedDocuments field to given value.

### HasRelatedDocuments

`func (o *InvoiceIngresoInput) HasRelatedDocuments() bool`

HasRelatedDocuments returns a boolean if a field has been set.

### GetGlobal

`func (o *InvoiceIngresoInput) GetGlobal() InvoiceIngresoInputAllOfGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *InvoiceIngresoInput) GetGlobalOk() (*InvoiceIngresoInputAllOfGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *InvoiceIngresoInput) SetGlobal(v InvoiceIngresoInputAllOfGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *InvoiceIngresoInput) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetExport

`func (o *InvoiceIngresoInput) GetExport() string`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *InvoiceIngresoInput) GetExportOk() (*string, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *InvoiceIngresoInput) SetExport(v string)`

SetExport sets Export field to given value.

### HasExport

`func (o *InvoiceIngresoInput) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetComplements

`func (o *InvoiceIngresoInput) GetComplements() []CustomComplementInput`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *InvoiceIngresoInput) GetComplementsOk() (*[]CustomComplementInput, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *InvoiceIngresoInput) SetComplements(v []CustomComplementInput)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *InvoiceIngresoInput) HasComplements() bool`

HasComplements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


