package github

import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequest(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Github API",
		Description: "This is your Github API repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))
}
