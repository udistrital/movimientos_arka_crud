package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	"github.com/udistrital/utils_oas/request"
)

//@opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

// @resBody JSON repuesta Delete
var resDelete string

//@resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

//@especificacion estructura de la fecha
const especificacion = "Jan 2, 2006 at 3:04pm (MST)"

var savepostres map[string]interface{}

var IntentosAPI = 1

var Id float64

// @estructura de las tablas parametricas
type Parametrica struct {
	Nombre            string
	Descripcion       string
	CodigoAbreviacion string
	Activo            bool
	NumeroOrden       float64
	FechaCreacion     time.Time
	FechaModificacion time.Time
}

//@exe_cmd ejecuta comandos en la terminal
func exe_cmd(cmd string, wg *sync.WaitGroup) {

	parts := strings.Fields(cmd)
	out, err := exec.Command(parts[0], parts[1]).Output()

	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
}

// @deleteFile Borrar archivos
func deleteFile(path string) {
	// delete file
	err := os.Remove(path)
	if err != nil {
		fmt.Errorf("no se pudo eliminar el archivo")
	}

}

//@run_bee activa el servicio de la api para realizar los test
func run_bee() {
	var resultado map[string]interface{}

	parametros := "MOVIMIENTOS_ARKA_CRUD_PGPORT=" + beego.AppConfig.String("PGport") + " MOVIMIENTOS_ARKA_CRUD_HTTP_PORT=" + beego.AppConfig.String("httpport") + " MOVIMIENTOS_ARKA_CRUD_PGUSER=" + beego.AppConfig.String("PGuser") + " MOVIMIENTOS_ARKA_CRUD_PGPASS=" + beego.AppConfig.String("PGpass") + " MOVIMIENTOS_ARKA_CRUD_PGURLS=" + beego.AppConfig.String("PGurls") + " MOVIMIENTOS_ARKA_CRUD_PGDB=" + beego.AppConfig.String("PGdb") + " MOVIMIENTOS_ARKA_CRUD_PGSCHEMA=" + beego.AppConfig.String("PGschemas") + " bee run"
	file, err := os.Create("script.sh")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintln(file, "cd ..")
	fmt.Fprintln(file, parametros)

	wg := new(sync.WaitGroup)
	commands := []string{"sh script.sh &"}
	for _, str := range commands {
		wg.Add(1)
		go exe_cmd(str, wg)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("Obteniendo respuesta de http://" + beego.AppConfig.String("appurl") + ":" + beego.AppConfig.String("httpport"))
	errApi := request.GetJson("http://"+beego.AppConfig.String("PGurls")+":"+beego.AppConfig.String("httpport"), &resultado)
	if errApi == nil && resultado != nil {
		fmt.Println("El API se Encuentra en Estado OK")
	} else if IntentosAPI <= 3 {

		stri := strconv.Itoa(IntentosAPI)
		fmt.Println("Intento de subir el API numero: " + stri)
		IntentosAPI++
		run_bee()
	} else {
		fmt.Println("Numero de intentos maximos alcanzados, revise por favor variables de entorno o si no esta ocupado el puerto")
	}

	deleteFile("script.sh")
	wg.Done()
}

//@init inicia la aplicacion para realizar los test
func init() {
	fmt.Println("Inicio de pruebas Unitarias al API")

	// gen_files()
	run_bee()
	//pasa las banderas al comando godog
	godog.BindFlags("godog.", flag.CommandLine, &opt)

}

//@TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	// init()
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: "progress",
		Paths:  []string{"features"},
		//Randomize: time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)

}

//@gen_files genera los archivos de ejemplos
func gen_files() {
	fmt.Println("Genera los archivos")
	t := time.Now()

	nombre := "Prueba_test" // t.Format(especificacion) //se cambia para que cumpla con la especificacion de varying(20)
	atributo := Parametrica{
		Nombre:            nombre,
		Descripcion:       "string",
		CodigoAbreviacion: "string",
		Activo:            true,
		NumeroOrden:       1,
		FechaCreacion:     t,
		FechaModificacion: t,
	}

	rankingsJson, _ := json.Marshal(atributo)
	ioutil.WriteFile("./assets/requests/BodyGen1.json", rankingsJson, 0644)
	ioutil.WriteFile("./assets/requests/BodyGen2.json", rankingsJson, 0644)
	ioutil.WriteFile("./assets/requests/BodyGen3.json", rankingsJson, 0644)
	ioutil.WriteFile("./assets/requests/BodyGen4.json", rankingsJson, 0644)
}
