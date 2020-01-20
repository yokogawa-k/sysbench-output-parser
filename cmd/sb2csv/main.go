package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/urfave/cli"
	"github.com/yokogawa-k/sysbench-output-parser/dbreport"
)

type output struct {
	Version           string  `csv:"version"`
	Threads           uint64  `csv:"threads"`
	TPS               float64 `csv:"tps"`
	QPS               float64 `csv:"qps"`
	Read              uint64  `csv:"read"`
	Write             uint64  `csv:"write"`
	Other             uint64  `csv:"other"`
	Total             uint64  `csv:"total"`
	IgnoredErrors     uint64  `csv:"ignored-errors"`
	Recoonects        uint64  `csv:"recoonects"`
	MinLatency        float64 `csv:"Latency(min)"`
	AvgLatency        float64 `csv:"Latency(avg)"`
	MaxLatency        float64 `csv:"Latency(max)"`
	PercentileLatency string  `csv:"Latency(Percentile)"`
}

func main() {
	app := &cli.App{
		Name:   "sb2csv",
		Action: convertCSV,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func convertCSV(c *cli.Context) error {
	fd, err := os.Open(c.String("file"))
	if err != nil {
		return err
	}

	s, err := dbreport.ParseDBReport(fd)
	if err != nil {
		return err
	}

	pl := fmt.Sprintf("%.2f (%d%%)", *s.Latency.Percentile, *s.Latency.LatencyPct)

	o := []*output{
		{
			Version:           s.Version.VersionNumber,
			Threads:           s.NumThreads,
			TPS:               *s.SQLStatistics.TransactionsInfo.TransactionsPerSecond,
			QPS:               *s.SQLStatistics.Queries.QueriesPerSecond,
			Read:              *s.SQLStatistics.ReadQueries,
			Write:             *s.SQLStatistics.WriteQueries,
			Other:             *s.SQLStatistics.OtherQueries,
			Total:             *s.SQLStatistics.TotalQueries,
			IgnoredErrors:     *s.SQLStatistics.IgnoredErrors.IgnoredErrors,
			Recoonects:        *s.SQLStatistics.Reconnects.Reconnects,
			MinLatency:        *s.Latency.Minimum,
			AvgLatency:        *s.Latency.Average,
			MaxLatency:        *s.Latency.Maximum,
			PercentileLatency: pl,
		},
	}

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '\t'
		return gocsv.NewSafeCSVWriter(writer)
	})

	cs, err := gocsv.MarshalString(o)
	if err != nil {
		return err
	}

	fmt.Println(cs)

	return nil
}
