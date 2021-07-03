package storage

import (
	"encoding/json"
	"github.com/mimose/gcosy/lib"
	"mo-for-desktop/model/errs"
)

func GenerateToStorageData(v interface{}, cipherBuilder lib.CipherBuilder) ([]byte, error) {
	vJson, marshalErr := json.Marshal(v)
	if marshalErr != nil {
		return nil, lib.NewError(errs.Marshal, errs.MarshalDesc, marshalErr)
	}
	encrypted, encryptedErr := cipherBuilder.AesEncrypt(vJson)
	if encryptedErr != nil {
		return nil, lib.NewError(errs.Encrypt, errs.EncryptDesc, encryptedErr)
	}
	return encrypted, nil
}
