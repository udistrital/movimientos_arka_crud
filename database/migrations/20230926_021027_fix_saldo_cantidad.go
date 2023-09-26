package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type FixSaldoCantidad_20230926_021027 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &FixSaldoCantidad_20230926_021027{}
	m.Created = "20230926_021027"

	migration.Register("FixSaldoCantidad_20230926_021027", m)
}

// Run the migrations
func (m *FixSaldoCantidad_20230926_021027) Up() {
	file, err := os.ReadFile("../scripts/20230926_021027_fix_saldo_cantidad_up.sql")

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
func (m *FixSaldoCantidad_20230926_021027) Down() {
	file, err := os.ReadFile("../scripts/20230926_021027_fix_saldo_cantidad_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
