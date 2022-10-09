package main

import "testing"

func TestShellCheck(t *testing.T) {
	json1, err := ShellCheck(`#!/bin/bash
rm -rf /
echo ${hello}
`)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%+v\n", json1)
}
