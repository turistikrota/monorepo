package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/list"
	"github.com/turistikrota/api/pkg/query"
	"github.com/turistikrota/api/pkg/rescode"
	"gorm.io/gorm"
)

type roleRepo struct {
	syncRepo
	txnGormRepo
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) abstracts.RoleRepo {
	return &roleRepo{
		db:          db,
		txnGormRepo: newTxnGormRepo(db),
	}
}

func (r *roleRepo) Save(ctx context.Context, role *entities.Role) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	if err := r.adapter.GetCurrent(ctx).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	var role entities.Role
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Role{}).Where("id = ?", id).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rescode.NotFound(err)
		}
		return nil, err
	}
	return &role, nil
}

func (r *roleRepo) FindByIds(ctx context.Context, ids []uuid.UUID) ([]*entities.Role, error) {
	var roles []*entities.Role
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Role{}).Where("id IN (?) AND is_active = ?", ids, true).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepo) Filter(ctx context.Context, req *list.PagiRequest, filters *valobj.BaseFilters) (*list.PagiResponse[*entities.Role], error) {
	conds := []query.Item{
		{
			Key:    "title ILIKE ?",
			Values: query.V{"%" + filters.Search + "%"},
			Skip:   filters.Search == "",
		},
		{
			Key:    "is_active = ?",
			Values: query.V{filters.IsActive == "1"},
			Skip:   filters.IsActive == "",
		},
	}
	res, err := query.RunList[*entities.Role](r.adapter.GetCurrent(ctx), &entities.Role{}, conds, req)
	if err != nil {
		return nil, rescode.Failed(err)
	}
	return res, nil
}
