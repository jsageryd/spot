package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	ok, err := isDotInstalled()
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		fmt.Println("dot executable not found.")
		return
	}
	http.HandleFunc("/", httpGenerate)
	listenAddr := ":4649"
	log.Printf("Listening at %s\n", listenAddr)
	err = http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func httpGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}
	err := makePDF(w, r.Body)
	if err == nil {
		log.Printf("Generated PDF for %s\n", r.RemoteAddr)
	} else {
		log.Printf("Error generating PDF for %s: %s\n", r.RemoteAddr, err)
	}
}

func makePDF(w io.Writer, r io.Reader) error {
	dot := exec.Command("dot", "-Tpdf")
	stdin, err := dot.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := dot.StdoutPipe()
	if err != nil {
		return err
	}
	_, err = io.Copy(stdin, r)
	if err != nil {
		return err
	}
	stdin.Close()
	err = dot.Start()
	if err != nil {
		return err
	}
	_, err = io.Copy(w, stdout)
	if err != nil {
		return err
	}
	err = dot.Wait()
	if err != nil {
		return err
	}
	return nil
}

func isDotInstalled() (bool, error) {
	cmd := exec.Command("which", "dot")
	if err := cmd.Start(); err != nil {
		return false, err
	}
	if err := cmd.Wait(); err != nil {
		return false, nil
	}
	return true, nil
}
