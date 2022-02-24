package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Amortizacion_20220222_233226 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Amortizacion_20220222_233226{}
	m.Created = "20220222_233226"

	migration.Register("Amortizacion_20220222_233226", m)
}

// Run the migrations
func (m *Amortizacion_20220222_233226) Up() {
	file, err := ioutil.ReadFile("../scripts/20220222_233226_amortizacion_up.sql")

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
func (m *Amortizacion_20220222_233226) Down() {
	file, err := ioutil.ReadFile("../scripts/20220222_233226_amortizacion_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
