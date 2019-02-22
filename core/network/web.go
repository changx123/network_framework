package network

import (
	"blog/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var Wroutes *gin.Engine

func init() {
	if config.WEB_DEBUG {
		gin.SetMode(gin.DebugMode)
		Wroutes = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		Wroutes = gin.New()
	}
}

type Engine struct {
	*gin.Engine
}

type Listener struct {
	net.Listener
}

var (
	server   *http.Server
	listener net.Listener
)

func WRun() {
	if config.HTTPS_OPEN {
		go Wroutes.RunTLS(config.HTTPS_LISTEN_ADDRESS, config.HTTPS_CERTFILE_PATH, config.HTTPS_KEYFILE_PATH)
	}
	if !config.HTTP_HOT_UPDATE && config.HTTP_OPEN {
		go Wroutes.Run(config.HTTP_LISTEN_ADDRESS)
	}
	if config.HTTP_HOT_UPDATE && config.HTTP_OPEN {
		en := &Engine{Wroutes}
		s := en.RunHotUpdate(config.HTTP_LISTEN_ADDRESS)
		listener = s.Listener
	}
}

func (engine *Engine) RunHotUpdate(addr ...string) (*Listener) {
	debugPrint("Listening and serving HTTP on %s\n", addr)
	address := resolveAddress(addr)
	server = &http.Server{Addr: address, Handler: engine}
	var listener net.Listener
	var err error
	if len(os.Args) >= 2 && os.Args[1] == "-child" {
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
		if err != nil {
			panic(err)
		}
	} else {
		listener, err = net.Listen("tcp4", server.Addr)
	}
	if err != nil {
		panic(err)
	}
	go func() {
		defer func() { debugPrintError(err) }()
		server.Serve(listener)
	}()
	return &Listener{listener}
}

func debugPrint(format string, values ...interface{}) {
	if gin.IsDebugging() {
		log.Printf("[GIN-debug] "+format, values...)
	}
}

func debugPrintError(err error) {
	if err != nil {
		debugPrint("[ERROR] %v\n", err)
	}
}

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		if port := os.Getenv("PORT"); port != "" {
			debugPrint("Environment variable PORT=\"%s\"", port)
			return ":" + port
		}
		debugPrint("Environment variable PORT is undefined. Using port :8080 by default")
		return ":8080"
	case 1:
		return addr[0]
	default:
		panic("too much parameters")
	}
}

func SingalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-ch
		fmt.Printf("signal: %v\n", sig)

		ctx, _ := context.WithTimeout(context.Background(), config.HTTP_HOT_UPDATE_TIMEOUT*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			// reload
			log.Printf("restart")
			err := restart()
			if err != nil {
				fmt.Printf("graceful restart failed: %v\n", err)
			}
			//更新当前pidfile
			UpdatePidFile()
			server.Shutdown(ctx)
			os.Exit(1)
			fmt.Printf("graceful reload\n")
			return
		}
	}
}

func restart() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return fmt.Errorf("listener is not tcp listener")
	}
	f, err := tl.File()
	if err != nil {
		return err
	}
	cmd := exec.Command(os.Args[0], "-child")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}

func UpdatePidFile() {
	sPid := fmt.Sprint(os.Getpid())
	if len(os.Args) >= 2 && os.Args[1] == "-child" {
		if err := procExsit(); err != nil {
			fmt.Printf("pid file exists, update\n")
		} else {
			fmt.Printf("pid file NOT exists, create\n")
		}
	}
	pidFile, _ := os.Create("./process.pid")
	defer pidFile.Close()
	pidFile.WriteString(sPid)
}

// 判断进程是否启动
func procExsit() (err error) {
	pidFile, err := os.Open("./process.pid")
	defer pidFile.Close()
	if err != nil {
		return nil
	}
	filePid, err := ioutil.ReadAll(pidFile)
	if err != nil {
		return nil
	}
	pidStr := fmt.Sprintf("%s", filePid)
	pid, _ := strconv.Atoi(pidStr)
	if _, err := os.FindProcess(pid); err != nil {
		fmt.Printf("Failed to find process: %v\n", err)
		return nil
	}

	return nil
}
