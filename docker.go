package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"log"
	"net/http"
)

type UpdateServiceObject struct {
	Service string `json:"service"`
	Image   string `json:"image"`
}

type ServiceListObject struct {
	Servicename string `json:"servicename"`
	Image       string `json:"image"`
}

type ContainerListObject struct {
	Name   []string `json:"name"`
	Image  string   `json:"image"`
	Status string   `json:"status"`
}

func (c Configuration) updateService(w http.ResponseWriter, r *http.Request) {

	updateServiceObject := UpdateServiceObject{}

	err := json.NewDecoder(r.Body).Decode(&updateServiceObject)
	if err != nil {
		log.Print("Failed parsing body.")
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services, err := findService(updateServiceObject.Service)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	auth := getAuthConfig(c.RegistryUser, c.RegistryPassword)

	services[0].Spec.TaskTemplate.ContainerSpec.Image = updateServiceObject.Image

	response, err := cli.ServiceUpdate(context.Background(),
		services[0].ID,
		services[0].Version,
		services[0].Spec,
		types.ServiceUpdateOptions{
			QueryRegistry:       true,
			EncodedRegistryAuth: auth,
		})
	if err != nil {
		log.Print(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	log.Print("INFO: ServiceUpdate, ", response)

	fmt.Fprintf(w, "OK")
}

func findService(name string) ([]swarm.Service, error) {

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
		err := errors.New("ERROR: Failed listing services")
		return service, err
	}

	if len(service) != 1 {
		err := errors.New("ERROR: Service not found")
		return service, err
	}
	return service, nil
}

func (c Configuration) listServices(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})

	if err != nil {
		log.Print("ERROR: Failed listing services")
	}

	var serviceList []ServiceListObject

	for _, service := range services {
		serviceList = append(serviceList, ServiceListObject{
			Servicename: service.Spec.Name,
			Image:       service.Spec.TaskTemplate.ContainerSpec.Image,
		})

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(serviceList)
}

func (c Configuration) listContainers(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var containerListObj []ContainerListObject

	for _, container := range containers {
		containerListObj = append(containerListObj, ContainerListObject{
			Name:   container.Names,
			Image:  container.Image,
			Status: container.Status,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(containerListObj)

}

// func generating auth string for registry
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
