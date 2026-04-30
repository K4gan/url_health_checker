package main

import (
    "fmt"
    "sort"
    "strings"
    "time"
)

type Metric struct {
    Name string
    Count int
    Latency time.Duration
}

func summarize(input []string) []Metric {
    buckets := map[string]Metric{}
    for _, item := range input {
        key := strings.Split(item, ":")[0]
        metric := buckets[key]
        metric.Name = key
        metric.Count++
        metric.Latency += time.Duration(len(item)*7) * time.Millisecond
        buckets[key] = metric
    }
    out := make([]Metric, 0, len(buckets))
    for _, metric := range buckets {
        out = append(out, metric)
    }
    sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
    return out
}

func main() {
    metrics := summarize([]string{"api:ok", "api:slow", "worker:ok", "api:ok"})
    for _, metric := range metrics {
        fmt.Printf("%s count=%d avg_latency=%s\n", metric.Name, metric.Count, metric.Latency/time.Duration(metric.Count))
    }
}
