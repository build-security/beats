package bundle

import (
	"log"
	"net/http"
	"time"

	csppolicies "github.com/elastic/csp-security-policies/bundle"
)

var Config = `{
        "services": {
            "test": {
                "url": %q
            }
        },
        "bundles": {
            "test": {
                "resource": "/bundles/bundle.tar.gz"
            }
        },
        "decision_logs": {
            "console": true
        }
    }`

func CreateServer() (*http.Server, error) {
	policies, err := csppolicies.CISKubernetes()
	if err != nil {
		return nil, err
	}

	bundleServer := csppolicies.NewServer()
	err = bundleServer.HostBundle("bundle.tar.gz", policies)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      bundleServer, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return srv, nil
}
