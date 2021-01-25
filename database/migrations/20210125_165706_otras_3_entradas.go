package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Otras3Entradas_20210125_165706 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Otras3Entradas_20210125_165706{}
	m.Created = "20210125_165706"

	migration.Register("Otras3Entradas_20210125_165706", m)
}

// Run the migrations
func (m *Otras3Entradas_20210125_165706) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210125_165706_otras_3_entradas_up.sql")

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
func (m *Otras3Entradas_20210125_165706) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210125_165706_otras_3_entradas_down.sql")

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
