package gql

import (
    "github.com/graph-gophers/graphql-go"
)

func GetSchema() *graphql.Schema {
    schema := `
    type Query {
        hello: String!
    }`
    return graphql.MustParseSchema(schema, &Resolver{})
}
