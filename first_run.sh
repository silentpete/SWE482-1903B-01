#!/bin/bash

# anytime the script experiences an error, fail out of the script.
set -e

# install_docker_ce is used to install the docker package needed to host the Bundy's environment.
function install_docker_ce() {
  echo "install docker-ce"
  sudo yum install -y yum-utils device-mapper-persistent-data lvm2
  sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
  sudo yum-config-manager --enable docker-ce-edge
  sudo yum install -y docker-ce
  sudo systemctl enable docker
  sudo systemctl start docker
}

# install_docker_compose is used to install the docker-compose package needed to start the stack.
function install_docker_compose() {
  echo "install docker-compose"
  sudo curl -L https://github.com/docker/compose/releases/download/1.16.1/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
  sudo chmod +x /usr/local/bin/docker-compose
  sudo curl -L https://raw.githubusercontent.com/docker/compose/1.16.1/contrib/completion/bash/docker-compose -o /etc/bash_completion.d/docker-compose
}

# check if docker is an available command, if not, install docker
type which docker >/dev/null 2>&1
if [[ $? -ne 0 ]]; then
  install_docker_ce
fi

# check if docker-compose is an available command, if not, install docker-compose
type which docker-compose >/dev/null 2>&1
if [[ $? -ne 0 ]]; then
  install_docker_compose
fi

# Look for any build.sh file and then execute it.
for VAR in $(find . -type f -name "build.sh"); do
  echo "building ${VAR}"
  ${VAR}
done

# the first time we run the environment, we need to load the database from the inventory.json file.
# the following will stand up the mysql container, wait for it to finish starting for the first time,
# then it will start the bundys container with the load-database argument.
# this will only run correctly once. After than docker-compose up -d or docker-compose stop.
if [[ $(docker ps -a | grep -c "mysql") -eq 0 ]]; then
  docker run -dit --name mysql -e MYSQL_ROOT_PASSWORD=bundys -e MYSQL_DATABASE=bundys -p 3306:3306 -v mysql:/var/lib/mysql mysql:8.0.17 --default-authentication-plugin=mysql_native_password --skip-mysqlx
fi

# wait for mysql to start before continuing
mysql_starting=true
while ${mysql_starting}; do
  if [[ $(docker logs mysql 2>&1 | grep -c "3306") -eq 1 ]]; then
    mysql_starting=false
  else
    echo "waiting for mysql to finish starting up"
    sleep 1
  fi
done

# load the default inventory
docker run -it --rm --link=mysql -v $PWD/bundys/:/bundys/ bundys-api:latest --load-database --inventory-file=/bundys/inventory.json --sql-db-host=mysql

# stop the mysql container
docker stop mysql

# remove the mysql container
docker rm mysql

# start the bundys stack
docker-compose up -d
