# backup-verify [![Build Status](https://cloud.drone.io/api/badges/d-fal/backup-verify/status.svg)](https://cloud.drone.io/d-fal/backup-verify) [![codecov](https://codecov.io/gh/d-fal/backup-verify/branch/main/graph/badge.svg?token=MDn8ffuCnQ)](https://codecov.io/gh/d-fal/backup-verify)

This application compares two dumps of any database rows without imposing unnecessary computational and memory load to the host machine.

### General idea:

Comparing two list of length `n` leads to an exhaustive `nxn` search. We should pick an item from the first list and search for the same one in the opposite one. We can even dismiss the items we find from the second list, but it comes with a painful memory cost: we should store everything into memory and it may result in a disaster.

### Memory leak issue

Comparing two gigantic json files, specially when they become large in size is not a trivial task. When we load the whole object in memory, it can easily lead to memory leak and os dump!. To alleviate that problem, we resort to streams of files, instead of dealing with the whole.

In this project, we have created a simplified streamer that starts reading a file and sends json blocks one by one.

> Note: we are aware that tools like [Automi](https://github.com/vladimirvivien/automi) would fulfil our goal, however we implemented by ourselves because we found it super pedantic.

### Preparing streamer :

```go

    streamer, err := streamming.NewStreamer(path)
	if err != nil {
		// take an action
    }

    // streaming blocks
    decoder, err := streamer.Stream()

	for { // forever loop that would be broken inside

		block, ok := decoder.Next()
		if !ok {
			break
        }

        fmt.Println(block) // data block derived from json file
	}

```

### Comparing two objects

`reflection` is one of the life savers when we want to compare two structs. However, using `reflection` is costly specially in heavy duty operations such as `n^2` search. To lessen the reflection pain, we have adopted [`go-cmp`](https://github.com/google/go-cmp) instead.

### CLI

build the app via

```bash
    make build
```

then run the app in the following way:

```bash
   $ ./bverify --source [YOUR SOURCE FILE] --destination [THE PATH TO DUMP FILE]

```

In the case you want to store diff data to a file, add the `--diff-path` flag and provide it with the destination of the intended diff file.

## Example run:

```bash
   $ ./bverify --source ./sample/original.json --destination ./sample/duplicate.json --diff-path ./out.txt

```

## Room for improvement

Currently, we are using a single process to run the app. In order to improve the performance of the app, I would suggest add concurrency to the app. However, using resources should be taken with full alacrity of devouring resources.
