package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UniqueElementoActa_20230331_182119 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UniqueElementoActa_20230331_182119{}
	m.Created = "20230331_182119"

	migration.Register("UniqueElementoActa_20230331_182119", m)
}

// Run the migrations
func (m *UniqueElementoActa_20230331_182119) Up() {
	file, err := ioutil.ReadFile("../scripts/20230331_182119_unique_elemento_acta_up.sql")

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
func (m *UniqueElementoActa_20230331_182119) Down() {
	file, err := ioutil.ReadFile("../scripts/20230331_182119_unique_elemento_acta_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
