package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Wong801/gin-api/src/config"
	"github.com/Wong801/gin-api/src/db"
	entity "github.com/Wong801/gin-api/src/entities"
	model "github.com/Wong801/gin-api/src/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *db.Adapter
}

func InitUserService() *UserService {
	return &UserService{
		db: db.InitDB(),
	}
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
		Secure:   true,
		HttpOnly: true,
	}, nil
}

func (us UserService) GetUser(u *model.UserLogin) (int, *model.UserBase, error) {
	var user model.UserBase
	db.Open(us.db)

	if err := us.db.Database.First(&user, "email = ? OR username = ?", u.Email, u.Username).Error; err != nil {
		return http.StatusNotFound, nil, err
	}

	us.db.Close()

	return http.StatusFound, &user, nil
}

func (us UserService) Register(u *model.User) (int, error) {
	u.Password, _ = hashPassword(u.Password)
	db.Open(us.db)

	if err := us.db.Database.Create(&u).Error; err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			columnName := strings.Split(perr.ConstraintName, "_")
			return http.StatusConflict, errors.New(columnName[len(columnName)-1] + " is already used")
		}
		return http.StatusInternalServerError, err
	}

	us.db.Close()
	return http.StatusCreated, nil
}

func (us UserService) Login(u *model.UserLogin) (int, *entity.Token, error) {
	var user model.User
	db.Open(us.db)

	if err := us.db.Database.First(&user, "email = ? OR username = ?", u.Email, u.Username).Error; err != nil {
		return http.StatusNotFound, nil, errors.New("user not found")
	}

	us.db.Close()

	correctPassword := checkPasswordHash(u.Password, user.Password)

	if !correctPassword {
		return http.StatusBadRequest, nil, errors.New("incorrect username or password")
	}

	token, err := craftToken(user.Username, user.Id)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, token, nil
}

func (us UserService) ChangePassword(id int, u *model.UserChangePassword) (int, error) {
	if u.OldPassword != u.VerifyOldPassword {
		return http.StatusBadRequest, errors.New("old password doesn't verified")
	}

	var user model.User
	db.Open(us.db)

	us.db.Database.Where("id = ?", id).First(&user)
	correctPassword := checkPasswordHash(u.OldPassword, user.Password)

	if !correctPassword {
		return http.StatusBadRequest, errors.New("incorrect password")
	}

	user.Password, _ = hashPassword(u.NewPassword)

	us.db.Database.Save(&user)
	us.db.Close()

	return http.StatusOK, nil
}
