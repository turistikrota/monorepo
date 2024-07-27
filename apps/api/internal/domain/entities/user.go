package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/ptr"
)

type User struct {
	Base
	valobj.Audit
	Name       string         `json:"name" gorm:"type:varchar(255);not null"`
	Email      string         `json:"email" gorm:"type:varchar(255);not null;unique"`
	IsActive   bool           `json:"is_active" gorm:"type:boolean;not null;default:true"`
	Roles      pq.StringArray `json:"roles" gorm:"type:text[]"`
	TempToken  *string        `json:"temp_token" gorm:"type:varchar(255);default:null;index:idx_verifier"`
	VerifiedAt *time.Time     `json:"verified_at" gorm:"type:timestamp;default:null;index:idx_verifier"`
}

func (u *User) AddRole(userId uuid.UUID, role string) {
	u.Roles = append(u.Roles, role)
	u.Audit.UpdatedBy = &userId
}

func (u *User) CheckRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) RemoveRole(userId uuid.UUID, role string) {
	for i, r := range u.Roles {
		if r == role {
			u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			break
		}
	}
	u.Audit.UpdatedBy = &userId
}

func (u *User) Verify() {
	u.VerifiedAt = ptr.Time(time.Now())
	u.IsActive = true
	u.TempToken = nil
}

func (u *User) UnVerify() {
	u.VerifiedAt = nil
	u.IsActive = false
	u.TempToken = ptr.String(uuid.New().String())
}

func (u *User) Enable(userId uuid.UUID) {
	u.IsActive = true
	u.Audit.UpdatedBy = &userId
}

func (u *User) Disable(userId uuid.UUID) {
	u.IsActive = false
	u.Audit.UpdatedBy = &userId
}

func NewUser(name string, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Roles:     pq.StringArray{},
		TempToken: ptr.String(uuid.New().String()),
	}
}

func NewUserFromAdmin(name string, email string, adminId uuid.UUID) *User {
	return &User{
		Audit: valobj.Audit{
			MakedBy: &adminId,
		},
		Name:      name,
		Email:     email,
		Roles:     pq.StringArray{},
		TempToken: ptr.String(uuid.New().String()),
	}
}
