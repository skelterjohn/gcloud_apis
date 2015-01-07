package properties

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type SimpleTest struct {
	section  string
	property string
	value    string
}

type TestCase struct {
	install string
	home    string
	working string
	tests   []SimpleTest
}

var (
	pfile1 string = "[core]\ndisable_usage_reporting = True\n account = dgk@research.att.com\nproject = cloudsdktest"
	pfile2 string = "[korn]\ndisable_usage_reporting = True\n account = dgk@research.att.com\nproject = cloudsdktest\n[compute] \n zone = us-east "
	pfile3 string = " [compute]\nzone = us-central \n account = dgk@research.att.com\n"
	pfile4 string = "[core]\ndisable_usage_reporting = False\n account = potus@gmail.com\nproject = cloudsdktest"
	pfile5 string = "[core]\ndisable_usage_reporting = True\n \n[core] \n account = dgk@research.att.com\nproject = cloudsdktest"
	pfile6 string = "[core]\ndisable_usage_reporting = True\n \n[core] ]n account = dgk@research.att.com\nproject = cloudsdktest"
)
var TestCases = []TestCase{
	{"", pfile1, "", []SimpleTest{
		{"core", "project", "cloudsdktest"},
		{"core", "disable_usage_reporting", "True"},
	}},
	{"", pfile2, "", []SimpleTest{
		{"core", "project", ""},
		{"compute", "zone", "us-east"},
	}},
	{pfile4, pfile1, "", []SimpleTest{
		{"core", "account", "dgk@research.att.com"},
		{"core", "disable_usage_reporting", "True"},
	}},
	{pfile1, pfile4, "", []SimpleTest{
		{"core", "account", "potus@gmail.com"},
		{"core", "disable_usage_reporting", "False"},
	}},
	{pfile4, pfile1, pfile3, []SimpleTest{
		{"compute", "zone", "us-central"},
		{"core", "account", "dgk@research.att.com"},
	}},
	{"", pfile5, "", []SimpleTest{
		{"core", "project", "cloudsdktest"},
		{"core", "disable_usage_reporting", "True"},
	}},
	{"", pfile6, "", []SimpleTest{
		{"core", "project", "cloudsdktest"},
		{"core", "disable_usage_reporting", "True"},
	}},
}

func SetupDirs(tmpdir string) ([3]string, error) {
	var err error
	var paths [3]string
	dir := filepath.Join(tmpdir, "home", ".config", "gcloud")
	err = os.Setenv("CLOUDSDK_CONFIG", dir)
	if err = os.MkdirAll(dir, 0755); err != nil {
		return paths, err
	}
	os.Args[0] = filepath.Join(dir, "gcloud")
	dir = filepath.Join(tmpdir, "install/google-cloud-sdk/bin")
	if err = os.MkdirAll(dir, 0755); err != nil {
		return paths, err
	}
	dir = filepath.Join(tmpdir, "workspace/.gcloud/foobar")
	if err = os.MkdirAll(dir, 0755); err != nil {
		return paths, err
	}
	if err = os.Chdir(dir); err != nil {
		return paths, err
	}
	paths[0] = filepath.Join(tmpdir, "install/google-cloud-sdk/properties")
	paths[1] = filepath.Join(tmpdir, "home/.config/gcloud/properties")
	paths[2] = filepath.Join(tmpdir, "workspace/.gcloud/properties")
	return paths, err
}

func CreateOrRemove(path string, value string) error {
	var err error
	if value != "" {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(value)
		if err != nil {
			return err
		}
	} else {
		err = os.Remove(path)
	}
	return err
}

func runTest(p *Properties, t *testing.T, item SimpleTest, i, j int) *Properties {
	err := p.LoadPropertiesFiles()
	if err != nil {
		t.Errorf("LoadProperties failed err=%s", err)
	}
	value, err := p.Get(item.section, item.property)
	if err != nil {
		t.Errorf("Test %d.%d: In section %s, property %s should be %s but has err=%s", i, j, item.section, item.property, item.value, err)
	}
	if value != item.value {
		t.Errorf("Test %d.%d: In section %s, property %s should be '%s' but is '%s'", i, j, item.section, item.property, item.value, value)
	}
	return p
}

func TestPropertiesGetSimple(t *testing.T) {
	var p *Properties
	var item SimpleTest
	var j int
	tmpdir, err := ioutil.TempDir("", "prop")
	if err != nil {
		t.Fatalf("unable to setup tests %s", err)
	}
	defer os.RemoveAll(tmpdir)
	paths, err := SetupDirs(tmpdir)
	if err != nil {
		t.Fatalf("unable to setup tests %s", err)
	}
	for i, test := range TestCases {
		CreateOrRemove(paths[0], test.install)
		CreateOrRemove(paths[1], test.home)
		CreateOrRemove(paths[2], test.working)
		for j, item = range test.tests {
			p = NewProperties()
			p = runTest(p, t, item, i, j)
		}
		j := len(test.tests) + 1
		CreateOrRemove(paths[0], "")
		CreateOrRemove(paths[1], "")
		file, err := os.Create(paths[2])
		if err != nil {
			t.Errorf("unable to create file %s", paths[2])
		}
		defer file.Close()
		_, err = p.WriteTo(file)
		if err != nil {
			t.Errorf("unable to dump properties into %s", paths[2])
		}
		runTest(p, t, item, i, j)
	}
}
