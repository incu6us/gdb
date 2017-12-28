package barkup

// TODO: prepare a tests
import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	// TarCmd is the path to the `tar` executable
	CopyCmd = "cp"
)

// Location is an `Exporter` interface that backs up a local file or directory
type Location struct {
	Path string
}

// Export produces a specified location or directory, and creates a gzip compressed tarball archive.
func (x Location) Export() *ExportResult {
	result := &ExportResult{MIME: "application/x-tar"}

	origPath := strings.TrimSuffix(x.Path, string(os.PathSeparator))
	dumpName := strings.Split(origPath, string(os.PathSeparator))
	dumpPath := fmt.Sprintf(`%v_%v`, dumpName[len(dumpName)-1], time.Now().Unix())

	cpOut, err := exec.Command(CopyCmd, "-r", origPath, dumpPath).Output()
	if err != nil {
		result.Error = makeErr(err, string(cpOut))
		return result
	}

	result.Path = dumpPath + ".tar.gz"
	out, err := exec.Command(TarCmd, "-czf", result.Path, dumpPath).Output()
	if err != nil {
		result.Error = makeErr(err, string(out))
		return result
	}
	if err := os.RemoveAll("./"+dumpPath); err != nil {
		result.Error = makeErr(err, "")
		return result
	}

	return result
}
