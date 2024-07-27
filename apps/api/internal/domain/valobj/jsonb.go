package valobj

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonbArray[T any] []*T

func (j *JsonbArray[T]) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, &j)
	case string:
		if v != "" {
			return j.Scan([]byte(v))
		}
	default:
		return fmt.Errorf("invalid type: %T", v)
	}
	return nil
}

func (j JsonbArray[T]) Value() (driver.Value, error) {
	data, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}
