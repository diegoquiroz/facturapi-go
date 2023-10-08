# RetentionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [**InvoiceCommonInputPropertiesAllOfCustomer**](InvoiceCommonInputPropertiesAllOfCustomer.md) |  | 
**CveRetenc** | **string** | Clave de la retención o información de pagos de acuerdo al [catálogo del SAT](#clave-de-retencion). | 
**FechaExp** | Pointer to **time.Time** | Fecha de expedición del comprobante en formato ISO8601 (UTC String). | [optional] 
**DescRetenc** | Pointer to **string** | Si la clave de la retención es “25” (Otro tipo de retenciones), este campo se usa para registrar la descripción de la retención. | [optional] 
**FolioInt** | Pointer to **string** | Identificador alfanumérico para control interno de la empresa y sin relevancia fiscal. | [optional] 
**Periodo** | [**RetentionInputPeriodo**](RetentionInputPeriodo.md) |  | 
**Totales** | [**RetentionInputTotales**](RetentionInputTotales.md) |  | 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar esta retención con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento.  | [optional] 
**Complements** | Pointer to **[]string** |  | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite insertar código HTML con tu propio contenido. | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]InvoiceableCommonInputNamespacesInner**](InvoiceableCommonInputNamespacesInner.md) |  | [optional] 

## Methods

### NewRetentionInput

`func NewRetentionInput(customer InvoiceCommonInputPropertiesAllOfCustomer, cveRetenc string, periodo RetentionInputPeriodo, totales RetentionInputTotales, ) *RetentionInput`

NewRetentionInput instantiates a new RetentionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionInputWithDefaults

`func NewRetentionInputWithDefaults() *RetentionInput`

NewRetentionInputWithDefaults instantiates a new RetentionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *RetentionInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *RetentionInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *RetentionInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer)`

SetCustomer sets Customer field to given value.


### GetCveRetenc

`func (o *RetentionInput) GetCveRetenc() string`

GetCveRetenc returns the CveRetenc field if non-nil, zero value otherwise.

### GetCveRetencOk

`func (o *RetentionInput) GetCveRetencOk() (*string, bool)`

GetCveRetencOk returns a tuple with the CveRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveRetenc

`func (o *RetentionInput) SetCveRetenc(v string)`

SetCveRetenc sets CveRetenc field to given value.


### GetFechaExp

`func (o *RetentionInput) GetFechaExp() time.Time`

GetFechaExp returns the FechaExp field if non-nil, zero value otherwise.

### GetFechaExpOk

`func (o *RetentionInput) GetFechaExpOk() (*time.Time, bool)`

GetFechaExpOk returns a tuple with the FechaExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaExp

`func (o *RetentionInput) SetFechaExp(v time.Time)`

SetFechaExp sets FechaExp field to given value.

### HasFechaExp

`func (o *RetentionInput) HasFechaExp() bool`

HasFechaExp returns a boolean if a field has been set.

### GetDescRetenc

`func (o *RetentionInput) GetDescRetenc() string`

GetDescRetenc returns the DescRetenc field if non-nil, zero value otherwise.

### GetDescRetencOk

`func (o *RetentionInput) GetDescRetencOk() (*string, bool)`

GetDescRetencOk returns a tuple with the DescRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescRetenc

`func (o *RetentionInput) SetDescRetenc(v string)`

SetDescRetenc sets DescRetenc field to given value.

### HasDescRetenc

`func (o *RetentionInput) HasDescRetenc() bool`

HasDescRetenc returns a boolean if a field has been set.

### GetFolioInt

`func (o *RetentionInput) GetFolioInt() string`

GetFolioInt returns the FolioInt field if non-nil, zero value otherwise.

### GetFolioIntOk

`func (o *RetentionInput) GetFolioIntOk() (*string, bool)`

GetFolioIntOk returns a tuple with the FolioInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioInt

`func (o *RetentionInput) SetFolioInt(v string)`

SetFolioInt sets FolioInt field to given value.

### HasFolioInt

`func (o *RetentionInput) HasFolioInt() bool`

HasFolioInt returns a boolean if a field has been set.

### GetPeriodo

`func (o *RetentionInput) GetPeriodo() RetentionInputPeriodo`

GetPeriodo returns the Periodo field if non-nil, zero value otherwise.

### GetPeriodoOk

`func (o *RetentionInput) GetPeriodoOk() (*RetentionInputPeriodo, bool)`

GetPeriodoOk returns a tuple with the Periodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodo

`func (o *RetentionInput) SetPeriodo(v RetentionInputPeriodo)`

SetPeriodo sets Periodo field to given value.


### GetTotales

`func (o *RetentionInput) GetTotales() RetentionInputTotales`

GetTotales returns the Totales field if non-nil, zero value otherwise.

### GetTotalesOk

`func (o *RetentionInput) GetTotalesOk() (*RetentionInputTotales, bool)`

GetTotalesOk returns a tuple with the Totales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotales

`func (o *RetentionInput) SetTotales(v RetentionInputTotales)`

SetTotales sets Totales field to given value.


### GetExternalId

`func (o *RetentionInput) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *RetentionInput) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *RetentionInput) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *RetentionInput) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *RetentionInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *RetentionInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *RetentionInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *RetentionInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetComplements

`func (o *RetentionInput) GetComplements() []string`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *RetentionInput) GetComplementsOk() (*[]string, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *RetentionInput) SetComplements(v []string)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *RetentionInput) HasComplements() bool`

HasComplements returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *RetentionInput) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *RetentionInput) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *RetentionInput) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *RetentionInput) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *RetentionInput) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *RetentionInput) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *RetentionInput) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *RetentionInput) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *RetentionInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *RetentionInput) GetNamespacesOk() (*[]InvoiceableCommonInputNamespacesInner, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *RetentionInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *RetentionInput) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


