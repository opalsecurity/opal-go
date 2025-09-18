module github.com/opalsecurity/opal-go

go 1.18

require (
	gopkg.in/validator.v2 v2.0.1
)

retract (
	v1.1.1 // cannot be installed on macos
)
