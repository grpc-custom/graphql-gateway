package scalar

import (
	"github.com/graphql-go/graphql"
)

// Empty a generic empty message
// see https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/empty.proto
var Empty = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Empty",
	Description: "google.protobuf.Empty",
	Serialize: func(_ interface{}) interface{} {
		empty := struct{}{}
		return empty
	},
})

// Any
// see https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/any.proto
var Any = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Any",
	Description: "google.protobuf.Any",
	Fields: graphql.Fields{
		"type_url": &graphql.Field{
			Type: String,
		},
		"value": &graphql.Field{
			Type: Bytes,
		},
	},
})

// Timestamp
// see https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/timestamp.proto
var Timestamp = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Timestamp",
	Description: "google.protobuf.Timestamp",
	Fields: graphql.Fields{
		"seconds": &graphql.Field{
			Type: graphql.NewNonNull(Int64),
		},
		"nanos": &graphql.Field{
			Type: graphql.NewNonNull(Int32),
		},
	},
})

// Duration
// see https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/duration.proto
var Duration = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Duration",
	Description: "google.protobuf.Duration",
	Fields: graphql.Fields{
		"seconds": &graphql.Field{
			Type: graphql.NewNonNull(Int64),
		},
		"nanos": &graphql.Field{
			Type: graphql.NewNonNull(Int32),
		},
	},
})
