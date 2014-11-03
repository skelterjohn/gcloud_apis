package properties

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func newSection() *Section {
	return &Section{
		values: make(map[string]*Property),
	}
}

func NewProperties() *Properties {
	properties := Properties{}
	properties.sections = map[string]*Section{}
	auth := newSection()
	auth.values["auth_host"] = &Property{}
	auth.values["token_host"] = &Property{}
	auth.values["disable_ssl_validation"] = &Property{}
	auth.values["client_id"] = &Property{}
	auth.values["client_secret"] = &Property{}
	properties.sections["auth"] = auth
	component_manager := newSection()
	component_manager.values["additional_repositories"] = &Property{}
	component_manager.values["disable_update_check"] = &Property{}
	component_manager.values["fixed_sdk_version"] = &Property{}
	component_manager.values["snapshop_url"] = &Property{}
	properties.sections["component_manager"] = component_manager
	core := newSection()
	core.values["account"] = &Property{}
	core.values["credentialed_hosted_repo_domains"] = &Property{}
	core.values["disable_color"] = &Property{}
	core.values["disable_prompts"] = &Property{}
	core.values["disable_usage_reporting"] = &Property{}
	core.values["api_host"] = &Property{}
	core.values["project"] = &Property{}
	core.values["user_output_enabled"] = &Property{}
	core.values["verbosity"] = &Property{}
	properties.sections["core"] = core
	compute := newSection()
	compute.values["region"] = &Property{}
	compute.values["zone"] = &Property{}
	properties.sections["compute"] = compute
	return &properties
}

func (prop *Properties) ReadFrom(r io.Reader) error {
	var err error
	s := bufio.NewScanner(r)
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
	return err
}

// return value of given property name from given section
func (properties *Properties) Get(section string, propname string) (string, error) {
	var prop *Property
	sect, ok := properties.sections[section]
	if !ok {
		return "", fmt.Errorf("Unknown Properties section %s", section)
	}
	prop, ok = sect.values[propname]
	if !ok {
		return "", fmt.Errorf("Unknown property in section %s", propname, section)
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
		return fmt.Errorf("Property assignment to unknown property %s in section", propname, section)
	}
	prop.value = value
	return nil
}

// Write properties to properties file
func (properties *Properties) WriteTo(w io.Writer) error {
	for sname, section := range properties.sections {
		_, err := fmt.Fprintf(w, "[%s]\n", sname)
		if err != nil {
			return err
		}
		for name, prop := range section.values {
			_, err := fmt.Fprintf(w, "%s = %s\n", name, prop.value)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
	} else {
		return ""
	}
}

func (properties *Properties) LoadPropertiesFiles() error {
	var path string
	var i int
	var err error
	for i = 0; i <= 2; i++ {
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
				log.Fatal(err)
			}
			defer file.Close()
			err = properties.ReadFrom(file)
		}
	}
	return err
}
