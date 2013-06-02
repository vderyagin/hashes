package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"log"
	"os"
)

var format = `md5:          %x
sha1:         %x
sha256:       %x
sha512:       %x
adler32:      %x
crc32 (IEEE): %x
crc64 (ECMA): %x
crc64 (ISO):  %x
fnv32:        %x
fnv32a:       %x
fnv64:        %x
fnv64a:       %x

`

func main() {
	files := os.Args[1:]

	for _, file := range files {
		f, err := os.Open(file)

		if err != nil {
			log.Fatalf("could not open file %s\n", file)
		}

		hashes := []io.Writer{
			md5.New(),
			sha1.New(),
			sha256.New(),
			sha512.New(),
			adler32.New(),
			crc32.NewIEEE(),
			crc64.New(crc64.MakeTable(crc64.ISO)),
			crc64.New(crc64.MakeTable(crc64.ECMA)),
			fnv.New32(),
			fnv.New32a(),
			fnv.New64(),
			fnv.New64a(),
		}

		io.Copy(io.MultiWriter(hashes...), f)

		var sums = make([]interface{}, len(hashes))

		for idx, hsh := range hashes {
			sums[idx] = hsh.(hash.Hash).Sum(nil)
		}

		fmt.Printf("%s:\n", file)
		fmt.Printf(format, sums...)
	}
}
