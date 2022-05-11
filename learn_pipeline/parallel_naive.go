package main

import (
    "crypto/md5"
    "os"
    "path/filepath"
    "io/ioutil"
    "fmt"
    "sync"
)

type _PathResult struct {
    path string
    err error
}

func walkPath(root string) chan _PathResult {
    out := make(chan _PathResult)

    visitor := func(path string, info os.FileInfo, err error) error {
        if err != nil {
            out <- _PathResult{path, err}
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        
        out <- _PathResult{path, nil}
        return nil

    }

    go func() {
        err := filepath.Walk(root, visitor)
        if err != nil {
    //      @TODO: how to handle final error returned from Walk?
        }
        close(out)
    }()


    return out
}

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5_Parallel(root string) {

    var wg sync.WaitGroup

    for r := range walkPath(root) {
        if r.err != nil {
            continue
        }

        wg.Add(1)

        go func(r _PathResult) {
            defer wg.Done()
            data, err := ioutil.ReadFile(r.path)
            if err != nil {
                return
            }
            fmt.Printf("%x %s\n", md5.Sum(data), r.path)
        }(r)
    }

    wg.Wait()
}

func main() {
    root := os.Args[1]
    MD5_Parallel(root)
}
