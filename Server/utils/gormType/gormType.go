package gormType

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
    if s == nil {
        return "[]", nil
    }
    return json.Marshal(s)
}

func (s *StringSlice) Scan(value interface{}) error {
    if value == nil {
        *s = nil
        return nil
    }

    switch v := value.(type) {
    case []byte:
        return json.Unmarshal(v, s)
    case string:
        return json.Unmarshal([]byte(v), s)
    default:
        return errors.New("unsupported data type")
    }
}
