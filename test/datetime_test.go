package test

import (
	"testing"
	"time"

	"github.com/rafael180496/pocolab-utils/casting"
	"github.com/rafael180496/pocolab-utils/utils"
)

/*TestFNulo : Envie una fecha nula.*/
func TestFActual(t *testing.T) {
	t.Logf("Factual:%s", utils.FActual())
}

/*TestFNulo : Envie una fecha nula.*/
func TestFNulo(t *testing.T) {
	t.Logf("NULO:%s", utils.FNulo())
}

/*TestTimetoString : Convierte un date en string */
func TestTimetoString(t *testing.T) {
	tiempo := time.Now()
	t.Logf("Tiempo:%s", tiempo)
	t.Logf("TiempoStr:%s", casting.TimetoStr(tiempo))
}
