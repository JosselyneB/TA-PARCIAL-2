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

func TestLogin(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("username=user&password=password"))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, handlers.Login(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "token")
    }
}
