package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xfnv1alpha1 "github.com/crossplane/crossplane/apis/apiextensions/fn/io/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func main() {

	funcIO := xfnv1alpha1.FunctionIO{}
	//composite := pgsqlv1.VSHNPostgreSQL{}
	//m := make(map[interface{}]interface{})
	// detect if os.Stdin is not empty
	x, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(x, &funcIO)
	if err != nil {
		log.Fatal(err)
	}

	object := `
apiVersion: kubernetes.crossplane.io/v1alpha1
kind: Object
metadata:
  name: testingconfigmaps
spec:
  providerConfigRef:
    name: kubernetes
  forProvider:
    manifest:
      apiVersion: v1
      kind: ConfigMap
      metadata:
        name: url-config
        namespace: default
      data:
        fullURL: "https://google.pl"`

	k8sapproved, err := yaml.YAMLToJSON([]byte(object))
	if err != nil {
		log.Fatal("from k8sapproved", err)
	}

	funcIO.Desired.Composite.Resource.Raw = funcIO.Observed.Composite.Resource.Raw

	funcIO.Desired.Resources = append(funcIO.Desired.Resources, xfnv1alpha1.DesiredResource{
		Name: "examplename",
		Resource: runtime.RawExtension{
			Raw: k8sapproved,
		},
	},
	)

	funcIO.Results = append(funcIO.Results,
		xfnv1alpha1.Result{
			Severity: xfnv1alpha1.SeverityNormal,
			Message:  fmt.Sprintf("\n\n\n%s\n\n\n", string(x)),
		},
	)

	d1, err := yaml.Marshal(funcIO)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(d1))
}
