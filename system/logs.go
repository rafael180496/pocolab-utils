package system

import (
	"fmt"
	"log"
	"time"

	"github.com/rafael180496/pocolab-utils/casting"
	"github.com/rafael180496/pocolab-utils/utils"
)

type (

	/*StLog : Estructura para crear log personalizados por medio de la fecha
	directorio/fecha.log
	*/
	StLog struct {
		//Directorio contenido
		Dir string
		//Nombre del log mas fecha
		Name string
		//Fecha
		Fe time.Time
		//Prefijo
		Prefix string
	}
	//StLogMa : Estructura para crear logs de errores y warnings e info
	StLogMa struct {
		War  *StLog
		Err  *StLog
		Info *StLog
	}
)

/*Printf : Ingresa un texto en los logs asignado. */
func (p *StLog) Printf(format string, args ...interface{}) error {
	err := p.Init()
	if err != nil {
		return err
	}
	log.Printf(format, args...)
	return nil
}

/*ValidF :  valida la fecha del log cargado*/
func (p *StLog) ValidF() {
	if !utils.DateIdent(p.Fe, time.Now()) {
		p.Fe = time.Now()
	}
}

// ReloadNow : Recarga el log actual para que se pueda usar
func (p *StLog) ReloadNow() error {
	p.ValidF()
	return p.Init()
}

// ReloadNow : Recarga el log actual para que se pueda usar
func (p *StLog) Reload(date time.Time) error {
	p.Fe = time.Now()
	return p.Init()
}

/*Init : Inicializa el log para comenzarlo a usar */
func (p *StLog) Init() error {
	NameArch := fmt.Sprintf("%s/%s%s.log", utils.Trim(p.Dir), utils.Trim(p.Name), utils.Trim(casting.TimetoStr(p.Fe)))
	if !FileExist(p.Dir, true) {
		errCreate := DirNew(p.Dir)
		if errCreate != nil {
			return errCreate
		}
	}
	log.SetPrefix(p.Prefix)
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	file, err := FileNew(NameArch)
	if err != nil {
		return err
	}
	log.SetOutput(file)
	return nil
}

/*NewStLog : Crea una nueva intancia de StLog */
func NewStLog(dir, name, prefix string, fe time.Time) (StLog, error) {
	LogNew := StLog{
		Dir:    dir,
		Name:   name,
		Prefix: prefix,
		Fe:     fe,
	}
	err := LogNew.Init()
	if err != nil {
		return LogNew, err
	}

	return LogNew, nil
}

// NewStLogMa : Crea una nueva intancia de StLogMa
func NewStLogMa(dir, name string, fe time.Time) (StLogMa, error) {
	var LogNew StLogMa
	tpErr, err := NewStLog(dir, name, "ERROR", fe)
	if err != nil {
		return LogNew, err
	}
	LogNew.Err = &tpErr

	tpWar, err := NewStLog(dir, name, "WARNING", fe)
	if err != nil {
		return LogNew, err
	}
	LogNew.War = &tpWar

	tpInf, err := NewStLog(dir, name, "INFO", fe)
	if err != nil {
		return LogNew, err
	}
	LogNew.Info = &tpInf
	return LogNew, nil
}

/*Init : Inicializa el log para comenzarlo a usar */
func (p *StLogMa) Init() error {
	err := p.Err.Init()
	if err != nil {
		return err
	}
	err = p.War.Init()
	if err != nil {
		return err
	}
	err = p.Info.Init()
	if err != nil {
		return err
	}
	return nil
}

// ReloadNow : Recarga el log actual para que se pueda usar
func (p *StLogMa) ReloadNow() error {
	err := p.Err.ReloadNow()
	if err != nil {
		return err
	}
	err = p.War.ReloadNow()
	if err != nil {
		return err
	}
	err = p.Info.ReloadNow()
	if err != nil {
		return err
	}
	return nil
}

// ReloadNow : Recarga el log actual para que se pueda usar
func (p *StLogMa) Reload(date time.Time) error {
	err := p.Err.Reload(date)
	if err != nil {
		return err
	}
	err = p.War.Reload(date)
	if err != nil {
		return err
	}
	err = p.Info.Reload(date)
	if err != nil {
		return err
	}
	return nil
}
