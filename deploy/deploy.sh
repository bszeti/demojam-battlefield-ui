oc delete project battlefield-ui
oc new-project --skip-config-write=true battlefield-ui
oc adm policy add-cluster-role-to-user cluster-admin -z default -n battlefield-ui
oc create -f deployment.yaml -n battlefield-ui
oc create -f service.yaml -n battlefield-ui
oc expose service battlefield-ui -n battlefield-ui


