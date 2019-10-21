package main

import(
	"os"
	"net/http"
	"path/filepath"
	"fmt"
);

func main(){

	var pathArg = "/var/www/";
	var portArg = "3000";


	if (len(os.Args) > 2){

		pathArg = os.Args[1];
		portArg = os.Args[2];

	} else if (len(os.Args) > 1){

		exec, err := os.Executable();

		if(err != nil){
			panic(err);
		}

		pathArg = filepath.Dir(exec);
		portArg = os.Args[1];

	}

	fmt.Println("Serving: "+ pathArg);
	
	fs := http.FileServer(http.Dir(pathArg));
	http.Handle("/", fs);

	http.ListenAndServe(":" + portArg, nil);

}
