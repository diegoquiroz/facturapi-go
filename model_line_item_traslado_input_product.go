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
	"fmt"
)

// LineItemTrasladoInputProduct - Objeto con información del producto o servicio facturado.
type LineItemTrasladoInputProduct struct {
	LineItemTrasladoProductInput *LineItemTrasladoProductInput
	String *string
}

// LineItemTrasladoProductInputAsLineItemTrasladoInputProduct is a convenience function that returns LineItemTrasladoProductInput wrapped in LineItemTrasladoInputProduct
func LineItemTrasladoProductInputAsLineItemTrasladoInputProduct(v *LineItemTrasladoProductInput) LineItemTrasladoInputProduct {
	return LineItemTrasladoInputProduct{
		LineItemTrasladoProductInput: v,
	}
}

// stringAsLineItemTrasladoInputProduct is a convenience function that returns string wrapped in LineItemTrasladoInputProduct
func StringAsLineItemTrasladoInputProduct(v *string) LineItemTrasladoInputProduct {
	return LineItemTrasladoInputProduct{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LineItemTrasladoInputProduct) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LineItemTrasladoProductInput
	err = newStrictDecoder(data).Decode(&dst.LineItemTrasladoProductInput)
	if err == nil {
		jsonLineItemTrasladoProductInput, _ := json.Marshal(dst.LineItemTrasladoProductInput)
		if string(jsonLineItemTrasladoProductInput) == "{}" { // empty struct
			dst.LineItemTrasladoProductInput = nil
		} else {
			match++
		}
	} else {
		dst.LineItemTrasladoProductInput = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LineItemTrasladoProductInput = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LineItemTrasladoInputProduct)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LineItemTrasladoInputProduct)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LineItemTrasladoInputProduct) MarshalJSON() ([]byte, error) {
	if src.LineItemTrasladoProductInput != nil {
		return json.Marshal(&src.LineItemTrasladoProductInput)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LineItemTrasladoInputProduct) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LineItemTrasladoProductInput != nil {
		return obj.LineItemTrasladoProductInput
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableLineItemTrasladoInputProduct struct {
	value *LineItemTrasladoInputProduct
	isSet bool
}

func (v NullableLineItemTrasladoInputProduct) Get() *LineItemTrasladoInputProduct {
	return v.value
}

func (v *NullableLineItemTrasladoInputProduct) Set(val *LineItemTrasladoInputProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemTrasladoInputProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemTrasladoInputProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemTrasladoInputProduct(val *LineItemTrasladoInputProduct) *NullableLineItemTrasladoInputProduct {
	return &NullableLineItemTrasladoInputProduct{value: val, isSet: true}
}

func (v NullableLineItemTrasladoInputProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemTrasladoInputProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


