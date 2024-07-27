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

type placeFeatureRepo struct {
	syncRepo
	txnGormRepo
	db *gorm.DB
}

func NewPlaceFeatureRepo(db *gorm.DB) abstracts.PlaceFeatureRepo {
	return &placeFeatureRepo{
		db:          db,
		txnGormRepo: newTxnGormRepo(db),
	}
}

func (r *placeFeatureRepo) Save(ctx context.Context, feature *entities.PlaceFeature) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	if err := r.adapter.GetCurrent(ctx).Save(feature).Error; err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (r *placeFeatureRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.PlaceFeature, error) {
	var feature entities.PlaceFeature
	if err := r.adapter.GetCurrent(ctx).Model(&entities.PlaceFeature{}).Where("id = ?", id).First(&feature).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return &feature, nil
}

func (r *placeFeatureRepo) FindByIds(ctx context.Context, ids []uuid.UUID) ([]*entities.PlaceFeature, error) {
	var features []*entities.PlaceFeature
	if err := r.adapter.GetCurrent(ctx).Model(&entities.PlaceFeature{}).Where("id IN (?)", ids).Find(&features).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return features, nil
}

func (r *placeFeatureRepo) Filter(ctx context.Context, req *list.PagiRequest, filters *valobj.BaseFilters) (*list.PagiResponse[*entities.PlaceFeature], error) {
	conds := []query.Item{
		{
			Key:    "title ILIKE ?",
			Values: []interface{}{"%" + filters.Search + "%"},
			Skip:   filters.Search == "",
		},
		{
			Key:    "is_active = ?",
			Values: []interface{}{filters.IsActive == "1"},
			Skip:   filters.IsActive == "",
		},
	}
	res, err := query.RunList[*entities.PlaceFeature](r.adapter.GetCurrent(ctx), &entities.PlaceFeature{}, conds, req)
	if err != nil {
		return nil, rescode.Failed(err)
	}
	return res, nil
}
