/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the CollectionOfActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfActivities{}

// CollectionOfActivities struct for CollectionOfActivities
type CollectionOfActivities struct {
	Value []Activity `json:"value,omitempty"`
}

// NewCollectionOfActivities instantiates a new CollectionOfActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfActivities() *CollectionOfActivities {
	this := CollectionOfActivities{}
	return &this
}

// NewCollectionOfActivitiesWithDefaults instantiates a new CollectionOfActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfActivitiesWithDefaults() *CollectionOfActivities {
	this := CollectionOfActivities{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfActivities) GetValue() []Activity {
	if o == nil || IsNil(o.Value) {
		var ret []Activity
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfActivities) GetValueOk() ([]Activity, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfActivities) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Activity and assigns it to the Value field.
func (o *CollectionOfActivities) SetValue(v []Activity) {
	o.Value = v
}

func (o CollectionOfActivities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCollectionOfActivities struct {
	value *CollectionOfActivities
	isSet bool
}

func (v NullableCollectionOfActivities) Get() *CollectionOfActivities {
	return v.value
}

func (v *NullableCollectionOfActivities) Set(val *CollectionOfActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfActivities(val *CollectionOfActivities) *NullableCollectionOfActivities {
	return &NullableCollectionOfActivities{value: val, isSet: true}
}

func (v NullableCollectionOfActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


