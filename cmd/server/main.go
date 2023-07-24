package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/cmd", cmdHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var cmd Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Println(cmd.Command)

	parts := strings.Split(cmd.Command, " ")
	execCmd := exec.Command(parts[0], parts[1:]...)

	var stdOut bytes.Buffer
	execCmd.Stdout = &stdOut

	var stdErr bytes.Buffer
	execCmd.Stderr = &stdErr

	err = execCmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	result := Result{
		Stdout: stdOut.String(),
		Stderr: stdErr.String(),
	}

	jsonResp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonResp)
	return
}

type Command struct {
	Command string
}

type Result struct {
	Stdout string
	Stderr string
}
