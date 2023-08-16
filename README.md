# GitHub Actions Workshops

## first-action.yaml

Following is a simple GitHub Actions workflow named as `First Workflow`. It defines one job named `first-job` to print a greeting and a goodbye message.

```yaml
name: First Workflow 

# Define the events that trigger this workflow to run.
# 'workflow_dispatch' event allows manual triggering of the workflow from the GitHub Actions UI.
on: workflow_dispatch

jobs:
  first-job:
    # Define the runner environment for the 'first-job' job.
    runs-on: ubuntu-latest
    steps:
      # Step 1: Print a greeting message. (Multiple Command)
      - name: Print greeting
        run: |
          echo "Hello World"
          echo "Have a great day!"

      # Step 2: Print a goodbye message.
      - name: Print goodbye
        run: echo "Done - bye!"
```


## simple-task.yaml
Following is a GitHub Actions workflow for a Go project. It defines one job named `build` to build and test the Go project.
The job is triggered when there's a push event that modifies files in the `simple-task/` directory.

```yaml
name: Go

# Define the events that trigger this workflow to run.
# The workflow is triggered when there's a push event affecting files in the "simple-task/" directory.
on:
  push:
    paths:
      - simple-task/**

jobs:
  build:
    # Define the runner environment for the "build" job.
    runs-on: ubuntu-latest
    steps:
      # Step 1: Checkout the code from the repository.
      - uses: actions/checkout@v3

      # Step 2: Set up the Go environment with the specified version.
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'

      # Step 3: Build the Go project using the "go build" command.
      - name: Build
        working-directory: ./simple-task
        run: go build -v .

      # Step 4: Run tests for the Go project using the "go test" command.
      - name: Test
        working-directory: ./simple-task
        run: go test -v .

```


## multiple_job_parallel.yaml
Following is a GitHub Actions workflow for a Go project. It defines two jobs: `test` and `build`, to run tests and build the project respectively. These jobs run in parallel. And also the following workflow will be triggered in two events. 

- `push` event triggers the workflow when new code is pushed to the repository.
- `workflow_dispatch` event allows manual triggering of the workflow from the GitHub Actions UI.

```yaml
name: Go

# 'push' event triggers the workflow when new code is pushed to the repository.
# 'workflow_dispatch' event allows manual triggering of the workflow from the GitHub Actions UI.

on: [push, workflow_dispatch]

jobs:
  test:
    # Define the runner environment for the 'test' job.
    runs-on: ubuntu-latest
    steps:
      # Step 1: Download the code from the repository.
      - name: Download Code
        uses: actions/checkout@v3
  
      # Step 2: Set up Golang environment with the specified version and cache dependencies.
      - name: Set up Golang
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
          cache-dependency-path: simple-task/go.sum
        
      # Step 3: Run tests using 'go test' command.
      - name: Test
        working-directory: ./simple-task
        run: go test -v .
   
  build:
    # Define the runner environment for the 'build' job.
    runs-on: ubuntu-latest
    steps: 
      # Step 1: Download the code from the repository.
      - name: Download Code
        uses: actions/checkout@v3
      
      # Step 2: Set up Golang environment with the specified version and cache dependencies.
      - name: Set up Golang
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
          cache-dependency-path: simple-task/go.sum
      
      # Step 3: Build the project using 'go build' command.
      - name: Build
        working-directory: ./simple-task
        run: go build -v .

```
### Display: 

![image](https://github.com/shamimice03/github-actions-lab/assets/19708705/444ab558-0fb0-49a7-9b85-8c1b05ccf682)

## multiple_job_sequential.yaml
This is a GitHub Actions workflow for a Go project. It defines two jobs: 'test' and 'build', to run tests and build the project respectively. These jobs runs sequentially. 

- use `needs` keyword to run jobs sequentially

```yaml
name: Go

# Define the events that trigger this workflow to run.
# 'push' event triggers the workflow when new code is pushed to the repository.
# 'workflow_dispatch' event allows manual triggering of the workflow from the GitHub Actions UI.

# on: [push, workflow_dispatch]

on: workflow_dispatch

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      # Step 1: Download the code from the repository.
      - name: Download Code
        uses: actions/checkout@v3
  
      # Step 2: Set up Golang environment with the specified version and cache dependencies.
      - name: Set up Golang
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
          cache-dependency-path: simple-task/go.sum
        
      # Step 3: Run tests using 'go test' command.
      - name: Test
        working-directory: ./simple-task
        run: go test -v .
   
  build:
    # This job depends on the successful completion of the 'test' job.
    # It will run only if the 'test' job finishes successfully.
    needs: test
    runs-on: ubuntu-latest
    steps: 
      # Step 1: Download the code from the repository.
      - name: Download Code
        uses: actions/checkout@v3
      
      # Step 2: Set up Golang environment with the specified version and cache dependencies.
      - name: Set up Golang
        uses: actions/setup-go@v4
        with:
          go-version: '1.20'
          cache-dependency-path: simple-task/go.sum
      
      # Step 3: Build the project using 'go build' command.
      - name: Build
        working-directory: ./simple-task
        run: go build -v .
```

### Display:

![image](https://github.com/shamimice03/github-actions-lab/assets/19708705/e21234e6-f36d-4b2f-a1de-a359eeec18bb)

## contexts.yaml
This is a GitHub Actions workflow named `view contexts`.  The workflow is triggered manually using the `workflow_dispatch` event.The `view` job is defined to view and output various GitHub Actions contexts.

```yaml
name: view contexts

on:
  workflow_dispatch

jobs:
  view:
    # Define the runner environment for the 'view' job.
    runs-on: ubuntu-latest
    steps: 
      # Step 1: View and output the GitHub context using the 'github' context.
      - name: view GitHub Context
        run: echo "${{ toJSON(github) }}"

      # Step 2: View and output the job context using the 'job' context.
      - name: view job Context
        run: echo "${{ toJSON(job) }}"

      # Step 3: View and output the secrets context using the 'secrets' context.
      - name: view secrets Context
        run: echo "${{ toJSON(secrets) }}"

      # Step 4: View and output the needs context using the 'needs' context.
      - name: view needs Context
        run: echo "${{ toJSON(needs) }}"
```

## Docs:

- [about-github-hosted-runners](https://docs.github.com/en/actions/using-github-hosted-runners/about-github-hosted-runners)



## 008_artifact.yaml
- [OIDC](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)

- Create a policy:

  ```json
    {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Effect": "Allow",
              "Action": [
                  "s3:GetObject",
                  "s3:PutObject",
                  "s3:DeleteObject"
              ],
              "Resource": [
                  "arn:aws:s3:::<BUCKET-NAME>/*"
              ]
          },
          {
              "Effect": "Allow",
              "Action": "s3:ListBucket",
              "Resource": "arn:aws:s3:::<BUCKET-NAME>"
          }
      ]
    }
  ```
- Create a role with above policy and following `Trust relationships`

  ```json
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Federated": "arn:aws:iam::111111111132:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action": "sts:AssumeRoleWithWebIdentity",
        "Condition": {
          "StringEquals": {
            "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
          },
          "StringLike": {
            "token.actions.githubusercontent.com:sub": "repo:shamimice03/github-actions-lab:ref:refs/heads/main"
          }
        }
      }
    ]
  }
  ```

  - Add `Role_ARN`, `Region`, `Bucket-Name` to the repo `secrets`

  ***

  # Reusable workflow
  - https://docs.github.com/en/actions/using-workflows/reusing-workflows
  - Secrets: https://docs.github.com/en/actions/using-workflows/reusing-workflows#passing-inputs-and-secrets-to-a-reusable-workflow
    

  ## `018_reusable_2.yaml` and `019_use_reusabel_2.yaml`

  -  Similar steps: `008_artifact.yaml`
    - OIDC
    - S3 Bucket
    - IAM Policy
    - IAM Role
    - Repo `secrets` setup
