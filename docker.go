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

type UpdateObject struct {
	Secret  string
	Service string
	Image   string
}

func (c Configuration) updateService(w http.ResponseWriter, r *http.Request) {

	updateObject := UpdateObject{}

	err := json.NewDecoder(r.Body).Decode(&updateObject)
	if err != nil {
		log.Print("Failed parsing body.")
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	if updateObject.Service == "" {
		log.Print("Missing Service name")
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	} else if updateObject.Image == "" {
		log.Print("Missing Image name")
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	} else if updateObject.Secret != c.Secret {
		log.Print("Unauthorized")
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services, err := findService(updateObject.Service)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	auth := getAuthConfig(c.RegistryUser, c.RegistryPassword)

	services[0].Spec.TaskTemplate.ContainerSpec.Image = updateObject.Image

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
		err := errors.New("Failed listing services.")
		return service, err
	}

	if len(service) != 1 {
		err := errors.New("Service not found.")
		return service, err
	}
	return service, nil
}

func listServices(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	services, err := cli.ServiceList(context.Background(), types.ServiceListOptions{})

	if err != nil {
		log.Print("Failed listing services.")
	}

	for _, service := range services {
		fmt.Fprintf(w, "name: %q, image %q\n", service.Spec.Name, service.Spec.TaskTemplate.ContainerSpec.Image)
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
