package viewmodels

import valueObjects "jwemanager/pkg/domain/value_objects"

type EncryptViewModel struct {
	UserID string                 `json:"user_id" validate:"required,uuid4"`
	KeyID  string                 `json:"key_id" validate:"required,uuid4"`
	Data   map[string]interface{} `json:"data" validate:"required"`
}
type EncryptedViewModel struct {
	EncryptedData string `json:"encrypted_data"`
}

func (pst EncryptViewModel) ToValueObject() valueObjects.EncryptValueObject {
	return valueObjects.EncryptValueObject{
		UserID: pst.UserID,
		KeyID:  pst.KeyID,
		Data:   pst.Data,
	}
}

func NewEncryptedViewModel(vo valueObjects.EncryptedValueObject) EncryptedViewModel {
	return EncryptedViewModel{
		EncryptedData: vo.EncryptedData,
	}
}
