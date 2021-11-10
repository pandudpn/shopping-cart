package userusecase

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

func (uu *UserUseCase) getUserByEmail(email string, isRegister bool) (*model.User, error) {
	user, err := uu.UserRepo.FindByEmail(email)
	if err != nil {
		logger.Log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			if isRegister {
				return model.NewUser(), nil
			}
			return nil, errors.New("user.email.not_found")
		}

		return nil, errors.New("query.find.error")
	}

	if !user.IsVerified() || !user.Enabled {
		return nil, errors.New("user.not_active")
	}

	return user, nil
}

func (uu *UserUseCase) getUserByPhone(phone string, isRegister bool) (*model.User, error) {
	user, err := uu.UserRepo.FindByPhone(phone)
	if err != nil {
		logger.Log.Error(err)

		if errors.Is(err, sql.ErrNoRows) {
			if isRegister {
				return model.NewUser(), nil
			}
			return nil, errors.New("user.phone.not_found")
		}

		return nil, errors.New("query.find.error")
	}

	if !user.IsVerified() || !user.Enabled {
		return nil, errors.New("user.not_active")
	}

	return user, nil
}

func (uu *UserUseCase) register(req *model.RequestRegister) error {
	isRegister := true
	user, err := uu.getUserByEmail(req.Email, isRegister)
	if err != nil {
		logger.Log.Debugf("error get user by email %v", err)
		return err
	}

	if user.Id != 0 {
		err = errors.New("user.email.taken")

		return err
	}

	phone, err := validationPhone(req.Phone)
	if err != nil {
		return err
	}

	user, err = uu.getUserByPhone(phone, isRegister)
	if err != nil {
		logger.Log.Debugf("error get user by phone %v", err)
		return err
	}

	if user.Id != 0 {
		err = errors.New("user.phone.taken")

		return err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error(err)
		err = errors.New("user.password.error")

		return err
	}
	now := time.Now().UTC()

	user.Email = req.Email
	user.Enabled = true // langsung buat enable. jika ingin melakukan verifikasi by email, cukup dibuat false
	user.Name = req.Name
	user.Phone = phone
	user.Password = string(pwd) // hasil yg sudah di hash menggunakan bcrypt
	user.EmailVerifiedAt = &now
	user.CreatedAt = now

	err = uu.UserRepo.InsertUser(user)
	if err != nil {
		logger.Log.Errorf("error insert user %v", err)
		err = errors.New("query.insert.error")

		return err
	}

	return nil
}

func validationPhone(phone string) (string, error) {
	var err error
	if len(phone) < 10 {
		logger.Log.Errorf("nomer telpon dibawah 10 digit %s", phone)
		err = errors.New("phone.invalid")
		return "", err
	}

	if phone[:2] == "62" {
		return phone, nil
	} else if phone[:1] == "0" && phone[:2] == "08" {
		phone = fmt.Sprintf("62%s", phone[1:])
	} else if phone[:1] == "8" {
		phone = fmt.Sprintf("62%s", phone)
	} else {
		logger.Log.Errorf("nomer telpon tidak sesuai %s", phone)
		err = errors.New("phone.invalid")
		return "", err
	}

	return phone, nil
}
