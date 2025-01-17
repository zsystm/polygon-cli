# `polycli dbbench`

> Auto-generated documentation.

## Table of Contents

- [Description](#description)
- [Usage](#usage)
- [Flags](#flags)
- [See Also](#see-also)

## Description

Perform a level/pebble db benchmark

```bash
polycli dbbench [flags]
```

## Usage

This command is meant to give us a sense of the system level
performance for leveldb:

```bash
polycli dbbench --degree-of-parallelism 2 | jq '.' > result.json
```

In many cases, we'll want to emulate the performance characteristics
of `bor` or `geth`. This is the basic IO pattern when `bor` is in sync:

```text
Process Name = bor
     Kbytes              : count     distribution
         0 -> 1          : 0        |                                        |
         2 -> 3          : 0        |                                        |
         4 -> 7          : 10239    |****************                        |
         8 -> 15         : 25370    |****************************************|
        16 -> 31         : 7082     |***********                             |
        32 -> 63         : 1241     |*                                       |
        64 -> 127        : 58       |                                        |
       128 -> 255        : 11       |                                        |
```

This is the IO pattern when `bor` is getting in sync.

```text
Process Name = bor
     Kbytes              : count     distribution
         0 -> 1          : 0        |                                        |
         2 -> 3          : 0        |                                        |
         4 -> 7          : 23089    |*************                           |
         8 -> 15         : 70350    |****************************************|
        16 -> 31         : 11790    |******                                  |
        32 -> 63         : 1193     |                                        |
        64 -> 127        : 204      |                                        |
       128 -> 255        : 271      |                                        |
       256 -> 511        : 1381     |                                        |
```

This gives us a sense of the relative size of the IOPs. We'd also want
to get a sense of the read/write ratio. This is some sample data from
bor while syncing:

```text
12:48:08 loadavg: 5.86 6.22 7.13 16/451 56297

READS  WRITES R_Kb     W_Kb     PATH
307558 1277   4339783  30488    /var/lib/bor/data/bor/chaindata/

12:48:38 loadavg: 6.46 6.32 7.14 3/452 56298

READS  WRITES R_Kb     W_Kb     PATH
309904 946    4399349  26051    /var/lib/bor/data/bor/chaindata/

```

During the same period of time this is what the IO looks like from a
node that's in sync.

```text
12:48:05 loadavg: 1.55 1.85 2.03 18/416 88371

READS  WRITES R_Kb     W_Kb     PATH
124530 488    1437436  12165    /var/lib/bor/data/bor/chaindata/

12:48:35 loadavg: 4.14 2.44 2.22 1/416 88371

READS  WRITES R_Kb     W_Kb     PATH
81282  215    823530   4610     /var/lib/bor/data/bor/chaindata/

```

If we want to simulate `bor` behavior, we can leverage this data to
configure the leveldb benchmark tool.


| Syncing | Reads   | Writes | Read (kb) | Write (kb) | RW Ratio | kb/r | kb/w |
|---------|---------|--------|-----------|------------|----------|------|------|
| TRUE    | 307,558 |  1,277 | 4,339,783 | 30,488     |      241 | 14.1 | 23.9 |
| TRUE    | 309,904 |    946 | 7,399,349 | 26,051     |      328 | 23.9 | 27.5 |
| FALSE   | 124,530 |    488 | 1,437,436 | 12,165     |      255 | 11.5 | 24.9 |
| FALSE   | 51,282  |    215 | 823,530   | 4,610      |      239 | 16.1 | 21.4 |

The number of IOps while syncing is a lot higher. The only other
obvious difference is that the IOp size is a bit larger while syncing
as well.

- Syncing
  - Read Write Ratio - 275:1 
  - Small IOp - 10kb
  - Large IOp - 256kb
  - Small Large Ratio - 10:1
- Synced
  - Read Write Ratio - 250:1
  - Small IOp - 10kb
  - Larg IOp - 32kb
  - Small Large Ratio - 10:1

```text
7:58PM DBG buckets bucket=0 count=9559791821 end=1 start=0
7:58PM DBG buckets bucket=1 count=141033 end=3 start=2
7:58PM DBG buckets bucket=2 count=92899 end=7 start=4
7:58PM DBG buckets bucket=3 count=256655 end=15 start=8
7:58PM DBG buckets bucket=4 count=262589 end=31 start=16
7:58PM DBG buckets bucket=5 count=191353 end=63 start=32
7:58PM DBG buckets bucket=6 count=99519 end=127 start=64
7:58PM DBG buckets bucket=7 count=74161 end=255 start=128
7:58PM DBG buckets bucket=8 count=17426 end=511 start=256
7:58PM DBG buckets bucket=9 count=692 end=1023 start=512
7:58PM DBG buckets bucket=10 count=989 end=2047 start=1024
7:58PM DBG buckets bucket=13 count=1 end=16383 start=8192
7:58PM INF recorded result desc="full scan" testDuration=10381196.479925
7:58PM DBG recorded result result={"Description":"full scan","EndTime":"2023-07-17T19:58:05.396257711Z","OpCount":9557081144,"OpRate":920614.609547304,"StartTime":"2023-07-17T17:05:04.199777776Z","Stats":{"AliveIterators":0,"AliveSnapshots":0,"BlockCache":{"Buckets":2048,"DelCount":259134854,"GrowCount":9,"HitCount":4,"MissCount":262147633,"Nodes":33294,"SetCount":259168148,"ShrinkCount":2,"Size":268427343},"BlockCacheSize":268427343,"FileCache":{"Buckets":16,"DelCount":536037,"GrowCount":0,"HitCount":2,"MissCount":536537,"Nodes":500,"SetCount":536537,"ShrinkCount":0,"Size":500},"IORead":1092651461848,"IOWrite":13032122717,"Level0Comp":0,"LevelDurations":[0,0,546151937,15675194130,100457643600,40581548153,0],"LevelRead":[0,0,45189458,1233235440,8351239571,3376108236,0],"LevelSizes":[0,103263963,1048356844,10484866671,104856767171,180600915234,797187827055],"LevelTablesCounts":[0,51,665,7066,53522,95777,371946],"LevelWrite":[0,0,45159786,1230799439,8328970986,3371359447,0],"MemComp":0,"NonLevel0Comp":1433,"OpenedTablesCount":500,"SeekComp":0,"WriteDelayCount":0,"WriteDelayDuration":0,"WritePaused":false},"TestDuration":10381196479925,"ValueDist":null}

```

## Flags

```bash
      --cache-size int                the number of megabytes to use as our internal cache size (default 512)
      --db-mode string                The mode to use: leveldb or pebbledb (default "leveldb")
      --db-path string                the path of the database that we'll use for testing (default "_benchmark_db")
      --degree-of-parallelism uint8   The number of concurrent goroutines we'll use (default 2)
      --dont-fill-read-cache          if false, then random reads will be cached
      --full-scan-mode                if true, the application will scan the full database as fast as possible and print a summary
      --handles int                   defines the capacity of the open files caching. Use -1 for zero, this has same effect as specifying NoCacher to OpenFilesCacher. (default 500)
  -h, --help                          help for dbbench
      --key-size uint                 The byte length of the keys that we'll use (default 32)
      --nil-read-opts                 if true we'll use nil read opt (this is what geth/bor does)
      --no-merge-write                allows disabling write merge
      --overwrite-count uint          the number of times to overwrite the data (default 5)
      --read-limit uint               the number of reads will attempt to complete in a given test (default 10000000)
      --read-only                     if true, we'll skip all the write operations and open the DB in read only mode
      --read-strict                   if true the rand reads will be made in strict mode
      --sequential-reads              if true we'll perform reads sequentially
      --sequential-writes             if true we'll perform writes in somewhat sequential manner
      --size-distribution string      the size distribution to use while testing (default "0-1:2347864,2-3:804394856,4-7:541267689,8-15:738828593,16-31:261122372,32-63:1063470933,64-127:3584745195,128-255:1605760137,256-511:316074206,512-1023:312887514,1024-2047:328894149,2048-4095:141180,4096-8191:92789,8192-16383:256060,16384-32767:261806,32768-65535:191032,65536-131071:99715,131072-262143:73782,262144-524287:17552,524288-1048575:717,1048576-2097151:995,2097152-4194303:1,8388608-16777215:1")
      --sync-writes                   sync each write
      --write-limit uint              The number of entries to write in the db (default 1000000)
      --write-zero                    if true, we'll write 0s rather than random data
```

The command also inherits flags from parent commands.

```bash
      --config string   config file (default is $HOME/.polygon-cli.yaml)
      --pretty-logs     Should logs be in pretty format or JSON (default true)
  -v, --verbosity int   0 - Silent
                        100 Panic
                        200 Fatal
                        300 Error
                        400 Warning
                        500 Info
                        600 Debug
                        700 Trace (default 500)
```

## See also

- [polycli](polycli.md) - A Swiss Army knife of blockchain tools.
