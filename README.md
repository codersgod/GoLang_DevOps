# AWS Cost Optimizer (Go + Vercel)

A lightweight Go application that reads AWS cost and infrastructure data and shows it in a web dashboard.

## Features

- AWS cost breakdown (EC2, S3, Lambda, Total)
- EC2 instance inventory with CPU utilization
- Security resources overview (Security Groups, Key Pairs)
- Detailed tables for Security Groups and Key Pairs
- Local run support (Go server on port 8080)
- Container support (Docker and docker-compose)
- Vercel deployment support using `@vercel/go`

## Tech Stack

- Go (backend + API routes)
- AWS SDK for Go v2
- Static HTML/CSS/JS frontend
- Docker / Docker Compose
- Vercel

## Project Structure

```text
.
├── cmd/
│   └── main.go
├── api/
│   └── index.go
├── internal/
│   ├── api/
│   │   └── handler.go
│   ├── aws/
│   │   └── cost.go
│   └── service/
│       └── cost_service.go
├── pkg/
│   └── utils.go
├── web/
│   └── index.html
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── vercel.json
└── README.md
```

## Prerequisites

- Go 1.22+ installed and available in PATH
- AWS credentials configured (environment variables or AWS profile)
- Optional: Docker Desktop
- Optional: Vercel account linked to GitHub

## Required AWS Permissions

At minimum, the AWS identity used by this app should be able to call:

- `ce:GetCostAndUsage`
- `ec2:DescribeInstances`
- `ec2:DescribeRegions`
- `ec2:DescribeSecurityGroups`
- `ec2:DescribeKeyPairs`
- `cloudwatch:GetMetricStatistics`

## Environment Variables

Set these before running:

- `AWS_REGION`
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN` (only if you use temporary credentials)

### PowerShell example

```powershell
$env:AWS_REGION="us-east-1"
$env:AWS_ACCESS_KEY_ID="YOUR_ACCESS_KEY"
$env:AWS_SECRET_ACCESS_KEY="YOUR_SECRET_KEY"
# Optional for temporary credentials
$env:AWS_SESSION_TOKEN="YOUR_SESSION_TOKEN"
```

## Run Locally (Go)

```powershell
go mod tidy
go run ./cmd
```

App URL:

- `http://localhost:8080`

## Build Locally

```powershell
go build ./...
```

## Run with Docker

### Build and run container

```powershell
docker build -t cost-optimizer:local .
docker run --rm -p 8080:8080 `
  -e AWS_REGION `
  -e AWS_ACCESS_KEY_ID `
  -e AWS_SECRET_ACCESS_KEY `
  -e AWS_SESSION_TOKEN `
  cost-optimizer:local
```

### Using docker-compose

```powershell
docker compose up --build
```

## API Endpoints

Base URL (local): `http://localhost:8080`

- `GET /cost` - Cost summary for last 30 days
- `GET /ec2` - EC2 count and instance details
- `GET /services` - All services with costs
- `GET /security` - Count summary of security resources
- `GET /security-details` - Detailed Security Groups and Key Pairs

## Deploy to Vercel from GitHub

1. Initialize git repository (if not already done):

```powershell
git init
git add .
git commit -m "Initial commit"
```

2. Create a GitHub repository, then push:

```powershell
git branch -M main
git remote add origin https://github.com/<your-user>/<your-repo>.git
git push -u origin main
```

3. In Vercel:

- Create New Project
- Import your GitHub repository
- Leave Root Directory blank (repo root, where `vercel.json` is present)
- Build settings can remain default for this configuration

4. In Vercel Project Settings -> Environment Variables, add:

- `AWS_REGION`
- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`
- `AWS_SESSION_TOKEN` (if needed)

5. Deploy and verify:

- Open the deployed root URL `/`
- Check API routes `/cost`, `/ec2`, `/services`, `/security`, `/security-details`

## Notes and Troubleshooting

- If `go` command is not recognized, install Go and reopen terminal.
- If AWS calls fail, verify credentials and IAM permissions.
- If cost data is empty, check account billing access and Cost Explorer permissions.
- If CPU utilization shows `N/A`, ensure CloudWatch metrics are available for instances.

## License

Use this project as needed for internal/dev purposes. Add a formal license file if required by your organization.
