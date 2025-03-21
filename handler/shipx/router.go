package shipx

import "github.com/xgfone/ship/v5"

type RouteRegister interface {
	RegisterRoute(r *ship.RouteGroupBuilder) error
}

func RegisterRoutes(r *ship.RouteGroupBuilder, routes []RouteRegister) error {
	for _, route := range routes {
		if route == nil {
			continue
		}
		if err := route.RegisterRoute(r); err != nil {
			return err
		}
	}

	return nil
}
