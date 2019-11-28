package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasMovimientosArka_20191108_131137 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasMovimientosArka_20191108_131137{}
	m.Created = "20191108_131137"

	migration.Register("CrearTablasMovimientosArka_20191108_131137", m)
}

// Run the migrations
func (m *CrearTablasMovimientosArka_20191108_131137) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191108_131137_crear_tablas_movimientos_arka.up.sql")

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
func (m *CrearTablasMovimientosArka_20191108_131137) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191108_131137_crear_tablas_movimientos_arka.down.sql")

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
