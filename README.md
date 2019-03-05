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
4. Tag the build for releases with the build command go build -o testproject -i -v -ldflags="-X main.version=$(git describe --always --long)" .

## Deployment
1. Trying github releases because we can create a draft release build (stage) and production releases
2. Install travis cli https://github.com/travis-ci/travis.rb#installation
3. Configure the cli client to generate OAUTH API keys by running the command 'travis setup releases'
travis setup releases --com <br/>
Username: spremraj84 <br/>
Password for spremraj84: <br/>
File to Upload: (outfput-file) <br/>
Deploy only from spremraj84/travisci-golang? |yes| yes <br/>
Encrypt API key? |yes| yes <br/>
4. The key will be automatically setup in .travis.yml file <br/>
