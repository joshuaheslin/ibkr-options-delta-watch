go mod vendor
mv -f vendor myfunction/vendor
cd terraform
terraform apply -auto-approve