package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstadosBajas_20211126_204708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstadosBajas_20211126_204708{}
	m.Created = "20211126_204708"

	migration.Register("EstadosBajas_20211126_204708", m)
}

// Run the migrations
func (m *EstadosBajas_20211126_204708) Up() {
	file, err := ioutil.ReadFile("../scripts/20211126_204708_estados_bajas_up.sql")

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
func (m *EstadosBajas_20211126_204708) Down() {
	file, err := ioutil.ReadFile("../scripts/20211126_204708_estados_bajas_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
