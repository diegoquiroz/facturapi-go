# RelatedDocumentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationship** | **string** | Clave de relación del catálogo del SAT (que puedes consultar en [esta tabla](#relacion-entre-facturas)). Es requerido cuando se envíe el parámetro &#x60;related&#x60;. | 
**Documents** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRelatedDocumentInput

`func NewRelatedDocumentInput(relationship string, ) *RelatedDocumentInput`

NewRelatedDocumentInput instantiates a new RelatedDocumentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedDocumentInputWithDefaults

`func NewRelatedDocumentInputWithDefaults() *RelatedDocumentInput`

NewRelatedDocumentInputWithDefaults instantiates a new RelatedDocumentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationship

`func (o *RelatedDocumentInput) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RelatedDocumentInput) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RelatedDocumentInput) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.


### GetDocuments

`func (o *RelatedDocumentInput) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *RelatedDocumentInput) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *RelatedDocumentInput) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *RelatedDocumentInput) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


