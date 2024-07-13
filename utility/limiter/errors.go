package limiter

import "github.com/xmplusdev/xmcore/common/errors"

func newError(values ...interface{}) *errors.Error {
	return errors.New(values...)
}
