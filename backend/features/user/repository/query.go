package repository

import (
	"errors"
	"time"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/user"
	"github.com/dimasyudhana/simple-chat/utils/identity"
	"github.com/dimasyudhana/simple-chat/utils/password"
	"gorm.io/gorm"
)

var log = middlewares.Log()

type Query struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &Query{
		db: db,
	}
}

// Register implements user.UserData.
func (uq *Query) Register(request user.UserCore) (user.UserCore, error) {
	userId, err := identity.GenerateID()
	if err != nil {
		log.Error("error while creating id for user")
		return user.UserCore{}, errors.New("error while creating id for user")
	}

	hashed, err := password.HashPassword(request.Password)
	if err != nil {
		log.Error("error while hashing password")
		return user.UserCore{}, errors.New("error while hashing password")
	}

	request.UserID = userId
	request.Password = hashed
	req := coreToModel(request)
	query := uq.db.Exec(`
		INSERT INTO users (user_id, username, email, password, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING user_id, username, email, created_at, updated_at
	`, req.UserID, req.Username, req.Email, req.Password, time.Now(), time.Now())

	if errors.Is(query.Error, gorm.ErrRegistered) {
		log.Error("error insert data, duplicate input")
		return user.UserCore{}, errors.New("error insert data, duplicate input")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no row affected")
		return user.UserCore{}, errors.New("no row affected")
	}

	log.Sugar().Infof("new user has been created: %s", req.UserID)
	return modelToCore(req), nil
}

// Login implements user.UserData.
func (uq *Query) Login(request user.UserCore) (user.UserCore, string, error) {
	result := User{}
	query := uq.db.Raw(`
		SELECT user_id, email, password
		FROM users
		WHERE email = ?
		LIMIT 1
	`, request.Email).Scan(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		log.Error("user record not found")
		return user.UserCore{}, "", errors.New("invalid email and password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		log.Warn("no row affected")
		return user.UserCore{}, "", errors.New("no row affected")
	}

	if !password.MatchPassword(request.Password, result.Password) {
		return user.UserCore{}, "", errors.New("password does not match")
	}

	token, err := middlewares.GenerateToken(result.UserID)
	if err != nil {
		log.Error("error while creating jwt token")
		return user.UserCore{}, "", errors.New("error while creating jwt token")
	}

	log.Sugar().Infof("user has been logged in: %s", result.UserID)
	return modelToCore(result), token, nil
}
