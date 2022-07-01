package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstadosCierres_20220630_195444 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstadosCierres_20220630_195444{}
	m.Created = "20220630_195444"

	migration.Register("EstadosCierres_20220630_195444", m)
}

// Run the migrations
func (m *EstadosCierres_20220630_195444) Up() {
	file, err := ioutil.ReadFile("../scripts/20220630_195444_estados_cierres_up.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *EstadosCierres_20220630_195444) Down() {
	file, err := ioutil.ReadFile("../scripts/20220630_195444_estados_cierres_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
