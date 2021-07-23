# CMD

Main applications for this project.

The directory name for each application should match the name of the executable you want to have (e.g.,** **`/cmd/myapp`).

Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the** **`/pkg`** **directory. If the code is not reusable or if you don't want others to reuse it, put that code in the** **`/internal`** **directory. You'll be surprised what others will do, so be explicit about your intentions!

It's common to have a small** **`main`** **function that imports and invokes the code from the** **`/internal`** **and** **`/pkg`directories and nothing else.

See the** **[`/cmd`](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md)** **directory for examples.

###

Packaging and Continuous Integration.

Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the** **`/build/package`** **directory.

Put your CI (travis, circle, drone) configurations and scripts in the** **`/build/ci`** **directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the** **`/build/ci`** **directory linking them to the location where the CI tools expect them (when possible).
