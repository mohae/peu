Peu
=====
[![Build Status](https://travis-ci.org/mohae/peu.png)](https://travis-ci.org/mohae/peu)
un peu: a little

## About
Peu makes things smaller, hopefully.

Peu is a simple compression tool for compressing files individually. Peu does not create an archive, but compresses each file individually.

This is mainly to provide a simple framework for me to test various compression algorithms.

## Usage
Peu is written in Go.

Compile `peu` and either place the executable in your path or run it in the compile directory with `./peu`.

### Running `peu`
Compressed files do not retain path information; only the filename is preserved.

Peu will accept 1 or more filenames, each file is processed individually: they will not be put into an archive.

#### Compress
To compress:

    peu c [algorithm] <filename> ...

To compress and direct output to a specific directory:

    peu c -o=dest/path [algorithm] <filename>...

#### Decompress
To decompress:

    peu d [algorithm] <filename> ...

To decompress and direct output to a specific directory:
    peu d -o=dest/path [algorithm] <filename>...

## Supported compression algorithms

* lz4

## License
Copyright 2015 by Joel Scoble.

This is provided under the MIT License. For more details, please check the included LICENSE file.
