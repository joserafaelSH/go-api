package service


type ServiceInterface[I any, O any] interface {
    Execute(input I) O
}

