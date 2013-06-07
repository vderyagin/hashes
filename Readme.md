# Hashes #

Calculate values of different hash functions for some piece of data
at the same time. Rather efficient, as input is read only once.

## Installation ##

**You'll need Go workspace set up, read
  [here](http://golang.org/doc/code.html) on how to do that.**


```shell
$ go get github.com/vderyagin/hashes
```

To update installation just add `-u` flag.

## Usage ##

`hashes` accepts single filename as an argument. If no arguments
provided, data is read from stdin.

### Output formats ###

Three output formats are supported: plain, json and xml:

```shell
Usage of hashes:
  -format="plain": format of output (allowed values: plain, json, xml)
```

### Usage examples ###

Plain output format:

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

JSON output:

```shell
$ hashes --format=json someotherfile
[
  {
    "Name": "md5",
    "Value": "f2b31d3a6a5c2827da2bb8832a082fa2"
  },
  {
    "Name": "sha1",
    "Value": "f1378a19c1ff20cfa5fef8103835084b9bdb7b2d"
  },
  {
    "Name": "sha256",
    "Value": "1c1c2a5993c6bd88588fa8f936b3e90eb64249de0088273e0479b494afaa36a0"
  },
  {
    "Name": "sha512",
    "Value": "3be02bd92fc20d8551a1a3e3bd71c9cbda4fddead75770b0aa867a64c80784c0958d03cc6eface14855dd36835108fdb79fae5e5aa2e2ea0211739329431edb5"
  },
  {
    "Name": "adler32",
    "Value": "933507dd"
  },
  {
    "Name": "crc32 (IEEE)",
    "Value": "7c992293"
  },
  {
    "Name": "crc32 (Castagnoli)",
    "Value": "28f9621e"
  },
  {
    "Name": "crc32 (Koopman)",
    "Value": "f98f7d0a"
  },
  {
    "Name": "crc64 (ISO)",
    "Value": "a559d223584f0660"
  },
  {
    "Name": "crc64 (ECMA)",
    "Value": "eb0b411de9b9a2bf"
  },
  {
    "Name": "fnv32-1",
    "Value": "1ff0b102"
  },
  {
    "Name": "fnv32-1a",
    "Value": "9e82a844"
  },
  {
    "Name": "fnv64-1",
    "Value": "f44436c1175a84a2"
  },
  {
    "Name": "fnv64-1a",
    "Value": "584d896cf9a92744"
  }
]
```

## Supported hash functions ##

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
