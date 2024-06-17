package tests

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    "myapp/handlers"
)

func TestCreateUser(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Jon","email":"jon@example.com"}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, handlers.CreateUser(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Equal(t, `{"name":"Jon","email":"jon@example.com"}`, rec.Body.String())
    }
}
