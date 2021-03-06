package utils

import (
	"github.com/satori/go.uuid"

)

func NewUuid()string{
	u := uuid.NewV4().String()
	return u
}