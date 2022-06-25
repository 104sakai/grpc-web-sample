FROM mariadb:10.9-rc
====

Overview

## Description

## Usage
```
docker build -t mariadb .
docker run -it -d --rm --name mariadb -e MYSQL_ROOT_PASSWORD=password -p 13306:3306 mariadb
```