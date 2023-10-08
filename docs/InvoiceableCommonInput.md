# InvoiceableCommonInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolioNumber** | Pointer to **int32** | Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización. | [optional] [default to autoincremental]
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: &#x60;h1&#x60;, &#x60;h2&#x60;, &#x60;h3&#x60;, &#x60;h4&#x60;, &#x60;h5&#x60;, &#x60;h6&#x60;, &#x60;div&#x60;, &#x60;p&#x60;, &#x60;span&#x60;, &#x60;small&#x60;, &#x60;br&#x60;, &#x60;b&#x60;, &#x60;i&#x60;, &#x60;ul&#x60;, &#x60;ol&#x60;, &#x60;li&#x60;, &#x60;strong&#x60;, &#x60;table&#x60;, &#x60;thead&#x60;, &#x60;tbody&#x60;, &#x60;tfoot&#x60;, &#x60;tr&#x60;, &#x60;th&#x60; y &#x60;td&#x60;. No se permiten atributos ni estilos.  | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]InvoiceableCommonInputNamespacesInner**](InvoiceableCommonInputNamespacesInner.md) |  | [optional] 

## Methods

### NewInvoiceableCommonInput

`func NewInvoiceableCommonInput() *InvoiceableCommonInput`

NewInvoiceableCommonInput instantiates a new InvoiceableCommonInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceableCommonInputWithDefaults

`func NewInvoiceableCommonInputWithDefaults() *InvoiceableCommonInput`

NewInvoiceableCommonInputWithDefaults instantiates a new InvoiceableCommonInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolioNumber

`func (o *InvoiceableCommonInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *InvoiceableCommonInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *InvoiceableCommonInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *InvoiceableCommonInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *InvoiceableCommonInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InvoiceableCommonInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InvoiceableCommonInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InvoiceableCommonInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *InvoiceableCommonInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *InvoiceableCommonInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *InvoiceableCommonInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *InvoiceableCommonInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *InvoiceableCommonInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *InvoiceableCommonInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *InvoiceableCommonInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *InvoiceableCommonInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *InvoiceableCommonInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InvoiceableCommonInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InvoiceableCommonInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InvoiceableCommonInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


