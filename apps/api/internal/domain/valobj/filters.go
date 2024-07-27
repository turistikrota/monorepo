package valobj

type BaseFilters struct {
	Search string `query:"search" validate:"omitempty,max=255"`
}
