package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/abmid/icanvas-analytics/pkg/user/entity"
	"github.com/abmid/icanvas-analytics/pkg/user/usecase"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type loginUseCase struct {
	userUC usecase.UserUseCase
}

func New(userUC usecase.UserUseCase) *loginUseCase {
	return &loginUseCase{
		userUC: userUC,
	}
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	fmt.Printf("Hashed : %s", hashedPwd)
	fmt.Printf("PLaine : %s", plainPwd)
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (UC *loginUseCase) Login(email, password string) (res *entity.User, httpStatus int, err error) {

	user, err := UC.userUC.Find(email)
	if err != nil {
		logrus.Error(err)
		if err == sql.ErrNoRows {
			return nil, http.StatusUnauthorized, errors.New("User not found")
		}
		return nil, http.StatusUnauthorized, errors.New("Failed get resource")
	}

	if user == nil {
		return nil, http.StatusUnauthorized, errors.New("User not found")
	}

	checkPassword := comparePasswords(user.Password, password)
	if checkPassword {
		return user, http.StatusOK, nil
	}
	return nil, http.StatusUnauthorized, errors.New("Wrong password !")
}
