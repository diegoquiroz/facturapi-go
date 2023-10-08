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

// checks if the LineItemTrasladoProductInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemTrasladoProductInput{}

// LineItemTrasladoProductInput struct for LineItemTrasladoProductInput
type LineItemTrasladoProductInput struct {
	// Descripción del bien o servicio como aparecerá en la factura.
	Description string `json:"description"`
	// Clave de producto/servicio, del catálogo del SAT. Nosotros te proporcionamos una manera más conveniente de encontrarlo utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys).
	ProductKey *string `json:"product_key,omitempty"`
	// Clave de unidad de medida, del catálogo del SAT. El valor por default `\"H87\"` (elemento) es la clave para representar una pieza o unidad de venta (lápiz, cuaderno, televisión, etc). Si la unidad de tu producto es kilogramos, litros, horas u otra unidad, te proporcionamos una manera conveniente de encontrar la clave utilizando nuestra [herramienta de búsqueda de claves](https://dashboard.facturapi.io/tools/keys). 
	UnitKey *string `json:"unit_key,omitempty"`
	// Palabra que representa la unidad de medida de tu producto. Debe estar relacionada con la clave de unidad `unit_key`.
	UnitName *string `json:"unit_name,omitempty"`
	// Identificador de uso interno designado por la empresa. Puede tener cualquier valor.
	Sku *string `json:"sku,omitempty"`
}

// NewLineItemTrasladoProductInput instantiates a new LineItemTrasladoProductInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemTrasladoProductInput(description string) *LineItemTrasladoProductInput {
	this := LineItemTrasladoProductInput{}
	this.Description = description
	var unitKey string = "H87"
	this.UnitKey = &unitKey
	var unitName string = "Elemento"
	this.UnitName = &unitName
	return &this
}

// NewLineItemTrasladoProductInputWithDefaults instantiates a new LineItemTrasladoProductInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemTrasladoProductInputWithDefaults() *LineItemTrasladoProductInput {
	this := LineItemTrasladoProductInput{}
	var unitKey string = "H87"
	this.UnitKey = &unitKey
	var unitName string = "Elemento"
	this.UnitName = &unitName
	return &this
}

// GetDescription returns the Description field value
func (o *LineItemTrasladoProductInput) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LineItemTrasladoProductInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LineItemTrasladoProductInput) SetDescription(v string) {
	o.Description = v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *LineItemTrasladoProductInput) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemTrasladoProductInput) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *LineItemTrasladoProductInput) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *LineItemTrasladoProductInput) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetUnitKey returns the UnitKey field value if set, zero value otherwise.
func (o *LineItemTrasladoProductInput) GetUnitKey() string {
	if o == nil || IsNil(o.UnitKey) {
		var ret string
		return ret
	}
	return *o.UnitKey
}

// GetUnitKeyOk returns a tuple with the UnitKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemTrasladoProductInput) GetUnitKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UnitKey) {
		return nil, false
	}
	return o.UnitKey, true
}

// HasUnitKey returns a boolean if a field has been set.
func (o *LineItemTrasladoProductInput) HasUnitKey() bool {
	if o != nil && !IsNil(o.UnitKey) {
		return true
	}

	return false
}

// SetUnitKey gets a reference to the given string and assigns it to the UnitKey field.
func (o *LineItemTrasladoProductInput) SetUnitKey(v string) {
	o.UnitKey = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *LineItemTrasladoProductInput) GetUnitName() string {
	if o == nil || IsNil(o.UnitName) {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemTrasladoProductInput) GetUnitNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitName) {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *LineItemTrasladoProductInput) HasUnitName() bool {
	if o != nil && !IsNil(o.UnitName) {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *LineItemTrasladoProductInput) SetUnitName(v string) {
	o.UnitName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *LineItemTrasladoProductInput) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemTrasladoProductInput) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *LineItemTrasladoProductInput) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *LineItemTrasladoProductInput) SetSku(v string) {
	o.Sku = &v
}

func (o LineItemTrasladoProductInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemTrasladoProductInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
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

type NullableLineItemTrasladoProductInput struct {
	value *LineItemTrasladoProductInput
	isSet bool
}

func (v NullableLineItemTrasladoProductInput) Get() *LineItemTrasladoProductInput {
	return v.value
}

func (v *NullableLineItemTrasladoProductInput) Set(val *LineItemTrasladoProductInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemTrasladoProductInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemTrasladoProductInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemTrasladoProductInput(val *LineItemTrasladoProductInput) *NullableLineItemTrasladoProductInput {
	return &NullableLineItemTrasladoProductInput{value: val, isSet: true}
}

func (v NullableLineItemTrasladoProductInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemTrasladoProductInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


