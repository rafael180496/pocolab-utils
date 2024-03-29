package system

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/rafael180496/pocolab-utils/utils"

	color "github.com/fatih/color"
)

type (
	/*Pc : es un tipo para las paletas de colores.*/
	Pc string
)

const (
	ShellToUse = "bash"

	/*colores disponibles */

	/*Green : verde */
	Green Pc = "g"
	/*Red : rojo */
	Red Pc = "r"
	/*Blue : azul */
	Blue Pc = "b"
	/*Cyan : celeste */
	Cyan Pc = "c"
	/*White : blanco */
	White Pc = "w"
	/*Black : negro */
	Black Pc = "bl"
	/*Yellow : amarillo*/
	Yellow Pc = "y"
	/*Magenta : magenta */
	Magenta Pc = "m"
	/*HBlack : negro fuerte */
	HBlack Pc = "hbl"
	/*HRed : rojo fuerte */
	HRed Pc = "hr"
	/*HGreen : verde fuerte */
	HGreen Pc = "hg"
	/*HYellow : amarrillo fuerte */
	HYellow Pc = "hy"
	/*HBlue : azul fuerte */
	HBlue Pc = "hb"
	/*HMagenta : magenta fuerte*/
	HMagenta Pc = "hm"
	/*HCyan : celeste fuerte */
	HCyan Pc = "hc"
	/*HWhite : blanco fuerte */
	HWhite Pc = "hw"
	/*FORMFE : Formato de fecha para los archivo YYYYMMDD*/
	FORMFE = "%d%02d%02d"
)

/*MaxCPUtaskLimit : maxima cantidad de core a usar  con limitaciones TASKCORE es el limite*/
func MaxCPUtaskLimit(TASKCORE int) int {
	TASKCORE = utils.ReturnIf(TASKCORE <= 0, 1, TASKCORE).(int)
	core := MaxCPUtask()
	return utils.ReturnIf(TASKCORE > core, core, TASKCORE).(int)
}

/*MaxCPUtask : maxima multi hilos que puede obtener   */
func MaxCPUtask() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	return utils.ReturnIf(maxProcs < numCPU, maxProcs, numCPU).(int)
}

/*NumProcSet : setea los numeros de proceso validos */
func NumProcSet(n int) int {
	return runtime.GOMAXPROCS(n)
}

/*GetCPU : envia la cantidad de CPU disponible en el procesador */
func GetCPU() int {
	n := runtime.NumCPU()
	if n <= 0 {
		return 1
	}
	return n
}

/*
MsjPc : envia un string de color personalizado de las
constantes disponible en la libreria
*/
func MsjPc(c Pc, format string, arg ...interface{}) string {
	var (
		menj strings.Builder
	)
	if IsLinux() || IsWindows() {
		d := color.New(sendColor(c), color.Bold)
		d.Fprintf(&menj, format, arg...)
	} else {
		fmt.Fprintf(&menj, format, arg...)
	}
	return menj.String()
}

/*
PrintPc : muestra un printf con color personalizado para consolas
basadas en linux
*/
func PrintPc(c Pc, format string, arg ...interface{}) {
	if IsLinux() || IsWindows() {
		d := color.New(sendColor(c), color.Bold)
		d.Printf(format, arg...)
	} else {
		fmt.Printf(format, arg...)
	}
}

/*
MsjBlue : Enviar un string de color celeste funciona para consolas
basadas en linux
*/
func MsjBlue(format string, arg ...interface{}) string {
	return MsjPc(Blue, format, arg...)
}

/*
MsjRed : Enviar un string de color rojo para consolas
basadas en linux
*/
func MsjRed(format string, arg ...interface{}) string {
	return MsjPc(Red, format, arg...)
}

/*
MsjGreen : Enviar un string de color verde para consolas
basadas en linux
*/
func MsjGreen(format string, arg ...interface{}) string {
	return MsjPc(Green, format, arg...)
}

/*
PrintGreen : muestra un printf con color verde para consolas
basadas en linux
*/
func PrintGreen(format string, arg ...interface{}) {
	PrintPc(Green, format, arg...)
}

/*
PrintRed : muestra un printf con color rojo para consolas
basadas en linux
*/
func PrintRed(format string, arg ...interface{}) {
	PrintPc(Red, format, arg...)
}

/*sendColor : envia el color correcto en atributo */
func sendColor(item Pc) color.Attribute {
	switch item {
	/*Green : verde */
	case Green:
		return color.FgGreen
		/*Red : rojo */
	case Red:
		return color.FgRed
		/*Blue : azul */
	case Blue:
		return color.FgBlue
		/*Cyan : celeste */
	case Cyan:
		return color.FgCyan
		/*White : blanco */
	case White:
		return color.FgWhite
		/*Black : negro */
	case Black:
		return color.FgBlack
		/*Yellow : amarillo*/
	case Yellow:
		return color.FgYellow
		/*Magenta : magenta */
	case Magenta:
		return color.FgMagenta
		/*HiBlack : negro fuerte */
	case HBlack:
		return color.FgHiBlack
		/*HRed : rojo fuerte */
	case HRed:
		return color.FgHiRed
		/*HGreen : verde fuerte */
	case HGreen:
		return color.FgHiGreen
		/*HYellow : amarrillo fuerte */
	case HYellow:
		return color.FgHiYellow
		/*HBlue : azul fuerte */
	case HBlue:
		return color.FgHiBlue
		/*HMagenta : magenta fuerte*/
	case HMagenta:
		return color.FgHiMagenta
		/*HCyan : celeste fuerte */
	case HCyan:
		return color.FgHiCyan
		/*HWhite : blanco fuerte */
	case HWhite:
		return color.FgHiWhite
	default:
		return color.FgWhite

	}

}

/*GetLocalIPV4 : te envia la ip local que contiene la maquina que estas ejecutando el programa */
func GetLocalIPV4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

/*IsMac : Valida si estas en un sistema operativo Imac */
func IsMac() bool {
	return IsSO("darwin")
}

/*IsWindows : Valida si estas en un sistema operativo windows */
func IsWindows() bool {
	return IsSO("windows")
}

/*IsLinux : Valida si estas en un sistema operativo linux */
func IsLinux() bool {
	return IsSO("linux")
}

/*
IsSO : valida en que sistema operativo estas lista:
android,
darwin,
dragonfly,
linux,
freebsd,
openbsd,
solaris,
netbsd,
plan9,
windows
*/
func IsSO(so string) bool {
	return utils.ReturnIf(runtime.GOOS == so, true, false).(bool)
}

/*
ExecuteBashLinux : ejecuta un comando en el bash de linux devolviendo primero
salida de la ejecucion
salida si da error al ejecucion
error nativo de golang
*/
func ExecuteBashLinux(command string) (string, string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func ExecutePrintLinux(command string) {
	out, outErr, err := ExecuteBashLinux(command)
	if !utils.IsEmptyStr(outErr) {
		PrintGreen("%s\n", outErr)
	}
	if !utils.IsEmptyStr(out) {
		PrintGreen("%s\n", out)
	}
	if err != nil {
		if err.Error() != "exit status 1" {
			PrintRed("%s\n", err.Error())
		}
	}
}
