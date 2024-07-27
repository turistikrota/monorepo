package validation

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/turistikrota/api/internal/domain/valobj"
	"github.com/turistikrota/api/pkg/currency"
	"github.com/turistikrota/api/pkg/iban"
)

func validateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		if valuer.String() == uuid.Nil.String() {
			return nil
		}
		return valuer.String()
	}
	return nil
}

func validatePlaceKind(fl validator.FieldLevel) bool {
	t := fl.Field().String()
	return t == valobj.PlaceKindEating.String() || t == valobj.PlaceKindCoffee.String() || t == valobj.PlaceKindBar.String() || t == valobj.PlaceKindBeach.String() || t == valobj.PlaceKindAmaze.String() || t == valobj.PlaceKindShopping.String() || t == valobj.PlaceKindTransport.String() || t == valobj.PlaceKindCulture.String() || t == valobj.PlaceKindNature.String() || t == valobj.PlaceKindHealth.String() || t == valobj.PlaceKindSport.String() || t == valobj.PlaceKindNightlife.String() || t == valobj.PlaceKindGarden.String() || t == valobj.PlaceKindTemple.String() || t == valobj.PlaceKindMuseum.String() || t == valobj.PlaceKindAntique.String() || t == valobj.PlaceKindPark.String() || t == valobj.PlaceKindThemePark.String() || t == valobj.PlaceKindOther.String()
}

func validateIban(fl validator.FieldLevel) bool {
	return iban.Validate(fl.Field().String())
}

func validateAmount(fl validator.FieldLevel) bool {
	d, err := decimal.NewFromString(fl.Field().String())
	if err != nil {
		return false
	}
	return d.GreaterThanOrEqual(decimal.Zero)
}

func validateCurrency(fl validator.FieldLevel) bool {
	return currency.IsValid(fl.Field().String())
}

func validateUserName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(userNameRegexp, fl.Field().String())
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(passwordRegexp, fl.Field().String())
	return matched
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(slugRegexp, fl.Field().String())
	return matched
}

func validateLocale(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(localeRegexp, fl.Field().String())
	return matched
}

func validateGender(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(genderRegexp, fl.Field().String())
	return matched
}

func validatePhone(f1 validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(phoneWithCountryCodeRegexp, f1.Field().String())
	return matched
}
