package stats_test

import (
	"testing"

	"github.com/colinnewell/stats-summary/stats"

	"github.com/google/go-cmp/cmp"
)

func TestAddStats(t *testing.T) {
	s := stats.New(3)
	s.RecordStats("a.b.c")
	s.RecordStats("d.e.f")
	s.RecordStats("a.d.c")
	s.RecordStats("a.d.c")
	s.RecordStats("a.d.c")
	s.RecordStats("a.b.c")
	s.RecordStats("d.e.f")

	want := []stats.Stats{
		{PartialName: "a.b", Count: 1},
		{PartialName: "d", Count: 2},
		{PartialName: "a.d", Count: 2},
		{PartialName: "a", Count: 4},
	}

	if diff := cmp.Diff(want, s.Summary()); diff != "" {
		t.Errorf("Got incorrect results (-want +got):\n%s\n", diff)
	}
}
