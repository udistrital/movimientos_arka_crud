package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/astaxie/beego"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

// @opt opciones de godog
var opt = godog.Options{Output: colors.Colored(os.Stdout)}

// @resStatus codigo de respuesta a las solicitudes a la api
var resStatus string

// @resBody JSON repuesta Delete
var resDelete string

// @resBody JSON de respuesta a las solicitudesde la api
var resBody []byte

// @especificacion estructura de la fecha
const especificacion = "Jan 2, 2006 at 3:04pm (MST)"

var savepostres map[string]interface{}

var IntentosAPI = 1

var Id float64

var integrationTestsEnabled bool
var beeCmd *exec.Cmd

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

// @run_bee activa el servicio de la api para realizar los test
func run_bee() error {
	if missing := missingIntegrationConfig(); len(missing) > 0 {
		return fmt.Errorf("faltan variables de configuracion para pruebas de integracion: %s", strings.Join(missing, ", "))
	}

	command := fmt.Sprintf(
		"cd .. && MOVIMIENTOS_ARKA_CRUD_PGPORT=%s MOVIMIENTOS_ARKA_CRUD_HTTP_PORT=%s MOVIMIENTOS_ARKA_CRUD_PGUSER=%s MOVIMIENTOS_ARKA_CRUD_PGPASS=%s MOVIMIENTOS_ARKA_CRUD_PGURLS=%s MOVIMIENTOS_ARKA_CRUD_PGDB=%s MOVIMIENTOS_ARKA_CRUD_SCHEMA=%s RUN_MODE=test bee run",
		beego.AppConfig.String("PGport"),
		beego.AppConfig.String("httpport"),
		beego.AppConfig.String("PGuser"),
		beego.AppConfig.String("PGpass"),
		beego.AppConfig.String("PGurls"),
		beego.AppConfig.String("PGdb"),
		beego.AppConfig.String("PGschemas"),
	)

	beeCmd = exec.Command("sh", "-c", command)
	beeCmd.Stdout = os.Stdout
	beeCmd.Stderr = os.Stderr
	if err := beeCmd.Start(); err != nil {
		return err
	}

	if err := waitForAPI(apiBaseURL(), 60*time.Second); err != nil {
		return err
	}

	fmt.Println("El API se Encuentra en Estado OK")
	return nil
}

func shouldRunIntegrationTests() bool {
	return strings.EqualFold(os.Getenv("RUN_INTEGRATION_TESTS"), "true") ||
		strings.EqualFold(os.Getenv("MOVIMIENTOS_ARKA_CRUD_RUN_INTEGRATION_TESTS"), "true")
}

func missingIntegrationConfig() []string {
	required := map[string]string{
		"PGurls":    beego.AppConfig.String("PGurls"),
		"httpport":  beego.AppConfig.String("httpport"),
		"PGuser":    beego.AppConfig.String("PGuser"),
		"PGpass":    beego.AppConfig.String("PGpass"),
		"PGdb":      beego.AppConfig.String("PGdb"),
		"PGport":    beego.AppConfig.String("PGport"),
		"PGschemas": beego.AppConfig.String("PGschemas"),
	}

	var missing []string
	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			missing = append(missing, key)
		}
	}
	return missing
}

func apiBaseURL() string {
	return "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport")
}

func waitForAPI(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	client := &http.Client{Timeout: 5 * time.Second}
	var lastErr error

	for time.Now().Before(deadline) {
		resp, err := client.Get(url)
		if err == nil {
			if resp.Body != nil {
				resp.Body.Close()
			}
			return nil
		}

		lastErr = err
		fmt.Println("Intento de subir el API numero: " + strconv.Itoa(IntentosAPI))
		IntentosAPI++
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("numero de intentos maximos alcanzados, revise variables de entorno o disponibilidad del puerto: %w", lastErr)
}

func stopBee() {
	if beeCmd != nil && beeCmd.Process != nil {
		_ = beeCmd.Process.Kill()
		_, _ = beeCmd.Process.Wait()
	}
}

// @init inicia la aplicacion para realizar los test
func init() {
	integrationTestsEnabled = shouldRunIntegrationTests()
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

// @TestMain para realizar la ejecucion con el comando go test ./test
func TestMain(m *testing.M) {
	if !integrationTestsEnabled {
		fmt.Println("Pruebas de integracion omitidas. Defina RUN_INTEGRATION_TESTS=true para habilitarlas.")
		os.Exit(0)
	}

	fmt.Println("Inicio de pruebas de integracion del API")
	if err := run_bee(); err != nil {
		fmt.Println(err.Error())
		stopBee()
		os.Exit(1)
	}

	ts := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "progress",
			Paths:  []string{"features"},
			Output: colors.Colored(os.Stdout),
		},
	}
	status := ts.Run()

	if st := m.Run(); st > status {
		status = st
	}
	stopBee()
	os.Exit(status)
}

// @gen_files genera los archivos de ejemplos
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
