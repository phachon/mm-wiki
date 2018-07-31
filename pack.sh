#!/bin/bash
VER=$1
if [ "$VER" = "" ]; then
    echo 'please input pack version!'
    exit 1
fi
RELEASE="release-${VER}"
rm -rf ${RELEASE}
mkdir ${RELEASE}

# windows amd64
echo 'Start pack windows amd64...'
GOOS=windows GOARCH=amd64 go get ./...
GOOS=windows GOARCH=amd64 go build ./
cd install
GOOS=windows GOARCH=amd64 go build ./
cd ..
tar -czvf "${RELEASE}/mm-wiki-windows-amd64.tar.gz" mm-wiki.exe conf/ docs/ logs/.gitignore static/ views/ install/install.exe LICENSE README.md
rm -rf mm-wiki.exe

echo 'Start pack windows X386...'
GOOS=windows GOARCH=386 go get ./...
GOOS=windows GOARCH=386 go build ./
cd install
GOOS=windows GOARCH=386 go build ./
cd ..
tar -czvf "${RELEASE}/mm-wiki-windows-386.tar.gz" mm-wiki.exe conf/ docs/ logs/.gitignore static/ views/ install/install.exe LICENSE README.md
rm -rf mm-wiki.exe

echo 'Start pack linux amd64'
GOOS=linux GOARCH=amd64 go get ./...
GOOS=linux GOARCH=amd64 go build ./
cd install
GOOS=linux GOARCH=amd64 go build ./
cd ..
tar -czvf "${RELEASE}/mm-wiki-linux-amd64.tar.gz" mm-wiki conf/ docs/ logs/.gitignore static/ views/ install/install LICENSE README.md
rm -rf mm-wiki

echo 'Start pack linux 386'
GOOS=linux GOARCH=386 go get ./...
GOOS=linux GOARCH=386 go build ./
cd install
GOOS=linux GOARCH=386 go build ./
cd ..
tar -czvf "${RELEASE}/mm-wiki-linux-386.tar.gz" mm-wiki conf/ docs/ logs/.gitignore static/ views/ install/install LICENSE README.md
rm -rf mm-wiki

echo 'Start pack mac amd64'
GOOS=darwin GOARCH=amd64 go get ./...
GOOS=darwin GOARCH=amd64 go build ./
cd install
GOOS=darwin GOARCH=amd64 go build ./
cd ..
tar -czvf "${RELEASE}/mm-wiki-mac-amd64.tar.gz" mm-wiki conf/ docs/ logs/.gitignore static/ views/ install/install LICENSE README.md
rm -rf mm-wiki

echo 'END'
