# NominaPercepcionesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percepcion** | [**[]NominaPercepcionInput**](NominaPercepcionInput.md) |  | 
**JubilacionPensionRetiro** | Pointer to [**NominaJubilacionInput**](NominaJubilacionInput.md) |  | [optional] 
**SeparacionIndemnizacion** | Pointer to [**NominaSeparacionInput**](NominaSeparacionInput.md) |  | [optional] 

## Methods

### NewNominaPercepcionesInput

`func NewNominaPercepcionesInput(percepcion []NominaPercepcionInput, ) *NominaPercepcionesInput`

NewNominaPercepcionesInput instantiates a new NominaPercepcionesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominaPercepcionesInputWithDefaults

`func NewNominaPercepcionesInputWithDefaults() *NominaPercepcionesInput`

NewNominaPercepcionesInputWithDefaults instantiates a new NominaPercepcionesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercepcion

`func (o *NominaPercepcionesInput) GetPercepcion() []NominaPercepcionInput`

GetPercepcion returns the Percepcion field if non-nil, zero value otherwise.

### GetPercepcionOk

`func (o *NominaPercepcionesInput) GetPercepcionOk() (*[]NominaPercepcionInput, bool)`

GetPercepcionOk returns a tuple with the Percepcion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercepcion

`func (o *NominaPercepcionesInput) SetPercepcion(v []NominaPercepcionInput)`

SetPercepcion sets Percepcion field to given value.


### GetJubilacionPensionRetiro

`func (o *NominaPercepcionesInput) GetJubilacionPensionRetiro() NominaJubilacionInput`

GetJubilacionPensionRetiro returns the JubilacionPensionRetiro field if non-nil, zero value otherwise.

### GetJubilacionPensionRetiroOk

`func (o *NominaPercepcionesInput) GetJubilacionPensionRetiroOk() (*NominaJubilacionInput, bool)`

GetJubilacionPensionRetiroOk returns a tuple with the JubilacionPensionRetiro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJubilacionPensionRetiro

`func (o *NominaPercepcionesInput) SetJubilacionPensionRetiro(v NominaJubilacionInput)`

SetJubilacionPensionRetiro sets JubilacionPensionRetiro field to given value.

### HasJubilacionPensionRetiro

`func (o *NominaPercepcionesInput) HasJubilacionPensionRetiro() bool`

HasJubilacionPensionRetiro returns a boolean if a field has been set.

### GetSeparacionIndemnizacion

`func (o *NominaPercepcionesInput) GetSeparacionIndemnizacion() NominaSeparacionInput`

GetSeparacionIndemnizacion returns the SeparacionIndemnizacion field if non-nil, zero value otherwise.

### GetSeparacionIndemnizacionOk

`func (o *NominaPercepcionesInput) GetSeparacionIndemnizacionOk() (*NominaSeparacionInput, bool)`

GetSeparacionIndemnizacionOk returns a tuple with the SeparacionIndemnizacion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparacionIndemnizacion

`func (o *NominaPercepcionesInput) SetSeparacionIndemnizacion(v NominaSeparacionInput)`

SetSeparacionIndemnizacion sets SeparacionIndemnizacion field to given value.

### HasSeparacionIndemnizacion

`func (o *NominaPercepcionesInput) HasSeparacionIndemnizacion() bool`

HasSeparacionIndemnizacion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


