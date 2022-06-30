/*
main

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MainCreateUserInput struct for MainCreateUserInput
type MainCreateUserInput struct {
	Name string `json:"name"`
	NickName string `json:"nickName"`
	Phone string `json:"phone"`
}

// NewMainCreateUserInput instantiates a new MainCreateUserInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainCreateUserInput(name string, nickName string, phone string) *MainCreateUserInput {
	this := MainCreateUserInput{}
	this.Name = name
	this.NickName = nickName
	this.Phone = phone
	return &this
}

// NewMainCreateUserInputWithDefaults instantiates a new MainCreateUserInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainCreateUserInputWithDefaults() *MainCreateUserInput {
	this := MainCreateUserInput{}
	return &this
}

// GetName returns the Name field value
func (o *MainCreateUserInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MainCreateUserInput) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MainCreateUserInput) SetName(v string) {
	o.Name = v
}

// GetNickName returns the NickName field value
func (o *MainCreateUserInput) GetNickName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value
// and a boolean to check if the value has been set.
func (o *MainCreateUserInput) GetNickNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NickName, true
}

// SetNickName sets field value
func (o *MainCreateUserInput) SetNickName(v string) {
	o.NickName = v
}

// GetPhone returns the Phone field value
func (o *MainCreateUserInput) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *MainCreateUserInput) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *MainCreateUserInput) SetPhone(v string) {
	o.Phone = v
}

func (o MainCreateUserInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["nickName"] = o.NickName
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableMainCreateUserInput struct {
	value *MainCreateUserInput
	isSet bool
}

func (v NullableMainCreateUserInput) Get() *MainCreateUserInput {
	return v.value
}

func (v *NullableMainCreateUserInput) Set(val *MainCreateUserInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMainCreateUserInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMainCreateUserInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainCreateUserInput(val *MainCreateUserInput) *NullableMainCreateUserInput {
	return &NullableMainCreateUserInput{value: val, isSet: true}
}

func (v NullableMainCreateUserInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainCreateUserInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

