package main

import (
	"fmt"
	"testing"
)

func TestGraphDefinition(t *testing.T) {
	var conntrack ConntrackPlugin

	graphdef := conntrack.GraphDefinition()
	if len(graphdef) != 1 {
		t.Errorf("GetTempfilename: %d should be 1", len(graphdef))
	}
}

func TestCurrentValue(t *testing.T) {
	samplePaths := []string{
		"./sample/nf_conntrack_count",
	}

	_, err := CurrentValue(samplePaths)
	if err != nil  {
		t.Errorf("%v", err)
	}

	failPaths := []string{
		"./sample/nf_conntrack_unknown",
	}

	expect := fmt.Sprintf("Cannot find files %s", failPaths)
	_, err = CurrentValue(failPaths)
	if err.Error() != expect {
		t.Errorf("expect %q to be equal %q", err, expect)
	}
}
