package services

import (
	// "flag"

	//"encoding/json"
	// "fmt"
	// "os"
	// "path/filepath"
	// "strings"

	// "k8s.io/client-go/tools/clientcmd"
	//clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"io/ioutil"
	"log"
	// "net/http"

	// "github.com/gorilla/mux"
	// "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	// "k8s.io/apimachinery/pkg/runtime/serializer"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
