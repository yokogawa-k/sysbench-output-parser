
### Reference

- [tsuna/sysbench\-tools: Tools to run sysbench, parse and graph its output](https://github.com/tsuna/sysbench-tools)
- [Sysbench test definition comparison \- eLinux\.org](https://elinux.org/Sysbench_test_definition_comparison#parser.py)

### Sample

#### Original Report

```text
sysbench 1.0.19 (using bundled LuaJIT 2.1.0-beta2)

Running the test with following options:
Number of threads: 8
Report intermediate results every 5 second(s)
Initializing random number generator from current time


Initializing worker threads...

Threads started!

[ 5s ] thds: 8 tps: 355.03 qps: 9952.34 (r/w/o: 4550.27/4687.20/714.87) lat (ms,95%): 48.34 err/s 1.20 reconn/s: 0.00
[ 10s ] thds: 8 tps: 338.41 qps: 9848.13 (r/w/o: 4495.75/4675.55/676.82) lat (ms,95%): 50.11 err/s 1.00 reconn/s: 0.00
[ 15s ] thds: 8 tps: 353.59 qps: 9829.03 (r/w/o: 4476.32/4645.52/707.19) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
[ 20s ] thds: 8 tps: 346.20 qps: 9940.86 (r/w/o: 4534.03/4714.43/692.40) lat (ms,95%): 50.11 err/s 2.00 reconn/s: 0.00
[ 25s ] thds: 8 tps: 343.80 qps: 9930.48 (r/w/o: 4541.84/4701.04/687.61) lat (ms,95%): 49.21 err/s 1.00 reconn/s: 0.00
[ 30s ] thds: 8 tps: 344.60 qps: 9962.88 (r/w/o: 4548.15/4725.55/689.19) lat (ms,95%): 49.21 err/s 1.80 reconn/s: 0.00
[ 35s ] thds: 8 tps: 338.40 qps: 9860.13 (r/w/o: 4509.06/4674.26/676.81) lat (ms,95%): 49.21 err/s 1.60 reconn/s: 0.00
[ 40s ] thds: 8 tps: 345.80 qps: 9833.00 (r/w/o: 4494.60/4646.80/691.60) lat (ms,95%): 49.21 err/s 2.00 reconn/s: 0.00
[ 45s ] thds: 8 tps: 345.40 qps: 10023.94 (r/w/o: 4573.77/4759.77/690.40) lat (ms,95%): 51.02 err/s 1.60 reconn/s: 0.00
[ 50s ] thds: 8 tps: 339.40 qps: 9883.44 (r/w/o: 4508.22/4696.02/679.20) lat (ms,95%): 51.02 err/s 0.40 reconn/s: 0.00
[ 55s ] thds: 8 tps: 343.40 qps: 9899.55 (r/w/o: 4523.98/4688.77/686.80) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
[ 60s ] thds: 8 tps: 344.59 qps: 9953.51 (r/w/o: 4539.47/4724.86/689.18) lat (ms,95%): 51.02 err/s 2.20 reconn/s: 0.00
[ 65s ] thds: 8 tps: 347.01 qps: 10172.29 (r/w/o: 4652.93/4825.54/693.82) lat (ms,95%): 49.21 err/s 1.80 reconn/s: 0.00
[ 70s ] thds: 8 tps: 351.40 qps: 9908.55 (r/w/o: 4519.18/4686.37/703.00) lat (ms,95%): 50.11 err/s 1.60 reconn/s: 0.00
[ 75s ] thds: 8 tps: 344.00 qps: 9763.47 (r/w/o: 4469.43/4606.03/688.01) lat (ms,95%): 49.21 err/s 1.20 reconn/s: 0.00
[ 80s ] thds: 8 tps: 346.20 qps: 10005.87 (r/w/o: 4576.43/4737.03/692.40) lat (ms,95%): 49.21 err/s 1.00 reconn/s: 0.00
[ 85s ] thds: 8 tps: 352.39 qps: 9785.99 (r/w/o: 4448.30/4632.90/704.78) lat (ms,95%): 50.11 err/s 1.60 reconn/s: 0.00
[ 90s ] thds: 8 tps: 341.80 qps: 9731.60 (r/w/o: 4445.20/4602.80/683.60) lat (ms,95%): 50.11 err/s 0.60 reconn/s: 0.00
[ 95s ] thds: 8 tps: 345.60 qps: 9823.52 (r/w/o: 4477.56/4654.76/691.19) lat (ms,95%): 51.02 err/s 1.20 reconn/s: 0.00
[ 100s ] thds: 8 tps: 350.01 qps: 10066.89 (r/w/o: 4593.33/4773.54/700.02) lat (ms,95%): 51.02 err/s 1.40 reconn/s: 0.00
[ 105s ] thds: 8 tps: 343.40 qps: 9898.00 (r/w/o: 4527.00/4684.20/686.80) lat (ms,95%): 50.11 err/s 1.40 reconn/s: 0.00
[ 110s ] thds: 8 tps: 345.00 qps: 9906.08 (r/w/o: 4507.75/4708.34/689.99) lat (ms,95%): 51.94 err/s 1.40 reconn/s: 0.00
[ 115s ] thds: 8 tps: 340.00 qps: 10084.30 (r/w/o: 4606.85/4797.45/680.01) lat (ms,95%): 52.89 err/s 3.60 reconn/s: 0.00
[ 120s ] thds: 8 tps: 347.20 qps: 9967.96 (r/w/o: 4534.78/4738.78/694.40) lat (ms,95%): 50.11 err/s 0.80 reconn/s: 0.00
[ 125s ] thds: 8 tps: 345.40 qps: 9977.85 (r/w/o: 4551.02/4736.02/690.80) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
[ 130s ] thds: 8 tps: 344.20 qps: 9928.38 (r/w/o: 4528.79/4711.19/688.40) lat (ms,95%): 50.11 err/s 0.60 reconn/s: 0.00
[ 135s ] thds: 8 tps: 335.60 qps: 9772.11 (r/w/o: 4447.65/4653.25/671.21) lat (ms,95%): 52.89 err/s 0.40 reconn/s: 0.00
[ 140s ] thds: 8 tps: 354.20 qps: 9857.32 (r/w/o: 4497.36/4651.56/708.39) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
[ 145s ] thds: 8 tps: 341.20 qps: 9858.97 (r/w/o: 4508.59/4667.99/682.40) lat (ms,95%): 51.02 err/s 1.80 reconn/s: 0.00
[ 150s ] thds: 8 tps: 347.80 qps: 9979.07 (r/w/o: 4552.23/4731.23/695.60) lat (ms,95%): 50.11 err/s 1.40 reconn/s: 0.00
[ 155s ] thds: 8 tps: 348.80 qps: 9841.54 (r/w/o: 4496.37/4647.57/697.60) lat (ms,95%): 50.11 err/s 1.00 reconn/s: 0.00
[ 160s ] thds: 8 tps: 343.80 qps: 9995.63 (r/w/o: 4564.41/4743.61/687.60) lat (ms,95%): 51.02 err/s 1.40 reconn/s: 0.00
[ 165s ] thds: 8 tps: 340.39 qps: 9823.99 (r/w/o: 4479.31/4663.90/680.79) lat (ms,95%): 51.94 err/s 0.40 reconn/s: 0.00
[ 170s ] thds: 8 tps: 351.00 qps: 9927.46 (r/w/o: 4515.74/4709.74/701.99) lat (ms,95%): 51.02 err/s 1.40 reconn/s: 0.00
[ 175s ] thds: 8 tps: 352.61 qps: 10050.58 (r/w/o: 4594.17/4751.18/705.23) lat (ms,95%): 50.11 err/s 1.00 reconn/s: 0.00
[ 180s ] thds: 8 tps: 342.20 qps: 9839.88 (r/w/o: 4482.95/4672.54/684.39) lat (ms,95%): 51.02 err/s 1.00 reconn/s: 0.00
[ 185s ] thds: 8 tps: 342.00 qps: 9843.80 (r/w/o: 4496.40/4663.40/684.00) lat (ms,95%): 51.02 err/s 1.80 reconn/s: 0.00
[ 190s ] thds: 8 tps: 348.00 qps: 9844.31 (r/w/o: 4493.56/4654.76/695.99) lat (ms,95%): 49.21 err/s 1.40 reconn/s: 0.00
[ 195s ] thds: 8 tps: 331.20 qps: 9917.51 (r/w/o: 4531.76/4723.36/662.39) lat (ms,95%): 51.94 err/s 1.20 reconn/s: 0.00
[ 200s ] thds: 8 tps: 356.60 qps: 10070.77 (r/w/o: 4604.79/4752.79/713.20) lat (ms,95%): 48.34 err/s 1.40 reconn/s: 0.00
[ 205s ] thds: 8 tps: 350.99 qps: 9732.01 (r/w/o: 4443.91/4586.11/701.99) lat (ms,95%): 49.21 err/s 1.40 reconn/s: 0.00
[ 210s ] thds: 8 tps: 336.40 qps: 9982.91 (r/w/o: 4574.45/4735.65/672.81) lat (ms,95%): 51.02 err/s 2.20 reconn/s: 0.00
[ 215s ] thds: 8 tps: 341.81 qps: 9961.22 (r/w/o: 4549.70/4727.91/683.62) lat (ms,95%): 51.94 err/s 2.00 reconn/s: 0.00
[ 220s ] thds: 8 tps: 333.19 qps: 9929.60 (r/w/o: 4525.82/4737.61/666.17) lat (ms,95%): 51.94 err/s 1.60 reconn/s: 0.00
[ 225s ] thds: 8 tps: 349.21 qps: 9941.46 (r/w/o: 4534.72/4708.12/698.62) lat (ms,95%): 50.11 err/s 0.80 reconn/s: 0.00
[ 230s ] thds: 8 tps: 349.40 qps: 9884.26 (r/w/o: 4497.43/4688.03/698.80) lat (ms,95%): 51.94 err/s 1.60 reconn/s: 0.00
[ 235s ] thds: 8 tps: 351.00 qps: 9885.45 (r/w/o: 4512.02/4671.42/702.00) lat (ms,95%): 51.02 err/s 2.20 reconn/s: 0.00
[ 240s ] thds: 8 tps: 350.00 qps: 9855.94 (r/w/o: 4492.37/4663.57/700.00) lat (ms,95%): 50.11 err/s 1.00 reconn/s: 0.00
[ 245s ] thds: 8 tps: 351.81 qps: 10002.18 (r/w/o: 4570.08/4728.49/703.61) lat (ms,95%): 49.21 err/s 2.20 reconn/s: 0.00
[ 250s ] thds: 8 tps: 340.80 qps: 9721.26 (r/w/o: 4429.43/4610.23/681.60) lat (ms,95%): 49.21 err/s 2.20 reconn/s: 0.00
[ 255s ] thds: 8 tps: 354.00 qps: 9868.69 (r/w/o: 4507.64/4653.04/708.01) lat (ms,95%): 50.11 err/s 1.80 reconn/s: 0.00
[ 260s ] thds: 8 tps: 342.80 qps: 9965.06 (r/w/o: 4541.94/4737.54/685.59) lat (ms,95%): 51.02 err/s 0.80 reconn/s: 0.00
[ 265s ] thds: 8 tps: 346.80 qps: 10065.20 (r/w/o: 4606.00/4765.60/693.60) lat (ms,95%): 49.21 err/s 1.80 reconn/s: 0.00
[ 270s ] thds: 8 tps: 350.20 qps: 9863.69 (r/w/o: 4483.64/4679.64/700.41) lat (ms,95%): 51.02 err/s 1.40 reconn/s: 0.00
[ 275s ] thds: 8 tps: 353.60 qps: 9952.75 (r/w/o: 4552.98/4692.58/707.20) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
[ 280s ] thds: 8 tps: 348.60 qps: 9790.18 (r/w/o: 4463.39/4629.59/697.20) lat (ms,95%): 51.02 err/s 2.20 reconn/s: 0.00
[ 285s ] thds: 8 tps: 344.80 qps: 9759.55 (r/w/o: 4453.18/4616.78/689.60) lat (ms,95%): 51.02 err/s 2.20 reconn/s: 0.00
[ 290s ] thds: 8 tps: 355.80 qps: 10017.66 (r/w/o: 4563.43/4742.63/711.60) lat (ms,95%): 49.21 err/s 1.20 reconn/s: 0.00
[ 295s ] thds: 8 tps: 339.80 qps: 9893.85 (r/w/o: 4533.42/4680.82/679.60) lat (ms,95%): 50.11 err/s 2.20 reconn/s: 0.00
[ 300s ] thds: 8 tps: 356.00 qps: 10030.59 (r/w/o: 4583.39/4735.19/712.00) lat (ms,95%): 50.11 err/s 1.20 reconn/s: 0.00
SQL statistics:
    queries performed:
        read:                            1357163
        write:                           1408638
        other:                           207680
        total:                           2973481
    transactions:                        103832 (346.07 per sec.)
    queries:                             2973481 (9910.63 per sec.)
    ignored errors:                      433    (1.44 per sec.)
    reconnects:                          0      (0.00 per sec.)

General statistics:
    total time:                          300.0278s
    total number of events:              103832

Latency (ms):
         min:                                    1.28
         avg:                                   23.11
         max:                                  133.17
         95th percentile:                       50.11
         sum:                              2399718.86

Threads fairness:
    events (avg/stddev):           12979.0000/64.33
    execution time (avg/stddev):   299.9649/0.01

```

#### CSV Output

```csv
version	threads	tps	qps	read	write	other	total	ignored-errors	recoonects	Latency(min)	Latency(avg)	Latency(max)	Latency(Percentile)
1.0.19	8	346.07	9910.63	1357163	1408638	207680	2973481	433	0	1.28	23.11	133.17	50.11 (95%)
```

