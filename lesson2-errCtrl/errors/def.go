package def

import "github.com/pkg/errors"

const (
	DataNotFound = iota + 10000
)

var DataNotFoundErr = errors.New("query db empty")
