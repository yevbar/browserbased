package vercel

import (
	"fmt"
	"strconv"
	"time"
)

const PROJECTS_URL = "https://api.vercel.com/v10/projects"
const BUILD_COMMAND = "npm run build"
const VERCEL_FRAMEWORK = "nextjs"

func GenerateProjectName() string {
	currentTime := strconv.Itoa(int(time.Now().Unix()))
	return fmt.Sprintf("project_%s", currentTime)
}

type VercelCreateProjectResponse struct {
	AccountID string `json:"accountId"`
	DirectoryListing bool `json:"directoryListing"`
	ID string `json:"id"`
	Name string `json:"name"`
	NodeVersion string `json:"nodeVersion"`
}

func (v *VercelHTTPClient) CreateNewProject(projectName string) (*VercelCreateProjectResponse, error) {
	response := VercelCreateProjectResponse{}

	err := v.MakePostRequest(PROJECTS_URL, map[string]interface{}{}, map[string]interface{}{
		"name": projectName,
		"buildCommand": BUILD_COMMAND,
		"framework": VERCEL_FRAMEWORK,
	}, &response)

	return &response, err
}
