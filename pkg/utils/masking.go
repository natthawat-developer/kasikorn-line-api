package utils

// MaskAccountNumber แปลงเลขบัญชีให้แสดงบางส่วน เช่น 56••••61
func MaskAccountNumber(accountNumber *string) *string {
	if accountNumber == nil || len(*accountNumber) < 6 {
		masked := "••••"
		return &masked
	}
	masked := (*accountNumber)[:2] + "••••" + (*accountNumber)[len(*accountNumber)-2:]
	return &masked
}




func MaskDebitCardNumber(cardNumber *string) string {
	// Check if cardNumber is nil or not exactly 19 characters (including spaces)
	if cardNumber == nil || len(*cardNumber) != 19 {
		return "" // Return empty string or handle error as needed
	}

	// Mask the middle digits
	return (*cardNumber)[:7] + "•• •••• " + (*cardNumber)[15:]
}
