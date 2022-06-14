module github.com/udistrital/movimientos_arka_crud

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/cucumber/godog v0.10.0
	github.com/lib/pq v1.10.2
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/udistrital/auditoria v0.0.0-20200115201815-9680ae9c2515
	github.com/udistrital/utils_oas v0.0.0-20211125230753-1091d2af48e2
	github.com/xeipuuv/gojsonschema v1.2.0
)

replace github.com/astaxie/beego v1.12.3 => github.com/udistrital/beego v1.12.4-0.20211126032252-ee78ca48b207
