package auth

import (
	"errors"
	"gin-simple-api/databases"
	"gin-simple-api/models"
	"gin-simple-api/src/types"
	"math"
	"os"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(perPage, page int, search string) ([]models.User, types.PaginationMeta, error) {
	var users []models.User
	var totalData int64
	var dbQuery = databases.DB.Model(&models.User{})

	if search != "" {
		dbQuery = dbQuery.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Menghitung total data yang cocok dengan query
	if err := dbQuery.Count(&totalData).Error; err != nil {
		return nil, types.PaginationMeta{}, err
	}

	if err := dbQuery.Limit(perPage).Offset((page - 1) * perPage).Find(&users).Error; err != nil {
		return nil, types.PaginationMeta{}, err
	}

	// if err := databases.DB.Model(&models.User{}).Count(&totalData).Error; err != nil {
	// 	return nil, types.PaginationMeta{}, err
	// }

	// if err := databases.DB.Limit(perPage).Offset((page - 1) * perPage).Find(&users).Error; err != nil {
	// 	return nil, types.PaginationMeta{}, err
	// }

	totalPages := int(math.Ceil(float64(totalData) / float64(perPage)))

	meta := types.PaginationMeta{
		PerPage:    perPage,
		Page:       page,
		TotalData:  int(totalData),
		TotalPages: totalPages,
	}

	return users, meta, nil
}

func RegisterUser(name, email, password string) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{Name: name, Email: email, Password: hashedPassword}
	if err := databases.DB.Create(&user).Error; err != nil {
		return models.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func LoginUser(email, password string) (models.User, error) {
	var user models.User

	if err := databases.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	if err := VerifyPassword(user.Password, password); err != nil {
		return models.User{}, errors.New("invalid email or password")
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenString, nil
}

func ValidateLogin(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email and password are required")
	}
	return nil
}

func ValidateRegister(name, email, password string) error {
	if name == "" || email == "" || password == "" {
		return errors.New("name, email and password are required")
	}

	emailRegex := `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	var user models.User
	if err := databases.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return errors.New("email already exists")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func GetUserProfile(userID string) (models.User, error) {
	var user models.User
	if err := databases.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
