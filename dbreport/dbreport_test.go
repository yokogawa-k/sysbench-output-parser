package dbreport

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadStringError(t *testing.T) {
	b := &badStringError{
		what: "foo",
		str:  "bar",
	}

	expected := "foo \"bar\""
	actual := b.Error()
	assert.Equal(t, expected, actual)
}

func TestParseDBReport(t *testing.T) {
	dummy := `sysbench 1.0.19 (using bundled LuaJIT 2.1.0-beta2)

Running the test with following options:
Number of threads: 8
Report intermediate results every 5 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 5s ] thds: 8 tps: 3.03 qps: 12.34 (r/w/o: 12.34/0.00/0.00) lat (ms,95%): 5.00 err/s 0.00 reconn/s: 0.00
[ 10s ] thds: 8 tps: 5.05 qps: 23.45 (r/w/o: 2.34/5.67/12.89) lat (ms,95%): 10.11 err/s 2.55 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            73
        write:                           28
        other:                           65
        total:                           166
    transactions:                        40     (4.00 per sec.)
    queries:                             179    (17.90 per sec.)
    ignored errors:                      13     (1.30 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          10.0278s
    total number of events:              40

Latency (ms):
         min:                                    1.28
         avg:                                    3.11
         max:                                   33.17
         95th percentile:                        8.11
         sum:                                  124.40

Threads fairness:
    events (avg/stddev):           5.0000/64.33
    execution time (avg/stddev):   9.9987/0.01

`
	expected := &SysbenchInfo{
		Version: &VersionInfo{
			VersionNumber: "1.0.19",
			AddInfo:       []string{"using", "bundled", "LuaJIT", "2.1.0-beta2"},
		},
		NumThreads:     8,
		ReportInterval: 5,
		Reports: []*DBReport{
			{
				ElapsedTime:    uint64Ptr(5),
				ThreadsRunning: uint64Ptr(8),
				TPS:            float64Ptr(3.03),
				QPS:            float64Ptr(12.34),
				ReadQuery:      float64Ptr(12.34),
				WriteQuery:     float64Ptr(0.00),
				OtherQuery:     float64Ptr(0.00),
				LatencyPct:     uint64Ptr(95),
				Latency:        float64Ptr(5.00),
				Errors:         float64Ptr(0.00),
				Reconnect:      float64Ptr(0.00),
			},
			{
				ElapsedTime:    uint64Ptr(10),
				ThreadsRunning: uint64Ptr(8),
				TPS:            float64Ptr(5.05),
				QPS:            float64Ptr(23.45),
				ReadQuery:      float64Ptr(2.34),
				WriteQuery:     float64Ptr(5.67),
				OtherQuery:     float64Ptr(12.89),
				LatencyPct:     uint64Ptr(95),
				Latency:        float64Ptr(10.11),
				Errors:         float64Ptr(2.55),
				Reconnect:      float64Ptr(0.00),
			},
		},
		SQLStatistics: &SQLStatistics{
			ReadQueries:  uint64Ptr(73),
			WriteQueries: uint64Ptr(28),
			OtherQueries: uint64Ptr(65),
			TotalQueries: uint64Ptr(166),
			TransactionsInfo: &SQLStatisticsTransactions{
				Transactions:          uint64Ptr(40),
				TransactionsPerSecond: float64Ptr(4.00),
			},
			Queries: &SQLStatisticsQueries{
				Queries:          uint64Ptr(179),
				QueriesPerSecond: float64Ptr(17.90),
			},
			IgnoredErrors: &SQLStatisticsIgnoredErrors{
				IgnoredErrors:          uint64Ptr(13),
				IgnoredErrorsPerSecond: float64Ptr(1.30),
			},
			Reconnects: &SQLStatisticsReconnects{
				Reconnects:          uint64Ptr(0),
				ReconnectsPerSecond: float64Ptr(0.00),
			},
		},
		GeneralStatistics: &GeneralStatistics{
			TotalTime:   float64Ptr(10.0278),
			TotalEvents: uint64Ptr(40),
		},
		Latency: &LatencyInfo{
			Minimum:    float64Ptr(1.28),
			Average:    float64Ptr(3.11),
			Maximum:    float64Ptr(33.17),
			LatencyPct: uint64Ptr(95),
			Percentile: float64Ptr(8.11),
			Sum:        float64Ptr(124.40),
		},
		ThreadsFairness: &ThreadsFairnessInfo{
			Events: &ThreadsFairnessEventsInfo{
				EventsAverage: float64Ptr(5.0000),
				EventsStddev:  float64Ptr(64.33),
			},
			ExecutionTime: &ThreadsFairnessExecutionInfo{
				ExecutionTimeAverage: float64Ptr(9.9987),
				ExecutionTimeStddev:  float64Ptr(0.01),
			},
		},
	}

	r := strings.NewReader(dummy)
	actual, err := ParseDBReport(r)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestParseDBReportFail(t *testing.T) {
	dummy := `sysbench dummy.version`
	r := strings.NewReader(dummy)
	_, err := ParseDBReport(r)
	assert.Error(t, err)
}

func float64Ptr(v float64) *float64 {
	return &v
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}
