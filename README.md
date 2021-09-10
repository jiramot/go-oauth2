## Generate Elliptic Curve keys
https://cloud.google.com/iot/docs/how-tos/credentials/keys
```
openssl ecparam -genkey -name prime256v1 -noout -out ec_private.pem
openssl ec -in ec_private.pem -pubout -out ec_public.pem
```