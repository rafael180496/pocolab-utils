package utils

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

// validateSQLInjection : Valida un string si no tiene palabras reservadas sql
func ValidateSQLInjection(input string) bool {
	// Lista de palabras clave que podrían ser utilizadas en una inyección SQL
	keywords := []string{
		"--",
		"/*",
		"*/",
		"@@",
		"char",
		"nchar",
		"varchar",
		"nvarchar",
		"alter",
		"begin",
		"cast",
		"create",
		"cursor",
		"declare",
		"delete",
		"drop",
		"end",
		"exec",
		"execute",
		"fetch",
		"insert",
		"kill",
		"open",
		"select",
		"sys",
		"sysobjects",
		"syscolumns",
		"table",
		"update",
	}
	// Convertir la entrada a minúsculas para la comparación
	inputLower := strings.ToLower(input)
	// Comprobar si la entrada contiene alguna de las palabras clave
	for _, keyword := range keywords {
		if strings.Contains(inputLower, keyword) {
			return false
		}
	}

	return true
}

/*EnvsLoad : carga un arreglo con variables de entorno y lo regresar enun map[string]string*/
func EnvsLoad(envs ...string) (map[string]string, error) {
	var m = map[string]string{}
	if len(envs) == 0 {
		return nil, fmt.Errorf("empty variable array")
	}
	godotenv.Load(".env")
	for _, v := range envs {
		m[v] = os.Getenv(v)
	}
	return m, nil
}

/*CreateEnvs : genera un .env solo con un map string*/
func CreateEnvs(data map[string]string) error {
	err := godotenv.Write(data, ".env")
	if err != nil {
		return err
	}
	return nil
}

/*Reverse :  manda la reversa de un string */
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

/*IsNilStr : valida si un string esta vacio*/
func IsNilStr(str string) bool {
	return ReturnIf(Trim(str) == "", false, true).(bool)
}

/*GetLines : consigue un arreglos por saltos de lineas en un string*/
func GetLines(s string) []string {
	return strings.Split(s, "\n")
}

/*ValidDuplidArrayStr : valida un arreglo de string si estan duplicados*/
func ValidDuplidArrayStr(strs []string) bool {
	for i, str := range strs {
		for j, straux := range strs {
			if i != j && straux == str {
				return false
			}
		}
	}
	return true
}
func accStr(str, acc string) string {
	return ReturnIf(acc == "u", strings.ToUpper(acc), strings.ToLower(str)).(string)
}
func accStrs(acc string, strs ...string) []string {
	var strsNew []string
	for _, str := range strs {
		strsNew = append(strsNew, accStr(str, acc))
	}
	return strsNew
}

/*IsNilArrayStr : valida si los string son vacio*/
func IsNilArrayStr(strs ...string) bool {
	for _, str := range strs {
		if !IsNilStr(str) {
			return false
		}
	}
	return true
}

/*UpperStrs : coloca en mayusculas un arreglo compleot de strings*/
func UpperStrs(strs ...string) []string {
	return accStrs("u", strs...)
}

/*LowerStrs : coloca en minusculas un arreglo compleot de strings*/
func LowerStrs(strs ...string) []string {
	return accStrs("l", strs...)
}

/*InStr : compara  varios string con uno especifico*/
func InStr(str string, strs ...string) bool {
	for _, item := range strs {
		if str == item {
			return true
		}
	}
	return false
}

/*
FilterExcl : excluye los registro de un arreglo A  con el B  ejemplo
A[a,b,c,d]
B[a,b]
Result
[c,d]
*/
func FilterExcl(strs []string, excl []string) []string {
	var ret []string
	for _, item := range strs {
		ind := true
		for _, item2 := range excl {
			if item2 == item {
				ind = false
			}
		}
		if ind {
			ret = append(ret, item)
		}
	}
	return ret
}

/*FilterStr : filtra un arreglo de string mediante un metodo definido */
func FilterStr(strs []string, valid func(string) bool) (ret []string) {
	for _, s := range strs {
		if valid(s) {
			ret = append(ret, s)
		}
	}
	return
}

/*
ReturnIf : retorna un if ternario
https://github.com/TheMushrr00m/go-ternary Doc
ReturnIf(<bool expression>, <result for true>, <result for false>)
ReturnIf(5 > 4, "It's true", "It's false :(")
*/
func ReturnIf(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
}

/*StrRand : genera una cadena de caracteres ramdon*/
func StrRand(cant int, Upper bool) string {
	cant = ReturnIf(cant <= 0, 1, cant).(int)
	str := ""
	for i := 0; i < cant; i++ {
		str += CharRand(Upper)
	}
	return str
}

/*CharRand : Genera una letra aleatoria upper indica si queres mayusculas o miniscula.*/
func CharRand(Upper bool) string {
	return ReturnIf(Upper, string(byte(RandInt(65, 90))), string(byte(RandInt(97, 122)))).(string)
}

/*RandInt : envia un numero aleatorio*/
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()
}

/*SubString : substring para un string en golang con runas*/
func SubString(s string, start, end int) string {
	start_str_idx := 0
	i := 0
	for j := range s {
		if i == start {
			start_str_idx = j
		}
		if i == end {
			return s[start_str_idx:j]
		}
		i++
	}
	return s[start_str_idx:]
}

/*Trim : Elimina el espacio de un string a nivel de runas.*/
func Trim(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

/*IsSpace : valida si la cadena contiene espacio. */
func IsSpace(str string) bool {
	for _, v := range str {
		if unicode.IsSpace(v) {
			return true
		}
	}
	return false
}

/*IsEmptyVl : valida si el valor esta vacio*/
func IsEmptyVl(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

/*Merge : combina variables array o string en un solo valor conjunto*/
func Merge(base, override interface{}) interface{} {
	b := reflect.ValueOf(base)
	o := reflect.ValueOf(override)
	ret := mergeBase(b, o)
	return ret.Interface()
}

func mergeBase(base, override reflect.Value) reflect.Value {
	var result reflect.Value
	switch base.Kind() {
	case reflect.Ptr:
		switch base.Elem().Kind() {
		case reflect.Ptr:
			fallthrough
		case reflect.Interface:
			fallthrough
		case reflect.Struct:
			fallthrough
		case reflect.Map:
			if base.IsNil() {
				result = override
			} else if override.IsNil() {
				result = base
			} else {
				result = mergeBase(base.Elem(), override.Elem())
			}
		default:
			if IsEmptyVl(override) {
				result = base
			} else {
				result = override
			}
		}
	case reflect.Interface:
		result = mergeBase(base.Elem(), override.Elem())
	case reflect.Struct:
		result = reflect.New(base.Type())
		for i, n := 0, base.NumField(); i < n; i++ {
			if result.Elem().Field(i).CanSet() {
				vl := mergeBase(base.Field(i), override.Field(i))
				if result.Elem().Field(i).CanSet() && vl.IsValid() {
					if vl.Kind() == reflect.Ptr && result.Elem().Field(i).Kind() != reflect.Ptr {
						vl = vl.Elem()
					} else if result.Elem().Field(i).Kind() == reflect.Ptr && vl.Kind() != reflect.Ptr && vl.CanAddr() {
						vl = vl.Addr()
					}
					result.Elem().Field(i).Set(vl)
				}
			}
		}

	case reflect.Map:
		element := base.Type().Elem().Kind() != reflect.Ptr
		result = reflect.MakeMap(base.Type())
		for _, key := range base.MapKeys() {
			result.SetMapIndex(key, base.MapIndex(key))
		}
		if override.Kind() == reflect.Map {
			for _, key := range override.MapKeys() {
				overrideVal := override.MapIndex(key)
				baseVal := base.MapIndex(key)
				if !overrideVal.IsValid() {
					continue
				}
				if !baseVal.IsValid() {
					result.SetMapIndex(key, overrideVal)
					continue
				}
				vl := mergeBase(baseVal, overrideVal)
				if element && vl.Kind() == reflect.Ptr {
					result.SetMapIndex(key, vl.Elem())

				} else {
					result.SetMapIndex(key, vl)
				}
			}
		}

	default:
		if IsEmptyVl(override) {
			result = base
		} else {
			result = override
		}
	}
	return result
}
