package main

const INT_TYPE = 0
const DBL_TYPE = 1
const STR_TYPE = 2
const GRP_TYPE = 3

type Column struct {
    alias string
    t int
}

func get_host_fields()(map[string]Column) {
    var host_columns map[string]Column = make(map[string]Column)
    
    host_columns["accept_passive_checks"] = Column{"passive_checks", INT_TYPE} // modified
    host_columns["acknowledged"] = Column{"acknowledged", INT_TYPE} // equal
    host_columns["action_url"] = Column{"action_url", STR_TYPE} // equal
    host_columns["action_url_expanded"] = Column{"null as action_url_expanded", STR_TYPE}  // modified - not found field
    host_columns["active_checks_enabled"] = Column{"active_checks", STR_TYPE}  // modified
    host_columns["address"] = Column{"address", STR_TYPE} // equal
    host_columns["alias"] = Column{"alias", STR_TYPE} // equal
    host_columns["check_command"] = Column{"check_command", STR_TYPE} // equal
    host_columns["check_freshness"] = Column{"check_freshness", INT_TYPE} // equal
    host_columns["check_interval"] = Column{"check_interval", DBL_TYPE} // equal
    host_columns["check_options"] = Column{"0 as check_options", INT_TYPE} // modified - probably always 0
    host_columns["check_period"] = Column{"check_period", STR_TYPE} // equal
    host_columns["check_type"] = Column{"check_type", INT_TYPE} // equal
    host_columns["checks_enabled"] = Column{"1 as checks_enabled", INT_TYPE} // modified - probably always 1
    host_columns["childs"] = Column{"null as childs", GRP_TYPE} // modified - treated as null for now
    host_columns["comments"] = Column{"null as comments", GRP_TYPE} // modified - treated as null for now
    host_columns["current_attempt"] = Column{"check_attempt", INT_TYPE} // modified
    host_columns["current_notification_number"] = Column{"0 as current_notification_number", INT_TYPE} // modified - treated as 0 for now
    host_columns["event_handler"] = Column{"event_handler", STR_TYPE} // equal
    host_columns["event_handler_enabled"] = Column{"event_handler_enabled", INT_TYPE} // equal
    host_columns["execution_time"] = Column{"execution_time", DBL_TYPE} // equal
    host_columns["custom_variable_names"] = Column{"null as custom_variable_names", GRP_TYPE} // modified - treated as null for now
    host_columns["custom_variable_values"] = Column{"null as custom_variable_values", GRP_TYPE} // modified - treated as null for now
    host_columns["first_notification_delay"] = Column{"first_notification_delay", DBL_TYPE} // equal
    host_columns["flap_detection_enabled"] = Column{"flap_detection", INT_TYPE} // modified
    host_columns["groups"] = Column{"null as groups", GRP_TYPE} // modified - treated as null for now
    host_columns["has_been_checked"] = Column{"checked", INT_TYPE} // modified
    host_columns["high_flap_threshold"] = Column{"high_flap_threshold", DBL_TYPE} // equal
    host_columns["icon_image"] = Column{"icon_image", STR_TYPE}  // equal
    host_columns["icon_image_alt"] = Column{"icon_image_alt", STR_TYPE} // equal
    host_columns["icon_image_expanded"] = Column{"icon_image", INT_TYPE} // modified
    host_columns["is_executing"] = Column{"0 as is_executing", INT_TYPE} // modified - needs be 0 for now
    host_columns["is_flapping"] = Column{"flapping", INT_TYPE} // modified
    host_columns["last_check"] = Column{"last_check", INT_TYPE} // equal
    host_columns["last_notification"] = Column{"last_notification", INT_TYPE} // equal
    host_columns["last_state_change"] = Column{"last_state_change", INT_TYPE} // equal
    host_columns["latency"] = Column{"latency", DBL_TYPE} // equal
    host_columns["low_flap_threshold"] = Column{"low_flap_threshold", DBL_TYPE} // equal
    host_columns["max_check_attempts"] = Column{"max_check_attempts", INT_TYPE} // equal
    host_columns["next_check"] = Column{"next_check", INT_TYPE} // equal
    host_columns["notes"] = Column{"notes", STR_TYPE} // equal
    host_columns["notes_expanded"] = Column{"notes", INT_TYPE} // modified
    host_columns["notes_url"] = Column{"notes_url", STR_TYPE} // equal
    host_columns["notes_url_expanded"] = Column{"notes_url", STR_TYPE} // modified
    host_columns["notification_interval"] = Column{"notification_interval", DBL_TYPE} // equal
    host_columns["notification_period"] = Column{"notification_period", STR_TYPE}  // equal
    host_columns["notifications_enabled"] = Column{"notify", INT_TYPE} // modified
    host_columns["name"] = Column{"name", STR_TYPE} // equal
    host_columns["num_services_crit"] = Column{"num_services_crit", INT_TYPE} // equal but generated using a new query
    host_columns["num_services_ok"] = Column{"num_services_ok", INT_TYPE} // equal but generated using a new query
    host_columns["num_services_pending"] = Column{"num_services_pending", INT_TYPE} // equal but generated using a new query
    host_columns["num_services_unknown"] = Column{"num_services_unknown", INT_TYPE} // equal but generated using a new query
    host_columns["num_services_warn"] = Column{"num_services_warn", INT_TYPE} // equal but generated using a new query
    host_columns["num_services"] = Column{"num_services", INT_TYPE} // equal but generated using a new query
    host_columns["obsess_over_host"] = Column{"obsess_over_host", INT_TYPE} // equal
    host_columns["parents"] = Column{"parents", GRP_TYPE} // equal but generated using a new query
    host_columns["percent_state_change"] = Column{"percent_state_change", INT_TYPE}
    host_columns["perf_data"] = Column{"perfdata", DBL_TYPE} // modified
    host_columns["plugin_output"] = Column{"TRIM(TRAILING '\n' FROM output) as output", STR_TYPE} // modified
    host_columns["long_plugin_output"] = Column{"TRIM(TRAILING '\n' FROM output) as output", STR_TYPE} // modified
    host_columns["process_performance_data"] = Column{"process_perfdata", INT_TYPE} // modified
    host_columns["retry_interval"] = Column{"retry_interval", DBL_TYPE} // equal
    host_columns["scheduled_downtime_depth"] = Column{"scheduled_downtime_depth", INT_TYPE} // equal
    host_columns["state"] = Column{"state", INT_TYPE} // equal
    host_columns["state_type"] = Column{"state_type", INT_TYPE} // equal
    host_columns["modified_attributes_list"] = Column{"null as modified_attributes_list", GRP_TYPE} // modified - treated as null for now
    host_columns["last_time_down"] = Column{"last_time_down", INT_TYPE} // equal
    host_columns["last_time_unreachable"] = Column{"last_time_unreachable", INT_TYPE} // equal
    host_columns["last_time_up"] = Column{"last_time_up", INT_TYPE} // equal
    host_columns["display_name"] = Column{"display_name", STR_TYPE} // equal
    host_columns["in_check_period"] = Column{"1 as in_check_period", INT_TYPE} // modified - treated as 1 for now
    host_columns["in_notification_period"] = Column{"1 as in_notification_period", INT_TYPE} // modified - treated as 1 for now
    host_columns["contacts"] = Column{"null as contacts", GRP_TYPE} // modified - treated as null for now    
    
    
    return host_columns
}

func get_service_fields()(map[string]Column) {
    var service_columns map[string]Column = make(map[string]Column)
    
    service_columns["accept_passive_checks"] = Column{"passive_checks", INT_TYPE} // modified
    service_columns["acknowledged"] = Column{"acknowledged", INT_TYPE} // equal
    service_columns["action_url"] = Column{"action_url", INT_TYPE} // equal
    service_columns["action_url_expanded"] = Column{"action_url", STR_TYPE} // modified
    service_columns["active_checks_enabled"] = Column{"active_checks", INT_TYPE} // modified
    service_columns["check_command"] = Column{"check_command", STR_TYPE} // modified
    service_columns["check_interval"] = Column{"check_interval", DBL_TYPE} // equal
    service_columns["check_options"] = Column{"0 as check_options", INT_TYPE} // modified - probably always 0
    service_columns["check_period"] = Column{"check_period", STR_TYPE} // equal
    service_columns["check_type"] = Column{"check_type", INT_TYPE} // equal
    service_columns["checks_enabled"] = Column{"1 as checks_enabled", INT_TYPE} // modified - probably always 1
    service_columns["comments"] = Column{"null as comments", GRP_TYPE} // modified - treated as null for now
    service_columns["current_attempt"] = Column{"check_attempt", INT_TYPE} // modified
    service_columns["current_notification_number"] = Column{"0 as current_notification_number", INT_TYPE} // modified - treated as 0 for now
    service_columns["description"] = Column{"description", STR_TYPE} // equal
    service_columns["event_handler"] = Column{"event_handler", STR_TYPE} // equal
    service_columns["event_handler_enabled"] = Column{"event_handler_enabled", INT_TYPE} // equal
    service_columns["custom_variable_names"] = Column{"null as custom_variable_names", GRP_TYPE} // modified - treated as null for now
    service_columns["custom_variable_values"] = Column{"null as custom_variable_values", GRP_TYPE} // modified - treated as null for now
    service_columns["execution_time"] = Column{"execution_time", DBL_TYPE} // equal
    service_columns["first_notification_delay"] = Column{"first_notification_delay", DBL_TYPE} // equal
    service_columns["flap_detection_enabled"] = Column{"flap_detection", INT_TYPE} // modified
    service_columns["groups"] = Column{"null as groups", GRP_TYPE} // modified - treated as null for now
    service_columns["has_been_checked"] = Column{"checked", INT_TYPE} // modified
    service_columns["high_flap_threshold"] = Column{"high_flap_threshold", DBL_TYPE} // equal
    service_columns["host_acknowledged"] = Column{"host_acknowledged", INT_TYPE} // equal but generated using a new query
    service_columns["host_action_url_expanded"] = Column{"host_action_url_expanded", STR_TYPE} // equal but generated using a new query
    service_columns["host_active_checks_enabled"] = Column{"host_active_checks_enabled", INT_TYPE} // equal but generated using a new query
    service_columns["host_address"] = Column{"host_address", STR_TYPE} // equal but generated using a new query
    service_columns["host_alias"] = Column{"host_alias", STR_TYPE} // equal but generated using a new query
    service_columns["host_checks_enabled"] = Column{"1 as host_checks_enabled", INT_TYPE} // modified - probably always 1
    service_columns["host_check_type"] = Column{"host_check_type", INT_TYPE} // equal but generated using a new query
    service_columns["host_latency"] = Column{"host_latency", INT_TYPE} // equal but generated using a new query
    service_columns["host_plugin_output"] = Column{"host_plugin_output", INT_TYPE} // equal but generated using a new query
    service_columns["host_perf_data"] = Column{"host_perf_data", INT_TYPE} // equal but generated using a new query
    service_columns["host_current_attempt"] = Column{"host_current_attempt", INT_TYPE} // equal but generated using a new query
    service_columns["host_check_command"] = Column{"host_check_command", STR_TYPE} // equal but generated using a new query
    service_columns["host_comments"] = Column{"null as comments", GRP_TYPE} // modified - treated as null for now
    service_columns["host_groups"] = Column{"null as host_groups", GRP_TYPE} // modified - treated as null for now
    service_columns["host_has_been_checked"] = Column{"host_has_been_checked", INT_TYPE} // modified - treated as null for now
    service_columns["host_icon_image_expanded"] = Column{"host_icon_image_expanded", STR_TYPE} // equal but generated using a new query
    service_columns["host_icon_image_alt"] = Column{"host_icon_image_alt", STR_TYPE} // equal but generated using a new query
    service_columns["host_is_executing"] = Column{"0 as host_is_executing", INT_TYPE} // modified - needs be 0 for now
    service_columns["host_is_flapping"] = Column{"host_is_flapping", INT_TYPE} // equal but generated using a new query
    service_columns["host_notes"] = Column{"host_notes", STR_TYPE} // equal but generated using a new query
    service_columns["host_name"] = Column{"host_name", STR_TYPE} // equal but generated using a new query
    service_columns["host_notes_url_expanded"] = Column{"host_notes_url_expanded", STR_TYPE} // equal but generated using a new query
    service_columns["host_notifications_enabled"] = Column{"host_notifications_enabled", INT_TYPE} // equal but generated using a new query
    service_columns["host_scheduled_downtime_depth"] = Column{"host_scheduled_downtime_depth", INT_TYPE} // equal but generated using a new query
    service_columns["host_state"] = Column{"host_state", INT_TYPE} // equal but generated using a new query
    service_columns["host_accept_passive_checks"] = Column{"host_accept_passive_checks", INT_TYPE} // equal but generated using a new query
    service_columns["host_last_state_change"] = Column{"host_last_state_change", INT_TYPE} // equal but generated using a new query
    service_columns["icon_image"] = Column{"icon_image", STR_TYPE} // equal
    service_columns["icon_image_alt"] = Column{"icon_image_alt", INT_TYPE} // equal
    service_columns["icon_image_expanded"] = Column{"icon_image", INT_TYPE} // modified
    service_columns["is_executing"] = Column{"0 as is_executing", INT_TYPE} // modified - needs be 0 for now
    service_columns["is_flapping"] = Column{"flapping", INT_TYPE} // modified
    service_columns["last_check"] = Column{"last_check", INT_TYPE} // equal
    service_columns["last_notification"] = Column{"last_notification", INT_TYPE} // equal
    service_columns["last_state_change"] = Column{"last_state_change", INT_TYPE} // equal
    service_columns["latency"] = Column{"latency", DBL_TYPE} // equal
    service_columns["low_flap_threshold"] = Column{"low_flap_threshold", DBL_TYPE} // equal
    service_columns["max_check_attempts"] = Column{"max_check_attempts", INT_TYPE} // equal
    service_columns["next_check"] = Column{"next_check", INT_TYPE} // equal
    service_columns["notes"] = Column{"notes", STR_TYPE} // equal
    service_columns["notes_expanded"] = Column{"notes", INT_TYPE} // modified
    service_columns["notes_url"] = Column{"notes_url", STR_TYPE} // equal
    service_columns["notes_url_expanded"] = Column{"notes_url", STR_TYPE} // modified
    service_columns["notification_interval"] = Column{"notification_interval", DBL_TYPE} // equal
    service_columns["notification_period"] = Column{"notification_period", STR_TYPE} // equal
    service_columns["notifications_enabled"] = Column{"notify", INT_TYPE} // modified
    service_columns["obsess_over_service"] = Column{"obsess_over_service", INT_TYPE} // equal
    service_columns["percent_state_change"] = Column{"percent_state_change", DBL_TYPE} // equal
    service_columns["perf_data"] = Column{"perfdata", INT_TYPE} // modified
    service_columns["plugin_output"] = Column{"TRIM(TRAILING '\n' FROM output) as output", STR_TYPE} // modified
    service_columns["process_performance_data"] = Column{"process_perfdata", INT_TYPE} // modified
    service_columns["retry_interval"] = Column{"retry_interval", DBL_TYPE} // equal
    service_columns["scheduled_downtime_depth"] = Column{"scheduled_downtime_depth", INT_TYPE} // equal
    service_columns["state"] = Column{"state", INT_TYPE} // equal
    service_columns["state_type"] = Column{"state_type", INT_TYPE} // equal
    service_columns["modified_attributes_list"] = Column{"null as modified_attributes_list", GRP_TYPE} // modified - treated as null for now
    service_columns["last_time_critical"] = Column{"last_time_critical", INT_TYPE} // equal
    service_columns["last_time_ok"] = Column{"last_time_ok", INT_TYPE} // equal
    service_columns["last_time_unknown"] = Column{"last_time_unknown", INT_TYPE} // equal
    service_columns["last_time_warning"] = Column{"last_time_warning", INT_TYPE} // equal
    service_columns["display_name"] = Column{"display_name", STR_TYPE} // equal
    service_columns["host_display_name"] = Column{"host_display_name", INT_TYPE} // equal but generated using a new query
    service_columns["host_custom_variable_names"] = Column{"null as custom_variable_names", GRP_TYPE} // modified - treated as null for now
    service_columns["host_custom_variable_values"] = Column{"null as custom_variable_values", GRP_TYPE} // modified - treated as null for now
    service_columns["in_check_period"] = Column{"1 as in_check_period", INT_TYPE} // modified - treated as 1 for now
    service_columns["in_notification_period"] = Column{"1 as in_notification_period", INT_TYPE} // modified - treated as 1 for now
    service_columns["host_parents"] = Column{"host_parents", GRP_TYPE} // equal but generated using a new query
    service_columns["long_plugin_output"] = Column{"TRIM(TRAILING '\n' FROM output) as output", STR_TYPE} // modified
    service_columns["contacts"] = Column{"null as contacts", GRP_TYPE} // modified - treated as null for now
    
    return service_columns
}

func get_column(name string, table string)(Column) {
    var c Column;
    
    if table == "hosts" {
        c = get_host_fields()[name]
    } else if table == "services" {
        c = get_service_fields()[name]
    }
    
    return c
}
