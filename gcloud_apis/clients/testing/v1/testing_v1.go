// Package testing provides access to the Google Cloud Testing API.
//
// See https://developers.google.com/cloud-test-lab/
//
// Usage example:
//
//   import "google.golang.org/api/testing/v1"
//   ...
//   testingService, err := testing.New(oauthHttpClient)
package testing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/googleapi"
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
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

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
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService

	TestEnvironmentCatalog *TestEnvironmentCatalogService
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

type AndroidDevice struct {
	// AndroidModelId: The id of the Android device to be used.
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	// Required
	AndroidModelId string `json:"androidModelId,omitempty"`

	// AndroidVersionId: The id of the Android OS version to be used.
	// Use
	// the EnvironmentDiscoveryService to get supported options.
	// Required
	AndroidVersionId string `json:"androidVersionId,omitempty"`

	// Locale: The locale the test device used for testing.
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	// Required
	Locale string `json:"locale,omitempty"`

	// Orientation: How the device is oriented during the test.
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	// Required
	Orientation string `json:"orientation,omitempty"`
}

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
}

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
	// Optional, default is determined by examining the
	// application's manifest.
	TestPackageId string `json:"testPackageId,omitempty"`

	// TestRunnerClass: The InstrumentationTestRunner class.
	// Optional,
	// default is determined by examining the application's manifest.
	TestRunnerClass string `json:"testRunnerClass,omitempty"`

	// TestTargets: Each target must be fully qualified with the package
	// name or class name,
	// in one of these formats:
	//  - "package
	// package_name"
	//  - "class package_name.class_name"
	//  - "class
	// package_name.class_name#method_name"
	//
	// If empty, all targets in the
	// module will be run.
	TestTargets []string `json:"testTargets,omitempty"`
}

type AndroidMatrix struct {
	// AndroidModelIds: The ids of the set of Android device to be used.
	// Use
	// the EnvironmentDiscoveryService to get supported options.
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
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	// Required
	Orientations []string `json:"orientations,omitempty"`
}

type AndroidModel struct {
	// Brand: The company that this device is branded with.
	// Example:
	// "Google", "Samsung"
	// @OutputOnly
	Brand string `json:"brand,omitempty"`

	// Codename: The name of the industrial design.
	// This corresponds to
	// android.os.Build.DEVICE
	// @OutputOnly
	Codename string `json:"codename,omitempty"`

	// Form: Whether this device is virtual or physical.
	// @OutputOnly
	Form string `json:"form,omitempty"`

	// Id: The unique opaque id for this model.
	// Use this for invoking the
	// TestExecutionService.
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
	// This
	// corresponds to either android.os.Build.SUPPORTED_ABIS (for API
	// level
	// 21 and above) or android.os.Build.CPU_ABI/CPU_ABI2.
	// @OutputOnly
	SupportedAbis []string `json:"supportedAbis,omitempty"`

	// SupportedVersionIds: The set of Android versions this device
	// supports.
	// Note that not all of these are necessarily supported in
	// physical devices.
	// @OutputOnly
	SupportedVersionIds []string `json:"supportedVersionIds,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview",
	// "deprecated"
	Tags []string `json:"tags,omitempty"`
}

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
	// Note that,
	// although specifying a seed causes the monkey to generate the
	// same
	// sequence of events, it does not guarantee that a particular
	// outcome
	// will be reproducible across runs.
	// Optional
	RandomSeed int64 `json:"randomSeed,omitempty"`
}

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

	// BootstrapApk: The APK used for bootstrapping (e.g., passing the login
	// screen).
	// Optional
	BootstrapApk *FileReference `json:"bootstrapApk,omitempty"`

	// BootstrapPackageId: The java package for the bootstrap.
	// Optional
	BootstrapPackageId string `json:"bootstrapPackageId,omitempty"`

	// BootstrapRunnerClass: The runner class for the bootstrap.
	// Optional
	BootstrapRunnerClass string `json:"bootstrapRunnerClass,omitempty"`

	// MaxDepth: The max depth of the traversal stack Robo can explore.
	// Needs to be at least
	// 2 to make Robo explore the app beyond the first
	// activity.
	// Default is 50.
	// Optional
	MaxDepth int64 `json:"maxDepth,omitempty"`

	// MaxSteps: The max number of steps Robo can execute.
	// Default is no
	// limit.
	// Optional
	MaxSteps int64 `json:"maxSteps,omitempty"`

	// RandomizeSteps: Whether Robo follows a random order of steps on a
	// given activity state.
	// Optional
	RandomizeSteps bool `json:"randomizeSteps,omitempty"`
}

type AndroidRuntimeConfiguration struct {
	// Locales: The set of available locales.
	// @OutputOnly
	Locales []*Locale `json:"locales,omitempty"`

	// Orientations: The set of available orientations.
	// @OutputOnly
	Orientations []*Orientation `json:"orientations,omitempty"`
}

type AndroidVersion struct {
	// ApiLevel: The API level for this Android version.
	// Examples: 18,
	// 19
	// @OutputOnly
	ApiLevel int64 `json:"apiLevel,omitempty"`

	// CodeName: The code name for this Android version.
	// Examples:
	// "JellyBean", "KitKat"
	// @OutputOnly
	CodeName string `json:"codeName,omitempty"`

	// Distribution: Market share for this version.
	// @OutputOnly
	Distribution *Distribution `json:"distribution,omitempty"`

	// Id: An opaque id for this Android version.
	// Use this id to invoke the
	// TestExecutionService.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// ReleaseDate: The date this Android version became available in the
	// market.
	// @OutputOnly
	ReleaseDate *Date `json:"releaseDate,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default", "preview",
	// "deprecated"
	Tags []string `json:"tags,omitempty"`

	// VersionString: A string representing this version of the Android
	// OS.
	// Examples: "4.3", "4.4"
	// @OutputOnly
	VersionString string `json:"versionString,omitempty"`
}

type BlobstoreFile struct {
	// BlobId: A blob ID.
	// Example:
	// /android_test/blobs/4e9AAT9sqHRY_oBBzIKHSEFgg
	BlobId string `json:"blobId,omitempty"`

	// Md5Hash: The MD5 hash of the referenced blob. (This is necessary to
	// create a
	// Bigstore object directly from the Blobstore reference.)
	Md5Hash string `json:"md5Hash,omitempty"`
}

type Browser struct {
	// AndroidCatalog: The catalog of Android devices for which we offer
	// this browser.
	// @OutputOnly
	AndroidCatalog *AndroidDeviceCatalog `json:"androidCatalog,omitempty"`

	// Id: A human readable id for this Browser version.
	// Use this id to
	// invoke the TestExecutionService.
	// Examples: "chrome-stable-channel",
	// "firefox-beta-channel"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// LinuxCatalog: The catalog of Linux machines which we offer this
	// browser.
	// @OutputOnly
	LinuxCatalog *LinuxMachineCatalog `json:"linuxCatalog,omitempty"`

	// Name: A string representing the browser name.
	// Examples: "chrome",
	// "firefox", "ie"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Release: The release of the browser.
	// Examples: "stable-channel",
	// "beta-channel", "10" (for ie), etc
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
}

type CancelTestMatrixResponse struct {
	// TestState: The rolled-up state of the test matrix just before it was
	// cancelled.
	TestState string `json:"testState,omitempty"`
}

type ClientInfo struct {
	// Name: Client name, such as gcloud.
	Name string `json:"name,omitempty"`
}

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
}

type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year/month where the day is not
	// significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`
}

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
	State string `json:"state,omitempty"`

	// StateDetails: Details about the state of the device.
	// NOT
	// user-specified
	StateDetails *DeviceStateDetails `json:"stateDetails,omitempty"`
}

type DeviceDetails struct {
	// ConnectionInfo: Details about the connection to the device.
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`

	// GceInstanceDetails: Details about the GCE instance backing the
	// device.
	GceInstanceDetails *GceInstanceDetails `json:"gceInstanceDetails,omitempty"`
}

type DeviceFile struct {
	ObbFile *ObbFile `json:"obbFile,omitempty"`
}

type DeviceStateDetails struct {
	// ErrorDetails: If the DeviceState is ERROR, then this string may
	// contain human-readable
	// details about the error.
	ErrorDetails string `json:"errorDetails,omitempty"`

	// ProgressDetails: A human-readable, detailed description of the
	// device's status.
	// For example: "Starting Device\n Device
	// Ready".
	//
	// During the device's lifespan data may be appended to the
	// progress.
	ProgressDetails string `json:"progressDetails,omitempty"`
}

type Distribution struct {
	// MarketShare: The estimated fraction (0-1) of the total market with
	// this configuration.
	// @OutputOnly
	MarketShare float64 `json:"marketShare,omitempty"`

	// MeasurementTime: The time this distribution was measured.
	// @OutputOnly
	MeasurementTime string `json:"measurementTime,omitempty"`
}

type Empty struct {
}

type Environment struct {
	// AndroidDevice: An Android device which must be used with an Android
	// test.
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`
}

type EnvironmentMatrix struct {
	// AndroidMatrix: A matrix of Android devices
	AndroidMatrix *AndroidMatrix `json:"androidMatrix,omitempty"`
}

type FileReference struct {
	// Blob: A blob in Blobstore.
	Blob *BlobstoreFile `json:"blob,omitempty"`

	// GcsPath: A path to a file in Google Cloud Storage.
	// Example:
	// gs://build-app-1414623860166/app-debug-unaligned.apk
	GcsPath string `json:"gcsPath,omitempty"`
}

type GceInstanceDetails struct {
	// Name: Desired instance name of the device.
	// May be user-specified. If
	// not, the backend will choose a name.
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
}

type GoogleCloudStorage struct {
	// GcsPath: The path to a directory in GCS that will
	// eventually contain
	// the results for this test.
	// The requesting user must have write access
	// on the bucket in the supplied
	// path.
	GcsPath string `json:"gcsPath,omitempty"`
}

type LinuxMachine struct {
	// VersionId: The version id of the Linux OS to be used.
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	VersionId string `json:"versionId,omitempty"`
}

type LinuxMachineCatalog struct {
	// Versions: The set of supported Linux versions.
	// @OutputOnly
	Versions []*LinuxVersion `json:"versions,omitempty"`
}

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
}

type ListDevicesResponse struct {
	// Devices: The GCE virtual Android devices to be returned.
	Devices []*Device `json:"devices,omitempty"`

	// NextPageToken: The pagination token to retrieve the next page of
	// device results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListTestMatricesResponse struct {
	// TestMatrices: The set of test matrices.
	TestMatrices []*TestMatrix `json:"testMatrices,omitempty"`
}

type ListWebDriverResponse struct {
	// NextPageToken: The pagination token to retrieve the next page of
	// WebDriver results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// WebdriverEnvironments: The WebDriver environments to be returned.
	WebdriverEnvironments []*WebDriver `json:"webdriverEnvironments,omitempty"`
}

type Locale struct {
	// Id: The id for this locale.
	// Example: "en_US"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this language/locale.
	// Example:
	// "English"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Region: A human-friendy string representing the region for this
	// locale.
	// Example: "United States"
	// Not present for every
	// locale.
	// @OutputOnly
	Region string `json:"region,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`
}

type ObbFile struct {
	// Obb: Opaque Binary Blob (OBB) file(s) to install on the device
	Obb *FileReference `json:"obb,omitempty"`

	// ObbFileName: OBB file name which must conform to the format as
	// specified by
	// Android
	// e.g.
	// [main|patch].0300110.com.example.android.obb
	// which will be installed
	// into
	//   <shared-storage>/Android/obb/<package-name>/
	// on the device
	ObbFileName string `json:"obbFileName,omitempty"`
}

type Orientation struct {
	// Id: The id for this orientation.
	// Example: "portrait"
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Name: A human-friendly name for this orientation.
	// Example:
	// "portrait"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// Tags: Tags for this dimension.
	// Examples: "default"
	Tags []string `json:"tags,omitempty"`
}

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
	// If not provided the
	// service will choose an appropriate value.
	ToolResultsHistory *ToolResultsHistory `json:"toolResultsHistory,omitempty"`
}

type TestDetails struct {
	// ErrorMessage: If the TestState is ERROR, then this string will
	// contain human-readable
	// details about the error.
	// @OutputOnly
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ProgressMessages: Human-readable, detailed descriptions of the test's
	// progress.
	// For example: "Provisioning a device", "Starting
	// Test".
	//
	// During the course of execution new data may be appended
	// to
	// the end of progress_messages.
	// @OutputOnly
	ProgressMessages []string `json:"progressMessages,omitempty"`
}

type TestEnvironmentCatalog struct {
	// AndroidDeviceCatalog: Android devices suitable for running Android
	// Instrumentation Tests.
	AndroidDeviceCatalog *AndroidDeviceCatalog `json:"androidDeviceCatalog,omitempty"`

	// WebDriverCatalog: WebDriver environments suitable for running web
	// tests.
	WebDriverCatalog *WebDriverCatalog `json:"webDriverCatalog,omitempty"`
}

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
}

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
}

type TestSetup struct {
	FilesToPush []*DeviceFile `json:"filesToPush,omitempty"`
}

type TestSpecification struct {
	// AndroidInstrumentationTest: An Android instrumentation test.
	AndroidInstrumentationTest *AndroidInstrumentationTest `json:"androidInstrumentationTest,omitempty"`

	// AndroidMonkeyTest: An Android monkey test.
	AndroidMonkeyTest *AndroidMonkeyTest `json:"androidMonkeyTest,omitempty"`

	// AndroidRoboTest: An Android robo test.
	AndroidRoboTest *AndroidRoboTest `json:"androidRoboTest,omitempty"`

	// AutoGoogleLogin: Enables automatic Google account login.
	// If set, the
	// service will automatically generate a Google test account
	// and use it
	// to log into the device, before executing the test. Note that
	// test
	// accounts might be reused.
	// Many applications can be tested more
	// effectively in the context of
	// such an account.
	// Default is
	// false.
	// Optional
	AutoGoogleLogin bool `json:"autoGoogleLogin,omitempty"`

	// TestSetup: Test setup requirements e.g. files to install, bootstrap
	// scripts
	TestSetup *TestSetup `json:"testSetup,omitempty"`

	// TestTimeout: Max time a test execution is allowed to run before it
	// is
	// automatically cancelled.
	TestTimeout string `json:"testTimeout,omitempty"`
}

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
}

type ToolResultsHistory struct {
	// HistoryId: A tool results history ID.
	HistoryId string `json:"historyId,omitempty"`

	// ProjectId: The cloud project that owns the tool results history.
	ProjectId string `json:"projectId,omitempty"`
}

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
	State string `json:"state,omitempty"`

	// StateDetails: Details about the state of the device.
	// @OutputOnly
	StateDetails *DeviceStateDetails `json:"stateDetails,omitempty"`
}

type WebDriver struct {
	// AndroidDevice: An Android device.
	AndroidDevice *AndroidDevice `json:"androidDevice,omitempty"`

	// BrowserId: The id of the browser to be used.
	// Use the
	// EnvironmentDiscoveryService to get supported values.
	// Required
	BrowserId string `json:"browserId,omitempty"`

	// Endpoint: The endpoint in host:port format where the target running
	// the specified
	// browser accepts WebDriver protocol
	// commands.
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
}

type WebDriverCatalog struct {
	// Browsers: The set of supported browsers.
	// @OutputOnly
	Browsers []*Browser `json:"browsers,omitempty"`
}

type WebDriverKeepAliveRequest struct {
}

type WindowsMachine struct {
	// VersionId: The version id of the Windows OS to be used.
	// Use the
	// EnvironmentDiscoveryService to get supported options.
	VersionId string `json:"versionId,omitempty"`
}

type WindowsMachineCatalog struct {
	// Versions: The set of supported Windows versions.
	// @OutputOnly
	Versions []*WindowsVersion `json:"versions,omitempty"`
}

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
}

// method id "testing.projects.devices.create":

type ProjectsDevicesCreateCall struct {
	s         *Service
	projectId string
	device    *Device
	opt_      map[string]interface{}
}

// Create: Creates a new GCE Android device.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to write to project
// - INVALID_ARGUMENT - if the
// request is malformed
// - NOT_FOUND - if the device type or project does
// not exist
func (r *ProjectsDevicesService) Create(projectId string, device *Device) *ProjectsDevicesCreateCall {
	c := &ProjectsDevicesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.device = device
	return c
}

// SshPublicKey sets the optional parameter "sshPublicKey": The public
// key to be set on the device in order to SSH into it.
func (c *ProjectsDevicesCreateCall) SshPublicKey(sshPublicKey string) *ProjectsDevicesCreateCall {
	c.opt_["sshPublicKey"] = sshPublicKey
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesCreateCall) Fields(s ...googleapi.Field) *ProjectsDevicesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesCreateCall) Do() (*Device, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.device)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["sshPublicKey"]; ok {
		params.Set("sshPublicKey", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Device
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
	s         *Service
	projectId string
	deviceId  string
	opt_      map[string]interface{}
}

// Delete: Deletes a GCE Android device instance.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to read project
// - INVALID_ARGUMENT - if the request
// is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) Delete(projectId string, deviceId string) *ProjectsDevicesDeleteCall {
	c := &ProjectsDevicesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesDeleteCall) Fields(s ...googleapi.Field) *ProjectsDevicesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
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
	s         *Service
	projectId string
	deviceId  string
	opt_      map[string]interface{}
}

// Get: Returns the GCE Android device.
//
// May return any of the following
// canonical error codes:
//
// - PERMISSION_DENIED - if the user is not
// authorized to read project
// - INVALID_ARGUMENT - if the request is
// malformed
// - NOT_FOUND - if the device type or project does not exist
func (r *ProjectsDevicesService) Get(projectId string, deviceId string) *ProjectsDevicesGetCall {
	c := &ProjectsDevicesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesGetCall) Fields(s ...googleapi.Field) *ProjectsDevicesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesGetCall) Do() (*Device, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Device
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
	s         *Service
	projectId string
	deviceId  string
	opt_      map[string]interface{}
}

// Keepalive: Issues a keep-alive to a GCE Android device instance.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to read project
// -
// INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the
// project does not exist
func (r *ProjectsDevicesService) Keepalive(projectId string, deviceId string) *ProjectsDevicesKeepaliveCall {
	c := &ProjectsDevicesKeepaliveCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.deviceId = deviceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesKeepaliveCall) Fields(s ...googleapi.Field) *ProjectsDevicesKeepaliveCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesKeepaliveCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices/{deviceId}/keepalive")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"deviceId":  c.deviceId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
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
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all the current devices.
//
// May return any of the following
// canonical error codes:
//
// - PERMISSION_DENIED - if the user is not
// authorized to read project
// - INVALID_ARGUMENT - if the request is
// malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsDevicesService) List(projectId string) *ProjectsDevicesListCall {
	c := &ProjectsDevicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Used to specify the
// max number of device results to be returned.
func (c *ProjectsDevicesListCall) PageSize(pageSize int64) *ProjectsDevicesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Used to request a
// specific page of the device results list.
func (c *ProjectsDevicesListCall) PageToken(pageToken string) *ProjectsDevicesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDevicesListCall) Fields(s ...googleapi.Field) *ProjectsDevicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDevicesListCall) Do() (*ListDevicesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/devices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListDevicesResponse
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

// method id "testing.projects.testMatrices.cancel":

type ProjectsTestMatricesCancelCall struct {
	s            *Service
	projectId    string
	testMatrixId string
	opt_         map[string]interface{}
}

// Cancel: Cancels unfinished test executions in a test matrix.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to read project
// -
// INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the
// Test Matrix does not exist
func (r *ProjectsTestMatricesService) Cancel(projectId string, testMatrixId string) *ProjectsTestMatricesCancelCall {
	c := &ProjectsTestMatricesCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCancelCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesCancelCall) Do() (*CancelTestMatrixResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CancelTestMatrixResponse
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
	opt_       map[string]interface{}
}

// Create: Request to run a matrix of tests according to the given
// specifications.
// Unsupported environments will be returned in the
// state UNSUPPORTED.
// Matrices are limited to at most 200 supported
// executions.
//
// May return any of the following canonical error
// codes:
//
// - PERMISSION_DENIED - if the user is not authorized to write
// to project
// - INVALID_ARGUMENT - if the request is malformed or if the
// matrix expands
//                      to more than 200 supported
// executions
func (r *ProjectsTestMatricesService) Create(projectId string, testmatrix *TestMatrix) *ProjectsTestMatricesCreateCall {
	c := &ProjectsTestMatricesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testmatrix = testmatrix
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesCreateCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesCreateCall) Do() (*TestMatrix, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testmatrix)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TestMatrix
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
	opt_         map[string]interface{}
}

// Delete: Delete all record of a test matrix plus any associated test
// executions.
//
// May return any of the following canonical error
// codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read
// project
// - INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND
// - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Delete(projectId string, testMatrixId string) *ProjectsTestMatricesDeleteCall {
	c := &ProjectsTestMatricesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesDeleteCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
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
	opt_         map[string]interface{}
}

// Get: Check the status of a test matrix.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to read project
// - INVALID_ARGUMENT - if the request
// is malformed
// - NOT_FOUND - if the Test Matrix does not exist
func (r *ProjectsTestMatricesService) Get(projectId string, testMatrixId string) *ProjectsTestMatricesGetCall {
	c := &ProjectsTestMatricesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.testMatrixId = testMatrixId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesGetCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesGetCall) Do() (*TestMatrix, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices/{testMatrixId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":    c.projectId,
		"testMatrixId": c.testMatrixId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TestMatrix
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
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List test matrices.
// The matrices are returned in the order of
// newest first by submit time.
//
// May return any of the following
// canonical error codes:
//
// - PERMISSION_DENIED - if the user is not
// authorized to read project
// - INVALID_ARGUMENT - if the request is
// malformed
func (r *ProjectsTestMatricesService) List(projectId string) *ProjectsTestMatricesListCall {
	c := &ProjectsTestMatricesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestMatricesListCall) Fields(s ...googleapi.Field) *ProjectsTestMatricesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestMatricesListCall) Do() (*ListTestMatricesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/testMatrices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListTestMatricesResponse
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
	s         *Service
	projectId string
	webdriver *WebDriver
	opt_      map[string]interface{}
}

// Create: Creates a new WebDriver environment and returns the endpoint
// for the user
// to access.
//
// May return any of the following canonical
// error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to
// write to project
// - INVALID_ARGUMENT - if the request is malformed
// -
// NOT_FOUND - if the WebDriver environment or project does not exist
func (r *ProjectsWebdriverService) Create(projectId string, webdriver *WebDriver) *ProjectsWebdriverCreateCall {
	c := &ProjectsWebdriverCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.webdriver = webdriver
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverCreateCall) Fields(s ...googleapi.Field) *ProjectsWebdriverCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWebdriverCreateCall) Do() (*WebDriver, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.webdriver)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *WebDriver
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
	opt_        map[string]interface{}
}

// Delete: Deletes a WebDriver environment instance.
//
// May return any of
// the following canonical error codes:
//
// - PERMISSION_DENIED - if the
// user is not authorized to read project
// - INVALID_ARGUMENT - if the
// request is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsWebdriverService) Delete(projectId string, webdriverId string) *ProjectsWebdriverDeleteCall {
	c := &ProjectsWebdriverDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.webdriverId = webdriverId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverDeleteCall) Fields(s ...googleapi.Field) *ProjectsWebdriverDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWebdriverDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
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
	s           *Service
	projectId   string
	webdriverId string
	opt_        map[string]interface{}
}

// Get: Returns the WebDriver environment.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to read project
// - INVALID_ARGUMENT - if the request
// is malformed
// - NOT_FOUND - if the WebDriver environment or project
// does not exist
func (r *ProjectsWebdriverService) Get(projectId string, webdriverId string) *ProjectsWebdriverGetCall {
	c := &ProjectsWebdriverGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.webdriverId = webdriverId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverGetCall) Fields(s ...googleapi.Field) *ProjectsWebdriverGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWebdriverGetCall) Do() (*WebDriver, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *WebDriver
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
	opt_                      map[string]interface{}
}

// Keepalive: Issues a keep-alive to the WebDriver environment
// instance.
//
// May return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to read project
// -
// INVALID_ARGUMENT - if the request is malformed
// - NOT_FOUND - if the
// project does not exist
func (r *ProjectsWebdriverService) Keepalive(projectId string, webdriverId string, webdriverkeepaliverequest *WebDriverKeepAliveRequest) *ProjectsWebdriverKeepaliveCall {
	c := &ProjectsWebdriverKeepaliveCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.webdriverId = webdriverId
	c.webdriverkeepaliverequest = webdriverkeepaliverequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverKeepaliveCall) Fields(s ...googleapi.Field) *ProjectsWebdriverKeepaliveCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWebdriverKeepaliveCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.webdriverkeepaliverequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver/{webdriverId}:keepalive")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"webdriverId": c.webdriverId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
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
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all the WebDriver environments.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to read project
// - INVALID_ARGUMENT - if the request
// is malformed
// - NOT_FOUND - if the project does not exist
func (r *ProjectsWebdriverService) List(projectId string) *ProjectsWebdriverListCall {
	c := &ProjectsWebdriverListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Used to specify the
// max number of results to be returned.
func (c *ProjectsWebdriverListCall) PageSize(pageSize int64) *ProjectsWebdriverListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Used to request a
// specific page of the results list.
func (c *ProjectsWebdriverListCall) PageToken(pageToken string) *ProjectsWebdriverListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWebdriverListCall) Fields(s ...googleapi.Field) *ProjectsWebdriverListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWebdriverListCall) Do() (*ListWebDriverResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/webdriver")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListWebDriverResponse
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

// method id "testing.testEnvironmentCatalog.get":

type TestEnvironmentCatalogGetCall struct {
	s               *Service
	environmentType string
	opt_            map[string]interface{}
}

// Get: Get the catalog of supported test environments.
//
// May return any
// of the following canonical error codes:
//
// - INVALID_ARGUMENT - if the
// request is malformed
// - NOT_FOUND - if the environment type does not
// exist
// - INTERNAL - if an internal error occurred
func (r *TestEnvironmentCatalogService) Get(environmentType string) *TestEnvironmentCatalogGetCall {
	c := &TestEnvironmentCatalogGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.environmentType = environmentType
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TestEnvironmentCatalogGetCall) Fields(s ...googleapi.Field) *TestEnvironmentCatalogGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TestEnvironmentCatalogGetCall) Do() (*TestEnvironmentCatalog, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/testEnvironmentCatalog/{environmentType}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"environmentType": c.environmentType,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TestEnvironmentCatalog
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
