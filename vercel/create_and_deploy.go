package vercel

func CreateAndDeploy(vercelToken string, files map[string]string) (*VercelDeploymentResponse, error) {
	client := CreateClient(vercelToken)

	project, err  := client.CreateNewProject(GenerateProjectName())

	if err != nil {
		return nil, err
	}

	return client.CreateNewDeployment(project.Name, files)
}
