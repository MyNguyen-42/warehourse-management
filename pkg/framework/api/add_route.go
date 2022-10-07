package api

import (
	"context"
	"github.com/labstack/echo/v4"
)

type NhanhAPI[Req interface{}, Resp interface{}] func(ctx context.Context, req Req) (Resp, error)

func AddRoute[Req interface{}, Resp interface{}](router *echo.Group, path string, api NhanhAPI[Req, Resp]) {
	router.POST(path, func(c echo.Context) error {
		var req = new(Req)
		ctx := c.Request().Context()
		if err := c.Bind(req); err != nil {
			return err
		}

		resp, err := api(ctx, *req)
		if err != nil {
			return err
		}

		return ResponseOK(c, resp)
	})
	//if docsGenerator != nil {
	//	var req = new(Req)
	//	_ = json.Unmarshal([]byte("{}"), req)
	//	docsGenerator.AddRequest("POST", path, *req)
	//}
}

//var docsGenerator *postman.DocGenerator
//
//func InitDocGenerator(name string, host string, header map[string]string) {
//	docsGenerator = postman.NewDocGenerator(name, host, header)
//}
//
//func ExportDocsToFile(path string) error {
//	return docsGenerator.GenerateFile(path)
//}
