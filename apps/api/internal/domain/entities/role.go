package entities

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/turistikrota/api/internal/domain/valobj"
)

type Role struct {
	Base
	valobj.Audit
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text;default:null"`
	IsActive    bool           `json:"is_active" gorm:"type:boolean;not null;default:true"`
	IsLocked    bool           `json:"is_locked" gorm:"type:boolean;not null;default:false"`
	Claims      pq.StringArray `json:"claims" gorm:"type:text[]"`
}

func (r *Role) AddClaim(claim string) {
	if r.CheckClaim(claim) {
		return
	}
	r.Claims = append(r.Claims, claim)
}

func (r *Role) CheckClaim(claim string) bool {
	for _, c := range r.Claims {
		if c == claim {
			return true
		}
	}
	return false
}

func (r *Role) RemoveClaim(claim string) {
	for i, c := range r.Claims {
		if c == claim {
			r.Claims = append(r.Claims[:i], r.Claims[i+1:]...)
			break
		}
	}
}

func (r *Role) Enable(userId uuid.UUID) {
	r.IsActive = true
	r.Audit.UpdatedBy = &userId
}

func (r *Role) Disable(userId uuid.UUID) {
	r.IsActive = false
	r.Audit.UpdatedBy = &userId
}

func (r *Role) Update(userId uuid.UUID, name string, description string, claims []string) {
	r.Name = name
	r.Description = description
	r.Claims = pq.StringArray(claims)
	r.Audit.UpdatedBy = &userId
}

func NewRole(userId uuid.UUID, name string, description string, claims []string) *Role {
	return &Role{
		Audit: valobj.Audit{
			MakedBy: &userId,
		},
		Name:        name,
		Description: description,
		Claims:      pq.StringArray(claims),
		IsActive:    true,
		IsLocked:    false,
	}
}
