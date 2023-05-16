package test

import (
	"testing"

	"github.com/rafael180496/pocolab-utils/system"
)

/*TestMsjBlue : Envie un string en color celeste*/
func TestMsjBlue(t *testing.T) {
	t.Logf("Texto:%s", system.MsjBlue("Hola"))
}

/*TestMsjRed : Envie un string en color rojo*/
func TestMsjRed(t *testing.T) {
	t.Logf("Texto:%s", system.MsjRed("Hola"))
}

/*TestMsjGreen : Envie un string en color verde*/
func TestMsjGreen(t *testing.T) {
	t.Logf("Texto:%s\n", system.MsjGreen("Hola"))
	system.PrintPc(system.Green, "Texto:Hola\n")
}

/*TestMsjGreen : Envie un string en color verde*/
func TestIP4(t *testing.T) {
	t.Logf("IPLOCAL:%s\n", system.GetLocalIPV4())
}

/*TestMsjPc : prueba todos los texto disponible */
func TestMsjPc(t *testing.T) {
	t.Logf("Texto:%s", system.MsjPc(system.Green, "%s", system.Green))
	t.Logf("Texto:%s", system.MsjPc(system.Red, "%s", system.Red))
	t.Logf("Texto:%s", system.MsjPc(system.Blue, "%s", system.Blue))
	t.Logf("Texto:%s", system.MsjPc(system.Cyan, "%s", system.Cyan))
	t.Logf("Texto:%s", system.MsjPc(system.White, "%s", system.White))
	t.Logf("Texto:%s", system.MsjPc(system.Black, "%s", system.Black))
	t.Logf("Texto:%s", system.MsjPc(system.Yellow, "%s", system.Yellow))
	t.Logf("Texto:%s", system.MsjPc(system.Magenta, "%s", system.Magenta))
	t.Logf("Texto:%s", system.MsjPc(system.HBlack, "%s", system.HBlack))
	t.Logf("Texto:%s", system.MsjPc(system.HRed, "%s", system.HRed))
	t.Logf("Texto:%s", system.MsjPc(system.HGreen, "%s", system.HGreen))
	t.Logf("Texto:%s", system.MsjPc(system.HYellow, "%s", system.HYellow))
	t.Logf("Texto:%s", system.MsjPc(system.HBlue, "%s", system.HBlue))
	t.Logf("Texto:%s", system.MsjPc(system.HMagenta, "%s", system.HMagenta))
	t.Logf("Texto:%s", system.MsjPc(system.HCyan, "%s", system.HCyan))
}
