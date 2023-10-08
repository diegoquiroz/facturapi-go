/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi-go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ToolsAPIService ToolsAPI service
type ToolsAPIService service

type ApiCheckGetRequest struct {
	ctx context.Context
	ApiService *ToolsAPIService
}

func (r ApiCheckGetRequest) Execute() (*CheckGet200Response, *http.Response, error) {
	return r.ApiService.CheckGetExecute(r)
}

/*
CheckGet Health check (Pulso)

Indica el estatus de disponibilidad de la API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCheckGetRequest
*/
func (a *ToolsAPIService) CheckGet(ctx context.Context) ApiCheckGetRequest {
	return ApiCheckGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckGet200Response
func (a *ToolsAPIService) CheckGetExecute(r ApiCheckGetRequest) (*CheckGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolsAPIService.CheckGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToolsTaxIdValidationGetRequest struct {
	ctx context.Context
	ApiService *ToolsAPIService
	taxId *string
}

func (r ApiToolsTaxIdValidationGetRequest) TaxId(taxId string) ApiToolsTaxIdValidationGetRequest {
	r.taxId = &taxId
	return r
}

func (r ApiToolsTaxIdValidationGetRequest) Execute() (*TaxIdValidationResult, *http.Response, error) {
	return r.ApiService.ToolsTaxIdValidationGetExecute(r)
}

/*
ToolsTaxIdValidationGet Validar RFC

Consulta el estado de un RFC en la lista de **EFOS** (Empresas que
Facturan Operaciones Simuladas). Al aparecer en esta lista, el RFC es o
fue sospechoso de incurrir en simulación de operaciones fiscales
(empresas factureras).

La respuesta (detallada más abajo) incluye los resultados de esta
validación. Se incluye la propiedad
booleana `is_valid`, que Facturapi resuelve interpretando la respuesta.
Un valor de `true` para esta propiedad indica que el RFC no tiene asuntos
por resolver y está libre de problemas; y lo contrario para `false`.
Adicionalmente puedes consultar la propiedad data para ver los valores
en bruto de la consulta al SAT.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolsTaxIdValidationGetRequest
*/
func (a *ToolsAPIService) ToolsTaxIdValidationGet(ctx context.Context) ApiToolsTaxIdValidationGetRequest {
	return ApiToolsTaxIdValidationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TaxIdValidationResult
func (a *ToolsAPIService) ToolsTaxIdValidationGetExecute(r ApiToolsTaxIdValidationGetRequest) (*TaxIdValidationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaxIdValidationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolsAPIService.ToolsTaxIdValidationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tools/tax_id_validation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxId == nil {
		return localVarReturnValue, nil, reportError("taxId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tax_id", r.taxId, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
