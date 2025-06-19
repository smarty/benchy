package strategies

import (
	"bytes"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"testing"
	"time"
)

// PProfCPUStrategy records CPU PProf readings and writes them to disk.
type PProfCPUStrategy interface {
	// StartRecording starts recording PProf for the CPU.
	StartRecording()

	// StopRecording stops the recording PProf for the CPU.
	StopRecording()

	// WriteRecording writes the most recent recording to disk.
	WriteRecording()

	// WriteRunnerScript writes a bash script that will gather all the PProf
	// files when running for convenient analysis.
	WriteRunnerScript()
}

// ---- Active ------

type ActivePProfCPU struct {
	b              *testing.B
	name           string
	saveDirectory  string
	currentProfile bytes.Buffer
	totalFiles     int
}

func NewActivePProfCPU(b *testing.B, name string) PProfCPUStrategy {
	// if a pprof is already running, then fail the benchmark and return the
	// null version
	err := pprof.StartCPUProfile(&bytes.Buffer{})
	if err != nil {
		b.Errorf("a pprof is already running, cannot profile benchmark '%s'", name)
		return &NullPProfCPU{}
	}

	pprof.StopCPUProfile()

	pathName := fmt.Sprintf("./workspace/%s/%s/", time.Now().Format("20060102-150405"), name)
	os.MkdirAll(pathName, 0777)

	return &ActivePProfCPU{
		b:             b,
		name:          name,
		saveDirectory: pathName,
	}
}

func (this *ActivePProfCPU) StartRecording() {
	this.currentProfile.Reset()
	pprof.StartCPUProfile(&this.currentProfile)
}

func (this *ActivePProfCPU) StopRecording() {
	pprof.StopCPUProfile()
}

func (this *ActivePProfCPU) WriteRecording() {
	err := os.WriteFile(
		fmt.Sprintf("%s/cpu_%d.pprof", this.saveDirectory, this.totalFiles),
		this.currentProfile.Bytes(),
		0777)

	this.totalFiles++
	if err != nil {
		this.b.Error(err)
		return
	}
}

func (this *ActivePProfCPU) WriteRunnerScript() {
	sb := strings.Builder{}
	sb.WriteString("#!/bin/bash\ngo tool pprof -http localhost:8080")
	for i := range this.totalFiles {
		sb.WriteString(fmt.Sprintf(" cpu_%d.pprof", i))
	}

	os.WriteFile(fmt.Sprintf("%s/%s", this.saveDirectory, "pprof.sh"), []byte(sb.String()), 0777)
}

// ---- NULL ------

type NullPProfCPU struct {
}

func NewNullPProfCPU() *NullPProfCPU {
	return &NullPProfCPU{}
}

func (this *NullPProfCPU) StartRecording()    {}
func (this *NullPProfCPU) StopRecording()     {}
func (this *NullPProfCPU) WriteRecording()    {}
func (this *NullPProfCPU) WriteRunnerScript() {}
