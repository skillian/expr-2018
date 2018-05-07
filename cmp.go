package expr

import "github.com/skillian/errors"

// Cmper implementers can compare themselves with other values.
type Cmper interface {
	Value
	Cmp(v Value) (int, error)
}

// Cmp compares the left expression to the right according the the following
// rules:
//
//	left == right	->	result == 0
//	left != right	->	result != 0
//	left > right	->	result > 0
//	left >= right	->	result >= 0
//	left < right	-> result < 0
//	left <= right	-> result <= 0
//
// An error is returned if the two expressions cannot be compared.
func Cmp(left, right Expr) (result int, err error) {
	var comparands [2]Value
	err = SimplifyIntoValues([]Expr{left, right}, comparands[:])
	if err != nil {
		return 0, err
	}
	if cmper, ok := comparands[0].(Cmper); ok {
		result, err = cmper.Cmp(comparands[1])
		if err == nil {
			return result, nil
		}
	}
	if cmper, ok := comparands[1].(Cmper); ok {
		inverted, err2 := cmper.Cmp(comparands[0])
		if err2 == nil {
			if err != nil {
				logger.Warn(
					"comparing %v to %v succeeded but comparing %v to %v failed: %v",
					comparands[1], comparands[0],
					comparands[0], comparands[1], err)
			}
			return -inverted, nil
		}
		// preserve err to see why we failed over to comparing right to left:
		return 0, errors.CreateError(err2, nil, err, 0)
	}
	if left == right {
		return 0, nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", left, right)
}
