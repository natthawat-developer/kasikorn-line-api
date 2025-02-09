package utils

func MaskDebitCardNumber(cardNumber *string) string {

	if cardNumber == nil || len(*cardNumber) != 19 {
		return ""
	}

	return (*cardNumber)[:7] + "•• •••• " + (*cardNumber)[15:]
}
