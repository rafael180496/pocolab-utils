package test

import (
	"fmt"
	"testing"

	"github.com/rafael180496/pocolab-utils/utils"
)

/*TestSendError : Envia un error con mensaje */
func TestSendtTrycatch(t *testing.T) {
	fmt.Println("Comienza")
	utils.Block{
		Try: func() {
			fmt.Println("Entra al Try")
			utils.Throw("Error")
		},
		Catch: func(e utils.Exception) {
			fmt.Printf("Error a capturar %v\n", e)
		},
		Finally: func() {
			fmt.Println("Finalizacion...")
		},
	}.Do()
	fmt.Println("Termina")
}
