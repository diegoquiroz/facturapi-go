/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi

import (
	"encoding/json"
	"time"
)

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoice{}

// Invoice struct for Invoice
type Invoice struct {
	// ID del objeto
	Id *string `json:"id,omitempty"`
	// Fecha de registro
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Si el valor es `true`, indica que el objeto fue creado en ambiente Live; o si es `false`, en ambiente Test.
	Livemode *bool `json:"livemode,omitempty"`
	// Estado actual de la factura. 
	Status *string `json:"status,omitempty"`
	// Estado actual de la solicitud de cancelación, en caso de haberla realizado. Puedes leer más a detalle en la sección de [Cancelar Factura](#operation/deleteInvoice)). 
	CancellationStatus *string `json:"cancellation_status,omitempty"`
	// Dirección URL para verificar el estado del CFDI en el portal del SAT. Este link es el mismo que aparece en el código QR, en el PDF de la factura.
	VerificationUrl *string `json:"verification_url,omitempty"`
	// Fecha de expedición del comprobante en formato ISO8601 (UTC String).
	Date *time.Time `json:"date,omitempty"`
	Address *InvoicePropertiesAddress `json:"address,omitempty"`
	// Tipo de comprobante. Puede tener los valores `\"I\"`: Ingreso, `\"P\"`: Pago, `\"E\"`: Egreso, `\"N\"`: Nómina, `\"T\"`: Traslado. 
	Type *string `json:"type,omitempty"`
	Customer *CuustomerInfo `json:"customer,omitempty"`
	// Monto total facturado.
	Total *float32 `json:"total,omitempty"`
	// Folio fiscal de la factura, asignado por el SAT.
	Uuid *string `json:"uuid,omitempty"`
	// Número de folio autoincremental para control interno y sin validez fiscal.
	FolioNumber *int32 `json:"folio_number,omitempty"`
	// Serie. Caracteres designados por la empresa para control interno y sin validez fiscal. En el PDF se imprime junto al //www.facturapi.io/img/logo.svg
	Series *string `json:"series,omitempty"`
	// Identificador que puedes usar para relacionar esta factura con tus registros para después buscar por este número.
	ExternalId *string `json:"external_id,omitempty"`
	// Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento.
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago).
	PaymentForm *string `json:"payment_form,omitempty"`
	Items []LineItem `json:"items,omitempty"`
	RelatedDocuments []RelatedDocument `json:"related_documents,omitempty"`
	// Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
	// Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo `currency`.
	Exchange *float32 `json:"exchange,omitempty"`
	Complements []NominaOrCustomComplementProperties `json:"complements,omitempty"`
	// En caso de que necesites incluir más información en el PDF, este campo te permite insertar código HTML con tu propio contenido.
	PdfCustomSection *string `json:"pdf_custom_section,omitempty"`
	// Código XML con la Addenda que se necesite agregar a la factura.
	Addenda *string `json:"addenda,omitempty"`
	Namespaces []NamespaceProperties `json:"namespaces,omitempty"`
	Stamp *Stamp `json:"stamp,omitempty"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice() *Invoice {
	this := Invoice{}
	var date time.Time = "now"
	this.Date = &date
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	var date time.Time = "now"
	this.Date = &date
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invoice) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invoice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invoice) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invoice) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invoice) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Invoice) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLivemode returns the Livemode field value if set, zero value otherwise.
func (o *Invoice) GetLivemode() bool {
	if o == nil || IsNil(o.Livemode) {
		var ret bool
		return ret
	}
	return *o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLivemodeOk() (*bool, bool) {
	if o == nil || IsNil(o.Livemode) {
		return nil, false
	}
	return o.Livemode, true
}

// HasLivemode returns a boolean if a field has been set.
func (o *Invoice) HasLivemode() bool {
	if o != nil && !IsNil(o.Livemode) {
		return true
	}

	return false
}

// SetLivemode gets a reference to the given bool and assigns it to the Livemode field.
func (o *Invoice) SetLivemode(v bool) {
	o.Livemode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Invoice) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Invoice) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Invoice) SetStatus(v string) {
	o.Status = &v
}

// GetCancellationStatus returns the CancellationStatus field value if set, zero value otherwise.
func (o *Invoice) GetCancellationStatus() string {
	if o == nil || IsNil(o.CancellationStatus) {
		var ret string
		return ret
	}
	return *o.CancellationStatus
}

// GetCancellationStatusOk returns a tuple with the CancellationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCancellationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CancellationStatus) {
		return nil, false
	}
	return o.CancellationStatus, true
}

// HasCancellationStatus returns a boolean if a field has been set.
func (o *Invoice) HasCancellationStatus() bool {
	if o != nil && !IsNil(o.CancellationStatus) {
		return true
	}

	return false
}

// SetCancellationStatus gets a reference to the given string and assigns it to the CancellationStatus field.
func (o *Invoice) SetCancellationStatus(v string) {
	o.CancellationStatus = &v
}

// GetVerificationUrl returns the VerificationUrl field value if set, zero value otherwise.
func (o *Invoice) GetVerificationUrl() string {
	if o == nil || IsNil(o.VerificationUrl) {
		var ret string
		return ret
	}
	return *o.VerificationUrl
}

// GetVerificationUrlOk returns a tuple with the VerificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetVerificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationUrl) {
		return nil, false
	}
	return o.VerificationUrl, true
}

// HasVerificationUrl returns a boolean if a field has been set.
func (o *Invoice) HasVerificationUrl() bool {
	if o != nil && !IsNil(o.VerificationUrl) {
		return true
	}

	return false
}

// SetVerificationUrl gets a reference to the given string and assigns it to the VerificationUrl field.
func (o *Invoice) SetVerificationUrl(v string) {
	o.VerificationUrl = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Invoice) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Invoice) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *Invoice) SetDate(v time.Time) {
	o.Date = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Invoice) GetAddress() InvoicePropertiesAddress {
	if o == nil || IsNil(o.Address) {
		var ret InvoicePropertiesAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAddressOk() (*InvoicePropertiesAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Invoice) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given InvoicePropertiesAddress and assigns it to the Address field.
func (o *Invoice) SetAddress(v InvoicePropertiesAddress) {
	o.Address = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Invoice) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Invoice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Invoice) SetType(v string) {
	o.Type = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Invoice) GetCustomer() CuustomerInfo {
	if o == nil || IsNil(o.Customer) {
		var ret CuustomerInfo
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomerOk() (*CuustomerInfo, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Invoice) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CuustomerInfo and assigns it to the Customer field.
func (o *Invoice) SetCustomer(v CuustomerInfo) {
	o.Customer = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Invoice) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Invoice) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *Invoice) SetTotal(v float32) {
	o.Total = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Invoice) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Invoice) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Invoice) SetUuid(v string) {
	o.Uuid = &v
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *Invoice) GetFolioNumber() int32 {
	if o == nil || IsNil(o.FolioNumber) {
		var ret int32
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetFolioNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioNumber) {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *Invoice) HasFolioNumber() bool {
	if o != nil && !IsNil(o.FolioNumber) {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given int32 and assigns it to the FolioNumber field.
func (o *Invoice) SetFolioNumber(v int32) {
	o.FolioNumber = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *Invoice) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *Invoice) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *Invoice) SetSeries(v string) {
	o.Series = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Invoice) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Invoice) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Invoice) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *Invoice) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *Invoice) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *Invoice) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetPaymentForm returns the PaymentForm field value if set, zero value otherwise.
func (o *Invoice) GetPaymentForm() string {
	if o == nil || IsNil(o.PaymentForm) {
		var ret string
		return ret
	}
	return *o.PaymentForm
}

// GetPaymentFormOk returns a tuple with the PaymentForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPaymentFormOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentForm) {
		return nil, false
	}
	return o.PaymentForm, true
}

// HasPaymentForm returns a boolean if a field has been set.
func (o *Invoice) HasPaymentForm() bool {
	if o != nil && !IsNil(o.PaymentForm) {
		return true
	}

	return false
}

// SetPaymentForm gets a reference to the given string and assigns it to the PaymentForm field.
func (o *Invoice) SetPaymentForm(v string) {
	o.PaymentForm = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Invoice) GetItems() []LineItem {
	if o == nil || IsNil(o.Items) {
		var ret []LineItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetItemsOk() ([]LineItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Invoice) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []LineItem and assigns it to the Items field.
func (o *Invoice) SetItems(v []LineItem) {
	o.Items = v
}

// GetRelatedDocuments returns the RelatedDocuments field value if set, zero value otherwise.
func (o *Invoice) GetRelatedDocuments() []RelatedDocument {
	if o == nil || IsNil(o.RelatedDocuments) {
		var ret []RelatedDocument
		return ret
	}
	return o.RelatedDocuments
}

// GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetRelatedDocumentsOk() ([]RelatedDocument, bool) {
	if o == nil || IsNil(o.RelatedDocuments) {
		return nil, false
	}
	return o.RelatedDocuments, true
}

// HasRelatedDocuments returns a boolean if a field has been set.
func (o *Invoice) HasRelatedDocuments() bool {
	if o != nil && !IsNil(o.RelatedDocuments) {
		return true
	}

	return false
}

// SetRelatedDocuments gets a reference to the given []RelatedDocument and assigns it to the RelatedDocuments field.
func (o *Invoice) SetRelatedDocuments(v []RelatedDocument) {
	o.RelatedDocuments = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Invoice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Invoice) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Invoice) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Invoice) GetExchange() float32 {
	if o == nil || IsNil(o.Exchange) {
		var ret float32
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetExchangeOk() (*float32, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Invoice) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given float32 and assigns it to the Exchange field.
func (o *Invoice) SetExchange(v float32) {
	o.Exchange = &v
}

// GetComplements returns the Complements field value if set, zero value otherwise.
func (o *Invoice) GetComplements() []NominaOrCustomComplementProperties {
	if o == nil || IsNil(o.Complements) {
		var ret []NominaOrCustomComplementProperties
		return ret
	}
	return o.Complements
}

// GetComplementsOk returns a tuple with the Complements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetComplementsOk() ([]NominaOrCustomComplementProperties, bool) {
	if o == nil || IsNil(o.Complements) {
		return nil, false
	}
	return o.Complements, true
}

// HasComplements returns a boolean if a field has been set.
func (o *Invoice) HasComplements() bool {
	if o != nil && !IsNil(o.Complements) {
		return true
	}

	return false
}

// SetComplements gets a reference to the given []NominaOrCustomComplementProperties and assigns it to the Complements field.
func (o *Invoice) SetComplements(v []NominaOrCustomComplementProperties) {
	o.Complements = v
}

// GetPdfCustomSection returns the PdfCustomSection field value if set, zero value otherwise.
func (o *Invoice) GetPdfCustomSection() string {
	if o == nil || IsNil(o.PdfCustomSection) {
		var ret string
		return ret
	}
	return *o.PdfCustomSection
}

// GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPdfCustomSectionOk() (*string, bool) {
	if o == nil || IsNil(o.PdfCustomSection) {
		return nil, false
	}
	return o.PdfCustomSection, true
}

// HasPdfCustomSection returns a boolean if a field has been set.
func (o *Invoice) HasPdfCustomSection() bool {
	if o != nil && !IsNil(o.PdfCustomSection) {
		return true
	}

	return false
}

// SetPdfCustomSection gets a reference to the given string and assigns it to the PdfCustomSection field.
func (o *Invoice) SetPdfCustomSection(v string) {
	o.PdfCustomSection = &v
}

// GetAddenda returns the Addenda field value if set, zero value otherwise.
func (o *Invoice) GetAddenda() string {
	if o == nil || IsNil(o.Addenda) {
		var ret string
		return ret
	}
	return *o.Addenda
}

// GetAddendaOk returns a tuple with the Addenda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAddendaOk() (*string, bool) {
	if o == nil || IsNil(o.Addenda) {
		return nil, false
	}
	return o.Addenda, true
}

// HasAddenda returns a boolean if a field has been set.
func (o *Invoice) HasAddenda() bool {
	if o != nil && !IsNil(o.Addenda) {
		return true
	}

	return false
}

// SetAddenda gets a reference to the given string and assigns it to the Addenda field.
func (o *Invoice) SetAddenda(v string) {
	o.Addenda = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *Invoice) GetNamespaces() []NamespaceProperties {
	if o == nil || IsNil(o.Namespaces) {
		var ret []NamespaceProperties
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetNamespacesOk() ([]NamespaceProperties, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Invoice) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []NamespaceProperties and assigns it to the Namespaces field.
func (o *Invoice) SetNamespaces(v []NamespaceProperties) {
	o.Namespaces = v
}

// GetStamp returns the Stamp field value if set, zero value otherwise.
func (o *Invoice) GetStamp() Stamp {
	if o == nil || IsNil(o.Stamp) {
		var ret Stamp
		return ret
	}
	return *o.Stamp
}

// GetStampOk returns a tuple with the Stamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetStampOk() (*Stamp, bool) {
	if o == nil || IsNil(o.Stamp) {
		return nil, false
	}
	return o.Stamp, true
}

// HasStamp returns a boolean if a field has been set.
func (o *Invoice) HasStamp() bool {
	if o != nil && !IsNil(o.Stamp) {
		return true
	}

	return false
}

// SetStamp gets a reference to the given Stamp and assigns it to the Stamp field.
func (o *Invoice) SetStamp(v Stamp) {
	o.Stamp = &v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Livemode) {
		toSerialize["livemode"] = o.Livemode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CancellationStatus) {
		toSerialize["cancellation_status"] = o.CancellationStatus
	}
	if !IsNil(o.VerificationUrl) {
		toSerialize["verification_url"] = o.VerificationUrl
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.FolioNumber) {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if !IsNil(o.PaymentForm) {
		toSerialize["payment_form"] = o.PaymentForm
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.RelatedDocuments) {
		toSerialize["related_documents"] = o.RelatedDocuments
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.Complements) {
		toSerialize["complements"] = o.Complements
	}
	if !IsNil(o.PdfCustomSection) {
		toSerialize["pdf_custom_section"] = o.PdfCustomSection
	}
	if !IsNil(o.Addenda) {
		toSerialize["addenda"] = o.Addenda
	}
	if !IsNil(o.Namespaces) {
		toSerialize["namespaces"] = o.Namespaces
	}
	if !IsNil(o.Stamp) {
		toSerialize["stamp"] = o.Stamp
	}
	return toSerialize, nil
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


