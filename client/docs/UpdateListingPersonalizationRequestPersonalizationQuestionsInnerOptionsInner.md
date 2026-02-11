# UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionId** | **NullableInt64** | The ID of the option. Note: This value may change if the option or question is updated. | 
**Label** | **string** | The option label. Note: For &#39;dropdown&#39; questions, max length is 20 characters. For &#39;labeled_upload&#39; questions, max length is 45 characters. | 

## Methods

### NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner

`func NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner(optionId NullableInt64, label string, ) *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner`

NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner instantiates a new UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInnerWithDefaults

`func NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInnerWithDefaults() *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner`

NewUpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInnerWithDefaults instantiates a new UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionId

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) GetOptionId() int64`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) GetOptionIdOk() (*int64, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) SetOptionId(v int64)`

SetOptionId sets OptionId field to given value.


### SetOptionIdNil

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) SetOptionIdNil(b bool)`

 SetOptionIdNil sets the value for OptionId to be an explicit nil

### UnsetOptionId
`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) UnsetOptionId()`

UnsetOptionId ensures that no value is present for OptionId, not even an explicit nil
### GetLabel

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateListingPersonalizationRequestPersonalizationQuestionsInnerOptionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


