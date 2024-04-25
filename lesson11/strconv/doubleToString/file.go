package doubletostring

import (
	"lesson11/math/mul"
	"strconv"
)

func DtoS(a int) string {
	return strconv.Itoa(mul.Double(a))
}
