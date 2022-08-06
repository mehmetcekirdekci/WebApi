package application

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mehmetcekirdekci/WebApi/app/customer/application/helper"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/repositories"
	"github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"
)

type Service interface {
	Register(dto *types.CustomerDto) error
}

type service struct {
	customerRepository repositories.CustomerRepository
	accountInformationRepository repositories.AccountInformationRepository
}

func NewService(customerRepository repositories.CustomerRepository, accountInformationRepository repositories.AccountInformationRepository) Service {
	if customerRepository == nil || accountInformationRepository == nil {
		return nil
	}

	return &service{
		customerRepository: customerRepository,
		accountInformationRepository: accountInformationRepository,
	}
}

func (receiver *service) Register(dto *types.CustomerDto) error {
	isEnabledCustomerActivate := false
	customerFromDb, err := receiver.customerRepository.GetByMail(dto.Email)
	if err != nil {
		return errors.New("Something went wrong.")
	} else if customerFromDb != nil && customerFromDb.IsActive == true {
		return errors.New("This customer already registered.")
	} else if customerFromDb != nil && customerFromDb.IsActive == false {
		isEnabledCustomerActivate = true
	}
	customer := dto.ToCustomer()
	if isEnabledCustomerActivate {
		err = receiver.customerRepository.Activate(customer, customerFromDb.CustomerId)
		if err != nil {
			return err
		}
	}
	customer.CustomerId = uuid.New()
	err = receiver.customerRepository.Register(customer)
	if err != nil {
		return err
	}
	passwordHash := helper.CreatePasswordHash(dto.Password)
	accountInformation := dto.ToAccountInformation(customer.CustomerId, passwordHash)
	err = receiver.accountInformationRepository.InsertAccountInformation(accountInformation)
	if err != nil {
		return err
	}
	// TODO: Rollback mechanism will add.
	return nil
}