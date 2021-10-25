package mysql_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alfathaulia/ca_ecommerce_api/domain"
	"github.com/alfathaulia/ca_ecommerce_api/user/repository"
	userMysqlRepo "github.com/alfathaulia/ca_ecommerce_api/user/repository/mysql"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var user = &domain.User{
	ID:             1,
	Username:       "user1",
	Email:          "user1@gmail.com",
	HashedPassword: "haspass",
	IsVerified:     true,
	Role:           "user",
	UpdatedAt:      time.Now(),
	CreatedAt:      time.Now(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFetch(t *testing.T) {
	db, mock := NewMock()

	mockUsers := []domain.User{
		{
			ID: user.ID, Username: user.Username, Email: user.Email, HashedPassword: user.HashedPassword, Role: user.Role, IsVerified: user.IsVerified, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt,
		},
		{
			ID: 2, Username: "User 2", Email: "122456", HashedPassword: "user2", Role: "user", IsVerified: true, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "username", "email", "hashed_password", "role", "is_verified", "updated_at", "created_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Username, mockUsers[0].Email, mockUsers[0].HashedPassword, mockUsers[0].Role, mockUsers[0].IsVerified, mockUsers[0].UpdatedAt, mockUsers[0].CreatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Username, mockUsers[1].Email, mockUsers[1].HashedPassword, mockUsers[1].Role, mockUsers[1].IsVerified, mockUsers[1].UpdatedAt, mockUsers[1].CreatedAt)

	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at FROM user  ORDER BY created_at LIMIT ?`

	mock.ExpectQuery(query).WillReturnRows(rows)

	a := userMysqlRepo.NewMysqlUserRepo(db)
	cursor := repository.EncodeCursor(mockUsers[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NoError(t, err)
	assert.NotEmpty(t, nextCursor)
	assert.Len(t, list, 2)

	assert.Equal(t, user.Username, list[0].Username)
	assert.Equal(t, user.HashedPassword, list[0].HashedPassword)
	assert.Equal(t, user.Email, list[0].Email)
	assert.Equal(t, user.Role, list[0].Role)
	assert.Equal(t, user.IsVerified, list[0].IsVerified)
	assert.Equal(t, user.CreatedAt, list[0].CreatedAt)
	assert.Equal(t, user.UpdatedAt, list[0].UpdatedAt)
	assert.NotZero(t, list[0].CreatedAt)

}

func TestGetByID(t *testing.T) {
	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"id", "username", "email", "hashed_password", "role", "is_verified", "updated_at", "created_at"}).
		AddRow(1, user.Username, user.Email, user.HashedPassword, user.Role, user.IsVerified, user.UpdatedAt, user.CreatedAt)

	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at FROM user WHERE ID = ?`

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := userMysqlRepo.NewMysqlUserRepo(db)

	num := int64(5)
	user, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user.Username, user.Username)
	assert.Equal(t, user.HashedPassword, user.HashedPassword)
	assert.Equal(t, user.Email, user.Email)
	assert.Equal(t, user.Role, user.Role)
	assert.Equal(t, user.IsVerified, user.IsVerified)
	assert.Equal(t, user.CreatedAt, user.CreatedAt)
	assert.Equal(t, user.UpdatedAt, user.CreatedAt)
	assert.NotZero(t, user.CreatedAt)
}

func TestGetByUsername(t *testing.T) {
	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"id", "username", "email", "hashed_password", "role", "is_verified", "updated_at", "created_at"}).AddRow(1, user.Username, user.Email, user.HashedPassword, user.Role, user.IsVerified, user.CreatedAt, user.UpdatedAt)

	query := `SELECT id, username, email, hashed_password, role, is_verified, updated_at, created_at FROM user WHERE username = ?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := userMysqlRepo.NewMysqlUserRepo(db)
	userName := "user1"
	user, err := a.GetByUsername(context.TODO(), userName)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestStore(t *testing.T) {
	db, mock := NewMock()

	query := "INSERT  user SET username=\\? , email=\\? , hashed_password=\\?, role=\\?, is_verified=\\?, updated_at=\\? , created_at=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Email, user.HashedPassword, user.Role, user.IsVerified, user.UpdatedAt, user.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	a := userMysqlRepo.NewMysqlUserRepo(db)
	err := a.Store(context.TODO(), user)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE  user SET username=\\? , email=\\? , hashed_password=\\?, role=\\?, updated_at=\\?  WHERE id=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(user.Username, user.Email, user.HashedPassword, user.Role, user.UpdatedAt, user.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := userMysqlRepo.NewMysqlUserRepo(db)

	err = a.Update(context.TODO(), user)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()

	query := "DELETE FROM user WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := userMysqlRepo.NewMysqlUserRepo(db)

	num := int64(12)
	err := a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}
