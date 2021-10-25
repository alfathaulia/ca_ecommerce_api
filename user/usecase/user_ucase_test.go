package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alfathaulia/ca_ecommerce_api/domain"
	"github.com/alfathaulia/ca_ecommerce_api/domain/mocks"
	util "github.com/alfathaulia/ca_ecommerce_api/user/repository/utils"
	ucase "github.com/alfathaulia/ca_ecommerce_api/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListUser, "next-cursor", nil).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListUser))
		mockUserRepo.AssertExpectations(t)

	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(nil, "", errors.New("Unexpected Error")).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)

		assert.Empty(t, nextCursor)
		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockUserRepo.AssertExpectations(t)

	})
}

func TestGetByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		a, err := u.GetByID(context.TODO(), mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.User{}, errors.New("Unexpected ")).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		a, err := u.GetByID(context.TODO(), mockUser.ID)

		assert.Error(t, err)
		assert.Equal(t, domain.User{}, a)

		mockUserRepo.AssertExpectations(t)

	})
}

func TestStore(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = 0
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()
		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &tempMockUser)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, tempMockUser.Username)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("existing-username", func(t *testing.T) {
		existingUser := mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(existingUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &mockUser)

		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)

	})
}

func TestDelete(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockUser, nil).Once()
		mockUserRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Delete(context.TODO(), mockUser.ID)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("user-is-not-exist", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.User{}, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Delete(context.TODO(), mockUser.ID)
		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockUserRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.User{}, errors.New("Unexpected Error")).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Delete(context.TODO(), mockUser.ID)

		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Update", mock.Anything, &mockUser).Once().Return(nil)

		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)

		err := u.Update(context.TODO(), &mockUser)
		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	pass := util.RandomString(6)
	hashPass, err := util.HashPassword(pass)
	assert.NoError(t, err)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: hashPass,
		Role:           "user",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = 0
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()

		mockUserRepo.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Register(context.TODO(), &tempMockUser)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, tempMockUser.Username)
		assert.NotEqual(t, mockUser.HashedPassword, pass)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("existing-username", func(t *testing.T) {
		existingUser := mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(existingUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Register(context.TODO(), &mockUser)

		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)

	})
}

// func TestLogin(t *testing.T) {
// 	mockUserRepo := new(mocks.UserRepository)
// 	pass := util.RandomString(6)
// 	hashPass, err := util.HashPassword(pass)
// 	assert.NoError(t, err)
// 	mockUser := domain.User{
// 		ID:             1,
// 		Username:       "user1",
// 		Email:          "user1@gmail.com",
// 		HashedPassword: hashPass,
// 		Role:           "user",
// 		UpdatedAt:      time.Now(),
// 		CreatedAt:      time.Now(),
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		tempMockUser := mockUser
// 		tempMockUser.ID = 1
// 		mockUserRepo.On("GetByUsername", mock.Anything, &mockUser.Username).Once().Return(nil)

// 		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
// 		user, err := u.Login(context.TODO(), mockUser.Username, mockUser.HashedPassword)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, user)
// 		assert.Equal(t, mockUser.Username, tempMockUser.Username)
// 		assert.NotEqual(t, mockUser.HashedPassword, user.HashedPassword)
// 		mockUserRepo.AssertExpectations(t)
// 	})
// 	t.Run("user-not-found", func(t *testing.T) {
// 		tempMockUser := mockUser
// 		tempMockUser.ID = 0
// 		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()

// 		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, u)

// 		mockUserRepo.AssertExpectations(t)
// 	})

// 	t.Run("password-wrong", func(t *testing.T) {
// 		tempMockUser := mockUser
// 		tempMockUser.ID = 0
// 		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrBadParamInput).Once()

// 		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
// 		user, err := u.Login(context.TODO(), mockUser.Username, mockUser.HashedPassword)
// 		assert.NoError(t, err)
// 		assert.NotNil(t, user)
// 		assert.NotEqual(t, tempMockUser.HashedPassword, user.HashedPassword)

// 		mockUserRepo.AssertExpectations(t)
// 	})

// }

func TestCreateAdmin(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "admin",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = 0
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()
		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &tempMockUser)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, tempMockUser.Username)
		assert.Equal(t, mockUser.Role, tempMockUser.Role)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("existing-username", func(t *testing.T) {
		existingUser := mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(existingUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &mockUser)

		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)

	})

}

func TestCreateStaff(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockUser := domain.User{
		ID:             1,
		Username:       "user1",
		Email:          "user1@gmail.com",
		HashedPassword: "sdfasdf",
		Role:           "staff",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.ID = 0
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(domain.User{}, domain.ErrNotFound).Once()
		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &tempMockUser)
		assert.NoError(t, err)
		assert.Equal(t, mockUser.Username, tempMockUser.Username)
		assert.Equal(t, mockUser.Role, tempMockUser.Role)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("existing-username", func(t *testing.T) {
		existingUser := mockUser
		mockUserRepo.On("GetByUsername", mock.Anything, mock.AnythingOfType("string")).Return(existingUser, nil).Once()
		u := ucase.NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Store(context.TODO(), &mockUser)

		assert.Error(t, err)
		mockUserRepo.AssertExpectations(t)

	})

}
