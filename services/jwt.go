package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issuer:    "warezap-api",
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}

func (service *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			Issuer:    service.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	generate := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	token, err := generate.SignedString([]byte(service.secretKey))

	if err != nil {
		return "", err
	}

	return token, nil

}

// function to check validate token jwt
func (service *jwtService) ValidateToken(token_string string) bool {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(service.secretKey), nil
	})

	if err != nil {
		return false
	}

	// fmt.Printf("token is: %v", token.Valid)

	return token.Valid
}
