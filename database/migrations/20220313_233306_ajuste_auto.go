package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteAuto_20220313_233306 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteAuto_20220313_233306{}
	m.Created = "20220313_233306"

	migration.Register("AjusteAuto_20220313_233306", m)
}

// Run the migrations
func (m *AjusteAuto_20220313_233306) Up() {
	file, err := ioutil.ReadFile("../scripts/20220313_233306_ajuste_auto_up.sql")

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
func (m *AjusteAuto_20220313_233306) Down() {
	file, err := ioutil.ReadFile("../scripts/20220313_233306_ajuste_auto_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
