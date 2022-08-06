package controller

import "github.com/mehmetcekirdekci/WebApi/app/customer/domain/types"

func (request *RegisterCustomerRequest) ToDto() *types.CustomerDto {
	dto := types.CustomerDto{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		BirthDate: request.BirthDate,
		Gender:    request.Gender,
		Adress:    request.Adress,
		Password:  request.Password,
	}
	return &dto
}