package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambioJsonbFormatoTipoMovimiento_20220405_084621 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambioJsonbFormatoTipoMovimiento_20220405_084621{}
	m.Created = "20220405_084621"

	migration.Register("CambioJsonbFormatoTipoMovimiento_20220405_084621", m)
}

// Run the migrations
func (m *CambioJsonbFormatoTipoMovimiento_20220405_084621) Up() {
	file, err := ioutil.ReadFile("../scripts/20220405_084621_cambio_jsonb_formato_tipo_movimiento.up.sql")

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
func (m *CambioJsonbFormatoTipoMovimiento_20220405_084621) Down() {
	file, err := ioutil.ReadFile("../scripts/20220405_084621_cambio_jsonb_formato_tipo_movimiento.down.sql")

	if err != nil {
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}
}
