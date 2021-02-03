package validation

import (
	ov "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateEAN validates an EAN
func ValidateEAN(v interface{}) error {
	return ov.Validate(v,
		ov.Length(13, 13),
		is.Digit,
	)
}
