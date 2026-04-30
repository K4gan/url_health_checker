package main

import "testing"

func TestSummarize(t *testing.T) {
    got := summarize([]string{"api:ok", "api:slow", "worker:ok"})
    if len(got) != 2 {
        t.Fatalf("expected 2 groups, got %d", len(got))
    }
}
