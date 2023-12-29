# IBKR-delta-tracker

Automatically track delta for optionSymbols by marketData API KEY.

# How to run

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

The GCP CloudFunction will take care of the rest every x days :D

# Improvements

- Connect with google sheets so we can take the OptionSymbol source for a list.