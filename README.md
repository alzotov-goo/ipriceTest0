### Scripts to work with TeamCity on Go(lang)  
* demonstartes Go working with http connection  
* TeamCity REST API used for communication  
* Essentially bash is enough to get required info  
* However to make it look more like a "script on programming language" I've done some job with Go language  

build: 
    go build *.go 

help
    *.exe --help

*builds* script:  
get successful builds in all branches since 2020/01 (by default)

   builds -url=teamcitybaseurl -login=test@user:password  

*changes* script:  
list of changes for sertain build (more details of every change is TODO if needed)

   changes -url=teamcitybaseurl -login=test@user:password -buildId=12345

buildId can be obtained from *builds* output

