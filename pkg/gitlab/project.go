package gitlab

import (
	"fmt"
	"github.com/gopaytech/go-commons/pkg/types"

	gl "github.com/xanzy/go-gitlab"
)

type Project interface {
	Get(id NameOrId) (*gl.Project, error)
	GetDefaultBranch(id NameOrId) (*gl.Branch, error)
	Create(name string, parentID int) (*gl.Project, error)
}

type project struct {
	client *gl.Client
}

func (p *project) GetDefaultBranch(id NameOrId) (*gl.Branch, error) {
	branches, _, err := p.client.Branches.ListBranches(id.Get(), &gl.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}

	for _, branch := range branches {
		if branch.Default {
			return branch, nil
		}
	}

	return nil, fmt.Errorf("default branch for project %v is not found", id.Get())
}

func (p *project) Get(id NameOrId) (*gl.Project, error) {
	project, _, err := p.client.Projects.GetProject(id.Get(), &gl.GetProjectOptions{})
	if err != nil {
		return nil, err
	}
	return project, err
}

func (p *project) Create(name string, parentID int) (*gl.Project, error) {
	project, _, err := p.client.Projects.CreateProject(&gl.CreateProjectOptions{
		Name:        types.StringToPointerString(name),
		NamespaceID: types.IntToPointerInt(parentID),
		Path:        nil,
		Visibility:  gl.Visibility(gl.PrivateVisibility),
	})

	return project, err
}

func NewProject(client *gl.Client) Project {
	return &project{client: client}
}
