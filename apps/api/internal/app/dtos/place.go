package dtos

import (
	"time"

	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/list"
)

type PlaceList struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	FeatureIds   []string         `json:"feature_ids"`
	Slug         string           `json:"slug"`
	Kind         valobj.PlaceKind `json:"kind"`
	Point        []float64        `json:"point"`
	Images       []*valobj.Image  `json:"images"`
	MinTimeSpent int16            `json:"min_time_spent"`
	MaxTimeSpent int16            `json:"max_time_spent"`
	IsPayed      bool             `json:"is_payed"`
}

type PlaceView struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	FeatureIds   []string         `json:"feature_ids"`
	Slug         string           `json:"slug"`
	Seo          valobj.Seo       `json:"seo"`
	Kind         valobj.PlaceKind `json:"kind"`
	Point        []float64        `json:"point"`
	Images       []*valobj.Image  `json:"images"`
	MinTimeSpent int16            `json:"min_time_spent"`
	MaxTimeSpent int16            `json:"max_time_spent"`
	IsPayed      bool             `json:"is_payed"`
}

type PlaceAdminList struct {
	Id           string           `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	FeatureIds   []string         `json:"feature_ids"`
	Slug         string           `json:"slug"`
	Kind         valobj.PlaceKind `json:"kind"`
	Point        []float64        `json:"point"`
	Images       []*valobj.Image  `json:"images"`
	MinTimeSpent int16            `json:"min_time_spent"`
	MaxTimeSpent int16            `json:"max_time_spent"`
	IsPayed      bool             `json:"is_payed"`
	IsActive     bool             `json:"is_active"`
	CreatedAt    string           `json:"created_at"`
}

func NewPlaceList(res *list.PagiResponse[*entities.Place]) *list.PagiResponse[*PlaceList] {
	items := make([]*PlaceList, 0, len(res.List))
	for _, item := range res.List {
		items = append(items, &PlaceList{
			Title:        item.Title,
			Description:  item.Description,
			FeatureIds:   item.FeatureIds.ToStringArray(),
			Slug:         item.Slug,
			Kind:         item.Kind,
			Point:        []float64{item.Point.Lat(), item.Point.Lon()},
			Images:       item.Images,
			MinTimeSpent: item.MinTimeSpent,
			MaxTimeSpent: item.MaxTimeSpent,
			IsPayed:      item.IsPayed,
		})
	}
	return &list.PagiResponse[*PlaceList]{
		List:          items,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		Limit:         res.Limit,
		Page:          res.Page,
		TotalPage:     res.TotalPage,
	}
}

func NewPlaceView(item *entities.Place) *PlaceView {
	return &PlaceView{
		Title:        item.Title,
		Description:  item.Description,
		FeatureIds:   item.FeatureIds.ToStringArray(),
		Slug:         item.Slug,
		Seo:          item.Seo,
		Kind:         item.Kind,
		Point:        []float64{item.Point.Lat(), item.Point.Lon()},
		Images:       item.Images,
		MinTimeSpent: item.MinTimeSpent,
		MaxTimeSpent: item.MaxTimeSpent,
		IsPayed:      item.IsPayed,
	}
}

func NewPlaceAdminList(res *list.PagiResponse[*entities.Place]) *list.PagiResponse[*PlaceAdminList] {
	items := make([]*PlaceAdminList, 0, len(res.List))
	for _, item := range res.List {
		items = append(items, &PlaceAdminList{
			Id:           item.Id.String(),
			Title:        item.Title,
			Description:  item.Description,
			FeatureIds:   item.FeatureIds.ToStringArray(),
			Slug:         item.Slug,
			Kind:         item.Kind,
			Point:        []float64{item.Point.Lat(), item.Point.Lon()},
			Images:       item.Images,
			MinTimeSpent: item.MinTimeSpent,
			MaxTimeSpent: item.MaxTimeSpent,
			IsPayed:      item.IsPayed,
			IsActive:     item.IsActive,
			CreatedAt:    item.CreatedAt.Format(time.RFC3339),
		})
	}
	return &list.PagiResponse[*PlaceAdminList]{
		List:          items,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		Limit:         res.Limit,
		Page:          res.Page,
		TotalPage:     res.TotalPage,
	}
}
