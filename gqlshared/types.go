package gqlshared

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// ScalarType - Generic GraphQL Scalar type
var ScalarType = graphql.NewScalar(graphql.ScalarConfig{
	Name: "Scalar",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		return valueAST
	},
	ParseValue: func(value interface{}) interface{} {
		return value
	},
})
