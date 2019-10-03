package services

import (
	// "flag"

	"encoding/json"
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
//StartBattlefield creates a battlefield with given name in namespace from yaml file
func StartBattlefield( battlefieldname string, namespace string, yamlname string, client *rest.RESTClient) ([]byte){
	log.Println("startBattlefield is called")

	dat, err := ioutil.ReadFile("resource/" + yamlname + ".yaml")
	if err != nil {
        panic(err)
    }

	//Read from yaml file
	battlefield := &v1alpha1.Battlefield{}
	yaml.Unmarshal(dat,battlefield)
	battlefield.ObjectMeta.Name = battlefieldname

	// Send to api
	result := &v1alpha1.Battlefield{}
	client.Post().
		Namespace(namespace).
		Resource("battlefields").
		Body(battlefield).
		Do().
      Into(result)

	  json, _ := json.Marshal(result)

	return json
	// yaml.
}