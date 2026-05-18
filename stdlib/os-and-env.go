package stdlib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ---- Environment variables ----

func envDemo() {
	// get a single env var (returns "" if not set)
	home := os.Getenv("HOME")
	fmt.Println("HOME:", home)

	// get with exists check
	val, ok := os.LookupEnv("MY_VAR")
	if !ok {
		fmt.Println("MY_VAR not set")
	} else {
		fmt.Println("MY_VAR:", val)
	}

	// set for this process
	os.Setenv("MY_VAR", "hello")

	// unset
	os.Unsetenv("MY_VAR")

	// all env vars as []string of "KEY=VALUE"
	for _, e := range os.Environ() {
		_ = e
	}
}

// ---- Process info ----

func processInfo() {
	fmt.Println("PID:", os.Getpid())
	fmt.Println("Parent PID:", os.Getppid())
	fmt.Println("Args:", os.Args) // os.Args[0] = binary name, [1:] = flags/args

	// hostname
	hostname, _ := os.Hostname()
	fmt.Println("Host:", hostname)

	// user info
	uid := os.Getuid()
	gid := os.Getgid()
	fmt.Println("UID:", uid, "GID:", gid)
}

// ---- Working directory ----

func dirDemo() {
	cwd, _ := os.Getwd()
	fmt.Println("CWD:", cwd)

	os.Chdir("/tmp") // change working directory

	// Stat: get file info
	info, err := os.Stat("somefile.txt")
	if os.IsNotExist(err) {
		fmt.Println("file not found")
	} else if err == nil {
		fmt.Println("size:", info.Size(), "modified:", info.ModTime())
		fmt.Println("is dir:", info.IsDir())
		fmt.Println("mode:", info.Mode())
	}
}

// ---- Directory operations ----

func dirOps() {
	// create directory
	os.Mkdir("mydir", 0755)
	os.MkdirAll("a/b/c", 0755) // creates intermediate dirs

	// list directory
	entries, err := os.ReadDir(".")
	if err == nil {
		for _, e := range entries {
			fmt.Printf("%s\t%v\n", e.Name(), e.IsDir())
		}
	}

	// temp directory
	tmpDir, _ := os.MkdirTemp("", "myapp-*")
	defer os.RemoveAll(tmpDir) // clean up

	// temp file
	tmpFile, _ := os.CreateTemp(tmpDir, "data-*.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// remove
	os.Remove("file.txt")       // single file or empty dir
	os.RemoveAll("dir")         // recursive
}

// ---- Path manipulation (filepath package) ----

func filepathDemo() {
	p := "/home/user/docs/file.txt"

	fmt.Println(filepath.Dir(p))      // /home/user/docs
	fmt.Println(filepath.Base(p))     // file.txt
	fmt.Println(filepath.Ext(p))      // .txt

	// join paths (OS-aware separator)
	joined := filepath.Join("/home", "user", "docs", "file.txt")
	fmt.Println(joined) // /home/user/docs/file.txt

	// absolute path
	abs, _ := filepath.Abs("../relative/path")
	fmt.Println(abs)

	// clean: resolve . and .. and double slashes
	fmt.Println(filepath.Clean("/a//b/../c")) // /a/c

	// Walk: traverse directory tree
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.IsDir())
		return nil
	})

	// WalkDir (Go 1.16+, more efficient than Walk)
	filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, d.IsDir())
		return nil
	})

	// Glob: match files by pattern
	matches, _ := filepath.Glob("*.go")
	fmt.Println(matches)
}

// ---- Running external commands ----

func execDemo() {
	// simple command
	out, err := exec.Command("echo", "hello").Output()
	if err == nil {
		fmt.Println(string(out))
	}

	// with environment
	cmd := exec.Command("printenv", "MY_VAR")
	cmd.Env = append(os.Environ(), "MY_VAR=injected")
	out, _ = cmd.Output()
	fmt.Println(string(out))

	// run and wait (discards output)
	exec.Command("touch", "/tmp/test.txt").Run()

	// pipe stdin/stdout
	cmd2 := exec.Command("cat")
	cmd2.Stdin = os.Stdin
	cmd2.Stdout = os.Stdout
	cmd2.Run()
}
