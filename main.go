package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	tmpFp := flag.String("f", "notfound", "filepath")
	flag.Parse()
	filePath := *tmpFp //pointer to string

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hashSha1 := sha1.New()
	if _, err := io.Copy(hashSha1, file); err != nil {
		panic(err)
	}
	strHashSha1 := hex.EncodeToString(hashSha1.Sum(nil))

	hashSha256 := sha256.New() //ハッシュ値がなぜが違う
	if _, err := io.Copy(hashSha256, file); err != nil {
		panic(err)
	}
	strHashSha256 := hex.EncodeToString(hashSha256.Sum(nil))

	hashSha384 := sha512.New384()
	if _, err := io.Copy(hashSha384, file); err != nil {
		panic(err)
	}
	strHashSha384 := hex.EncodeToString(hashSha384.Sum(nil))
	//hashSha256 := sha256.New()
	//hashSha384 := sha512.New384()
	//hashSha512 := sha512.New()
	//
	//hashRmd160 := crypto.RIPEMD160.New()
	//io.WriteString(hashRmd160, string())
	//if _, err = io.Copy(hashMd5, file); err != nil{
	//	panic(err)
	fmt.Println(strHashSha1, strHashSha256, strHashSha384)
}
