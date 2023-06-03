package auth

import (
	"net/http"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {

	// // Initialize database connection
	// db := config.ConnectDB()

	// // Initialize blog model
	// var user models.User

	// // Get the JSON body and decode into credentials
	// err := json.NewDecoder(r.Body).Decode(&user)
	// baseUtility.Catch(err)

	// // Validate the user credentials
	// if user.Username != "admin" || user.Password != "admin" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	fmt.Fprintf(w, "Wrong credentials")
	// 	return
	// }

	// // Generate token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"username": user.Username,
	// 	"password": user.Password,
	// 	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	// })

	// // Sign token with secret key
	// tokenString, err := token.SignedString([]byte("secret"))
	// baseUtility.Catch(err)

	// // Return token
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(tokenString))

	// // Insert user to database
	// db.Create(&user)

}

// Register function
func Register(w http.ResponseWriter, r *http.Request) {

	// // Initialize database connection
	// db := config.ConnectDB()

	// // Initialize blog model
	// var user models.User

	// // Get the JSON body and decode into credentials
	// err := json.NewDecoder(r.Body).Decode(&user)
	// baseUtility.Catch(err)

	// // Insert user to database
	// db.Create(&user)

	// // Return user
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(user)

}

// Logout function

func Logout(w http.ResponseWriter, r *http.Request) {

	// // Initialize database connection
	// db := config.ConnectDB()

	// // Initialize blog model
	// var user models.User

	// // Get the JSON body and decode into credentials
	// err := json.NewDecoder(r.Body).Decode(&user)
	// baseUtility.Catch(err)

	// // Delete user from database
	// db.Delete(&user)

	// // Return user
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(user)

}
