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
	"errors"
	"fmt"
	"regexp"
	"time"

	gitlab "github.com/xanzy/go-gitlab"
)

type GitlabApi struct {
	client *gitlab.Client
}

// NewGitlab 初始化
func NewGitlab(token, baseUrl string) (*GitlabApi, error) {
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return nil, err
	}

	return &GitlabApi{client: git}, nil
}

func (l *GitlabApi) GetProjectById(projectId int) (*gitlab.Project, error) {
	project, _, err := l.client.Projects.GetProject(projectId, &gitlab.GetProjectOptions{})
	return project, err
}

// GetProject 获取 ProjectID
func (l *GitlabApi) GetProject(gitUrl string) (*gitlab.Project, error) {
	reg := regexp.MustCompile("^(http|https|git)://.*/(.*)/(.*).git$")
	if reg == nil {
		return nil, errors.New("regex regular regenerate failed")
	}

	regStringList := reg.FindStringSubmatch(gitUrl)
	if regStringList == nil {
		return nil, errors.New("string match fail")
	}

	nameSpaceStr := fmt.Sprintf("%s/%s", regStringList[2], regStringList[3])
	project, _, err := l.client.Projects.GetProject(nameSpaceStr, &gitlab.GetProjectOptions{})
	if err != nil {
		return nil, err
	}

	return project, err
}

// GetBranches 获取Branch
func (l *GitlabApi) GetBranches(projectId int) ([]string, error) {
	branches, _, err := l.client.Branches.ListBranches(projectId, &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 20,
		},
	})
	if err != nil {
		return nil, err
	}

	branchList := make([]string, 0)
	for _, data := range branches {
		branchList = append(branchList, data.Name)
	}

	return branchList, nil
}

// GetTags 获取Tag
func (l *GitlabApi) GetTags(projectId int) ([]string, error) {
	tags, _, err := l.client.Tags.ListTags(projectId, &gitlab.ListTagsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 20,
		},
	})
	if err != nil {
		return nil, err
	}

	tagList := make([]string, 0)
	for _, data := range tags {
		tagList = append(tagList, data.Name)
	}

	return tagList, nil
}

type ProjectMeta struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//SSHURLToRepo      string `json:"sshurl_to_repo"`
	//HTTPURLToRepo     string `json:"httpurl_to_repo"`
	NameWithNamespace string `json:"nameWithNamespace"`
	Path              string `json:"path"`
	PathWithNamespace string `json:"pathWithNamespace"`
	DefaultBranch     string `json:"defaultBranch"`
}

func (l *GitlabApi) ListProjects(search string) ([]ProjectMeta, error) {
	projects, _, err := l.client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Search: &search,
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 10,
		},
	})
	if err != nil {
		return nil, err
	}
	projectsMeta := make([]ProjectMeta, len(projects))
	for i := 0; i < len(projects); i++ {
		p := projects[i]
		projectsMeta[i] = ProjectMeta{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			//SSHURLToRepo:      p.SSHURLToRepo,
			//HTTPURLToRepo:     p.HTTPURLToRepo,
			NameWithNamespace: p.NameWithNamespace,
			Path:              p.Path,
			PathWithNamespace: p.PathWithNamespace,
			DefaultBranch:     p.DefaultBranch,
		}
	}
	return projectsMeta, nil
}

func (l *GitlabApi) GetProjectMeta(projectId int) (*ProjectMeta, error) {
	res, _, err := l.client.Projects.GetProject(projectId, &gitlab.GetProjectOptions{})
	if err != nil {
		return nil, err
	}

	return &ProjectMeta{
		ID:                res.ID,
		Name:              res.Name,
		Description:       res.Description,
		NameWithNamespace: res.NameWithNamespace,
		Path:              res.Path,
		PathWithNamespace: res.PathWithNamespace,
		DefaultBranch:     res.DefaultBranch,
	}, nil

}

type Node struct {
	Value  string `json:"value"`
	Label  string `json:"label"`
	IsLeaf bool   `json:"isLeaf"`
}

func (l *GitlabApi) ListRepoTree(projectId int, path string, ref string) ([]Node, error) {
	recursive := false
	nodes, _, err := l.client.Repositories.ListTree(projectId, &gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 100, // TODO: config if need
		},
		Path:      &path,      // 子目录
		Ref:       &ref,       // branch or tag
		Recursive: &recursive, // 是否递归 默认 false
	})

	if err != nil {
		return nil, err
	}

	treeNodes := make([]Node, len(nodes))
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		treeNodes[i] = Node{
			Value: node.Path,
			Label: node.Name,
			IsLeaf: func(nodeType string) bool {
				if nodeType == "tree" {
					return false
				}
				// 其他类型返回 true 表示叶子节点
				return true
			}(node.Type),
		}
	}

	return treeNodes, nil
}

type Commit struct {
	CommitId      string `json:"commitId"`
	Title         string `json:"title"`
	Message       string `json:"message"`
	AuthorName    string `json:"authorName"`
	CommitterName string `json:"committerName"`
	CommitterDate string `json:"committerDate"`
}

func (l *GitlabApi) ListCommitsRecent(projectId int, ref string, pageNum int) ([]Commit, error) {
	res, _, err := l.client.Commits.ListCommits(projectId, &gitlab.ListCommitsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: pageNum,
		},
		RefName: &ref,
	})
	if err != nil {
		return nil, err
	}

	commits := make([]Commit, len(res))
	for i := 0; i < len(res); i++ {
		singleCommit := res[i]
		commits[i] = Commit{
			CommitId:      singleCommit.ID,
			Title:         singleCommit.Title,
			Message:       singleCommit.Message,
			AuthorName:    singleCommit.AuthorName,
			CommitterName: singleCommit.CommitterName,
			CommitterDate: singleCommit.CommittedDate.In(time.Local).Format(time.RFC3339),
		}
	}

	return commits, nil
}

type MergeRequestMeta struct {
	Title        string `json:"title"`
	SourceBranch string `json:"sourceBranch"`
	TargetBranch string `json:"targetBranch"`
	CreateAt     string `json:"createAt"`
	UpdateAt     string `json:"updateAt"`
	AuthorName   string `json:"authorName"`
}

func (l *GitlabApi) ListMergeRequestsRecent(projectId int) ([]MergeRequestMeta, error) {
	res, _, err := l.client.MergeRequests.ListProjectMergeRequests(projectId, &gitlab.ListProjectMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 10,
		},
	})
	if err != nil {
		return nil, err
	}

	mergeRequests := make([]MergeRequestMeta, len(res))
	for i := 0; i < len(res); i++ {
		singleMergeRequest := res[i]
		mergeRequests[i] = MergeRequestMeta{
			Title:        singleMergeRequest.Title,
			SourceBranch: singleMergeRequest.SourceBranch,
			TargetBranch: singleMergeRequest.TargetBranch,
			CreateAt:     singleMergeRequest.CreatedAt.In(time.Local).Format(time.RFC3339),
			UpdateAt:     singleMergeRequest.UpdatedAt.In(time.Local).Format(time.RFC3339),
			AuthorName:   singleMergeRequest.Author.Name,
		}
	}

	return mergeRequests, nil
}

type EventMeta struct {
	CreateAt            string `json:"createAt"`
	AuthorName          string `json:"authorName"`
	ActionName          string `json:"actionName"`
	PushDataRef         string `json:"pushDataRef"`
	PushDataCommitTitle string `json:"pushDataCommitTitle"`
}

func (l *GitlabApi) ListEventsRecent(projectId int) ([]EventMeta, error) {
	res, _, err := l.client.Events.ListProjectVisibleEvents(projectId, &gitlab.ListContributionEventsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    0,
			PerPage: 10,
		},
	})

	if err != nil {
		return nil, err
	}

	events := make([]EventMeta, len(res))
	for i := 0; i < len(res); i++ {
		singleEvent := res[i]
		events[i] = EventMeta{
			CreateAt:            singleEvent.CreatedAt.In(time.Local).Format(time.RFC3339),
			AuthorName:          singleEvent.Author.Name,
			ActionName:          singleEvent.ActionName,
			PushDataRef:         singleEvent.PushData.Ref,
			PushDataCommitTitle: singleEvent.PushData.CommitTitle,
		}
	}

	return events, nil

}
