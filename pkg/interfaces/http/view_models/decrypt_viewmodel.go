package viewmodels

import valueObjects "jwemanager/pkg/domain/value_objects"

type DecryptViewModel struct {
	UserID        string `json:"user_id" validate:"required,uuid4"`
	KeyID         string `json:"key_id" validate:"required,uuid4"`
	EncryptedData string `json:"encrypted_data" validate:"required"`
}

type DecryptedViewModel struct {
	Data map[string]interface{} `json:"data"`
}

func (pst DecryptViewModel) ToValueObject() valueObjects.DecryptValueObject {
	return valueObjects.DecryptValueObject{
		UserID:        pst.UserID,
		KeyID:         pst.KeyID,
		EncryptedData: []byte(pst.EncryptedData),
	}
}

func NewDecryptedViewModel(vo valueObjects.DecryptedValueObject) DecryptedViewModel {
	return DecryptedViewModel{
		Data: vo.Data,
	}
}
