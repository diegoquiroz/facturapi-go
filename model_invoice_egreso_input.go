/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi-go

import (
	"encoding/json"
	"time"
)

// checks if the InvoiceEgresoInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceEgresoInput{}

// InvoiceEgresoInput struct for InvoiceEgresoInput
type InvoiceEgresoInput struct {
	Customer *InvoiceCommonInputPropertiesAllOfCustomer `json:"customer,omitempty"`
	// Fecha de expedición del comprobante en formato ISO8601 (UTC String). No puede ser anterior a 72 horas en el pasado, ni posterior al presente.
	Date *time.Time `json:"date,omitempty"`
	Address *InvoiceCommonInputPropertiesAllOfAddress `json:"address,omitempty"`
	// Identificador opcional que puedes usar para relacionar esta factura con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único.
	ExternalId *string `json:"external_id,omitempty"`
	// Identificador único que puedes usar para evitar duplicados al reintentar una petición. Puede ser cualquier cadena de texto, mientras sea única para cada documento. Si se deja en blanco, no se tomará en cuenta. 
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// Número de folio asignado por la empresa para control interno. Si se omite, se asignará el valor autoincremental de la organización.
	FolioNumber *int32 `json:"folio_number,omitempty"`
	// Serie. Caracteres designados por la empresa para control interno y sin validez fiscal.
	Series *string `json:"series,omitempty"`
	// En caso de que necesites incluir más información en el PDF, este campo te permite enviar código HTML con tu propio contenido.  Por seguridad, el código que puedes enviar está limitado a las siguientes etiquetas: `h1`, `h2`, `h3`, `h4`, `h5`, `h6`, `div`, `p`, `span`, `small`, `br`, `b`, `i`, `ul`, `ol`, `li`, `strong`, `table`, `thead`, `tbody`, `tfoot`, `tr`, `th` y `td`. No se permiten atributos ni estilos. 
	PdfCustomSection *string `json:"pdf_custom_section,omitempty"`
	// Código XML con la Addenda que se necesite agregar a la factura.
	Addenda *string `json:"addenda,omitempty"`
	Namespaces []InvoiceableCommonInputNamespacesInner `json:"namespaces,omitempty"`
	Type string `json:"type"`
	// Código que representa la forma de pago, de acuerdo al [catálogo del SAT](#forma-de-pago).
	PaymentForm string `json:"payment_form"`
	RelatedDocuments []RelatedDocumentInput `json:"related_documents,omitempty"`
	Items []LineItemEgresoInput `json:"items"`
	// Código de Uso CFDI según el catálogo del SAT. Puedes ver los códigos en [esta tabla](#uso-cfdi), o utilizar las constantes incluídas en nuestras librerías.
	Use *string `json:"use,omitempty"`
	// Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
	// Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos (MXN) que equivalen a una unidad de la divisa señalada en el atributo `currency`.
	Exchange *float32 `json:"exchange,omitempty"`
	Complements []CustomComplementInput `json:"complements,omitempty"`
}

// NewInvoiceEgresoInput instantiates a new InvoiceEgresoInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceEgresoInput(type_ string, paymentForm string, items []LineItemEgresoInput) *InvoiceEgresoInput {
	this := InvoiceEgresoInput{}
	var date time.Time = "now"
	this.Date = &date
	var folioNumber int32 = autoincremental
	this.FolioNumber = &folioNumber
	this.Type = type_
	this.PaymentForm = paymentForm
	this.Items = items
	var use string = "G01"
	this.Use = &use
	var currency string = "MXN"
	this.Currency = &currency
	var exchange float32 = 1
	this.Exchange = &exchange
	return &this
}

// NewInvoiceEgresoInputWithDefaults instantiates a new InvoiceEgresoInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceEgresoInputWithDefaults() *InvoiceEgresoInput {
	this := InvoiceEgresoInput{}
	var date time.Time = "now"
	this.Date = &date
	var folioNumber int32 = autoincremental
	this.FolioNumber = &folioNumber
	var use string = "G01"
	this.Use = &use
	var currency string = "MXN"
	this.Currency = &currency
	var exchange float32 = 1
	this.Exchange = &exchange
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetCustomer() InvoiceCommonInputPropertiesAllOfCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret InvoiceCommonInputPropertiesAllOfCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetCustomerOk() (*InvoiceCommonInputPropertiesAllOfCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given InvoiceCommonInputPropertiesAllOfCustomer and assigns it to the Customer field.
func (o *InvoiceEgresoInput) SetCustomer(v InvoiceCommonInputPropertiesAllOfCustomer) {
	o.Customer = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *InvoiceEgresoInput) SetDate(v time.Time) {
	o.Date = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetAddress() InvoiceCommonInputPropertiesAllOfAddress {
	if o == nil || IsNil(o.Address) {
		var ret InvoiceCommonInputPropertiesAllOfAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetAddressOk() (*InvoiceCommonInputPropertiesAllOfAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given InvoiceCommonInputPropertiesAllOfAddress and assigns it to the Address field.
func (o *InvoiceEgresoInput) SetAddress(v InvoiceCommonInputPropertiesAllOfAddress) {
	o.Address = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *InvoiceEgresoInput) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *InvoiceEgresoInput) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetFolioNumber() int32 {
	if o == nil || IsNil(o.FolioNumber) {
		var ret int32
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetFolioNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioNumber) {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasFolioNumber() bool {
	if o != nil && !IsNil(o.FolioNumber) {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given int32 and assigns it to the FolioNumber field.
func (o *InvoiceEgresoInput) SetFolioNumber(v int32) {
	o.FolioNumber = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *InvoiceEgresoInput) SetSeries(v string) {
	o.Series = &v
}

// GetPdfCustomSection returns the PdfCustomSection field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetPdfCustomSection() string {
	if o == nil || IsNil(o.PdfCustomSection) {
		var ret string
		return ret
	}
	return *o.PdfCustomSection
}

// GetPdfCustomSectionOk returns a tuple with the PdfCustomSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetPdfCustomSectionOk() (*string, bool) {
	if o == nil || IsNil(o.PdfCustomSection) {
		return nil, false
	}
	return o.PdfCustomSection, true
}

// HasPdfCustomSection returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasPdfCustomSection() bool {
	if o != nil && !IsNil(o.PdfCustomSection) {
		return true
	}

	return false
}

// SetPdfCustomSection gets a reference to the given string and assigns it to the PdfCustomSection field.
func (o *InvoiceEgresoInput) SetPdfCustomSection(v string) {
	o.PdfCustomSection = &v
}

// GetAddenda returns the Addenda field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetAddenda() string {
	if o == nil || IsNil(o.Addenda) {
		var ret string
		return ret
	}
	return *o.Addenda
}

// GetAddendaOk returns a tuple with the Addenda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetAddendaOk() (*string, bool) {
	if o == nil || IsNil(o.Addenda) {
		return nil, false
	}
	return o.Addenda, true
}

// HasAddenda returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasAddenda() bool {
	if o != nil && !IsNil(o.Addenda) {
		return true
	}

	return false
}

// SetAddenda gets a reference to the given string and assigns it to the Addenda field.
func (o *InvoiceEgresoInput) SetAddenda(v string) {
	o.Addenda = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetNamespaces() []InvoiceableCommonInputNamespacesInner {
	if o == nil || IsNil(o.Namespaces) {
		var ret []InvoiceableCommonInputNamespacesInner
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetNamespacesOk() ([]InvoiceableCommonInputNamespacesInner, bool) {
	if o == nil || IsNil(o.Namespaces) {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasNamespaces() bool {
	if o != nil && !IsNil(o.Namespaces) {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []InvoiceableCommonInputNamespacesInner and assigns it to the Namespaces field.
func (o *InvoiceEgresoInput) SetNamespaces(v []InvoiceableCommonInputNamespacesInner) {
	o.Namespaces = v
}

// GetType returns the Type field value
func (o *InvoiceEgresoInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InvoiceEgresoInput) SetType(v string) {
	o.Type = v
}

// GetPaymentForm returns the PaymentForm field value
func (o *InvoiceEgresoInput) GetPaymentForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentForm
}

// GetPaymentFormOk returns a tuple with the PaymentForm field value
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetPaymentFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentForm, true
}

// SetPaymentForm sets field value
func (o *InvoiceEgresoInput) SetPaymentForm(v string) {
	o.PaymentForm = v
}

// GetRelatedDocuments returns the RelatedDocuments field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetRelatedDocuments() []RelatedDocumentInput {
	if o == nil || IsNil(o.RelatedDocuments) {
		var ret []RelatedDocumentInput
		return ret
	}
	return o.RelatedDocuments
}

// GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetRelatedDocumentsOk() ([]RelatedDocumentInput, bool) {
	if o == nil || IsNil(o.RelatedDocuments) {
		return nil, false
	}
	return o.RelatedDocuments, true
}

// HasRelatedDocuments returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasRelatedDocuments() bool {
	if o != nil && !IsNil(o.RelatedDocuments) {
		return true
	}

	return false
}

// SetRelatedDocuments gets a reference to the given []RelatedDocumentInput and assigns it to the RelatedDocuments field.
func (o *InvoiceEgresoInput) SetRelatedDocuments(v []RelatedDocumentInput) {
	o.RelatedDocuments = v
}

// GetItems returns the Items field value
func (o *InvoiceEgresoInput) GetItems() []LineItemEgresoInput {
	if o == nil {
		var ret []LineItemEgresoInput
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetItemsOk() ([]LineItemEgresoInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *InvoiceEgresoInput) SetItems(v []LineItemEgresoInput) {
	o.Items = v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *InvoiceEgresoInput) SetUse(v string) {
	o.Use = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InvoiceEgresoInput) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetExchange() float32 {
	if o == nil || IsNil(o.Exchange) {
		var ret float32
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetExchangeOk() (*float32, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given float32 and assigns it to the Exchange field.
func (o *InvoiceEgresoInput) SetExchange(v float32) {
	o.Exchange = &v
}

// GetComplements returns the Complements field value if set, zero value otherwise.
func (o *InvoiceEgresoInput) GetComplements() []CustomComplementInput {
	if o == nil || IsNil(o.Complements) {
		var ret []CustomComplementInput
		return ret
	}
	return o.Complements
}

// GetComplementsOk returns a tuple with the Complements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceEgresoInput) GetComplementsOk() ([]CustomComplementInput, bool) {
	if o == nil || IsNil(o.Complements) {
		return nil, false
	}
	return o.Complements, true
}

// HasComplements returns a boolean if a field has been set.
func (o *InvoiceEgresoInput) HasComplements() bool {
	if o != nil && !IsNil(o.Complements) {
		return true
	}

	return false
}

// SetComplements gets a reference to the given []CustomComplementInput and assigns it to the Complements field.
func (o *InvoiceEgresoInput) SetComplements(v []CustomComplementInput) {
	o.Complements = v
}

func (o InvoiceEgresoInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceEgresoInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if !IsNil(o.FolioNumber) {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
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
	toSerialize["type"] = o.Type
	toSerialize["payment_form"] = o.PaymentForm
	if !IsNil(o.RelatedDocuments) {
		toSerialize["related_documents"] = o.RelatedDocuments
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
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
	return toSerialize, nil
}

type NullableInvoiceEgresoInput struct {
	value *InvoiceEgresoInput
	isSet bool
}

func (v NullableInvoiceEgresoInput) Get() *InvoiceEgresoInput {
	return v.value
}

func (v *NullableInvoiceEgresoInput) Set(val *InvoiceEgresoInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceEgresoInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceEgresoInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceEgresoInput(val *InvoiceEgresoInput) *NullableInvoiceEgresoInput {
	return &NullableInvoiceEgresoInput{value: val, isSet: true}
}

func (v NullableInvoiceEgresoInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceEgresoInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


