package repository

import (
	"testing"
	"time"

	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/utils"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ================ Save ========================
func saveRandomLender(t *testing.T) LenderRepository {
	currentTime := time.Now()

	arg := model.Lender{
		Name:       utils.RandomOwner(),
		Created_At: currentTime,
	}

	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	repo := NewLenderRepositoryImpl(db)
	user, err := repo.Save(arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)

	// require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.Created_At)

	return repo
}

func TestSaveLender(t *testing.T) {
	saveRandomLender(t)
}

// ================ Update ========================
func TestUpdateBorrower(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBorrowerRepositoryImpl(db)

	// Create a random borrower
	borrower := createBorrower(t, repo)

	// Update the borrower
	borrower.Username = utils.RandomOwner()
	borrower.Password = utils.RandomString(8)
	borrower.Name = utils.RandomOwner()
	borrower.Alamat = "Jakarta"
	borrower.Phone_Number = "123456789"

	// Test Update
	repo.Update(borrower)

	// Retrieve the updated borrower
	updatedBorrower, err := repo.FindById(borrower.Id)
	require.NoError(t, err)

	// Ensure the borrower is updated
	require.Equal(t, borrower.Username, updatedBorrower.Username)
	require.Equal(t, borrower.Password, updatedBorrower.Password)
	require.Equal(t, borrower.Name, updatedBorrower.Name)
	require.Equal(t, borrower.Alamat, updatedBorrower.Alamat)
	require.Equal(t, borrower.Phone_Number, updatedBorrower.Phone_Number)
}

func createLender(t *testing.T, repo BorrowerRepository) model.Borrower {
	currentTime := time.Now()
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	borrower := model.Borrower{
		Username:     utils.RandomOwner(),
		Password:     hashedPassword,
		Name:         utils.RandomOwner(),
		Alamat:       "Bandung",
		Phone_Number: "084758754875",
		Created_At:   currentTime,
	}

	// Save the borrower
	createdBorrower, err := repo.Save(borrower)
	require.NoError(t, err)
	require.NotEmpty(t, createdBorrower)
	require.Equal(t, borrower.Username, createdBorrower.Username)
	require.Equal(t, borrower.Password, createdBorrower.Password)
	require.Equal(t, borrower.Name, createdBorrower.Name)
	require.Equal(t, borrower.Alamat, createdBorrower.Alamat)
	require.Equal(t, borrower.Phone_Number, createdBorrower.Phone_Number)
	require.NotZero(t, createdBorrower.Created_At)

	return createdBorrower
}
