# ConsolidateShopReturnPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceReturnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 
**DestinationReturnPolicyId** | **int64** | The numeric ID of the [Return Policy](/documentation/reference#operation/getShopReturnPolicies). | 

## Methods

### NewConsolidateShopReturnPoliciesRequest

`func NewConsolidateShopReturnPoliciesRequest(sourceReturnPolicyId int64, destinationReturnPolicyId int64, ) *ConsolidateShopReturnPoliciesRequest`

NewConsolidateShopReturnPoliciesRequest instantiates a new ConsolidateShopReturnPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolidateShopReturnPoliciesRequestWithDefaults

`func NewConsolidateShopReturnPoliciesRequestWithDefaults() *ConsolidateShopReturnPoliciesRequest`

NewConsolidateShopReturnPoliciesRequestWithDefaults instantiates a new ConsolidateShopReturnPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceReturnPolicyId

`func (o *ConsolidateShopReturnPoliciesRequest) GetSourceReturnPolicyId() int64`

GetSourceReturnPolicyId returns the SourceReturnPolicyId field if non-nil, zero value otherwise.

### GetSourceReturnPolicyIdOk

`func (o *ConsolidateShopReturnPoliciesRequest) GetSourceReturnPolicyIdOk() (*int64, bool)`

GetSourceReturnPolicyIdOk returns a tuple with the SourceReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReturnPolicyId

`func (o *ConsolidateShopReturnPoliciesRequest) SetSourceReturnPolicyId(v int64)`

SetSourceReturnPolicyId sets SourceReturnPolicyId field to given value.


### GetDestinationReturnPolicyId

`func (o *ConsolidateShopReturnPoliciesRequest) GetDestinationReturnPolicyId() int64`

GetDestinationReturnPolicyId returns the DestinationReturnPolicyId field if non-nil, zero value otherwise.

### GetDestinationReturnPolicyIdOk

`func (o *ConsolidateShopReturnPoliciesRequest) GetDestinationReturnPolicyIdOk() (*int64, bool)`

GetDestinationReturnPolicyIdOk returns a tuple with the DestinationReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationReturnPolicyId

`func (o *ConsolidateShopReturnPoliciesRequest) SetDestinationReturnPolicyId(v int64)`

SetDestinationReturnPolicyId sets DestinationReturnPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


