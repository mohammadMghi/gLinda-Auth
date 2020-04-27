package model

type User struct {
	ID        int64      `gorm:"primary_key" json:"id"`
	//Name      string    `json:"name"`
	//Age       string    `json:"age"`
	Email     string    `gorm="email" json:"email"`
	Password  string    `gorm="password" json:"password"`
	Token     string    `gorm:"-", json:"token,omitempty"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt time.Time `json:"deleted_at"`
}

func (User) TableName() string { 
	return "users" 
}
const (
	// Key (Should come from somewhere else).
	Key = "secret"
)