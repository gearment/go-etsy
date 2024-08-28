# ShippingCarrier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShippingCarrierId** | Pointer to **int64** | The numeric ID of this shipping carrier. | [optional] 
**Name** | Pointer to **string** | The name of this shipping carrier. | [optional] 
**DomesticClasses** | Pointer to [**[]ShippingCarrierMailClass**](ShippingCarrierMailClass.md) | Set of domestic mail classes of this shipping carrier. | [optional] 
**InternationalClasses** | Pointer to [**[]ShippingCarrierMailClass**](ShippingCarrierMailClass.md) | Set of international mail classes of this shipping carrier. | [optional] 

## Methods

### NewShippingCarrier

`func NewShippingCarrier() *ShippingCarrier`

NewShippingCarrier instantiates a new ShippingCarrier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingCarrierWithDefaults

`func NewShippingCarrierWithDefaults() *ShippingCarrier`

NewShippingCarrierWithDefaults instantiates a new ShippingCarrier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShippingCarrierId

`func (o *ShippingCarrier) GetShippingCarrierId() int64`

GetShippingCarrierId returns the ShippingCarrierId field if non-nil, zero value otherwise.

### GetShippingCarrierIdOk

`func (o *ShippingCarrier) GetShippingCarrierIdOk() (*int64, bool)`

GetShippingCarrierIdOk returns a tuple with the ShippingCarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierId

`func (o *ShippingCarrier) SetShippingCarrierId(v int64)`

SetShippingCarrierId sets ShippingCarrierId field to given value.

### HasShippingCarrierId

`func (o *ShippingCarrier) HasShippingCarrierId() bool`

HasShippingCarrierId returns a boolean if a field has been set.

### GetName

`func (o *ShippingCarrier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShippingCarrier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShippingCarrier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShippingCarrier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomesticClasses

`func (o *ShippingCarrier) GetDomesticClasses() []ShippingCarrierMailClass`

GetDomesticClasses returns the DomesticClasses field if non-nil, zero value otherwise.

### GetDomesticClassesOk

`func (o *ShippingCarrier) GetDomesticClassesOk() (*[]ShippingCarrierMailClass, bool)`

GetDomesticClassesOk returns a tuple with the DomesticClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticClasses

`func (o *ShippingCarrier) SetDomesticClasses(v []ShippingCarrierMailClass)`

SetDomesticClasses sets DomesticClasses field to given value.

### HasDomesticClasses

`func (o *ShippingCarrier) HasDomesticClasses() bool`

HasDomesticClasses returns a boolean if a field has been set.

### GetInternationalClasses

`func (o *ShippingCarrier) GetInternationalClasses() []ShippingCarrierMailClass`

GetInternationalClasses returns the InternationalClasses field if non-nil, zero value otherwise.

### GetInternationalClassesOk

`func (o *ShippingCarrier) GetInternationalClassesOk() (*[]ShippingCarrierMailClass, bool)`

GetInternationalClassesOk returns a tuple with the InternationalClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalClasses

`func (o *ShippingCarrier) SetInternationalClasses(v []ShippingCarrierMailClass)`

SetInternationalClasses sets InternationalClasses field to given value.

### HasInternationalClasses

`func (o *ShippingCarrier) HasInternationalClasses() bool`

HasInternationalClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


