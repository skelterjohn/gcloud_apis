package discovery

import (
	"github.com/googlecloudplatform/gcloud/gcloud_apis/discovery_docs"
	"strings"
	"testing"
)

var simpleDoc = `
{
  "discoveryVersion": "v1"
}
`

func TestSimpleLoad(t *testing.T) {
	doc, err := LoadDiscovery(strings.NewReader(simpleDoc))
	if err != nil {
		t.Errorf(err.Error())
	}
	if doc.DiscoveryVersion != "v1" {
		t.Errorf("did not get version, expected %q, got %q", "v1", doc.DiscoveryVersion)
	}
}

func TestLoadAutoscalerV1beta2(t *testing.T) {
	doc, err := LoadDiscovery(strings.NewReader(discovery_docs.AutoscalerV1beta2))
	if err != nil {
		t.Errorf(err.Error())
	}
	if doc.DiscoveryVersion != "v1" {
		t.Errorf("did not get version, expected %q, got %q", "v1", doc.DiscoveryVersion)
	}
}
