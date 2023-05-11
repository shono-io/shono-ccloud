# BillingV1CostListDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the schema version of this representation of a resource. | [optional] [readonly] 
**Kind** | Pointer to **string** | Kind defines the object this REST resource represents. | [optional] [readonly] 
**Id** | **string** | ID is the \&quot;natural identifier\&quot; for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\&quot;time\&quot;); however, it may collide with IDs for other object &#x60;kinds&#x60; or objects of the same &#x60;kind&#x60; within a different scope/namespace (\&quot;space\&quot;). | [readonly] 
**StartDate** | **string** | Start date of time period (inclusive) to retrieve billing costs. It is represented in RFC3339 format and is in UTC. | 
**EndDate** | **string** | End date of time period (exclusive) to retrieve billing costs. It is represented in RFC3339 format and is in UTC. | 
**Granularity** | Pointer to **string** | Granularity at which each line item is aggregated. | [optional] [default to "DAILY"]
**NetworkAccessType** | Pointer to **string** | Network access type for the cluster. | [optional] 
**Product** | Pointer to **string** | Product name. | [optional] 
**LineType** | Pointer to **string** | Type of the line item. | [optional] 
**Price** | Pointer to **float64** | Price for the line item in dollars. | [optional] 
**Unit** | **string** | Unit of the line item. | 
**Quantity** | Pointer to **float64** | Quantity of the line item. | [optional] 
**OriginalAmount** | **float64** | Original amount accrued for this line item. | 
**DiscountAmount** | Pointer to **float64** | Amount discounted from the original amount in dollars. | [optional] 
**Amount** | Pointer to **float64** | Final amount after deducting discounts. | [optional] 
**Resource** | Pointer to [**BillingV1CostResource**](BillingV1CostResource.md) |  | [optional] 

## Methods

### NewBillingV1CostListDataInner

`func NewBillingV1CostListDataInner(id string, startDate string, endDate string, unit string, originalAmount float64, ) *BillingV1CostListDataInner`

NewBillingV1CostListDataInner instantiates a new BillingV1CostListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingV1CostListDataInnerWithDefaults

`func NewBillingV1CostListDataInnerWithDefaults() *BillingV1CostListDataInner`

NewBillingV1CostListDataInnerWithDefaults instantiates a new BillingV1CostListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *BillingV1CostListDataInner) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *BillingV1CostListDataInner) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *BillingV1CostListDataInner) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *BillingV1CostListDataInner) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *BillingV1CostListDataInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BillingV1CostListDataInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BillingV1CostListDataInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *BillingV1CostListDataInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetId

`func (o *BillingV1CostListDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingV1CostListDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingV1CostListDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetStartDate

`func (o *BillingV1CostListDataInner) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingV1CostListDataInner) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingV1CostListDataInner) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *BillingV1CostListDataInner) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingV1CostListDataInner) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingV1CostListDataInner) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetGranularity

`func (o *BillingV1CostListDataInner) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *BillingV1CostListDataInner) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *BillingV1CostListDataInner) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *BillingV1CostListDataInner) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetNetworkAccessType

`func (o *BillingV1CostListDataInner) GetNetworkAccessType() string`

GetNetworkAccessType returns the NetworkAccessType field if non-nil, zero value otherwise.

### GetNetworkAccessTypeOk

`func (o *BillingV1CostListDataInner) GetNetworkAccessTypeOk() (*string, bool)`

GetNetworkAccessTypeOk returns a tuple with the NetworkAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessType

`func (o *BillingV1CostListDataInner) SetNetworkAccessType(v string)`

SetNetworkAccessType sets NetworkAccessType field to given value.

### HasNetworkAccessType

`func (o *BillingV1CostListDataInner) HasNetworkAccessType() bool`

HasNetworkAccessType returns a boolean if a field has been set.

### GetProduct

`func (o *BillingV1CostListDataInner) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *BillingV1CostListDataInner) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *BillingV1CostListDataInner) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *BillingV1CostListDataInner) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetLineType

`func (o *BillingV1CostListDataInner) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *BillingV1CostListDataInner) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *BillingV1CostListDataInner) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *BillingV1CostListDataInner) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetPrice

`func (o *BillingV1CostListDataInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingV1CostListDataInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingV1CostListDataInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingV1CostListDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetUnit

`func (o *BillingV1CostListDataInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BillingV1CostListDataInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BillingV1CostListDataInner) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetQuantity

`func (o *BillingV1CostListDataInner) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingV1CostListDataInner) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingV1CostListDataInner) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillingV1CostListDataInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *BillingV1CostListDataInner) GetOriginalAmount() float64`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *BillingV1CostListDataInner) GetOriginalAmountOk() (*float64, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *BillingV1CostListDataInner) SetOriginalAmount(v float64)`

SetOriginalAmount sets OriginalAmount field to given value.


### GetDiscountAmount

`func (o *BillingV1CostListDataInner) GetDiscountAmount() float64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *BillingV1CostListDataInner) GetDiscountAmountOk() (*float64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *BillingV1CostListDataInner) SetDiscountAmount(v float64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *BillingV1CostListDataInner) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetAmount

`func (o *BillingV1CostListDataInner) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingV1CostListDataInner) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingV1CostListDataInner) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingV1CostListDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetResource

`func (o *BillingV1CostListDataInner) GetResource() BillingV1CostResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BillingV1CostListDataInner) GetResourceOk() (*BillingV1CostResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BillingV1CostListDataInner) SetResource(v BillingV1CostResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *BillingV1CostListDataInner) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

