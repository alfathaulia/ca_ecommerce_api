package usecase

import (
	"context"
	"time"

	"github.com/alfathaulia/ca_ecommerce_api/domain"
	util "github.com/alfathaulia/ca_ecommerce_api/user/repository/utils"
)

const timeout = time.Second * 10

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (m *userUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, nextCursor, err = m.userRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}
	return

}

func (m *userUsecase) GetByID(ctx context.Context, id int64) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err = m.userRepo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return
}
func (m *userUsecase) GetByUsername(ctx context.Context, username string) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	res, err = m.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return domain.User{}, nil
	}
	return
}

func (m *userUsecase) Update(ctx context.Context, ar *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	ar.UpdatedAt = time.Now()
	return m.userRepo.Update(ctx, ar)
}

func (m *userUsecase) Store(ctx context.Context, a *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	user, _ := m.GetByUsername(ctx, a.Username)
	if user != (domain.User{}) {
		return domain.ErrConflict
	}
	err = m.userRepo.Store(ctx, a)
	return
}
func (m *userUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	existedArticle, err := m.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (domain.User{}) {
		return domain.ErrNotFound
	}
	return m.userRepo.Delete(ctx, id)
}

func (m *userUsecase) Register(ctx context.Context, user *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	hashPass, err := util.HashPassword(user.HashedPassword)
	if err != nil {
		return err
	}
	user.HashedPassword = hashPass
	existedArticle, _ := m.GetByUsername(ctx, user.Username)
	if existedArticle != (domain.User{}) {
		return domain.ErrConflict
	}
	return m.userRepo.Register(ctx, user)
}

func (m *userUsecase) Login(ctx context.Context, username string, password string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	res, err := m.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return domain.User{}, domain.ErrNotFound
	}
	checkPass := util.CheckPassword(password, res.HashedPassword)
	if checkPass != nil {
		return domain.User{}, domain.ErrBadParamInput
	}
	if res.Username != username {
		return domain.User{}, domain.ErrInternalServerError
	}

	return res, err
}

func (m *userUsecase) CreateAdmin(ctx context.Context, a *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	existedArticle, _ := m.GetByUsername(ctx, a.Username)
	if existedArticle != (domain.User{}) {
		return domain.ErrConflict
	}
	a.Role = "admin"
	err = m.userRepo.Store(ctx, a)
	return
}

func (m *userUsecase) CreateStaff(ctx context.Context, a *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	existedArticle, _ := m.GetByUsername(ctx, a.Username)
	if existedArticle != (domain.User{}) {
		return domain.ErrConflict
	}
	a.Role = "staff"
	err = m.userRepo.Store(ctx, a)
	return
}
