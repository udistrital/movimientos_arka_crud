# movimientos_arka_crud

Api CRUD para el manejo de los movimiento contables o generales en el Sistema de Gestión de Almacén e Inventarios ARKA II

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
MOVIMIENTOS_ARKA_CRUD_HTTP_PORT: Puerto asignado para la ejecución del API
MOVIMIENTOS_ARKA_CRUD_PGUSER: Usuario de la base de datos
MOVIMIENTOS_ARKA_CRUD_PGPASS: Clave del usuario para la conexión a la base de datos  
MOVIMIENTOS_ARKA_CRUD_PGURLS: Host de conexión
MOVIMIENTOS_ARKA_CRUD_PGDB: Nombre de la base de datos
MOVIMIENTOS_ARKA_CRUD_SCHEMA: Esquema a utilizar en la base de datos
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con MOVIMIENTOS_ARKA_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/movimientos_arka_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/movimientos_arka_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
MOVIMIENTOS_ARKA_CRUD_HTTP_PORT=8080 MOVIMIENTOS_ARKA_CRUD_DB_HOST=127.0.0.1:27017 MOVIMIENTOS_ARKA_CRUD_SOME_VARIABLE=some_value bee run
```
### Ejecución Dockerfile
```shell
# docker build --tag=movimientos_arka_crud . --no-cache
# docker run -p 80:80 movimientos_arka_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/movimientos_arka_crud

#2. Moverse a la carpeta del repositorio
cd movimientos_arka_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.6.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/movimientos_arka_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/movimientos_arka_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/movimientos_arka_crud/status.svg?ref=refs/heads/release/0.6.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/movimientos_arka_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/movimientos_arka_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/movimientos_arka_crud) |

## Modelo de Datos
[Modelo de Datos MOVIMIENTOS ARKA CRUD](https://drive.google.com/drive/u/2/folders/1VXw4sg2JiGH8PvLzff3J2gmW5igcq4V6)\
[PGModeler](models/modelo.dbm) - [SVG](models/modelo.svg)


## Licencia

This file is part of movimientos_arka_crud.

movimientos_arka_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

movimientos_arka_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with proyecto_academico_crud. If not, see https://www.gnu.org/licenses/.