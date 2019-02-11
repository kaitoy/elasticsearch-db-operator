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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IndexSpec defines the desired state of Index
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html
type IndexSpec struct {
	Settings *Settings          `json:"settings,omitempty"`
	Mappings map[string]Mapping `json:"mappings"`
}

// Settings type represents settings of an index
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules.html#index-modules-settings
type Settings struct {
	Index IndexSettings `json:"index"`
}

// IndexSettings represents the index property of Settings
type IndexSettings struct {
	NumberOfShards          int32            `json:"number_of_shards,omitempty"`
	NumberOfReplicas        int32            `json:"number_of_replicas,omitempty"`
	Codec                   string           `json:"codec,omitempty"`
	RoutingPartitionSize    int32            `json:"routing_partition_size,omitempty"`
	AutoExpandReplicas      string           `json:"auto_expand_replicas,omitempty"`
	RefreshInterval         string           `json:"refresh_interval,omitempty"`
	MaxResultWindow         int32            `json:"max_result_window,omitempty"`
	MaxInnerResultWindow    int32            `json:"max_inner_result_window,omitempty"`
	MaxRescoreWindow        int32            `json:"max_rescore_window,omitempty"`
	MaxDocvalueFieldsSearch int32            `json:"max_docvalue_fields_search,omitempty"`
	MaxScriptFields         int32            `json:"max_script_fields,omitempty"`
	MaxNgramDiff            int32            `json:"max_ngram_diff,omitempty"`
	MaxShingleDiff          int32            `json:"max_shingle_diff,omitempty"`
	Blocks                  *Blocks          `json:"blocks,omitempty"`
	MaxRefreshListeners     int32            `json:"max_refresh_listeners,omitempty"`
	MaxAnalyzedOffset       int32            `json:"max_analyzed_offset,omitempty"`
	MaxTermsCount           int32            `json:"max_terms_count,omitempty"`
	Routing                 *RoutingSetting  `json:"routing,omitempty"`
	GCDeletes               *metav1.Duration `json:"gc_deletes,omitempty"`
	MaxRegexLength          int32            `json:"max_regex_length,omitempty"`
	DefaultPipeline         string           `json:"default_pipeline,omitempty"`
	Mapping                 *MappingSetting  `json:"mapping,omitempty"`
}

// Blocks represents the blocks property of an IndexSettings
type Blocks struct {
	ReadOnly            bool `json:"read_only,omitempty"`
	ReadOnlyAllowDelete bool `json:"read_only_allow_delete,omitempty"`
	Read                bool `json:"read,omitempty"`
	Write               bool `json:"write,omitempty"`
	Metadata            bool `json:"metadata,omitempty"`
}

// RoutingSetting represents the routing property of an IndexSettings
type RoutingSetting struct {
	Allocation *Allocation `json:"allocation,omitempty"`
	Rebalance  *Rebalance  `json:"rebalance,omitempty"`
}

// Allocation represents the allocation property of a RoutingSetting
type Allocation struct {
	Enable string `json:"enable,omitempty"`
}

// Rebalance represents the rebalance property of a RoutingSetting
type Rebalance struct {
	Enable string `json:"enable,omitempty"`
}

// MappingSetting represents the mapping property of an IndexSettings
type MappingSetting struct {
	TotalFields  *TotalFields  `json:"total_fields,omitempty"`
	Depth        *Depth        `json:"depth,omitempty"`
	NestedFields *NestedFields `json:"nested_fields,omitempty"`
}

// TotalFields represents the total_fields property of a MappingSetting
type TotalFields struct {
	Limit int32 `json:"limit,omitempty"`
}

// Depth represents the depth property of a MappingSetting
type Depth struct {
	Limit int32 `json:"limit,omitempty"`
}

// NestedFields represents the total_fields property of a MappingSetting
type NestedFields struct {
	Limit int32 `json:"limit,omitempty"`
}

// Mapping represents a mapping
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping.html
type Mapping struct {
	Source     *Source             `json:"_source,omitempty"`
	All        *All                `json:"_all,omitempty"`
	FieldNames *FieldNames         `json:"_field_names,omitempty"`
	Routing    *Routing            `json:"_routing,omitempty"`
	Meta       map[string]string   `json:"meta,omitempty"`
	Properties map[string]Property `json:"properties"`
}

// Routing represents the _routing property of a mapping
type Routing struct {
	Enabled bool `json:"enabled"`
}

// FieldNames represents the _field_names property of a mapping
type FieldNames struct {
	Enabled bool `json:"enabled"`
}

// All represents the _all property of a mapping
type All struct {
	Enabled bool `json:"enabled"`
}

// Source represents the _source property of a mapping
type Source struct {
	Enabled bool `json:"enabled"`
}

// Property represents a field mapping
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-types.html
// https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-params.html
type Property struct {
	Type                     string                    `json:"type,omitempty"`
	Format                   string                    `json:"format,omitempty"`
	Path                     string                    `json:"path,omitempty"`
	DocValues                bool                      `json:"doc_values,omitempty"`
	Store                    bool                      `json:"store,omitempty"`
	Coerce                   bool                      `json:"coerce,omitempty"`
	Boost                    float32                   `json:"boost,omitempty"`
	Index                    bool                      `json:"index,omitempty"`
	NullValue                string                    `json:"null_value,omitempty"`
	Locale                   string                    `json:"locale,omitempty"`
	IgnoreMalformed          bool                      `json:"ignore_malformed,omitempty"`
	IgnoreZValue             bool                      `json:"ignore_z_value,omitempty"`
	Tree                     string                    `json:"tree,omitempty"`
	Precision                string                    `json:"precision,omitempty"`
	TreeLevels               int32                     `json:"tree_levels,omitempty"`
	Strategy                 string                    `json:"strategy,omitempty"`
	DistanceErrorPct         float64                   `json:"distance_error_pct,omitempty"`
	Orientation              string                    `json:"orientation,omitempty"`
	PointsOnly               bool                      `json:"points_only,omitempty"`
	EagerGlobalOrdinals      bool                      `json:"eager_global_ordinals,omitempty"`
	Fields                   map[string]Property       `json:"fields,omitempty"`
	IgnoreAbove              int32                     `json:"ignore_above,omitempty"`
	IndexOptions             string                    `json:"index_options,omitempty"`
	Norms                    bool                      `json:"norms,omitempty"`
	Similarity               string                    `json:"similarity,omitempty"`
	Normalizer               string                    `json:"normalizer,omitempty"`
	SplitQueriesOnWhitespace bool                      `json:"split_queries_on_whitespace,omitempty"`
	Dynamic                  string                    `json:"dynamic,omitempty"`
	Properties               map[string]Property       `json:"properties,omitempty"`
	ScalingFactor            float64                   `json:"scaling_factor,omitempty"`
	Enabled                  bool                      `json:"enabled,omitempty"`
	Analyzer                 string                    `json:"analyzer,omitempty"`
	Fielddata                bool                      `json:"fielddata,omitempty"`
	FielddataFrequencyFilter *FielddataFrequencyFilter `json:"fielddata_frequency_filter,omitempty"`
	IndexPrefixes            *IndexPrefixes            `json:"index_prefixes,omitempty"`
	IndexPhrases             bool                      `json:"index_phrases,omitempty"`
	PositionIncrementGap     int32                     `json:"position_increment_gap,omitempty"`
	SearchAnalyzer           string                    `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer      string                    `json:"search_quote_analyzer,omitempty"`
	TermVector               string                    `json:"term_vector,omitempty"`
	EnablePositionIncrements bool                      `json:"enable_position_increments,omitempty"`
	Relations                map[string][]string       `json:"relations,omitempty"`
	CopyTo                   map[string][]string       `json:"copy_to,omitempty"`
}

// FielddataFrequencyFilter represents a fielddata_frequency_filter parameter of a mapping
// https://www.elastic.co/guide/en/elasticsearch/reference/current/fielddata.html#field-data-filtering
type FielddataFrequencyFilter struct {
	Min            float64 `json:"min"`
	Max            float64 `json:"max"`
	MinSegmentSize int32   `json:"min_segment_size,omitempty"`
}

// IndexPrefixes represents a index_prefixes parameter of a mapping
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-prefixes.html
type IndexPrefixes struct {
	MinChars int32 `json:"min_chars,omitempty"`
	MaxChars int32 `json:"max_chars,omitempty"`
}

// IndexStatus defines the observed state of Index
type IndexStatus struct {
	Closed     bool             `json:"closed"`
	Conditions []IndexCondition `json:"conditions"`
}

// IndexCondition represents information about the status of an index
type IndexCondition struct {
	StatusCode         int         `json:"statusCode"`
	Status             string      `json:"status"`
	LastProbeTime      metav1.Time `json:"lastProbeTime"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Index is the Schema for the indices API
// +k8s:openapi-gen=true
type Index struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	URL struct {
		ElasticsearchEndpoint string `json:"elasticsearchEndpoint"`
		Index                 string `json:"index"`
	} `json:"url"`
	Spec   IndexSpec   `json:"spec,omitempty"`
	Status IndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IndexList contains a list of Index
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Index `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Index{}, &IndexList{})
}
