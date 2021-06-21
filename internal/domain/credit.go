package credit

import (
)

type Credit struct {
	Value       int
	Debenture	string
	EndDate		int
	StartDate	int
	Status		int
}

func NewCredit(value, status, endDate, startDate int, debenture string) (Credit, error) {

	return Credit{
		Value	:  	value,
		Status	:	status,
		Debenture:	debenture,
		EndDate	:	endDate,
		StartDate: startDate,
	}, nil
}