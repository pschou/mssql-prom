package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	// Import the Azure AD driver module (also imports the regular driver package)
	"github.com/denisenkom/go-mssqldb/azuread"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/service"
)

func ConnectWithMSI() (*sql.DB, error) {
	return sql.Open(azuread.DriverName, "sqlserver://azuresql.database.windows.net?database=yourdb&fedauth=ActiveDirectoryMSI")

	cfg, err := config.Load("/path/to/config/file")

	config.Load("/path/to/config/file")

	keytab.Load("/path/to/file.keytab")
	keytab.Parse(b)

	cl := client.NewWithPassword("username", "REALM.COM", "password", cfg)
	cl := client.NewWithKeytab("username", "REALM.COM", kt, cfg)

	s := service.NewSettings(&kt) // kt is a keytab and optional settings can also be provided.
	if ok, creds, err := service.VerifyAPREQ(&APReq, s); ok {
		// Perform application specific actions
		// creds object has details about the client identity
	}
}

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "", "the database server")
	user          = flag.String("user", "", "the database user")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
}
