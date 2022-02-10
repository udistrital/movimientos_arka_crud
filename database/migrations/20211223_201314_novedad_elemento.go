package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type NovedadElemento_20211223_201314 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NovedadElemento_20211223_201314{}
	m.Created = "20211223_201314"

	migration.Register("NovedadElemento_20211223_201314", m)
}

// Run the migrations
func (m *NovedadElemento_20211223_201314) Up() {
	file, err := ioutil.ReadFile("../scripts/20211223_201314_novedad_elemento_up.sql")

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
func (m *NovedadElemento_20211223_201314) Down() {
	file, err := ioutil.ReadFile("../scripts/20211223_201314_novedad_elemento_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
