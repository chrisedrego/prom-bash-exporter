package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	var path = "/"
	var port = ":8080"
	ExecCommand()
	http.HandleFunc(path, MetricsServer)
	http.ListenAndServe(port, nil)
}

func MetricsServer(w http.ResponseWriter, r *http.Request) {
	var PATH = r.URL.Path[1:]
	fmt.Print(PATH)
	if PATH == "" {
		fmt.Fprintf(w, "Bash-Exporter")
	} else if PATH == "metrics" {
		fmt.Fprintf(w, "# HELP bash-exporter based logs\n")
		fmt.Fprintf(w, "bash_exporter_metric_name{id=\"attrib_value\"} 0\n")
	}
}

func ExecCommand() {
	cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
