package gqlshared

import (
	"reflect"

	"github.com/graphql-go/graphql"
)

// BuildFields accepts any struct and will use reflection to construct the equivalent GraphQL fields
func BuildFields(obj interface{}) *graphql.Fields {
	fields := graphql.Fields{}

	if objType := reflect.TypeOf(obj); objType.Kind() == reflect.Struct {
		for index := 0; index < objType.NumField(); index++ {
			field := objType.Field(index)

			key := field.Tag.Get("gql")
			if len(key) == 0 {
				key = field.Tag.Get("json")
				if len(key) == 0 {
					key = field.Name
				}
			}

			fields[key] = &graphql.Field{
				Type:        getGraphQLType(field),
				Description: key,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					v := reflect.ValueOf(p.Source)
					val := reflect.Indirect(v).FieldByName(field.Name)

					switch val.Kind() {
					case reflect.Int, reflect.Int32, reflect.Int64:
						return val.Int(), nil
					case reflect.Bool:
						return val.Bool(), nil
					case reflect.Float64, reflect.Float32:
						return val.Float(), nil
					case reflect.String:
						return val.String(), nil
					}

					return val.Interface(), nil
				},
			}
		}
	}

	return &fields
}

func getGraphQLType(field reflect.StructField) *graphql.Scalar {
	switch field.Type.Name() {
	case reflect.Int.String(), reflect.Int32.String(), reflect.Int64.String():
		return graphql.Int
	case reflect.String.String():
		return graphql.String
	case reflect.Float32.String(), reflect.Float64.String():
		return graphql.Float
	case reflect.Bool.String():
		return graphql.Boolean
	default:
		return ScalarType
	}
}
