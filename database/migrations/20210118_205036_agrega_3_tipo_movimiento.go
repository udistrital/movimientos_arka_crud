package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Agrega3TipoMovimiento_20210118_205036 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Agrega3TipoMovimiento_20210118_205036{}
	m.Created = "20210118_205036"

	migration.Register("Agrega3TipoMovimiento_20210118_205036", m)
}

// Up Run the migrations
func (m *Agrega3TipoMovimiento_20210118_205036) Up() {
	file, err := ioutil.ReadFile("../scripts/20210118_205036_agrega_3_tipo_movimiento_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Down Reverse the migrations
func (m *Agrega3TipoMovimiento_20210118_205036) Down() {
	file, err := ioutil.ReadFile("../scripts/20210118_205036_agrega_3_tipo_movimiento_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
