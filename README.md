# ds-agent

Overview
---
A small agent written in Golang for handling image updates in Docker swarm mode services.

Why
---

Usage
---

export DS_AGENT_REGISTRY_USERNAME=testuser  
export DS_AGENT_REGISTRY_PASSWORD=test123<br>
export DS_AGENT_PORT=3000  

Update service:

curl -H "Content-Type: application/json" -X POST -d '{"service":"ubuntu_service","image":"httpd"}' http://localhost:3000/update

List services:

curl http://localhost:3000/services
