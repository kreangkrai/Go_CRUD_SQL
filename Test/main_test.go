package Test

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/go-playground/assert/v2"
	"github.com/kriangkrai/SQL/RUNSQL/Models"
	"github.com/kriangkrai/SQL/RUNSQL/Router"

	"testing"
)

func TestGetLocation(t *testing.T) {

	r, _ := Router.SetupRouter()

	//Get
	req := httptest.NewRequest("GET", "/location/get", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	result := []Models.Location{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := Models.Location{
		ID:       9,
		Name:     "XXX",
		Location: "HQ",
	}

	assert.Equal(t, expect, result[0])

}
func TestAddLocation(t *testing.T) {

	r, _ := Router.SetupRouter()
	//Add

	data := `{
				"Name" : "XXX",
				"Location" : "HQ"
			}`

	req := httptest.NewRequest("POST", "/location/insert", strings.NewReader(data))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	result := Models.Location{}

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := Models.Location{
		Name:     "XXX",
		Location: "HQ",
	}

	assert.Equal(t, expect, result)

}

func TestUpdateLocation(t *testing.T) {

	r, _ := Router.SetupRouter()

	//Update

	data := `{
				"ID" : 2,
				"Name" : "Mee",
				"Location" : "HQ"
			}`

	req := httptest.NewRequest("PUT", "/location/update", strings.NewReader(data))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var result string

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := "Update Success"

	assert.Equal(t, expect, result)

}
func TestDeleteLocation(t *testing.T) {

	r, _ := Router.SetupRouter()
	//Delete

	req := httptest.NewRequest("DELETE", "/location/delete/2", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)
	var result string

	err := json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	expect := "Delete Success"

	assert.Equal(t, expect, result)

}
