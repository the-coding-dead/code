package testutils

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func DeferRollbackDir(t *testing.T, srcDir string) {
	t.Helper()

	memo := setup(t, srcDir)

	t.Cleanup(func() { teardown(t, memo) })
}

type memo struct {
	dir  []file
	file map[file]string
}

type file struct {
	path string
	mode os.FileMode
}

func setup(t *testing.T, srcDir string) *memo {
	t.Helper()

	memo := &memo{
		dir:  make([]file, 0),
		file: make(map[file]string),
	}

	var fileNameCnt int

	dir := t.TempDir()

	err := filepath.Walk(srcDir, func(src string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			memo.dir = append(memo.dir, file{src, info.Mode()})

			return nil
		}

		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		dst := filepath.Join(dir, strconv.Itoa(fileNameCnt))

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = io.Copy(out, in); err != nil {
			return err
		}

		memo.file[file{src, info.Mode()}] = dst

		fileNameCnt++

		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	return memo
}

func teardown(t *testing.T, memo *memo) {
	for _, dir := range memo.dir {
		if _, err := os.Stat(dir.path); os.IsNotExist(err) {
			if err := os.Mkdir(dir.path, dir.mode); err != nil {
				t.Fatal(err)
			}
		}
	}

	for src, dst := range memo.file {
		in, err := os.OpenFile(src.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, src.mode)
		if err != nil {
			t.Fatal(err)
		}
		defer in.Close()

		out, err := os.Open(dst)
		if err != nil {
			t.Fatal(err)
		}
		defer out.Close()

		if _, err = io.Copy(in, out); err != nil {
			t.Fatal(err)
		}
	}
}
