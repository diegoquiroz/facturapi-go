/*
Facturapi

<br /> <br />  En esta página enlistamos todos los métodos disponibles en la API de Facturapi, así como la referencia completa de los parámetros que acepta cada uno. Para ver las propiedades anidadas de un objeto o arreglo de objetos, puedes hacer clic sobre el nombre del campo y expandirlo.  La API de Facturapi está diseñada con el estándar [REST](https://developer.mozilla.org/en-US/docs/Glossary/REST) en mente. Los endpoints de la API están agrupados por recursos, tienen URLs predecibles, las respuestas tienen formato JSON y usamos códigos HTTP de respuesta, autenticación y verbos estándar.  Durante el desarrollo, puedes usar la API de Facturapi en ambiente Test y las facturas que emitas no se enviarán al SAT ni tendrán validez fiscal.  La llave secreta que utilices para autenticarte determinará tanto el ambiente en el que se creará la factura (Test o Live), así como la organización a utilizar como emisor de tu factura, o bien como dueña del recurso que solicites crear. 

API version: 2.0
Contact: soporte@facturapi.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package facturapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CustomerAPIService CustomerAPI service
type CustomerAPIService service

type ApiCreateCustomerRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	customerCreateInput *CustomerCreateInput
}

func (r ApiCreateCustomerRequest) CustomerCreateInput(customerCreateInput CustomerCreateInput) ApiCreateCustomerRequest {
	r.customerCreateInput = &customerCreateInput
	return r
}

func (r ApiCreateCustomerRequest) Execute() (*Customer, *http.Response, error) {
	return r.ApiService.CreateCustomerExecute(r)
}

/*
CreateCustomer Crear cliente

Registra un nuevo cliente en Facturapi.

Esta llamada valida que los datos fiscales coincidan con
los registros del SAT para ese RFC, de lo contrario, la llamada
devolverá un error indicando el problema.

Una vez creado el cliente y obtenido un objeto de respuesta,
te recomendamos guardar el ID en tu base de datos junto a la información
de tu cliente. Posteriormente, puedes llamar al encpoint de Crear Factura
pasando el ID del cliente en lugar de repetir la información.

Por último, ten en cuenta que los clientes que crees en ambiente _Test_ **no se
comparten** con el ambiente _Live_.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCustomerRequest
*/
func (a *CustomerAPIService) CreateCustomer(ctx context.Context) ApiCreateCustomerRequest {
	return ApiCreateCustomerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerAPIService) CreateCustomerExecute(r ApiCreateCustomerRequest) (*Customer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.CreateCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers"

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
	localVarPostBody = r.customerCreateInput
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

type ApiDeleteCustomerRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	customerId string
}

func (r ApiDeleteCustomerRequest) Execute() (*Customer, *http.Response, error) {
	return r.ApiService.DeleteCustomerExecute(r)
}

/*
DeleteCustomer Eliminar cliente

Elimina el cliente de tu organización. Las facturas asociadas al cliente **no** se eliminarán.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId ID del objeto a eliminar
 @return ApiDeleteCustomerRequest
*/
func (a *CustomerAPIService) DeleteCustomer(ctx context.Context, customerId string) ApiDeleteCustomerRequest {
	return ApiDeleteCustomerRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerAPIService) DeleteCustomerExecute(r ApiDeleteCustomerRequest) (*Customer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.DeleteCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)

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

type ApiEditCustomerRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	customerId string
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiEditCustomerRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiEditCustomerRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiEditCustomerRequest) Execute() (*Customer, *http.Response, error) {
	return r.ApiService.EditCustomerExecute(r)
}

/*
EditCustomer Editar cliente

Actualiza la información de un cliente existente, asignando los valores de los parámetros enviados. Los parámetros que no se envíen en la petición no se modificarán.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId ID del objeto a editar
 @return ApiEditCustomerRequest
*/
func (a *CustomerAPIService) EditCustomer(ctx context.Context, customerId string) ApiEditCustomerRequest {
	return ApiEditCustomerRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerAPIService) EditCustomerExecute(r ApiEditCustomerRequest) (*Customer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.EditCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)

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
	localVarPostBody = r.uNKNOWNBASETYPE
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

type ApiGetCustomerRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	customerId string
}

func (r ApiGetCustomerRequest) Execute() (*Customer, *http.Response, error) {
	return r.ApiService.GetCustomerExecute(r)
}

/*
GetCustomer Obtener cliente por ID

Regresa el objeto 'Customer' relacionado al `id` especificado.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId ID del objeto a obtener
 @return ApiGetCustomerRequest
*/
func (a *CustomerAPIService) GetCustomer(ctx context.Context, customerId string) ApiGetCustomerRequest {
	return ApiGetCustomerRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return Customer
func (a *CustomerAPIService) GetCustomerExecute(r ApiGetCustomerRequest) (*Customer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Customer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.GetCustomer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)

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

type ApiListCustomersRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	q *string
	date *DateRange
	page *interface{}
	limit *interface{}
}

// Consulta. Texto a buscar en &#x60;legal_name&#x60; (nombre fiscal) o en &#x60;tax_id&#x60; (RFC).
func (r ApiListCustomersRequest) Q(q string) ApiListCustomersRequest {
	r.q = &q
	return r
}

// Objeto con rango de fechas solicitado.
func (r ApiListCustomersRequest) Date(date DateRange) ApiListCustomersRequest {
	r.date = &date
	return r
}

// Página de resultados a regresar, empezando desde la página 1.
func (r ApiListCustomersRequest) Page(page interface{}) ApiListCustomersRequest {
	r.page = &page
	return r
}

// Número del 1 al 100 que representa la cantidad máxima de resultados a regresar con motivos de paginación.
func (r ApiListCustomersRequest) Limit(limit interface{}) ApiListCustomersRequest {
	r.limit = &limit
	return r
}

func (r ApiListCustomersRequest) Execute() (*CustomerSearchResult, *http.Response, error) {
	return r.ApiService.ListCustomersExecute(r)
}

/*
ListCustomers Listar clientes

Regresa una lista paginada de todos los clientes de una organización o realiza una búsqueda de acuerdo a parámetros

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCustomersRequest
*/
func (a *CustomerAPIService) ListCustomers(ctx context.Context) ApiListCustomersRequest {
	return ApiListCustomersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomerSearchResult
func (a *CustomerAPIService) ListCustomersExecute(r ApiListCustomersRequest) (*CustomerSearchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomerSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.ListCustomers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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

type ApiValidateCustomerTaxInfoRequest struct {
	ctx context.Context
	ApiService *CustomerAPIService
	customerId string
}

func (r ApiValidateCustomerTaxInfoRequest) Execute() (*ValidateCustomerTaxInfo200Response, *http.Response, error) {
	return r.ApiService.ValidateCustomerTaxInfoExecute(r)
}

/*
ValidateCustomerTaxInfo Validar información fiscal

Valida que la información fiscal del cliente coincida con los registros del SAT.

Su función principal es validar que los datos del cliente registraado siguen cumpliendo la validación del SAT.

:::tip
  Las operaciones de crear cliente, editar cliente y crear factura ya realizan una
  validación de la información del cliente, por lo que **no** es necesario llamar a este endpoint
  antes de realizar dichas operaciones.
:::


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId ID del objeto `Customer` a validar
 @return ApiValidateCustomerTaxInfoRequest
*/
func (a *CustomerAPIService) ValidateCustomerTaxInfo(ctx context.Context, customerId string) ApiValidateCustomerTaxInfoRequest {
	return ApiValidateCustomerTaxInfoRequest{
		ApiService: a,
		ctx: ctx,
		customerId: customerId,
	}
}

// Execute executes the request
//  @return ValidateCustomerTaxInfo200Response
func (a *CustomerAPIService) ValidateCustomerTaxInfoExecute(r ApiValidateCustomerTaxInfoRequest) (*ValidateCustomerTaxInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ValidateCustomerTaxInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAPIService.ValidateCustomerTaxInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customer_id}/tax-info-validation"
	localVarPath = strings.Replace(localVarPath, "{"+"customer_id"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)

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
