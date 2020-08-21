Only passing builds will be accepted in PRs. 

This means that your changes must keep the project to the standard of linting, whether with golangci-lint for the Go part of the application or with ng lint for the Angular app.

The project must also still be able to build both the front-end and back-end.

Notice that in the Angular app, the environment files can be edited to suit your needs. Of course you will need an instance of GoScope working, with some example logs in the DB in order to see anything in the SPA.
Change that file to match your port, in my case `http://localhost:7004/`.

Any change to assets in the `frontend` folder, obviously building the SPA too, will require a rebuild of the `bindata.go`. 
For this you will require to have the package installed (via `go get -u github.com/shuLhan/go-bindata/...`).

If you make changes to the Angular app, navigate in terminal to the root of the project and run (following commands all assume your pwd is the root of project): 
- Build the Angular app run `cd frontend && npm run build`
- To create the bindata file: `cd ../goscope && go-bindata -nomemcopy  ../frontend/dist/...`
- Edit the `bindata.go` file in `/goscope` so that the package is `goscope` instead of `main`
