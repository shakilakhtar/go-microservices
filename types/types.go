package types


import (
"time"
)

// Data Binding structs
type (
	User struct {
		Id           string `json:"id"`
		FirstName    string `json:"firstname"`
		LastName     string `json:"lastname"`
		Email        string `json:"email"`
		Password     string `json:"password,omitempty"`
		HashPassword []byte `json:"hashpassword,omitempty"`
	}

	Asset struct {
		Id          int       `sql:"AUTO_INCREMENT" gorm:"primary_key" json:"id"`
		Name        string    `sql:"not null;unique" gorm:"column:name" json:"name"`
		Description string    `gorm:"column:description" json:"description"`
		Status      string    `sql:"DEFAULT:'Active'" gorm:"column:status" json:"status,omitempty"`
		CreatedBy   string    `sql:"not null" gorm:"column:createdby" json:"createdby"`
		CreatedOn   time.Time `gorm:"column:createdon" json:"createdon,omitempty"`
	}
)
