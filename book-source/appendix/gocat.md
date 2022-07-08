# Cat in Go

```go
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

var nFlag = flag.Bool("n", false, "Number the output lines, starting at 1.")

func main() {
	flag.Parse()

	// stop if the user didn't provide any arguments
	if len(flag.Args()) == 0 {
		return
	}

	encounterederrors := []error{}
	// read and print each file the user provided.
	for _, f := range flag.Args() {
		fileData, err := os.ReadFile(f)
		if err != nil {
			// If we hit an error with a specific file, just skip it and move on
			// and report the error later.
			encounterederrors = append(encounterederrors, err)
			break
		}

		if *nFlag {
			// split at newlines so that we can number each line.
			fileDataSplit := bytes.Split(fileData, []byte("\n"))
			// there's always an extra newline at the end when we split,
			// so remove that.
			fileDataSplit = fileDataSplit[0 : len(fileDataSplit)-1]

			i := 1
			for _, line := range fileDataSplit {
				fmt.Fprintf(os.Stdout, "\t%d\t%s\n", i, string(line))
				i++
			}

			continue
		}
		
		fmt.Fprintln(os.Stdout, string(fileData))
	}

	if len(encounterederrors) > 0 {
		for _, e := range encounterederrors {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}

	os.Exit(0)
}

```

## Output

```
$ ./gocat -h
Usage of ./gocat:
  -n    Number the output lines, starting at 1.
```

```
$ ./gocat main.go 
package main

import (
        "bytes"
        "flag"
        "fmt"
        "os"
)

var nFlag = flag.Bool("n", false, "Number the output lines, starting at 1.")

func main() {
        flag.Parse()

        // stop if the user didn't provide any arguments
        if len(flag.Args()) == 0 {
                return
        }

        encounterederrors := []error{}
        // read and print each file the user provided.
        for _, f := range flag.Args() {
                fileData, err := os.ReadFile(f)
                if err != nil {
                        // If we hit an error with a specific file, just skip it and move on
                        // and report the error later.
                        encounterederrors = append(encounterederrors, err)
                        break
                }

                if *nFlag {
                        // split at newlines so that we can number each line.
                        fileDataSplit := bytes.Split(fileData, []byte("\n"))
                        // there's always an extra newline at the end when we split,
                        // so remove that.
                        fileDataSplit = fileDataSplit[0 : len(fileDataSplit)-1]

                        i := 1
                        for _, line := range fileDataSplit {
                                fmt.Fprintf(os.Stdout, "\t%d\t%s\n", i, string(line))
                                i++
                        }
                } else {
                        fmt.Fprintln(os.Stdout, string(fileData))
                }
        }

        if len(encounterederrors) > 0 {
                for _, e := range encounterederrors {
                        fmt.Fprintln(os.Stderr, e)
                }
                os.Exit(1)
        }

        os.Exit(0)
}

```

```
# ./gocat -n main.go 
        1       package main
        2
        3       import (
        4               "bytes"
        5               "flag"
        6               "fmt"
        7               "os"
        8       )
        9
        10      var nFlag = flag.Bool("n", false, "Number the output lines, starting at 1.")
        11
        12      func main() {
        13              flag.Parse()
        14
        15              // stop if the user didn't provide any arguments
        16              if len(flag.Args()) == 0 {
        17                      return
        18              }
        19
        20              encounterederrors := []error{}
        21              // read and print each file the user provided.
        22              for _, f := range flag.Args() {
        23                      fileData, err := os.ReadFile(f)
        24                      if err != nil {
        25                              // If we hit an error with a specific file, just skip it and move on
        26                              // and report the error later.
        27                              encounterederrors = append(encounterederrors, err)
        28                              break
        29                      }
        30
        31                      if *nFlag {
        32                              // split at newlines so that we can number each line.
        33                              fileDataSplit := bytes.Split(fileData, []byte("\n"))
        34                              // there's always an extra newline at the end when we split,
        35                              // so remove that.
        36                              fileDataSplit = fileDataSplit[0 : len(fileDataSplit)-1]
        37
        38                              i := 1
        39                              for _, line := range fileDataSplit {
        40                                      fmt.Fprintf(os.Stdout, "\t%d\t%s\n", i, string(line))
        41                                      i++
        42                              }
        43                      } else {
        44                              fmt.Fprintln(os.Stdout, string(fileData))
        45                      }
        46              }
        47
        48              if len(encounterederrors) > 0 {
        49                      for _, e := range encounterederrors {
        50                              fmt.Fprintln(os.Stderr, e)
        51                      }
        52                      os.Exit(1)
        53              }
        54
        55              os.Exit(0)
        56      }
```