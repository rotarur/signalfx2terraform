# 1.6.2, Unreleased

## Added

## Updated

## Bugfixes

## Removed

# 1.6.1, 2019-08-16

## Updated

Adjusted detector.CreateUpdateDetectorRequest to use pointer for Rules

# 1.6.0, 2019-08-16

## Added

* Added `*GCPIntegration` methods
* Added `*Opsgenie` methods
* Added `*PagerDutyIntegration` methods
* Added `*SlackIntegration` methods

## Updated

* `Detector.Rules` now uses `Notification` as it's type instead of an untyped `[]map[string]interface{}`.
* Renamed `integration.GcpIntegration` and it's sub-types to `GCP`, fixing case.

# 1.5.0, 2019-08-05

## Added

* Add OrgToken methods

## Bugfixes

* Properly recognize the SignalFlow keep alive event message and ignore it.

## Updated

* Moved various notification bits into a `notification` package

# 1.4.0, 2019-07-29

## Added

* Add `*AzureIntregration` functions to client.

## Updated

## Bugfixes

## Removed

# 1.3.0, 2019-07-24

## Added

* Added OpenAPI code for integrations, experimental for now.
* Add `*AwsCloudWatchIntegration` functions to client.

## Removed

* Removed `credentialName` from Opsgenie notifications, not a real field in the API.

# 1.2.0, 2019-07-16

## Updated

* Many numeric properties have been adjusted to pointers to play better with Go's JSON (un)marshaling.

# 1.1.0, 2019-07-15

## Added
* Added `DashboardConfigs` to `CreateUpdateDashboardRequest`
* `DashboardGroupCreate` now has an option to create an empty group.

## Updated
* Many types have been changed to pointers to add (de)serialization
* Moved `StringOrSlice` into a `util` package, cuz all projects must have one

## Bugfixes
* Switched to `StringOrSlice` for some fields that needed it.
* Added `StringOrInteger` to handle failures in some Chart filter responses, thanks to (doctornkz)[https://github.com/doctornkz] for flagging!

## Removed

# 1.0.0, Forgot the date

Tagged!

## Added

## Updated

## Bugfixes

## Removed