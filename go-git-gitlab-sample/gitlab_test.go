/*
 * MIT License
 *
 * Copyright (c) 2021 ashing
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main

import "testing"

var token = "your-gitlab-token"
var baseUrl = "https://gitlab.com" // or other private gitlab site

func TestGitlabApi_ListProjects(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	projects, err := l.ListProjects("iptable")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("projects is", projects)

}

func TestListRepoTree(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	nodes, err := l.ListRepoTree(128, "", "master")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", nodes)
}

func TestListCommentsRecent(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}
	commits, err := l.ListCommitsRecent(238, "master", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("commits is %+v\n", commits)
}

func TestListMergeRequestRecent(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	res, err := l.ListMergeRequestsRecent(238)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res is %+v\n", res)
}

func TestListEventsRecent(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	res, err := l.ListEventsRecent(238)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res is %+v\n", res)
}

func TestGetProjectMeta(t *testing.T) {
	l, err := NewGitlab(token, baseUrl)
	if err != nil {
		t.Fatal(err)
	}

	res, err := l.GetProjectMeta(238)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res is %+v\n", res)
}
