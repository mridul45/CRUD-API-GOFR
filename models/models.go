package models

type Customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Spending string `json:"spending"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrderItem struct {
	ID       int `json:"id"`
	DishID   int `json:"dish_id"`
	Quantity int `json:"quantity"`
	OrderID  int `json:"order_id" gorm:"index"` // Add this line for the foreign key
}

type Dish struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Price      string   `json:"price"`
	CategoryID int      `json:"category_id" gorm:"index"`
	Category   Category `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Bill struct {
	ID         int      `json:"id"`
	CustomerID int      `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignkey:CustomerID"`
	Amount     float64  `json:"amount"`
}
