package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/meech/auth_user_api/pkg/config"
	"github.com/meech/auth_user_api/pkg/database"
	"github.com/meech/auth_user_api/pkg/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	database.DB.Create(&user)

	response := models.Response{Message: "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateJWT(userID uint) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate the token with the secret key
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	response := models.Response{Message: "Login successful", Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get user ID from the token
	userID, err := getUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Fetch the existing user from the database
	var existingUser models.User
	result := database.DB.First(&existingUser, userID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update user profile fields
	existingUser.Email = updatedUser.Email
	existingUser.Name = updatedUser.Name
	database.DB.Save(&existingUser)

	response := models.Response{Message: "Profile updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Get user ID from the token
	userID, err := getUserIDFromToken(r)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Delete user account from the database
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	database.DB.Delete(&user)

	response := models.Response{Message: "Account deleted successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getUserIDFromToken(r *http.Request) (uint, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userID := uint(claims["user_id"].(float64))
	return userID, nil
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
