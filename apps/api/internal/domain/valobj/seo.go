package valobj

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/9ssi7/slug"
)

type Seo struct {
	Title       string `json:"title" validate:"required,max=255" example:"My page title"`
	Description string `json:"description" validate:"required,max=255" example:"My page description"`
	Keywords    string `json:"keywords" validate:"required,max=255" example:"keyword1, keyword2, keyword3"`
	IsIndex     bool   `json:"is_index"`
	IsFollow    bool   `json:"is_follow"`
}

func (obj *Seo) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, obj)
	case string:
		if v != "" {
			return obj.Scan([]byte(v))
		}
	default:
		return errors.New("not supported")
	}

	return nil
}

func (obj Seo) Value() (driver.Value, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (obj Seo) MakeSlug() string {
	return slug.New(obj.Title, slug.TR)
}
