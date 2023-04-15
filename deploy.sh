go mod vendor
rm -rf myfunction/vendor
mv -f vendor myfunction/vendor
cd terraform
terraform apply -auto-approve