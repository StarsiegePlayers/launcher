package main

//go:generate go-winres make

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	"github.com/lxn/win"
)

var (
	VERSION string
	DATE    string
	TIME    string
	DEBUG   string
)

// main is the body of the program
// technical note: by and large we fail silently, unless we're debugging
func main() {
	// if we weren't called with any args, provide syntax
	if len(os.Args) <= 1 {
		syntax()
		return
	}

	// normalize our commandline
	argc := strings.ToLower(os.Args[1])

	// process the main command line switches and exit
	switch argc {
	case "version":
		version()
		return

	case "help":
		syntax()
		return

	case "checkconfig":
		checkConfig()
		return
	}

	// pull the registration information for our target executable
	// TODO: lets replace this with checkConfig in the future
	targetExe, err := walk.RegistryKeyString(walk.ClassesRootKey(), "starsiege", "Executable")
	if err != nil {
		debugError("Unable to find key HKCR\\starsiege\\Executable")
		return
	}

	// get the path for the executable
	targetFilePath := filepath.Dir(targetExe)

	// continue parsing our command line
	u, err := url.Parse(argc)
	if err != nil {
		debugError("Unable to parse URI")
		return
	}

	// did we accidentally get sent something we can't handle?
	if u.Scheme != "starsiege" {
		debugError("Invalid URI scheme")
		return
	}

	// split the host:port combination up
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		debugError("Unable to parse host and port")
		return
	}

	// sanity check the ip address
	// TODO: support for hostnames in the future, would likely require a DNS lookup
	ip := net.ParseIP(host)
	if ip == nil {
		debugError("Invalid IP address provided")
		return
	}

	// ensure the port is between 1 and 65535 [16 bits]
	uintport, err := strconv.ParseUint(port, 10, 16)
	if err != nil || uintport == 0 {
		debugError("Invalid port provided")
		return
	}

	// prepare our final arguments
	args := fmt.Sprintf("+connect IP:%s:%d", ip.String(), uintport)

	// display a preview of the shellexec call if debugging
	if DEBUG != "" {
		result := win.MessageBox(win.HWND_DESKTOP, L(targetExe, args), L("Execute?"), win.MB_OKCANCEL|win.MB_ICONQUESTION)
		if result != win.IDOK {
			return
		}
	}

	// yeet
	win.ShellExecute(win.HWND_DESKTOP, L(""), L(targetExe), L(args), L(targetFilePath), win.SW_SHOW)
}

// debugError is a quick shortcut for a messagebox
func debugError(text ...string) {
	if DEBUG != "" {
		win.MessageBox(win.HWND_DESKTOP, L(text...), L("Errors Encountered"), win.MB_OK|win.MB_ICONERROR)
	}
}

// checkConfig ensures our key is readable
// TODO: some duplicated code here, but I'm a bit too tired to refactor this
func checkConfig() {
	runExe, err := walk.RegistryKeyString(walk.ClassesRootKey(), "starsiege", "Executable")
	if err != nil {
		debugError("Unable to find key HKCR\\starsiege\\Executable")
		return
	}
	win.MessageBox(win.HWND_DESKTOP, L(runExe), L("Configuration Check"), win.MB_OK|win.MB_ICONINFORMATION)
}

// syntax shows what command line arguments we can accept
func syntax() {
	out := []string{
		"syntax:",
		filepath.Base(os.Args[0]),
		"<help|version|checkconfig|{starsiege:// uri}>",
	}
	win.MessageBox(win.HWND_DESKTOP, L(out...), L("Syntax"), win.MB_OK|win.MB_ICONINFORMATION)
}

// version shows the number associated with the build and its date and time of compiling
func version() {
	out := []string{
		"Neo's Starsiege Launcher\n",
		"Copyright (C) 2021 Starsiege Players Community, et al.\n",
		fmt.Sprintf("Version: %s %s\n", VERSION, DEBUG),
		fmt.Sprintf("Compiled: %s %s", DATE, TIME),
	}

	win.MessageBox(win.HWND_DESKTOP, L(out...), L("Invalid commandline options"), win.MB_OK|win.MB_ICONINFORMATION)
}

// L is a flavored replica of the msvc L"" macro
// it provides the extra feature of joining
// multiple strings together with spaces
func L(input ...string) *uint16 {
	hold := strings.Join(input, " ")
	return win.StringToBSTR(hold)
}
