package graph

import (
	"github.com/sakhamoori/mytolet/api/internal/auth"
	"gorm.io/gorm"
)

// Resolver serves as a dependency injection container for your services
type Resolver struct{
	DB           *gorm.DB
	AuthProvider *auth.JWTProvider
}