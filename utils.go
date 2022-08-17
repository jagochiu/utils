package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
EpochToHumanReadable -
*/
func Date(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

/*
ShortDate -
*/
func ShortDate(t time.Time) string { return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day()) }

/*
FileDate -
*/
func FileDate(t time.Time) string {
	return fmt.Sprintf("%04d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func CreateDir(path string) error {
	if exist, _ := PathExists(path); exist {
		return nil
	}
	err := os.Mkdir(path, 0770)
	if err != nil {
		log.Printf(`[DIR] create directory fail - %v`, err)
	}
	return err
}

func Bool[V any](condition bool, a, b V) V {
	if !condition {
		return b
	}
	return a
}
