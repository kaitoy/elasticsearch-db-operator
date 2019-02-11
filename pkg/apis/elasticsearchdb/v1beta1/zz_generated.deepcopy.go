// +build !ignore_autogenerated

/*
Copyright 2019 kaitoy.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *All) DeepCopyInto(out *All) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new All.
func (in *All) DeepCopy() *All {
	if in == nil {
		return nil
	}
	out := new(All)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Allocation) DeepCopyInto(out *Allocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Allocation.
func (in *Allocation) DeepCopy() *Allocation {
	if in == nil {
		return nil
	}
	out := new(Allocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blocks) DeepCopyInto(out *Blocks) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blocks.
func (in *Blocks) DeepCopy() *Blocks {
	if in == nil {
		return nil
	}
	out := new(Blocks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Depth) DeepCopyInto(out *Depth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Depth.
func (in *Depth) DeepCopy() *Depth {
	if in == nil {
		return nil
	}
	out := new(Depth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FieldNames) DeepCopyInto(out *FieldNames) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FieldNames.
func (in *FieldNames) DeepCopy() *FieldNames {
	if in == nil {
		return nil
	}
	out := new(FieldNames)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FielddataFrequencyFilter) DeepCopyInto(out *FielddataFrequencyFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FielddataFrequencyFilter.
func (in *FielddataFrequencyFilter) DeepCopy() *FielddataFrequencyFilter {
	if in == nil {
		return nil
	}
	out := new(FielddataFrequencyFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Index) DeepCopyInto(out *Index) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.URL = in.URL
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Index.
func (in *Index) DeepCopy() *Index {
	if in == nil {
		return nil
	}
	out := new(Index)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Index) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexCondition) DeepCopyInto(out *IndexCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexCondition.
func (in *IndexCondition) DeepCopy() *IndexCondition {
	if in == nil {
		return nil
	}
	out := new(IndexCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexList) DeepCopyInto(out *IndexList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Index, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexList.
func (in *IndexList) DeepCopy() *IndexList {
	if in == nil {
		return nil
	}
	out := new(IndexList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexPrefixes) DeepCopyInto(out *IndexPrefixes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexPrefixes.
func (in *IndexPrefixes) DeepCopy() *IndexPrefixes {
	if in == nil {
		return nil
	}
	out := new(IndexPrefixes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexSettings) DeepCopyInto(out *IndexSettings) {
	*out = *in
	if in.Blocks != nil {
		in, out := &in.Blocks, &out.Blocks
		*out = new(Blocks)
		**out = **in
	}
	if in.Routing != nil {
		in, out := &in.Routing, &out.Routing
		*out = new(RoutingSetting)
		(*in).DeepCopyInto(*out)
	}
	if in.GCDeletes != nil {
		in, out := &in.GCDeletes, &out.GCDeletes
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Mapping != nil {
		in, out := &in.Mapping, &out.Mapping
		*out = new(MappingSetting)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexSettings.
func (in *IndexSettings) DeepCopy() *IndexSettings {
	if in == nil {
		return nil
	}
	out := new(IndexSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexSpec) DeepCopyInto(out *IndexSpec) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Mappings != nil {
		in, out := &in.Mappings, &out.Mappings
		*out = make(map[string]Mapping, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexSpec.
func (in *IndexSpec) DeepCopy() *IndexSpec {
	if in == nil {
		return nil
	}
	out := new(IndexSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexStatus) DeepCopyInto(out *IndexStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]IndexCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexStatus.
func (in *IndexStatus) DeepCopy() *IndexStatus {
	if in == nil {
		return nil
	}
	out := new(IndexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mapping) DeepCopyInto(out *Mapping) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(Source)
		**out = **in
	}
	if in.All != nil {
		in, out := &in.All, &out.All
		*out = new(All)
		**out = **in
	}
	if in.FieldNames != nil {
		in, out := &in.FieldNames, &out.FieldNames
		*out = new(FieldNames)
		**out = **in
	}
	if in.Routing != nil {
		in, out := &in.Routing, &out.Routing
		*out = new(Routing)
		**out = **in
	}
	if in.Meta != nil {
		in, out := &in.Meta, &out.Meta
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]Property, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mapping.
func (in *Mapping) DeepCopy() *Mapping {
	if in == nil {
		return nil
	}
	out := new(Mapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MappingSetting) DeepCopyInto(out *MappingSetting) {
	*out = *in
	if in.TotalFields != nil {
		in, out := &in.TotalFields, &out.TotalFields
		*out = new(TotalFields)
		**out = **in
	}
	if in.Depth != nil {
		in, out := &in.Depth, &out.Depth
		*out = new(Depth)
		**out = **in
	}
	if in.NestedFields != nil {
		in, out := &in.NestedFields, &out.NestedFields
		*out = new(NestedFields)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MappingSetting.
func (in *MappingSetting) DeepCopy() *MappingSetting {
	if in == nil {
		return nil
	}
	out := new(MappingSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NestedFields) DeepCopyInto(out *NestedFields) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NestedFields.
func (in *NestedFields) DeepCopy() *NestedFields {
	if in == nil {
		return nil
	}
	out := new(NestedFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make(map[string]Property, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]Property, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.FielddataFrequencyFilter != nil {
		in, out := &in.FielddataFrequencyFilter, &out.FielddataFrequencyFilter
		*out = new(FielddataFrequencyFilter)
		**out = **in
	}
	if in.IndexPrefixes != nil {
		in, out := &in.IndexPrefixes, &out.IndexPrefixes
		*out = new(IndexPrefixes)
		**out = **in
	}
	if in.Relations != nil {
		in, out := &in.Relations, &out.Relations
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.CopyTo != nil {
		in, out := &in.CopyTo, &out.CopyTo
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rebalance) DeepCopyInto(out *Rebalance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rebalance.
func (in *Rebalance) DeepCopy() *Rebalance {
	if in == nil {
		return nil
	}
	out := new(Rebalance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Routing) DeepCopyInto(out *Routing) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Routing.
func (in *Routing) DeepCopy() *Routing {
	if in == nil {
		return nil
	}
	out := new(Routing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutingSetting) DeepCopyInto(out *RoutingSetting) {
	*out = *in
	if in.Allocation != nil {
		in, out := &in.Allocation, &out.Allocation
		*out = new(Allocation)
		**out = **in
	}
	if in.Rebalance != nil {
		in, out := &in.Rebalance, &out.Rebalance
		*out = new(Rebalance)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutingSetting.
func (in *RoutingSetting) DeepCopy() *RoutingSetting {
	if in == nil {
		return nil
	}
	out := new(RoutingSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	in.Index.DeepCopyInto(&out.Index)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.URL = in.URL
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Template) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateCondition) DeepCopyInto(out *TemplateCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateCondition.
func (in *TemplateCondition) DeepCopy() *TemplateCondition {
	if in == nil {
		return nil
	}
	out := new(TemplateCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateList) DeepCopyInto(out *TemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateList.
func (in *TemplateList) DeepCopy() *TemplateList {
	if in == nil {
		return nil
	}
	out := new(TemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	if in.IndexPatterns != nil {
		in, out := &in.IndexPatterns, &out.IndexPatterns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Mappings != nil {
		in, out := &in.Mappings, &out.Mappings
		*out = make(map[string]Mapping, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStatus) DeepCopyInto(out *TemplateStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TemplateCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStatus.
func (in *TemplateStatus) DeepCopy() *TemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TotalFields) DeepCopyInto(out *TotalFields) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TotalFields.
func (in *TotalFields) DeepCopy() *TotalFields {
	if in == nil {
		return nil
	}
	out := new(TotalFields)
	in.DeepCopyInto(out)
	return out
}
