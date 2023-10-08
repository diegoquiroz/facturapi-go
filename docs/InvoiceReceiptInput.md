# InvoiceReceiptInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolioNumber** | Pointer to **int32** | Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización. | [optional] [default to autoincremental]
**Series** | Pointer to **string** | Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: &#x60;h1&#x60;, &#x60;h2&#x60;, &#x60;h3&#x60;, &#x60;h4&#x60;, &#x60;h5&#x60;, &#x60;h6&#x60;, &#x60;div&#x60;, &#x60;p&#x60;, &#x60;span&#x60;, &#x60;small&#x60;, &#x60;br&#x60;, &#x60;b&#x60;, &#x60;i&#x60;, &#x60;ul&#x60;, &#x60;ol&#x60;, &#x60;li&#x60;, &#x60;strong&#x60;, &#x60;table&#x60;, &#x60;thead&#x60;, &#x60;tbody&#x60;, &#x60;tfoot&#x60;, &#x60;tr&#x60;, &#x60;th&#x60; y &#x60;td&#x60;. No se permiten atributos ni estilos.  | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]InvoiceableCommonInputNamespacesInner**](InvoiceableCommonInputNamespacesInner.md) |  | [optional] 
**Customer** | [**InvoiceCommonInputPropertiesAllOfCustomer**](InvoiceCommonInputPropertiesAllOfCustomer.md) |  | 
**Use** | Pointer to **string** | Código de Uso CFDI según el catálogo del SAT. Puedes ver los códigos en [esta tabla](#uso-cfdi), o utilizar las constantes incluídas en nuestras librerías. | [optional] [default to "G01"]
**Conditions** | Pointer to **string** | Condiciones de pago | [optional] 

## Methods

### NewInvoiceReceiptInput

`func NewInvoiceReceiptInput(customer InvoiceCommonInputPropertiesAllOfCustomer, ) *InvoiceReceiptInput`

NewInvoiceReceiptInput instantiates a new InvoiceReceiptInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReceiptInputWithDefaults

`func NewInvoiceReceiptInputWithDefaults() *InvoiceReceiptInput`

NewInvoiceReceiptInputWithDefaults instantiates a new InvoiceReceiptInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolioNumber

`func (o *InvoiceReceiptInput) GetFolioNumber() int32`

GetFolioNumber returns the FolioNumber field if non-nil, zero value otherwise.

### GetFolioNumberOk

`func (o *InvoiceReceiptInput) GetFolioNumberOk() (*int32, bool)`

GetFolioNumberOk returns a tuple with the FolioNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioNumber

`func (o *InvoiceReceiptInput) SetFolioNumber(v int32)`

SetFolioNumber sets FolioNumber field to given value.

### HasFolioNumber

`func (o *InvoiceReceiptInput) HasFolioNumber() bool`

HasFolioNumber returns a boolean if a field has been set.

### GetSeries

`func (o *InvoiceReceiptInput) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *InvoiceReceiptInput) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *InvoiceReceiptInput) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *InvoiceReceiptInput) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *InvoiceReceiptInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *InvoiceReceiptInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *InvoiceReceiptInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *InvoiceReceiptInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *InvoiceReceiptInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *InvoiceReceiptInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *InvoiceReceiptInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *InvoiceReceiptInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *InvoiceReceiptInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *InvoiceReceiptInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *InvoiceReceiptInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *InvoiceReceiptInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetCustomer

`func (o *InvoiceReceiptInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InvoiceReceiptInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InvoiceReceiptInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer)`

SetCustomer sets Customer field to given value.


### GetUse

`func (o *InvoiceReceiptInput) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *InvoiceReceiptInput) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *InvoiceReceiptInput) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *InvoiceReceiptInput) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetConditions

`func (o *InvoiceReceiptInput) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *InvoiceReceiptInput) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *InvoiceReceiptInput) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *InvoiceReceiptInput) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


