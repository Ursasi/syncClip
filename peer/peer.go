package peer

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	constant "syncClip"
	"syncClip/peer/handler"
	"syncClip/server/service"
	"syncClip/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
)

type Config struct {
	Address       string
	Port          int
	ServerAddress string
	ServerPort    int
	Concurrency   int
	Logging       bool
}

type State struct {
	ID     string
	Boards []service.Board
	mu     sync.Mutex
	Cond   *sync.Cond
}

var PeerState State
var peerConfig Config
var clipChan = make(chan string, 2048)

func StartPeer(config Config) {
	_, _ = maxprocs.Set()
	peerConfig = config
	initFlags()

	go startBackgroundTask()
	startHTTPServer()
}

func startBackgroundTask() {
	var err error
	PeerState.ID, PeerState.Boards, err = registerPeer()
	PeerState.Cond = sync.NewCond(&PeerState.mu)
	if err != nil {
		return
	}
	go func() {
		err := startProbePeriodically()
		if err != nil {
			log.Printf("start probe error: %v", err)
		}
	}()
	watchClip()
}
func initFlags() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(errors.Wrap(err, "fatal error viper bindPFlags"))
	}
}
func startHTTPServer() {
	r := gin.Default()
	r.POST("/receive", handler.Receive)
	addr := fmt.Sprintf("%s:%d", peerConfig.Address, peerConfig.Port)
	if err := r.Run(addr); err != nil {
		panic(errors.Wrap(err, "failed to start the server"))
	}
}

func watchClip() {
	clipChan = make(chan string)
	if runtime.GOOS == "darwin" {
		go WatchClipboard()
	}
	go func() {
		for msg := range clipChan {
			if msg != "" {
				fmt.Println("send to server:", msg)
				// TODO: Implement sending to server
			}
		}
	}()
}

func registerPeer() (string, []service.Board, error) {
	client := util.NewRestyClient(2 * time.Second)
	resp, err := client.R().
		SetBody(map[string]string{
			"IP":   peerConfig.Address,
			"Port": fmt.Sprintf("%d", peerConfig.Port),
			"MAC":  "00:00:00:00:00:00",
		}).
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("http://%s:%d/register", peerConfig.ServerAddress, peerConfig.ServerPort))
	if err != nil || resp == nil || resp.RawResponse.StatusCode != http.StatusOK {
		log.Println(errors.Wrap(err, "failed to register"))
		return "", nil, errors.Wrap(err, "failed to register")
	}
	var response util.RegisterResponse
	bodyReader := resp.RawResponse.Body
	bodyBytes, err := io.ReadAll(bodyReader)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Println(errors.Wrap(err, "json unmarshal failed"))
		return "", nil, errors.Wrap(err, "json unmarshal failed")
	}
	return response.ID, response.Boards, nil
}

func startProbe() error {
	client := util.NewRestyClient(2 * time.Second)
	resp, err := client.R().
		SetBody(map[string]string{
			"ID": PeerState.ID,
		}).
		SetHeader("Content-Type", "application/json").
		Post(fmt.Sprintf("http://%s:%d/probe", peerConfig.ServerAddress, peerConfig.ServerPort))
	if err != nil || resp == nil || resp.RawResponse.StatusCode != http.StatusOK {
		log.Println(errors.Wrap(err, "failed to probe"))
		return errors.Wrap(err, "failed to probe")
	}
	var response util.ProbeResponse
	bodyReader := resp.RawResponse.Body
	bodyBytes, err := io.ReadAll(bodyReader)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Println(errors.Wrap(err, "json unmarshal failed"))
		return errors.Wrap(err, "json unmarshal failed")
	}
	PeerState.mu.Lock()
	defer PeerState.mu.Unlock()

	changed := hasBoardsChanged(PeerState.Boards, response.Boards)
	if changed {
		PeerState.Boards = response.Boards
		PeerState.Cond.Broadcast()
	}
	return nil
}

func startProbePeriodically() error {
	ticker := time.NewTicker(constant.DefaultProbeInterval * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := startProbe(); err != nil {
			log.Printf("probe error：%v", err)
		}
	}
	return nil
}

func hasBoardsChanged(oldBoards, newBoards []service.Board) bool {
	if len(oldBoards) != len(newBoards) {
		return true
	}
	oldMap := make(map[string]service.Board)
	for _, board := range oldBoards {
		oldMap[board.ID] = board
	}
	for _, board := range newBoards {
		if oldBoard, exists := oldMap[board.ID]; !exists || !isBoardEqual(oldBoard, board) {
			return true
		}
	}
	return false
}

func isBoardEqual(a, b service.Board) bool {
	return a.ID == b.ID && a.IP == b.IP && a.Port == b.Port
}
