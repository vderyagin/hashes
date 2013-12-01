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
	"runtime"
	"sync"
)

const stdin = "*stdin*"

type Results map[string][]HashSum

type HashSum struct {
	Name  string
	Value string
	hash  hash.Hash
}

var format = flag.String("format", "plain", "format of output (allowed values: plain, json, xml)")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	results := make(Results)

	files := inputFiles()

	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, f := range files {
		results[f] = emptyHashes()

		writers := make([]io.Writer, len(results[f]))

		for idx := range results[f] {
			writers[idx] = results[f][idx].hash
		}

		input := inputFrom(f)

		go func(ws []io.Writer, in io.Reader, res []HashSum) {
			defer wg.Done()
			io.Copy(MultiGoroutineWriter(ws...), in)
			calculateSums(res)
		}(writers, input, results[f])
	}

	wg.Wait()

	outputResults(results)
}

type multiGoroutineWriter struct {
	writers []io.Writer
}

func MultiGoroutineWriter(writers ...io.Writer) multiGoroutineWriter {
	return multiGoroutineWriter{writers}
}

func (w multiGoroutineWriter) Write(p []byte) (n int, err error) {
	var wg sync.WaitGroup

	wg.Add(len(w.writers))

	for _, writer := range w.writers {
		go func(w io.Writer) {
			defer wg.Done()
			w.Write(p)
		}(writer)
	}

	wg.Wait()

	return len(p), nil
}

func inputFiles() []string {
	args := flag.Args()
	files := make([]string, 0, len(args))

	for _, f := range args {
		fi, err := os.Stat(f)

		if err != nil {
			fmt.Fprintf(os.Stderr, "could not read file '%s'\n", f)
			os.Exit(1)
		}

		if !fi.Mode().IsRegular() {
			fmt.Fprintf(os.Stderr, "'%s' is not a regular file\n", f)
			os.Exit(1)
		}

		files = append(files, f)
	}

	if len(files) == 0 {
		files = append(files, stdin) // read from stdin when no files specified
	}

	return files
}

func inputFrom(s string) io.Reader {
	if s == stdin {
		return os.Stdin
	} else {
		input, err := os.Open(s)

		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open file '%s'\n", flag.Arg(0))
			os.Exit(1)
		}

		return input
	}
}

func emptyHashes() []HashSum {
	return []HashSum{
		{Name: "md5", hash: md5.New()},
		{Name: "sha1", hash: sha1.New()},
		{Name: "sha256", hash: sha256.New()},
		{Name: "sha512", hash: sha512.New()},
		{Name: "adler32", hash: adler32.New()},
		{Name: "crc32 (IEEE)", hash: crc32.New(crc32.MakeTable(crc32.IEEE))},
		{Name: "crc32 (Castagnoli)", hash: crc32.New(crc32.MakeTable(crc32.Castagnoli))},
		{Name: "crc32 (Koopman)", hash: crc32.New(crc32.MakeTable(crc32.Koopman))},
		{Name: "crc64 (ISO)", hash: crc64.New(crc64.MakeTable(crc64.ISO))},
		{Name: "crc64 (ECMA)", hash: crc64.New(crc64.MakeTable(crc64.ECMA))},
		{Name: "fnv32-1", hash: fnv.New32()},
		{Name: "fnv32-1a", hash: fnv.New32a()},
		{Name: "fnv64-1", hash: fnv.New64()},
		{Name: "fnv64-1a", hash: fnv.New64a()},
	}
}

func calculateSums(hashes []HashSum) {
	for idx := range hashes {
		hashes[idx].Value = fmt.Sprintf("%x", hashes[idx].hash.Sum(nil))
	}
}

func outputResults(results Results) {
	switch *format {
	case "plain":
		for file, hashes := range results {
			fmt.Printf("%s:\n", file)
			for _, hsh := range hashes {
				fmt.Printf("  %s: %s\n", hsh.Name, hsh.Value)
			}
		}
	case "json":
		j, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(j))
	case "xml":
		x, _ := xml.MarshalIndent(results, "", "  ")
		fmt.Println(string(x))
	default:
		fmt.Fprintf(os.Stderr, "unknown output format: '%s'\n", *format)
		os.Exit(1)
	}
}
