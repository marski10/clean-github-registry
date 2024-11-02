package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type PackageID struct {
	ID       int       `json:"id"`
	UPDATED  time.Time `json:"updated_at"`
	METADATA struct {
		CONTAINER struct {
			TAGS []string `json:"tags"`
		} `json:"container"`
	} `json:"metadata"`
}

func main() {

	if repository != '' {
	
		var packageId []PackageID
		packageId = PackageID{}.GetPackagesId(repository)
		for _, id := range packageId {

			if len(id.METADATA.CONTAINER.TAGS) > 1 {
				fmt.Println("You can't delete latest version")
			} else {
				PackageID{}.DeletePackage(id, repo.NAME)
			}
		}
	}else if {
		var repoNames []RepoName
		repoNames = RepoName{}.GetListRepository()
		for _, repo := range repoNames {

			fmt.Println(repo.NAME)
			var packageId []PackageID
			packageId = PackageID{}.GetPackagesId(repo.NAME)
	
			for _, id := range packageId {
	
				if len(id.METADATA.CONTAINER.TAGS) > 1 {
					fmt.Println("You can't delete latest version")
				} else {
					PackageID{}.DeletePackage(id, repo.NAME)
				}
			}
		}
	}

	


}

func (PackageID) DeletePackage(id PackageID, repository string) (bool, error) {

	tokengithub := os.Getenv("GITHUB_TOKEN")
	organization := os.Getenv("ORGANIZATION")
		
	endpoint := "https://api.github.com/orgs/" + agro-amazonia + "/packages/container/" + repository + "/versions/" + fmt.Sprintf("%d", id.ID)
	request, err := http.NewRequest("DELETE", endpoint, nil)

	request.Header.Add("Authorization", "Bearer "+tokengithub)

	if err != nil {
		log.Fatalln("Error in treatment endpoint")
		return false, err
	}

	consumer := &http.Client{}
	resp, err := consumer.Do(request)

	if err != nil {
		log.Fatalln("Error while making request")
		return false, err
	}

	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 204 {
		return true, nil

	} else {
		log.Fatalln("Error while delete package")
		return false, nil
	}

}

func (PackageID) GetPackagesId(repository string) []PackageID {

	tokengithub := os.Getenv("GITHUB_TOKEN")
	organization := os.Getenv("ORGANIZATION")
	endpoint := "https://api.github.com/orgs/" + organization + "/packages/container/" + repository + "/versions"

	fmt.Println(endpoint)

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		log.Fatalln("error in treatment endpoint")
		return []PackageID{}
	}

	request.Header.Add("Authorization", "Bearer "+tokengithub)
	request.Header.Set("Accept", "application/json")

	consumer := &http.Client{}

	resp, err := consumer.Do(request)

	if err != nil {
		log.Fatalln("Error in consume endpoint")
	}

	body, err := io.ReadAll(resp.Body)

	var JsonBody []PackageID

	err = json.Unmarshal(body, &JsonBody)

	return JsonBody

}

type RepoName struct {
	NAME string `json:"name"`
}

func (RepoName) GetListRepository() []RepoName {
	tokengithub := os.Getenv("GITHUB_TOKEN")
	organization := os.Getenv("ORGANIZATION")
	endpoint := "https://api.github.com/orgs/" + organization + "/repos"

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		log.Fatalln("error in treatment endpoint")
		return []RepoName{}
	}

	request.Header.Add("Authorization", "Bearer "+tokengithub)
	request.Header.Set("Accept", "application/json")

	consumer := &http.Client{}

	resp, err := consumer.Do(request)

	if err != nil {
		log.Fatalln("Error in consume endpoint")
	}

	body, err := io.ReadAll(resp.Body)

	var JsonBody []RepoName

	err = json.Unmarshal(body, &JsonBody)

	return JsonBody

}
