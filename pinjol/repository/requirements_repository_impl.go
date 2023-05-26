package repository

import (
	"errors"
	"time"

	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/pinjol/model"
	"gorm.io/gorm"
)

type RequirementsRepositoryImpl struct {
	Db *gorm.DB
}

func NewRequirementsRepositoryImpl(Db *gorm.DB) RequirementsRepository {
	return &RequirementsRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository
func (r *RequirementsRepositoryImpl) Delete(id int64) (model.Requirements, error) {
	var bor model.Requirements
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements UsersRepository
func (r *RequirementsRepositoryImpl) FindAll() ([]model.Requirements, error) {
	var bor []model.Requirements
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements UsersRepository
func (r *RequirementsRepositoryImpl) FindById(id int64) (model.Requirements, error) {
	var bor model.Requirements
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("requirements is not found")
	}
	return bor, nil
}

// Save implements UsersRepository
func (r *RequirementsRepositoryImpl) Save(newRequirements model.Requirements) (model.Requirements, error) {
	currentTime := time.Now()
	newRequirements.Created_At = currentTime
	result := r.Db.Create(&newRequirements)
	helper.ErrorPanic(result.Error)
	return newRequirements, nil
}

// Update implements UsersRepository
func (r *RequirementsRepositoryImpl) Update(updatedRequirements model.Requirements) (model.Requirements, error) {
	result := r.Db.Model(&model.Requirements{}).Where("id = ?", updatedRequirements.Id).Updates(updatedRequirements)
	helper.ErrorPanic(result.Error)
	return updatedRequirements, nil
}

// FindByName implements UsersRepository
func (r *RequirementsRepositoryImpl) FindByName(name string) (model.Requirements, error) {
	var bor model.Requirements
	result := r.Db.First(&bor, "name = ?", name)

	if result.Error != nil {
		return bor, errors.New("invalid name")
	}
	return bor, nil
}
