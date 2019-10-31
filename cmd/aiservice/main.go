package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/spacepal/Spacepal-AI/internal/ai"
	"github.com/spacepal/Spacepal-AI/internal/ai/list"
	"github.com/spacepal/Spacepal-AI/internal/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	var addr = fmt.Sprint(":", os.Getenv("PORT"))
	flag.Parse()
	var aiManager = ai.NewManager()
	list.RegisterAll(aiManager)
	http.Handle("/ai/names", server.NewAINamesHandler(aiManager))
	var turnHander = ai.NewTurnHandler(aiManager)
	go turnHander.Start()
	http.Handle("/ai/do", server.NewDoHandler(turnHander, aiManager))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
