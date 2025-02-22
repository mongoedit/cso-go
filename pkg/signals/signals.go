package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Gracefully handle SIGINT and SIGTERM

Pass in a an optional list of cleanup funcs to make sure your app exists cleanly
*/
func HandleInterrupt(cleanupJobs ...func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Print("\u001b[2k")
		fmt.Println("\u001b[1A")
		fmt.Println("Cleaning up and exiting...")
		for _, job := range cleanupJobs {
			job()
		}
		os.Exit(0)
	}()
}
