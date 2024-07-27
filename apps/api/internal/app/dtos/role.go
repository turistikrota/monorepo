package dtos

import (
	"time"

	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/list"
)

type RoleList struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsLocked    bool   `json:"is_locked"`
	IsActive    bool   `json:"is_active"`
	Createdat   string `json:"created_at"`
}

func NewRoleList(res *list.PagiResponse[*entities.Role]) *list.PagiResponse[*RoleList] {
	items := make([]*RoleList, 0, len(res.List))
	for _, item := range res.List {
		items = append(items, &RoleList{
			Id:          item.Id.String(),
			Name:        item.Name,
			Description: item.Description,
			IsLocked:    item.IsLocked,
			IsActive:    item.IsActive,
			Createdat:   item.CreatedAt.Format(time.RFC3339),
		})
	}
	return &list.PagiResponse[*RoleList]{
		List:          items,
		Total:         res.Total,
		FilteredTotal: res.FilteredTotal,
		Limit:         res.Limit,
		Page:          res.Page,
		TotalPage:     res.TotalPage,
	}
}
