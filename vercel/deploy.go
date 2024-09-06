package vercel

import (
	"fmt"
	"strconv"
	"time"
)

const DEPLOYMENT_URL = "https://api.vercel.com/v13/deployments"

func GenerateDeploymentName() string {
	currentTime := strconv.Itoa(int(time.Now().Unix()))
	return fmt.Sprintf("deployment_%s", currentTime)
}

type VercelDeploymentFile struct {
	Data string `json:"data"`
	Encoding string `json:"encoding"` // Either 'base64' or 'utf-8'
	File string `json:"file"` // The filename including the whole path ie folder/file.js
}

func FilesToVercelDeploymentFiles(files map[string]string) []*VercelDeploymentFile {
	result := []*VercelDeploymentFile{}

	for filePath, fileContents := range files {
		result = append(result, &VercelDeploymentFile{
			Data: fileContents,
			Encoding: "utf-8",
			File: filePath,
		})
	}

	return result
}

type VercelDeploySettings struct {
	BuildCommand *string `json:"buildCommand,omitempty"` // Could be null in response hence pointer
	InstallCommand *string `json:"installCommand"` // Ditto ^
	Framework *string `json:"framework,omitempty"` // Ditto ^
}

func DefaultVercelDeploySettings() VercelDeploySettings {
	buildCommand := BUILD_COMMAND
	return VercelDeploySettings{
		BuildCommand: &buildCommand,
	}
}

type VercelBuild struct {
	Env []interface{} `json:"env"`
}

type VercelCreator struct {
	UID string `json:"uid"`
}

type VercelDeploymentResponse struct {
	AliasAssigned bool `json:"aliasAssigned"`
	BootedAt int `json:"bootedAt"`
	Build VercelBuild `json:"build"`
	BuildSkipped bool `json:"buildSkipped"`
	BuildingAt int `json:"buildingAt"`
	CreatedAt int `json:"createdAt"`
	CreatedIn string `json:"createdIn"`
	Creator VercelCreator `json:"creator"`
	Env []interface{} `json:"env"`
	ID string `json:"id"`
	InspectorURL *string `json:"inspectorUrl"` // Is string | null hence the pointer
	IsInCurrentBuildsQueue bool `json:"isInConcurrentBuildsQueue"`
	Meta map[string]interface{} `json:"meta"`
	Name string `json:"name"`
	OwnerID string `json:"ownerId"`
	Plan string `json:"plan"` // One of "pro", "enterprise", "hobby"
	ProjectID string `json:"projectId"`
	ProjectSettings VercelDeploySettings `json:"projectSettings"`
	Public bool `json:"public"`
	ReadyState string `json:"readyState"` // One of "CANCELED", "ERROR", "QUEUED", "BUILDING", "INITIALIZING", "READY"
	Regions []interface{} `json:"regions"`
	Routes *[]interface{} `json:"routes"`
	Status string `json:"status"` // One of "CANCELED", "ERROR", "QUEUED", "BUILDING", "INITIALIZING", "READY"
	Type string `json:"type"` // Should just be "LAMBDAS"
	URL string `json:"url"`
	Version int `json:"version"`
}

// Creates a deployment based on a dictionary mapping filepath to filecontent
func (v *VercelHTTPClient) CreateNewDeployment(projectName string, files map[string]string) (*VercelDeploymentResponse, error) {
	response := VercelDeploymentResponse{}

	err := v.MakePostRequest(DEPLOYMENT_URL, map[string]interface{}{
		"forceNew": 1,
		"skipAutoDetectionConfirmation": 1,
	}, map[string]interface{}{
		"name": GenerateDeploymentName(),
		"files": FilesToVercelDeploymentFiles(files),
		"project": projectName,
		"projectSettings": DefaultVercelDeploySettings(),
		"target": "production",
	}, &response)

	return &response, err
}
