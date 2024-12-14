package cqrs

import (
	"context"

	"gorm.io/gorm"
)

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	GetTransaction() *gorm.DB
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type unitOfWork struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (uow *unitOfWork) Begin() error {
	if uow.tx != nil {
		return nil
	}
	uow.tx = uow.db.Begin()
	return uow.tx.Error
}

func (uow *unitOfWork) Commit() error {
	if uow.tx == nil {
		return nil
	}
	err := uow.tx.Commit().Error
	uow.tx = nil
	return err
}

func (uow *unitOfWork) Rollback() error {
	if uow.tx == nil {
		return nil
	}
	err := uow.tx.Rollback().Error
	uow.tx = nil
	return err
}

func (uow *unitOfWork) GetTransaction() *gorm.DB {
	if uow.tx != nil {
		return uow.tx
	}
	return uow.db
}

func (uow *unitOfWork) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	if uow.tx != nil {
		return fn(uow.tx)
	}

	return uow.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		uow.tx = tx
		defer func() {
			uow.tx = nil
		}()
		return fn(tx)
	})
}
