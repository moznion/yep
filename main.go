package main

import "flag"
import "fmt"
import "math"
import "time"

func main() {
	flag.Usage = func() {
		fmt.Printf(`Usage:
  yep [options] [expletive (optional)]

Options:
`)
		flag.PrintDefaults()
	}

	interval := flag.Float64("interval", 0, "Interval as seconds to output expletives. If this value is zero or negative, it doesn't have any breaks")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()
	args := flag.Args()

	if *help {
		flag.Usage()
		return
	}

	expletive := "y"
	if len(args) > 0 {
		expletive = args[0]
	}

	i := 0.0
	decimal := 0.0
	if *interval > 0.0 {
		i = math.Trunc(*interval)
		decimal = *interval - i
	}

	duration := time.Duration(i)*time.Second + time.Duration(decimal*1000)*time.Millisecond
	for {
		fmt.Printf("%s\n", expletive)
		time.Sleep(duration)
	}
}
