package dbreport

// SysbenchInfo is sysbench report information struct
type SysbenchInfo struct {
	Version           *VersionInfo         `@@`
	NumThreads        uint64               `"Number" "of" "threads:" @Int`
	ReportInterval    uint64               `"Report" "intermediate" "results" "every" @Int "second" "(" "s" ")"`
	Reports           []*DBReport          `("[" @@ )*`
	SQLStatistics     *SQLStatistics       `"SQL" "statistics:" "queries" "performed:" @@`
	GeneralStatistics *GeneralStatistics   `"General" "statistics:" @@`
	Latency           *LatencyInfo         `"Latency" "(" "ms" ")" ":" @@`
	ThreadsFairness   *ThreadsFairnessInfo `"Threads" "fairness:" @@`
}

// VersionInfo is sysbench version information struct
type VersionInfo struct {
	VersionNumber string   `"sysbench" @VersionNumber`
	AddInfo       []string `"(" @( Ident | VersionNumber )* ")"`
}

// DBReport is sysbench dbreport struct
type DBReport struct {
	ElapsedTime    *uint64  `@Int "s" "]"`
	ThreadsRunning *uint64  `"thds:" @Int`
	TPS            *float64 `"tps:" @Float`
	QPS            *float64 `"qps:" @Float`
	ReadQuery      *float64 `"(" "r/w/o:" @Float`
	WriteQuery     *float64 `"/" @Float`
	OtherQuery     *float64 `"/" @Float ")"`
	LatencyPct     *uint64  `"lat" LetancyPictPrefix @Int LetancyPictSuffix`
	Latency        *float64 `@Float`
	Errors         *float64 `"err/s" @Float`
	Reconnect      *float64 `"reconn/s:" @Float`
}

// SQLStatistics is sysbench report sql statistics struct
type SQLStatistics struct {
	ReadQueries      *uint64                     `"read:" @Int`
	WriteQueries     *uint64                     `"write:" @Int`
	OtherQueries     *uint64                     `"other:" @Int`
	TotalQueries     *uint64                     `"total:" @Int`
	TransactionsInfo *SQLStatisticsTransactions  `"transactions:" ( @@ )`
	Queries          *SQLStatisticsQueries       `"queries:" ( @@ )`
	IgnoredErrors    *SQLStatisticsIgnoredErrors `"ignored" "errors:" ( @@ )`
	Reconnects       *SQLStatisticsReconnects    `"reconnects:" ( @@ )`
}

// SQLStatisticsTransactions is sysbench report transactions struct in sql statistics
type SQLStatisticsTransactions struct {
	Transactions          *uint64  `@Int`
	TransactionsPerSecond *float64 `"(" @Float "per" "sec." ")"`
}

// SQLStatisticsQueries is sysbench report queries struct in sql statistics
type SQLStatisticsQueries struct {
	Queries          *uint64  `@Int`
	QueriesPerSecond *float64 `"(" @Float "per" "sec." ")"`
}

// SQLStatisticsIgnoredErrors is sysbench report ignore errors struct in sql statistics
type SQLStatisticsIgnoredErrors struct {
	IgnoredErrors          *uint64  `@Int`
	IgnoredErrorsPerSecond *float64 `"(" @Float "per" "sec." ")"`
}

// SQLStatisticsReconnects is sysbench report reconnects struct in sql statistics
type SQLStatisticsReconnects struct {
	Reconnects          *uint64  `@Int`
	ReconnectsPerSecond *float64 `"(" @Float "per" "sec." ")"`
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
