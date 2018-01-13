# blook

`blook` is a tool developed to save your time during text file analisys.
It makes very fast search in sorted files and returns as many results as you wish.
Please see `Use cases` section.

## Problem

You have a big (several gigabytes or more) sorted text file.
You want to find all the strings which starts with some prefix and to the end of the file.
The file is too big to use _look_ Unix utility and _look_ returns only strings with the given prefix, but you want all the strings since the given position.
_sed_ is too slow.

## The solution

As your text files are sorted then you are happy!
`blook` utility makes _binary_ search in your files. You can save a lot of time during file analisys.

## Usage

`blook --pattern from --file file.log`

## Use cases

1. Search in logs

Suppose you have log files with the following content:

```
...
2018-01-13T10:10:23+00:00 fetch user with id = 100500
2018-01-13T10:10:23+00:00 check if user with id = 100500 is eligible to some cool stuff
2018-01-13T10:10:24+00:00 send email to user with id = 100500
...
```

The following command returns all the lines from `/var/log/remote/your_service.log` file from `2018-01-13T10:10:23+00:00` timestamp to the end of the file:

`./blook --pattern 2018-01-13T10:10:23+00:00 --file /var/log/remote/your_service.log`

2. Search in sorted files of any kind

Suppose you have text files with the following content:

```
...
AB100500800 ok
AB100500801 fail
AB100500807 fail
...
```

The following command returns all fail cases `/var/log/remote/your_device.log` file from `AB100500800` stamp to the end of the file:

`./blook --pattern AB100500800 --file /var/log/remote/your_device.log | grep fail`


## License

Copyright (c) Andrey Butov. All rights reserved. The use and
distribution terms for this software are covered by the Eclipse
Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
which can be found in the file epl-v10.html at the root of this
distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must
not remove this notice, or any other, from this software.
