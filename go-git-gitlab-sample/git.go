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
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type vType string

const (
	branchType vType = "branch"
	tagType    vType = "tag"
)

type GitPull struct {
	ctx         context.Context
	gitlabToken string
	gitlabUrl   string
	task        *Task
	gitDir      string
}

type Task struct {
	Id           int    // task id
	ProjectId    int    // project id
	VersionType  vType  // branch or tag
	VersionValue string // eg: v1.5.0
	CommitId     string // 如果是 branch 需要进行 checkout
}

type GitlabInfo struct {
	token string
	url   string
}

func NewGitPull(ctx context.Context, info *GitlabInfo, task *Task) *GitPull {
	return &GitPull{
		ctx:         ctx,
		gitlabToken: info.token,
		gitlabUrl:   info.url,
		gitDir:      "/tmp",
		task:        task,
	}
}

func (g *GitPull) pullTag(gitlabNameDir, gitUrl, tag string) (string, error) {
	logBuf := new(bytes.Buffer)
	r, err := git.PlainClone(gitlabNameDir, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "",            // 使用 token 校验的话，username 可以为空
			Password: g.gitlabToken, // gitlab access token
		},
		URL:           gitUrl,
		Progress:      logBuf,
		ReferenceName: plumbing.ReferenceName(tag),
		Depth:         1, // tag 可以使用 depth=1 因为不需要 checkout commit
		SingleBranch:  true,
		Tags:          git.TagFollowing,
	})
	if err != nil {
		return "", err
	}

	// retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n%s\n%s\n", logBuf.String(), ref.Name(), ref.Hash().String()), nil
}

func (g *GitPull) pullBranchCommit(gitlabNameDir, gitUrl, branch, commitHash string) (string, error) {
	logBuf := new(bytes.Buffer)
	r, err := git.PlainClone(gitlabNameDir, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "",            // 使用 token 校验的话，username 可以为空
			Password: g.gitlabToken, // gitlab access token
		},
		URL:           gitUrl,
		Progress:      logBuf,
		ReferenceName: plumbing.ReferenceName(branch),
		//Depth:         1, // 晚点需要进行 checkout
		SingleBranch: true,
		Tags:         git.TagFollowing,
	})
	if err != nil {
		return "", err
	}

	// retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		log.Println("refHead1", err)
		return "", err
	}
	log.Println("ref is", ref.Hash())

	workTree, err := r.Worktree()
	if err != nil {
		log.Println("workTree", err)
		return "", err
	}

	err = workTree.Checkout(&git.CheckoutOptions{
		Hash:  plumbing.NewHash(commitHash),
		Force: true,
	})
	if err != nil {
		log.Println("checkout", err)
		return "", err
	}

	ref, err = r.Head()
	if err != nil {
		log.Println("refHead2", err)
		return "", err
	}
	log.Println("ref is", ref.Hash())

	return fmt.Sprintf("%s\n%s\n%s\n", logBuf.String(), ref.Name(), ref.Hash().String()), nil
}

func (g *GitPull) Pull() error {
	log.Println("开始拉取代码")

	gitlabApi, err := NewGitlab(g.gitlabToken, g.gitlabUrl)
	if err != nil {
		return err
	}
	project, err := gitlabApi.GetProjectById(g.task.ProjectId)
	if err != nil {
		return err
	}
	// 获取组名和仓库名
	gitUrl := project.HTTPURLToRepo
	gitGroups, gitName := project.Namespace.Path, project.Path
	log.Println(gitUrl, gitGroups, gitName)

	// 判断gitlab组目录
	gitlabNameDir := path.Join(g.gitDir, gitGroups, gitName)
	isExist, err := pathExists(gitlabNameDir)
	if err != nil {
		return err
	}

	// 如果存在则删除目录
	if isExist {
		err = os.RemoveAll(gitlabNameDir)
		if err != nil {
			return err
		}
	}

	// 新建目录
	if err = os.MkdirAll(gitlabNameDir, os.ModePerm); err != nil {
		return err
	}

	// 生成referenceName
	var referenceName string
	var result string
	switch g.task.VersionType {
	case branchType:
		referenceName = fmt.Sprintf("refs/heads/%s", g.task.VersionValue)
		result, err = g.pullBranchCommit(gitlabNameDir, gitUrl, referenceName, g.task.CommitId)
		if err != nil {
			return err
		}
	case tagType:
		referenceName = fmt.Sprintf("refs/tags/%s", g.task.VersionValue)
		result, err = g.pullTag(gitlabNameDir, gitUrl, referenceName)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("错误的 git 类型, %s\n", g.task.VersionType)
	}

	// 打印结果
	log.Printf("taskId:%d stdout:%s stderr:%s", g.task.Id, result, err)

	// 返回数据
	return err
}

// 检查文件是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
