package domain

type User struct {
	ID       string `json:"id" bson:"_id"`            // Unique identifier
	Username string `json:"username" bson:"username"` // Login username
	Password string `json:"password" bson:"password"` // Hashed password
	Role     string `json:"role" bson:"role"`         // User role (e.g., admin, inventory, warehouse)
}
