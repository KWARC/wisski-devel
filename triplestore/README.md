# GraphDB Free Docker Image

This folder contains a Dockerfile for [GraphDB Free](https://graphdb.ontotext.com/).

>  GraphDB is an enterprise ready Semantic Graph Database, compliant with W3C Standards. Semantic graph databases (also called RDF triplestores) provide the core infrastructure for solutions where modelling agility, data integration, relationship exploration and cross-enterprise data publishing and consumption are important. 

## Using the Docker Image

For licensing reasons, it is not possible to distribute the resulting docker image. 
Instead the GraphDB Server Binaries must be retrieved from Onotext directly.

To build the Docker Image:
1. [Request a Download Link for GraphDB Free](https://www.ontotext.com/products/graphdb/graphdb-free/)
2. In the email, select 'Download as a stand-alone server'. 
3. Rename the downloaded file to `graphdb.zip` and place it in this folder. 
4. Build the docker image using `docker build -t graphdb .` or similar. 

## Legal

The code found in this folder is neither affiliated with nor endorsed by Onotext.
This code is licensed under the terms of the [The Unlicense](LICENSE) and (where applicable) in the public domain. 