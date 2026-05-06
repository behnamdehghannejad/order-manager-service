package utility

import "errors"

func ValidateOrderInput(userID int64, amount float64, status string) error {
	if userID <= 0 {
		return errors.New("user_id must be greater than 0")
	}

	if amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if status == "" {
		return errors.New("status is required")
	}

	return nil
}
