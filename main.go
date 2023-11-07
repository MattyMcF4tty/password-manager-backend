package main

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

func main() {
  supabaseUrl := "https://lfzkmvojietwzqsrphid.supabase.co"
  supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Imxmemttdm9qaWV0d3pxc3JwaGlkIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTkzNTY2NzMsImV4cCI6MjAxNDkzMjY3M30.mF80rqrlnvKJwdA0mq5TDpZihXoy1815RVm-lSceUAM"
  supabase := supa.CreateClient(supabaseUrl, supabaseKey)

  var results map[string]interface{}
  err := supabase.DB.From("Users").Select("*").Single().Execute(&results)
  if err != nil {
    panic(err)
  }

  fmt.Println(results) // Selected rows
}