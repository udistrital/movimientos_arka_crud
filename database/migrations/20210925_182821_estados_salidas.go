package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstadosSalidas_20210925_182821 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstadosSalidas_20210925_182821{}
	m.Created = "20210925_182821"

	migration.Register("EstadosSalidas_20210925_182821", m)
}

// Run the migrations
func (m *EstadosSalidas_20210925_182821) Up() {
	file, err := ioutil.ReadFile("../scripts/20210925_182821_estados_salidas_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *EstadosSalidas_20210925_182821) Down() {
	file, err := ioutil.ReadFile("../scripts/20210925_182821_estados_salidas_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
