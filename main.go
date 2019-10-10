package main

import (
	// "flag"

	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	// "strings"

	"k8s.io/client-go/tools/clientcmd"
	//clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	// "io/ioutil"
	"log"
	"net/http"
		"github.com/gorilla/handlers"


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
	name := vars["name"]
	player := vars["player"]

	status, err := strconv.ParseBool(vars["status"])
	if err != nil {
		log.Println("Wrong status",status)
	} else {
		log.Printf("ShieldHandler request: %s",name)
		err = services.ShieldPlayer(name,namespace,player,status,client)
	}

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(status)
	}

}

//DisqualifyHandler is
func DisqualifyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	player := vars["player"]
	status, err := strconv.ParseBool(vars["status"])
	if err != nil {
		log.Println("Wrong status",status)
	} else {
		log.Printf("DisqualifyHandler request: %s",name)
		err = services.DisqualifyPlayer(name,namespace,player,status,client)
	}

	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(status)
	}
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
	router.HandleFunc("/api/battlefield/{name}", GetBattlefield).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/battlefield/{name}", DeleteBattlefieldWithName).Methods("DELETE")
	router.HandleFunc("/api/start/{name}/{type}", StartBattlefieldWithNameAndType).Methods("GET")
	router.HandleFunc("/api/start/{name}", StartBattlefieldWithName).Methods("GET") //default-type
	router.HandleFunc("/api/start", StartBattlefield).Methods("GET", "OPTIONS") //demofield - default type
	router.HandleFunc("/api/battlefield/{name}/{player}/shield/{status}", ShieldHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/battlefield/{name}/{player}/disqualified/{status}", DisqualifyHandler).Methods("GET", "OPTIONS")


	//Serve static filed at root from "static" directory
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/").Handler(fs)

	// router.Handle("/", fs)

	// staticDir := "/static/"
	// router.
	// 	PathPrefix(staticDir).
	// 	Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	err = http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
