Only passing builds will be accepted in PRs. 

This means that your changes must keep the project to the standard of linting, whether with the Go part of the application or with the Vue.js part.

The project must also still be able to build both the front-end and back-end.

Notice that in the Vue.js app, the environment files can be edited to suit your needs. Of course you will need an instance of GoScope working, with some example logs in the DB in order to see anything in the SPA.
Change that file (`.env.development`) to match your port, in my case `http://localhost:7005/`.

Any change to assets in the `frontend` folder, obviously building the SPA too, will require a rebuild of the `bindata.go`. 
For this you will require to have the package installed (via `go get -u github.com/shuLhan/go-bindata/...`).

If you make changes to the Vue.js app, navigate in terminal to the root of the project and run (following commands all assume your pwd is the root of project): 
- Build the Vue.js app run `cd frontend && npm run build`
- To create the bindata file, from the `src/utils` directory: `go-bindata -nomemcopy  ../../frontend/dist/...`
- Edit the `bindata.go` file in `src/utils` so that the package is `utils` instead of `main`
