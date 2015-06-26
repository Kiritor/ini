/**
 * Parse .ini config file
 * @Author            LCore
 * @Blog              http://kiritor.github.io
*/


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	var data Dict
	var selection string
	buf :=bufio.NewReader(file)
	for{
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err !=nil {
			if err!=io.EOF {
				fmt.Println("Error:",err.Error())
			}
			if len(line) == 0 {
				//readed file And no line to parse
				break
			}
		}

		switch {
			case len(line) == 0:
			case line[0] == '[' && line[len(line)-1]== ']':
			     selection = strings.TrimSpace(line[1:len(line)-1]) //得到section
			     data = make(Dict)
		         data[selection] = make(map[string]string)
			default:
			     //parse key-value
			     i :=strings.IndexAny(line,"=")
			     value :=strings.TrimSpace(line[i+1:len(line)])
			     key := strings.TrimSpace(line[0:i])
			     //set key-value into data
			     data[selection][key] = value
			     if ini.validateSectionUniq(selection) == true {
					 ini.dicts = append(ini.dicts,data)
				 }
		}
	}
	fmt.Println(ini.dicts)
	defer file.Close()
	return ini
}

//validate the section is uniq,to append or add
func (ini *Ini) validateSectionUniq(section string) bool {
	//section is no key
	for _,v :=range ini.dicts {
		for key,_ :=range v {
			if key == section {
				return false
			}
		}
	}
	return true
}

func main(){
	NewIni("config.ini")
}

