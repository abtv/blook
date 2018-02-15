# blook

[![Build Status](https://travis-ci.org/abtv/blook.svg?branch=master)](https://travis-ci.org/abtv/blook)

## What is it?

**blook** is a tool developed to save your time during text file analysis.
It makes very fast binary **range search by prefix** in sorted files.

## When to use `blook`

When you need to find something in big logs, time series data, etc.

You can say there are **sed**, **awk** and **look**. There are reasons against them:

* **sed** and **awk** are slow because they make linear processing. If files are sorted we should use binary search instead. Please, check `Benchmarks` section for more info.

* we could use Unix **look** utility, but it has 2 disadvantages:
  - it doesn't support range search (for example, find all the lines between `2018-01-13T10:10` and `2018-01-13T10:12`)
  - it doesn't work with very big files (because it works via mmap)

## When not to use `blook`

* if you need to search in non-sorted files, just use **sed** or **awk**
* if you need only prefix search **without** range search **and** your files can be mmap-ed, just use **look** (because it's already included in any Unix system)

## Usage

`blook from_pattern to_pattern file1.log [file2.log ...]`

## Example

Suppose you have log files with the following content:

```
...
2018-01-13T10:10:23+00:00 fetch user with id = 100500
2018-01-13T10:10:23+00:00 check if user with id = 100500 is eligible to some cool stuff
2018-01-13T10:10:24+00:00 send email to user with id = 100500
...
```

The following command returns all the lines from `/var/log/remote/your_service.log` file between `2018-01-13T10:10:23+00:00` and `2018-01-13T10:20:23+00:00` timestamps:

`./blook 2018-01-13T10:10:23+00:00 2018-01-13T10:20:23+00:00 /var/log/remote/your_service.log`

## Build

Build for MacOS:

```
env GOOS=darwin go build
```

Build for Linux:
```
env GOOS=linux go build
```

## Benchmarks

Just run ./benchmark.sh . Be careful: benchmark will create a file about 260MB.

My results (blook works at least **1500 times faster** for this file):

```

blook results:

real	0m0.012s
user	0m0.002s
sys	0m0.010s
----------------
sed results:

real	0m18.939s
user	0m15.280s
sys	0m0.476s
----------------
awk results:

real	0m23.160s
user	0m20.372s
sys	0m0.406s

```

## License

Copyright (c) Andrey Butov. All rights reserved. The use and
distribution terms for this software are covered by the Eclipse
Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
which can be found in the file epl-v10.html at the root of this
distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must
not remove this notice, or any other, from this software.
