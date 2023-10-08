# PaymentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentForm** | **string** | Código de la forma de pago según el [catálogo del SAT](#forma-de-pago). También puedes utilizar la constante &#x60;PaymentForm&#x60; incluída en nuestras librerías. | 
**RelatedDocuments** | [**[]PaymentInputRelatedDocumentsInner**](PaymentInputRelatedDocumentsInner.md) |  | 
**Currency** | Pointer to **string** | Código de la moneda, acorde al estándar [ISO 4217](https://es.wikipedia.org/wiki/ISO_4217). | [optional] [default to "MXN"]
**Exchange** | Pointer to **float32** | Tipo de cambio conforme a la moneda usada. Representa el número de pesos mexicanos que equivalen a una unidad de la divisa señalada en el atributo &#x60;currency&#x60;. | [optional] [default to 1]
**Date** | Pointer to **time.Time** | Fecha en que se recibió el pago. Sólo es necesario incluirla si el pago se efectuó en una fecha anterior a la emisión de este comprobante. No se permiten fechas futuras. | [optional] [default to "now"]
**NumOperacion** | Pointer to **string** | Número de cheque, de autorización, de referencia, clave de rastreo SPEI, línea de captura o algún número de referencia que permita identificar la operación correspondiente al pago efectuado. | [optional] 
**RfcEmisorCtaOrd** | Pointer to **string** | RFC de la entidad emisora de la cuenta de origen, es decir, la operadora, banco, institución financiera, emisor de monedero electrónico, etc. | [optional] 
**NomBancoOrdExt** | Pointer to **string** | Nombre del banco ordenante. | [optional] 
**CtaOrdenante** | Pointer to **string** | Número de cuenta con la que se realizó el pago. | [optional] 
**RfcEmisorCtaBen** | Pointer to **string** | RFC de la entidad de la cuenta operadora destino, es decir, la operadora, banco, institución financiera, emisor de monedero electrónico, etc. | [optional] 
**CtaBeneficiario** | Pointer to **string** | Número de cuenta donde se recibió el pago. | [optional] 
**TipoCadPago** | Pointer to **string** | Clave del tipo de cadena de pago que genera la entidad receptora del pago. Si existe este campo, es obligatorio registrar los campos &#x60;certPago&#x60;, &#x60;cadPago&#x60; y &#x60;selloPago&#x60;.  | [optional] 
**CertPago** | Pointer to **string** | Certificado que corresponde al pago, como una cadena de texto en formato base 64. | [optional] 
**CadPago** | Pointer to **string** | Cadena original del comprobante de pago generado por la entidad emisora de la cuenta beneficiaria. | [optional] 
**SelloPago** | Pointer to **string** | Sello digital que se asocie al pago expresado como una cadena de texto en formato base 64. | [optional] 

## Methods

### NewPaymentInput

`func NewPaymentInput(paymentForm string, relatedDocuments []PaymentInputRelatedDocumentsInner, ) *PaymentInput`

NewPaymentInput instantiates a new PaymentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInputWithDefaults

`func NewPaymentInputWithDefaults() *PaymentInput`

NewPaymentInputWithDefaults instantiates a new PaymentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentForm

`func (o *PaymentInput) GetPaymentForm() string`

GetPaymentForm returns the PaymentForm field if non-nil, zero value otherwise.

### GetPaymentFormOk

`func (o *PaymentInput) GetPaymentFormOk() (*string, bool)`

GetPaymentFormOk returns a tuple with the PaymentForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentForm

`func (o *PaymentInput) SetPaymentForm(v string)`

SetPaymentForm sets PaymentForm field to given value.


### GetRelatedDocuments

`func (o *PaymentInput) GetRelatedDocuments() []PaymentInputRelatedDocumentsInner`

GetRelatedDocuments returns the RelatedDocuments field if non-nil, zero value otherwise.

### GetRelatedDocumentsOk

`func (o *PaymentInput) GetRelatedDocumentsOk() (*[]PaymentInputRelatedDocumentsInner, bool)`

GetRelatedDocumentsOk returns a tuple with the RelatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDocuments

`func (o *PaymentInput) SetRelatedDocuments(v []PaymentInputRelatedDocumentsInner)`

SetRelatedDocuments sets RelatedDocuments field to given value.


### GetCurrency

`func (o *PaymentInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *PaymentInput) GetExchange() float32`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *PaymentInput) GetExchangeOk() (*float32, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *PaymentInput) SetExchange(v float32)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *PaymentInput) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetDate

`func (o *PaymentInput) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PaymentInput) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PaymentInput) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *PaymentInput) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNumOperacion

`func (o *PaymentInput) GetNumOperacion() string`

GetNumOperacion returns the NumOperacion field if non-nil, zero value otherwise.

### GetNumOperacionOk

`func (o *PaymentInput) GetNumOperacionOk() (*string, bool)`

GetNumOperacionOk returns a tuple with the NumOperacion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOperacion

`func (o *PaymentInput) SetNumOperacion(v string)`

SetNumOperacion sets NumOperacion field to given value.

### HasNumOperacion

`func (o *PaymentInput) HasNumOperacion() bool`

HasNumOperacion returns a boolean if a field has been set.

### GetRfcEmisorCtaOrd

`func (o *PaymentInput) GetRfcEmisorCtaOrd() string`

GetRfcEmisorCtaOrd returns the RfcEmisorCtaOrd field if non-nil, zero value otherwise.

### GetRfcEmisorCtaOrdOk

`func (o *PaymentInput) GetRfcEmisorCtaOrdOk() (*string, bool)`

GetRfcEmisorCtaOrdOk returns a tuple with the RfcEmisorCtaOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEmisorCtaOrd

`func (o *PaymentInput) SetRfcEmisorCtaOrd(v string)`

SetRfcEmisorCtaOrd sets RfcEmisorCtaOrd field to given value.

### HasRfcEmisorCtaOrd

`func (o *PaymentInput) HasRfcEmisorCtaOrd() bool`

HasRfcEmisorCtaOrd returns a boolean if a field has been set.

### GetNomBancoOrdExt

`func (o *PaymentInput) GetNomBancoOrdExt() string`

GetNomBancoOrdExt returns the NomBancoOrdExt field if non-nil, zero value otherwise.

### GetNomBancoOrdExtOk

`func (o *PaymentInput) GetNomBancoOrdExtOk() (*string, bool)`

GetNomBancoOrdExtOk returns a tuple with the NomBancoOrdExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomBancoOrdExt

`func (o *PaymentInput) SetNomBancoOrdExt(v string)`

SetNomBancoOrdExt sets NomBancoOrdExt field to given value.

### HasNomBancoOrdExt

`func (o *PaymentInput) HasNomBancoOrdExt() bool`

HasNomBancoOrdExt returns a boolean if a field has been set.

### GetCtaOrdenante

`func (o *PaymentInput) GetCtaOrdenante() string`

GetCtaOrdenante returns the CtaOrdenante field if non-nil, zero value otherwise.

### GetCtaOrdenanteOk

`func (o *PaymentInput) GetCtaOrdenanteOk() (*string, bool)`

GetCtaOrdenanteOk returns a tuple with the CtaOrdenante field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaOrdenante

`func (o *PaymentInput) SetCtaOrdenante(v string)`

SetCtaOrdenante sets CtaOrdenante field to given value.

### HasCtaOrdenante

`func (o *PaymentInput) HasCtaOrdenante() bool`

HasCtaOrdenante returns a boolean if a field has been set.

### GetRfcEmisorCtaBen

`func (o *PaymentInput) GetRfcEmisorCtaBen() string`

GetRfcEmisorCtaBen returns the RfcEmisorCtaBen field if non-nil, zero value otherwise.

### GetRfcEmisorCtaBenOk

`func (o *PaymentInput) GetRfcEmisorCtaBenOk() (*string, bool)`

GetRfcEmisorCtaBenOk returns a tuple with the RfcEmisorCtaBen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfcEmisorCtaBen

`func (o *PaymentInput) SetRfcEmisorCtaBen(v string)`

SetRfcEmisorCtaBen sets RfcEmisorCtaBen field to given value.

### HasRfcEmisorCtaBen

`func (o *PaymentInput) HasRfcEmisorCtaBen() bool`

HasRfcEmisorCtaBen returns a boolean if a field has been set.

### GetCtaBeneficiario

`func (o *PaymentInput) GetCtaBeneficiario() string`

GetCtaBeneficiario returns the CtaBeneficiario field if non-nil, zero value otherwise.

### GetCtaBeneficiarioOk

`func (o *PaymentInput) GetCtaBeneficiarioOk() (*string, bool)`

GetCtaBeneficiarioOk returns a tuple with the CtaBeneficiario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaBeneficiario

`func (o *PaymentInput) SetCtaBeneficiario(v string)`

SetCtaBeneficiario sets CtaBeneficiario field to given value.

### HasCtaBeneficiario

`func (o *PaymentInput) HasCtaBeneficiario() bool`

HasCtaBeneficiario returns a boolean if a field has been set.

### GetTipoCadPago

`func (o *PaymentInput) GetTipoCadPago() string`

GetTipoCadPago returns the TipoCadPago field if non-nil, zero value otherwise.

### GetTipoCadPagoOk

`func (o *PaymentInput) GetTipoCadPagoOk() (*string, bool)`

GetTipoCadPagoOk returns a tuple with the TipoCadPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipoCadPago

`func (o *PaymentInput) SetTipoCadPago(v string)`

SetTipoCadPago sets TipoCadPago field to given value.

### HasTipoCadPago

`func (o *PaymentInput) HasTipoCadPago() bool`

HasTipoCadPago returns a boolean if a field has been set.

### GetCertPago

`func (o *PaymentInput) GetCertPago() string`

GetCertPago returns the CertPago field if non-nil, zero value otherwise.

### GetCertPagoOk

`func (o *PaymentInput) GetCertPagoOk() (*string, bool)`

GetCertPagoOk returns a tuple with the CertPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPago

`func (o *PaymentInput) SetCertPago(v string)`

SetCertPago sets CertPago field to given value.

### HasCertPago

`func (o *PaymentInput) HasCertPago() bool`

HasCertPago returns a boolean if a field has been set.

### GetCadPago

`func (o *PaymentInput) GetCadPago() string`

GetCadPago returns the CadPago field if non-nil, zero value otherwise.

### GetCadPagoOk

`func (o *PaymentInput) GetCadPagoOk() (*string, bool)`

GetCadPagoOk returns a tuple with the CadPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCadPago

`func (o *PaymentInput) SetCadPago(v string)`

SetCadPago sets CadPago field to given value.

### HasCadPago

`func (o *PaymentInput) HasCadPago() bool`

HasCadPago returns a boolean if a field has been set.

### GetSelloPago

`func (o *PaymentInput) GetSelloPago() string`

GetSelloPago returns the SelloPago field if non-nil, zero value otherwise.

### GetSelloPagoOk

`func (o *PaymentInput) GetSelloPagoOk() (*string, bool)`

GetSelloPagoOk returns a tuple with the SelloPago field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelloPago

`func (o *PaymentInput) SetSelloPago(v string)`

SetSelloPago sets SelloPago field to given value.

### HasSelloPago

`func (o *PaymentInput) HasSelloPago() bool`

HasSelloPago returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


