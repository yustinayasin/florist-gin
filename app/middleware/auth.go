package middleware

import (
	"errors"
	"net/http"
	"time"

	"florist-gin/business/users"
	"florist-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId uint32 `json:"id"`
	jwt.MapClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

type JWTConfig struct {
	Claims                  *JwtCustomClaims
	SigningKey              []byte
	ErrorHandlerWithContext func(error, *gin.Context) error
}

type GeneratorToken interface {
	GenerateToken(userId uint32) string
}

func (jwtConf *ConfigJWT) Init() JWTConfig {
	return JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c *gin.Context) error {
			utils.ErrorResponseWithoutMessages(c, http.StatusForbidden, e.Error())
			return nil
		},
	}
}

func (configJWT ConfigJWT) GenerateToken(userId uint32) string {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Minute * time.Duration(configJWT.ExpiresDuration)).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(configJWT.SecretJWT))

	return token
}

func verifyToken(tokenString string, jwtConf ConfigJWT, c *gin.Context) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(jwtConf.SecretJWT), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

// Auth for private routes
func RequireAuth(next gin.HandlerFunc, jwtConf ConfigJWT, userRepoInterface users.UserRepoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You need to login first")
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// Verify token
		token, err := verifyToken(tokenString, jwtConf, c)
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}

		// Check the expiry date
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to extract claims")
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Token expired")
			c.Abort()
			return
		}

		// Find the user
		// Access the "subs" claim and convert it to int
		subsClaim, ok := claims["userId"].(uint32)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to parse user ID")
			c.Abort()
			return
		}

		// Convert subsClaim to int
		// subsInt := int(subsClaim)

		userUseCase := &users.UserUseCase{
			Repo: userRepoInterface,
			Jwt:  jwtConf,
		}

		// Ensure userUseCase is not nil before calling methods on it
		if userUseCase == nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to initialize user use case")
			c.Abort()
			return
		}

		user, err := userUseCase.GetUser(subsClaim)
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to fetch user data")
			c.Abort()
			return
		}

		// Attach the user to the request context
		c.Set("user", user)

		// Call the next handler
		next(c)
	}
}

func RequireAuthAdmin(next gin.HandlerFunc, jwtConf ConfigJWT, userRepoInterface users.UserRepoInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You need to login first")
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// Verify token
		token, err := verifyToken(tokenString, jwtConf, c)
		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Invalid Token")
			c.Abort()
			return
		}

		// Check the expiry date
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to extract claims")
			c.Abort()
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "Token expired")
			c.Abort()
			return
		}

		// Find the user
		// Access the "subs" claim and convert it to int
		subsClaim, ok := claims["userId"].(uint32)
		if !ok {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to parse user ID")
			c.Abort()
			return
		}

		// Convert subsClaim to int
		//subsInt := int(subsClaim)

		userUseCase := &users.UserUseCase{
			Repo: userRepoInterface,
			Jwt:  jwtConf,
		}

		// Ensure userUseCase is not nil before calling methods on it
		if userUseCase == nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to initialize user use case")
			c.Abort()
			return
		}

		user, err := userUseCase.GetUser(subsClaim)

		if err != nil {
			utils.ErrorResponseWithoutMessages(c, http.StatusInternalServerError, "Failed to fetch user data")
			c.Abort()
			return
		}

		if user.TypeId != 1 {
			utils.ErrorResponseWithoutMessages(c, http.StatusUnauthorized, "You don't have the authorization")
			c.Abort()
			return
		}

		// Attach the user to the request context
		c.Set("user", user)

		// Call the next handler
		next(c)
	}
}
