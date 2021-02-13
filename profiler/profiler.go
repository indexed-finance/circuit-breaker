package profiler

import (
	"net/http"
	"net/http/pprof"

	"github.com/arl/statsviz"
)

// Profiler bundles together net/http/pprof and statsviz
type Profiler struct {
	srv *http.Server
}

// New returns a new profiler server ready for starting
func New() *Profiler {
	return &Profiler{}
}

// Start runs a lightweight http server for net/http/pprof with optional stats visualization
// and launches the http server as a goroutine
func (p *Profiler) Start(addr string, enableStatsviz bool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	if enableStatsviz {
		statsviz.Register(mux)
	}
	p.srv = &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go p.srv.ListenAndServe()
}

// Close stops the profiler http server ignoring err server closed
func (p *Profiler) Close() error {
	if err := p.srv.Close(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
