// +build windows

package mssql

import (
	"github.com/microsoft/go-mssqldb/integratedauth"
	_ "github.com/microsoft/go-mssqldb/integratedauth/ntlm"
	_ "github.com/microsoft/go-mssqldb/integratedauth/winsspi"
)

func init() {
	integratedauth.DefaultProviderName = "winsspi"
}
