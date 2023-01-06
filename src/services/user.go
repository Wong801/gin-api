package service

import (
	"errors"
	"fmt"
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

type UserService struct {
}

func InitUserService() *UserService {
	return &UserService{}
}

func hashPassword(password string) (string, error) {
	salt, _ := strconv.Atoi(config.GetEnv("JWT_SALT", "10"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func craftToken(username string, id int) (*entity.Token, error) {
	duration, _ := strconv.Atoi(config.GetEnv("JWT_DURATION", "24"))

	expiration := time.Now().UTC().Add(time.Hour * time.Duration(duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    username,
		ID:        fmt.Sprint(id),
	})

	tokenString, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET", "secret")))

	if err != nil {
		return nil, err
	}

	return &entity.Token{
		Jwt:      tokenString,
		MaxAge:   int((time.Hour / time.Second) * time.Duration(duration)),
		Domain:   config.GetEnv("CORS_DOMAIN", "localhost"),
		Secure:   false,
		HttpOnly: true,
	}, nil
}

func (us UserService) Register(u *model.User) (int, error) {
	DB := db.InitDB()

	u.Password, _ = hashPassword(u.Password)

	if err := DB.Database.Create(&u).Error; err != nil {
		return http.StatusInternalServerError, err
	}

	DB.Close()
	return http.StatusOK, nil
}

func (us UserService) Login(u *model.UserLogin) (int, *entity.Token, error) {
	DB := db.InitDB()
	var user model.User

	DB.Database.Where("email = ?", u.Email).Or("username = ?", u.Username).First(&user)

	correctPassword := checkPasswordHash(u.Password, user.Password)

	DB.Close()

	if !correctPassword {
		return http.StatusBadRequest, nil, errors.New("incorrect username or password")
	}

	token, err := craftToken(user.Username, user.Id)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, token, nil
}
