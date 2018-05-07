package expr

import (
	"github.com/skillian/errors"
	"github.com/skillian/logging"
)

var (
	errNilPtr error = errors.ErrorMessage{Fmt: "nil pointer"}

	logger = logging.GetLogger("github.com/skillian/expr")
)

// TypeError represents an error from an expression with the wrong type.
type TypeError errors.Error

// TypeErrorFromExpectedAndActual creates a TypeError from a variable with an
// expected type and the actual value.  The returned error's call stack
// excludes TypeErrorFromExpectedAndActual itself.
func TypeErrorFromExpectedAndActual(expected, actual interface{}) TypeError {
	return TypeError(errors.CreateError(
		errors.ErrorMessage{
			Fmt: "expected value of type %T but got %v (type: %T)",
			Args: []interface{}{
				expected,
				actual,
				actual,
			},
		},
		nil,
		nil,
		1))
}

func (te TypeError) Error() string {
	return errors.Error(te).Error()
}
