package test

import (
	"fmt"
	"testing"

	"github.com/rafael180496/pocolab-utils/encrip"
)

/*TestAes : prueba las funciones de encriptamiento aes */
func TestAes(t *testing.T) {
	key := "abc123"
	text := "hola mundo"
	t.Logf("Key:%s,Text:%s", key, text)
	bloque, err := encrip.EncripAES(key, text)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("BloqueEncrip:%s", bloque)
	bloque, err = encrip.DesencripAES(key, bloque)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("BloqueDesencrip:%s", bloque)
}

/*TestMd5 : Genera un hash Md5 */
func TestMd5(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquemd5:%s", encrip.GeneredHashMd5(key))
}

/*Test256 : Genera un hash sha256 */
func Test256(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloque256:%s", encrip.GeneredHashSha256(key))
}

/*TestToken : Genera un token con hash sha 256 unico */
func TestToken(t *testing.T) {
	key := "abc123"
	t.Logf("key:%s", key)
	t.Logf("bloquetoken:%s", encrip.GenToken(key))
}

/*TestUUID : Genera una clave unica*/
func TestUUID(t *testing.T) {
	text := encrip.GeneredUUID()
	t.Logf("text:[%s]", text)
}

/*TestEncryptCBC : Encriptacion AES CBC*/
func TestEncryptCBC(t *testing.T) {
	OrigData := "user=PEDRO|pass=super.123|correo=test@gmail.com"
	Key := "dde4b1f8a9e6b815"
	init := "dde4b1f8a9e6b815"
	encrypted := encrip.EncryptCBC(OrigData, Key, init)
	fmt.Printf("\nCiphertext(base64):%s", encrypted)
	decrypted := encrip.DecryptCBC(encrypted, Key, init)
	fmt.Printf("\nDecryption result:%s", string(decrypted))
}
