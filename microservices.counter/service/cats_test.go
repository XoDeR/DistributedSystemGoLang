package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"microservices.counter/common"
	"microservices.counter/common/clientTest"

	"github.com/NYTimes/gizmo/server"
)

func TestGetCats(t *testing.T) {
	tests := []struct {
		givenURI    string
		givenClient *clientTest.Client

		wantCode int
		wantBody interface{}
	}{
		{
			"/svc/nyt/cats",
			&clientTest.Client{
				TestSemanticConceptSearch: func(conceptType, concept string) ([]*nyt.SemanticConceptArticle, error) {
					return []*nyt.SemanticConceptArticle{
						&nyt.SemanticConceptArticle{
							Url: "https://www.nytimes.com/awesome-article",
						},
					}, nil
				},
			},

			http.StatusOK,
			[]interface{}{
				map[string]interface{}{
					"url": "https://www.nytimes.com/awesome-article",
				},
			},
		},
		{
			"/svc/nyt/cats",
			&clientTest.Client{
				TestSemanticConceptSearch: func(conceptType, concept string) ([]*nyt.SemanticConceptArticle, error) {
					return nil, errors.New("NOPE!")
				},
			},

			http.StatusServiceUnavailable,
			map[string]interface{}{
				"error": "sorry, this service is unavailable",
			},
		},
	}

	for _, test := range tests {

		srvr := server.NewRPCServer(nil)
		srvr.Register(&RPCService{client: test.givenClient})

		r, _ := http.NewRequest("GET", test.givenURI, nil)
		w := httptest.NewRecorder()
		srvr.ServeHTTP(w, r)

		if w.Code != test.wantCode {
			t.Errorf("expected response code of %d; got %d", test.wantCode, w.Code)
		}

		var got interface{}
		err := json.NewDecoder(w.Body).Decode(&got)
		if err != nil {
			t.Error("unable to JSON decode response body: ", err)
		}

		if !reflect.DeepEqual(got, test.wantBody) {
			t.Errorf("expected response body of\n%#v;\ngot\n%#v", test.wantBody, got)
		}
	}

}
