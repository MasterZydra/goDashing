#!/bin/bash
#https://gist.github.com/bclinkinbeard/1331790
versionLabel=$1
productName="goDashing"
releaseFolder="release/"

mkdir ${releaseFolder} 2>/dev/null

rm ${releaseFolder}*.tgz 2>/dev/null
rm ${releaseFolder}*.zip 2>/dev/null

rice embed-go

arch=amd64
os=windows
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o ${product}.exe ./cmd/godashing/...
zip -r ${product}.zip ${product}.exe
rm $product.exe
mv $product.zip ${releaseFolder}/

arch=amd64
os=darwin
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/godashing/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/

arch=arm
os=linux
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/godashing/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/

arch=amd64
os=linux
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/godashing/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/

arch=386
os=linux
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o $product ./cmd/godashing/...
tar czfv $product.tgz $product
rm $product
mv $product.tgz ${releaseFolder}/

arch=386
os=windows
product=${productName}_${versionLabel}_${os}_${arch}
env GOOS=$os GOARCH=$arch go build -ldflags="-s -w -X main.version=${1} -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` " -o ${product}.exe ./cmd/godashing/...
zip -r ${product}.zip ${product}.exe
rm $product.exe
mv $product.zip ${releaseFolder}/
