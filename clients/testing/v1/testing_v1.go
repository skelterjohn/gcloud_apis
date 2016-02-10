// Package testing provides access to the Google Cloud Testing API.
//
// See https://developers.google.com/cloud-test-lab/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/testing/v1"
//   ...
//   testingService, err := testing.New(oauthHttpClient)
package testing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "testing:v1"
const apiName = "testing"
const apiVersion = "v1"
const basePath = "https://testing.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	s.TestEnvironmentCatalog = NewTestEnvironmentCatalogService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService

	TestEnvironmentCatalog *TestEnvironmentCatalogService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Devices = NewProjectsDevicesService(s)
	rs.TestMatrices = NewProjectsTestMatricesService(s)
	rs.Webdriver = NewProjectsWebdriverService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Devices *ProjectsDevicesService

	TestMatrices *ProjectsTestMatricesService

	Webdriver *ProjectsWebdriverService
}

func NewProjectsDevicesService(s *Service) *ProjectsDevicesService {
	rs := &ProjectsDevicesService{s: s}
	return rs
}

type ProjectsDevicesService struct {
	s *Service
}

func NewProjectsTestMatricesService(s *Service) *ProjectsTestMatricesService {
	rs := &ProjectsTestMatricesService{s: s}
	return rs
}

type ProjectsTestMatricesService struct {
	s *Service
}

func NewProjectsWebdriverService(s *Service) *ProjectsWebdriverService {
	rs := &ProjectsWebdriverService{s: s}
	return rs
}

type ProjectsWebdriverService struct {
	s *Service
}

func NewTestEnvironmentCatalogService(s *Service) *TestEnvironmentCatalogService {
	rs := &TestEnvironmentCatalogService{s: s}
	return rs
}

type TestEnvironmentCatalogService struct {
	s *Service
}

// AndroidDevice: A single Android device.
type AndroidDevice struct {
	// AndroidModelId: The id of the Android device to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidModelId string `json:"androidModelId,omitempty"`

	// AndroidVersionId: The id of the Android OS version to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidVersionId string `json:"androidVersionId,omitempty"`

	// Locale: The locale the test device used for testing.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Locale string `json:"locale,omitempty"`

	// Orientation: How the device is oriented during the test.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Orientation string `json:"orientation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidModelId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidDevice) MarshalJSON() ([]byte, error) {
	type noMethod AndroidDevice
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidDeviceCatalog: The currently supported Android devices.
type AndroidDeviceCatalog struct {
	// Models: The set of supported Android device models.
	// @OutputOnly
	Models []*AndroidModel `json:"models,omitempty"`

	// RuntimeConfiguration: The set of supported runtime
	// configurations.
	// @OutputOnly
	RuntimeConfiguration *AndroidRuntimeConfiguration `json:"runtimeConfiguration,omitempty"`

	// Versions: The set of supported Android OS versions.
	// @OutputOnly
	Versions []*AndroidVersion `json:"versions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Models") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidDeviceCatalog) MarshalJSON() ([]byte, error) {
	type noMethod AndroidDeviceCatalog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidInstrumentationTest: A test of an Android application that can
// control an Android component
// independently of its normal lifecycle.
// Android instrumentation tests run an application APK and test APK
// inside the
// same process on a virtual or physical AndroidDevice.  They also
// specify
// a test runner class, such as com.google.GoogleTestRunner, which can
// vary
// on the specific instrumentation framework chosen.
//
// See <http://developer.android.com/tools/testing/testing_android.html>
// for
// more information on types of Android tests.
type AndroidInstrumentationTest struct {
	// AppApk: The APK for the application under test.
	// Required
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// TestApk: The APK containing the test code to be executed.
	// Required
	TestApk *FileReference `json:"testApk,omitempty"`

	// TestPackageId: The java package for the test to be
	// executed.
	// Optional, default is determined by examining the application's
	// manifest.
	TestPackageId string `json:"testPackageId,omitempty"`

	// TestRunnerClass: The InstrumentationTestRunner class.
	// Optional, default is determined by examining the application's
	// manifest.
	TestRunnerClass string `json:"testRunnerClass,omitempty"`

	// TestTargets: Each target must be fully qualified with the package
	// name or class name,
	// in one of these formats:
	//  - "package package_name"
	//  - "class package_name.class_name"
	//  - "class package_name.class_name#method_name"
	//
	// If empty, all targets in the module will be run.
	TestTargets []string `json:"testTargets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidInstrumentationTest) MarshalJSON() ([]byte, error) {
	type noMethod AndroidInstrumentationTest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidMatrix: A set of Android device configuration permutations is
// defined by the
// the cross-product of the given axes.  Internally, the given
// AndroidMatrix
// will be expanded into a set of AndroidDevices.
//
// Only supported permutations will be instantiated.  Invalid
// permutations
// (e.g., incompatible models/versions) are ignored.
type AndroidMatrix struct {
	// AndroidModelIds: The ids of the set of Android device to be used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidModelIds []string `json:"androidModelIds,omitempty"`

	// AndroidVersionIds: The ids of the set of Android OS version to be
	// used.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	AndroidVersionIds []string `json:"androidVersionIds,omitempty"`

	// Locales: The set of locales the test device will enable for
	// testing.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Locales []string `json:"locales,omitempty"`

	// Orientations: The set of orientations to test with.
	// Use the EnvironmentDiscoveryService to get supported
	// options.
	// Required
	Orientations []string `json:"orientations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidModelIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidMatrix) MarshalJSON() ([]byte, error) {
	type noMethod AndroidMatrix
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidModel: A description of an Android device tests may be run on.
type AndroidModel struct {
	// Brand: The company that this device is branded with.
	// Example: "Google", "Samsung"
	// @OutputOnly
	Brand string `json:"brand,omitempty"`

	// Codename: The name of the industrial design.
	// This corresponds to android.os.Build.DEVICE
	// @OutputOnly
	Codename string `json:"codename,omitempty"`

	// Form: Whether this device is virtual or physical.
	// @OutputOnly
	//
	// Possible values:
	//   "DEVICE_FORM_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VIRTUAL" - A software stack that simulates the device
	//   "PHYSICAL" - Actual hardware
	Form string `json:"form,omitempty"`

	// Id: The unique opaque id for this model.
	// Use this for invoking the TestExecutionService.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Manufacturer: The manufacturer of this device.
	// @OutputOnly
	Manufacturer string `json:"manufacturer,omitempty"`

	// Name: The human-readable marketing name for this device
	// model.
	// Examples: "Nexus 5", "Galaxy S5"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// ScreenX: Screen size in the horizontal (X) dimension measured in
	// pixels.
	// @OutputOnly
	ScreenX int64 `json:"screenX,omitempty"`

	// ScreenY: Screen size in the vertical (Y) dimension measured in
	// pixels.
	// @OutputOnly
	ScreenY int64 `json:"screenY,omitempty"`

	// SupportedAbis: The list of supported ABIs for this device.
	// This corresponds to either android.os.Build.SUPPORTED_ABIS (for API
	// level
	// 21 and above) or android.os.Build.CPU_ABI/CPU_ABI2.
	// @OutputOnly
	SupportedAbis []string `json:"supportedAbis,omitempty"`

	// SupportedVersionIds: The set of Android versions this device
	// supports.
	// @OutputOnly
	SupportedVersionIds []string `json:"supportedVersionIds,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Brand") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidModel) MarshalJSON() ([]byte, error) {
	type noMethod AndroidModel
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidMonkeyTest: A test of an Android application that uses the
// UI/Application Exerciser
// Monkey from the Android SDK. (Not to be confused with the
// "monkeyrunner"
// tool, which is also included in the SDK.)
//
// See http://developer.android.com/tools/help/monkey.html for details.
type AndroidMonkeyTest struct {
	// AppApk: The APK for the application under test.
	// Required
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// EventCount: Number of random monkey events (e.g. clicks, touches) to
	// generate.
	// Defaults to 2000.
	EventCount int64 `json:"eventCount,omitempty"`

	// EventDelay: Fixed delay between events.
	// Defaults to 10ms.
	EventDelay string `json:"eventDelay,omitempty"`

	// RandomSeed: Seed value for pseudo-random number generator.
	// Note that, although specifying a seed causes the monkey to generate
	// the
	// same sequence of events, it does not guarantee that a particular
	// outcome
	// will be reproducible across runs.
	// Optional
	RandomSeed int64 `json:"randomSeed,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidMonkeyTest) MarshalJSON() ([]byte, error) {
	type noMethod AndroidMonkeyTest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidRoboTest: A test of an android application that explores the
// application on a virtual
// or physical Android Device, finding culprits and crashes as it goes.
type AndroidRoboTest struct {
	// AppApk: The APK for the application under test.
	// Required
	AppApk *FileReference `json:"appApk,omitempty"`

	// AppInitialActivity: The initial activity that should be used to start
	// the app.
	// Optional
	AppInitialActivity string `json:"appInitialActivity,omitempty"`

	// AppPackageId: The java package for the application under
	// test.
	// Optional, default is determined by examining the application's
	// manifest.
	AppPackageId string `json:"appPackageId,omitempty"`

	// MaxDepth: The max depth of the traversal stack Robo can explore.
	// Needs to be at least
	// 2 to make Robo explore the app beyond the first activity.
	// Default is 50.
	// Optional
	MaxDepth int64 `json:"maxDepth,omitempty"`

	// MaxSteps: The max number of steps Robo can execute.
	// Default is no limit.
	// Optional
	MaxSteps int64 `json:"maxSteps,omitempty"`

	// RandomizeSteps: Whether Robo follows a random order of steps on a
	// given activity state.
	// Optional
	RandomizeSteps bool `json:"randomizeSteps,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppApk") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidRoboTest) MarshalJSON() ([]byte, error) {
	type noMethod AndroidRoboTest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidRuntimeConfiguration: Configuration that can be selected at
// the time a test is run.
type AndroidRuntimeConfiguration struct {
	// Locales: The set of available locales.
	// @OutputOnly
	Locales []*Locale `json:"locales,omitempty"`

	// Orientations: The set of available orientations.
	// @OutputOnly
	Orientations []*Orientation `json:"orientations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Locales") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidRuntimeConfiguration) MarshalJSON() ([]byte, error) {
	type noMethod AndroidRuntimeConfiguration
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidVersion: A version of the Android OS
type AndroidVersion struct {
	// ApiLevel: The API level for this Android version.
	// Examples: 18, 19
	// @OutputOnly
	ApiLevel int64 `json:"apiLevel,omitempty"`

	// CodeName: The code name for this Android version.
	// Examples: "JellyBean", "KitKat"
	// @OutputOnly
	CodeName string `json:"codeName,omitempty"`

	// Distribution: Market share for this version.
	// @OutputOnly
	Distribution *Distribution `json:"distribution,omitempty"`

	// Id: An opaque id for this Android version.
	// Use this id to invoke the TestExecutionService.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// ReleaseDate: The date this Android version became available in the
	// market.
	// @OutputOnly
	ReleaseDate *Date `json:"releaseDate,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview", "deprecated"
	Tags []string `json:"tags,omitempty"`

	// VersionString: A string representing this version of the Android
	// OS.
	// Examples: "4.3", "4.4"
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiLevel") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidVersion) MarshalJSON() ([]byte, error) {
	type noMethod AndroidVersion
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BlobstoreFile: Reference to a blob in Blobstore.
type BlobstoreFile struct {
	// BlobId: A blob ID.
	// Example: /android_test/blobs/4e9AAT9sqHRY_oBBzIKHSEFgg
	BlobId string `json:"blobId,omitempty"`

	// Md5Hash: The MD5 hash of the referenced blob. (This is necessary to
	// create a
	// Bigstore object directly from the Blobstore reference.)
	Md5Hash string `json:"md5Hash,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BlobId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BlobstoreFile) MarshalJSON() ([]byte, error) {
	type noMethod BlobstoreFile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Browser: An available browser.
type Browser struct {
	// AndroidCatalog: The catalog of Android devices for which we offer
	// this browser.
	// @OutputOnly
	AndroidCatalog *AndroidDeviceCatalog `json:"androidCatalog,omitempty"`

	// Id: A human readable id for this Browser version.
	// Use this id to invoke the TestExecutionService.
	// Examples: "chrome-stable-channel", "firefox-beta-channel"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// LinuxCatalog: The catalog of Linux machines which we offer this
	// browser.
	// @OutputOnly
	LinuxCatalog *LinuxMachineCatalog `json:"linuxCatalog,omitempty"`

	// Name: A string representing the browser name.
	// Examples: "chrome", "firefox", "ie"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Release: The release of the browser.
	// Examples: "stable-channel", "beta-channel", "10" (for ie),
	// etc
	// @OutputOnly
	Release string `json:"release,omitempty"`

	// VersionString: A string representing the version of the
	// browser.
	// Examples:  "42.12.34.1234", "37.01", "10.0.9200.16384" (for
	// ie)
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`

	// WindowsCatalog: The catalog of Windows machines which we offer this
	// browser.
	// @OutputOnly
	WindowsCatalog *WindowsMachineCatalog `json:"windowsCatalog,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidCatalog") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Browser) MarshalJSON() ([]byte, error) {
	type noMethod Browser
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CancelTestMatrixResponse: Response containing the state of a
// cancelled test matrix.
type CancelTestMatrixResponse struct {
	// TestState: The rolled-up state of the test matrix just before it was
	// cancelled.
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	TestState string `json:"testState,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "TestState") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CancelTestMatrixResponse) MarshalJSON() ([]byte, error) {
	type noMethod CancelTestMatrixResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ClientInfo: Information about the client which invoked the test.
type ClientInfo struct {
	// Name: Client name, such as gcloud.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ClientInfo) MarshalJSON() ([]byte, error) {
	type noMethod ClientInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ConnectionInfo: Information needed to connect to services running on
// the virtual device. The
// ssh_port is used to connect to the device, and then the adb_port and
// vnc_port
// on the device can be forwarded to two local ports, to which adb and
// vnc can
// connect, respectively.
//
// All of the fields in this message are provided by the backend.
type ConnectionInfo struct {
	// AdbPort: Port for ADB (e.g. 5555)
	// NOT user-specified
	// Required
	AdbPort int64 `json:"adbPort,omitempty"`

	// IpAddress: IP address of the device.
	// NOT user-specified
	// Required
	IpAddress string `json:"ipAddress,omitempty"`

	// SshPort: Port for SSH (e.g. 22)
	// NOT user-specified
	// Required
	SshPort int64 `json:"sshPort,omitempty"`

	// VncPassword: Password for VNC
	// NOT user-specified
	// Required
	VncPassword string `json:"vncPassword,omitempty"`

	// VncPort: Port for VNC (e.g. 6444)
	// NOT user-specified
	// Required
	VncPort int64 `json:"vncPort,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdbPort") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ConnectionInfo) MarshalJSON() ([]byte, error) {
	type noMethod ConnectionInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Date: Represents a whole calendar date, e.g. date of birth. The time
// of day and
// time zone are either specified elsewhere or are not significant. The
// date
// is relative to the Proleptic Gregorian Calendar. The day may be 0
// to
// represent a year and month where the day is not significant, e.g.
// credit card
// expiration date. The year may be 0 to represent a month and day
// independent
// of year, e.g. anniversary date. Related types are
// google.type.TimeOfDay
// and `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year/month where the day is not significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type noMethod Date
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Device: A GCE virtual Android device instance.
type Device struct {
	// AndroidDevice: The Android device
	// configuration.
	// User-specified
	// Required
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	// CreationTime: The time the request to create this device was
	// received.
	// NOT user-specified
	CreationTime string `json:"creationTime,omitempty"`

	// DeviceDetails: Information about the backing GCE instance and
	// connection.
	// NOT user-specified
	DeviceDetails *DeviceDetails `json:"deviceDetails,omitempty"`

	// Id: Unique id set by the backend.
	// NOT user-specified
	Id string `json:"id,omitempty"`

	// ProjectId: Project id set by the backend.
	// NOT user-specified
	ProjectId string `json:"projectId,omitempty"`

	// State: State of the device.
	// NOT user-specified
	//
	// Possible values:
	//   "DEVICE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "PREPARING" - The device is in the process of spinning up.
	//   "READY" - The device is created and ready to use.
	//   "CLOSED" - The device has been closed.
	//   "DEVICE_ERROR" - There has been an error.
	State string `json:"state,omitempty"`

	// StateDetails: Details about the state of the device.
	// NOT user-specified
	StateDetails *DeviceStateDetails `json:"stateDetails,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AndroidDevice") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Device) MarshalJSON() ([]byte, error) {
	type noMethod Device
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DeviceDetails: Details about the GCE instance and connection.
type DeviceDetails struct {
	// ConnectionInfo: Details about the connection to the device.
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`

	// GceInstanceDetails: Details about the GCE instance backing the
	// device.
	GceInstanceDetails *GceInstanceDetails `json:"gceInstanceDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConnectionInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DeviceDetails) MarshalJSON() ([]byte, error) {
	type noMethod DeviceDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DeviceFile: A single device file description.
type DeviceFile struct {
	ObbFile *ObbFile `json:"obbFile,omitempty"`

	RegularFile *RegularFile `json:"regularFile,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ObbFile") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DeviceFile) MarshalJSON() ([]byte, error) {
	type noMethod DeviceFile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DeviceStateDetails: Additional details about the status of the
// device.
type DeviceStateDetails struct {
	// ErrorDetails: If the DeviceState is ERROR, then this string may
	// contain human-readable
	// details about the error.
	ErrorDetails string `json:"errorDetails,omitempty"`

	// ProgressDetails: A human-readable, detailed description of the
	// device's status.
	// For example: "Starting Device\n Device Ready".
	//
	// During the device's lifespan data may be appended to the progress.
	ProgressDetails string `json:"progressDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorDetails") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DeviceStateDetails) MarshalJSON() ([]byte, error) {
	type noMethod DeviceStateDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Distribution: Data about the relative number of devices running
// a
// given configuration of the Android platform.
type Distribution struct {
	// MarketShare: The estimated fraction (0-1) of the total market with
	// this configuration.
	// @OutputOnly
	MarketShare float64 `json:"marketShare,omitempty"`

	// MeasurementTime: The time this distribution was measured.
	// @OutputOnly
	MeasurementTime string `json:"measurementTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MarketShare") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Distribution) MarshalJSON() ([]byte, error) {
	type noMethod Distribution
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Environment: The environment in which the test is run.
type Environment struct {
	// AndroidDevice: An Android device which must be used with an Android
	// test.
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidDevice") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Environment) MarshalJSON() ([]byte, error) {
	type noMethod Environment
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// EnvironmentMatrix: The matrix of environments in which the test is to
// be executed.
type EnvironmentMatrix struct {
	// AndroidMatrix: A matrix of Android devices
	AndroidMatrix *AndroidMatrix `json:"androidMatrix,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AndroidMatrix") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *EnvironmentMatrix) MarshalJSON() ([]byte, error) {
	type noMethod EnvironmentMatrix
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// FileReference: A reference to a file, used for user inputs.
type FileReference struct {
	// Blob: A blob in Blobstore.
	Blob *BlobstoreFile `json:"blob,omitempty"`

	// GcsPath: A path to a file in Google Cloud Storage.
	// Example: gs://build-app-1414623860166/app-debug-unaligned.apk
	GcsPath string `json:"gcsPath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Blob") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FileReference) MarshalJSON() ([]byte, error) {
	type noMethod FileReference
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GceInstanceDetails: This information is provided for the user to look
// up additional details of
// the backing GCE instance. It is assumed the user does not modify
// this
// instance. If so, then the device service makes no guarantees
// about
// device functionality.
type GceInstanceDetails struct {
	// Name: Desired instance name of the device.
	// May be user-specified. If not, the backend will choose a name.
	Name string `json:"name,omitempty"`

	// ProjectId: The GCE project that contains the instance backing this
	// device. If
	// user-specified, must be the same as the project_id in
	// the
	// CreateDeviceRequest.
	ProjectId string `json:"projectId,omitempty"`

	// Zone: Desired GCE zone for the device
	// user-specified
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GceInstanceDetails) MarshalJSON() ([]byte, error) {
	type noMethod GceInstanceDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudStorage: A storage location within Google cloud storage
// (GCS).
type GoogleCloudStorage struct {
	// GcsPath: The path to a directory in GCS that will
	// eventually contain the results for this test.
	// The requesting user must have write access on the bucket in the
	// supplied
	// path.
	GcsPath string `json:"gcsPath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcsPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudStorage) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudStorage
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LinuxMachine: A single Linux machine.
type LinuxMachine struct {
	// VersionId: The version id of the Linux OS to be used.
	// Use the EnvironmentDiscoveryService to get supported options.
	VersionId string `json:"versionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "VersionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LinuxMachine) MarshalJSON() ([]byte, error) {
	type noMethod LinuxMachine
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LinuxMachineCatalog: The currently supported Linux machines.
type LinuxMachineCatalog struct {
	// Versions: The set of supported Linux versions.
	// @OutputOnly
	Versions []*LinuxVersion `json:"versions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Versions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LinuxMachineCatalog) MarshalJSON() ([]byte, error) {
	type noMethod LinuxMachineCatalog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LinuxVersion: A verison of a Linux OS.
type LinuxVersion struct {
	// Id: The unique opaque id for this Linux Version.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Tags: Tags for this version.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// VersionString: A string representing this version of the Linux
	// OS.
	// Examples: "debian-7-wheezy-v20150325",
	// "debian-7-wheezy-v30150325"
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LinuxVersion) MarshalJSON() ([]byte, error) {
	type noMethod LinuxVersion
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListDevicesResponse: Response containing a list of devices. Supports
// pagination.
type ListDevicesResponse struct {
	// Devices: The GCE virtual Android devices to be returned.
	Devices []*Device `json:"devices,omitempty"`

	// NextPageToken: The pagination token to retrieve the next page of
	// device results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Devices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListDevicesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListDevicesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListTestMatricesResponse: Response contain a list of Test Matrices.
type ListTestMatricesResponse struct {
	// TestMatrices: The set of test matrices.
	TestMatrices []*TestMatrix `json:"testMatrices,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "TestMatrices") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListTestMatricesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListTestMatricesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListWebDriverResponse: Response containing a list of WebDriver
// environments. Supports pagination.
type ListWebDriverResponse struct {
	// NextPageToken: The pagination token to retrieve the next page of
	// WebDriver results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// WebdriverEnvironments: The WebDriver environments to be returned.
	WebdriverEnvironments []*WebDriver `json:"webdriverEnvironments,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListWebDriverResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListWebDriverResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Locale: A location/region designation for language.
type Locale struct {
	// Id: The id for this locale.
	// Example: "en_US"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this language/locale.
	// Example: "English"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Region: A human-friendy string representing the region for this
	// locale.
	// Example: "United States"
	// Not present for every locale.
	// @OutputOnly
	Region string `json:"region,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Locale) MarshalJSON() ([]byte, error) {
	type noMethod Locale
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type ObbFile struct {
	// Obb: Opaque Binary Blob (OBB) file(s) to install on the device
	Obb *FileReference `json:"obb,omitempty"`

	// ObbFileName: OBB file name which must conform to the format as
	// specified by
	// Android
	// e.g. [main|patch].0300110.com.example.android.obb
	// which will be installed into
	//   <shared-storage>/Android/obb/<package-name>/
	// on the device
	ObbFileName string `json:"obbFileName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Obb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ObbFile) MarshalJSON() ([]byte, error) {
	type noMethod ObbFile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Orientation: Screen orientation of the device.
type Orientation struct {
	// Id: The id for this orientation.
	// Example: "portrait"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this orientation.
	// Example: "portrait"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Orientation) MarshalJSON() ([]byte, error) {
	type noMethod Orientation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RegularFile: A file or directory to install on the device before the
// test starts
type RegularFile struct {
	Content *FileReference `json:"content,omitempty"`

	// DevicePath: Where to put the content on the device, must be a full,
	// whitelisted path.
	// If it exists, it will be completely replaced.
	// TODO(fsteen): Make the following path substitutions available:
	// <p> ${EXTERNAL_STORAGE} - the external storage mount point
	// (/sdcard)
	// <p> ${ANDROID_DATA} - the userdata partition mount point
	// (/data)
	// Note: /data/local/tmp is whitelisted, but /data is not.
	//
	// <p> The corresponding paths (in parentheses) will be made available
	// and
	// treated as implicit path substitutions, so the user may use
	// them
	// interchangeably. E.g. if /sdcard on a particular device does not map
	// to
	// external storage, the system will replace it with the external
	// storage path
	// prefix for that device and copy the file there.
	//
	// <p> It is strongly advised to use the <a
	// href=
	// "http://developer.android.com/reference/android/os/Environment.h
	// tml">
	// Environment API</a> in app and test code to access files on the
	// device in a
	// portable way.
	DevicePath string `json:"devicePath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RegularFile) MarshalJSON() ([]byte, error) {
	type noMethod RegularFile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ResultStorage: Locations where the results of running the test are
// stored.
type ResultStorage struct {
	// GoogleCloudStorage: Required.
	GoogleCloudStorage *GoogleCloudStorage `json:"googleCloudStorage,omitempty"`

	// ToolResultsExecution: The tool results execution that results are
	// written to.
	// @OutputOnly
	ToolResultsExecution *ToolResultsExecution `json:"toolResultsExecution,omitempty"`

	// ToolResultsHistory: The tool results history that contains the tool
	// results execution that
	// results are written to.
	//
	// If not provided the service will choose an appropriate value.
	ToolResultsHistory *ToolResultsHistory `json:"toolResultsHistory,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GoogleCloudStorage")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ResultStorage) MarshalJSON() ([]byte, error) {
	type noMethod ResultStorage
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestDetails: Additional details about the progress of the running
// test.
type TestDetails struct {
	// ErrorMessage: If the TestState is ERROR, then this string will
	// contain human-readable
	// details about the error.
	// @OutputOnly
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ProgressMessages: Human-readable, detailed descriptions of the test's
	// progress.
	// For example: "Provisioning a device", "Starting Test".
	//
	// During the course of execution new data may be appended
	// to the end of progress_messages.
	// @OutputOnly
	ProgressMessages []string `json:"progressMessages,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestDetails) MarshalJSON() ([]byte, error) {
	type noMethod TestDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestEnvironmentCatalog: A description of a test environment.
type TestEnvironmentCatalog struct {
	// AndroidDeviceCatalog: Android devices suitable for running Android
	// Instrumentation Tests.
	AndroidDeviceCatalog *AndroidDeviceCatalog `json:"androidDeviceCatalog,omitempty"`

	// WebDriverCatalog: WebDriver environments suitable for running web
	// tests.
	WebDriverCatalog *WebDriverCatalog `json:"webDriverCatalog,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AndroidDeviceCatalog") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestEnvironmentCatalog) MarshalJSON() ([]byte, error) {
	type noMethod TestEnvironmentCatalog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestExecution: Specifies a single test to be executed in a single
// environment.
type TestExecution struct {
	// Environment: How the host machine(s) are configured.
	// @OutputOnly
	Environment *Environment `json:"environment,omitempty"`

	// Id: Unique id set by the backend.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// MatrixId: Id of the containing TestMatrix.
	// @OutputOnly
	MatrixId string `json:"matrixId,omitempty"`

	// ProjectId: The cloud project that owns the test
	// execution.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// State: Indicates the current progress of the test execution (e.g.,
	// FINISHED).
	// @OutputOnly
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	State string `json:"state,omitempty"`

	// TestDetails: Additional details about the running test.
	// @OutputOnly
	TestDetails *TestDetails `json:"testDetails,omitempty"`

	// TestSpecification: How to run the test.
	// @OutputOnly
	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`

	// Timestamp: The time this test execution was initially
	// created.
	// @OutputOnly
	Timestamp string `json:"timestamp,omitempty"`

	// ToolResultsStep: Where the results for this execution are
	// written.
	// @OutputOnly
	ToolResultsStep *ToolResultsStep `json:"toolResultsStep,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Environment") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestExecution) MarshalJSON() ([]byte, error) {
	type noMethod TestExecution
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestMatrix: A group of one or more TestExecutions, built by taking
// a
// product of values over a pre-defined set of axes.
type TestMatrix struct {
	// ClientInfo: Information about the client which invoked the test.
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`

	// EnvironmentMatrix: How the host machine(s) are configured.
	// Required
	EnvironmentMatrix *EnvironmentMatrix `json:"environmentMatrix,omitempty"`

	// ProjectId: The cloud project that owns the test matrix.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// ResultStorage: Where the results for the matrix are written.
	// Required
	ResultStorage *ResultStorage `json:"resultStorage,omitempty"`

	// State: Indicates the current progress of the test matrix (e.g.,
	// FINISHED)
	// @OutputOnly
	//
	// Possible values:
	//   "TEST_STATE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "VALIDATING" - The execution or matrix is being validated.
	//   "PENDING" - The execution or matrix is waiting for resources to
	// become available.
	//   "RUNNING" - The execution is currently being processed.
	//
	// Can only be set on an execution.
	//   "FINISHED" - The execution or matrix has terminated normally.
	//
	// On a matrix this means that the matrix level processing completed
	// normally,
	// but individual executions may be in an ERROR state.
	//   "ERROR" - The execution or matrix has stopped because it
	// encountered an
	// infrastructure failure.
	//   "UNSUPPORTED_ENVIRONMENT" - The execution was not run because it
	// corresponds to a unsupported
	// environment.
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ENVIRONMENT" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested environment.
	//
	// Example: requested AndroidVersion is lower than APK's
	// minSdkVersion
	//
	// Can only be set on an execution.
	//   "INCOMPATIBLE_ARCHITECTURE" - The execution was not run because the
	// provided inputs are incompatible with
	// the requested architecture.
	//
	// Example: requested device does not support running the native code
	// in
	// the supplied APK
	//
	// Can only be set on an execution.
	//   "CANCELLED" - The user cancelled the execution.
	//
	// Can only be set on an execution.
	//   "INVALID" - The execution or matrix was not run because the
	// provided inputs are not
	// valid.
	//
	// Examples: input file is not of the expected type, is
	// malformed/corrupt, or
	// was flagged as malware
	State string `json:"state,omitempty"`

	// TestExecutions: The list of test executions that the service creates
	// for this matrix.
	// @OutputOnly
	TestExecutions []*TestExecution `json:"testExecutions,omitempty"`

	// TestMatrixId: Unique id set by the service.
	// @OutputOnly
	TestMatrixId string `json:"testMatrixId,omitempty"`

	// TestSpecification: How to run the test.
	// Required
	TestSpecification *TestSpecification `json:"testSpecification,omitempty"`

	// Timestamp: The time this test matrix was initially
	// created.
	// @OutputOnly
	Timestamp string `json:"timestamp,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestMatrix) MarshalJSON() ([]byte, error) {
	type noMethod TestMatrix
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestSetup: A description of how to set up the device prior to running
// the test
type TestSetup struct {
	FilesToPush []*DeviceFile `json:"filesToPush,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FilesToPush") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestSetup) MarshalJSON() ([]byte, error) {
	type noMethod TestSetup
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestSpecification: A description of how to run the test.
type TestSpecification struct {
	// AndroidInstrumentationTest: An Android instrumentation test.
	AndroidInstrumentationTest *AndroidInstrumentationTest `json:"androidInstrumentationTest,omitempty"`

	// AndroidMonkeyTest: An Android monkey test.
	AndroidMonkeyTest *AndroidMonkeyTest `json:"androidMonkeyTest,omitempty"`

	// AndroidRoboTest: An Android robo test.
	AndroidRoboTest *AndroidRoboTest `json:"androidRoboTest,omitempty"`

	// AutoGoogleLogin: Enables automatic Google account login.
	// If set, the service will automatically generate a Google test
	// account
	// and use it to log into the device, before executing the test. Note
	// that
	// test accounts might be reused.
	// Many applications can be tested more effectively in the context
	// of
	// such an account.
	// Default is false.
	// Optional
	AutoGoogleLogin bool `json:"autoGoogleLogin,omitempty"`

	// TestSetup: Test setup requirements e.g. files to install, bootstrap
	// scripts
	TestSetup *TestSetup `json:"testSetup,omitempty"`

	// TestTimeout: Max time a test execution is allowed to run before it
	// is
	// automatically cancelled.
	TestTimeout string `json:"testTimeout,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AndroidInstrumentationTest") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestSpecification) MarshalJSON() ([]byte, error) {
	type noMethod TestSpecification
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ToolResultsExecution: Represents a tool results execution
// resource.
//
// This has the results of a TestMatrix.
type ToolResultsExecution struct {
	// ExecutionId: A tool results execution ID.
	// @OutputOnly
	ExecutionId string `json:"executionId,omitempty"`

	// HistoryId: A tool results history ID.
	// @OutputOnly
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results
	// execution.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExecutionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ToolResultsExecution) MarshalJSON() ([]byte, error) {
	type noMethod ToolResultsExecution
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ToolResultsHistory: Represents a tool results history resource.
type ToolResultsHistory struct {
	// HistoryId: A tool results history ID.
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results history.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HistoryId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ToolResultsHistory) MarshalJSON() ([]byte, error) {
	type noMethod ToolResultsHistory
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ToolResultsStep: Represents a tool results step resource.
//
// This has the results of a TestExecution.
type ToolResultsStep struct {
	// ExecutionId: A tool results execution ID.
	// @OutputOnly
	ExecutionId string `json:"executionId,omitempty"`

	// HistoryId: A tool results history ID.
	// @OutputOnly
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results
	// step.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// StepId: A tool results step ID.
	// @OutputOnly
	StepId string `json:"stepId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExecutionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ToolResultsStep) MarshalJSON() ([]byte, error) {
	type noMethod ToolResultsStep
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type VMDetails struct {
	// CreationTime: The time this device was initially created.
	// @OutputOnly
	CreationTime string `json:"creationTime,omitempty"`

	// DeviceDetails: Information about the backing GCE instance and
	// connection.
	// @OutputOnly
	DeviceDetails *DeviceDetails `json:"deviceDetails,omitempty"`

	// State: State of the device.
	// @OutputOnly
	//
	// Possible values:
	//   "DEVICE_UNSPECIFIED" - Do not use.  For proto versioning only.
	//   "PREPARING" - The device is in the process of spinning up.
	//   "READY" - The device is created and ready to use.
	//   "CLOSED" - The device has been closed.
	//   "DEVICE_ERROR" - There has been an error.
	State string `json:"state,omitempty"`

	// StateDetails: Details about the state of the device.
	// @OutputOnly
	StateDetails *DeviceStateDetails `json:"stateDetails,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreationTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *VMDetails) MarshalJSON() ([]byte, error) {
	type noMethod VMDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WebDriver: A WebDriver environment.
type WebDriver struct {
	// AndroidDevice: An Android device.
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	// BrowserId: The id of the browser to be used.
	// Use the EnvironmentDiscoveryService to get supported values.
	// Required
	BrowserId string `json:"browserId,omitempty"`

	// Endpoint: The endpoint in host:port format where the target running
	// the specified
	// browser accepts WebDriver protocol commands.
	// @OutputOnly
	Endpoint string `json:"endpoint,omitempty"`

	// Id: Unique id set by the system.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// LinuxMachine: A Linux virtual machine.
	LinuxMachine *LinuxMachine `json:"linuxMachine,omitempty"`

	// ProjectId: The GCE project for this WebDriver test
	// environment.
	// @OutputOnly
	ProjectId string `json:"projectId,omitempty"`

	// SshPublicKey: The public key to be set on the VM in order to SSH into
	// it.
	SshPublicKey string `json:"sshPublicKey,omitempty"`

	// VmDetails: The state details of the target
	// device/machine.
	// @OutputOnly
	VmDetails *VMDetails `json:"vmDetails,omitempty"`

	// WindowsMachine: A Windows virtual machine.
	WindowsMachine *WindowsMachine `json:"windowsMachine,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AndroidDevice") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WebDriver) MarshalJSON() ([]byte, error) {
	type noMethod WebDriver
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WebDriverCatalog: The currently supported WebDriver VM resources.
type WebDriverCatalog struct {
	// Browsers: The set of supported browsers.
	// @OutputOnly
	Browsers []*Browser `json:"browsers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Browsers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WebDriverCatalog) MarshalJSON() ([]byte, error) {
	type noMethod WebDriverCatalog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WebDriverKeepAliveRequest: Request to issue a keep-alive to a
// WebDriver environment instance by
// project and webdriver ids.
type WebDriverKeepAliveRequest struct {
}

// WindowsMachine: A single Windows machine.
type WindowsMachine struct {
	// VersionId: The version id of the Windows OS to be used.
	// Use the EnvironmentDiscoveryService to get supported options.
	VersionId string `json:"versionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "VersionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WindowsMachine) MarshalJSON() ([]byte, error) {
	type noMethod WindowsMachine
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WindowsMachineCatalog: The currently supported Windows machines.
type WindowsMachineCatalog struct {
	// Versions: The set of supported Windows versions.
	// @OutputOnly
	Versions []*WindowsVersion `json:"versions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Versions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WindowsMachineCatalog) MarshalJSON() ([]byte, error) {
	type noMethod WindowsMachineCatalog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WindowsVersion: A version of a Windows OS.
type WindowsVersion struct {
	// Id: The unique opaque id for this Windows Version.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Tags: Tags for this version.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`

	// VersionString: A string representing this version of the Windows
	// OS.
	// Examples: "windows-server-2008-r2-dc-v20150331",
	// windows-7"
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WindowsVersion) MarshalJSON() ([]byte, error) {
	type noMethod WindowsVersion
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "testing.projects.devices.create":

type ProjectsDevicesCreateCall struct {
	s          *Service
	projectId  string
	device     *Device
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new GCE Android device.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to write to
// project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the device type or project does not exist
func (r *ProjectsDevicesService) Create(projectId string, device *Device) *ProjectsDevicesCreateCall {
	c := &ProjectsDevicesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.device = device
	return c
}

// SshPublicKey sets the optional parameter "sshPublicKey": The public
// key to be set on the device in order to SSH into it.
func (c *ProjectsDevicesCreateCall) SshPublicKey(sshPublicKey string) *ProjectsDevicesCreateCall {
	c.urlParams_.Set("sshPublicKey", sshPublicKey)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesCreateCall) Fields(s ...googleapi.Field) *ProjectsDevicesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDevicesCreateCall) Context(ctx context.Context) *ProjectsDevicesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDevicesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.device)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.devices.create" call.
// Exactly one of *Device or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Device.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDevicesCreateCall) Do(opts ...googleapi.CallOption) (*Device, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Device{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new GCE Android device.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the device type or project does not exist",
	//   "flatPath": "v1/projects/{projectId}/devices",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.devices.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCE project under which to create the device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sshPublicKey": {
	//       "description": "The public key to be set on the device in order to SSH into it.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/devices",
	//   "request": {
	//     "$ref": "Device"
	//   },
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.devices.delete":

type ProjectsDevicesDeleteCall struct {
	s          *Service
	projectId  string
	deviceId   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a GCE Android device instance.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) Delete(projectId string, deviceId string) *ProjectsDevicesDeleteCall {
	c := &ProjectsDevicesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesDeleteCall) Fields(s ...googleapi.Field) *ProjectsDevicesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDevicesDeleteCall) Context(ctx context.Context) *ProjectsDevicesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDevicesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.devices.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDevicesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a GCE Android device instance.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/devices/{deviceId}",
	//   "httpMethod": "DELETE",
	//   "id": "testing.projects.devices.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The GCE virtual Android device to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The GCE project that contains the device to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/devices/{deviceId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.devices.get":

type ProjectsDevicesGetCall struct {
	s            *Service
	projectId    string
	deviceId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns the GCE Android device.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the device type or project does not exist
func (r *ProjectsDevicesService) Get(projectId string, deviceId string) *ProjectsDevicesGetCall {
	c := &ProjectsDevicesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesGetCall) Fields(s ...googleapi.Field) *ProjectsDevicesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDevicesGetCall) IfNoneMatch(entityTag string) *ProjectsDevicesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDevicesGetCall) Context(ctx context.Context) *ProjectsDevicesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDevicesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.devices.get" call.
// Exactly one of *Device or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Device.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDevicesGetCall) Do(opts ...googleapi.CallOption) (*Device, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Device{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the GCE Android device.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the device type or project does not exist",
	//   "flatPath": "v1/projects/{projectId}/devices/{deviceId}",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.devices.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The id of the GCE Android virtual device.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The GCE project that contains this device instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/devices/{deviceId}",
	//   "response": {
	//     "$ref": "Device"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "testing.projects.devices.keepalive":

type ProjectsDevicesKeepaliveCall struct {
	s          *Service
	projectId  string
	deviceId   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Keepalive: Issues a keep-alive to a GCE Android device instance.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) Keepalive(projectId string, deviceId string) *ProjectsDevicesKeepaliveCall {
	c := &ProjectsDevicesKeepaliveCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesKeepaliveCall) Fields(s ...googleapi.Field) *ProjectsDevicesKeepaliveCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDevicesKeepaliveCall) Context(ctx context.Context) *ProjectsDevicesKeepaliveCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDevicesKeepaliveCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}/keepalive")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.devices.keepalive" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsDevicesKeepaliveCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Issues a keep-alive to a GCE Android device instance.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/devices/{deviceId}/keepalive",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.devices.keepalive",
	//   "parameterOrder": [
	//     "projectId",
	//     "deviceId"
	//   ],
	//   "parameters": {
	//     "deviceId": {
	//       "description": "The GCE virtual Android device to be issued the keep-alive.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The GCE project that contains the device to be issued the keep-alive.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/devices/{deviceId}/keepalive",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.devices.list":

type ProjectsDevicesListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all the current devices.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) List(projectId string) *ProjectsDevicesListCall {
	c := &ProjectsDevicesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Used to specify the
// max number of device results to be returned.
func (c *ProjectsDevicesListCall) PageSize(pageSize int64) *ProjectsDevicesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Used to request a
// specific page of the device results list.
func (c *ProjectsDevicesListCall) PageToken(pageToken string) *ProjectsDevicesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesListCall) Fields(s ...googleapi.Field) *ProjectsDevicesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDevicesListCall) IfNoneMatch(entityTag string) *ProjectsDevicesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDevicesListCall) Context(ctx context.Context) *ProjectsDevicesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDevicesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.devices.list" call.
// Exactly one of *ListDevicesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListDevicesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDevicesListCall) Do(opts ...googleapi.CallOption) (*ListDevicesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListDevicesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the current devices.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/devices",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.devices.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Used to specify the max number of device results to be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Used to request a specific page of the device results list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The GCE project to list the devices from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/devices",
	//   "response": {
	//     "$ref": "ListDevicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDevicesListCall) Pages(ctx context.Context, f func(*ListDevicesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "testing.projects.testMatrices.cancel":

type ProjectsTestMatricesCancelCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
}

// Cancel: Cancels unfinished test executions in a test matrix.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Cancel(projectId string, testMatrixId string) *ProjectsTestMatricesCancelCall {
	c := &ProjectsTestMatricesCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCancelCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesCancelCall) Context(ctx context.Context) *ProjectsTestMatricesCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsTestMatricesCancelCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.testMatrices.cancel" call.
// Exactly one of *CancelTestMatrixResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *CancelTestMatrixResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTestMatricesCancelCall) Do(opts ...googleapi.CallOption) (*CancelTestMatrixResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CancelTestMatrixResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels unfinished test executions in a test matrix.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the Test Matrix does not exist",
	//   "flatPath": "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.testMatrices.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the test.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "description": "Test matrix that will be canceled.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel",
	//   "response": {
	//     "$ref": "CancelTestMatrixResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.create":

type ProjectsTestMatricesCreateCall struct {
	s          *Service
	projectId  string
	testmatrix *TestMatrix
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Request to run a matrix of tests according to the given
// specifications.
// Unsupported environments will be returned in the state
// UNSUPPORTED.
// Matrices are limited to at most 200 supported executions.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to write to
// project
// - INVALID_ARGUMENT - if the request is malformed or if the matrix
// expands
//                      to more than 200 supported executions
func (r *ProjectsTestMatricesService) Create(projectId string, testmatrix *TestMatrix) *ProjectsTestMatricesCreateCall {
	c := &ProjectsTestMatricesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testmatrix = testmatrix
	return c
}

// RequestId sets the optional parameter "requestId": A string id used
// to detect duplicated requests.
// Ids are automatically scoped to a project, so
// users should ensure the ID is unique per-project.
// A UUID is recommended.
//
// Optional, but strongly recommended.
func (c *ProjectsTestMatricesCreateCall) RequestId(requestId string) *ProjectsTestMatricesCreateCall {
	c.urlParams_.Set("requestId", requestId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCreateCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesCreateCall) Context(ctx context.Context) *ProjectsTestMatricesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsTestMatricesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testmatrix)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.testMatrices.create" call.
// Exactly one of *TestMatrix or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TestMatrix.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTestMatricesCreateCall) Do(opts ...googleapi.CallOption) (*TestMatrix, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestMatrix{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request to run a matrix of tests according to the given specifications.\nUnsupported environments will be returned in the state UNSUPPORTED.\nMatrices are limited to at most 200 supported executions.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project\n- INVALID_ARGUMENT - if the request is malformed or if the matrix expands\n                     to more than 200 supported executions",
	//   "flatPath": "v1/projects/{projectId}/testMatrices",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.testMatrices.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCE project under which this job will run.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestId": {
	//       "description": "A string id used to detect duplicated requests.\nIds are automatically scoped to a project, so\nusers should ensure the ID is unique per-project.\nA UUID is recommended.\n\nOptional, but strongly recommended.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices",
	//   "request": {
	//     "$ref": "TestMatrix"
	//   },
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.delete":

type ProjectsTestMatricesDeleteCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
}

// Delete: Delete all record of a test matrix plus any associated test
// executions.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Delete(projectId string, testMatrixId string) *ProjectsTestMatricesDeleteCall {
	c := &ProjectsTestMatricesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesDeleteCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesDeleteCall) Context(ctx context.Context) *ProjectsTestMatricesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsTestMatricesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.testMatrices.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTestMatricesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Delete all record of a test matrix plus any associated test executions.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the Test Matrix does not exist",
	//   "flatPath": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "httpMethod": "DELETE",
	//   "id": "testing.projects.testMatrices.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the test.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "description": "Test matrix that will be canceled.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.get":

type ProjectsTestMatricesGetCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Check the status of a test matrix.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Get(projectId string, testMatrixId string) *ProjectsTestMatricesGetCall {
	c := &ProjectsTestMatricesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesGetCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTestMatricesGetCall) IfNoneMatch(entityTag string) *ProjectsTestMatricesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesGetCall) Context(ctx context.Context) *ProjectsTestMatricesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsTestMatricesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.testMatrices.get" call.
// Exactly one of *TestMatrix or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *TestMatrix.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsTestMatricesGetCall) Do(opts ...googleapi.CallOption) (*TestMatrix, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestMatrix{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Check the status of a test matrix.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the Test Matrix does not exist",
	//   "flatPath": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.testMatrices.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "testMatrixId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the test matrix.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "testMatrixId": {
	//       "description": "Unique test matrix id which was assigned by the service.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices/{testMatrixId}",
	//   "response": {
	//     "$ref": "TestMatrix"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "testing.projects.testMatrices.list":

type ProjectsTestMatricesListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List test matrices.
// The matrices are returned in the order of newest first by submit
// time.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
func (r *ProjectsTestMatricesService) List(projectId string) *ProjectsTestMatricesListCall {
	c := &ProjectsTestMatricesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesListCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTestMatricesListCall) IfNoneMatch(entityTag string) *ProjectsTestMatricesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTestMatricesListCall) Context(ctx context.Context) *ProjectsTestMatricesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsTestMatricesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.testMatrices.list" call.
// Exactly one of *ListTestMatricesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListTestMatricesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTestMatricesListCall) Do(opts ...googleapi.CallOption) (*ListTestMatricesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListTestMatricesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List test matrices.\nThe matrices are returned in the order of newest first by submit time.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed",
	//   "flatPath": "v1/projects/{projectId}/testMatrices",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.testMatrices.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Cloud project that owns the tests.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/testMatrices",
	//   "response": {
	//     "$ref": "ListTestMatricesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "testing.projects.webdriver.create":

type ProjectsWebdriverCreateCall struct {
	s          *Service
	projectId  string
	webdriver  *WebDriver
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new WebDriver environment and returns the endpoint
// for the user
// to access.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to write to
// project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the WebDriver environment or project does not exist
func (r *ProjectsWebdriverService) Create(projectId string, webdriver *WebDriver) *ProjectsWebdriverCreateCall {
	c := &ProjectsWebdriverCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.webdriver = webdriver
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverCreateCall) Fields(s ...googleapi.Field) *ProjectsWebdriverCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWebdriverCreateCall) Context(ctx context.Context) *ProjectsWebdriverCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWebdriverCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.webdriver)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.webdriver.create" call.
// Exactly one of *WebDriver or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *WebDriver.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsWebdriverCreateCall) Do(opts ...googleapi.CallOption) (*WebDriver, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WebDriver{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new WebDriver environment and returns the endpoint for the user\nto access.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the WebDriver environment or project does not exist",
	//   "flatPath": "v1/projects/{projectId}/webdriver",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.webdriver.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCP project under which to create the WebDriver environment.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/webdriver",
	//   "request": {
	//     "$ref": "WebDriver"
	//   },
	//   "response": {
	//     "$ref": "WebDriver"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.webdriver.delete":

type ProjectsWebdriverDeleteCall struct {
	s           *Service
	projectId   string
	webdriverId string
	urlParams_  gensupport.URLParams
	ctx_        context.Context
}

// Delete: Deletes a WebDriver environment instance.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsWebdriverService) Delete(projectId string, webdriverId string) *ProjectsWebdriverDeleteCall {
	c := &ProjectsWebdriverDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.webdriverId = webdriverId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverDeleteCall) Fields(s ...googleapi.Field) *ProjectsWebdriverDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWebdriverDeleteCall) Context(ctx context.Context) *ProjectsWebdriverDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWebdriverDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.webdriver.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsWebdriverDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a WebDriver environment instance.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/webdriver/{webdriverId}",
	//   "httpMethod": "DELETE",
	//   "id": "testing.projects.webdriver.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "webdriverId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCP project that contains the WebDriver endpoint to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webdriverId": {
	//       "description": "The GCE WebDriver environment to be deleted specified from the\nWebDriver id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/webdriver/{webdriverId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.webdriver.get":

type ProjectsWebdriverGetCall struct {
	s            *Service
	projectId    string
	webdriverId  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns the WebDriver environment.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the WebDriver environment or project does not exist
func (r *ProjectsWebdriverService) Get(projectId string, webdriverId string) *ProjectsWebdriverGetCall {
	c := &ProjectsWebdriverGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.webdriverId = webdriverId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverGetCall) Fields(s ...googleapi.Field) *ProjectsWebdriverGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsWebdriverGetCall) IfNoneMatch(entityTag string) *ProjectsWebdriverGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWebdriverGetCall) Context(ctx context.Context) *ProjectsWebdriverGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWebdriverGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.webdriver.get" call.
// Exactly one of *WebDriver or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *WebDriver.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsWebdriverGetCall) Do(opts ...googleapi.CallOption) (*WebDriver, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &WebDriver{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the WebDriver environment.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the WebDriver environment or project does not exist",
	//   "flatPath": "v1/projects/{projectId}/webdriver/{webdriverId}",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.webdriver.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "webdriverId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCP project that contains this WebDriver instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webdriverId": {
	//       "description": "The GCE WebDriver environment to be deleted specified from the\nWebDriver id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/webdriver/{webdriverId}",
	//   "response": {
	//     "$ref": "WebDriver"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "testing.projects.webdriver.keepalive":

type ProjectsWebdriverKeepaliveCall struct {
	s                         *Service
	projectId                 string
	webdriverId               string
	webdriverkeepaliverequest *WebDriverKeepAliveRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
}

// Keepalive: Issues a keep-alive to the WebDriver environment
// instance.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsWebdriverService) Keepalive(projectId string, webdriverId string, webdriverkeepaliverequest *WebDriverKeepAliveRequest) *ProjectsWebdriverKeepaliveCall {
	c := &ProjectsWebdriverKeepaliveCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.webdriverId = webdriverId
	c.webdriverkeepaliverequest = webdriverkeepaliverequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverKeepaliveCall) Fields(s ...googleapi.Field) *ProjectsWebdriverKeepaliveCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWebdriverKeepaliveCall) Context(ctx context.Context) *ProjectsWebdriverKeepaliveCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWebdriverKeepaliveCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.webdriverkeepaliverequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}:keepalive")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.webdriver.keepalive" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsWebdriverKeepaliveCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Issues a keep-alive to the WebDriver environment instance.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/webdriver/{webdriverId}:keepalive",
	//   "httpMethod": "POST",
	//   "id": "testing.projects.webdriver.keepalive",
	//   "parameterOrder": [
	//     "projectId",
	//     "webdriverId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The GCP project that contains the webdriver to be issued the keep-alive.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "webdriverId": {
	//       "description": "The WebDriver environment to be issued the keep-alive.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/webdriver/{webdriverId}:keepalive",
	//   "request": {
	//     "$ref": "WebDriverKeepAliveRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "testing.projects.webdriver.list":

type ProjectsWebdriverListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all the WebDriver environments.
//
// May return any of the following canonical error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsWebdriverService) List(projectId string) *ProjectsWebdriverListCall {
	c := &ProjectsWebdriverListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Used to specify the
// max number of results to be returned.
func (c *ProjectsWebdriverListCall) PageSize(pageSize int64) *ProjectsWebdriverListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Used to request a
// specific page of the results list.
func (c *ProjectsWebdriverListCall) PageToken(pageToken string) *ProjectsWebdriverListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverListCall) Fields(s ...googleapi.Field) *ProjectsWebdriverListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsWebdriverListCall) IfNoneMatch(entityTag string) *ProjectsWebdriverListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWebdriverListCall) Context(ctx context.Context) *ProjectsWebdriverListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWebdriverListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.projects.webdriver.list" call.
// Exactly one of *ListWebDriverResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListWebDriverResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsWebdriverListCall) Do(opts ...googleapi.CallOption) (*ListWebDriverResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListWebDriverResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the WebDriver environments.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the project does not exist",
	//   "flatPath": "v1/projects/{projectId}/webdriver",
	//   "httpMethod": "GET",
	//   "id": "testing.projects.webdriver.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Used to specify the max number of results to be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Used to request a specific page of the results list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The GCP project to list the environments from.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/webdriver",
	//   "response": {
	//     "$ref": "ListWebDriverResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsWebdriverListCall) Pages(ctx context.Context, f func(*ListWebDriverResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "testing.testEnvironmentCatalog.get":

type TestEnvironmentCatalogGetCall struct {
	s               *Service
	environmentType string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// Get: Get the catalog of supported test environments.
//
// May return any of the following canonical error codes:
//
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the environment type does not exist
// - INTERNAL - if an internal error occurred
func (r *TestEnvironmentCatalogService) Get(environmentType string) *TestEnvironmentCatalogGetCall {
	c := &TestEnvironmentCatalogGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.environmentType = environmentType
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TestEnvironmentCatalogGetCall) Fields(s ...googleapi.Field) *TestEnvironmentCatalogGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *TestEnvironmentCatalogGetCall) IfNoneMatch(entityTag string) *TestEnvironmentCatalogGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TestEnvironmentCatalogGetCall) Context(ctx context.Context) *TestEnvironmentCatalogGetCall {
	c.ctx_ = ctx
	return c
}

func (c *TestEnvironmentCatalogGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/testEnvironmentCatalog/{environmentType}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"environmentType": c.environmentType,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "testing.testEnvironmentCatalog.get" call.
// Exactly one of *TestEnvironmentCatalog or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *TestEnvironmentCatalog.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *TestEnvironmentCatalogGetCall) Do(opts ...googleapi.CallOption) (*TestEnvironmentCatalog, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestEnvironmentCatalog{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the catalog of supported test environments.\n\nMay return any of the following canonical error codes:\n\n- INVALID_ARGUMENT - if the request is malformed\n- NOT_FOUND - if the environment type does not exist\n- INTERNAL - if an internal error occurred",
	//   "flatPath": "v1/testEnvironmentCatalog/{environmentType}",
	//   "httpMethod": "GET",
	//   "id": "testing.testEnvironmentCatalog.get",
	//   "parameterOrder": [
	//     "environmentType"
	//   ],
	//   "parameters": {
	//     "environmentType": {
	//       "description": "The type of environment that should be listed.",
	//       "enum": [
	//         "ENVIRONMENT_TYPE_UNSPECIFIED",
	//         "ANDROID",
	//         "WEBDRIVER"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/testEnvironmentCatalog/{environmentType}",
	//   "response": {
	//     "$ref": "TestEnvironmentCatalog"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}
