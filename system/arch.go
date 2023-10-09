package system

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"

	"github.com/rafael180496/pocolab-utils/casting"
	"github.com/rafael180496/pocolab-utils/utils"
)

var (
	/*EXT : extensiones de archivos */
	EXT = map[string]string{
		"JSON": ".json",
		"INI":  ".ini",
		"XML":  ".xml",
		"RUT":  ".rut",
		"SQL":  ".sql",
		"DB":   ".db",
		"CSV":  ".csv",
		"TXT":  ".txt",
		"DBX":  ".dbx",
	}
)

type (
	/*StArchMa : estructura maestra que crea directorios o archivos masivos.*/
	StArchMa []StArch
	/*StArch : estructura para registra los archivos y generarlos.*/
	StArch struct {
		Path   string `json:"file"`
		IndDir bool   `json:"indir"`
	}
)

/*Create : crear el archivo o directorio vacio si existe no lo crea */
func (p *StArchMa) Create() error {
	for _, item := range *p {
		err := item.Create()
		if err != nil {
			return nil
		}
	}
	return nil
}

/*Create : crear el archivo o directorio vacio si existe no lo crea */
func (p *StArch) Create() error {
	if !FileExist(p.Path, p.IndDir) {
		var err error
		if p.IndDir {
			err = DirNew(p.Path)
		} else {
			_, err = FileNew(p.Path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

/*ReadFileStr : lee un archivo de texto y lo pasa a string*/
func ReadFileStr(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	dataStr := casting.BytetoStr(data)
	if !utils.IsNilStr(dataStr) {
		return "", utils.StrErr("el contenido del archivo esta vacio")
	}
	return dataStr, nil
}

/*FileExt : Valida las extenciones de archivos.*/
func FileExt(Path string, ext string) bool {
	if !FileExist(Path, false) {
		return false
	}
	return utils.ReturnIf(strings.Index(Path, EXT[ext]) > 0, true, false).(bool)
}

/*FileRename : Renombra a un archivo como tambien lo puede mover a otro directorio de manera nativa.*/
func FileRename(PathOrigen, PathNuevo string) error {
	err := os.Rename(PathOrigen, PathNuevo)
	if err != nil {
		return err
	}
	return nil
}

/*
FileExist : Valida si el archivo del path existe antes de procesarlo.
Valida tambien si existe un directorio con el inddir en true
*/
func FileExist(Path string, inddir bool) bool {
	if !utils.IsNilStr(Path) {
		return false
	}
	if inddir {
		Path = PlecaAdd(Path)
	}
	info, err := os.Stat(Path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	if inddir {
		return info.IsDir()
	}
	return true
}

/*FileNew : crea un archivo X*/
func FileNew(p string) (*os.File, error) {
	f, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

/*
TrimFile : Renombra a una carpeta o archivo quitandole todos los espacio regresando
el path del nuevo archivo.
*/
func TrimFile(Path string) (string, error) {
	if !FileExist(Path, false) {
		return "", fmt.Errorf("the file does not exist")
	}
	PathOrig := Path
	Path = utils.Trim(strings.Replace(Path, "\r", "", -1))
	err := FileRename(PathOrig, Path)
	if err != nil {
		return "", err
	}
	return Path, nil
}

/*DirNew : Crea una carpeta vacia en el sistema*/
func DirNew(Path string) error {
	err := os.MkdirAll(PlecaAdd(Path), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

/*Open :  abre un archivo X*/
func Open(Path string) (*os.File, error) {
	if !FileExist(Path, false) {
		return nil, fmt.Errorf("the file does not exist")
	}
	fileOrig, err := os.Open(Path)
	return fileOrig, err
}

/*Write : escrive los datos en un archivo X*/
func Write(Path string, data []byte) error {
	f, err := Open(Path)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

/*CpFile : copia un archivo Origen a un directorio destino*/
func CpFile(PathOrig, PathDest string) error {
	PathDest = PlecaAdd(PathDest)
	if !FileExist(PathDest, true) {
		return fmt.Errorf("directory does not exist")
	}
	fileOrig, err := Open(PathOrig)
	if err != nil {
		return err
	}
	infoFile, err := os.Stat(fileOrig.Name())
	if err != nil {
		return err
	}
	pathFinal := PathDest + infoFile.Name()
	fileNew, err := FileNew(pathFinal)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileNew, fileOrig)
	if err != nil {
		return err
	}
	return nil
}

/*CpDir : copia una carpeta entera a una carpeta destino*/
func CpDir(PathOrig, PathDest string) error {
	PathOrig = PlecaAdd(PathOrig)
	PathDest = PlecaAdd(PathDest)
	if !FileExist(PathOrig, true) {
		return fmt.Errorf("directory does not exist")
	}
	if !FileExist(PathDest, true) {
		err := DirNew(PathDest)
		if err != nil {
			return err
		}
	}
	archs, err := ListDir(PathOrig)
	if err != nil {
		return err
	}
	for _, arch := range archs {
		if arch.IsDir() {
			err = CpDir(PathOrig+arch.Name(), PathDest+arch.Name())
			if err != nil {
				return err
			}
		} else {
			err = CpFile(PathOrig+arch.Name(), PathDest)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

/*RmDir : Elimina un directorio entero*/
func RmDir(src string) error {
	src = PlecaAdd(src)
	if !FileExist(src, true) {
		return fmt.Errorf("directory does not exist")
	}

	archs, err := ListDir(src)
	if err != nil {
		return err
	}
	for _, arch := range archs {
		if arch.IsDir() {
			err = RmDir(src + arch.Name())
			if err != nil {
				return err
			}
		} else {
			err = RmFile(src + arch.Name())
			if err != nil {
				return err
			}
		}
	}
	err = RmFile(src)
	if err != nil {
		return err
	}

	return nil
}

/*RmFile : elimina un archivo exacto*/
func RmFile(file string) error {
	if !FileExist(file, false) {
		return fmt.Errorf("the file does not exist")
	}
	err := os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

/*ListDir : lista la infomacion que contiene una carpeta*/
func ListDir(src string) ([]fs.DirEntry, error) {
	src = PlecaAdd(src)
	if !FileExist(src, true) {
		return nil, fmt.Errorf("directory does not exist")
	}
	files, err := os.ReadDir(src)
	if err != nil {
		return nil, err
	}
	return files, nil
}

/*PlecaAdd : coloca la pleca de un directorio "/" */
func PlecaAdd(src string) string {
	if src[len(src)-1] != '/' {
		src = src + "/"
	}
	return src
}
