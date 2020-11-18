module datumbrain/my-project

go 1.15

replace (
	datumbrain/my-project/apis => ./cache
	datumbrain/my-project/constants => ./constants
	datumbrain/my-project/log => ./log
	datumbrain/my-project/models => ./models
	datumbrain/my-project/utils => ./utils
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi v1.5.0
	github.com/go-chi/cors v1.1.1
	github.com/go-chi/jwtauth v4.0.4+incompatible
)
