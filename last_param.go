package main

import ("fmt" 
"os" )

func last_param(s string) string {
    if len(os.Args)>1{
        return os.Args[len(os.Args)-1]
    }
    return "" 
}

