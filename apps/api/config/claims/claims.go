package claims

type roleClaims struct {
	Super   string
	Create  string
	Update  string
	Enable  string
	Disable string
	View    string
	List    string
}

type placeClaims struct {
	Super   string
	Create  string
	Update  string
	Enable  string
	Disable string
	View    string
	List    string
}

type placeFeatureClaims struct {
	Super   string
	Create  string
	Update  string
	Enable  string
	Disable string
	View    string
	List    string
}

type userClaims struct {
	Super      string
	Disable    string
	Enable     string
	List       string
	View       string
	RoleAdd    string
	RoleRemove string
}

var (
	Admin      string = "admin"
	AdminSuper string = "admin_super"

	Role = roleClaims{
		Super:   "role_super",
		Create:  "role_create",
		Update:  "role_update",
		Enable:  "role_enable",
		Disable: "role_disable",
		View:    "role_view",
		List:    "role_list",
	}

	Place = placeClaims{
		Super:   "place_super",
		Create:  "place_create",
		Update:  "place_update",
		Enable:  "place_enable",
		Disable: "place_disable",
		View:    "place_view",
		List:    "place_list",
	}

	PlaceFeature = placeFeatureClaims{
		Super:   "place_feature_super",
		Create:  "place_feature_create",
		Update:  "place_feature_update",
		Enable:  "place_feature_enable",
		Disable: "place_feature_disable",
		View:    "place_feature_view",
		List:    "place_feature_list",
	}

	User = userClaims{
		Super:      "user_super",
		Disable:    "user_disable",
		Enable:     "user_enable",
		List:       "user_list",
		View:       "user_view",
		RoleAdd:    "user_role_add",
		RoleRemove: "user_role_remove",
	}
)

func IsReal(claim string) bool {
	switch claim {
	case Admin, AdminSuper:
		return true
	case Role.Super, Role.Create, Role.Update, Role.Enable, Role.Disable, Role.View, Role.List:
		return true
	case Place.Super, Place.Create, Place.Update, Place.Enable, Place.Disable, Place.View, Place.List:
		return true
	case PlaceFeature.Super, PlaceFeature.Create, PlaceFeature.Update, PlaceFeature.Enable, PlaceFeature.Disable, PlaceFeature.View, PlaceFeature.List:
		return true
	}
	return false
}
