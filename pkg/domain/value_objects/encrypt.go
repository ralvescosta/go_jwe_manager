package valueObjects

type EncryptValueObject struct {
	UserID string
	KeyID  string
	Data   map[string]interface{}
}

type EncryptedValueObject struct {
	EncryptedData string
}
