package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alfathaulia/ca_ecommerce_api/domain"
	"github.com/alfathaulia/ca_ecommerce_api/user/repository"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

// NewmysqlUserRepo will create an object that represent the domain.UserRepository interface
func NewMysqlUserRepo(DB *sql.DB) domain.UserRepository {
	return &mysqlUserRepo{DB: DB}
}

func (m *mysqlUserRepo) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.Email,
			&t.HashedPassword,
			&t.Role,
			&t.IsVerified,
			&t.UpdatedAt,
			&t.CreatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlUserRepo) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at FROM user ORDER BY created_at LIMIT ?`

	decodeCursor, err := repository.DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", domain.ErrBadParamInput
	}
	res, err = m.fetch(ctx, query, decodeCursor, num)
	if err != nil {
		return nil, "", err
	}
	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}
	return
}
func (m *mysqlUserRepo) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at
  						FROM user WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
func (m *mysqlUserRepo) GetByUsername(ctx context.Context, username string) (res domain.User, err error) {
	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at
  						FROM user WHERE username = ?`

	list, err := m.fetch(ctx, query, username)
	if err != nil {
		return res, domain.ErrNotFound
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}
func (m *mysqlUserRepo) Store(ctx context.Context, data *domain.User) (err error) {
	query := `INSERT  user SET username=? , email=? , hashed_password=?, role=?, is_verified=?, updated_at=? , created_at=?`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, data.Username, data.Email, data.HashedPassword, data.Role, data.IsVerified, data.UpdatedAt, data.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	data.ID = lastID
	return
}
func (m *mysqlUserRepo) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM user WHERE id = ?"

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  behavior. total affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysqlUserRepo) Update(ctx context.Context, dataUpdate *domain.User) (err error) {
	query := `UPDATE  user SET username=? , email=? , hashed_password=?, role=?, updated_at=? WHERE id=?`

	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, dataUpdate.Username, dataUpdate.Email, dataUpdate.HashedPassword, dataUpdate.Role, dataUpdate.UpdatedAt, dataUpdate.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  behavior. total affected: %d", affect)
		return
	}

	return
}

func (m *mysqlUserRepo) Register(ctx context.Context, users *domain.User) (err error) {
	query := `INSERT  user SET username=? , email=? , hashed_password=?, role=?, is_verified=?, updated_at=? , created_at=?`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, users.Username, users.Email, users.HashedPassword, "user", users.UpdatedAt, users.CreatedAt)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	users.ID = lastID

	return
}

func (m *mysqlUserRepo) Login(ctx context.Context, username string, password string) (res domain.User, err error) {
	// query := `SELECT id, username, unit_code, hashed_password, role, updated_at, created_at FROM user WHERE username = ? `

	// query := `SELECT id, username, unit_code, hashed_password, role, updated_at, created_at
	// 						FROM user WHERE username = ?`

	list, err := m.GetByUsername(ctx, username)
	if err != nil {
		return
	}

	res = list
	if res.HashedPassword != password {
		return domain.User{}, domain.ErrBadParamInput
	}
	if res.Username != username {
		return domain.User{}, domain.ErrInternalServerError
	}
	return res, nil

}
