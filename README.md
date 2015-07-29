# Benchmark Go Channel Performance

## How?

1 producer goroutine and P consumer goroutines with P * T channel
send/receive operations in total. Each consumer goroutine also sleeps 1
microsecond after receiving an item.

Run as:

    git clone https://github.com/shishkander/go-channel-benchmark.git
    cd go-channel-benchmark
    # To vary parameters, do:
    # $TEXTEDITOR experimental.go
    go run experimental.go

### How to read result tables

* **Columns** are number of channel receive operations per goroutine. The
*none* column shows the cost of setting up and dismantling of goroutines, which
is insignificant in other columns.
* **Rows** are sizes of the channel.
* **Values** is the time in microseconds per 1 send/receive operation.  numbers
are not important though, only their relation.
* **GOMAXPROCS** means how many goroutines can really run in parallel.
* **Consumers** is the number of consuming goroutines.

## Results

### GOMAXPROCS = 04, Producers = 01, Consumers = 01

    |===================================================================|
    | buf\tasks |  none |    50 |   100 |   150 |   250 |   400 |   800 |
    |===================================================================|
    | unbuf     |   5.2 |  18.7 |  18.6 |  18.3 |  18.5 |  18.4 |  18.1 |
    |    50     |   4.8 |  50.8 |  38.6 |  30.2 |  24.4 |  19.9 |  17.1 |
    |   100     |   5.2 |  51.3 |  57.5 |  45.1 |  34.1 |  27.9 |  21.5 |
    |   150     |   5.1 |  51.2 |  57.1 |  57.3 |  42.2 |  36.0 |  24.6 |
    |   250     |   4.9 |  53.0 |  57.4 |  57.5 |  59.9 |  44.5 |  30.3 |
    |   400     |   5.1 |  52.5 |  56.7 |  57.2 |  57.8 |  61.6 |  39.2 |
    |   800     |   5.0 |  50.9 |  57.3 |  58.4 |  60.5 |  58.9 |  59.1 |
    |===================================================================|

### GOMAXPROCS = 04, Producers = 01, Consumers = 02

    |===================================================================|
    | buf\tasks |  none |   100 |   200 |   300 |   500 |   800 |  1600 |
    |===================================================================|
    | unbuf     |   4.3 |  14.7 |  13.5 |  13.4 |  13.4 |  13.4 |  13.1 |
    |    50     |   5.6 |  17.8 |  14.3 |  12.7 |  10.9 |   9.1 |   7.5 |
    |   100     |   5.8 |  21.6 |  17.0 |  14.7 |  12.2 |   9.7 |   7.6 |
    |   150     |   5.9 |  21.4 |  19.6 |  17.0 |  13.5 |  11.1 |   8.5 |
    |   250     |   5.9 |  21.5 |  21.4 |  20.1 |  16.1 |  12.7 |   9.6 |
    |   400     |   5.7 |  21.5 |  20.9 |  21.4 |  20.1 |  15.7 |  10.7 |
    |   800     |   6.0 |  21.7 |  22.0 |  21.2 |  21.0 |  21.3 |  14.6 |
    |===================================================================|

### GOMAXPROCS = 04, Producers = 01, Consumers = 03

    |===================================================================|
    | buf\tasks |  none |   150 |   300 |   450 |   750 |  1200 |  2400 |
    |===================================================================|
    | unbuf     |   6.5 |  13.4 |  13.2 |  13.3 |  13.6 |  13.7 |  13.5 |
    |    50     |   6.1 |  12.6 |  10.3 |   9.1 |   8.2 |   6.9 |   6.2 |
    |   100     |   6.2 |  14.4 |  11.8 |  10.2 |   8.9 |   7.0 |   6.4 |
    |   150     |   6.6 |  16.4 |  12.9 |  11.3 |   9.2 |   7.3 |   6.5 |
    |   250     |   6.6 |  15.5 |  14.7 |  12.9 |  10.1 |   8.6 |   6.6 |
    |   400     |   6.5 |  15.4 |  15.2 |  14.8 |  12.1 |   9.9 |   7.4 |
    |   800     |   6.4 |  16.2 |  15.9 |  15.9 |  15.9 |  13.5 |   9.5 |
    |===================================================================|


### GOMAXPROCS = 03, Producers = 01, Consumers = 01

    |===================================================================|
    | buf\tasks |  none |    50 |   100 |   150 |   250 |   400 |   800 |
    |===================================================================|
    | unbuf     |   4.2 |  16.8 |  15.4 |  16.8 |  16.7 |  17.4 |  16.2 |
    |    50     |   4.7 |  51.7 |  37.6 |  28.8 |  23.8 |  19.9 |  15.9 |
    |   100     |   4.6 |  51.5 |  54.5 |  43.5 |  32.4 |  25.1 |  19.8 |
    |   150     |   4.7 |  51.0 |  54.3 |  58.1 |  43.7 |  31.0 |  20.8 |
    |   250     |   4.7 |  50.7 |  54.2 |  56.1 |  59.4 |  43.9 |  30.2 |
    |   400     |   5.0 |  52.2 |  57.1 |  58.7 |  59.6 |  61.1 |  40.7 |
    |   800     |   5.0 |  52.2 |  57.5 |  58.7 |  60.8 |  61.1 |  61.8 |
    |===================================================================|

### GOMAXPROCS = 03, Producers = 01, Consumers = 02

    |===================================================================|
    | buf\tasks |  none |   100 |   200 |   300 |   500 |   800 |  1600 |
    |===================================================================|
    | unbuf     |   5.4 |  14.3 |  14.0 |  13.8 |  13.8 |  14.3 |  13.8 |
    |    50     |   4.7 |  18.2 |  15.7 |  14.4 |  13.0 |  11.9 |  10.9 |
    |   100     |   5.6 |  21.3 |  17.7 |  15.9 |  14.0 |  13.4 |  10.9 |
    |   150     |   5.5 |  21.5 |  19.4 |  17.5 |  15.0 |  13.4 |  11.5 |
    |   250     |   5.5 |  21.2 |  20.9 |  20.0 |  17.4 |  14.5 |  12.5 |
    |   400     |   5.5 |  21.3 |  21.1 |  21.1 |  19.9 |  17.0 |  13.3 |
    |   800     |   5.5 |  21.5 |  22.0 |  21.5 |  21.1 |  20.9 |  16.7 |
    |===================================================================|

### GOMAXPROCS = 03, Producers = 01, Consumers = 03

    |===================================================================|
    | buf\tasks |  none |   150 |   300 |   450 |   750 |  1200 |  2400 |
    |===================================================================|
    | unbuf     |   5.9 |  13.6 |  13.5 |  13.5 |  13.4 |  13.5 |  13.5 |
    |    50     |   5.6 |  13.4 |  11.9 |  11.6 |  10.8 |   9.3 |   9.5 |
    |   100     |   6.0 |  14.7 |  12.9 |  12.2 |  10.8 |   9.7 |   9.6 |
    |   150     |   6.0 |  15.5 |  14.2 |  12.6 |  11.3 |  10.5 |   9.4 |
    |   250     |   6.0 |  15.7 |  15.0 |  13.7 |  12.2 |  10.8 |   9.2 |
    |   400     |   5.9 |  15.6 |  15.7 |  15.2 |  13.5 |  12.0 |   9.9 |
    |   800     |   5.8 |  15.5 |  16.1 |  15.8 |  15.5 |  14.2 |  12.1 |
    |===================================================================|

### GOMAXPROCS = 02, Producers = 01, Consumers = 01

    |===================================================================|
    | buf\tasks |  none |    50 |   100 |   150 |   250 |   400 |   800 |
    |===================================================================|
    | unbuf     |   4.2 |  20.3 |  19.9 |  19.4 |  20.1 |  19.4 |  20.6 |
    |    50     |   4.4 |  53.6 |  38.4 |  32.5 |  27.2 |  23.4 |  21.0 |
    |   100     |   4.4 |  54.2 |  58.8 |  46.0 |  35.8 |  29.4 |  23.5 |
    |   150     |   4.2 |  55.0 |  57.9 |  60.9 |  44.2 |  34.9 |  26.7 |
    |   250     |   4.4 |  55.8 |  59.8 |  59.7 |  61.7 |  44.9 |  31.6 |
    |   400     |   4.3 |  55.4 |  58.1 |  59.2 |  60.8 |  61.4 |  40.0 |
    |   800     |   4.4 |  55.8 |  60.7 |  61.5 |  62.3 |  60.9 |  61.1 |
    |===================================================================|

### GOMAXPROCS = 02, Producers = 01, Consumers = 02

    |===================================================================|
    | buf\tasks |  none |   100 |   200 |   300 |   500 |   800 |  1600 |
    |===================================================================|
    | unbuf     |   4.7 |  19.7 |  18.7 |  18.8 |  18.6 |  18.4 |  18.6 |
    |    50     |   4.9 |  17.8 |  14.6 |  13.2 |  12.2 |  11.8 |  11.3 |
    |   100     |   5.0 |  23.8 |  17.8 |  15.0 |  13.2 |  12.5 |  11.6 |
    |   150     |   4.8 |  24.0 |  20.8 |  18.3 |  14.7 |  13.1 |  12.1 |
    |   250     |   4.8 |  24.0 |  23.8 |  21.9 |  17.4 |  15.0 |  13.0 |
    |   400     |   5.2 |  23.9 |  23.9 |  23.8 |  21.3 |  17.7 |  14.2 |
    |   800     |   5.1 |  24.1 |  23.8 |  23.8 |  23.9 |  23.8 |  17.4 |
    |===================================================================|

### GOMAXPROCS = 02, Producers = 01, Consumers = 03

    |===================================================================|
    | buf\tasks |  none |   150 |   300 |   450 |   750 |  1200 |  2400 |
    |===================================================================|
    | unbuf     |   5.5 |  14.2 |  14.1 |  13.9 |  13.9 |  13.9 |  13.9 |
    |    50     |   5.8 |  14.8 |  13.9 |  13.8 |  13.6 |  13.2 |  12.1 |
    |   100     |   5.6 |  15.3 |  14.4 |  14.0 |  13.6 |  12.6 |  12.9 |
    |   150     |   5.9 |  16.1 |  14.9 |  14.3 |  13.8 |  13.4 |  13.0 |
    |   250     |   5.8 |  16.0 |  15.6 |  15.0 |  14.2 |  14.0 |  13.5 |
    |   400     |   5.6 |  15.9 |  15.7 |  15.5 |  14.8 |  14.4 |  13.7 |
    |   800     |   5.9 |  15.9 |  15.8 |  15.7 |  15.9 |  15.1 |  14.1 |
    |===================================================================|


### GOMAXPROCS = 01, Producers = 01, Consumers = 01

    |===================================================================|
    | buf\tasks |  none |    50 |   100 |   150 |   250 |   400 |   800 |
    |===================================================================|
    | unbuf     |   1.0 |  15.0 |  15.2 |  15.2 |  14.9 |  15.1 |  14.9 |
    |    50     |   1.0 |  59.6 |  32.0 |  21.2 |  14.2 |  10.6 |   8.4 |
    |   100     |   1.0 |  59.8 |  60.2 |  43.4 |  27.2 |  18.4 |  10.3 |
    |   150     |   1.0 |  60.6 |  59.7 |  61.8 |  37.5 |  24.8 |  14.0 |
    |   250     |   1.0 |  61.0 |  57.4 |  60.8 |  61.0 |  39.1 |  20.7 |
    |   400     |   1.0 |  61.1 |  60.6 |  60.1 |  60.8 |  57.1 |  31.6 |
    |   800     |   1.0 |  60.0 |  60.3 |  57.7 |  60.3 |  58.0 |  57.5 |
    |===================================================================|

### GOMAXPROCS = 01, Producers = 01, Consumers = 02

    |===================================================================|
    | buf\tasks |  none |   100 |   200 |   300 |   500 |   800 |  1600 |
    |===================================================================|
    | unbuf     |   1.2 |  18.6 |  18.4 |  18.5 |  18.3 |  18.4 |  18.5 |
    |    50     |   1.3 |  13.0 |   8.0 |   7.2 |   5.6 |   5.0 |   3.8 |
    |   100     |   1.2 |  23.3 |  12.6 |   9.1 |   6.6 |   5.8 |   4.7 |
    |   150     |   1.2 |  23.4 |  18.0 |  12.7 |   8.3 |   6.5 |   5.2 |
    |   250     |   1.2 |  23.2 |  23.4 |  19.8 |  12.5 |   8.6 |   6.3 |
    |   400     |   1.3 |  23.2 |  23.4 |  23.7 |  19.1 |  12.7 |   7.8 |
    |   800     |   1.2 |  23.9 |  23.5 |  23.4 |  23.2 |  23.3 |  12.6 |
    |===================================================================|

### GOMAXPROCS = 01, Producers = 01, Consumers = 03

    |===================================================================|
    | buf\tasks |  none |   150 |   300 |   450 |   750 |  1200 |  2400 |
    |===================================================================|
    | unbuf     |   1.5 |  12.9 |  12.9 |  12.9 |  12.8 |  12.8 |  12.9 |
    |    50     |   1.5 |   6.9 |   5.6 |   4.2 |   3.5 |   3.3 |   3.0 |
    |   100     |   1.5 |  10.4 |   6.8 |   5.9 |   4.5 |   3.9 |   3.1 |
    |   150     |   1.5 |  15.0 |   8.4 |   6.8 |   4.9 |   4.1 |   3.4 |
    |   250     |   1.5 |  14.8 |  13.1 |   9.0 |   7.0 |   4.8 |   4.1 |
    |   400     |   1.6 |  14.7 |  15.8 |  13.6 |   9.2 |   6.2 |   4.8 |
    |   800     |   1.5 |  14.6 |  14.5 |  14.7 |  14.4 |  10.2 |   6.2 |
    |===================================================================|