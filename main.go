package main

import (
   log "github.com/sirupsen/logrus"
   // "github.com/zserge/lorca"

   // "runtime"
   "os"
	// "net"
	// "net/http"
	// "os/signal"
	// "fmt"
	// "embed"
	// "sync"
)

/// go:embed www
// var fs embed.FS

func initLogger() {
   // The TextFormatter is default, you don't actually have to do this.
   log.SetFormatter(&log.TextFormatter{
      ForceColors: true,
      DisableLevelTruncation: true,
      PadLevelText: true,
      FullTimestamp: false,
   })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type install struct {
   installLocation string
   fileName string
}

type fileContents struct {
	// sync.Mutex
	baseImage string
	installer string
   workDir string
   installs []install
}

func main() {
   initLogger()
   log.Info("DockerfileGen starting...")

   // args := []string{}
   // if runtime.GOOS == "linux" {
   //   args = append(args, "--class=Lorca")
   // }
   //
   // ui, err := lorca.New("", "", 480, 320, args...)
   //
   // if err != nil {
 	// 	log.Fatal(err)
   // }
 	// defer ui.Close()
   //
   // ui.Bind("start", func() {
	// 	log.Info("UI is ready")
	// })
   //
   // err = ui.Bind("getBaseImage", func(acBaseImage string) string { return acBaseImage})
   // if err != nil {
   //    log.Fatal("Failed to bind to UI")
   // }
   //
   // path := ui.Eval(`getBaseImage("arch")`).String()
   // log.Debug(path)
   //
   // ln, err := net.Listen("tcp", "127.0.0.1:0")
   // if err != nil {
   //    log.Fatal(err)
   // }
   // defer ln.Close()
   // go http.Serve(ln, http.FileServer(http.FS(fs)))
   // ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
   //
   // defer func(ui lorca.UI) {
   //    err := ui.Close()
   //    if err != nil {
   //       log.Fatal("Failed to close UI")
   //    }
   // }(ui)

   // Wait for the browser window to be closed
   // sigc := make(chan os.Signal)

   // signal.Notify(sigc, os.Interrupt)

   generateDockerfile()

   // select {
   // case <-sigc:
   // // case <-ui.Done():
   // }
   log.Info("DockerfileGen exiting...")
}
