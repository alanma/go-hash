package main

import (
        "fmt"
        "strings"
        "io"
        "os"
        goopt     "github.com/droundy/goopt"
        "hash"
        "crypto/md5"
        "crypto/sha1"
        "crypto/sha256"
        "crypto/sha512"
        ripemd160 "golang.org/x/crypto/ripemd160"
)

var algorithms = [9]string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512", "sha512224", 
                            "sha512256", "ripemd160"}

var Version string = "0.1.0"

func main() {
        var progname = os.Args[0][strings.LastIndex(os.Args[0], "/")+1:]

        var algorithm = goopt.StringWithLabel([]string{"-a", "--algorithm"}, algorithms[0],
                "ALG", fmt.Sprintf("Hashing algorithm: %s", algorithms))
        var version = goopt.Flag([]string{"-V", "--version"}, []string{}, "Display version", "")


        goopt.Summary = fmt.Sprintf("%s [OPTIONS] FILENAME\n\nMessage digest calculator with various hashing algorithms.\n\nArguments:\n  FILENAME                 File(s) to hash\n", progname)
        goopt.Parse(nil)

        var files []string = goopt.Args

        if *version {
                fmt.Printf("%s version %s\n", progname, Version)
                os.Exit(0)
        }

        validateAlgorithm(goopt.Usage(), *algorithm)
        valildateFiles(goopt.Usage(), files)

        calculateDigest(files, *algorithm)

        os.Exit(0)
}

func calculateDigest(files []string, algorithm string) {
        var hasher hash.Hash
        switch algorithm {
        case "md5":
                hasher = md5.New()
        case "sha1":
                hasher = sha1.New()
        case "sha224":
                hasher = sha256.New224()
        case "sha256":
                hasher = sha256.New()
        case "sha384":
                hasher = sha512.New384()
        case "sha512":
                hasher = sha512.New()
        case "sha512224":
                hasher = sha512.New512_224()
        case "sha512256":
                hasher = sha512.New512_256()
        case "ripemd160":
                hasher = ripemd160.New()
        default:
                fmt.Printf("Algorithm %s not implemented yet\n", algorithm)
                os.Exit(1)
        }
        for _, fn := range files {
                hasher.Reset()
                file, err := os.Open(fn)
                if err != nil {
                        panic(err.Error())
                }
                if _, err := io.Copy(hasher, file); err != nil {
                        panic(err.Error())
                }       
                fmt.Printf("%x  %s\n", hasher.Sum(nil), fn)
                file.Close()
        }
}

func validateAlgorithm(usage, algorithm string) {
        if ! contains(algorithms[:], algorithm) {
                fmt.Printf("\n***** Unexpected algorithm: %s *****\n\n", algorithm)
                fmt.Println(usage)
                os.Exit(1)
        }
}

func valildateFiles(usage string, files []string) {
        if len(files) == 0 {
                fmt.Printf("\n***** Must pass at least 1 filename. *****\n\n")
                fmt.Println(usage)
                os.Exit(1)
        }
}

// Checks if an array/slice contains a given string
func contains(s []string, str string) bool {
        for _, a := range s {
                if a == str {
                        return true
                }
        }
        return false
}
