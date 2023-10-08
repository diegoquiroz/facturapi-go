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
	"fmt"
)

// LineItemEgresoInputProduct - Objeto con información del producto o servicio facturado.
type LineItemEgresoInputProduct struct {
	LineItemProductEgresoInput *LineItemProductEgresoInput
	String *string
}

// LineItemProductEgresoInputAsLineItemEgresoInputProduct is a convenience function that returns LineItemProductEgresoInput wrapped in LineItemEgresoInputProduct
func LineItemProductEgresoInputAsLineItemEgresoInputProduct(v *LineItemProductEgresoInput) LineItemEgresoInputProduct {
	return LineItemEgresoInputProduct{
		LineItemProductEgresoInput: v,
	}
}

// stringAsLineItemEgresoInputProduct is a convenience function that returns string wrapped in LineItemEgresoInputProduct
func StringAsLineItemEgresoInputProduct(v *string) LineItemEgresoInputProduct {
	return LineItemEgresoInputProduct{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LineItemEgresoInputProduct) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LineItemProductEgresoInput
	err = newStrictDecoder(data).Decode(&dst.LineItemProductEgresoInput)
	if err == nil {
		jsonLineItemProductEgresoInput, _ := json.Marshal(dst.LineItemProductEgresoInput)
		if string(jsonLineItemProductEgresoInput) == "{}" { // empty struct
			dst.LineItemProductEgresoInput = nil
		} else {
			match++
		}
	} else {
		dst.LineItemProductEgresoInput = nil
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
		dst.LineItemProductEgresoInput = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LineItemEgresoInputProduct)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LineItemEgresoInputProduct)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LineItemEgresoInputProduct) MarshalJSON() ([]byte, error) {
	if src.LineItemProductEgresoInput != nil {
		return json.Marshal(&src.LineItemProductEgresoInput)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LineItemEgresoInputProduct) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LineItemProductEgresoInput != nil {
		return obj.LineItemProductEgresoInput
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableLineItemEgresoInputProduct struct {
	value *LineItemEgresoInputProduct
	isSet bool
}

func (v NullableLineItemEgresoInputProduct) Get() *LineItemEgresoInputProduct {
	return v.value
}

func (v *NullableLineItemEgresoInputProduct) Set(val *LineItemEgresoInputProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemEgresoInputProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemEgresoInputProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemEgresoInputProduct(val *LineItemEgresoInputProduct) *NullableLineItemEgresoInputProduct {
	return &NullableLineItemEgresoInputProduct{value: val, isSet: true}
}

func (v NullableLineItemEgresoInputProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemEgresoInputProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


