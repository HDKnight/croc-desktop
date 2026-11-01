package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/schollz/croc/v11/src/models"
	"github.com/schollz/croc/v11/src/tcp"
)

// croc-relay is the relay from this repository. Senders must present -pass.
// A receiver running stock croc can join with the relay address and code;
// that client sends the public default password and cannot open a room.
func main() {
	pass := flag.String("pass", os.Getenv("CROC_PASS"), "password required to open a room")
	portsFlag := flag.String("ports", "9009,9010,9011,9012,9013", "comma-separated relay ports; the first is the control port")
	flag.Parse()

	password := strings.TrimSpace(*pass)
	if password == "" {
		password = models.DEFAULT_PASSPHRASE
	}
	var ports []string
	for _, port := range strings.Split(*portsFlag, ",") {
		if port = strings.TrimSpace(port); port != "" {
			ports = append(ports, port)
		}
	}
	if len(ports) < 2 {
		fmt.Fprintln(os.Stderr, "relay requires at least two ports")
		os.Exit(2)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	capability, err := tcp.NewRelayCapabilitySet(min(len(ports)-1, 8))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	banner := strings.Join(ports[1:], ",")
	errc := make(chan error, len(ports))
	for i, port := range ports {
		go func(port string, control bool) {
			if control {
				errc <- tcp.RunWithOptionsAsync("", port, password,
					tcp.WithBanner(banner),
					tcp.WithLogLevel("info"),
					tcp.WithFastAdmission(capability),
					tcp.WithCtx(ctx),
				)
				return
			}
			errc <- tcp.RunWithOptionsAsync("", port, password,
				tcp.WithLogLevel("info"),
				tcp.WithFastAdmission(capability),
				tcp.WithCtx(ctx),
			)
		}(port, i == 0)
	}
	if err := <-errc; err != nil && ctx.Err() == nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
