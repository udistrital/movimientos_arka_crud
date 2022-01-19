package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EstandarFormatos_20220119_002057 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EstandarFormatos_20220119_002057{}
	m.Created = "20220119_002057"

	migration.Register("EstandarFormatos_20220119_002057", m)
}

// Run the migrations
func (m *EstandarFormatos_20220119_002057) Up() {
	file, err := ioutil.ReadFile("../scripts/20220119_002057_estandar_formatos_up.sql")

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
func (m *EstandarFormatos_20220119_002057) Down() {
	file, err := ioutil.ReadFile("../scripts/20220119_002057_estandar_formatos_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
