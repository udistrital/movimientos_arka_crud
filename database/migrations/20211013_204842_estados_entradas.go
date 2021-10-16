package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstadosEntradas_20211013_204842 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstadosEntradas_20211013_204842{}
	m.Created = "20211013_204842"

	migration.Register("EstadosEntradas_20211013_204842", m)
}

// Run the migrations
func (m *EstadosEntradas_20211013_204842) Up() {
	file, err := ioutil.ReadFile("../scripts/20211013_204842_estados_entradas_up.sql")

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
func (m *EstadosEntradas_20211013_204842) Down() {
	file, err := ioutil.ReadFile("../scripts/20211013_204842_estados_entradas_down.sql")

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
