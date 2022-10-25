package main

import (
    "strings"
    "strconv"
)

func filter_handle(in []string)([]string) {
    var out []string
    size := len(in)
    i := 0
    
    for {
        if i < size {
            if (strings.HasPrefix(in[i], "Or: ") || strings.HasPrefix(in[i], "And: ")) {
                t := strings.Split(in[i], ": ")
                n, _ := strconv.Atoi(t[1])
                out = filter_handle_helper(out, t[0], n)
            } else {
                out = append(out, in[i])
            }
        } else {
            break
        }
        i++
    }       
    
    return out
}

func filter_handle_helper(in []string, con string, n int)([]string) {
    p := len(in) - 1
    
    if (p < (n -1)) || n <= 1 { // ignore for now, should be a invalid input
        return in
    }
    
    s := ""
       
    for i := n; i >= 1; i -- {
        e := in[p]
        if i == 1 {
           s = "(" + e + " " + con + " " + s
           in[p] = s
        } else if i == n {
            s = e + ")"
            in = in[:p]
        } else {
            s = e + " " + con + " " + s
            in = in[:p]
        }
        p--
    }
    
    return in
}
