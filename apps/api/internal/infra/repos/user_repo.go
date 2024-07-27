package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/turistikrota/api/internal/domain/abstracts"
	"github.com/turistikrota/api/internal/domain/entities"
	"github.com/turistikrota/api/pkg/list"
	"github.com/turistikrota/api/pkg/rescode"
	"gorm.io/gorm"
)

type userRepo struct {
	syncRepo
	txnGormRepo
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) abstracts.UserRepo {
	return &userRepo{
		db:          db,
		txnGormRepo: newTxnGormRepo(db),
	}
}

func (r *userRepo) Save(ctx context.Context, user *entities.User) error {
	r.syncRepo.Lock()
	defer r.syncRepo.Unlock()
	if err := r.adapter.GetCurrent(ctx).Save(user).Error; err != nil {
		return rescode.Failed(err)
	}
	return nil
}

func (r *userRepo) FindByToken(ctx context.Context, token string) (*entities.User, error) {
	var user entities.User
	if err := r.adapter.GetCurrent(ctx).Model(&entities.User{}).Where("temp_token = ? AND verified_at IS NULL", token).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, rescode.NotFound(err)
		}
		return nil, rescode.Failed(err)
	}
	return &user, nil
}

func (r *userRepo) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.adapter.GetCurrent(ctx).Model(&entities.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return &user, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.adapter.GetCurrent(ctx).Model(&entities.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return &user, nil
}

func (r *userRepo) IsExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.adapter.GetCurrent(ctx).Model(&entities.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, rescode.Failed(err)
	}
	return count > 0, nil
}

func (r *userRepo) FindByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var user entities.User
	if err := r.adapter.GetCurrent(ctx).Model(&entities.User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return &user, nil
}

func (r *userRepo) Filter(ctx context.Context, req *list.PagiRequest, search string, isActive string) (*list.PagiResponse[*entities.User], error) {
	var users []*entities.User
	query := r.adapter.GetCurrent(ctx).Model(&entities.User{})
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if isActive != "" {
		query = query.Where("is_active = ?", isActive)
	}
	var filteredTotal int64
	if err := query.Count(&filteredTotal).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	if err := query.Limit(*req.Limit).Offset(req.Offset()).Find(&users).Error; err != nil {
		return nil, rescode.Failed(err)
	}
	return list.NewPagiResponse(req, users, total, filteredTotal), nil
}
