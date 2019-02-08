package tests

import (
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/anmolbabu/go-syscall-ex/pkg/utils"
)

func TestGetPID(t *testing.T) {
	var cmdBytes []byte
	var err error
	t.Run("GetPID", func(t *testing.T) {
		gotPID := utils.GetPID()
		cmd := "p=`echo $$`;ps -f $p --no-headers | awk '{print $3}'"
		cmdBytes, err = exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatalf("error executing command. Error %+v", err)
		}

		cmdOp := strings.TrimSpace(string(cmdBytes))

		if cmdOp == string(gotPID) {
			t.Fatalf("expected %+v, got %+v", gotPID, cmdOp)
		}
	})
}

func TestGetUID(t *testing.T) {
	var cmdBytes []byte
	var err error
	t.Run("GetUID", func(t *testing.T) {
		gotUID := utils.GetUID()

		gotPID := utils.GetPID()
		cmd := "stat -c '%u' /proc/" + strconv.Itoa(gotPID) + "/"
		cmdBytes, err = exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			t.Fatalf("error executing command. Error %+v", err)
		}

		cmdOp := strings.TrimSpace(string(cmdBytes))

		if cmdOp == string(gotUID) {
			t.Fatalf("expected %+v, got %+v", gotUID, cmdOp)
		}
	})
}
