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
