# ds-agent

Overview
---
A small agent written in Golang for handling image updates in Docker swarm mode services. Support for private registry with basic authentication(enabled as default). Server also runs https as default.

Build
---

go build -o ds-agent 

Usage
---

export DS_AGENT_REGISTRY_USERNAME=testuser  
export DS_AGENT_REGISTRY_PASSWORD=test123   
export DS_AGENT_PORT=3000  
export DS_AGENT_SECRET=pwd123


HTTPS
---

Mount/copy your certificate and matching private key into:

/root/crt.pem  
/root/key.pem

Usage
---

Update service:

curl -H "Content-Type: application/json" -X POST -d '{"secret":"pwd123", "service":"ubuntu_service","image":"httpd"}' https://server.example.com:3000/services/update

List services:

curl https://server.example.com:3000/services
