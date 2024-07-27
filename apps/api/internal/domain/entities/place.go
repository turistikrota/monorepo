package entities

import (
	"github.com/9ssi7/slug"
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/turistikrota/api/internal/domain/valobj"
)

type Place struct {
	Base
	valobj.Audit
	FeatureUUIDs valobj.UUIDArray                `json:"feature_uuids" gorm:"column:feature_uuids;type:jsonb;not null"`
	Title        string                          `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Description  string                          `json:"description" gorm:"column:description;type:text;not null"`
	Slug         string                          `json:"slug" gorm:"column:slug;type:varchar(255);not null"`
	Kind         valobj.PlaceKind                `json:"kind" gorm:"type:varchar(255);not null"`
	Seo          valobj.Seo                      `json:"seo" gorm:"column:seo;type:jsonb;not null"`
	Point        orb.Point                       `json:"point" gorm:"type:geography(POINT,4326);not null"`
	Images       valobj.JsonbArray[valobj.Image] `json:"images" gorm:"column:images;type:jsonb;not null"`
	MinTimeSpent int16                           `json:"min_time_spent" gorm:"column:min_time_spent;type:smallint;not null"`
	MaxTimeSpent int16                           `json:"max_time_spent" gorm:"column:max_time_spent;type:smallint;not null"`
	IsActive     bool                            `json:"is_active" gorm:"type:boolean;not null;default:true"`
	IsPayed      bool                            `json:"is_payed" gorm:"type:boolean;not null;default:false"`
}

func (p *Place) Enable() {
	p.IsActive = true
}

func (p *Place) Disable() {
	p.IsActive = false
}

func (p *Place) Update(adminId uuid.UUID, featureIds []uuid.UUID, kind valobj.PlaceKind, title string, description string, seo valobj.Seo, points []float64, images []*valobj.Image, minTimeSpent int16, maxTimeSpent int16, isPayed bool) {
	p.Slug = slug.New(seo.Title, slug.TR)
	p.FeatureUUIDs = valobj.UUIDArray(featureIds)
	p.Title = title
	p.Description = description
	p.Kind = kind
	p.Seo = seo
	p.Point = orb.Point{points[0], points[1]}
	p.Images = valobj.JsonbArray[valobj.Image](images)
	p.MinTimeSpent = minTimeSpent
	p.MaxTimeSpent = maxTimeSpent
	p.IsPayed = isPayed
	p.Audit.UpdatedBy = &adminId
}

func NewPlace(adminId uuid.UUID, featureIds []uuid.UUID, kind valobj.PlaceKind, title string, description string, seo valobj.Seo, points []float64, images []*valobj.Image, minTimeSpent int16, maxTimeSpent int16, isPayed bool) *Place {
	return &Place{
		Audit: valobj.Audit{
			MakedBy: &adminId,
		},
		FeatureUUIDs: valobj.UUIDArray(featureIds),
		Title:        title,
		Description:  description,
		Slug:         slug.New(seo.Title, slug.TR),
		Kind:         kind,
		Seo:          seo,
		Point:        orb.Point{points[0], points[1]},
		Images:       valobj.JsonbArray[valobj.Image](images),
		MinTimeSpent: minTimeSpent,
		MaxTimeSpent: maxTimeSpent,
		IsActive:     true,
		IsPayed:      isPayed,
	}
}
