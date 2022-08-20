package api

import "errors"

var ErrRequiredFieldsEmpty = errors.New("required fields is empty")
var ErrBadRequiredFields = errors.New("bad required fields")
