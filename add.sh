for i in "$@"; do
    echo 'add domain '$i''
    mkdir -p $i/delivery $i/repository $i/usecase
    touch $i/repository.go $i/usecase.go $i/usecase/usecase.go $i/usecase/usecase_test.go
    echo 'package '$i'' > $i/repository.go
    echo 'package '$i'' > $i/usecase.go
    echo 'package '$i'' > $i/usecase/usecase.go
    echo 'package '$i'' > $i/usecase/usecase_test.go
done