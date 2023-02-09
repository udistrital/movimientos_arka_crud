package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Consecutivo_20230208_220542 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Consecutivo_20230208_220542{}
	m.Created = "20230208_220542"

	migration.Register("Consecutivo_20230208_220542", m)
}

// Run the migrations
func (m *Consecutivo_20230208_220542) Up() {
	file, err := ioutil.ReadFile("../scripts/20230208_220542_consecutivo_up.sql")

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
func (m *Consecutivo_20230208_220542) Down() {
	file, err := ioutil.ReadFile("../scripts/20230208_220542_consecutivo_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
