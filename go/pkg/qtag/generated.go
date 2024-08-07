package qtag

import (
	"encoding/json"
)

type Dbt struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewDbt() *Dbt {
	x := Dbt{}
	x.init()
	return &x
}

func (c *Dbt) init() {
	c.name = "dbt"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"fields\":{\"app\":\"dbt\"}}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Dbt) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Dbt) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Dbt) Merge(other *Dbt) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Dbt) AddApp(value any) {
	c.values["app"] = value
}

func (c *Dbt) AddDbt_version(value any) {
	c.values["dbt_version"] = value
}

func (c *Dbt) AddConnection_name(value any) {
	c.values["connection_name"] = value
}

func (c *Dbt) AddProfile_name(value any) {
	c.values["profile_name"] = value
}

func (c *Dbt) AddTarget_name(value any) {
	c.values["target_name"] = value
}

func (c *Dbt) AddNode_id(value any) {
	c.values["node_id"] = value
}

func (c *Dbt) AddDbt_snowflake_query_tags_version(value any) {
	c.values["dbt_snowflake_query_tags_version"] = value
}

func (c *Dbt) AddProject_name(value any) {
	c.values["project_name"] = value
}

func (c *Dbt) AddTarget_database(value any) {
	c.values["target_database"] = value
}

func (c *Dbt) AddTarget_schema(value any) {
	c.values["target_schema"] = value
}

func (c *Dbt) AddInvocation_id(value any) {
	c.values["invocation_id"] = value
}

func (c *Dbt) AddNode_name(value any) {
	c.values["node_name"] = value
}

func (c *Dbt) AddNode_alias(value any) {
	c.values["node_alias"] = value
}

func (c *Dbt) AddNode_package_name(value any) {
	c.values["node_package_name"] = value
}

func (c *Dbt) AddNode_original_file_path(value any) {
	c.values["node_original_file_path"] = value
}

func (c *Dbt) AddNode_database(value any) {
	c.values["node_database"] = value
}

func (c *Dbt) AddNode_schema(value any) {
	c.values["node_schema"] = value
}

func (c *Dbt) AddNode_resource_type(value any) {
	c.values["node_resource_type"] = value
}

func (c *Dbt) AddMaterialized(value any) {
	c.values["materialized"] = value
}

func (c *Dbt) AddIs_incremental(value any) {
	c.values["is_incremental"] = value
}

func (c *Dbt) AddNode_refs(value any) {
	c.values["node_refs"] = value
}

func (c *Dbt) AddDbt_cloud_project_id(value any) {
	c.values["dbt_cloud_project_id"] = value
}

func (c *Dbt) AddDbt_cloud_job_id(value any) {
	c.values["dbt_cloud_job_id"] = value
}

func (c *Dbt) AddDbt_cloud_run_id(value any) {
	c.values["dbt_cloud_run_id"] = value
}

func (c *Dbt) AddDbt_cloud_run_reason(value any) {
	c.values["dbt_cloud_run_reason"] = value
}

func (c *Dbt) AddDbt_cloud_run_reason_category(value any) {
	c.values["dbt_cloud_run_reason_category"] = value
}

type Hex struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewHex() *Hex {
	x := Hex{}
	x.init()
	return &x
}

func (c *Hex) init() {
	c.name = "hex"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"prefix\":\"Hex query metadata:\"}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Hex) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Hex) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Hex) Merge(other *Hex) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Hex) AddCategories(value any) {
	c.values["categories"] = value
}

func (c *Hex) AddCell_type(value any) {
	c.values["cell_type"] = value
}

func (c *Hex) AddConnection(value any) {
	c.values["connection"] = value
}

func (c *Hex) AddContext(value any) {
	c.values["context"] = value
}

func (c *Hex) AddProject_name(value any) {
	c.values["project_name"] = value
}

func (c *Hex) AddProject_url(value any) {
	c.values["project_url"] = value
}

func (c *Hex) AddUser_email(value any) {
	c.values["user_email"] = value
}

type Metabase struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewMetabase() *Metabase {
	x := Metabase{}
	x.init()
	return &x
}

func (c *Metabase) init() {
	c.name = "metabase"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"fields\":{\"client\":\"Metabase\"}}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Metabase) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Metabase) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Metabase) Merge(other *Metabase) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Metabase) AddClient(value any) {
	c.values["client"] = value
}

func (c *Metabase) AddServerId(value any) {
	c.values["serverId"] = value
}

func (c *Metabase) AddQueryHash(value any) {
	c.values["queryHash"] = value
}

func (c *Metabase) AddQueryType(value any) {
	c.values["queryType"] = value
}

func (c *Metabase) AddCardId(value any) {
	c.values["cardId"] = value
}

func (c *Metabase) AddDashboardId(value any) {
	c.values["dashboardId"] = value
}

func (c *Metabase) AddContext(value any) {
	c.values["context"] = value
}

func (c *Metabase) AddUserId(value any) {
	c.values["userId"] = value
}

func (c *Metabase) AddDatabaseId(value any) {
	c.values["databaseId"] = value
}

func (c *Metabase) AddPulseId(value any) {
	c.values["pulseId"] = value
}

type Mode struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewMode() *Mode {
	x := Mode{}
	x.init()
	return &x
}

func (c *Mode) init() {
	c.name = "mode"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"bodymatch\":\"https://modeanalytics.com/\"}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Mode) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Mode) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Mode) Merge(other *Mode) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Mode) AddUser(value any) {
	c.values["user"] = value
}

func (c *Mode) AddEmail(value any) {
	c.values["email"] = value
}

func (c *Mode) AddUrl(value any) {
	c.values["url"] = value
}

func (c *Mode) AddScheduled(value any) {
	c.values["scheduled"] = value
}

type Sigma struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewSigma() *Sigma {
	x := Sigma{}
	x.init()
	return &x
}

func (c *Sigma) init() {
	c.name = "sigma"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"prefix\":\"Sigma Σ\"}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Sigma) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Sigma) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Sigma) Merge(other *Sigma) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Sigma) AddKind(value any) {
	c.values["kind"] = value
}

func (c *Sigma) AddRequest__id(value any) {
	c.values["request__id"] = value
}

func (c *Sigma) AddSourceUrl(value any) {
	c.values["sourceUrl"] = value
}

func (c *Sigma) AddEmail(value any) {
	c.values["email"] = value
}

type Sundeck struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewSundeck() *Sundeck {
	x := Sundeck{}
	x.init()
	return &x
}

func (c *Sundeck) init() {
	c.name = "sundeck"
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"fields\":{\"app\":\"sundeck\"}}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *Sundeck) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *Sundeck) UnknownValue(name string, value any) {
	c.values[name] = value
}

func (c *Sundeck) Merge(other *Sundeck) {
	for k, v := range other.values {
		c.values[k] = v
	}
}

func (c *Sundeck) AddApp(value any) {
	c.values["app"] = value
}

func (c *Sundeck) AddWarehouse_updated(value any) {
	c.values["warehouse_updated"] = value
}

func (c *Sundeck) AddQuery_text_updated(value any) {
	c.values["query_text_updated"] = value
}

func (c *Sundeck) AddFlow_name(value any) {
	c.values["flow_name"] = value
}

func (c *Sundeck) AddQuery_rejected(value any) {
	c.values["query_rejected"] = value
}

func (c *Sundeck) AddAccount(value any) {
	c.values["account"] = value
}

func (c *Sundeck) AddUser(value any) {
	c.values["user"] = value
}

func (c *Sundeck) AddService_name(value any) {
	c.values["service_name"] = value
}

func (c *Sundeck) AddService_operation(value any) {
	c.values["service_operation"] = value
}

func (c *Sundeck) AddHooks_applied(value any) {
	c.values["hooks_applied"] = value
}

func (c *Sundeck) AddReject_message(value any) {
	c.values["reject_message"] = value
}

func (c *Sundeck) AddParent_query_id(value any) {
	c.values["parent_query_id"] = value
}

func (c *Sundeck) AddChild_query_ids(value any) {
	c.values["child_query_ids"] = value
}

func (c *Sundeck) AddWorkload(value any) {
	c.values["workload"] = value
}

func (c *Sundeck) AddKind(value any) {
	c.values["kind"] = value
}

func (c *Sundeck) AddAuto_routing_matched(value any) {
	c.values["auto_routing_matched"] = value
}

func (c *Sundeck) AddAuto_routing_matched_warehouse(value any) {
	c.values["auto_routing_matched_warehouse"] = value
}

func (c *Sundeck) AddAuto_routing_matched_warehouse_size(value any) {
	c.values["auto_routing_matched_warehouse_size"] = value
}

func (c *Sundeck) AddAuto_routing_num_computed_signatures(value any) {
	c.values["auto_routing_num_computed_signatures"] = value
}

func (c *Sundeck) AddLookup_routing_matched(value any) {
	c.values["lookup_routing_matched"] = value
}

func (c *Sundeck) AddLookup_routing_matched_warehouse(value any) {
	c.values["lookup_routing_matched_warehouse"] = value
}

func (c *Sundeck) AddLookup_routing_num_computed_signatures(value any) {
	c.values["lookup_routing_num_computed_signatures"] = value
}

func (c *Sundeck) AddAuto_routing_warehouse_pool(value any) {
	c.values["auto_routing_warehouse_pool"] = value
}

func (c *Sundeck) AddQuery_id(value any) {
	c.values["query_id"] = value
}
