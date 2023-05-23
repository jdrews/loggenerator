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
		err2 := file.Close()
		if err2 != nil {
			log.Fatalf("failed closing file on shutdown: %s", err2)
		}
		fmt.Println("\rGoodbye!")
		os.Exit(0)
	}()

	err = file.Close()
	if err != nil {
		log.Fatalf("failed closing file on startup: %s", err)
	}

	// Begin opening file, write line, flush, and close file.
	i := 0
	for {
		openFile, err2 := os.OpenFile(*logPtr, os.O_APPEND, 0644)
		if err2 != nil {
			log.Fatalf("failed opening file: %s", err2)
		}

		datawriter := bufio.NewWriter(openFile)

		_, err3 := datawriter.WriteString(fmt.Sprint(i, ": ", prependStr, " ", generator.LogLine()))
		if err3 != nil {
			log.Fatalf("failed writing string: %s", err3)
		}

		err4 := datawriter.Flush()
		if err4 != nil {
			log.Fatalf("failed flushing file: %s", err4)
		}

		err5 := openFile.Close()
		if err5 != nil {
			log.Fatalf("failed closing file: %s", err5)
		}

		time.Sleep(time.Duration(*intervalPtr) * time.Millisecond)
		i++
	}
}
