package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/pinjol/model"
	"gorm.io/gorm"
)

type LendersRepositoryImpl struct {
	Db *gorm.DB
}

func NewLendersRepositoryImpl(Db *gorm.DB) LendersRepository {
	return &LendersRepositoryImpl{Db: Db}
}

// Delete implements LendersRepository
func (r *LendersRepositoryImpl) Delete(id int64) (model.Lenders, error) {
	var bor model.Lenders
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements LendersRepository
func (r *LendersRepositoryImpl) FindAll() ([]model.Lenders, error) {
	var bor []model.Lenders
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements LendersRepository
func (r *LendersRepositoryImpl) FindById(id int64) (model.Lenders, error) {
	var bor model.Lenders
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("lenders is not found")
	}
	return bor, nil
}

// Save implements LendersRepository
func (r *LendersRepositoryImpl) Save(newLenders model.Lenders) (model.Lenders, error) {
	currentTime := time.Now()
	newLenders.Created_At = currentTime
	result := r.Db.Create(&newLenders)
	helper.ErrorPanic(result.Error)
	return newLenders, nil
}

// Update implements LendersRepository
func (r *LendersRepositoryImpl) Update(updatedLenders model.Lenders) (model.Lenders, error) {
	result := r.Db.Model(&model.Lenders{}).Where("id = ?", updatedLenders.Id).Updates(updatedLenders)
	helper.ErrorPanic(result.Error)
	return updatedLenders, nil
}

// FindByName implements LendersRepository
func (r *LendersRepositoryImpl) FindByName(name string) (model.Lenders, error) {
	var bor model.Lenders
	result := r.Db.First(&bor, "name = ?", name)

	if result.Error != nil {
		return bor, errors.New("invalid name")
	}
	return bor, nil
}
