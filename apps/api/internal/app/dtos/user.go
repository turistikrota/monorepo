package dtos

import (
	"time"

	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/list"
)

type UserAdminList struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsActive  bool   `json:"is_active"`
	Createdat string `json:"created_at"`
}

func NewUserAdminList(res *list.PagiResponse[*entities.User]) *list.PagiResponse[*UserAdminList] {
	items := make([]*UserAdminList, 0, len(res.List))
	for _, item := range res.List {
		items = append(items, &UserAdminList{
			Id:        item.Id.String(),
			Name:      item.Name,
			Email:     item.Email,
			IsActive:  item.IsActive,
			Createdat: item.CreatedAt.Format(time.RFC3339),
		})
	}
	return &list.PagiResponse[*UserAdminList]{
		List:          items,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		Limit:         res.Limit,
		Page:          res.Page,
		TotalPage:     res.TotalPage,
	}
}
