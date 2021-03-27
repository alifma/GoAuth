package main

import "net/http"


const USERNAME = "batman"
const PASSWORD = "secret"

// func Auth(w http.ResponseWriter, r *http.Request) bool {
//     username, password, ok := r.BasicAuth()
//     if !ok {
//         w.Write([]byte(`something went wrong`))
//         return false
//     }

//     isValid := (username == USERNAME) && (password == PASSWORD)
//     if !isValid {
//         w.Write([]byte(`wrong username/password`))
//         return false
//     }

//     return true
// }



func MiddlewareAuth (next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if !ok {
            w.Write([]byte(`Something Went Wrong`))
            return
        }
        isValid := (username == USERNAME) && (password == PASSWORD)
        if !isValid {
            w.Write([]byte(`Missmatch Username & Password`))
            return
        }
        next.ServeHTTP(w, r)
    })
}

// func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
// 	if r.Method != "GET" {
// 			w.Write([]byte("Only GET is allowed"))
// 			return false
// 	}

// 	return true
// }

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            w.Write([]byte("Only GET is allowed"))
            return
        }

        next.ServeHTTP(w, r)
    })
}