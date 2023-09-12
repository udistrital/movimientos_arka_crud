package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RefactorCentroCostos_20230907_113125 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RefactorCentroCostos_20230907_113125{}
	m.Created = "20230907_113125"

	migration.Register("RefactorCentroCostos_20230907_113125", m)
}

// Run the migrations
func (m *RefactorCentroCostos_20230907_113125) Up() {
	file, err := os.ReadFile("../scripts/20230907_113125_refactor_centro_costos_up.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}

// Reverse the migrations
func (m *RefactorCentroCostos_20230907_113125) Down() {
	file, err := os.ReadFile("../scripts/20230907_113125_refactor_centro_costos_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
