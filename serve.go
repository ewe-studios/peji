package peji

import (
	"github.com/ewestudio/sabuhp/radar"
)

func PageToRadarEvent(router *radar.Mux, master *Pages) OnPages {
	return func(route string, _ *Pages) {
		router.Event(route, master)
	}
}
