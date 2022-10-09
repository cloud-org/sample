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

import (
	"context"
	"testing"
)

var gitlabInfo = GitlabInfo{
	token: "your-gitlab-token",
	url:   "https://gitlab.com", // or other private gitlab site
}

func TestPullCode(t *testing.T) {
	g := NewGitPull(context.TODO(), &gitlabInfo, nil)
	res, err := g.pullTag(
		"/tmp/xxx",
		"https://xxxxx/xxxx/xxx.git",
		"refs/tags/v0.0.1",
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("res", res)
	return
}

func TestPullCommit(t *testing.T) {
	g := NewGitPull(context.TODO(), &gitlabInfo, nil)
	res, err := g.pullBranchCommit(
		"/tmp/xxx",
		"https://xxxx/xxx/xxxx.git",
		"refs/heads/master",
		"xxxxx",
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("res", res)
	return
}

func TestGitPull_Pull(t *testing.T) {

	g := NewGitPull(context.TODO(), &gitlabInfo, &Task{
		Id:           1,
		ProjectId:    238,
		VersionType:  branchType,
		VersionValue: "master",
		CommitId:     "xxx",
	})

	err := g.Pull()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
