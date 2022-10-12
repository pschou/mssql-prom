module testproj

go 1.16

require (
	github.com/bet365/go-mssqldb-auth-krb5 v0.0.0-20220816165549-b8d5174b1ce8
	github.com/denisenkom/go-mssqldb v0.12.0
	github.com/hashicorp/go-uuid v1.0.3
	github.com/jcmturner/gofork v1.7.6
	github.com/jcmturner/gokrb5/v8 v8.4.3
	github.com/microsoft/go-mssqldb v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	gopkg.in/jcmturner/aescts.v1 v1.0.1
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1
	gopkg.in/jcmturner/goidentity.v3 v3.0.0
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0
	gopkg.in/jcmturner/rpc.v1 v1.1.0
)

replace github.com/microsoft/go-mssqldb => github.com/bet365/go-mssqldb v0.12.1-0.20220816164333-86ae3f6abb64
