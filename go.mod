module github.com/onsd/Equipment-Management-2

// + heroku goVersion go1.11
// + heroku install ./cmd/... ./vendor
require (
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7
	github.com/gin-gonic/gin v0.0.0-20170702092826-d459835d2b07
	github.com/go-sql-driver/mysql v1.4.0
	github.com/golang/protobuf v0.0.0-20170601230230-5a0f697c9ed9
	github.com/jinzhu/gorm v1.9.1
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a
	github.com/mattn/go-isatty v0.0.0-20170307163044-57fdcb988a5c
	github.com/onsd/Equipment-Management v0.0.0-20181117070445-344ec1574827 // indirect
	github.com/ugorji/go v0.0.0-20170215201144-c88ee250d022
	golang.org/x/sys v0.0.0-20180810173357-98c5dad5d1a0
	google.golang.org/appengine v1.1.0
	gopkg.in/go-playground/validator.v8 v8.18.1
	gopkg.in/yaml.v2 v2.0.0-20160928153709-a5b47d31c556
)
