package main

import (
	"net/http"
	"os/exec"
	"time"
)

func main() {
	lmao := time.Now().Format("🌈 02 Jan")
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", lmao).Run()
	exec.Command("git", "push", "origin", "main").Run()
	http.Get("https://betteruptime.com/api/v1/heartbeat/zfnS1uFSeYdSwQY41Na7mMRW")
}