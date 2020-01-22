package ioreport

type SysbenchInfo struct {
	Version           *VersionInfo         `@@`
	NumThreads        uint64               `"Number" "of" "threads:" @Int`
	ReportInterval    uint64               `"Report" "intermediate" "results" "every" @Int "second" "(" "s" ")"`
	GeneralStatistics *GeneralStatistics   `"General" "statistics:" @@`
	Latency           *LatencyInfo         `"Latency" "(" "ms" ")" ":" @@`
	ThreadsFairness   *ThreadsFairnessInfo `"Threads" "fairness:" @@`
}

// VersionInfo is sysbench version information struct
type VersionInfo struct {
	VersionNumber string   `"sysbench" @VersionNumber`
	AddInfo       []string `"(" @( Ident | VersionNumber )* ")"`
}

// GeneralStatistics is sysbench report general statistcs struct
type GeneralStatistics struct {
	TotalTime   *float64 `"total" "time:" @Float "s"`
	TotalEvents *uint64  `"total" "number" "of" "events:" @Int`
}

// LatencyInfo is sysbench report latency information struct
type LatencyInfo struct {
	Minimum    *float64 `"min:" @Float`
	Average    *float64 `"avg:" @Float`
	Maximum    *float64 `"max:" @Float`
	LatencyPct *uint64  `@Int`
	Percentile *float64 `"th" "percentile:" @Float`
	Sum        *float64 `"sum:" @Float`
}

// ThreadsFairnessInfo is sysbench report threads fairness information struct
type ThreadsFairnessInfo struct {
	Events        *ThreadsFairnessEventsInfo    `"events" "(" "avg/stddev" ")" ":" @@`
	ExecutionTime *ThreadsFairnessExecutionInfo `"execution" "time" "(" "avg/stddev" ")" ":" @@`
}

// ThreadsFairnessEventsInfo is sysbench report events information struct in threads fairness
type ThreadsFairnessEventsInfo struct {
	EventsAverage *float64 `@Float`
	EventsStddev  *float64 `"/" @Float`
}

// ThreadsFairnessExecutionInfo is sysbench report execution information struct in threads fairness
type ThreadsFairnessExecutionInfo struct {
	ExecutionTimeAverage *float64 `@Float`
	ExecutionTimeStddev  *float64 `"/" @Float`
}
