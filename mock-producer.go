// =======================================================================================
// mock-producer: This program pretends to read a data source and send the data to a queue.
// @author Flávio Gomes da Silva Lisboa <flavio.lisboa@fgsl.eti.br>
// @license LGPL-2.1
// =======================================================================================
package main

import (
    "fmt"
    "os"    
    "time"

    stomp "github.com/go-stomp/stomp"
)

func main() {
    fmt.Println("MOCK-PRODUCER: initializing data producing...")
    fmt.Println("MOCK-PRODUCER: version 1.0.0")

    time.Sleep(30 * time.Second)

    // continue even some error occurs
    for {
        listEvents()
        time.Sleep(10 * time.Second)
    }
}

func listEvents() {
    var data string
    var conn *stomp.Conn
    var err error

    data = getDataForLog()

    activemqHost := "mock-queue"
    if os.Getenv("QUEUE_HOST") != "" {
        activemqHost = os.Getenv("QUEUE_HOST")
    }
    activemqPort := "61616"
    if os.Getenv("QUEUE_PORT") != "" {
        activemqPort = os.Getenv("QUEUE_PORT")
    }
    activemqServer := activemqHost + ":" + activemqPort
    conn, err = stomp.Dial("tcp", activemqServer,stomp.ConnOpt.Login(os.Getenv("QUEUE_USERNAME"), os.Getenv("QUEUE_PASSWORD")))
    if err != nil {
        fmt.Println("ERROR: stomp.Dial: " + err.Error())
    } else {
        fmt.Println("MOCK-PRODUCER: Has connectivity with " + activemqServer)
    }
    err = conn.Send(
        "/queue/pods",// destination
        "application/json",// content-type
        []byte(data))// body
    if err != nil {
        fmt.Println("ERROR WHEN SENDING TO QUEUE " + err.Error())
        conn.Disconnect()
    }

    msg := data

    fmt.Println(msg)
}

func getDataForLog() string {
    var data string

    now := time.Now()
    data = "{" +
        "\"class\": \"audit\"," +
        "\"subclass\": \"pod_ip\"," +
        "\"origin\":\"mock\"," +
        "\"dc\":\"mock\"," +
        "\"host\":\"mock\"," +
        "\"pod_namespace\":\"mock\"," +
        "\"pod\":\"mock\"," +
        "\"pod_ip\":\"mock\"," +
        "\"pod_ip_status\":\"mock\"," +
        "\"short_message\":\"mock/mock:mock\"," +
        "\"full_message\":\"mock/mock:mock:mock:" + now.String() + "\"," +
        "\"timestamp\":\"" +     now.String() + "\"," +
        "\"logtype\": \"mock\"}";
    return data;
}
