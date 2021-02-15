package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type FixEntradaAprovechamiento_20210129_003900 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &FixEntradaAprovechamiento_20210129_003900{}
	m.Created = "20210129_003900"

	migration.Register("FixEntradaAprovechamiento_20210129_003900", m)
}

// Run the migrations
func (m *FixEntradaAprovechamiento_20210129_003900) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210129_003900_fix_entrada_aprovechamiento_up.sql")
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
func (m *FixEntradaAprovechamiento_20210129_003900) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210129_003900_fix_entrada_aprovechamiento_down.sql")

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
