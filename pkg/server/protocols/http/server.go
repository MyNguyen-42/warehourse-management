package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
)

func EchoServer() *echo.Echo {
	// Echo instance
	server := echo.New()

	//server.HTTPErrorHandler = errorHandler
	server.Binder = &auditBinder{defaultBinder: new(echo.DefaultBinder)}

	// Middleware
	server.Use(echoMiddleware.CORS())

	server.Any("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	return server
}

type auditBinder struct {
	defaultBinder echo.Binder
}

// Bind Implement echo#Binder
func (binder *auditBinder) Bind(i interface{}, ctx echo.Context) error {
	if jsonBin := ctx.Request().Context().Value(auditBinder{}); jsonBin != nil {
		return json.Unmarshal(jsonBin.([]byte), i)
	}

	if err := binder.defaultBinder.Bind(i, ctx); err != nil {
		return fmt.Errorf("bind: %w", err)
	}

	jsonBin, err := json.Marshal(i)
	if err != nil {
		return err
	}
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), auditBinder{}, jsonBin)))
	return nil
}

//func errorHandler(err error, c echo.Context) {
//	he, ok := err.(*echo.HTTPError)
//	if ok {
//		if he.Internal != nil {
//			if herr, ok := he.Internal.(*echo.HTTPError); ok {
//				he = herr
//			}
//		}
//
//		if msg, ok := he.Message.(string); ok {
//			log.Logger(c.Request().Context()).Sugar().Errorf("got echo internal code %d error %s", he.Code, msg)
//		}
//		err = he
//	}
//
//	// Send response
//	if !c.Response().Committed {
//		if c.Request().Method == http.MethodHead {
//			err = c.NoContent(he.Code)
//		} else {
//			err = api.ResponseError(c, err)
//		}
//	}
//}
