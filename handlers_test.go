package main

import (
	"bytes"
	"fmt"
	"github.com/asdine/storm"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const appjson = "application/json"

func TestCreateAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)

	body := ParseResponseBody(res)

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, address, body)
}

func TestFindAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = BasicGet(url + "/address/1")

	body := ParseResponseBody(res)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, address, body)
}

// func TestFindAllAddress(t *testing.T) {
// 	app := createTestApplication()
// 	engine := setupRouter(app)

// 	w := httptest.NewRecorder()
// 	add := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
// 	new := bytes.NewBufferString(add)

// 	req, _ := http.NewRequest("POST", "/address", new)
// 	engine.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)

// 	w = httptest.NewRecorder()
// 	add = `{"id":2,"firstname":"sam","lastname":"john","email":"john@sam.com", "phone": "333-231-3456"}`
// 	new := bytes.NewBufferString(add)

// 	req, _ := http.NewRequest("POST", "/address", new)
// 	engine.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)

// 	w = httptest.NewRecorder()

// 	req, _ = http.NewRequest("GET", "/address/", bytes.NewBuffer([]byte{}))
// 	engine.ServeHTTP(w, req)

// 	expected := `[{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"},{"id":2,"firstname":"sam","lastname":"john","email":"john@sam.com", "phone": "333-231-3456"}]`

// 	assert.Equal(t, 200, w.Code)
// 	assert.JSONEq(t, expected, w.Body.String())

// }

func TestUpdateAddress(t *testing.T) {

}

func TestDeleteAddress(t *testing.T) {
	httpserver := createTestApplication()
	defer httpserver.Close()
	url := httpserver.URL

	address := `{"id":1,"firstname":"john","lastname":"snow","email":"snow@winterfell.com", "phone": "133-333-1313"}`
	new := bytes.NewBufferString(address)
	res, err := BasicPost(url+"/address", appjson, new)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	res, err = BasicDelete(url+"/address/1", appjson, nil)
	assert.NoError(t, err)

	body := ParseResponseBody(res)

	expected := "Success"
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expected, body)
}

func createTestDatabase() *storm.DB {
	Dir := path.Join("/tmp/gin-rest-api", fmt.Sprintf("%d", time.Now().UnixNano()))
	err := os.MkdirAll(Dir, os.FileMode(0700))
	if err != nil {
		panic(err)
	}

	Dir = path.Join(Dir, "test.db")
	db, _ := storm.Open(Dir)

	return db
}

func createTestApplication() *httptest.Server {
	db := createTestDatabase()
	app := &Application{Database: db}
	api := newApiServer(app)
	return api
}

func newApiServer(app *Application) *httptest.Server {
	apiEngine := setupRouter(app)
	return httptest.NewServer(apiEngine)
}

func ParseResponseBody(resp *http.Response) string {
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(resp.Body)
	mustNotErr(resp.Body.Close())
	return buf.String()
}
