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

type placeRepo struct {
	syncRepo
	txnGormRepo
	db *gorm.DB
}

func NewPlaceRepo(db *gorm.DB) abstracts.PlaceRepo {
	return &placeRepo{
		db:          db,
		txnGormRepo: newTxnGormRepo(db),
	}
}

func (r *placeRepo) Save(ctx context.Context, place *entities.Place) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	if err := r.adapter.GetCurrent(ctx).Save(place).Error; err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (r *placeRepo) IsExistsBySlug(ctx context.Context, slug string) (bool, error) {
	var count int64
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Place{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
		return false, rescode.Failed(err)
	}
	return count > 0, nil
}

func (r *placeRepo) FindBySlug(ctx context.Context, slug string) (*entities.Place, error) {
	var place entities.Place
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Place{}).Where("slug = ? AND is_active = ?", slug, true).First(&place).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rescode.NotFound(err)
		}
		return nil, rescode.Failed(err)
	}
	return &place, nil
}

func (r *placeRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.Place, error) {
	var place entities.Place
	if err := r.adapter.GetCurrent(ctx).Model(&entities.Place{}).Where("id = ?", id).First(&place).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rescode.NotFound(err)
		}
		return nil, rescode.Failed(err)
	}
	return &place, nil
}

func (r *placeRepo) Filter(ctx context.Context, req *list.PagiRequest, filters *valobj.PlaceFilters) (*list.PagiResponse[*entities.Place], error) {
	conds := []query.Item{
		{
			Key:    "title ILIKE ?",
			Values: query.V{"%" + filters.Search + "%"},
			Skip:   filters.Search == "",
		},
		{
			Key:    "is_payed = ?",
			Values: query.V{filters.IsPayed == "1"},
			Skip:   filters.IsPayed == "",
		},
		{
			Key:    "is_active = ?",
			Values: query.V{filters.IsActive == "1"},
			Skip:   filters.IsActive == "",
		},
		{
			Key:    "kind = ?",
			Values: query.V{filters.Kind},
			Skip:   filters.Kind == "",
		},
		{
			Key:    "WHERE ST_DWithin(ST_MakePoint(longitude,latitude)::geography,ST_MakePoint(?, ?)::geography,?)",
			Values: query.V{filters.Lng, filters.Lat, filters.Distance},
			Skip:   filters.Lat == "" || filters.Lng == "" || filters.Distance == "",
		},
		{
			Key:    "min_time_spent >= ?",
			Values: query.V{filters.MinTimeSpent},
			Skip:   filters.MinTimeSpent == "",
		},
		{
			Key:    "max_time_spent <= ?",
			Values: query.V{filters.MaxTimeSpent},
			Skip:   filters.MaxTimeSpent == "",
		},
	}
	res, err := query.RunList[*entities.Place](r.adapter.GetCurrent(ctx), &entities.Place{}, conds, req)
	if err != nil {
		return nil, rescode.Failed(err)
	}
	return res, nil
}
