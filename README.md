# Repo Initer

Experimentation for a repo init-er for kubernetes

Steps:

Process originally proposed here: https://groups.google.com/d/msg/kubernetes-dev/o6E1u-orDK8/SAqal_CeCgAJ

## How it works.

This script creates:

* A **LICENSE** file, which by default uses the Apache 2.0 format.
* A **code-of-conduct.md** file (which defers to the Kubernetes Code of
  Conduct.)
* A **CONTRIBUTING.md** file, which specifies the repo's policies for contributing.
* An **OWNERS.yaml** YAML file, which allows for automatic routing.

* A **RELEASE** file TODO: (what goes here?)
* A **README.md**

Note: If the files exist, this tool will not create them.
