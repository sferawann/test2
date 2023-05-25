package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createRandomBorrower(t *testing.T) BorrowerRepository {
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
	createRandomBorrower(t)
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

func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjaman_online port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

func TestFindAll(t *testing.T) {
	// Setup mocking
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewBorrowerRepositoryImpl(db)

	// Expected rows returned from the database
	rows := sqlmock.NewRows([]string{"id", "username", "password", "name", "alamat", "phone_number", "created_at"}).
		AddRow(1, "Jono", "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq", "Jono Joni", "Bandung", "084578956", "2023-05-25 00:00:23").
		AddRow(2, "Jane", "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq", "Jane Doe", "Jakarta", "084578957", "2023-05-25 00:00:45")

	// Expect the query to fetch all borrowers
	mock.ExpectQuery("SELECT \\* FROM borrowers").WillReturnRows(rows)

	// Call FindAll
	foundBorrowers, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, foundBorrowers, 2)

	// Assert the expected borrowers
	expectedBorrowers := []model.Borrower{
		{
			Id:           1,
			Username:     "Jono",
			Password:     "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq",
			Name:         "Jono Joni",
			Alamat:       "Bandung",
			Phone_Number: "084578956",
			Created_At:   time.Date(2023, time.May, 25, 0, 0, 23, 0, time.UTC),
		},
		{
			Id:           2,
			Username:     "Jane",
			Password:     "$2a$10$ycv4VS.TiL.USMExDTpQZuAExOoUYdUWP8ahQWGMacscM/KR4Thzq",
			Name:         "Jane Doe",
			Alamat:       "Jakarta",
			Phone_Number: "084578957",
			Created_At:   time.Date(2023, time.May, 25, 0, 0, 45, 0, time.UTC),
		},
	}

	assert.ElementsMatch(t, expectedBorrowers, foundBorrowers)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
