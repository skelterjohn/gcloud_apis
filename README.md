# gcloud_apis

`gcloud_apis` is an auto-generated command-line interface (CLI) to interact with any Google Cloud Platform REST API.

## Dependencies

* [Go](https://golang.org/doc/install): To set up your environment correctly, read [How to Write Go Code](https://golang.org/doc/code.html).  Make sure your `GOPATH` environment variable is set.

* [Cloud SDK](https://cloud.google.com/sdk/)

## Installation & Setup

1. Download and install `gcloud_apis` with the Go tool:
	```sh
	$ go get github.com/skelterjohn/gcloud_apis
	```

2. Make sure you have an authenticated user account with `gcloud`:
	```sh
	$ gcloud auth login
	```

3. To set up auth for `gcloud_apis`, export the refresh token:

	```sh
	$ export GCLOUD_APIS_REFRESH_TOKEN=$(gcloud auth print-refresh-token)
	```

## Usage

### List methods

```sh
$ gcloud_apis list
autoscaler
autoscaler:v1beta2
bigquery
bigquery:v2
compute
compute:v1
```

```sh
$ gcloud_apis list autoscaler
autoscaler.autoscalers
autoscaler.zoneOperations
autoscaler.zones
```

```sh
$ gcloud_apis list autoscaler.autoscalers
autoscaler.autoscalers.delete
autoscaler.autoscalers.get
autoscaler.autoscalers.insert
autoscaler.autoscalers.list
autoscaler.autoscalers.patch
autoscaler.autoscalers.update
```

### Make requests

```sh
$ gcloud_apis compute.instances.list <YOUR PROJECT>/us-central1-a
{
 "id": "projects/<YOUR PROJECT>/zones/us-central1-a/instances",
 "kind": "compute#instanceList",
 "selfLink": "https://www.googleapis.com/compute/v1/projects/<YOUR PROJECT>/zones/us-central1-a/instances"
}
```

```sh
$ gcloud_apis compute.projects.setCommonInstanceMetadata <YOUR PROJECT> --items[0].key=foo --items[0].value=bar
{
  "id": "4687621082493875678",
  "insertTime": "2014-11-02T14:03:25.131-08:00",
  "kind": "compute#operation",
  "name": "operation-1418418066582-29739ad979c21-7e7dac48-b6146833",
  "operationType": "setMetadata",
  "selfLink": "https://www.googleapis.com/compute/v1/projects/<YOUR PROJECT>/global/operations/operation-1418418066582-29739ad979c21-7e7dac48-b6146833",
  "startTime": "2014-11-02T14:03:25.131-08:00",
  "status": "PENDING",
  "targetId": "14043977863849352249",
  "targetLink": "https://www.googleapis.com/compute/v1/projects/<YOUR PROJECT>",
  "user": "set
}
```

```sh
$ gcloud_apis compute.projects.get <YOUR PROJECT>
{
  "commonInstanceMetadata": {
    "fingerprint": "Vm9SEZQeWH0=",
    "kind": "compute#metadata"
  },
  "creationTimestamp": "2014-11-02T14:03:25.131-08:00",
  "id": "14043977863193752249",
  "kind": "compute#project",
  "name": "<YOUR PROJECT>",
  "quotas": [],
  "selfLink": "https://www.googleapis.com/compute/v1/projects/<YOUR PROJECT>"
}
```
