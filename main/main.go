package main

import (
	"ini"
	"fmt"
)

func main(){
	iniC:=ini.NewIni("config.ini")
	fmt.Println(iniC)
	fmt.Println(iniC.GetValue("database","username"))
	fmt.Println(iniC.DictList())
/*
	fmt.Println(iniC.dicts)
	iniC.DeleteValue("database","username")
	fmt.Println(iniC.dicts)
	iniC.SetValue("databases","passwords","lcoress")
	fmt.Println(iniC.dicts)
*/
}
