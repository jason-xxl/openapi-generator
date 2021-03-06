/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// OuterEnumDefaultValue the model 'OuterEnumDefaultValue'
type OuterEnumDefaultValue string

// List of OuterEnumDefaultValue
const (
	OUTERENUMDEFAULTVALUE_PLACED OuterEnumDefaultValue = "placed"
	OUTERENUMDEFAULTVALUE_APPROVED OuterEnumDefaultValue = "approved"
	OUTERENUMDEFAULTVALUE_DELIVERED OuterEnumDefaultValue = "delivered"
)

// Ptr returns reference to OuterEnumDefaultValue value
func (v OuterEnumDefaultValue) Ptr() *OuterEnumDefaultValue {
	return &v
}


type NullableOuterEnumDefaultValue struct {
	Value OuterEnumDefaultValue
	ExplicitNull bool
}

func (v NullableOuterEnumDefaultValue) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOuterEnumDefaultValue) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
