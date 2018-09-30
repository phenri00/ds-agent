# ds-agent

Overview
---
A small agent written in Golang for handling image updates in Docker swarm mode services.

Build
---

go build -o ds-agent 

Usage
---

export DS_AGENT_REGISTRY_USERNAME=testuser  
export DS_AGENT_REGISTRY_PASSWORD=test123   
export DS_AGENT_PORT=3000  
export DS_AGENT_SECRET=pwd123

Update service:

curl -H "Content-Type: application/json" -X POST -d '{"secret":"pwd123", "service":"ubuntu_service","image":"httpd"}' http://localhost:3000/services/update

List services:

curl http://localhost:3000/services
