package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RegistroInmuebles_20221222_005703 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RegistroInmuebles_20221222_005703{}
	m.Created = "20221222_005703"

	migration.Register("RegistroInmuebles_20221222_005703", m)
}

// Run the migrations
func (m *RegistroInmuebles_20221222_005703) Up() {
	file, err := ioutil.ReadFile("../scripts/20221222_005703_registro_inmuebles_up.sql")

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
func (m *RegistroInmuebles_20221222_005703) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20221222_005703_registro_inmuebles_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
