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
	"fmt"
)

// CreateInvoiceRequest - struct for CreateInvoiceRequest
type CreateInvoiceRequest struct {
	InvoiceEgresoInput *InvoiceEgresoInput
	InvoiceIngresoInput *InvoiceIngresoInput
	InvoiceNominaInput *InvoiceNominaInput
	InvoicePagoInput *InvoicePagoInput
	InvoiceTrasladoInput *InvoiceTrasladoInput
}

// InvoiceEgresoInputAsCreateInvoiceRequest is a convenience function that returns InvoiceEgresoInput wrapped in CreateInvoiceRequest
func InvoiceEgresoInputAsCreateInvoiceRequest(v *InvoiceEgresoInput) CreateInvoiceRequest {
	return CreateInvoiceRequest{
		InvoiceEgresoInput: v,
	}
}

// InvoiceIngresoInputAsCreateInvoiceRequest is a convenience function that returns InvoiceIngresoInput wrapped in CreateInvoiceRequest
func InvoiceIngresoInputAsCreateInvoiceRequest(v *InvoiceIngresoInput) CreateInvoiceRequest {
	return CreateInvoiceRequest{
		InvoiceIngresoInput: v,
	}
}

// InvoiceNominaInputAsCreateInvoiceRequest is a convenience function that returns InvoiceNominaInput wrapped in CreateInvoiceRequest
func InvoiceNominaInputAsCreateInvoiceRequest(v *InvoiceNominaInput) CreateInvoiceRequest {
	return CreateInvoiceRequest{
		InvoiceNominaInput: v,
	}
}

// InvoicePagoInputAsCreateInvoiceRequest is a convenience function that returns InvoicePagoInput wrapped in CreateInvoiceRequest
func InvoicePagoInputAsCreateInvoiceRequest(v *InvoicePagoInput) CreateInvoiceRequest {
	return CreateInvoiceRequest{
		InvoicePagoInput: v,
	}
}

// InvoiceTrasladoInputAsCreateInvoiceRequest is a convenience function that returns InvoiceTrasladoInput wrapped in CreateInvoiceRequest
func InvoiceTrasladoInputAsCreateInvoiceRequest(v *InvoiceTrasladoInput) CreateInvoiceRequest {
	return CreateInvoiceRequest{
		InvoiceTrasladoInput: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateInvoiceRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvoiceEgresoInput
	err = newStrictDecoder(data).Decode(&dst.InvoiceEgresoInput)
	if err == nil {
		jsonInvoiceEgresoInput, _ := json.Marshal(dst.InvoiceEgresoInput)
		if string(jsonInvoiceEgresoInput) == "{}" { // empty struct
			dst.InvoiceEgresoInput = nil
		} else {
			match++
		}
	} else {
		dst.InvoiceEgresoInput = nil
	}

	// try to unmarshal data into InvoiceIngresoInput
	err = newStrictDecoder(data).Decode(&dst.InvoiceIngresoInput)
	if err == nil {
		jsonInvoiceIngresoInput, _ := json.Marshal(dst.InvoiceIngresoInput)
		if string(jsonInvoiceIngresoInput) == "{}" { // empty struct
			dst.InvoiceIngresoInput = nil
		} else {
			match++
		}
	} else {
		dst.InvoiceIngresoInput = nil
	}

	// try to unmarshal data into InvoiceNominaInput
	err = newStrictDecoder(data).Decode(&dst.InvoiceNominaInput)
	if err == nil {
		jsonInvoiceNominaInput, _ := json.Marshal(dst.InvoiceNominaInput)
		if string(jsonInvoiceNominaInput) == "{}" { // empty struct
			dst.InvoiceNominaInput = nil
		} else {
			match++
		}
	} else {
		dst.InvoiceNominaInput = nil
	}

	// try to unmarshal data into InvoicePagoInput
	err = newStrictDecoder(data).Decode(&dst.InvoicePagoInput)
	if err == nil {
		jsonInvoicePagoInput, _ := json.Marshal(dst.InvoicePagoInput)
		if string(jsonInvoicePagoInput) == "{}" { // empty struct
			dst.InvoicePagoInput = nil
		} else {
			match++
		}
	} else {
		dst.InvoicePagoInput = nil
	}

	// try to unmarshal data into InvoiceTrasladoInput
	err = newStrictDecoder(data).Decode(&dst.InvoiceTrasladoInput)
	if err == nil {
		jsonInvoiceTrasladoInput, _ := json.Marshal(dst.InvoiceTrasladoInput)
		if string(jsonInvoiceTrasladoInput) == "{}" { // empty struct
			dst.InvoiceTrasladoInput = nil
		} else {
			match++
		}
	} else {
		dst.InvoiceTrasladoInput = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvoiceEgresoInput = nil
		dst.InvoiceIngresoInput = nil
		dst.InvoiceNominaInput = nil
		dst.InvoicePagoInput = nil
		dst.InvoiceTrasladoInput = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateInvoiceRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateInvoiceRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	if src.InvoiceEgresoInput != nil {
		return json.Marshal(&src.InvoiceEgresoInput)
	}

	if src.InvoiceIngresoInput != nil {
		return json.Marshal(&src.InvoiceIngresoInput)
	}

	if src.InvoiceNominaInput != nil {
		return json.Marshal(&src.InvoiceNominaInput)
	}

	if src.InvoicePagoInput != nil {
		return json.Marshal(&src.InvoicePagoInput)
	}

	if src.InvoiceTrasladoInput != nil {
		return json.Marshal(&src.InvoiceTrasladoInput)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateInvoiceRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InvoiceEgresoInput != nil {
		return obj.InvoiceEgresoInput
	}

	if obj.InvoiceIngresoInput != nil {
		return obj.InvoiceIngresoInput
	}

	if obj.InvoiceNominaInput != nil {
		return obj.InvoiceNominaInput
	}

	if obj.InvoicePagoInput != nil {
		return obj.InvoicePagoInput
	}

	if obj.InvoiceTrasladoInput != nil {
		return obj.InvoiceTrasladoInput
	}

	// all schemas are nil
	return nil
}

type NullableCreateInvoiceRequest struct {
	value *CreateInvoiceRequest
	isSet bool
}

func (v NullableCreateInvoiceRequest) Get() *CreateInvoiceRequest {
	return v.value
}

func (v *NullableCreateInvoiceRequest) Set(val *CreateInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceRequest(val *CreateInvoiceRequest) *NullableCreateInvoiceRequest {
	return &NullableCreateInvoiceRequest{value: val, isSet: true}
}

func (v NullableCreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


