package utils

func MaskDebitCardNumber(cardNumber *string) string {
	// Check if cardNumber is nil or not exactly 19 characters (including spaces)
	if cardNumber == nil || len(*cardNumber) != 19 {
		return ""
	}

	// Mask the middle digits
	return (*cardNumber)[:7] + "•• •••• " + (*cardNumber)[15:]
}
