# ds-agent

Overview
---
A small agent written in Golang for handling image updates in Docker swarm mode services. Support for private registry with basic authentication(enabled as default). 

Build
---

go build -o ds-agent 

Usage
---

export DS_AGENT_REGISTRY_USERNAME=testuser  
export DS_AGENT_REGISTRY_PASSWORD=test123   
export DS_AGENT_PORT=3000  
export DS_AGENT_SECRET=pwd123

Optional:

export DS_AGENT_TLS=true|false


HTTPS
---

If TLS is enabled you need your certificate and matching private key in pem format. Put these in same folder as the app.

Using docker? Copy/mount into following:   

/root/crt.pem  
/root/key.pem

Docker
---
Dockerfile and docker-compose.yml are included.

Usage
---

Update service:

curl -H "Content-Type: application/json" -X POST -d '{"secret":"pwd123", "service":"ubuntu_service","image":"httpd"}' https://server.example.com:3000/services/update

List services:

curl -H "Content-Type: application/json" -X POST -d '{"secret":"pwd123"} ttps://se ver.example.com:3000/services

List containers:

curl -H "Content-Type: application/json" -X POST -d '{"secret":"pwd123"} ttps://se ver.example.com:3000/containers
