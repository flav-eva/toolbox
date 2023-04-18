package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

/*
PBKDF2(Password-Based Key Derivation Function)是一个用来导出密钥的函数，常用于生成加密的密码。
它的基本原理是通过一个伪随机函数（例如HMAC函数），把明文和一个盐值作为输入参数，然后重复进行运算，
并最终产生密钥。 如果重复的次数足够大，破解的成本就会变得很高

官方提供 “golang.org/x/crypto/pbkdf2″ 包封装此加密算法，形式如下：

func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	...
}
1. password：用户密码，加密的原材料，以 []byte 传入；
2. salt：随机盐，以 []byte 传入；
3. iter：迭代次数，次数越多加密与破解所需时间越长，
4. keyLen：期望得到的密文长度；
5. Hash：加密所用的 hash 函数，默认使用为 sha1，请使用sha256

随机盐生成
Go语言中有两个包提供rand，分别为 “math/rand” 和 “crypto/rand”, 对应两种应用场景：
* “math/rand” 包实现了伪随机数生成器，也就是生成整形和浮点型。
* “crypto/rand” 包实现了用于加解密的更安全的随机数生成器，这里使用的就是此种。


这里用于用户密码的加密和验证

*/

var (
	ErrUnknownKDF = errors.New("unknown KDF")
)

type KDFType string

const (
	KDF_PBKDF2 KDFType = "pbkdf2"

	PBKDF2_DEFAULT_ITERATION      = 20480
	PBKDF2_DEFAULT_SALT_BYTE_SIZE = 16
	PBKDF2_DEFAULT_KEY_BYTE_SIZE  = 64
)

type KeyHeader struct {
	KDF KDFType `json:"kdf"`
}

type pbkdf2Key struct {
	KeyHeader
	Iteration int    `json:"iteration"`
	Salt      []byte `json:"salt"`
	Key       []byte `json:"key"`
}

func HashPassword(password string) ([]byte, error) {
	key := pbkdf2Key{
		KeyHeader: KeyHeader{KDF_PBKDF2},
		Iteration: PBKDF2_DEFAULT_ITERATION,
		Salt:      make([]byte, PBKDF2_DEFAULT_SALT_BYTE_SIZE),
	}
	if _, err := rand.Read(key.Salt); err != nil {
		return nil, err
	}
	key.Key = pbkdf2.Key([]byte(password), key.Salt, key.Iteration, PBKDF2_DEFAULT_KEY_BYTE_SIZE, sha256.New)
	return json.Marshal(&key)
}

func VerifyPassword(password string, hashPassword []byte) (bool, error) {
	var key pbkdf2Key
	if err := json.Unmarshal(hashPassword, &key); err != nil {
		return false, err
	}
	if key.KDF != KDF_PBKDF2 {
		return false, ErrUnknownKDF
	}

	return bytes.Equal(key.Key, pbkdf2.Key([]byte(password), key.Salt, key.Iteration, len(key.Key), sha256.New)), nil
}
