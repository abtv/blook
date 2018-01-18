# blook

[![Build Status](https://travis-ci.org/abtv/blook.svg?branch=master)](https://travis-ci.org/abtv/blook)

`blook` is a tool developed to save your time during text file analisys.
It makes very fast search in sorted files and returns as many results as you wish.

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

## License

Copyright (c) Andrey Butov. All rights reserved. The use and
distribution terms for this software are covered by the Eclipse
Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
which can be found in the file epl-v10.html at the root of this
distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must
not remove this notice, or any other, from this software.
