package service

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Wong801/gin-api/src/config"
	"github.com/Wong801/gin-api/src/db"
	entity "github.com/Wong801/gin-api/src/entities"
	model "github.com/Wong801/gin-api/src/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	salt, _ := strconv.Atoi(config.GetEnv("JWT_SALT", "10"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(u *model.User) (int, error) {
	DB := db.InitDB()

	u.Password, _ = hashPassword(u.Password)

	if err := DB.Database.Create(&u).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	DB.Close()
	return http.StatusOK, nil
}

func LoginUser(u *model.UserLogin) (int, *entity.Token, error) {
	DB := db.InitDB()
	var user model.User

	DB.Database.Where("email = ?", u.Email).Or("username = ?", u.Username).First(&user)

	correctPassword := checkPasswordHash(u.Password, user.Password)

	DB.Close()

	if !correctPassword {
		return http.StatusBadRequest, nil, errors.New("incorrect username or password")
	}

	baseUser := &model.UserBase{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		DoB:       user.DoB,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	duration, _ := strconv.Atoi(config.GetEnv("JWT_DURATION", "24"))

	expiration := time.Now().UTC().Add(time.Hour * time.Duration(duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": baseUser,
		"iat":  time.Now().UTC(),
		"exp":  expiration,
	})

	tokenString, err := token.SignedString(config.GetEnv("JWT_SECRET", "secret"))

	return http.StatusOK, &entity.Token{
		Jwt:      tokenString,
		MaxAge:   int((time.Hour / time.Second) * time.Duration(duration)),
		Domain:   config.GetEnv("JWT_DOMAIN", "localhost"),
		Secure:   false,
		HttpOnly: true,
	}, err
}
