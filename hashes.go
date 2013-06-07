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
	"os"
)

type h struct {
	name string
	hash hash.Hash
}

func main() {
	var input io.Reader
	var err error

	switch {
	case len(os.Args) == 2:
		input, err = os.Open(os.Args[1])
	case len(os.Args) > 2:
		fmt.Fprintln(os.Stderr, "only one argument is allowed")
		os.Exit(1)
	default:
		input = os.Stdin
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file '%s'\n", os.Args[1])
		os.Exit(1)
	}

	hashes := []h{
		h{name: "md5", hash: md5.New()},
		h{name: "sha1", hash: sha1.New()},
		h{name: "sha256", hash: sha256.New()},
		h{name: "sha512", hash: sha512.New()},
		h{name: "adler32", hash: adler32.New()},
		h{name: "crc32 (IEEE)", hash: crc32.New(crc32.MakeTable(crc32.IEEE))},
		h{name: "crc32 (Castagnoli)", hash: crc32.New(crc32.MakeTable(crc32.Castagnoli))},
		h{name: "crc32 (Koopman)", hash: crc32.New(crc32.MakeTable(crc32.Koopman))},
		h{name: "crc64 (ISO)", hash: crc64.New(crc64.MakeTable(crc64.ISO))},
		h{name: "crc64 (ECMA)", hash: crc64.New(crc64.MakeTable(crc64.ECMA))},
		h{name: "fnv32-1", hash: fnv.New32()},
		h{name: "fnv32-1a", hash: fnv.New32a()},
		h{name: "fnv64-1", hash: fnv.New64()},
		h{name: "fnv64-1a", hash: fnv.New64a()},
	}

	writers := make([]io.Writer, len(hashes))

	for idx := range hashes {
		writers[idx] = hashes[idx].hash
	}

	io.Copy(io.MultiWriter(writers...), input)

	for _, hsh := range hashes {
		fmt.Printf("%s: %x\n", hsh.name, hsh.hash.Sum(nil))
	}
}
