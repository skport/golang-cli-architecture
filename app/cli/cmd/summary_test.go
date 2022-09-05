// Functional Test : Cmd Summary

package cmd

import (
	"regexp"
	"testing"

	"github.com/spf13/cobra"
)

func TestSummaryCmdRun(t *testing.T) {
	cmd := &cobra.Command{
		Use: "Test",
		Run: SummaryCmdRun,
	}

	// TestCase : Error 異常系 (TableDrivenTests)
	var equalTests = []struct {
		in     []string
		out    string
		errMsg string
	}{
		{
			[]string{""},
			"cannot be blank",
			"Checking empty param",
		},
		{
			[]string{"httpps://github.com"},
			"must be a valid URL",
			"Checking wrong URL",
		},
	}
	for _, testCase := range equalTests {
		cmd.SetArgs(testCase.in)
		stdout := extractStdout(t, func() {
			cmd.Execute()
		})
		if stdout != testCase.out {
			t.Errorf("Unexpected output: %s", testCase.errMsg)
		}
	}

	// TestCase : Normal　正常系
	cmd.SetArgs([]string{"https://github.com"})
	stdout := extractStdout(t, func() {
		cmd.Execute()
	})
	re := regexp.MustCompile(`^\[title`)
	if re.MatchString(stdout) == false {
		t.Errorf("Unexpected output: %s", "Checking body")
	}
}
