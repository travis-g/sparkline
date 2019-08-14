package main

import (
	"log"
	"os/exec"
	"testing"
)

var result interface{}

func benchmarkExec(cmd string, b *testing.B) {
	var (
		out  []byte
		err  error
		data = []string{"0", "1", "2", "8", "9", "30", "3", "18", "1", "3", "10"}
	)
	for n := 0; n < b.N; n++ {
		out, err = exec.Command(cmd, data...).Output()
		if err != nil {
			log.Fatalf("cmd failed: %s\n", err)
		}
	}
	result = out
}

func BenchmarkShell(b *testing.B) {
	benchmarkExec("./mkspark", b)
}
