[![Build Status](https://travis-ci.com/spremraj84/travisci-golang.svg?branch=master)](https://travis-ci.com/spremraj84/travisci-golang)
## Repo to test all functionalities of github and Travis CI
1. In settings for this repo protect the master branch
2. Enable - Require pull request reviews before merging
3. Require reviews from code owners
4. Make sure to enable the same permission for administrators

## CI/CD Pipeline
1. Create .travis.yml file
2. Create the build script that will be executed during CI
3. Make sure the build script is set as executable "git update-index --add --chmod=+x build.sh", else build will fail with permission denied
