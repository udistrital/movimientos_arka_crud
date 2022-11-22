package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type VidaUtilElementos_20221121_110038 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &VidaUtilElementos_20221121_110038{}
	m.Created = "20221121_110038"

	migration.Register("VidaUtilElementos_20221121_110038", m)
}

// Run the migrations
func (m *VidaUtilElementos_20221121_110038) Up() {
	file, err := ioutil.ReadFile("../scripts/20221121_110038_vida_util_elementos_up.sql")

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
func (m *VidaUtilElementos_20221121_110038) Down() {
	file, err := ioutil.ReadFile("../scripts/20221121_110038_vida_util_elementos_down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
