package query

import (
	"github.com/turistikrota/api/pkg/list"
	"gorm.io/gorm"
)

func RunList[T any](db *gorm.DB, model interface{}, filters []Item, req *list.PagiRequest) (*list.PagiResponse[T], error) {
	var items []T
	q := db.Model(model)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	filterq, values := Build(filters)
	var filteredTotal int64
	if filterq != "" {
		q = q.Where(filterq, values...)
		if err := q.Count(&filteredTotal).Error; err != nil {
			return nil, err
		}
	} else {
		filteredTotal = total
	}
	if err := q.Limit(*req.Limit).Offset(req.Offset()).Find(&items).Error; err != nil {
		return nil, err
	}
	return list.NewPagiResponse(req, items, total, filteredTotal), nil
}
