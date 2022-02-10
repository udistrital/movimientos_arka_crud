package main

import (
	"fmt"
	"io/ioutil"
	"strings"


	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type FuncionalidadTraslados_20211118_122319 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &FuncionalidadTraslados_20211118_122319{}
	m.Created = "20211118_122319"

	migration.Register("FuncionalidadTraslados_20211118_122319", m)
}

// Run the migrations
func (m *FuncionalidadTraslados_20211118_122319) Up() {
	file, err := ioutil.ReadFile("../scripts/20211118_122319_funcionalidad_traslados_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

	// use m.SQL("CREATE TABLE ...") to make schema update
	
}

// Reverse the migrations
func (m *FuncionalidadTraslados_20211118_122319) Down() {
	file, err := ioutil.ReadFile("../scripts/20211118_122319_funcionalidad_traslados_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}


	// use m.SQL("DROP TABLE ...") to reverse schema update

}
