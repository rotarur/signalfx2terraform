/*
 * Org Tokens API
 *
 * API for creating, retrieving, updating, and deleting an organization's org tokens.<br>  If you're an Enterprise customer, this API also lets you set limits and thresholds for org tokens. When you do this, the org token controls the amount of data the `/datapoint` and `/backfill` endpoints can  send to SignalFx.<br> ## Overview When you send a request to store metrics using the endpoint  `https://ingest.{REALM}.signalfx.com/v2/datapoint`, you have to authenticate the request with an **org token**. In the web UI, these are called **access tokens**. <br> Although you can work with org tokens in the web UI, you can also manage  them with the org token API, which allows you to do the following:<br>   * Create a token   * Retrieve one or more tokens using search criteria   * Retrieve a token by specifying its name   * Update a token specified by its name   * Delete a token specified by its name   * Update the secret for a token specified by its name   * Set limits for requests that use a specific token.<br>    ## Authentication To create, update, or delete an org token, or to rotate the token secret, you need to authenticate with a **session token** that's associated with a SignalFx user that has administrative privileges. <br> To *retrieve* an org token, you can use an org token or session token. The session token doesn't have to be associated with a user that has administrative privileges.    ## Org tokens and limits If you're an Enterprise customer, you can set usage limits on an org token. Use this feature to apply the limits to data sources that are using that particular token. <br>  For example, use org token limits to restrict the number of MTS you send using the `/datapoint` or `/backfill` endpoints, both of which require an org token for authentication. ### Limit settings For any org token you create, you can specify the following:  <br>   * A limit or limits.    * A threshold for the limit. When the token's usage exceeds the threshold     SignalFx issues a notification.    * The notification for the threshold.<br>   If you don't specify limits, SignalFx applies the appropriate overage charges to MTS that exceed your subscription agreement. <br> If you don't specify a threshold, SignalFx doesn't send out a notification.  When you reach the limit specified for the org token, SignalFx stops recording MTS, but keeps the MTS you've already sent. <br> You can turn on the threshold feature in the web UI, but you're restricted  to a threshold of 90%. <br> If you set a threshold and a notification for it, but the notification has no recipients, SignalFx doesn't create a detector or send out the notification. ### DPM limits If your organization uses DPM pricing, you can set a DPM limit, threshold and notification for an org token. The limit restricts the number of MTS you can send with a `/datapoint` or `/backfill` request that uses the token. <br>  This feature helps you control DPM costs for servers using a token that has limits. <br>  To learn more about token limits for DPM pricing, see the topic [Track Organization DPM Usage with Access Tokens](https://docs.signalfx.com/en/latest/_sidebars-and-includes/dpm-tokens.html#dpm-tokens) in the SignalFx product documentation. ### Host-based pricing limits If your organization uses host-based pricing, you can set limits,   thresholds, and a notification for an org token. Four different token limits are available: <br>   * Hosts: Limits the number of hosts that can use this token concurrently.   * Containers: Limits the number of Docker containers that can use this     token concurrently.   * Custom metrics: Limits the number of custom metrics you send from your     hosts.   * Hi-res metrics: Limits the number of high-resolution metrics you     send from your hosts.<br>  To learn more about token limits for host-based pricing, see the topic [Managing usage limits with access  tokens](https://docs.signalfx.com/en/latest/admin-guide/tokens.html#managing-usage-limits-with-access-tokens) in the SignalFx product documentation. ### Usage-based pricing limits Limits for usage-based pricing works the same as for host-based pricing, but you can only set limits, thresholds, and notifications for custom metrics and hi-res metrics. ### Using Org token limits Org tokens help you control spending you incur from your data sources. For example, when you assign an org token to a test host and set low limits, you avoid using up your SignalFx quotas.  Another use of org token limits is to manage resource usage for different  groups of users. When you have users in the U.S. and Canada sending data to  SignalFx, you can assign them specific org tokens and manage the amount of data coming from each region.
 *
 * API version: 3.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package orgtoken

// Usage limits for one or more aspects of host-based pricing. Each limit has a key that's a numeric string, and a numeric value: <br>   * \"1\": Max number of hosts that can use this token   * \"2\": Max number of Docker containers that can use this token   * \"3\": Max number of custom metrics that can be sent with this token   * \"4\": Max number of hi-res metrics that can be sent with this token
type UsageLimits struct {
	HostThreshold          *int64 `json:"1,omitempty"`
	ContainerThreshold     *int64 `json:"2,omitempty"`
	CustomMetricThreshold  *int64 `json:"3,omitempty"`
	HighResMetricThreshold *int64 `json:"4,omitempty"`
}