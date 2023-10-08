# RetentionProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveRetenc** | Pointer to **string** | Clave de la retención o información de pagos de acuerdo al catálogo del SAT. | [optional] 
**FechaExp** | Pointer to **time.Time** | Fecha de expedición del comprobante en formato ISO8601 (UTC String). | [optional] 
**DescRetenc** | Pointer to **string** | Si la clave de la retención es “25” (Otro tipo de retenciones), este campo se usa para registrar la descripción de la retención. | [optional] 
**FolioInt** | Pointer to **string** | Identificador alfanumérico para control interno de la empresa y sin relevancia fiscal. | [optional] 
**Periodo** | Pointer to [**RetentionPropertiesPeriodo**](RetentionPropertiesPeriodo.md) |  | [optional] 
**Totales** | Pointer to [**RetentionPropertiesTotales**](RetentionPropertiesTotales.md) |  | [optional] 
**ExternalId** | Pointer to **string** | Identificador opcional que puedes usar para relacionar esta retención con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único. | [optional] 
**IdempotencyKey** | Pointer to **string** | Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta.  | [optional] 
**Complements** | Pointer to **[]string** |  | [optional] 
**PdfCustomSection** | Pointer to **string** | En caso de que necesites incluir más información en el PDF, este campo te permite insertar código HTML con tu propio contenido. | [optional] 
**Addenda** | Pointer to **string** | Código XML con la Addenda que se necesite agregar a la factura. | [optional] 
**Namespaces** | Pointer to [**[]NamespaceProperties**](NamespaceProperties.md) |  | [optional] 

## Methods

### NewRetentionProperties

`func NewRetentionProperties() *RetentionProperties`

NewRetentionProperties instantiates a new RetentionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionPropertiesWithDefaults

`func NewRetentionPropertiesWithDefaults() *RetentionProperties`

NewRetentionPropertiesWithDefaults instantiates a new RetentionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveRetenc

`func (o *RetentionProperties) GetCveRetenc() string`

GetCveRetenc returns the CveRetenc field if non-nil, zero value otherwise.

### GetCveRetencOk

`func (o *RetentionProperties) GetCveRetencOk() (*string, bool)`

GetCveRetencOk returns a tuple with the CveRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveRetenc

`func (o *RetentionProperties) SetCveRetenc(v string)`

SetCveRetenc sets CveRetenc field to given value.

### HasCveRetenc

`func (o *RetentionProperties) HasCveRetenc() bool`

HasCveRetenc returns a boolean if a field has been set.

### GetFechaExp

`func (o *RetentionProperties) GetFechaExp() time.Time`

GetFechaExp returns the FechaExp field if non-nil, zero value otherwise.

### GetFechaExpOk

`func (o *RetentionProperties) GetFechaExpOk() (*time.Time, bool)`

GetFechaExpOk returns a tuple with the FechaExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaExp

`func (o *RetentionProperties) SetFechaExp(v time.Time)`

SetFechaExp sets FechaExp field to given value.

### HasFechaExp

`func (o *RetentionProperties) HasFechaExp() bool`

HasFechaExp returns a boolean if a field has been set.

### GetDescRetenc

`func (o *RetentionProperties) GetDescRetenc() string`

GetDescRetenc returns the DescRetenc field if non-nil, zero value otherwise.

### GetDescRetencOk

`func (o *RetentionProperties) GetDescRetencOk() (*string, bool)`

GetDescRetencOk returns a tuple with the DescRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescRetenc

`func (o *RetentionProperties) SetDescRetenc(v string)`

SetDescRetenc sets DescRetenc field to given value.

### HasDescRetenc

`func (o *RetentionProperties) HasDescRetenc() bool`

HasDescRetenc returns a boolean if a field has been set.

### GetFolioInt

`func (o *RetentionProperties) GetFolioInt() string`

GetFolioInt returns the FolioInt field if non-nil, zero value otherwise.

### GetFolioIntOk

`func (o *RetentionProperties) GetFolioIntOk() (*string, bool)`

GetFolioIntOk returns a tuple with the FolioInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioInt

`func (o *RetentionProperties) SetFolioInt(v string)`

SetFolioInt sets FolioInt field to given value.

### HasFolioInt

`func (o *RetentionProperties) HasFolioInt() bool`

HasFolioInt returns a boolean if a field has been set.

### GetPeriodo

`func (o *RetentionProperties) GetPeriodo() RetentionPropertiesPeriodo`

GetPeriodo returns the Periodo field if non-nil, zero value otherwise.

### GetPeriodoOk

`func (o *RetentionProperties) GetPeriodoOk() (*RetentionPropertiesPeriodo, bool)`

GetPeriodoOk returns a tuple with the Periodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodo

`func (o *RetentionProperties) SetPeriodo(v RetentionPropertiesPeriodo)`

SetPeriodo sets Periodo field to given value.

### HasPeriodo

`func (o *RetentionProperties) HasPeriodo() bool`

HasPeriodo returns a boolean if a field has been set.

### GetTotales

`func (o *RetentionProperties) GetTotales() RetentionPropertiesTotales`

GetTotales returns the Totales field if non-nil, zero value otherwise.

### GetTotalesOk

`func (o *RetentionProperties) GetTotalesOk() (*RetentionPropertiesTotales, bool)`

GetTotalesOk returns a tuple with the Totales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotales

`func (o *RetentionProperties) SetTotales(v RetentionPropertiesTotales)`

SetTotales sets Totales field to given value.

### HasTotales

`func (o *RetentionProperties) HasTotales() bool`

HasTotales returns a boolean if a field has been set.

### GetExternalId

`func (o *RetentionProperties) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *RetentionProperties) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *RetentionProperties) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *RetentionProperties) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *RetentionProperties) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *RetentionProperties) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *RetentionProperties) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *RetentionProperties) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetComplements

`func (o *RetentionProperties) GetComplements() []string`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *RetentionProperties) GetComplementsOk() (*[]string, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *RetentionProperties) SetComplements(v []string)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *RetentionProperties) HasComplements() bool`

HasComplements returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *RetentionProperties) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *RetentionProperties) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *RetentionProperties) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *RetentionProperties) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *RetentionProperties) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *RetentionProperties) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *RetentionProperties) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *RetentionProperties) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *RetentionProperties) GetNamespaces() []NamespaceProperties`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *RetentionProperties) GetNamespacesOk() (*[]NamespaceProperties, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *RetentionProperties) SetNamespaces(v []NamespaceProperties)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *RetentionProperties) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


