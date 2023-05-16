package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rafael180496/pocolab-utils/casting"
	"github.com/rafael180496/pocolab-utils/utils"
)

/*SubString : recorta una cadena */
func TestSubString(t *testing.T) {
	fmt.Printf("%s", utils.SubString("18:20", 0, 2))
	fmt.Printf("%s", utils.SubString("18:20", 3, 5))
}

/*TestReturnIf : retorna con un condicional */
func TestReturnIf(t *testing.T) {
	t.Logf("%s", utils.ReturnIf(5 > 4, "It's true", "It's false"))
}

/*TestTrim : Quita los espacio de un texto */
func TestTrim(t *testing.T) {
	text := "Hola Mundo TDA"
	t.Logf("text:[%s]", text)
	t.Logf("trim:[%s]", utils.Trim(text))
}

/*TestFloat64 : Redondea un valor float a x decimales*/
func TestFloat64(t *testing.T) {
	valor := 12.34661
	t.Logf("valor:[%f]", valor)
	valor = casting.RoundFloat64(valor, 2)
	t.Logf("valor:[%f]", valor)
}

/*Test_Merge_string : combina unas cadena de caracteres*/
func Test_Merge_string(t *testing.T) {
	a := ""
	b := "foo"
	expected := "foo"
	ret := utils.Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

/*Test_Merge_Array : combina los array de un mismo tipo de datos*/
func Test_Merge_Array(t *testing.T) {
	a := []int{10, 11, 12}
	b := []int{}
	expected := []int{10, 11, 12}
	ret := utils.Merge(a, b)

	if !reflect.DeepEqual(ret, expected) {
		t.Errorf("Actual ( %#v ) does not match expected ( %#v )", ret, expected)
	}
}

/*TestEncodeDecode :  tranforma base64 los string  y los decodifica*/
func TestEncodeDecode(t *testing.T) {
	code := casting.EncodeBase64("prueba")
	fmt.Printf("Encode:[%s]\n", code)
	encode := casting.DecodeBase64(code)
	fmt.Printf("Decode:[%s]\n", encode)
}
