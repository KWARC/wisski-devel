# wisski-devel

This repository contains a docker-compose file that can be used as a development environment for a WissKI setup. 

To create a wisski instance, do:

- Copy the graphdb sources into `triplestore/graphdb.zip`
- `mkdir data`
- `docker-compose build`
- `docker-compose up -d`
- `docker-compose exec barrel /user_shell.sh` -- to get a shell to install drupal via composer from scratch

The credentials for sql are 'root:root'. 

This starts:
- an apache instance on port 8080
- a phpmyadmin on port 8081
- a graphdb on port 7200