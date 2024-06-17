package handlers

import (
    "github.com/labstack/echo/v4"
    "github.com/graph-gophers/graphql-go/relay"
    "myapp/gql"
)

func GraphQL(c echo.Context) error {
    schema := gql.GetSchema()
    h := relay.Handler{Schema: schema}
    h.ServeHTTP(c.Response(), c.Request())
    return nil
}
