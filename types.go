package main

import "net/http"

type Middleware func(http.Handler) http.HandleFunc
