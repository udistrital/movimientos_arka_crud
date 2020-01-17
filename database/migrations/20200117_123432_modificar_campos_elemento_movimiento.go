package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModificarCamposElementoMovimiento_20200117_123432 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModificarCamposElementoMovimiento_20200117_123432{}
	m.Created = "20200117_123432"

	migration.Register("ModificarCamposElementoMovimiento_20200117_123432", m)
}

// Run the migrations
func (m *ModificarCamposElementoMovimiento_20200117_123432) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE movimientos_arka.elementos_movimiento ADD COLUMN elemento_catalogo_id integer;")

}

// Reverse the migrations
func (m *ModificarCamposElementoMovimiento_20200117_123432) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE movimientos_arka.elementos_movimiento DROP COLUMN elemento_catalogo_id integer;")
}
