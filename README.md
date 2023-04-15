# twitter-poster

Automatically make a tweet from content generated from OpenAI

## Run code

Develop the code in `myfunction` module and test it.

```bash
go run main.go
```

## Deploy code

Deployment will use local `terraform.tfvars` (remember to copy over from .env)

`vendor` dir must be in myfunction in favour or go.mod (GCP functions requirement for improved cold starts)

```bash
bash deploy.sh
```

The GCP CloudFunction will take care of the rest every 2 days :D
