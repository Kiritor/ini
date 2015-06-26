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
			     if ini.ValidateSectionUniq(selection) == true {
					 ini.dicts = append(ini.dicts,data)
				 }
		}
	}

	defer file.Close()
	return ini
}

//getValue by section and key
func (ini *Ini) GetValue(section,key string) string {

	data :=ini.dicts
	for _,v := range data {
		for k,value := range v {
			if k == section {
				for j,m:=range value {
					if j== key {
						return m
					}
				}
				return "The section:"+section +" is not exist key:"+key
			}
		}
	}

	return "The section is not exist!"

}

//validate the section is uniq,to append or not do
func (ini *Ini) ValidateSectionUniq(section string) bool {
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

//delete key-value by section and key
func (ini *Ini) DeleteValue (section,key string) bool {
	data :=ini.dicts
	for i,v :=range data {
		for k,_ :=range v {
			if k == section {
				delete(ini.dicts[i][k],key)
				return true
			}
		}
	}
	return false
}

//set value,if not section or key add,else value change

func (ini *Ini) SetValue(section,key,value string) bool {
	data :=ini.dicts
	var flag bool
	var keyFlag bool
	var index = make(map[int]bool)
	var conf = make(Dict)
	//
	for i,v := range data {
		_,flag := v[section]
		index[i] = flag
		if flag == true {
			//查看该section下是否存在该key
			_,keyBool := data[i][section][key]
			keyFlag = keyBool
			break
		}
	}
	i,flag := func (m map[int]bool) (i int,v bool) {
		for i,v:=range m {
			if v == true {
				return i,true
			}
		}
		return 0,false
	}(index)

	//if section and key exist
	if flag && keyFlag{
		ini.dicts[i][section][key] = value
		return true
	}else if flag &&!keyFlag{
		//section  exsit,key not exsit,add key-value beyond the section
		ini.dicts[i][section][key] = value
		return true

	}else {
		//直接添加section 和key
		conf[section] = make(map[string]string)
		conf[section][key] = value
		ini.dicts = append(ini.dicts,conf)
		return true
	}
	return false
}

func main(){
	ini:=NewIni("config.ini")
	fmt.Println(ini.GetValue("database","username"))
	fmt.Println(ini.dicts)
	ini.DeleteValue("database","username")
	fmt.Println(ini.dicts)
	ini.SetValue("databases","passwords","lcoress")
	fmt.Println(ini.dicts)
}



