package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jdrews/loggenerator/generator"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	logPtr := flag.String("logfile", "test/logfile.log", "Name of logfile that receives the generated log lines")
	intervalPtr := flag.Int("interval", 1000, "Log line generation interval in milliseconds")
	prependLognamePtr := flag.Bool("prependlogname", false, "Prepend the name of the logfile to the loglines")
	flag.Parse()
	prependStr := ""
	if *prependLognamePtr {
		prependStr = *logPtr
	}

	// Initialize the generator
	generator.Initialize()

	// Prepare file
	err := os.MkdirAll(filepath.Dir(*logPtr), 0777)
	if err != nil {
		log.Fatalf("failed creating directory path: %s", err)
	}
	file, err := os.OpenFile(*logPtr, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rCtrl+C pressed in Terminal, closing file...")
		file.Close()
		fmt.Println("\rGoodbye!")
		os.Exit(0)
	}()

	file.Close()

	// Begin opening file, write line, flush, and close file.
	i := 0
	for {
		file, err := os.OpenFile(*logPtr, os.O_APPEND, 0644)

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
			os.Exit(1)
		}
		datawriter := bufio.NewWriter(file)
		datawriter.WriteString(fmt.Sprint(i, ": ", prependStr, " (", time.Now().Format(time.RFC3339), ") [", generator.RandomSeverity(), "] ", generator.Paragraph(), ">>STOP\n"))
		datawriter.Flush()
		file.Close()
		time.Sleep(time.Duration(*intervalPtr) * time.Millisecond)
		i++
	}
}
