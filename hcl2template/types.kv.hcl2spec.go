// Code generated by "mapstructure-to-hcl2 -type NameValue,NameValues,KVFilter"; DO NOT EDIT.
package hcl2template

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatKVFilter is an auto-generated flat version of KVFilter.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatKVFilter struct {
	Filters map[string]string `cty:"filters"`
	Filter  []FlatNameValue   `cty:"filter"`
}

// FlatMapstructure returns a new FlatKVFilter.
// FlatKVFilter is an auto-generated flat version of KVFilter.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*KVFilter) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatKVFilter)
}

// HCL2Spec returns the hcl spec of a KVFilter.
// This spec is used by HCL to read the fields of KVFilter.
// The decoded values from this spec will then be applied to a FlatKVFilter.
func (*FlatKVFilter) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.BlockAttrsSpec{TypeName: "filters", ElementType: cty.String, Required: false},
		"filter":  &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*FlatNameValue)(nil).HCL2Spec())},
	}
	return s
}

// FlatNameValue is an auto-generated flat version of NameValue.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatNameValue struct {
	Name  *string `cty:"name"`
	Value *string `cty:"value"`
}

// FlatMapstructure returns a new FlatNameValue.
// FlatNameValue is an auto-generated flat version of NameValue.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*NameValue) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatNameValue)
}

// HCL2Spec returns the hcl spec of a NameValue.
// This spec is used by HCL to read the fields of NameValue.
// The decoded values from this spec will then be applied to a FlatNameValue.
func (*FlatNameValue) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"name":  &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"value": &hcldec.AttrSpec{Name: "value", Type: cty.String, Required: false},
	}
	return s
}
