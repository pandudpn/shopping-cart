package user

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_USER_INSERT = `insert into public.user (name, email, password, phone, enabled, email_verified_at, created_at)
						values ($1, $2, $3, $4, $5, $6, $7) returning id`
)

func (ur *UserRepository) InsertUser(user *model.User) error {
	stmt, err := ur.DB.Prepare(QUERY_USER_INSERT)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(user.Name, user.Email, user.Password, user.Phone, user.Enabled, user.EmailVerifiedAt, user.CreatedAt).Scan(&user.Id)
	if err != nil {
		return err
	}

	return nil
}
