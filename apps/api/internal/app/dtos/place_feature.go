package dtos

import (
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/list"
)

type PlaceFeatureList struct {
	Id          string `json:"id"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewPlaceFeatureList(res *list.PagiResponse[*entities.PlaceFeature]) *list.PagiResponse[*PlaceFeatureList] {
	items := make([]*PlaceFeatureList, 0, len(res.List))
	for _, item := range res.List {
		items = append(items, &PlaceFeatureList{
			Id:          item.Id.String(),
			Icon:        item.Icon,
			Title:       item.Title,
			Description: item.Description,
		})
	}
	return &list.PagiResponse[*PlaceFeatureList]{
		List:          items,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		Limit:         res.Limit,
		Page:          res.Page,
		TotalPage:     res.TotalPage,
	}
}
