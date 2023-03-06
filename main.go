package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xfnv1alpha1 "github.com/crossplane/crossplane/apis/apiextensions/fn/io/v1alpha1"
	pgsqlv1 "github.com/vshn/component-appcat/apis/vshn/v1"
	"gopkg.in/yaml.v3"
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

	psql := &pgsqlv1.VSHNPostgreSQL{}

	err = json.Unmarshal(funcIO.Observed.Composite.Resource.Raw, psql)
	if err != nil {
		log.Fatal(err, "FROM JSON.UNMARSHAL")
	}

	b1, _ := yaml.Marshal(funcIO)
	b2, _ := yaml.Marshal(psql)

	// /*
	// 	main login goes here
	// 	manipulate m object
	funcIO.Results = append(funcIO.Results,
		xfnv1alpha1.Result{
			Severity: xfnv1alpha1.SeverityWarning,
			Message:  string(b1),
		},
	)
	// */

	funcIO.Desired.Composite.Resource.Raw = b2

	d1, err := yaml.Marshal(funcIO)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(d1))
}
