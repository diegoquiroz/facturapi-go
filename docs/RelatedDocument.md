# RelatedDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationship** | Pointer to **string** | Clave de relación del catálogo del SAT (que puedes consultar en [esta tabla](#relacion-entre-facturas)). Es requerido cuando se envíe el parámetro &#x60;related&#x60;. | [optional] 
**Documents** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRelatedDocument

`func NewRelatedDocument() *RelatedDocument`

NewRelatedDocument instantiates a new RelatedDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedDocumentWithDefaults

`func NewRelatedDocumentWithDefaults() *RelatedDocument`

NewRelatedDocumentWithDefaults instantiates a new RelatedDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationship

`func (o *RelatedDocument) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RelatedDocument) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RelatedDocument) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *RelatedDocument) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetDocuments

`func (o *RelatedDocument) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *RelatedDocument) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *RelatedDocument) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *RelatedDocument) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


