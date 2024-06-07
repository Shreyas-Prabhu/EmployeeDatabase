package models

type Employee struct {
	ID       int     `gorm:"primaryKey" json:"id,omitempty" validate:"required"`
	Name     string  `json:"name,omitempty" validate:"required"`
	Position string  `json:"position,omitempty" validate:"required"`
	Salary   float64 `json:"salary,omitempty" validate:"required"`
}
