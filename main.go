package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Stu struct {
	Person
}

type QueryServerResponse struct {
	Code          int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                  // 错误码
	ServerId      uint64 `protobuf:"varint,2,opt,name=serverId,proto3" json:"serverId,omitempty"`          //      int    `json:"serverId"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                   //string `json:"serverName"`
	Ip            string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`                       //string `json:"ip"`
	LaunchTime    string `protobuf:"bytes,5,opt,name=launchTime,proto3" json:"launchTime,omitempty"`       //string `json:"launchTime"`
	FirstOpenTime string `protobuf:"bytes,6,opt,name=firstOpenTime,proto3" json:"firstOpenTime,omitempty"` //string `json:"firstOpenTime"`
	Version       string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`             //string `json:"version"`
	Git           string `protobuf:"bytes,8,opt,name=git,proto3" json:"git,omitempty"`                     //string `json:"commitId"`
	Online        string `protobuf:"bytes,9,opt,name=online,proto3" json:"online,omitempty"`               //string `json:"online"`
	Register      string `protobuf:"bytes,10,opt,name=register,proto3" json:"register,omitempty"`          //string `json:"register"`
	Cluster       string `protobuf:"bytes,11,opt,name=cluster,proto3" json:"cluster,omitempty"`            //string `json:"cluster"`
	OpGame        string `protobuf:"bytes,12,opt,name=opGame,proto3" json:"opGame,omitempty"`              //string `json:"op_game"`
	MergeServer   uint64 `protobuf:"varint,13,opt,name=mergeServer,proto3" json:"mergeServer,omitempty"`   //int    `json:"mergeServer"`
}

type t1 string
type t2 string

func main() {
	fmt.Println(456)
}

// quicksort
