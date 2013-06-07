package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"os"
)

type HashSum struct {
	Name  string
	Value string
	hash  hash.Hash
}

var format = flag.String("format", "plain", "format of output (allowed values: plain, json, xml)")

func main() {
	flag.Parse()

	input := determineInput()
	hashes := emptyHashes()

	writers := make([]io.Writer, len(hashes))

	for idx := range hashes {
		writers[idx] = hashes[idx].hash
	}

	io.Copy(io.MultiWriter(writers...), input)

	calculateSums(hashes)
	outputHashes(hashes)
}

func determineInput() io.Reader {
	var input io.Reader
	var err error

	switch {
	case len(flag.Args()) == 1:
		input, err = os.Open(flag.Arg(0))
	case len(flag.Args()) > 1:
		fmt.Fprintln(os.Stderr, "only one argument is allowed")
		os.Exit(1)
	default:
		input = os.Stdin
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open file '%s'\n", flag.Arg(0))
		os.Exit(1)
	}

	return input
}

func emptyHashes() []HashSum {
	return []HashSum{
		HashSum{Name: "md5", hash: md5.New()},
		HashSum{Name: "sha1", hash: sha1.New()},
		HashSum{Name: "sha256", hash: sha256.New()},
		HashSum{Name: "sha512", hash: sha512.New()},
		HashSum{Name: "adler32", hash: adler32.New()},
		HashSum{Name: "crc32 (IEEE)", hash: crc32.New(crc32.MakeTable(crc32.IEEE))},
		HashSum{Name: "crc32 (Castagnoli)", hash: crc32.New(crc32.MakeTable(crc32.Castagnoli))},
		HashSum{Name: "crc32 (Koopman)", hash: crc32.New(crc32.MakeTable(crc32.Koopman))},
		HashSum{Name: "crc64 (ISO)", hash: crc64.New(crc64.MakeTable(crc64.ISO))},
		HashSum{Name: "crc64 (ECMA)", hash: crc64.New(crc64.MakeTable(crc64.ECMA))},
		HashSum{Name: "fnv32-1", hash: fnv.New32()},
		HashSum{Name: "fnv32-1a", hash: fnv.New32a()},
		HashSum{Name: "fnv64-1", hash: fnv.New64()},
		HashSum{Name: "fnv64-1a", hash: fnv.New64a()},
	}
}

func calculateSums(hashes []HashSum) {
	for idx := range hashes {
		hashes[idx].Value = fmt.Sprintf("%x", hashes[idx].hash.Sum(nil))
	}
}

func outputHashes(hashes []HashSum) {
	switch *format {
	case "plain":
		for _, hsh := range hashes {
			fmt.Printf("%s: %s\n", hsh.Name, hsh.Value)
		}
	case "json":
		j, _ := json.MarshalIndent(hashes, "", "  ")
		fmt.Println(string(j))
	case "xml":
		x, _ := xml.MarshalIndent(hashes, "", "  ")
		fmt.Println(string(x))
	default:
		fmt.Fprintf(os.Stderr, "unknown output format: '%s'\n", *format)
		os.Exit(1)
	}
}
