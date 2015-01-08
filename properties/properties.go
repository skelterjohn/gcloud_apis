package properties

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

type Properties struct {
	sections map[string]*Section
}

type Section struct {
	values map[string]*Property
}

type Property struct {
	value string
}

var (
	sectHdr   = regexp.MustCompile(`^\s*\[([a-zA-Z_]+)\]\s*$`)
	keyVal    = regexp.MustCompile(`^\s*([a-zA-Z_]+)\s*=\s*([^\s]+)\s*$`)
	blankLine = regexp.MustCompile(`^\s*$`)
)

func NewProperties() *Properties {
	return &Properties{
		sections: map[string]*Section{
			"auth": &Section{
				values: map[string]*Property{
					"auth_host":              &Property{},
					"token_host":             &Property{},
					"disable_ssl_validation": &Property{},
					"client_id":              &Property{},
					"client_secret":          &Property{},
				},
			},
			"component_manager": &Section{
				values: map[string]*Property{
					"additional_repositories": &Property{},
					"disable_update_check":    &Property{},
					"fixed_sdk_version":       &Property{},
					"snapshop_url":            &Property{},
				},
			},
			"core": &Section{
				values: map[string]*Property{
					"account":                          &Property{},
					"credentialed_hosted_repo_domains": &Property{},
					"disable_color":                    &Property{},
					"disable_prompts":                  &Property{},
					"disable_usage_reporting":          &Property{},
					"api_host":                         &Property{},
					"project":                          &Property{},
					"user_output_enabled":              &Property{},
					"verbosity":                        &Property{},
				},
			},
			"compute": &Section{
				values: map[string]*Property{
					"region": &Property{},
					"zone":   &Property{},
				},
			},
		},
	}
}

type readCounter struct {
	io.Reader
	n int64
}

func (rc *readCounter) Read(buf []byte) (n int, err error) {
	n, err = rc.Reader.Read(buf)
	rc.n += int64(n)
	return
}

type writeCounter struct {
	io.Writer
	n int64
}

func (wc *writeCounter) Write(buf []byte) (n int, err error) {
	n, err = wc.Writer.Write(buf)
	wc.n += int64(n)
	return
}

func (prop *Properties) ReadFrom(r io.Reader) (int64, error) {
	rc := readCounter{r, 0}
	s := bufio.NewScanner(&rc)
	sect := &Section{}
	sectname := ""
	skip := true
	for s.Scan() {
		l := s.Text()
		if blankLine.FindStringSubmatch(l) != nil {
			continue
		}
		hdr := sectHdr.FindStringSubmatch(l)
		if hdr != nil {
			sectname = hdr[1]
			_, ok := prop.sections[sectname]
			skip = !ok
			if skip {
				continue
			}
			sect = prop.sections[sectname]
		} else {
			if skip {
				continue
			}
			val := keyVal.FindStringSubmatch(l)
			if val == nil {
				continue
			}
			propname := val[1]
			_, match := sect.values[propname]
			if !match {
				continue
			}
			prop := sect.values[propname]
			prop.value = val[2]
		}
	}
	return rc.n, s.Err()
}

// return value of given property name from given section
func (properties *Properties) Get(section string, propname string) (string, error) {
	sect, ok := properties.sections[section]
	if !ok {
		return "", fmt.Errorf("Unknown Properties section %s", section)
	}
	prop, ok := sect.values[propname]
	if !ok {
		return "", fmt.Errorf("Unknown property in section %s: %s", section, propname)
	}
	return prop.value, nil
}

func (properties *Properties) Set(section string, propname string, value string) error {
	sect, ok := properties.sections[section]
	if !ok {
		return fmt.Errorf("Property assignment to unknown section %s", section)
	}
	prop, ok := sect.values[propname]
	if !ok {
		return fmt.Errorf("Property assignment to unknown property %s in section %s", propname, section)
	}
	prop.value = value
	return nil
}

// Write properties to properties file
func (properties *Properties) WriteTo(w io.Writer) (int64, error) {
	wc := writeCounter{w, 0}
	for sname, section := range properties.sections {
		_, err := fmt.Fprintf(&wc, "[%s]\n", sname)
		if err != nil {
			return wc.n, err
		}
		for name, prop := range section.values {
			_, err := fmt.Fprintf(&wc, "%s = %s\n", name, prop.value)
			if err != nil {
				return wc.n, err
			}
		}
	}
	return wc.n, nil
}
func GetWorkspacePropFile() string {
	curdir, _ := filepath.Abs(".")
	for len(curdir) > 1 {
		path := filepath.Join(curdir, ".gcloud", "properties")
		if _, err := os.Stat(path); err == nil {
			return path
		}
		curdir = filepath.Dir(curdir)
	}
	return ""
}

// return properties under $HOME/config
func GetHomeDirPropFile() string {
	dir := os.Getenv("CLOUDSDK_CONFIG")
	if len(dir) == 0 {
		if usr, err := user.Current(); err == nil {
			dir = filepath.Join(usr.HomeDir, ".config", "gcloud")
		} else {
			return ""
		}
	}
	path := filepath.Join(dir, "properties")
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return ""
}

func GetInstallDirPropFile() string {
	path := ""
	progname := os.Args[0]
	if progname[0] == '/' {
		dir := filepath.Dir(progname)
		dir = filepath.Dir(dir)
		path = filepath.Join(dir, "properties")
	} else {
		pathvar := os.Getenv("PATH")
		for _, dir := range filepath.SplitList(pathvar) {
			dir = filepath.Dir(dir)
			_, name := filepath.Split(dir)
			if name == "google-cloud-sdk" {
				path = filepath.Join(dir, "properties")
				break
			}
		}
	}
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return ""
}

func (properties *Properties) LoadPropertiesFiles() error {
	for i := 0; i <= 2; i++ {
		var path string
		switch i {
		case 0:
			path = GetInstallDirPropFile()
		case 1:
			path = GetHomeDirPropFile()
		case 2:
			path = GetWorkspacePropFile()
		}
		if len(path) > 0 {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = properties.ReadFrom(file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
