package testutils_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/the-coding-dead/code/testutils"
)

func TestDeferRollbackDir(t *testing.T) {
	t.Parallel()

	t.Run("Text", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			src string
			raw string
		}{
			{"testdata/defer_rollback_dir/text/a.txt", "a\n"},
			{"testdata/defer_rollback_dir/text/b.txt", "b\n"},
			{"testdata/defer_rollback_dir/text/c/c.txt", "c\n"},
		}

		t.Run("Setup", func(t *testing.T) {
			testutils.DeferRollbackDir(t, "testdata/defer_rollback_dir/text")

			for _, tt := range tests {
				tt := tt
				t.Run(tt.src, func(t *testing.T) {
					t.Parallel()

					if err := ioutil.WriteFile(tt.src, []byte("hoge"), os.ModePerm); err != nil {
						t.Fatal(err)
					}
				})
			}
		})

		t.Run("CheckRollbacked", func(t *testing.T) {
			for _, tt := range tests {
				tt := tt
				t.Run(tt.src, func(t *testing.T) {
					t.Parallel()

					bs, err := ioutil.ReadFile(tt.src)
					if err != nil {
						t.Fatal(err)
					}

					if want, got := tt.raw, string(bs); want != got {
						t.Errorf("want: '%s', but got '%s'", want, got)
					}
				})
			}
		})
	})

	t.Run("Image", func(t *testing.T) {
		t.Parallel()

		var want []byte

		const (
			dir      = "testdata/defer_rollback_dir/image"
			fileName = "testdata/defer_rollback_dir/image/a.jpg"
		)

		t.Run("Setup", func(t *testing.T) {
			testutils.DeferRollbackDir(t, dir)
			var err error

			want, err = ioutil.ReadFile(fileName)
			if err != nil {
				t.Fatal(err)
			}

			if err := ioutil.WriteFile(fileName, []byte("hoge"), os.ModePerm); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("CheckRollbacked", func(t *testing.T) {
			got, err := ioutil.ReadFile(fileName)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	})
}
