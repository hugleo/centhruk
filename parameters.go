package main

import (
    "strings"
) 

type Parameters struct {
    command string
    table string
    columns[] string
    filters[] string
    stats[] string
    keepAlive bool
    limit string
}

func preProcessParameters(p[] string)(Parameters) {
    var command string = ""
    var table string = ""
    var keepAlive = false
    var columns[] string = nil
    var filters[] string
    var stats[] string
    var limit string = ""
    
    for _, s := range p {
        if strings.HasPrefix(s, "GET ") {
            command = "GET"
            table = strings.TrimPrefix(s, "GET ")
        } else if strings.HasPrefix(s, "KeepAlive: on") {
              keepAlive = false
        } else if strings.HasPrefix(s, "Limit: ") {
              t := strings.Split(s, ": ")
              limit = t[1]
        } else if strings.HasPrefix(s, "Columns: ") {
              columns = strings.Split(strings.TrimPrefix(s, "Columns: "), " ")
        } else if strings.HasPrefix(s, "Filter: ") {
              filters = append(filters, strings.TrimPrefix(s, "Filter: "))
        } else if strings.HasPrefix(s, "And: ") || strings.HasPrefix(s, "Or: ") {
              filters = append(filters, s)
        } else if strings.HasPrefix(s, "Stats: ") {
              stats = append(stats, strings.TrimPrefix(s, "Stats: "))
        } else if strings.HasPrefix(s, "StatsAnd: ") || strings.HasPrefix(s, "StatsOr: ") {
              stats = append(stats, strings.TrimPrefix(s, "Stats"))
        }
    }
    
    columns = preFormatColumns(columns, table)
    filters = preFormatFiltersStats(filters, table)
    stats = preFormatFiltersStats(stats, table)
    
    var pa Parameters
    pa.command = command
    pa.table = table
    pa.columns = columns
    pa.filters = filters
    pa.stats = stats
    pa.keepAlive = keepAlive
    pa.limit = limit
    
    return pa
}

func preFormatFiltersStats(cs[]string, table string)([] string) {
    var r[] string
    
    for _, s := range cs {
        t := strings.Split(s, " ")
        
        tt := ""
        if t[0] == "Or:" || t[0] == "And:" {
            tt = s
        } else {
            var c Column
            c = get_column(t[0], table)   
            column := c.alias
            conector := t[1]
    
            value := "" // Stats: name !=
            if len(t) >= 3 {
                value = t[2]
            }
    
            if (c.t == STR_TYPE || c.t == GRP_TYPE) {
                conector, value = normalize_to_sql(conector, value)
            }
         
            if strings.Contains(column, " as") {
                tt = "0"
            } else {
                tt = column + " " + conector + " " + value  
            }
                       
        }
        
        r = append(r, tt)
    }
    
    return r
}

func preFormatColumns(cs[]string, table string)([] string) {
    var r[] string
    
    for _, c := range cs {
        
        var s string
        s = get_column(c, table).alias
        r = append(r, s)
    }
    
    return r
}

func normalize_to_sql(oper string, content string)(string, string) {
    if oper == "~" {
        oper = "like"
        content = `%` + content + `%`
    }
    content = `'` +  content + `'`
    
    return oper, content
}
