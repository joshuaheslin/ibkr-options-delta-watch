# twitter-poster

Automatically make a tweet from content generated from OpenAI

## Setup

```bash
cd teraform/
terraform init
terrafrom apply
```

## Run code

Develop the code in `myfunction` module and test it.

```bash
go run main.go
```

## Deploy code

Deployment will use local `terraform.tfvars` (remember to copy over from .env)

```bash
bash deploy.sh
```

The GCP CloudFunction will take care of the rest every 2 days :D
