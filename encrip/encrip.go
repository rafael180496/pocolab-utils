package encrip

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rafael180496/pocolab-utils/casting"
	"github.com/rafael180496/pocolab-utils/utils"
)

// TokenData : estructura para devolver un token
type TokenInput struct {
	//Autor o empresa
	Signature string
	//timepo a expirar si es menor o igual a 0 no expirara
	Expired        int
	SecretKey      string
	UserData       map[string]interface{}
	AdditionalData map[string]interface{}
}

type TokenOutput struct {
	UserData       map[string]interface{} `json:"userData,omitempty"`
	AdditionalData map[string]interface{} `json:"additionalData,omitempty"`
	jwt.StandardClaims
}

// TokenifyJwt :  crea un token con data jwt
func TokenifyJwt(inputData TokenInput) (string, error) {
	dataJwtStandard := jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Issuer:   inputData.Signature,
	}
	if inputData.Expired > 0 {
		dataJwtStandard.ExpiresAt = time.Now().Add(time.Minute * time.Duration(inputData.Expired)).Unix()
	}
	claims := TokenOutput{
		UserData:       inputData.UserData,
		AdditionalData: inputData.AdditionalData,
		StandardClaims: dataJwtStandard,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(inputData.SecretKey))
}

// DetokenifyJwt : decodifica un token regresando la data correspondiente
func DetokenifyJwt(tokenStr, secretKey string) (tp *TokenOutput, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenOutput{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenOutput); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid-token")
	}
}

/*GeneredHashSha256 : Genera un hash con encriptacion sha256 */
func GeneredHashSha256(key string) string {
	hash := sha256.New()
	io.WriteString(hash, key)
	text := hex.EncodeToString(hash.Sum(nil))
	return text
}

/*GeneredHashMd5 : Genera un hash con encriptacion md5 */
func GeneredHashMd5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/*EncripAES : Encripta en aes 256 un texto donde la key es la llave noce una llave secundaria y el texto a ecriptar */
func EncripAES(key string, text string) (string, error) {
	textByte := casting.StrtoByte(text)
	bloque, err := aes.NewCipher([]byte(GeneredHashMd5(key)))
	if err != nil {
		fmt.Println(err.Error())
		return "", fmt.Errorf("error generating phisher")
	}
	gcm, err := cipher.NewGCM(bloque)
	if err != nil {
		return "", fmt.Errorf("error generating block")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to read the block")
	}
	textoEncrip := gcm.Seal(nonce, nonce, textByte, nil)

	return casting.BytetoStrHex(textoEncrip), nil
}

/*DesencripAES : Desencripta en aes 256 un texto donde la key es la llave noce una llave secundaria y el texto a ecriptar */
func DesencripAES(key string, text string) (string, error) {
	textByte := casting.StrtoByteHex(text)
	keyHash := []byte(GeneredHashMd5(key))
	bloque, err := aes.NewCipher(keyHash)
	if err != nil {
		return "", fmt.Errorf("error generating phisher")
	}
	gcm, err := cipher.NewGCM(bloque)
	if err != nil {
		return "", fmt.Errorf("error generating block")
	}
	nonceSize := gcm.NonceSize()
	if nonceSize > len(textByte) {
		return "", fmt.Errorf("failed to open the block")
	}
	nonce, ciphertext := textByte[:nonceSize], textByte[nonceSize:]
	textoDesencrip, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to open the block")
	}

	return casting.BytetoStr(textoDesencrip), nil
}

/*GenToken : Genera un token dependiendo de un string.*/
func GenToken(str string) string {
	return GeneredHashSha256(utils.StrRand(len(str), false) + string(rune(time.Now().Second())))
}

/*GeneredUUID : Genera un codigo uuid unico */
func GeneredUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

/*
Encriptacion AES AES/CBC/PKCS5Padding

*/
/*EncryptCBC : encrip AES/CBC/PKCS5Padding*/
func EncryptCBC(src, sKey, ivParameter string) string {
	key := []byte(sKey)
	iv := []byte(ivParameter)
	result, err := aes128Encrypt([]byte(src), key, iv)
	if err != nil {
		panic(err)
	}
	return base64.RawStdEncoding.EncodeToString(result)
}

/*DecryptCBC : encrip AES/CBC/PKCS5Padding*/
func DecryptCBC(src, sKey, ivParameter string) string {
	key := []byte(sKey)
	iv := []byte(ivParameter)
	var result []byte
	var err error
	result, err = base64.RawStdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	origData, err := aes128Decrypt(result, key, iv)
	if err != nil {
		panic(err)
	}
	return string(origData)
}
func aes128Encrypt(origData, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}
func aes128Decrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
