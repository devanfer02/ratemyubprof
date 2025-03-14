package config

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
)

type sonicJsonSerializer struct{}

func NewSonicJSONSerializer() *sonicJsonSerializer {
	return &sonicJsonSerializer{}
}

func (s *sonicJsonSerializer) Serialize(c echo.Context, i interface{}, indent string) error {
	enc := sonic.ConfigDefault.NewEncoder(c.Response())
	if indent != "" {
		enc.SetIndent("", indent)
	}
	return enc.Encode(i)
}

func (s *sonicJsonSerializer) Deserialize(c echo.Context, i interface{}) error {
	err := sonic.ConfigDefault.NewDecoder(c.Request().Body).Decode(i)
	if ute, ok := err.(*json.UnmarshalTypeError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unmarshal type error: expected=%v, got=%v, field=%v, offset=%v", ute.Type, ute.Value, ute.Field, ute.Offset)).SetInternal(err)
	} else if se, ok := err.(*json.SyntaxError); ok {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Syntax error: offset=%v, error=%v", se.Offset, se.Error())).SetInternal(err)
	}
	return err
}