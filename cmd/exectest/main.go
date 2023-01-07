package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(os.Getwd())
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "python", "scripts/listen_stream.py")

	b, err := cmd.CombinedOutput()
	fmt.Printf("output: %s\n", b)
	if err != nil {
		panic(err)
	}

	cancel()
	fmt.Println(cmd.CombinedOutput())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}
