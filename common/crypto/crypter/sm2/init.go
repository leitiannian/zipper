package sm2

import "github.com/zipper-project/zipper/common/crypto/crypter"

func init() {
	tcrypter := &Crypter{}
	crypter.RegisterCrypter(tcrypter.Name(), tcrypter)
}
