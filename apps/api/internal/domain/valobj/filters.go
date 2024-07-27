package valobj

type BaseFilters struct {
	Search   string `query:"search" validate:"omitempty,max=255"`
	IsActive string `query:"is_active" validate:"omitempty,oneof=0 1"`
}
