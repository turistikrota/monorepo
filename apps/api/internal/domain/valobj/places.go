package valobj

// PlaceKind defines the type of place
//
// Enum: eating, coffee, bar, beach, amaze, shopping, transport, culture, nature, health, sport, nightlife, garden, temple, museum, antique, park, themePark, other
type PlaceKind string

const (
	PlaceKindEating    PlaceKind = "eating"
	PlaceKindCoffee    PlaceKind = "coffee"
	PlaceKindBar       PlaceKind = "bar"
	PlaceKindBeach     PlaceKind = "beach"
	PlaceKindAmaze     PlaceKind = "amaze"
	PlaceKindShopping  PlaceKind = "shopping"
	PlaceKindTransport PlaceKind = "transport"
	PlaceKindCulture   PlaceKind = "culture"
	PlaceKindNature    PlaceKind = "nature"
	PlaceKindHealth    PlaceKind = "health"
	PlaceKindSport     PlaceKind = "sport"
	PlaceKindNightlife PlaceKind = "nightlife"
	PlaceKindGarden    PlaceKind = "garden"
	PlaceKindTemple    PlaceKind = "temple"
	PlaceKindMuseum    PlaceKind = "museum"
	PlaceKindAntique   PlaceKind = "antique"
	PlaceKindPark      PlaceKind = "park"
	PlaceKindThemePark PlaceKind = "themePark"
	PlaceKindOther     PlaceKind = "other"
)

type PlaceFilters struct {
	Search       string `query:"search" validate:"omitempty,max=255"`
	Lat          string `query:"lat" validate:"omitempty,latitude"`
	Lng          string `query:"lng" validate:"omitempty,longitude"`
	Distance     string `query:"distance" validate:"omitempty,numeric"`
	Kind         string `query:"kind" validate:"omitempty,oneof=eating coffee bar beach amaze shopping transport culture nature health sport nightlife garden temple museum antique park themePark other"`
	MinTimeSpent string `query:"min_time_spent" validate:"omitempty,numeric"`
	MaxTimeSpent string `query:"max_time_spent" validate:"omitempty,numeric"`
	IsPayed      string `query:"is_payed" validate:"omitempty,oneof=0 1"`
	IsActive     string `query:"is_active" validate:"omitempty,oneof=0 1"`
}
