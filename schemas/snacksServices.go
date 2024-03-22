package schemas

import (
	"gorm.io/gorm"
)

type SnackService struct {
	gorm.Model
	Name           string
	Company        string
	Location       string
	TimeOfDelivery int64
}
