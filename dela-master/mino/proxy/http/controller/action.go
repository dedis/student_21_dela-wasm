package controller

import (
	"fmt"
	"time"

	"go.dedis.ch/dela/cli/node"
	"go.dedis.ch/dela/mino/proxy"
	"go.dedis.ch/dela/mino/proxy/http"
	"golang.org/x/xerrors"
)

var defaultRetry = 10
var proxyFac func(string) proxy.Proxy = http.NewHTTP

type startAction struct{}

// Execute implements node.ActionTemplate. It starts and injects the proxy http
// server.
func (a startAction) Execute(ctx node.Context) error {

	addr := ctx.Flags.String("clientaddr")

	proxyhttp := proxyFac(addr)

	ctx.Injector.Inject(proxyhttp)

	go proxyhttp.Listen()

	for i := 0; i < defaultRetry && proxyhttp.GetAddr() == nil; i++ {
		time.Sleep(time.Second)
	}

	if proxyhttp.GetAddr() == nil {
		return xerrors.Errorf("failed to start proxy server")
	}

	// We assume the listen worked proprely, however it might not be the case.
	// The log should inform the user about that.
	fmt.Fprintf(ctx.Out, "started proxy server on %s", proxyhttp.GetAddr().String())

	return nil
}
