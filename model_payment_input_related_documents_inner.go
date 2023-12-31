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
)

// checks if the PaymentInputRelatedDocumentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentInputRelatedDocumentsInner{}

// PaymentInputRelatedDocumentsInner struct for PaymentInputRelatedDocumentsInner
type PaymentInputRelatedDocumentsInner struct {
	// Folio fiscal ó UUID del comprobante de ingreso relacionado.
	Uuid string `json:"uuid"`
	// Cantidad del pago correspondiente al comprobante relacionado, usando el método de pago indicado en este elemento del arreglo de pagos. Este valor debe ser expresado en la moneda definida en `related_documents[].currency`. 
	Amount float32 `json:"amount"`
	Taxes []PaymentInputRelatedDocumentsInnerTaxesInner `json:"taxes"`
	// Código que representa si el bien o servicio es objeto de impuesto o no. Este atributo corresponde al campo \"ObjetoImp\" en el CFDI.  - `01`: No objeto de impuesto. - `02`: Sí objeto de impuesto. - `03`: Sí objeto de impuesto, pero no obligado a desglose. - `04`: Sí objeto de impuesto, y no causa impuesto. 
	Taxability *string `json:"taxability,omitempty"`
	// Número de parcialidad del pago.
	Installment int32 `json:"installment"`
	// Cantidad que estaba pendiente por pagar antes de recibir este pago. Este valor se expresa en la moneda definida en `payments[].related[].currency`.
	LastBalance float32 `json:"last_balance"`
	// Si la moneda utilizada en la factura relacionada no es moneda nacional (MXN), debe especificarse su valor acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
	// Obligatorio cuando la moneda del documento relacionado es distinta a la moneda de pago. Tipo de cambio entre las dos monedas al momento del pago. Ejemplo: La factura de iingreso relacionada se registra en USD, mientras que el pago actual se realiza en MXN, este atributo debería registrarse como `0.45` (USD/MXN). 
	Exchange *float32 `json:"exchange,omitempty"`
	// Opcionalmente se puede incluir el número de folio del documento relacionado.
	FolioNumber *int32 `json:"folio_number,omitempty"`
	// Opcionalmente se puede incluir la serie del documento relacionado.
	Series *string `json:"series,omitempty"`
}

// NewPaymentInputRelatedDocumentsInner instantiates a new PaymentInputRelatedDocumentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInputRelatedDocumentsInner(uuid string, amount float32, taxes []PaymentInputRelatedDocumentsInnerTaxesInner, installment int32, lastBalance float32) *PaymentInputRelatedDocumentsInner {
	this := PaymentInputRelatedDocumentsInner{}
	this.Uuid = uuid
	this.Amount = amount
	this.Taxes = taxes
	var taxability string = "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "
	this.Taxability = &taxability
	this.Installment = installment
	this.LastBalance = lastBalance
	var currency string = "MXN"
	this.Currency = &currency
	return &this
}

// NewPaymentInputRelatedDocumentsInnerWithDefaults instantiates a new PaymentInputRelatedDocumentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInputRelatedDocumentsInnerWithDefaults() *PaymentInputRelatedDocumentsInner {
	this := PaymentInputRelatedDocumentsInner{}
	var taxability string = "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "
	this.Taxability = &taxability
	var currency string = "MXN"
	this.Currency = &currency
	return &this
}

// GetUuid returns the Uuid field value
func (o *PaymentInputRelatedDocumentsInner) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PaymentInputRelatedDocumentsInner) SetUuid(v string) {
	o.Uuid = v
}

// GetAmount returns the Amount field value
func (o *PaymentInputRelatedDocumentsInner) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentInputRelatedDocumentsInner) SetAmount(v float32) {
	o.Amount = v
}

// GetTaxes returns the Taxes field value
func (o *PaymentInputRelatedDocumentsInner) GetTaxes() []PaymentInputRelatedDocumentsInnerTaxesInner {
	if o == nil {
		var ret []PaymentInputRelatedDocumentsInnerTaxesInner
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetTaxesOk() ([]PaymentInputRelatedDocumentsInnerTaxesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Taxes, true
}

// SetTaxes sets field value
func (o *PaymentInputRelatedDocumentsInner) SetTaxes(v []PaymentInputRelatedDocumentsInnerTaxesInner) {
	o.Taxes = v
}

// GetTaxability returns the Taxability field value if set, zero value otherwise.
func (o *PaymentInputRelatedDocumentsInner) GetTaxability() string {
	if o == nil || IsNil(o.Taxability) {
		var ret string
		return ret
	}
	return *o.Taxability
}

// GetTaxabilityOk returns a tuple with the Taxability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetTaxabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Taxability) {
		return nil, false
	}
	return o.Taxability, true
}

// HasTaxability returns a boolean if a field has been set.
func (o *PaymentInputRelatedDocumentsInner) HasTaxability() bool {
	if o != nil && !IsNil(o.Taxability) {
		return true
	}

	return false
}

// SetTaxability gets a reference to the given string and assigns it to the Taxability field.
func (o *PaymentInputRelatedDocumentsInner) SetTaxability(v string) {
	o.Taxability = &v
}

// GetInstallment returns the Installment field value
func (o *PaymentInputRelatedDocumentsInner) GetInstallment() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Installment
}

// GetInstallmentOk returns a tuple with the Installment field value
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetInstallmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Installment, true
}

// SetInstallment sets field value
func (o *PaymentInputRelatedDocumentsInner) SetInstallment(v int32) {
	o.Installment = v
}

// GetLastBalance returns the LastBalance field value
func (o *PaymentInputRelatedDocumentsInner) GetLastBalance() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LastBalance
}

// GetLastBalanceOk returns a tuple with the LastBalance field value
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetLastBalanceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastBalance, true
}

// SetLastBalance sets field value
func (o *PaymentInputRelatedDocumentsInner) SetLastBalance(v float32) {
	o.LastBalance = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentInputRelatedDocumentsInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentInputRelatedDocumentsInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentInputRelatedDocumentsInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *PaymentInputRelatedDocumentsInner) GetExchange() float32 {
	if o == nil || IsNil(o.Exchange) {
		var ret float32
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetExchangeOk() (*float32, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *PaymentInputRelatedDocumentsInner) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given float32 and assigns it to the Exchange field.
func (o *PaymentInputRelatedDocumentsInner) SetExchange(v float32) {
	o.Exchange = &v
}

// GetFolioNumber returns the FolioNumber field value if set, zero value otherwise.
func (o *PaymentInputRelatedDocumentsInner) GetFolioNumber() int32 {
	if o == nil || IsNil(o.FolioNumber) {
		var ret int32
		return ret
	}
	return *o.FolioNumber
}

// GetFolioNumberOk returns a tuple with the FolioNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetFolioNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.FolioNumber) {
		return nil, false
	}
	return o.FolioNumber, true
}

// HasFolioNumber returns a boolean if a field has been set.
func (o *PaymentInputRelatedDocumentsInner) HasFolioNumber() bool {
	if o != nil && !IsNil(o.FolioNumber) {
		return true
	}

	return false
}

// SetFolioNumber gets a reference to the given int32 and assigns it to the FolioNumber field.
func (o *PaymentInputRelatedDocumentsInner) SetFolioNumber(v int32) {
	o.FolioNumber = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *PaymentInputRelatedDocumentsInner) GetSeries() string {
	if o == nil || IsNil(o.Series) {
		var ret string
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInputRelatedDocumentsInner) GetSeriesOk() (*string, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *PaymentInputRelatedDocumentsInner) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given string and assigns it to the Series field.
func (o *PaymentInputRelatedDocumentsInner) SetSeries(v string) {
	o.Series = &v
}

func (o PaymentInputRelatedDocumentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInputRelatedDocumentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["amount"] = o.Amount
	toSerialize["taxes"] = o.Taxes
	if !IsNil(o.Taxability) {
		toSerialize["taxability"] = o.Taxability
	}
	toSerialize["installment"] = o.Installment
	toSerialize["last_balance"] = o.LastBalance
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.FolioNumber) {
		toSerialize["folio_number"] = o.FolioNumber
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullablePaymentInputRelatedDocumentsInner struct {
	value *PaymentInputRelatedDocumentsInner
	isSet bool
}

func (v NullablePaymentInputRelatedDocumentsInner) Get() *PaymentInputRelatedDocumentsInner {
	return v.value
}

func (v *NullablePaymentInputRelatedDocumentsInner) Set(val *PaymentInputRelatedDocumentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInputRelatedDocumentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInputRelatedDocumentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInputRelatedDocumentsInner(val *PaymentInputRelatedDocumentsInner) *NullablePaymentInputRelatedDocumentsInner {
	return &NullablePaymentInputRelatedDocumentsInner{value: val, isSet: true}
}

func (v NullablePaymentInputRelatedDocumentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInputRelatedDocumentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


