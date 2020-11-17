package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// AuthInterface providers an interface to authenticate, generate tokens and validate tokens
type AuthInterface interface {
	GenerateToken(email string) string
	ValidateToken(token string) (*jwt.Token, error)
	LoginUser(email string, password string) bool
}

type authCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type authService struct {
	email     string
	password  string
	secretKey string
	issuer    string
}

// StaticAuthService TODO: Change to AuthService which provides a way to instanciate a new authservice
func StaticAuthService() AuthInterface {
	return &authService{
		email:     "admin@admin",
		password:  "admin",
		secretKey: getSecretKey(),
		issuer:    "boilerplate",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}

	return secret
}

func (a *authService) GenerateToken(email string) string {
	return ""
}

func (a *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return nil, errors.New("NEW")
}

func (a *authService) LoginUser(email string, password string) bool {
	return a.email == email && a.password == password
}
