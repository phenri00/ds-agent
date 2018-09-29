package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"log"
	"net/http"
)

type UpdateObject struct {
	Service string
	Image   string
}

func (c Configuration) updateService(w http.ResponseWriter, r *http.Request) {

	updateObject := UpdateObject{}

	err := json.NewDecoder(r.Body).Decode(&updateObject)
	if err != nil {
		panic(err)
	}

	//TODO
	//Verify json/struct
	serviceName := updateObject.Service
	imageName := updateObject.Image

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services := findServiceInfo(serviceName)

	auth := getAuthConfig(c.RegistryUser, c.RegistryPassword)

	services[0].Spec.TaskTemplate.ContainerSpec.Image = imageName

	response, err := cli.ServiceUpdate(context.Background(),
		services[0].ID,
		services[0].Version,
		services[0].Spec,
		types.ServiceUpdateOptions{
			QueryRegistry:       true,
			EncodedRegistryAuth: auth,
		})
	if err != nil {
		panic(err)
	}

	log.Print("INFO: ServiceUpdate, ", response)

	fmt.Fprintf(w, "OK")
}

func findServiceInfo(name string) []swarm.Service {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	filters := filters.NewArgs()
	filters.Add("name", name)

	service, err := cli.ServiceList(context.Background(), types.ServiceListOptions{
		Filters: filters,
	})
	if err != nil {
		panic(err)
	}
	if len(service) != 1 {
		log.Print("Service not found.")
	}

	return service
}

func listServices(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		fmt.Fprintf(w, service.ID, service.Spec.Name, service.Spec.TaskTemplate.ContainerSpec.Image,
			service.Version.Index)
	}
}

func getAuthConfig(userName string, password string) string {

	authConfig := types.AuthConfig{
		Username: userName,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	return authStr
}
