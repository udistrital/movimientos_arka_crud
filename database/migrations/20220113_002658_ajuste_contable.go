package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteContable_20220113_002658 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteContable_20220113_002658{}
	m.Created = "20220113_002658"

	migration.Register("AjusteContable_20220113_002658", m)
}

// Run the migrations
func (m *AjusteContable_20220113_002658) Up() {
	file, err := ioutil.ReadFile("../scripts/20220113_002658_ajuste_contable_up.sql")

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
func (m *AjusteContable_20220113_002658) Down() {
	file, err := ioutil.ReadFile("../scripts/20220113_002658_ajuste_contable_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
