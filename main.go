package main

import (
	// "flag"

	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	// "strings"

	"k8s.io/client-go/tools/clientcmd"
	//clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// yaml "gopkg.in/yaml.v2"

	"github.com/bszeti/battlefield-ui/pkg/apis/rhte/v1alpha1"
	"github.com/bszeti/battlefield-ui/pkg/services"

	
)

var namespace= os.Getenv("NAMESPACE")

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\": \"OK\"}")
}

//GetBattlefield gets a Battlefield resource by name
func GetBattlefield(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	log.Printf("GetBattlefield request: %s",name)

	battlefield, err :=  services.GetBattlefield(name,namespace,client)
	if err!=nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(battlefield) 
	}
}

//StartBattlefield is
func StartBattlefield(w http.ResponseWriter, r *http.Request) {
	name := "demofield"
	_, err :=  services.StartBattlefield(name,namespace,"default",client)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(name)
	}
}

//StartBattlefieldWithName is
func StartBattlefieldWithName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]


	_, err := services.StartBattlefield(name,namespace,"default",client)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(name)
	}
}

//StartBattlefieldWithNameAndType is
func StartBattlefieldWithNameAndType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	yamlname := vars["type"]

	_, err := services.StartBattlefield(name,namespace,yamlname,client)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(name)
	}
}

//DeleteBattlefieldWithName is
func DeleteBattlefieldWithName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log.Printf("DeleteBattlefieldWithName request: %s",name)
	err := services.DeleteBattlefield(name,namespace,client)

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	
}

//ShieldHandler is
func ShieldHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	status := vars["status"]
	json.NewEncoder(w).Encode("ShieldHandler: " + player + " " + status)
}

//DisqualifyHandler is
func DisqualifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player := vars["player"]
	status := vars["status"]
	json.NewEncoder(w).Encode("DisqualifyHandler: " + player + " " + status)
}

var client *rest.RESTClient

func main() {
	fmt.Println("Monitoring namespace:",namespace)
	if namespace=="" {
		panic("NAMESPACE env var is required.")
	}

	var config *rest.Config
	var err error

	log.Printf("Creating K8s client")
	log.Printf("try using in-cluster configuration")
	config, err = rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(
			os.Getenv("HOME"), ".kube", "config",
			)
		log.Printf("try using kubeconfig: %s",kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	if config == nil {
		panic(err)
	}

	v1alpha1.SchemeBuilder.AddToScheme(scheme.Scheme)

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &v1alpha1.SchemeGroupVersion
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err = rest.RESTClientFor(&crdConfig)
     if err != nil {
		panic(err)
	}


	log.Printf("Creating server")

	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/api/battlefield/{name}", GetBattlefield).Methods("GET")
	router.HandleFunc("/api/battlefield/{name}", DeleteBattlefieldWithName).Methods("DELETE")
	router.HandleFunc("/api/start/{name}/{type}", StartBattlefieldWithNameAndType).Methods("GET")
	router.HandleFunc("/api/start/{name}", StartBattlefieldWithNameAndType).Methods("GET") //default-type
	router.HandleFunc("/api/start", StartBattlefield).Methods("GET") //demofield - default type
	router.HandleFunc("/api/battlefield/{name}/{player}/shield/{status}", ShieldHandler).Methods("GET")
	router.HandleFunc("/api/battlefield/{name}/{player}/disqualify/{status}", DisqualifyHandler).Methods("GET")


	//Serve static filed at root from "static" directory
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/").Handler(fs)

	// router.Handle("/", fs)

	// staticDir := "/static/"
	// router.
	// 	PathPrefix(staticDir).
	// 	Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
