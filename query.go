package main

import (
    "encoding/json"
    "strconv"
)

func generateStatusCode(n int)(string) {
    sn := strconv.Itoa(n)
    
    for i := len(sn) + 4; i < 15; i++ {
        sn = " " + sn
    }
    
    return "200 " + sn + "\n"
}

func addRoot(i []interface{})([]interface{}) {
    var root[] interface{}
    
    if i != nil {
        root = append(root, i)
    } else {
        var empty[] interface{} = make([]interface{}, 0)
        return empty
    }
    
    return root
}

func generateJson(ss[] interface{})(string) {
    j, err := json.Marshal(ss)
    
    if err != nil {
        return `["error"]`
    } else {
        return string(j)
    }
}

func processTableStatus()(string) {
    var result[] interface{}
    
    result = append(result, 1)
    result = append(result, 1)
    result = append(result, 1)
    result = append(result, 0)
    result = append(result, 1)
    result = append(result, 1)
    result = append(result, 0)
    result = append(result, 1)
    result = append(result, 1)
    result = append(result, 1)
    result = append(result, 1629390707)
    result = append(result, -1)
    result = append(result, "1.2.8p27")
    result = append(result, 0)
    result = append(result, 0)
    result = append(result, 0)
    result = append(result, 0)
    result = append(result, 1629301880)
    result = append(result, "18.10.0")
    result = append(result, 50)
        
    return generateJson(addRoot(result))
}

func processTableContacts()(string) {
    return `[]`
}

func processTableContactGroups()(string) {
    return `[]`
}

func processTableComments()(string) {
    return `[]`
}

func processTableDowntimes()(string) {
    return `[]`
}

func processTableCommands()(string) {
    return `[]`
}

func processTableHosts(pa Parameters)(string) {
    sql_filters := ""
    var result[] interface{}
    
    if len(pa.filters) > 0 {
        sql_filters = filters_to_sql(pa.filters)
    }
    
    stats_filters := stats_to_sql(pa.stats)
    sql_str := generate_hosts_sql(pa.columns, sql_filters, stats_filters, pa.limit)
    result = exec_sql(sql_str, pa.table)
    
    return generateJson(result)
}

func processTableServices(pa Parameters)(string) {
    sql_filters := ""
    var result[] interface{}
    
    if len(pa.filters) > 0 {
        sql_filters = filters_to_sql(pa.filters)
    }
    
    stats_filters := stats_to_sql(pa.stats)
    sql_str := generate_services_sql(pa.columns, sql_filters, stats_filters, pa.limit)
    result = exec_sql(sql_str, pa.table)
 
    return generateJson(result)
}

func processCommand(pa Parameters)(string) {
    result := ""
    statusCode := ""
    
    switch pa.table {
        case "status":
            result = processTableStatus()
            break
        case "hosts":
            result = processTableHosts(pa)
            break
        case "services":
            result = processTableServices(pa)
            break            
        case "contacts":
            result = processTableContacts()
            break
        case "contactgroups":
            result = processTableContactGroups()
            break
        case "comments":
            result = processTableComments()
            break            
        case "downtimes":
            result = processTableDowntimes()
            break           
        case "commands":
            result = processTableCommands()
            break               
        default:
            break
    }
    
    statusCode = generateStatusCode(len(result) + 1)

    return statusCode + result + "\n"
}
