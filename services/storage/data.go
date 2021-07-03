package storage

import (
	"encoding/json"
	"github.com/mimose/gcosy/lib"
	. "mo-for-desktop/services/errs"
)

func GenerateToStorageData(v interface{}, cipherBuilder lib.CipherBuilder) ([]byte, error) {
	vJson, marshalErr := json.Marshal(v)
	if marshalErr != nil {
		return nil, lib.NewError(Marshal, MarshalDesc, marshalErr)
	}
	encrypted, encryptedErr := cipherBuilder.AesEncrypt(vJson)
	if encryptedErr != nil {
		return nil, lib.NewError(Encrypt, EncryptDesc, encryptedErr)
	}
	return encrypted, nil
}
