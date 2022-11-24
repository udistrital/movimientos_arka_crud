package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/cucumber/godog"
	"github.com/xeipuuv/gojsonschema"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

//@AreEqualJSON comparar dos JSON si son iguales retorna true de lo contrario false
func AreEqualJSON(s1, s2 string) (bool, error) {

	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 1 :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string 2 :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

//@toJson convierte string en JSON
func toJson(p interface{}) string {

	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

//@getPages convierte en un tipo el json
func getPages(ruta string) []byte {

	raw, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []byte
	c = raw
	return c
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
}

//@iSendRequestToWhereBodyIsJson realiza la solicitud a la API
func iSendRequestToWhereBodyIsJson(arg1, arg2, arg3 string) error {

	var url string

	if arg1 == "GET" || arg1 == "POST" {
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + arg2

	} else {
		if arg1 == "PUT" || arg1 == "DELETE" {
			str := strconv.FormatFloat(Id, 'f', 5, 64)
			url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + arg2 + "/" + str

		}
	}
	if arg1 == "GETID" {
		arg1 = "GET"
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + arg2 + "/" + str

	}
	if arg1 == "DELETE" {
		str := strconv.FormatFloat(Id, 'f', 0, 64)
		url = "http://" + beego.AppConfig.String("PGurls") + ":" + beego.AppConfig.String("httpport") + arg2 + "/" + str
		resDelete = "{\"Id\":" + str + "}"
		ioutil.WriteFile("./assets/responses/Ino.json", []byte(resDelete), 0644)

	}

	pages := getPages(arg3)

	req, err := http.NewRequest(arg1, url, bytes.NewBuffer(pages))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyr, _ := ioutil.ReadAll(resp.Body)

	resStatus = resp.Status
	resBody = bodyr

	if arg1 == "POST" && resStatus == "201 Created" {
		ioutil.WriteFile("./assets/requests/BodyRec2.json", resBody, 0644)
		json.Unmarshal([]byte(bodyr), &savepostres)
		Id = savepostres["Id"].(float64)

	}
	return nil

}

//@theResponseCodeShouldBe valida el codigo de respuesta
func theResponseCodeShouldBe(arg1 string) error {
	if resStatus != arg1 {
		return fmt.Errorf("se esperaba el codigo de respuesta .. %s .. y se obtuvo el codigo de respuesta .. %s .. ", arg1, resStatus)
	}
	return nil
}

//@theResponseShouldMatchJson valida el JSON de respuesta
func theResponseShouldMatchJson(arg1 string) error {
	div := strings.Split(arg1, "")

	pages := getPages(arg1)
	//areEqual, _ := AreEqualJSON(string(pages), string(resBody))
	if div[13] == "V" {
		schemaLoader := gojsonschema.NewStringLoader(string(pages))
		documentLoader := gojsonschema.NewStringLoader(string(resBody))
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if result.Valid() {
			return nil
		} else {
			return fmt.Errorf("Errores : %s", result.Errors())
		}
	}
	if div[13] == "I" {
		areEqual, _ := AreEqualJSON(string(pages), string(resBody))
		if areEqual {
			return nil
		} else {
			return fmt.Errorf(" se esperaba el body de respuesta %s y se obtuvo %s", string(pages), resBody)
		}

	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I send "([^"]*)" request to "([^"]*)" where body is json "([^"]*)"$`, iSendRequestToWhereBodyIsJson)
	s.Step(`^the response code should be "([^"]*)"$`, theResponseCodeShouldBe)
	s.Step(`^the response should match json "([^"]*)"$`, theResponseShouldMatchJson)
}
