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

// checks if the OrganizationCustomizationPdfExtra type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationCustomizationPdfExtra{}

// OrganizationCustomizationPdfExtra Configura qué campos opcionales se queiren mostrar en el PDF. El SAT no obliga a mostrar estos campos, pero pueden activarse según la preferencia de la organización.
type OrganizationCustomizationPdfExtra struct {
	// Mostrar códigos de catálogos del SAT junto a sus descripciones. Ejemplo: “KGM Kilogramo”. 
	Codes *bool `json:"codes,omitempty"`
	// Mostrar la clave de producto-servicio. 
	ProductKey *bool `json:"product_key,omitempty"`
	// Redondear el precio unitario en el PDF a 2 decimales, pero conservar los 6 decimales en el XML. 
	RoundUnitPrice *bool `json:"round_unit_price,omitempty"`
	// Mostrar el desglose de impuestos en el PDF. Si se desactiva, sólo se mostratán los impuestos en los totales, pero no en el detalle de cada concepto. 
	TaxBreakdown *bool `json:"tax_breakdown,omitempty"`
}

// NewOrganizationCustomizationPdfExtra instantiates a new OrganizationCustomizationPdfExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationCustomizationPdfExtra() *OrganizationCustomizationPdfExtra {
	this := OrganizationCustomizationPdfExtra{}
	var codes bool = true
	this.Codes = &codes
	var productKey bool = true
	this.ProductKey = &productKey
	var roundUnitPrice bool = false
	this.RoundUnitPrice = &roundUnitPrice
	var taxBreakdown bool = true
	this.TaxBreakdown = &taxBreakdown
	return &this
}

// NewOrganizationCustomizationPdfExtraWithDefaults instantiates a new OrganizationCustomizationPdfExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationCustomizationPdfExtraWithDefaults() *OrganizationCustomizationPdfExtra {
	this := OrganizationCustomizationPdfExtra{}
	var codes bool = true
	this.Codes = &codes
	var productKey bool = true
	this.ProductKey = &productKey
	var roundUnitPrice bool = false
	this.RoundUnitPrice = &roundUnitPrice
	var taxBreakdown bool = true
	this.TaxBreakdown = &taxBreakdown
	return &this
}

// GetCodes returns the Codes field value if set, zero value otherwise.
func (o *OrganizationCustomizationPdfExtra) GetCodes() bool {
	if o == nil || IsNil(o.Codes) {
		var ret bool
		return ret
	}
	return *o.Codes
}

// GetCodesOk returns a tuple with the Codes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomizationPdfExtra) GetCodesOk() (*bool, bool) {
	if o == nil || IsNil(o.Codes) {
		return nil, false
	}
	return o.Codes, true
}

// HasCodes returns a boolean if a field has been set.
func (o *OrganizationCustomizationPdfExtra) HasCodes() bool {
	if o != nil && !IsNil(o.Codes) {
		return true
	}

	return false
}

// SetCodes gets a reference to the given bool and assigns it to the Codes field.
func (o *OrganizationCustomizationPdfExtra) SetCodes(v bool) {
	o.Codes = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *OrganizationCustomizationPdfExtra) GetProductKey() bool {
	if o == nil || IsNil(o.ProductKey) {
		var ret bool
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomizationPdfExtra) GetProductKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *OrganizationCustomizationPdfExtra) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given bool and assigns it to the ProductKey field.
func (o *OrganizationCustomizationPdfExtra) SetProductKey(v bool) {
	o.ProductKey = &v
}

// GetRoundUnitPrice returns the RoundUnitPrice field value if set, zero value otherwise.
func (o *OrganizationCustomizationPdfExtra) GetRoundUnitPrice() bool {
	if o == nil || IsNil(o.RoundUnitPrice) {
		var ret bool
		return ret
	}
	return *o.RoundUnitPrice
}

// GetRoundUnitPriceOk returns a tuple with the RoundUnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomizationPdfExtra) GetRoundUnitPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.RoundUnitPrice) {
		return nil, false
	}
	return o.RoundUnitPrice, true
}

// HasRoundUnitPrice returns a boolean if a field has been set.
func (o *OrganizationCustomizationPdfExtra) HasRoundUnitPrice() bool {
	if o != nil && !IsNil(o.RoundUnitPrice) {
		return true
	}

	return false
}

// SetRoundUnitPrice gets a reference to the given bool and assigns it to the RoundUnitPrice field.
func (o *OrganizationCustomizationPdfExtra) SetRoundUnitPrice(v bool) {
	o.RoundUnitPrice = &v
}

// GetTaxBreakdown returns the TaxBreakdown field value if set, zero value otherwise.
func (o *OrganizationCustomizationPdfExtra) GetTaxBreakdown() bool {
	if o == nil || IsNil(o.TaxBreakdown) {
		var ret bool
		return ret
	}
	return *o.TaxBreakdown
}

// GetTaxBreakdownOk returns a tuple with the TaxBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationCustomizationPdfExtra) GetTaxBreakdownOk() (*bool, bool) {
	if o == nil || IsNil(o.TaxBreakdown) {
		return nil, false
	}
	return o.TaxBreakdown, true
}

// HasTaxBreakdown returns a boolean if a field has been set.
func (o *OrganizationCustomizationPdfExtra) HasTaxBreakdown() bool {
	if o != nil && !IsNil(o.TaxBreakdown) {
		return true
	}

	return false
}

// SetTaxBreakdown gets a reference to the given bool and assigns it to the TaxBreakdown field.
func (o *OrganizationCustomizationPdfExtra) SetTaxBreakdown(v bool) {
	o.TaxBreakdown = &v
}

func (o OrganizationCustomizationPdfExtra) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationCustomizationPdfExtra) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Codes) {
		toSerialize["codes"] = o.Codes
	}
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
	}
	if !IsNil(o.RoundUnitPrice) {
		toSerialize["round_unit_price"] = o.RoundUnitPrice
	}
	if !IsNil(o.TaxBreakdown) {
		toSerialize["tax_breakdown"] = o.TaxBreakdown
	}
	return toSerialize, nil
}

type NullableOrganizationCustomizationPdfExtra struct {
	value *OrganizationCustomizationPdfExtra
	isSet bool
}

func (v NullableOrganizationCustomizationPdfExtra) Get() *OrganizationCustomizationPdfExtra {
	return v.value
}

func (v *NullableOrganizationCustomizationPdfExtra) Set(val *OrganizationCustomizationPdfExtra) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationCustomizationPdfExtra) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationCustomizationPdfExtra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationCustomizationPdfExtra(val *OrganizationCustomizationPdfExtra) *NullableOrganizationCustomizationPdfExtra {
	return &NullableOrganizationCustomizationPdfExtra{value: val, isSet: true}
}

func (v NullableOrganizationCustomizationPdfExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationCustomizationPdfExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


