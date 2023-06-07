package auth

import (
	"go-learn-restapi-mysql/config"
	"go-learn-restapi-mysql/controllers/base"
	"go-learn-restapi-mysql/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

// Login function
func Login(c *gin.Context) {
	// Initialize blog model
	var user_params models.Users

	// Get the JSON body and decode into credentials
	err := c.ShouldBindJSON(&user_params)
	base.ResponseBindJson(err, c)

	var user models.Users
	if err := config.DB.Where("email = ?", user_params.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			base.Response(c, false, http.StatusUnauthorized, "Wrong credentials")
		} else {
			base.Response(c, false, http.StatusInternalServerError, "Error while querying the user")
		}
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user_params.Password)); err != nil {
		base.Response(c, false, http.StatusUnauthorized, "Wrong credentials")
		return
	}

	expTime := time.Now().Add(time.Minute * time.Duration(config.JWT_EXPIRE))
	claims := &config.JWTClaim{
		Email:    user.Email,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "learn-restapi-mysql",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		base.Response(c, false, http.StatusInternalServerError, "Error while signing the token")
		return
	}

	// Set cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenString,
		HttpOnly: true,
	})

	// Return token
	data := gin.H{
		"name":          user.Name,
		"token":         tokenString,
		"token_expired": expTime,
	}

	base.ResponseWithData(c, true, http.StatusOK, "Login successfully", data)

}

// Register function
func Register(c *gin.Context) {
	// Initialize user model
	var user models.Users

	// Get the JSON body and decode into user
	err := c.ShouldBindJSON(&user)
	base.ResponseBindJson(err, c)

	// Validate user input
	if err := validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		base.ResponseValidate(errors, c)
		return
	}

	// Insert user to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		base.Response(c, false, http.StatusInternalServerError, "Failed to hash password")

		return
	}
	user.Password = string(hashedPassword)
	err = config.DB.Create(&user).Error

	base.ResponseCreate(err, user, c)
}

// Logout function
func Logout(c *gin.Context) {
	// Clear cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	// Return token
	base.Response(c, true, http.StatusOK, "Logout successfully")
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				base.Response(c, false, http.StatusUnauthorized, "Unauthorized")
				return
			}
		}

		// Mengambil token value
		tokenString := cookie

		// parsing token jwt
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_KEY), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					// token invalid
					base.Response(c, false, http.StatusUnauthorized, "Unauthorized")
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// token expired
					base.Response(c, false, http.StatusUnauthorized, "Unauthorized, Token expired!")
					return
				}
			}

			base.Response(c, false, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if !token.Valid {
			base.Response(c, false, http.StatusUnauthorized, "Unauthorized")
			return
		}

		c.Next()
	}
}
