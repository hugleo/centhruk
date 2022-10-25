package main

import (
    "fmt"
)

func test_tree_sql() {
    //in := []string{"name != teste", "has_been_checked = 0", "state = 0", "has_been_checked = 1", "And: 2", "state = 1", "state = 1", "active_checks_enabled = 1", "acknowledged = 0", "scheduled_downtime_depth = 0", "And: 4", "state = 2", "state = 2", "active_checks_enabled = 1", "acknowledged = 0", "scheduled_downtime_depth = 0", "And: 4"}
    
    //in := []string{"description = Filesystem", "state = 0", "Or: 2"}
    
    /*GET hosts
Columns: accept_passive_checks acknowledged action_url action_url_expanded active_checks_enabled address alias check_command check_freshness check_interval check_options check_period check_type checks_enabled childs comments current_attempt current_notification_number event_handler event_handler_enabled execution_time custom_variable_names custom_variable_values first_notification_delay flap_detection_enabled groups has_been_checked high_flap_threshold icon_image icon_image_alt icon_image_expanded is_executing is_flapping last_check last_notification last_state_change latency low_flap_threshold max_check_attempts name next_check notes notes_expanded notes_url notes_url_expanded notification_interval notification_period notifications_enabled num_services_crit num_services_ok num_services_pending num_services_unknown num_services_warn num_services obsess_over_host parents percent_state_change perf_data plugin_output process_performance_data retry_interval scheduled_downtime_depth state state_type modified_attributes_list last_time_down last_time_unreachable last_time_up display_name in_check_period in_notification_period long_plugin_output
has_been_checked = 0
state = 1
has_been_checked = 1
And: 2
state = 2
has_been_checked = 1
And: 2
Or: 3
OutputFormat: json
ResponseHeader: fixed16*/
    
    in := []string{"has_been_checked = 0", "state = 1", "has_been_checked = 1", "And: 2", "state = 2", "has_been_checked = 1", "And: 2"}
    
    //s := stats_to_sql(in)
    
    //s := filters_to_sql(in)
    
    //s := tree_to_str(generate_tree(in))
 
    /*for _, s1 := range s {
        fmt.Printf("Sx = %s\n", s1)
    }*/
    
    //s := tree_to_str(generate_tree(in))
    s := filter_handle(in)
    
    for i := 0; i < len(s); i ++ {
        fmt.Printf("FILTER %d = %s\n", i, s[i])
    }
    
    
    
//TREE = PARENT (XXX PARENT (Or PARENT (XXX PARENT (XXX has_been_checked = 0) PARENT (And state = 1 has_been_checked = 1)) PARENT (And state = 2 has_been_checked = 1)))

    //como esta fazendo: ((checked = 0 And (state = 1 And checked = 1)) Or (state = 2 And checked = 1)))
    
    //FILTER = [(has_been_checked = 0 Or (state = 1 And has_been_checked = 1) Or (state = 2 And has_been_checked = 1))]

    
   /* has_been_checked = 0

state = 1 AND has_been_checked = 1


state = 2 AND has_been_checked = 1*/
    
/*or {
    - XXX
    - AND --
    - AND - -
}*/
   
    //return s
}


/*func test_func(pa Parameters)(string) {
    //return filters_to_sql(pa.filters)
    
//    s := ""
//    r := stats_to_sql(pa.stats)
    
//    for i := 0; i < len(r); i++ {
  //      s = s + " -> " + r[i]
//    }
    
    sql_filters := ""
    sql_stats := ""
    if len(pa.filters) > 0 {
        sql_filters = filters_to_sql(pa.filters)
    }
    
    s := generate_hosts_sql(pa.columns, sql_filters, sql_stats, "")
    
    return s
}*/
