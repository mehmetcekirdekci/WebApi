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
	isCustomerAlreadyExist, isEnabledCustomerActivate, customerFromDb, err := CustomerEnableCheck(dto.Email, receiver)
	if err != nil {
		return err
	}
	customer := dto.ToCustomer()
	if isCustomerAlreadyExist && !isEnabledCustomerActivate {
		return errors.New(types.IsCustomerAlreadyExistErrorMessage)
	}
	if isCustomerAlreadyExist && isEnabledCustomerActivate {
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

func CustomerEnableCheck(email string, receiver *service) (bool, bool, *types.Customer, error) {
	filter := types.Customer{
		Email: email,
	}
	customerFromDb, err := receiver.customerRepository.GetByFilter(&filter)
	if err != nil {
		return false, false, nil, errors.New(types.CustomerRegisterErrorMessage)
	} else if customerFromDb != nil && customerFromDb.IsActive == true {
		return true, false, customerFromDb, nil
	} else if customerFromDb != nil && customerFromDb.IsActive == false {
		return true, true, customerFromDb, nil
	}
	return false, false, customerFromDb, nil
}