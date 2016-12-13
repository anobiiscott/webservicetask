package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"testing"
)

type HandlerTest struct {
	url    string
	method string
	handler func(http.ResponseWriter, *http.Request)
	status int
	response string
}

var HandlerTests = []HandlerTest {
	{"/", "GET", HelloHandler, http.StatusUnprocessableEntity, `{"code":422,"result":"Unprocessable Entity"}`},
	{"/?name=foo", "GET", HelloHandler, http.StatusOK, `"hello foo"`},
	{"/developers", "GET", DevelopersGET, http.StatusOK, `{"Name":"Alex","Age":28,"Language":"Go","Floor":5}`},
	{"/developers/1", "GET", DeveloperGET, http.StatusOK, `{"Name":"Charlie","Age":23,"Language":"Go","Floor":5}`},
	{"/developers/12", "GET", DeveloperGET, http.StatusNotFound, `{"code":404,"result":"Developer not found"}`},
	{"/developers/0", "GET", DeveloperGET, http.StatusNotFound, `{"code":404,"result":"Developer not found"}`},
	{"/developers?name=foo&age=33&language=PHP&floor=5", "POST", DevelopersPOST, http.StatusOK, `{"Name":"foo","Age":33,"Language":"PHP","Floor":5}`},
	{"/developers", "POST", DevelopersPOST, http.StatusUnprocessableEntity, `{"code":422,"result":"Unprocessable Entity"}`},
}

func TestHandlers(t *testing.T) {

	for _, ht := range HandlerTests {
		req, err := http.NewRequest(ht.method, ht.url, nil)

		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		Router().ServeHTTP(w, req)

		//check status and output is correct
		assert.Equal(t, w.Code, ht.status)
		assert.Contains(t, w.Body.String(), ht.response)
	}
}

func TestSendResponse(t *testing.T) {
	res, _ := json.Marshal(Response{http.StatusNotFound, "Developer not found"})
	w := httptest.NewRecorder()

	SendResponse(res, w, http.StatusNotFound)
	assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"code":404,"result":"Developer not found"}`, w.Body.String())
}

func BenchmarkSendResponse(b *testing.B) {
	res, _ := json.Marshal(Response{http.StatusNotFound, "Developer not found"})
	w := httptest.NewRecorder()

	for n := 0; n < b.N; n++ {
		SendResponse(res, w, http.StatusNotFound)
	}
}
