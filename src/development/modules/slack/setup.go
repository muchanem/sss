package main

import (
	"fmt"
	"net/http"
	 u "github.com/whitman-colm/go-1/utils/other"
)


resp, err := http.Get("https://slack.com/api/api.test")
u.QuitAtError(err)
