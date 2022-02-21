package valueObjects

type DecryptValueObject struct {
	UserID        string
	KeyID         string
	EncryptedData []byte
}

type DecryptedValueObject struct {
	Data map[string]interface{}
}
