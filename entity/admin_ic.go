package entity

type AdminIc struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
	Password string `json:"password"`
}
