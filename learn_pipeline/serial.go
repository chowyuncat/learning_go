package main

import (
    "crypto/md5"
    "path/filepath"
    "io/ioutil"
    "os"
    "fmt"
)

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5All_Serial(root string) (map[string][md5.Size]byte, error) {
    m := make(map[string][md5.Size]byte)
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        data, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        m[path] = md5.Sum(data)
        return nil
    })
    if err != nil {
        return nil, err
    }
    return m, nil
}


func DoSerial(root string) {
    md5sums, err := MD5All_Serial(root)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    for pathname, hash := range md5sums {
        // match output format of the utility md5sum (md5 -r)
        fmt.Printf("%x %s\n", hash, pathname)
    }
}

func main() {
    root := os.Args[1]
    DoSerial(root)
}
