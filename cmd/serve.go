package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var (
	debugMode       bool
	serverAddress   string
	shutdownTimeout time.Duration

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Launches the svc-fizzbuzz service webserver",
		Long: fmt.Sprintf(`serve command launch a fizzbuzz api server.

Curl examples:
  $ curl -X GET    http://localhost%[1]s/
  $ curl -X GET    http://localhost%[1]s/version
  $ curl -X GET    http://localhost%[1]s/metrics
  $ curl -X GET    http://localhost%[1]s/status
  $ curl -X GET    http://localhost%[1]s/api/v1/fizzbuzz
  $ curl -X GET    http://localhost%[1]s/api/v1/fizzbuzz?limit=70&mul1=7&mul1=9&word1=bon&word2=coin
  $ curl -X GET    http://localhost%[1]s/api/v1/fizzbuzz/top
  $ curl -X GET    http://localhost%[1]s/api/v1/hits`, defautAddress),
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	// force debug mode
	serveCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Force debug mode")

	// address flag
	serveCmd.PersistentFlags().StringVarP(&serverAddress, "address", "a", defautAddress, "HTTP server address")

	// address flag
	serveCmd.PersistentFlags().DurationVarP(&shutdownTimeout, "shutdown-timeout", "t", 10*time.Second, "shutdown timeout (5s,5m,5h) before connections are cancelled)")

	// Here you will define your flags and configuration settings.
}

func serve() {
	initLogger()
	log.Infof("%s version %s - %s", svc.Name, svc.Version, svcName)

	srv := getServer()
	go launchServer(srv)
	waitForShutdown(srv)
}

func initLogger() {
	// initalize the debug mode
	if strings.HasSuffix(svc.Version, "+dev") || debugMode {
		log.SetLevel(log.DebugLevel)
		log.WithFields(log.Fields{
			"Name":     svc.Name,
			"Version":  svc.Version,
			"FullName": fmt.Sprintf("%s-%s", svc.Name, svc.Version),
		}).Debug("set log debug level")
	}

}

func getServer() *http.Server {
	log.WithFields(log.Fields{"address": serverAddress, "shutdownTimeout": shutdownTimeout}).Debug("Flags")
	// Got service router and launch a gracefull shutdown server
	mux := svc.NewRouter()
	return &http.Server{Addr: serverAddress, Handler: mux}
}

func waitForShutdown(srv *http.Server) {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	<-stop

	log.Printf("%s shutting down ...\n", svcName)

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s down\n", svcName)
}

func launchServer(srv *http.Server) {
	log.Printf("%s listening on %s with %v timeout", svcName, serverAddress, shutdownTimeout)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}
