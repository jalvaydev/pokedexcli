package main

import "os"

func commandExit(params []string, cfg *config) error {
	os.Exit(0)
	return nil
}
