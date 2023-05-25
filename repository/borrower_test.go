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
func saveRandomBorrower(t *testing.T) BorrowerRepository {
	currentTime := time.Now()
	hashedpassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)
	arg := model.Borrower{
		Username:     utils.RandomOwner(),
		Password:     hashedpassword,
		Name:         utils.RandomOwner(),
		Alamat:       "Bandung",
		Phone_Number: "084758754875",
		Created_At:   currentTime,
	}

	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	repo := NewBorrowerRepositoryImpl(db)
	user, err := repo.Save(arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Alamat, user.Alamat)
	require.Equal(t, arg.Phone_Number, user.Phone_Number)

	// require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.Created_At)

	return repo
}

func TestSaveBorrower(t *testing.T) {
	saveRandomBorrower(t)
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

func createBorrower(t *testing.T, repo BorrowerRepository) model.Borrower {
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

// ================ Find By ID ========================
func TestFindByIdBorrower(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBorrowerRepositoryImpl(db)

	// Test FindById for ID 1
	foundBorrower, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected borrower with ID 1
	expectedBorrower := model.Borrower{
		Id:           1,
		Username:     "Jono",
		Password:     "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq",
		Name:         "Jono Joni",
		Alamat:       "Bandung",
		Phone_Number: "084578956",
		Created_At:   time.Date(2023, time.May, 25, 0, 0, 23, 668047000, time.Local),
	}

	require.Equal(t, expectedBorrower, foundBorrower)
}

// ================ Find By Username ========================
func TestFindByUsernameBorrower(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBorrowerRepositoryImpl(db)

	// Test FindByUsername for username Jono
	foundBorrower, err := repo.FindByUsername("Jono")
	require.NoError(t, err)

	// Expected borrower with Username Jono
	expectedBorrower := model.Borrower{
		Id:           1,
		Username:     "Jono",
		Password:     "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq",
		Name:         "Jono Joni",
		Alamat:       "Bandung",
		Phone_Number: "084578956",
		Created_At:   time.Date(2023, time.May, 25, 0, 0, 23, 668047000, time.Local),
	}

	require.Equal(t, expectedBorrower, foundBorrower)

	// Test FindByUsername for non-existing username "Nonexistent"
	_, err = repo.FindByUsername("Nonexistent")
	require.Error(t, err)
	require.EqualError(t, err, "invalid username or Password")
}

// ================ Find All ID ========================
func TestFindAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBorrowerRepositoryImpl(db)

	// Create multiple borrowers in the database
	borrowers := []model.Borrower{
		{
			Id:           1,
			Username:     "Jono",
			Password:     "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq",
			Name:         "Jono Joni",
			Alamat:       "Bandung",
			Phone_Number: "084578956",
			Created_At:   time.Date(2023, time.May, 25, 0, 0, 23, 668047000, time.Local),
		},
		{
			Id:           2,
			Username:     "Ibnu",
			Password:     "$2a$10$XsA6dtRfHC..WjBfSXss1uSVCvBaWDI0CzrLMfikaj.rBp.czOt3S",
			Name:         "Ibnu Farhan",
			Alamat:       "Bandung",
			Phone_Number: "08548785659",
			Created_At:   time.Date(2023, time.May, 25, 0, 1, 2, 94739000, time.Local),
		},
		// Add more borrowers if needed
	}

	// Test FindAll
	foundBorrowers, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(borrowers), len(foundBorrowers))

	// Compare each borrower in the expected list with the found borrowers
	for _, expectedBorrower := range borrowers {
		found := false
		for _, actualBorrower := range foundBorrowers {
			if expectedBorrower.Id == actualBorrower.Id {
				require.Equal(t, expectedBorrower, actualBorrower)
				found = true
				break
			}
		}
		require.True(t, found, "Borrower not found: ID %d", expectedBorrower.Id)
	}
}

// ================ Delete ========================
func TestDeleteBorrower(t *testing.T) {
	db := setupTestDB(t)
	repo := NewBorrowerRepositoryImpl(db)

	// Create a borrower for deletion
	borrower := model.Borrower{
		Id:           3,
		Username:     "Jini",
		Password:     "123456",
		Name:         "Jin",
		Alamat:       "Bekasi",
		Phone_Number: "084579",
		Created_At:   time.Now().Local(),
	}

	// Save the borrower
	_, err := repo.Save(borrower)
	require.NoError(t, err)

	// Delete the borrower
	_, err = repo.Delete(3)
	require.NoError(t, err)
}

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}
