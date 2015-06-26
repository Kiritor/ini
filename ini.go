/**
 * Parse .ini config file
 * @Author            LCore
 * @Blog              http://kiritor.github.io
*/


package main

import (
	//"bufio"
	"fmt"
	//"io"
	"os"
	//"strings"
)

type Dict map[string]map[string]string

//ini Object
type Ini struct {
	filePath                     string
	dicts                        []Dict                  //data collection
}

//init ini Object with filePath
func NewIni(filePath string) *Ini {
	ini := new(Ini)
	ini.filePath = filePath
	file,err :=os.Open(ini.filePath)
	if err!= nil {
	    fmt.Println("Error:",err.Error())
		return nil
	}
	defer file.Close()
	return ini
}

