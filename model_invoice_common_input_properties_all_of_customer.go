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

// InvoiceCommonInputPropertiesAllOfCustomer - Cliente receptor de la factura.
type InvoiceCommonInputPropertiesAllOfCustomer struct {
	CustomerCreateInput *CustomerCreateInput
	String *string
}

// CustomerCreateInputAsInvoiceCommonInputPropertiesAllOfCustomer is a convenience function that returns CustomerCreateInput wrapped in InvoiceCommonInputPropertiesAllOfCustomer
func CustomerCreateInputAsInvoiceCommonInputPropertiesAllOfCustomer(v *CustomerCreateInput) InvoiceCommonInputPropertiesAllOfCustomer {
	return InvoiceCommonInputPropertiesAllOfCustomer{
		CustomerCreateInput: v,
	}
}

// stringAsInvoiceCommonInputPropertiesAllOfCustomer is a convenience function that returns string wrapped in InvoiceCommonInputPropertiesAllOfCustomer
func StringAsInvoiceCommonInputPropertiesAllOfCustomer(v *string) InvoiceCommonInputPropertiesAllOfCustomer {
	return InvoiceCommonInputPropertiesAllOfCustomer{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InvoiceCommonInputPropertiesAllOfCustomer) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CustomerCreateInput
	err = newStrictDecoder(data).Decode(&dst.CustomerCreateInput)
	if err == nil {
		jsonCustomerCreateInput, _ := json.Marshal(dst.CustomerCreateInput)
		if string(jsonCustomerCreateInput) == "{}" { // empty struct
			dst.CustomerCreateInput = nil
		} else {
			match++
		}
	} else {
		dst.CustomerCreateInput = nil
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
		dst.CustomerCreateInput = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InvoiceCommonInputPropertiesAllOfCustomer)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InvoiceCommonInputPropertiesAllOfCustomer)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InvoiceCommonInputPropertiesAllOfCustomer) MarshalJSON() ([]byte, error) {
	if src.CustomerCreateInput != nil {
		return json.Marshal(&src.CustomerCreateInput)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InvoiceCommonInputPropertiesAllOfCustomer) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CustomerCreateInput != nil {
		return obj.CustomerCreateInput
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableInvoiceCommonInputPropertiesAllOfCustomer struct {
	value *InvoiceCommonInputPropertiesAllOfCustomer
	isSet bool
}

func (v NullableInvoiceCommonInputPropertiesAllOfCustomer) Get() *InvoiceCommonInputPropertiesAllOfCustomer {
	return v.value
}

func (v *NullableInvoiceCommonInputPropertiesAllOfCustomer) Set(val *InvoiceCommonInputPropertiesAllOfCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCommonInputPropertiesAllOfCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCommonInputPropertiesAllOfCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCommonInputPropertiesAllOfCustomer(val *InvoiceCommonInputPropertiesAllOfCustomer) *NullableInvoiceCommonInputPropertiesAllOfCustomer {
	return &NullableInvoiceCommonInputPropertiesAllOfCustomer{value: val, isSet: true}
}

func (v NullableInvoiceCommonInputPropertiesAllOfCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCommonInputPropertiesAllOfCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


