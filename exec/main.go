package main

import (
        "../libexec/access_mysql"
)

func main() {
        db := access_mysql.Connect()
        defer db.Close()
}
