package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CentroCostos_20230824_175423 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CentroCostos_20230824_175423{}
	m.Created = "20230824_175423"

	migration.Register("CentroCostos_20230824_175423", m)
}

const script = "../scripts/20230824_175423_centro_costos_"

// Run the migrations
func (m *CentroCostos_20230824_175423) Up() {
	file, err := os.ReadFile(script + "up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *CentroCostos_20230824_175423) Down() {
	file, err := os.ReadFile(script + "down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
