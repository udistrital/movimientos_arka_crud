package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertBajasTraslados_20210212_125542 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertBajasTraslados_20210212_125542{}
	m.Created = "20210212_125542"

	migration.Register("InsertBajasTraslados_20210212_125542", m)
}

// Run the migrations
func (m *InsertBajasTraslados_20210212_125542) Up() {
	file, err := ioutil.ReadFile("../scripts/20210212_125542_insertBajasTraslados_up.sql")

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
func (m *InsertBajasTraslados_20210212_125542) Down() {
	file, err := ioutil.ReadFile("../scripts/20210212_125542_insertBajasTraslados_down.sql")

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
