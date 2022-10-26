<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [go-git-gitlab-sample](#go-git-gitlab-sample)
  - [usage](#usage)
  - [acknowledgment](#acknowledgment)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## go-git-gitlab-sample

```tree
.
├── git.go          // git 相关操作，checkout branch or tag
├── git_test.go     // git unit test
├── gitlab.go       // gitlab api
├── gitlab_test.go  // gitlab unit test
```

### usage

```go
package main

import (
	"context"
	"fmt"
	gitlab "go-git-gitlab-sample"
	"log"
	"os"
	"path"
	"time"
)

var token = "xxx"
var baseUrl = "xxx" // or other private gitlab site
var gitlabInfo = gitlab.GitlabInfo{
	Token: token,
	Url:   baseUrl, // or other private gitlab site
}

func main() {
	l, err := gitlab.NewGitlab(token, baseUrl)
	if err != nil {
		log.Println(err)
		return
	}

	projects, err := l.ListProjects("")
	if err != nil {
		log.Println(err)
		return
	}

	clone(projects)

	return
}

func clone(projects []gitlab.ProjectMeta) {
	ctx := context.TODO()
	g := gitlab.NewGitPull(ctx, &gitlabInfo, nil)
	for _, item := range projects {
		log.Println(item)
		//if !filterRepo(item) {
		//	continue
		//}
		// 判断gitlab组目录
		gitlabNameDir := path.Join("/tmp/export", item.Group, item.Path)
		isExist, err := gitlab.PathExists(gitlabNameDir)
		if err != nil {
			log.Println(err)
			return
		}

		// 如果存在则删除目录
		if isExist {
			err = os.RemoveAll(gitlabNameDir)
			if err != nil {
				log.Println(err)
				return
			}
		}

		// 新建目录
		if err = os.MkdirAll(gitlabNameDir, os.ModePerm); err != nil {
			log.Println(err)
			return
		}

		referenceName := fmt.Sprintf("refs/heads/%s", item.DefaultBranch)
		res, err := g.PullBranch(gitlabNameDir, item.HTTPURLToRepo, referenceName)
		if err != nil {
			log.Println(err)
			continue
		}

		time.Sleep(3 * time.Second)
		log.Println(res)
	}
}

```

### acknowledgment

- https://github.com/go-git/go-git
- https://github.com/xanzy/go-gitlab
