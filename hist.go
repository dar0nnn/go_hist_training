package main

import (
	"context"
	"fmt"
	mongoutils "mongo_hist/mongoutils"
	"reflect"
)

func main() {
	client, err := mongoutils.Create_client_with_conn()
	if err != nil {
		fmt.Println("fuck")
	}
	fmt.Println(mongoutils.Check_conn(client))
	data := mongoutils.Get_data_for_hist(client, "5e3869f0294b4d34a1043e65")
	fmt.Println(len(data))
	fmt.Println(reflect.TypeOf(data))
	err = client.Disconnect(context.TODO())

}
