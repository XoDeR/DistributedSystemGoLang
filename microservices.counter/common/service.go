package common

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gziphandler"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type (
	RPCService struct {
		client Client
	}
	Config struct {
		Server     *server.Config
		ItemsToken string
	}
)

func NewRPCService(cfg *Config) *RPCService {
	return &RPCService{
		NewClient(cfg.ItemsToken),
	}
}

func (s *RPCService) Prefix() string {
	return "/svc/nyt"
}

func (s *RPCService) Service() (*grpc.ServiceDesc, interface{}) {
	return &_NYTProxyService_serviceDesc, s
}

func (s *RPCService) Middleware(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(h)
}

func (s *RPCService) ContextMiddleware(h server.ContextHandler) server.ContextHandler {
	return h
}

func (s *RPCService) JSONMiddleware(j server.JSONContextEndpoint) server.JSONContextEndpoint {
	return func(ctx context.Context, r *http.Request) (int, interface{}, error) {
		status, res, err := j(ctx, r)
		if err != nil {
			server.LogWithFields(r).WithFields(logrus.Fields{
				"error": err,
			}).Error("problems with serving request")
			return http.StatusServiceUnavailable, nil, &jsonErr{"sorry, this service is unavailable"}
		}

		server.LogWithFields(r).Info("success!")
		return status, res, nil
	}
}

func (s *RPCService) ContextEndpoints() map[string]map[string]server.ContextHandlerFunc {
	return map[string]map[string]server.ContextHandlerFunc{}
}

func (s *RPCService) JSONEndpoints() map[string]map[string]server.JSONContextEndpoint {
	return map[string]map[string]server.JSONContextEndpoint{
		"/items/{itemName}/{tenantName}": map[string]server.JSONContextEndpoint{
			"POST": s.AddNewItemWithTenantJSON,
		},
		"/items/{tenantName}/count": map[string]server.JSONContextEndpoint{
			"GET": s.GetItemCountByTenantJSON,
		},
	}

	/*
			//set mime type to JSON
		response.Header().Set("Content-type", "application/json")

		err := request.ParseForm()
		if err != nil {
			http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
		}

		// fixed size 1000
		var result = make([]string, 1000)

		switch request.Method {
		case "GET":

			id := strings.Replace(request.URL.Path, "/items/", "", -1)

			// query items
			tenant := request.PostFormValue("tenant")

			var itemsCount uint32 = getItemsCount(tenant)

			//	b, err := json.Marshal(itemList)
			//	if err != nil {
			//	fmt.Println(err)
			//return
			//	}
			//result[i] = fmt.Sprintf("%s", string(b))

			result = result[:1]

		case "POST":
			id := request.PostFormValue("id")

			// insert new entry with id and tenant
			var res bool = true
			if res == true {
				result[0] = "true"
			}
			result = result[:1]

		default:
		}

		json, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprintf(response, "%v", string(json))
	*/
}

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string {
	return e.Err
}
