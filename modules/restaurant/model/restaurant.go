package model

type Restaurant struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name,omitempty" gorm:"column:name;"`
	Addr string `json:"address,omitempty" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name,omitempty" gorm:"column:name;"`
	Addr *string `json:"address,omitempty" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

type RestaurantCreate struct {
	Name string `json:"name,omitempty" gorm:"column:name;"`
	Addr string `json:"address,omitempty" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }
