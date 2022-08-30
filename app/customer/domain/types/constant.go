package types

type (
	GenderTypeEnum int8
)

// Customer entity's gender constant variables.
const (
	Male GenderTypeEnum = iota
	Female
	Unknown
)

// db table's constant variables
const (
	CustomersTable = "Customers"
	AccountInformationsTable = "AccountInformations"
)

const (
	CustomerRegisterErrorMessage = "Something went wrong."
	IsCustomerAlreadyExistErrorMessage = "Customer registered succesfuly."
)