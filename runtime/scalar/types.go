package scalar

import (
	"encoding/base64"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func coerceFloat32(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return float32(1)
		}
		return float32(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case int8:
		return float32(value)
	case *int8:
		if value == nil {
			return nil
		}
		return float32(*value)
	case int16:
		return float32(value)
	case *int16:
		if value == nil {
			return nil
		}
		return float32(*value)
	case int32:
		return value
	case *int32:
		if value == nil {
			return nil
		}
		return *value
	case int64:
		return float32(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case uint:
		return float32(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case uint8:
		return float32(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return float32(*value)
	case uint16:
		return float32(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return float32(*value)
	case uint32:
		return float32(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case uint64:
		return float32(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case float32:
		return value
	case *float32:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case float64:
		return float32(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceFloat32(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceFloat32(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

func coerceFloat64(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return float64(1)
		}
		return float64(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case int8:
		return float64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return float64(*value)
	case int16:
		return float64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return float64(*value)
	case int32:
		return value
	case *int32:
		if value == nil {
			return nil
		}
		return *value
	case int64:
		return float64(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case uint:
		return float64(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case uint8:
		return float64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return float64(*value)
	case uint16:
		return float64(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return float64(*value)
	case uint32:
		return float64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case uint64:
		return float64(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case float32:
		return float64(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case float64:
		return value
	case *float64:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceFloat64(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceFloat64(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

func coerceInt32(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return int32(1)
		}
		return int32(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case int8:
		return int32(value)
	case *int8:
		if value == nil {
			return nil
		}
		return int32(*value)
	case int16:
		return int32(value)
	case *int16:
		if value == nil {
			return nil
		}
		return int32(*value)
	case int32:
		return value
	case *int32:
		if value == nil {
			return nil
		}
		return *value
	case int64:
		return int32(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case uint:
		return int32(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case uint8:
		return int32(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return int32(*value)
	case uint16:
		return int32(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return int32(*value)
	case uint32:
		return int32(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case uint64:
		return int32(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case float32:
		return int32(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case float64:
		return int32(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceInt32(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceInt32(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

func coerceUint32(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return uint32(1)
		}
		return uint32(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case int8:
		return uint32(value)
	case *int8:
		if value == nil {
			return nil
		}
		return uint32(*value)
	case int16:
		return uint32(value)
	case *int16:
		if value == nil {
			return nil
		}
		return uint32(*value)
	case int32:
		return uint32(value)
	case *int32:
		if value == nil {
			return nil
		}
		return *value
	case int64:
		return uint32(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case uint:
		return uint32(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case uint8:
		return uint32(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return uint32(*value)
	case uint16:
		return uint32(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return uint32(*value)
	case uint32:
		return value
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case uint64:
		return uint32(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case float32:
		return uint32(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case float64:
		return uint32(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceUint32(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceUint32(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

func coerceInt64(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return int64(1)
		}
		return int64(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case int8:
		return int64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int16:
		return int64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int32:
		return int64(value)
	case *int32:
		if value == nil {
			return nil
		}
		return int64(*value)
	case int64:
		return value
	case *int64:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case uint:
		return int64(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case uint8:
		return int64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint16:
		return int64(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return int64(*value)
	case uint32:
		return int64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case uint64:
		return int64(value)
	case *uint64:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case float32:
		return int64(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case float64:
		return int64(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceInt64(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceInt64(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

func coerceUint64(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return uint64(1)
		}
		return uint64(0)
	case *bool:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case int:
		return value
	case *int:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case int8:
		return uint64(value)
	case *int8:
		if value == nil {
			return nil
		}
		return uint64(*value)
	case int16:
		return uint64(value)
	case *int16:
		if value == nil {
			return nil
		}
		return uint64(*value)
	case int32:
		return uint64(value)
	case *int32:
		if value == nil {
			return nil
		}
		return uint64(*value)
	case int64:
		return uint64(value)
	case *int64:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case uint:
		return uint64(value)
	case *uint:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case uint8:
		return uint64(value)
	case *uint8:
		if value == nil {
			return nil
		}
		return uint64(*value)
	case uint16:
		return uint64(value)
	case *uint16:
		if value == nil {
			return nil
		}
		return uint64(*value)
	case uint32:
		return uint64(value)
	case *uint32:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case uint64:
		return value
	case *uint64:
		if value == nil {
			return nil
		}
		return *value
	case float32:
		return uint64(value)
	case *float32:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case float64:
		return uint64(value)
	case *float64:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceUint64(val)
	case *string:
		if value == nil {
			return nil
		}
		return coerceUint64(*value)
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

var String = graphql.String

var Bool = graphql.Boolean

var Float32 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Float",
	Description: "float32",
	Serialize:   coerceFloat32,
	ParseValue:  coerceFloat32,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.FloatValue:
			if v, err := strconv.ParseFloat(val.Value, 32); err == nil {
				return float32(v)
			}
		}
		return nil
	},
})

var Float64 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Double",
	Description: "float64",
	Serialize:   coerceFloat64,
	ParseValue:  coerceFloat64,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.FloatValue:
			if v, err := strconv.ParseFloat(val.Value, 64); err == nil {
				return v
			}
		}
		return nil
	},
})

var Int32 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Int32",
	Description: "int32",
	Serialize:   coerceInt32,
	ParseValue:  coerceInt32,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.IntValue:
			if v, err := strconv.ParseInt(val.Value, 10, 32); err == nil {
				return int32(v)
			}
		}
		return nil
	},
})

var Uint32 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Uint32",
	Description: "uint32",
	Serialize:   coerceUint32,
	ParseValue:  coerceUint32,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.IntValue:
			if v, err := strconv.ParseUint(val.Value, 10, 32); err == nil {
				return uint32(v)
			}
		}
		return nil
	},
})

var Int64 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Int64",
	Description: "int64",
	Serialize:   coerceInt64,
	ParseValue:  coerceInt64,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.IntValue:
			if v, err := strconv.ParseInt(val.Value, 10, 64); err == nil {
				return v
			}
		}
		return nil
	},
})

var Uint64 = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Uint64",
	Description: "uint64",
	Serialize:   coerceUint64,
	ParseValue:  coerceUint64,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch val := valueAST.(type) {
		case *ast.IntValue:
			if v, err := strconv.ParseUint(val.Value, 10, 64); err == nil {
				return v
			}
		}
		return nil
	},
})

var Bytes = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "Bytes",
	Description: "bytes",
	Serialize: func(value interface{}) interface{} {
		src, ok := value.([]byte)
		if !ok {
			return "nil"
		}
		return base64.StdEncoding.EncodeToString(src)
	},
})
