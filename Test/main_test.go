package Test

import (
	"fmt"

	"github.com/kriangkrai/SQL/RUNSQL/Controller"
	"github.com/kriangkrai/SQL/RUNSQL/Router"

	"testing"
)

// func TestGetData(t *testing.T) {

// 	Controller.Connect()
// 	r := Router.SetupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/user/get", nil)
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotNil(t, w.Body)
// }

func TestDatas(t *testing.T) {

	r := Router.SetupRouter()

	fmt.Println(r)
	data := Controller.GetData()

	var datas = data[0].Device

	if datas != "Sensor4" {
		t.Errorf("%s", datas)
	}

	//r := Router.SetupRouter()

	//w := httptest.NewRecorder()
	//req, _ := http.NewRequest("GET", "/user/get", nil)

	//r.ServeHTTP(w, req)

	//assert.Equal(t, http.StatusOK, w.Code)
	//assert.NotNil(t, w.Body)
}
