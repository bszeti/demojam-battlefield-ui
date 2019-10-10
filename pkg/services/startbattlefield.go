package services

import (
	"errors"

	"io/ioutil"
	"log"
	"k8s.io/client-go/rest"

	yaml "gopkg.in/yaml.v2"

	"github.com/bszeti/battlefield-ui/pkg/apis/rhte/v1alpha1"
)


//GetBattlefield gets a Battlefield resource by name
func GetBattlefield( battlefieldname string, namespace string, client *rest.RESTClient) (*v1alpha1.Battlefield, error){
	log.Printf("GetBattlefield is called: %s",battlefieldname)

	battlefield := v1alpha1.Battlefield{}
	err := client.Get().
		Namespace(namespace).
		Resource("battlefields").
		Name(battlefieldname).
		Do().
		Into(&battlefield)
	if err!=nil {
		return nil, err
	}
	return &battlefield, nil
}


//StartBattlefield creates a battlefield with given name in namespace from yaml file
func StartBattlefield( battlefieldname string, namespace string, yamlname string, client *rest.RESTClient) (*v1alpha1.Battlefield, error){
	log.Printf("startBattlefield is called: %s %s %s", battlefieldname, namespace, yamlname)

	//Check if exists
	found, err := GetBattlefield(battlefieldname, namespace, client)
	if err==nil {
		//Exists...
		if found.Status.Phase != "done" {
			log.Println("Running...")
			return found, nil
		}
		//Delete if Game Over
		err := DeleteBattlefield(battlefieldname, namespace, client)
		if err!=nil {
			log.Println("Failed to delete battlefield", err)
		}

	}

	//Read from yaml file
	dat, err := ioutil.ReadFile("resource/" + yamlname + ".yaml")
	if err != nil {
		log.Println("Failed to read file", err)
        return nil, err
    }

	battlefield := &v1alpha1.Battlefield{}
	err = yaml.Unmarshal(dat,battlefield)
	if (err != nil) {
		log.Println("Unmarshall error",err)
		return nil, err
	}
	battlefield.ObjectMeta.Name = battlefieldname

	// Send to api
	result := &v1alpha1.Battlefield{}
	err = client.Post().
		Namespace(namespace).
		Resource("battlefields").
		Body(battlefield).
		Do().
	  Into(result)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}


//DeleteBattlefield creates a battlefield with given name in namespace from yaml file
func DeleteBattlefield( battlefieldname string, namespace string, client *rest.RESTClient) (error){
	log.Println("DeleteBattlefield is called")

	err := client.Delete().
		Namespace(namespace).
		Resource("battlefields").
		Name(battlefieldname).
		Do().
		Error()

	return err
}

//DisqualifyPlayer sets disqualify for player 
func DisqualifyPlayer( battlefieldname string, namespace string, playerName string, status bool, client *rest.RESTClient) (error) {
	log.Println("DisqualifyPlayer is called")

	battlefield,err := GetBattlefield(battlefieldname, namespace, client)
	if err != nil {
		log.Println("Failed to get Battlefield.",err)
		return err
	}

	foundPlayer := false
	for index, player := range battlefield.Spec.Players {
		if playerName == player.Name {
			battlefield.Spec.Players[index].Disqualified = status
			foundPlayer = true
		}
	}
	if !foundPlayer {
		return errors.New("Player not found")
	}

	err = client.Put().
		Namespace(namespace).
		Resource("battlefields").
		Name(battlefieldname).
		Body(battlefield).
		Do().
		Error()

	if err != nil {
		log.Println("Failed to update Battlefield.",err)
		return err
	}
	return nil
}

//ShieldPlayer sets shield for player 
func ShieldPlayer( battlefieldname string, namespace string, playerName string, status bool, client *rest.RESTClient) (error) {
	log.Println("ShieldPlayer is called")

	battlefield,err := GetBattlefield(battlefieldname, namespace, client)
	if err != nil {
		log.Println("Failed to get Battlefield.",err)
		return err
	}

	foundPlayer := false
	for index, player := range battlefield.Spec.Players {
		if playerName == player.Name {
			battlefield.Spec.Players[index].Shield = status
			foundPlayer = true
		}
	}
	if !foundPlayer {
		return errors.New("Player not found")
	}

	err = client.Put().
		Namespace(namespace).
		Resource("battlefields").
		Name(battlefieldname).
		Body(battlefield).
		Do().
		Error()

	if err != nil {
		log.Println("Failed to update Battlefield.",err)
		return err
	}
	return nil
}

