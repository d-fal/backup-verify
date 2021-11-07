package application

import (
	"fmt"
	"os"

	streamming "github.com/d-fal/bverify/internal/stream"
	"github.com/d-fal/bverify/utils"
	"github.com/logrusorgru/aurora"
)

type Opt func(*utils.Row) error

func Start(original, duplicate string, opts ...Opt) (bool, error) {
	var filesMatch bool = true

	origStreamer, err := streamming.NewStreamer(original)
	if err != nil {
		return false, err
	}

	duplicateStreamer, err := streamming.NewStreamer(duplicate)
	if err != nil {
		return false, err
	}

	decoder, err := origStreamer.Stream()

	if err != nil {
		return false, err
	}

	for {

		block, ok := decoder.Next()
		if !ok {
			break
		}

		w := utils.SetRow(block)

		if exists, _ := duplicateStreamer.Get().Match(block); !exists {
			filesMatch = false
			for _, opt := range opts {
				if err := opt(w.Get()); err != nil {
					fmt.Println("error: ", err)
				}
			}
			continue
		}

	}

	return filesMatch, nil
}

func WithConsolePrint() Opt {
	return func(u *utils.Row) error {
		fmt.Println("mismatch : ", aurora.Yellow(u))
		return nil
	}
}

func WithDiffSave(diffPath string) Opt {
	file, err := os.OpenFile(diffPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return func(r *utils.Row) error {
			return err
		}
	}

	return func(r *utils.Row) error {
		_, err := file.WriteString(r.String() + "\n")
		return err
	}
}
