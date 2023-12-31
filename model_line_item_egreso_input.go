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

// checks if the LineItemEgresoInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemEgresoInput{}

// LineItemEgresoInput Conceptos incluidos en el documento
type LineItemEgresoInput struct {
	// Cantidad de unidades incluídas del mismo concepto.
	Quantity *float32 `json:"quantity,omitempty"`
	// Monto total de descuento aplicado a este concepto.
	Discount *float32 `json:"discount,omitempty"`
	Product LineItemEgresoInputProduct `json:"product"`
	Parts []PartInput `json:"parts,omitempty"`
	CustomsKeys []string `json:"customs_keys,omitempty"`
	// Código XML de tu complemento.
	Complement *string `json:"complement,omitempty"`
	ThirdParty *LineItemInputThirdParty `json:"third_party,omitempty"`
}

// NewLineItemEgresoInput instantiates a new LineItemEgresoInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemEgresoInput(product LineItemEgresoInputProduct) *LineItemEgresoInput {
	this := LineItemEgresoInput{}
	var quantity float32 = 1
	this.Quantity = &quantity
	var discount float32 = 0
	this.Discount = &discount
	this.Product = product
	return &this
}

// NewLineItemEgresoInputWithDefaults instantiates a new LineItemEgresoInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemEgresoInputWithDefaults() *LineItemEgresoInput {
	this := LineItemEgresoInput{}
	var quantity float32 = 1
	this.Quantity = &quantity
	var discount float32 = 0
	this.Discount = &discount
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *LineItemEgresoInput) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *LineItemEgresoInput) SetDiscount(v float32) {
	o.Discount = &v
}

// GetProduct returns the Product field value
func (o *LineItemEgresoInput) GetProduct() LineItemEgresoInputProduct {
	if o == nil {
		var ret LineItemEgresoInputProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetProductOk() (*LineItemEgresoInputProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *LineItemEgresoInput) SetProduct(v LineItemEgresoInputProduct) {
	o.Product = v
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetParts() []PartInput {
	if o == nil || IsNil(o.Parts) {
		var ret []PartInput
		return ret
	}
	return o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetPartsOk() ([]PartInput, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given []PartInput and assigns it to the Parts field.
func (o *LineItemEgresoInput) SetParts(v []PartInput) {
	o.Parts = v
}

// GetCustomsKeys returns the CustomsKeys field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetCustomsKeys() []string {
	if o == nil || IsNil(o.CustomsKeys) {
		var ret []string
		return ret
	}
	return o.CustomsKeys
}

// GetCustomsKeysOk returns a tuple with the CustomsKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetCustomsKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomsKeys) {
		return nil, false
	}
	return o.CustomsKeys, true
}

// HasCustomsKeys returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasCustomsKeys() bool {
	if o != nil && !IsNil(o.CustomsKeys) {
		return true
	}

	return false
}

// SetCustomsKeys gets a reference to the given []string and assigns it to the CustomsKeys field.
func (o *LineItemEgresoInput) SetCustomsKeys(v []string) {
	o.CustomsKeys = v
}

// GetComplement returns the Complement field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetComplement() string {
	if o == nil || IsNil(o.Complement) {
		var ret string
		return ret
	}
	return *o.Complement
}

// GetComplementOk returns a tuple with the Complement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetComplementOk() (*string, bool) {
	if o == nil || IsNil(o.Complement) {
		return nil, false
	}
	return o.Complement, true
}

// HasComplement returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasComplement() bool {
	if o != nil && !IsNil(o.Complement) {
		return true
	}

	return false
}

// SetComplement gets a reference to the given string and assigns it to the Complement field.
func (o *LineItemEgresoInput) SetComplement(v string) {
	o.Complement = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *LineItemEgresoInput) GetThirdParty() LineItemInputThirdParty {
	if o == nil || IsNil(o.ThirdParty) {
		var ret LineItemInputThirdParty
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemEgresoInput) GetThirdPartyOk() (*LineItemInputThirdParty, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *LineItemEgresoInput) HasThirdParty() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given LineItemInputThirdParty and assigns it to the ThirdParty field.
func (o *LineItemEgresoInput) SetThirdParty(v LineItemInputThirdParty) {
	o.ThirdParty = &v
}

func (o LineItemEgresoInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemEgresoInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	toSerialize["product"] = o.Product
	if !IsNil(o.Parts) {
		toSerialize["parts"] = o.Parts
	}
	if !IsNil(o.CustomsKeys) {
		toSerialize["customs_keys"] = o.CustomsKeys
	}
	if !IsNil(o.Complement) {
		toSerialize["complement"] = o.Complement
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["third_party"] = o.ThirdParty
	}
	return toSerialize, nil
}

type NullableLineItemEgresoInput struct {
	value *LineItemEgresoInput
	isSet bool
}

func (v NullableLineItemEgresoInput) Get() *LineItemEgresoInput {
	return v.value
}

func (v *NullableLineItemEgresoInput) Set(val *LineItemEgresoInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemEgresoInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemEgresoInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemEgresoInput(val *LineItemEgresoInput) *NullableLineItemEgresoInput {
	return &NullableLineItemEgresoInput{value: val, isSet: true}
}

func (v NullableLineItemEgresoInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemEgresoInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


