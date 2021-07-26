package alp

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outw, errw := new(bytes.Buffer), new(bytes.Buffer)
	p := NewProfiler(outw, errw)

	data := []string{
		"time:2018-10-14T05:58:05+09:00",
		"method:POST",
		"uri:/foo/bar?token=xxx&uuid=1234",
		"status:200",
		"size:12",
		"request_time:0.247",
		"apptime:0.057",
	}

	ltsvData := strings.Join(data, "\t")

	tmpfile, err := os.CreateTemp(os.TempDir(), "alp-mock-stdin")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	if _, err := tmpfile.Write([]byte(ltsvData)); err != nil {
		t.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	p.SetInReader(tmpfile)

	err = p.Run([]string{"ltsv"})
	if err != nil {
		t.Fatal(err)
	}
}
