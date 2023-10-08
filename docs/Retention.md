# Retention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID del objeto | [optional] 
**CreatedAt** | Pointer to **time.Time** | Fecha de registro | [optional] 
**Livemode** | Pointer to **bool** | Si el valor es &#x60;true&#x60;, indica que el objeto fue creado en ambiente Live; o si es &#x60;false&#x60;, en ambiente Test. | [optional] 
**Status** | Pointer to **string** | Estado actual de la retención.  | [optional] 
**VerificationUrl** | Pointer to **string** | Dirección URL para verificar el estado de la retención en el portal del SAT. Este link es el mismo que aparece en el código QR, en el PDF de la retención. | [optional] 
**Type** | Pointer to **string** | Tipo de comprobante. | [optional] 
**Uuid** | Pointer to **string** | Folio fiscal de la retención, asignado por el SAT. | [optional] 
**Stamp** | Pointer to [**Stamp**](Stamp.md) |  | [optional] 
**Customer** | Pointer to [**CuustomerInfo**](CuustomerInfo.md) |  | [optional] 
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

### NewRetention

`func NewRetention() *Retention`

NewRetention instantiates a new Retention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionWithDefaults

`func NewRetentionWithDefaults() *Retention`

NewRetentionWithDefaults instantiates a new Retention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Retention) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Retention) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Retention) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Retention) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Retention) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Retention) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Retention) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Retention) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLivemode

`func (o *Retention) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Retention) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Retention) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *Retention) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetStatus

`func (o *Retention) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Retention) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Retention) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Retention) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationUrl

`func (o *Retention) GetVerificationUrl() string`

GetVerificationUrl returns the VerificationUrl field if non-nil, zero value otherwise.

### GetVerificationUrlOk

`func (o *Retention) GetVerificationUrlOk() (*string, bool)`

GetVerificationUrlOk returns a tuple with the VerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUrl

`func (o *Retention) SetVerificationUrl(v string)`

SetVerificationUrl sets VerificationUrl field to given value.

### HasVerificationUrl

`func (o *Retention) HasVerificationUrl() bool`

HasVerificationUrl returns a boolean if a field has been set.

### GetType

`func (o *Retention) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Retention) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Retention) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Retention) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *Retention) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Retention) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Retention) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Retention) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStamp

`func (o *Retention) GetStamp() Stamp`

GetStamp returns the Stamp field if non-nil, zero value otherwise.

### GetStampOk

`func (o *Retention) GetStampOk() (*Stamp, bool)`

GetStampOk returns a tuple with the Stamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStamp

`func (o *Retention) SetStamp(v Stamp)`

SetStamp sets Stamp field to given value.

### HasStamp

`func (o *Retention) HasStamp() bool`

HasStamp returns a boolean if a field has been set.

### GetCustomer

`func (o *Retention) GetCustomer() CuustomerInfo`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Retention) GetCustomerOk() (*CuustomerInfo, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Retention) SetCustomer(v CuustomerInfo)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Retention) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCveRetenc

`func (o *Retention) GetCveRetenc() string`

GetCveRetenc returns the CveRetenc field if non-nil, zero value otherwise.

### GetCveRetencOk

`func (o *Retention) GetCveRetencOk() (*string, bool)`

GetCveRetencOk returns a tuple with the CveRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveRetenc

`func (o *Retention) SetCveRetenc(v string)`

SetCveRetenc sets CveRetenc field to given value.

### HasCveRetenc

`func (o *Retention) HasCveRetenc() bool`

HasCveRetenc returns a boolean if a field has been set.

### GetFechaExp

`func (o *Retention) GetFechaExp() time.Time`

GetFechaExp returns the FechaExp field if non-nil, zero value otherwise.

### GetFechaExpOk

`func (o *Retention) GetFechaExpOk() (*time.Time, bool)`

GetFechaExpOk returns a tuple with the FechaExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFechaExp

`func (o *Retention) SetFechaExp(v time.Time)`

SetFechaExp sets FechaExp field to given value.

### HasFechaExp

`func (o *Retention) HasFechaExp() bool`

HasFechaExp returns a boolean if a field has been set.

### GetDescRetenc

`func (o *Retention) GetDescRetenc() string`

GetDescRetenc returns the DescRetenc field if non-nil, zero value otherwise.

### GetDescRetencOk

`func (o *Retention) GetDescRetencOk() (*string, bool)`

GetDescRetencOk returns a tuple with the DescRetenc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescRetenc

`func (o *Retention) SetDescRetenc(v string)`

SetDescRetenc sets DescRetenc field to given value.

### HasDescRetenc

`func (o *Retention) HasDescRetenc() bool`

HasDescRetenc returns a boolean if a field has been set.

### GetFolioInt

`func (o *Retention) GetFolioInt() string`

GetFolioInt returns the FolioInt field if non-nil, zero value otherwise.

### GetFolioIntOk

`func (o *Retention) GetFolioIntOk() (*string, bool)`

GetFolioIntOk returns a tuple with the FolioInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolioInt

`func (o *Retention) SetFolioInt(v string)`

SetFolioInt sets FolioInt field to given value.

### HasFolioInt

`func (o *Retention) HasFolioInt() bool`

HasFolioInt returns a boolean if a field has been set.

### GetPeriodo

`func (o *Retention) GetPeriodo() RetentionPropertiesPeriodo`

GetPeriodo returns the Periodo field if non-nil, zero value otherwise.

### GetPeriodoOk

`func (o *Retention) GetPeriodoOk() (*RetentionPropertiesPeriodo, bool)`

GetPeriodoOk returns a tuple with the Periodo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodo

`func (o *Retention) SetPeriodo(v RetentionPropertiesPeriodo)`

SetPeriodo sets Periodo field to given value.

### HasPeriodo

`func (o *Retention) HasPeriodo() bool`

HasPeriodo returns a boolean if a field has been set.

### GetTotales

`func (o *Retention) GetTotales() RetentionPropertiesTotales`

GetTotales returns the Totales field if non-nil, zero value otherwise.

### GetTotalesOk

`func (o *Retention) GetTotalesOk() (*RetentionPropertiesTotales, bool)`

GetTotalesOk returns a tuple with the Totales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotales

`func (o *Retention) SetTotales(v RetentionPropertiesTotales)`

SetTotales sets Totales field to given value.

### HasTotales

`func (o *Retention) HasTotales() bool`

HasTotales returns a boolean if a field has been set.

### GetExternalId

`func (o *Retention) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Retention) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Retention) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Retention) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *Retention) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *Retention) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *Retention) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *Retention) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetComplements

`func (o *Retention) GetComplements() []string`

GetComplements returns the Complements field if non-nil, zero value otherwise.

### GetComplementsOk

`func (o *Retention) GetComplementsOk() (*[]string, bool)`

GetComplementsOk returns a tuple with the Complements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplements

`func (o *Retention) SetComplements(v []string)`

SetComplements sets Complements field to given value.

### HasComplements

`func (o *Retention) HasComplements() bool`

HasComplements returns a boolean if a field has been set.

### GetPdfCustomSection

`func (o *Retention) GetPdfCustomSection() string`

GetPdfCustomSection returns the PdfCustomSection field if non-nil, zero value otherwise.

### GetPdfCustomSectionOk

`func (o *Retention) GetPdfCustomSectionOk() (*string, bool)`

GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfCustomSection

`func (o *Retention) SetPdfCustomSection(v string)`

SetPdfCustomSection sets PdfCustomSection field to given value.

### HasPdfCustomSection

`func (o *Retention) HasPdfCustomSection() bool`

HasPdfCustomSection returns a boolean if a field has been set.

### GetAddenda

`func (o *Retention) GetAddenda() string`

GetAddenda returns the Addenda field if non-nil, zero value otherwise.

### GetAddendaOk

`func (o *Retention) GetAddendaOk() (*string, bool)`

GetAddendaOk returns a tuple with the Addenda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddenda

`func (o *Retention) SetAddenda(v string)`

SetAddenda sets Addenda field to given value.

### HasAddenda

`func (o *Retention) HasAddenda() bool`

HasAddenda returns a boolean if a field has been set.

### GetNamespaces

`func (o *Retention) GetNamespaces() []NamespaceProperties`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *Retention) GetNamespacesOk() (*[]NamespaceProperties, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *Retention) SetNamespaces(v []NamespaceProperties)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *Retention) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


