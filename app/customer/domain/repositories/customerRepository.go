package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Register(customer *types.Customer) error
	GetByFilter(filter *types.Customer) (*types.Customer, error)
	Activate(customer *types.Customer, customerId uuid.UUID) error
}

type (
	customerRepository struct {
		db *gorm.DB
	}
)

func NewCustomerRepository(database *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: database,
	}
}

func (receiver *customerRepository) Register(customer *types.Customer) error  {
	err := receiver.db.Table(types.CustomersTable).Create(&customer).Error
	if err != nil {
		return errors.New("Customer can not be registered.")
	}
	return nil
}

func (receiver *customerRepository) Activate(customer *types.Customer, customerId uuid.UUID) error  {
	err := receiver.db.Table(types.CustomersTable).Update(customerId.String(), &customer).Error
	if err != nil {
		return errors.New("Customer can not be registered.")
	}
	return nil
}

func (receiver *customerRepository) Deactivate(customer *types.Customer) error  {
	err := receiver.db.Table(types.CustomersTable).Create(&customer).Error
	if err != nil {
		return errors.New("Customer can not be registered.")
	}
	return nil
}

func (receiver *customerRepository) GetByFilter(filter *types.Customer) (*types.Customer, error) {
	var customer types.Customer
	err := receiver.db.Table(types.CustomersTable).Where(filter).First(&customer).Error
	if customer.Id == 0 {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &customer, nil
}