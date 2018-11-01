## Retrieve keys from kops
```
aws s3 sync s3://kops-state-b429b/kubernetes.killabayte.com/pki/private/ca/ ca-key
aws s3 sync s3://kops-state-b429b/kubernetes.killabayte.com/pki/issued/ca/ ca-crt
mv ca-key/*.key ca.key
mv ca-crt/*.crt ca.crt
```
## Create new user
```
sudo apt install openssl
openssl genrsa -out {USER_NAME}.pem 2048
openssl req -new -key {USER_NAME}.pem -out {USER_NAME}-csr.pem -subj "/CN={USER_NAME}/O=devteam/"
openssl x509 -req -in {USER_NAME}-csr.pem -CA ca.crt -CAkey ca.key -CAcreateserial -out {USER_NAME}.crt -days 10000
```

## add new context
```
kubectl config set-credentials {USER_NAME} --client-certificate={USER_NAME}.crt --client-key={USER_NAME}.pem
kubectl config set-context {USER_NAME} --cluster=kubernetes.killabayte.com --user {USER_NAME}
```
