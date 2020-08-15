Notice that in the Angular app, the environment files can be edited to suit your needs. Of course you will need an instance of GoScope working, with some example logs in the DB in order to see anything in the SPA.
Change that file to match your port, in my case `http://localhost:7004/`.
Any change to assets in the `static` folder, obviously building the SPA too, will require a rebuild of the `bindata.go`. 
For this you will require to have the package installed (via `go get -u github.com/shuLhan/go-bindata/...`).

Then navigate in terminal to the root of the project and run (following commands all assume your pwd is the root of project): 
- Build the Angular app run `cd static/goscope && ng build --prod --output-hashing none --base-href /goscope/`
- To create the bindata file: `cd ../../goscope && go-bindata -nomemcopy  ../static/goscope/dist/...`
- Edit the `bindata.go` file in `/goscope` so that the package is `goscope` instead of `main`
