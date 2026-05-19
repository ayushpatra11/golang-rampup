package stdlib

import (
	"flag"
	"fmt"
	"os"
)

// flag package: command-line flag parsing (built-in)

/*
---- Defining flags ----

Flags are defined at package level (before main) or inside main/init.

flag.String(name, default, usage)    -> *string
flag.Int(name, default, usage)       -> *int
flag.Bool(name, default, usage)      -> *bool
flag.Float64(name, default, usage)   -> *float64
flag.Duration(name, default, usage)  -> *time.Duration

Var variants bind to an existing variable:
flag.StringVar(&myVar, name, default, usage)
*/

var (
	host    = flag.String("host", "localhost", "server host")
	port    = flag.Int("port", 8080, "server port")
	verbose = flag.Bool("verbose", false, "enable verbose output")
	timeout = flag.Duration("timeout", 0, "operation timeout (0 = no timeout)")
)

func runCLI() {
	// must call Parse before reading flag values
	flag.Parse()

	fmt.Printf("host=%s port=%d verbose=%v timeout=%v\n",
		*host, *port, *verbose, *timeout)

	// non-flag args after flags
	args := flag.Args()       // []string of remaining args
	nArgs := flag.NArg()      // count of remaining args
	nFlag := flag.NFlag()     // count of flags that were set
	fmt.Println(args, nArgs, nFlag)
}

// Custom usage message
func customUsage() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <input-file>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}
}

// ---- FlagSet: multiple subcommands ----

func subcommandDemo() {
	// each subcommand gets its own FlagSet
	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	serveHost := serveCmd.String("host", "0.0.0.0", "bind host")
	servePort := serveCmd.Int("port", 8080, "bind port")

	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	migrateDir := migrateCmd.String("dir", "./migrations", "migrations directory")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "expected subcommand: serve|migrate")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "serve":
		serveCmd.Parse(os.Args[2:])
		fmt.Printf("serving on %s:%d\n", *serveHost, *servePort)
	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		fmt.Printf("migrating from %s\n", *migrateDir)
	default:
		fmt.Fprintln(os.Stderr, "unknown subcommand:", os.Args[1])
		os.Exit(1)
	}
}

// ---- Custom flag type ----
// Implement flag.Value interface: Set(string) error, String() string

type StringSlice []string

func (s *StringSlice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

func (s *StringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

func customFlagDemo() {
	var tags StringSlice
	flag.Var(&tags, "tag", "tags (can be specified multiple times)")
	flag.Parse()
	// go run main.go -tag foo -tag bar -tag baz
	// tags = [foo bar baz]
}

/*
flag package conventions:
- Flags use - prefix: -verbose (single dash, unlike POSIX -v)
- Both -flag and --flag work
- -flag=value and -flag value both work (except bool flags with =)
- bool flags: -verbose=false must use =; -verbose alone sets to true

For complex CLIs with subcommands, nested flags, and shell completion,
consider third-party packages:
- github.com/spf13/cobra  (most popular - used by kubectl, hugo, etc.)
- github.com/urfave/cli   (simpler alternative)
*/
