package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RefactorTraslados_20220519_171140 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RefactorTraslados_20220519_171140{}
	m.Created = "20220519_171140"

	migration.Register("RefactorTraslados_20220519_171140", m)
}

// Run the migrations
func (m *RefactorTraslados_20220519_171140) Up() {
	file, err := ioutil.ReadFile("../scripts/20220519_171140_refactor_traslados_up.sql")

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
func (m *RefactorTraslados_20220519_171140) Down() {
	file, err := ioutil.ReadFile("../scripts/20220519_171140_refactor_traslados_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
