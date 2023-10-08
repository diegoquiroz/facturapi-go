/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// RetentionAPIService RetentionAPI service
type RetentionAPIService service

type ApiCancelRetentionRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	retentionId string
}

func (r ApiCancelRetentionRequest) Execute() (*Retention, *http.Response, error) {
	return r.ApiService.CancelRetentionExecute(r)
}

/*
CancelRetention Cancelar retención

Realiza una solicitud de cancelación de retención ante el SAT.

A diferencia de las facturas comúnes, la cancelación de la retención es inmediata y no requiere autorización de parte del receptor.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param retentionId ID de la retención a cancelar
 @return ApiCancelRetentionRequest
*/
func (a *RetentionAPIService) CancelRetention(ctx context.Context, retentionId string) ApiCancelRetentionRequest {
	return ApiCancelRetentionRequest{
		ApiService: a,
		ctx: ctx,
		retentionId: retentionId,
	}
}

// Execute executes the request
//  @return Retention
func (a *RetentionAPIService) CancelRetentionExecute(r ApiCancelRetentionRequest) (*Retention, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Retention
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.CancelRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions/{retention_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"retention_id"+"}", url.PathEscape(parameterValueToString(r.retentionId, "retentionId")), -1)

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

type ApiCreateRetentionRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	retentionInput *RetentionInput
}

func (r ApiCreateRetentionRequest) RetentionInput(retentionInput RetentionInput) ApiCreateRetentionRequest {
	r.retentionInput = &retentionInput
	return r
}

func (r ApiCreateRetentionRequest) Execute() (*Retention, *http.Response, error) {
	return r.ApiService.CreateRetentionExecute(r)
}

/*
CreateRetention Crear retención

Crea una nueva Retención. Si el comprobante es creado en ambiente Live, ésta será **timbrado y enviado al SAT**.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRetentionRequest
*/
func (a *RetentionAPIService) CreateRetention(ctx context.Context) ApiCreateRetentionRequest {
	return ApiCreateRetentionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Retention
func (a *RetentionAPIService) CreateRetentionExecute(r ApiCreateRetentionRequest) (*Retention, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Retention
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.CreateRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.retentionInput
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

type ApiDownloadRetentionRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	retentionId string
	format string
}

func (r ApiDownloadRetentionRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.DownloadRetentionExecute(r)
}

/*
DownloadRetention Descargar retención

Descarga una retención en PDF, XML o ambos en un archivo comprimido ZIP.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param retentionId ID del objeto a descargar
 @param format Formato del archivo de descarga
 @return ApiDownloadRetentionRequest
*/
func (a *RetentionAPIService) DownloadRetention(ctx context.Context, retentionId string, format string) ApiDownloadRetentionRequest {
	return ApiDownloadRetentionRequest{
		ApiService: a,
		ctx: ctx,
		retentionId: retentionId,
		format: format,
	}
}

// Execute executes the request
//  @return *os.File
func (a *RetentionAPIService) DownloadRetentionExecute(r ApiDownloadRetentionRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.DownloadRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions/{retention_id}/{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"retention_id"+"}", url.PathEscape(parameterValueToString(r.retentionId, "retentionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

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

type ApiGetRetentionRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	retentionId string
}

func (r ApiGetRetentionRequest) Execute() (*Retention, *http.Response, error) {
	return r.ApiService.GetRetentionExecute(r)
}

/*
GetRetention Obtener retención por ID

Regresa el objeto 'Retention' relacionado al `id` especificado.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param retentionId ID del objeto a obtener
 @return ApiGetRetentionRequest
*/
func (a *RetentionAPIService) GetRetention(ctx context.Context, retentionId string) ApiGetRetentionRequest {
	return ApiGetRetentionRequest{
		ApiService: a,
		ctx: ctx,
		retentionId: retentionId,
	}
}

// Execute executes the request
//  @return Retention
func (a *RetentionAPIService) GetRetentionExecute(r ApiGetRetentionRequest) (*Retention, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Retention
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.GetRetention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions/{retention_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"retention_id"+"}", url.PathEscape(parameterValueToString(r.retentionId, "retentionId")), -1)

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

type ApiListRetentionsRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	q *string
	customer *string
	date *DateRange
	page *interface{}
	limit *interface{}
}

// Consulta. Texto a buscar en el nombre fiscal del cliente o su RFC.
func (r ApiListRetentionsRequest) Q(q string) ApiListRetentionsRequest {
	r.q = &q
	return r
}

// Identificador del cliente. Útil para obtener las retenciones emitidas a un sólo cliente.
func (r ApiListRetentionsRequest) Customer(customer string) ApiListRetentionsRequest {
	r.customer = &customer
	return r
}

// Objeto con rango de fechas solicitado.
func (r ApiListRetentionsRequest) Date(date DateRange) ApiListRetentionsRequest {
	r.date = &date
	return r
}

// Página de resultados a regresar, empezando desde la página 1.
func (r ApiListRetentionsRequest) Page(page interface{}) ApiListRetentionsRequest {
	r.page = &page
	return r
}

// Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación.
func (r ApiListRetentionsRequest) Limit(limit interface{}) ApiListRetentionsRequest {
	r.limit = &limit
	return r
}

func (r ApiListRetentionsRequest) Execute() (*RetentionSearchResult, *http.Response, error) {
	return r.ApiService.ListRetentionsExecute(r)
}

/*
ListRetentions Listar retenciones

Regresa una lista paginada de todas las retenciones de una organización o realiza una búsqueda de acuerdo a parámetros

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRetentionsRequest
*/
func (a *RetentionAPIService) ListRetentions(ctx context.Context) ApiListRetentionsRequest {
	return ApiListRetentionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RetentionSearchResult
func (a *RetentionAPIService) ListRetentionsExecute(r ApiListRetentionsRequest) (*RetentionSearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RetentionSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.ListRetentions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.customer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customer", r.customer, "")
	}
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue interface{} = 50
		r.limit = &defaultValue
	}
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

type ApiSendRetentionByEmailRequest struct {
	ctx context.Context
	ApiService *RetentionAPIService
	retentionId string
	sendInvoiceByEmailRequest *SendInvoiceByEmailRequest
}

func (r ApiSendRetentionByEmailRequest) SendInvoiceByEmailRequest(sendInvoiceByEmailRequest SendInvoiceByEmailRequest) ApiSendRetentionByEmailRequest {
	r.sendInvoiceByEmailRequest = &sendInvoiceByEmailRequest
	return r
}

func (r ApiSendRetentionByEmailRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.SendRetentionByEmailExecute(r)
}

/*
SendRetentionByEmail Enviar retención por correo electrónico

Envía un correo electrónico a la dirección de tu cliente, con los archivos XML y PDF adjuntos al mensaje.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param retentionId ID del objeto a obtener
 @return ApiSendRetentionByEmailRequest
*/
func (a *RetentionAPIService) SendRetentionByEmail(ctx context.Context, retentionId string) ApiSendRetentionByEmailRequest {
	return ApiSendRetentionByEmailRequest{
		ApiService: a,
		ctx: ctx,
		retentionId: retentionId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *RetentionAPIService) SendRetentionByEmailExecute(r ApiSendRetentionByEmailRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RetentionAPIService.SendRetentionByEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/retentions/{retention_id}/email"
	localVarPath = strings.Replace(localVarPath, "{"+"retention_id"+"}", url.PathEscape(parameterValueToString(r.retentionId, "retentionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sendInvoiceByEmailRequest
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
