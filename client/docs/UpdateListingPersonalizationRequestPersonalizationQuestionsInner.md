# UpdateListingPersonalizationRequestPersonalizationQuestionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | **NullableInt64** | The ID of the personalization question. Note: This value may change if the personalization question is updated. | 
**QuestionText** | **string** | The title of the personalization question. Must be between 1 and 45 characters. Note: During the migration to the new personalization endpoints, if you&#39;re still using a legacy UI (without a title input),please send the default value &#39;Personalization&#39;. | 
**Instructions** | **NullableString** | Optional instructions for a personalization question. For legacy, single personalization, max length is 256 characters. Once multiple personalization questions are enabled, the max length will be 120 characters. | 
**QuestionType** | [**UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType**](UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType.md) |  | 
**Required** | **bool** | When true, the personalization question is required. | 
**MaxAllowedFiles** | **NullableInt64** | The maximum number of files the buyer may upload in response to a personalization question. Note: This value is only applicable to &#39;unlabeled_upload&#39; and &#39;labeled_upload&#39; questions. | 
**MaxAllowedCharacters** | **NullableInt64** | The maximum number of characters the buyer may enter in response to a personalization question. Note: This value is only applicable to &#39;text_input&#39; questions. | 
**Options** | [**[]UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner**](UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner.md) | The list of options for a personalization question. For &#39;dropdown&#39; questions, this list contains the options for the dropdown. For &#39;labeled_upload&#39; questions, this list contains the labels for the files that the buyer may upload, and must match the max_allowed_files value.. | 

## Methods

### NewUpdateListingPersonalizationRequestPersonalizationQuestionsInner

`func NewUpdateListingPersonalizationRequestPersonalizationQuestionsInner(questionId NullableInt64, questionText string, instructions NullableString, questionType UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType, required bool, maxAllowedFiles NullableInt64, maxAllowedCharacters NullableInt64, options []UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner, ) *UpdateListingPersonalizationRequestPersonalizationQuestionsInner`

NewUpdateListingPersonalizationRequestPersonalizationQuestionsInner instantiates a new UpdateListingPersonalizationRequestPersonalizationQuestionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerWithDefaults

`func NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerWithDefaults() *UpdateListingPersonalizationRequestPersonalizationQuestionsInner`

NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerWithDefaults instantiates a new UpdateListingPersonalizationRequestPersonalizationQuestionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionId() int64`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionIdOk() (*int64, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetQuestionId(v int64)`

SetQuestionId sets QuestionId field to given value.


### SetQuestionIdNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetQuestionIdNil(b bool)`

 SetQuestionIdNil sets the value for QuestionId to be an explicit nil

### UnsetQuestionId
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) UnsetQuestionId()`

UnsetQuestionId ensures that no value is present for QuestionId, not even an explicit nil
### GetQuestionText

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionText() string`

GetQuestionText returns the QuestionText field if non-nil, zero value otherwise.

### GetQuestionTextOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionTextOk() (*string, bool)`

GetQuestionTextOk returns a tuple with the QuestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionText

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetQuestionText(v string)`

SetQuestionText sets QuestionText field to given value.


### GetInstructions

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### SetInstructionsNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetQuestionType

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionType() UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType`

GetQuestionType returns the QuestionType field if non-nil, zero value otherwise.

### GetQuestionTypeOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetQuestionTypeOk() (*UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType, bool)`

GetQuestionTypeOk returns a tuple with the QuestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionType

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetQuestionType(v UpdateListingPersonalizationRequestPersonalizationQuestionsInnerQuestionType)`

SetQuestionType sets QuestionType field to given value.


### GetRequired

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetMaxAllowedFiles

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetMaxAllowedFiles() int64`

GetMaxAllowedFiles returns the MaxAllowedFiles field if non-nil, zero value otherwise.

### GetMaxAllowedFilesOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetMaxAllowedFilesOk() (*int64, bool)`

GetMaxAllowedFilesOk returns a tuple with the MaxAllowedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedFiles

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetMaxAllowedFiles(v int64)`

SetMaxAllowedFiles sets MaxAllowedFiles field to given value.


### SetMaxAllowedFilesNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetMaxAllowedFilesNil(b bool)`

 SetMaxAllowedFilesNil sets the value for MaxAllowedFiles to be an explicit nil

### UnsetMaxAllowedFiles
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) UnsetMaxAllowedFiles()`

UnsetMaxAllowedFiles ensures that no value is present for MaxAllowedFiles, not even an explicit nil
### GetMaxAllowedCharacters

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetMaxAllowedCharacters() int64`

GetMaxAllowedCharacters returns the MaxAllowedCharacters field if non-nil, zero value otherwise.

### GetMaxAllowedCharactersOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetMaxAllowedCharactersOk() (*int64, bool)`

GetMaxAllowedCharactersOk returns a tuple with the MaxAllowedCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedCharacters

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetMaxAllowedCharacters(v int64)`

SetMaxAllowedCharacters sets MaxAllowedCharacters field to given value.


### SetMaxAllowedCharactersNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetMaxAllowedCharactersNil(b bool)`

 SetMaxAllowedCharactersNil sets the value for MaxAllowedCharacters to be an explicit nil

### UnsetMaxAllowedCharacters
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) UnsetMaxAllowedCharacters()`

UnsetMaxAllowedCharacters ensures that no value is present for MaxAllowedCharacters, not even an explicit nil
### GetOptions

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetOptions() []UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) GetOptionsOk() (*[]UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetOptions(v []UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner)`

SetOptions sets Options field to given value.


### SetOptionsNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInner) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


