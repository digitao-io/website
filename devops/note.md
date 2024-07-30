# Setup the Server

## Install Docker:

Reference: https://docs.docker.com/engine/install/ubuntu/

```
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done

apt-get update
apt-get install ca-certificates curl
install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
chmod a+r /etc/apt/keyrings/docker.asc

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  tee /etc/apt/sources.list.d/docker.list > /dev/null

apt-get update

apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## Install MongoDB Shell:

Reference: https://www.mongodb.com/docs/mongocli/current/install/

```
apt-get install gnupg
wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list

apt-get update
apt-get install -y mongodb-mongosh

mongosh --version
```

## Install MC for Minio:

Reference: https://min.io/docs/minio/linux/reference/minio-mc.html

```
curl https://dl.min.io/client/mc/release/linux-amd64/mc --output mc
chmod +x mc
mc --help
```

## Install Nginx

Reference: https://ubuntu.com/tutorials/install-and-configure-nginx#2-installing-nginx

```
apt-get update
apt-get install nginx
```

## Install Certbot

Reference: https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-22-04

```
snap install --classic certbot
```
