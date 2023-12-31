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

// checks if the ProductProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductProperties{}

// ProductProperties struct for ProductProperties
type ProductProperties struct {
	// Descripción del bien o servicio como aparecerá en la factura.
	Description *string `json:"description,omitempty"`
	// Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys).
	ProductKey *string `json:"product_key,omitempty"`
	// Precio por unidad del bien o servicio. Este valor representará el precio con IVA incluído o sin él, dependiendo del valor de `tax_included`.
	Price *float32 `json:"price,omitempty"`
	// - `true`: Indica que todos los impuestos aplicables están incluídos en el precio (atributo price) y se desglosarán automáticamente al emitir la factura. - `false`: Indica que el atributo price no incluye impuestos, por lo que aquellos impuestos a aplicar se sumarán en el precio final. 
	TaxIncluded *bool `json:"tax_included,omitempty"`
	// Código que representa si el bien o servicio es objeto de impuesto o no. Este atributo corresponde al campo \"ObjetoImp\" en el CFDI.  - `01`: No objeto de impuesto. - `02`: Sí objeto de impuesto. - `03`: Sí objeto de impuesto, pero no obligado a desglose. - `04`: Sí objeto de impuesto, y no causa impuesto. 
	Taxability *string `json:"taxability,omitempty"`
	Taxes []BaseTax `json:"taxes,omitempty"`
	LocalTaxes []LocalTax `json:"local_taxes,omitempty"`
	// Clave de unidad de medida, del catálogo del SAT. El valor por default `\"H87\"` (elemento) es la clave para representar una pieza o unidad de venta (lápiz, cuaderno, televisión, etc). Si la unidad de tu producto es kilogramos, litros, horas u otra unidad, puedes encontrar la clave utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys). 
	UnitKey *string `json:"unit_key,omitempty"`
	// Palabra que representa la unidad de medida de tu producto. Debe estar relacionada con la clave de unidad `unit_key`.
	UnitName *string `json:"unit_name,omitempty"`
	// Identificador de uso interno designado por la empresa. Puede tener cualquier valor.
	Sku *string `json:"sku,omitempty"`
}

// NewProductProperties instantiates a new ProductProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductProperties() *ProductProperties {
	this := ProductProperties{}
	var taxIncluded bool = true
	this.TaxIncluded = &taxIncluded
	var taxability string = "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "
	this.Taxability = &taxability
	var unitKey string = "H87"
	this.UnitKey = &unitKey
	var unitName string = "Elemento"
	this.UnitName = &unitName
	return &this
}

// NewProductPropertiesWithDefaults instantiates a new ProductProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPropertiesWithDefaults() *ProductProperties {
	this := ProductProperties{}
	var taxIncluded bool = true
	this.TaxIncluded = &taxIncluded
	var taxability string = "'01' si el array `taxes` está vacío; '02' si el array `taxes` tiene por lo menos un elemento. "
	this.Taxability = &taxability
	var unitKey string = "H87"
	this.UnitKey = &unitKey
	var unitName string = "Elemento"
	this.UnitName = &unitName
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductProperties) SetDescription(v string) {
	o.Description = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *ProductProperties) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *ProductProperties) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *ProductProperties) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductProperties) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductProperties) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ProductProperties) SetPrice(v float32) {
	o.Price = &v
}

// GetTaxIncluded returns the TaxIncluded field value if set, zero value otherwise.
func (o *ProductProperties) GetTaxIncluded() bool {
	if o == nil || IsNil(o.TaxIncluded) {
		var ret bool
		return ret
	}
	return *o.TaxIncluded
}

// GetTaxIncludedOk returns a tuple with the TaxIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetTaxIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.TaxIncluded) {
		return nil, false
	}
	return o.TaxIncluded, true
}

// HasTaxIncluded returns a boolean if a field has been set.
func (o *ProductProperties) HasTaxIncluded() bool {
	if o != nil && !IsNil(o.TaxIncluded) {
		return true
	}

	return false
}

// SetTaxIncluded gets a reference to the given bool and assigns it to the TaxIncluded field.
func (o *ProductProperties) SetTaxIncluded(v bool) {
	o.TaxIncluded = &v
}

// GetTaxability returns the Taxability field value if set, zero value otherwise.
func (o *ProductProperties) GetTaxability() string {
	if o == nil || IsNil(o.Taxability) {
		var ret string
		return ret
	}
	return *o.Taxability
}

// GetTaxabilityOk returns a tuple with the Taxability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetTaxabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Taxability) {
		return nil, false
	}
	return o.Taxability, true
}

// HasTaxability returns a boolean if a field has been set.
func (o *ProductProperties) HasTaxability() bool {
	if o != nil && !IsNil(o.Taxability) {
		return true
	}

	return false
}

// SetTaxability gets a reference to the given string and assigns it to the Taxability field.
func (o *ProductProperties) SetTaxability(v string) {
	o.Taxability = &v
}

// GetTaxes returns the Taxes field value if set, zero value otherwise.
func (o *ProductProperties) GetTaxes() []BaseTax {
	if o == nil || IsNil(o.Taxes) {
		var ret []BaseTax
		return ret
	}
	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetTaxesOk() ([]BaseTax, bool) {
	if o == nil || IsNil(o.Taxes) {
		return nil, false
	}
	return o.Taxes, true
}

// HasTaxes returns a boolean if a field has been set.
func (o *ProductProperties) HasTaxes() bool {
	if o != nil && !IsNil(o.Taxes) {
		return true
	}

	return false
}

// SetTaxes gets a reference to the given []BaseTax and assigns it to the Taxes field.
func (o *ProductProperties) SetTaxes(v []BaseTax) {
	o.Taxes = v
}

// GetLocalTaxes returns the LocalTaxes field value if set, zero value otherwise.
func (o *ProductProperties) GetLocalTaxes() []LocalTax {
	if o == nil || IsNil(o.LocalTaxes) {
		var ret []LocalTax
		return ret
	}
	return o.LocalTaxes
}

// GetLocalTaxesOk returns a tuple with the LocalTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetLocalTaxesOk() ([]LocalTax, bool) {
	if o == nil || IsNil(o.LocalTaxes) {
		return nil, false
	}
	return o.LocalTaxes, true
}

// HasLocalTaxes returns a boolean if a field has been set.
func (o *ProductProperties) HasLocalTaxes() bool {
	if o != nil && !IsNil(o.LocalTaxes) {
		return true
	}

	return false
}

// SetLocalTaxes gets a reference to the given []LocalTax and assigns it to the LocalTaxes field.
func (o *ProductProperties) SetLocalTaxes(v []LocalTax) {
	o.LocalTaxes = v
}

// GetUnitKey returns the UnitKey field value if set, zero value otherwise.
func (o *ProductProperties) GetUnitKey() string {
	if o == nil || IsNil(o.UnitKey) {
		var ret string
		return ret
	}
	return *o.UnitKey
}

// GetUnitKeyOk returns a tuple with the UnitKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetUnitKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UnitKey) {
		return nil, false
	}
	return o.UnitKey, true
}

// HasUnitKey returns a boolean if a field has been set.
func (o *ProductProperties) HasUnitKey() bool {
	if o != nil && !IsNil(o.UnitKey) {
		return true
	}

	return false
}

// SetUnitKey gets a reference to the given string and assigns it to the UnitKey field.
func (o *ProductProperties) SetUnitKey(v string) {
	o.UnitKey = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *ProductProperties) GetUnitName() string {
	if o == nil || IsNil(o.UnitName) {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetUnitNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitName) {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *ProductProperties) HasUnitName() bool {
	if o != nil && !IsNil(o.UnitName) {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *ProductProperties) SetUnitName(v string) {
	o.UnitName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductProperties) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductProperties) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductProperties) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductProperties) SetSku(v string) {
	o.Sku = &v
}

func (o ProductProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.TaxIncluded) {
		toSerialize["tax_included"] = o.TaxIncluded
	}
	if !IsNil(o.Taxability) {
		toSerialize["taxability"] = o.Taxability
	}
	if !IsNil(o.Taxes) {
		toSerialize["taxes"] = o.Taxes
	}
	if !IsNil(o.LocalTaxes) {
		toSerialize["local_taxes"] = o.LocalTaxes
	}
	if !IsNil(o.UnitKey) {
		toSerialize["unit_key"] = o.UnitKey
	}
	if !IsNil(o.UnitName) {
		toSerialize["unit_name"] = o.UnitName
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableProductProperties struct {
	value *ProductProperties
	isSet bool
}

func (v NullableProductProperties) Get() *ProductProperties {
	return v.value
}

func (v *NullableProductProperties) Set(val *ProductProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableProductProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableProductProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductProperties(val *ProductProperties) *NullableProductProperties {
	return &NullableProductProperties{value: val, isSet: true}
}

func (v NullableProductProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


