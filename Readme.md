# Hashes #

Calculate different hashes for file at the same time. Rather
efficient, as input is read only once.

## Installation ##

**You'll need Go workspace set up, read [here](http://golang.org/doc/code.html) on how to do that.**


```shell
$ go get github.com/vderyagin/hashes
```

To update installation just add `-u` flag.

## Usage ##

```shell
$ hashes somefile
md5: 5bf68257840b05436466d4e815bd7bac
sha1: 9a96f5808e5289aef41015df6bf9c19b62adfa93
sha256: 54a1d9399132bc5910c2e72eefaa65b1de5bb358aad8ce90cd4fa870ba15d924
sha512: f201ffc5952a4eeed9adae369ed1b12f644fcd2810678604cb0703c37e64d2e655f44cc64cbcf47f70d90208bee1275eea77d0f559fec393ce403810e896ece4
adler32: fc22ecce
crc32 (IEEE): 6568dfd6
crc32 (Castagnoli): 33f3f518
crc32 (Koopman): e6b36317
crc64 (ISO): 3d8397b0979c506c
crc64 (ECMA): 12f17c708b67b6bc
fnv32-1: 23124a1a
fnv32-1a: ae662f74
fnv64-1: af92ca3d4f7dbfda
fnv64-1a: 0615cf2092b65c54
```

## Currently supported hash functions ##

* md5
* sha1
* sha256
* sha512
* adler32
* crc32 (IEEE)
* crc32 (Castagnoli)
* crc32 (Koopman)
* crc64 (ISO)
* crc64 (ECMA)
* fnv32-1
* fnv32-1a
* fnv64-1
* fnv64-1a
