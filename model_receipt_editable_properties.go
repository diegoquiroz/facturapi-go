/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ReceiptEditableProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiptEditableProperties{}

// ReceiptEditableProperties struct for ReceiptEditableProperties
type ReceiptEditableProperties struct {
	// Fecha de emisión del recibo. Por defecto se utiliza la fecha actual.
	Date *time.Time `json:"date,omitempty"`
	// Código que representa la forma de pago, según el [catálogo del SAT](#forma-de-pago).
	PaymentForm *string `json:"payment_form,omitempty"`
	// Autoincremental. Número de folio del recibo para control interno y sin validez fiscal.
	FolioNumber *int32 `json:"folio_number,omitempty"`
	// Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
	// Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo `currency`.
	Exchange *float32 `json:"exchange,omitempty"`
	// Nombre de la sucursal donde se expidió el recibo.
	Branch *string `json:"branch,omitempty"`
	// Identificador opcional que puedes usar para relacionar este recibo con tus registros y poder hacer búsquedas usando este identificador. Facturapi no valida que este campo sea único.
	ExternalId *string `json:"external_id,omitempty"`
}

// NewReceiptEditableProperties instantiates a new ReceiptEditableProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptEditableProperties() *ReceiptEditableProperties {
	this := ReceiptEditableProperties{}
	return &this
}

// NewReceiptEditablePropertiesWithDefaults instantiates a new ReceiptEditableProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptEditablePropertiesWithDefaults() *ReceiptEditableProperties {
	this := ReceiptEditableProperties{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *ReceiptEditableProperties) SetDate(v time.Time) {
	o.Date = &v
}

// GetPaymentForm returns the PaymentForm field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetPaymentForm() string {
	if o == nil || IsNil(o.PaymentForm) {
		var ret string
		return ret
	}
	return *o.PaymentForm
}

// GetPaymentFormOk returns a tuple with the PaymentForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetPaymentFormOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentForm) {
		return nil, false
	}
	return o.PaymentForm, true
}

// HasPaymentForm returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasPaymentForm() bool {
	if o != nil && !IsNil(o.PaymentForm) {
		return true
	}

	return false
}

// SetPaymentForm gets a reference to the given string and assigns it to the PaymentForm field.
func (o *ReceiptEditableProperties) SetPaymentForm(v string) {
	o.PaymentForm = &v
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetFolioNumber() int32 {
	if o == nil || IsNil(o.FolioNumber) {
		var ret int32
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetFolioNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioNumber) {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasFolioNumber() bool {
	if o != nil && !IsNil(o.FolioNumber) {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given int32 and assigns it to the FolioNumber field.
func (o *ReceiptEditableProperties) SetFolioNumber(v int32) {
	o.FolioNumber = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ReceiptEditableProperties) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetExchange() float32 {
	if o == nil || IsNil(o.Exchange) {
		var ret float32
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetExchangeOk() (*float32, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given float32 and assigns it to the Exchange field.
func (o *ReceiptEditableProperties) SetExchange(v float32) {
	o.Exchange = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ReceiptEditableProperties) SetBranch(v string) {
	o.Branch = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ReceiptEditableProperties) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiptEditableProperties) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ReceiptEditableProperties) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ReceiptEditableProperties) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o ReceiptEditableProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptEditableProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.PaymentForm) {
		toSerialize["payment_form"] = o.PaymentForm
	}
	if !IsNil(o.FolioNumber) {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableReceiptEditableProperties struct {
	value *ReceiptEditableProperties
	isSet bool
}

func (v NullableReceiptEditableProperties) Get() *ReceiptEditableProperties {
	return v.value
}

func (v *NullableReceiptEditableProperties) Set(val *ReceiptEditableProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptEditableProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptEditableProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptEditableProperties(val *ReceiptEditableProperties) *NullableReceiptEditableProperties {
	return &NullableReceiptEditableProperties{value: val, isSet: true}
}

func (v NullableReceiptEditableProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptEditableProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


