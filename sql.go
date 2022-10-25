package main

import (
//    "fmt"
    "strings"
    "strconv"
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
) 

func filters_to_sql(in[] string) (string) {
    s := strings.Join(filter_handle(in), " AND ")

    return s
}

func stats_to_sql(in[] string) ([]string) {
    
    out := filter_handle(in)
    
    return out
}

func generate_stats_sql_cases(stats_filters[] string)(string) {
    cases := ""
    i := 0
        
    for _, s := range stats_filters {
        if i > 0 {
            cases = cases + ","
        }
        cases = cases + "COUNT(CASE WHEN (" + s + ") THEN 1 ELSE NULL END  ) as st" + strconv.Itoa(i)
        i++
    }
    
    return cases
}

func add_parent_table_if_need(columns[] string)(string) {
    for _, s := range columns {
        if s == "parents" {

sql_str :=            
`
    LEFT JOIN
        (
            SELECT
                hosts_hosts_parents.child_id,
                GROUP_CONCAT(hosts.name SEPARATOR ',') AS parents 
            FROM
                hosts_hosts_parents 
            JOIN
                hosts 
                    ON (
                        hosts_hosts_parents.parent_id = hosts.host_id
                    ) 
            GROUP BY
                hosts_hosts_parents.child_id 
            ORDER BY
                NULL
        ) AS p 
            ON (
                H.host_id = p.child_id
            )
`
            return sql_str
        }
    }
    
    return ""
}

func add_parent_services_table_if_need(columns[] string)(string) {
    for _, s := range columns {
        if s == "host_parents" {

sql_str :=            
`
    LEFT JOIN
        (
            SELECT
                hosts_hosts_parents.child_id,
                GROUP_CONCAT(hosts.name SEPARATOR ',') AS host_parents 
            FROM
                hosts_hosts_parents 
            JOIN
                hosts 
                    ON (
                        hosts_hosts_parents.parent_id = hosts.host_id
                    ) 
            GROUP BY
                hosts_hosts_parents.child_id 
            ORDER BY
                NULL
        ) AS p 
            ON (
                S.host_id = p.child_id
            )
`
            return sql_str
        }
    }
    
    return ""
}

func add_services_num_table_if_need(columns[] string)(string) {
    for _, s := range columns {
        if (s == "num_services_crit" || s == "num_services_ok" || s == "num_services_pending" ||
            s == "num_services_unknown" || s == "num_services_warn" || s == "num_services") {

sql_str :=            
`
LEFT JOIN

(
SELECT 
  host_id as host_id_t,
  COUNT(CASE WHEN state=0 THEN 1 ELSE NULL END) as num_services_crit,
  COUNT(CASE WHEN state=1 THEN 1 ELSE NULL END) as num_services_ok,
  COUNT(CASE WHEN state=2 THEN 1 ELSE NULL END) as num_services_pending,
  COUNT(CASE WHEN state=3 THEN 1 ELSE NULL END) as num_services_unknown,
  COUNT(CASE WHEN state=4 THEN 1 ELSE NULL END) as num_services_warn,
  COUNT(state) as num_services
FROM services where enabled = 1 group by host_id
) as S on H.host_id = S.host_id_t
`
            return sql_str            
        }
    }
    
    return "";
}

func add_host_table_if_need()(string) {
    
sql_str :=
`
  JOIN (
      SELECT host_id,      
state as host_state, acknowledged as host_acknowledged, scheduled_downtime_depth as host_scheduled_downtime_depth,
display_name as host_display_name, notify as host_notifications_enabled, notes as host_notes, notes_url as host_notes_url_expanded,
name as host_name, flapping as host_is_flapping, icon_image_alt as host_icon_image_alt, address as host_address, alias as host_alias,
check_type as host_check_type, latency as host_latency, output as host_plugin_output, perfdata as host_perf_data, check_attempt as host_current_attempt,
check_command as host_check_command, checked as host_has_been_checked, icon_image as host_icon_image_expanded, last_state_change as host_last_state_change,
action_url as host_action_url_expanded, active_checks as host_active_checks_enabled, passive_checks as host_accept_passive_checks
  
  FROM hosts) as H on (S.host_id = H.host_id)
`

    return sql_str
}
    
func generate_hosts_sql(columns[] string, sql_filters string, stats_filters[] string, limit string) (string) {
    
    cases := generate_stats_sql_cases(stats_filters)
    
    sql_str := "SELECT " + strings.Join(columns, ", ") + " " + cases + 
    " FROM hosts as H"
    
    sql_str = sql_str + add_services_num_table_if_need(columns) + add_parent_table_if_need(columns)
    
    sql_str = sql_str + " where H.enabled = 1"
    
    if sql_filters != "" {
        sql_str = sql_str + " and"
        sql_str = sql_str + " " + sql_filters
    }
    
    if limit != "" {
        sql_str = sql_str + " limit " + limit
    }
    
//    fmt.Println(sql_str)
    
    return sql_str
}

func generate_services_sql(columns[] string, sql_filters string, stats_filters[] string, limit string) (string) {
    cases := generate_stats_sql_cases(stats_filters)
    
    sql_str := "SELECT " + strings.Join(columns, ", ") + " " + cases + 
    " FROM services as S"
    
    sql_str = sql_str + add_parent_services_table_if_need(columns) + add_host_table_if_need()
    
    sql_str = sql_str + " where S.enabled = 1"

    if sql_filters != "" {
        sql_str = sql_str + " and"
        sql_str = sql_str + " " + sql_filters
    }
    
    if limit != "" {
        sql_str = sql_str + " limit " + limit
    }    
    
//    fmt.Println(sql_str)
    
    return sql_str
}

func exec_sql(sql_str string, table string)([]interface{}) {
    config := get_config()    
    db, err := sql.Open("mysql", config.Database.Username + ":" + config.Database.Password +
                        "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" +
                        config.Database.Name)

    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }
    
    res, err := db.Query(sql_str)
        
    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }
    
    return getSQLResult(res, table)
}

func handle_field(name string, value string, table string)(interface{}) {
    var empty[] string = make([]string, 0)
    
    t := STR_TYPE
    
    c := get_column(name, table)
    
    if c.alias != "" {
        t = c.t
    }
       
    switch t {
        case GRP_TYPE:
            if value == "" {
                return empty
            }
            r := strings.Split(value, ",")
            return r
        default:
            return value
    }
    
}

func getSQLResult(rows *sql.Rows, table string) ([]interface{}) {
    columnTypes, err := rows.ColumnTypes()
    
    if err != nil {
        log.Fatal(err)
    }

    count := len(columnTypes)
    finalRows := []interface{}{}
    
    for rows.Next() {

        scanArgs := make([]interface{}, count)

        for i, v := range columnTypes {
            
            //fmt.Println(v.DatabaseTypeName())
            
            switch v.DatabaseTypeName() {
            case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
                scanArgs[i] = new(sql.NullString)
                break;
            case "BOOL":
                scanArgs[i] = new(sql.NullBool)
                break;
            case "TINYINT", "BIGINT", "SMALLINT", "INT4", "INT":
                scanArgs[i] = new(sql.NullInt64)
                break;
            case "DOUBLE":
                scanArgs[i] = new(sql.NullFloat64)
                break;                
            default:
                scanArgs[i] = new(sql.NullString)
            }
        }

        err := rows.Scan(scanArgs...)

        if err != nil {
            log.Fatal(err)
        }

        //masterData := map[string]interface{}{}
        var masterData[] interface{}

        for i, v := range columnTypes {
            
            //fmt.Println(v.DatabaseTypeName())

            // manipulate special field
                        
            if z, ok := (scanArgs[i]).(*sql.NullBool); ok  {
                //masterData[v.Name()] = z.Bool
                masterData = append(masterData, z.Bool)
                
                continue;
            }

            if z, ok := (scanArgs[i]).(*sql.NullString); ok  {
                //masterData[v.Name()] = z.String
                
                masterData = append(masterData, handle_field(v.Name(), z.String, table))
                continue;
            }

            if z, ok := (scanArgs[i]).(*sql.NullInt64); ok  {
                //masterData[v.Name()] = z.Int64
                masterData = append(masterData, z.Int64)
                continue;
            }

            if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok  {
                //masterData[v.Name()] = z.Float64
                masterData = append(masterData, z.Float64)
                continue;
            }
            
            if z, ok := (scanArgs[i]).(*sql.NullInt32); ok  {
                //masterData[v.Name()] = z.Int32
                masterData = append(masterData, z.Int32)
                continue;
            }

            //masterData[v.Name()] = scanArgs[i]
            masterData = append(masterData, scanArgs[i])
        }

        finalRows = append(finalRows, masterData)
    }

    //z, err := json.Marshal(finalRows)            
        
    return finalRows
}
