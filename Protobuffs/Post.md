# Practical Golang: Using Protobuffers.

## Introduction

Most apps we make need a means of communication. We usually use *JSON*, or just plain text. *JSON* has got especially popular because of the rise of *Node.js*. The truth though, is, that JSON isn't really a fast format. The marshaller in Go also isn't that fast. That's why in this article we'll learn how to use [google protocol buffers][1]. They are in fact very easy to use, and are much faster than JSON.

Regarding the performance gains, here they are, according to [this benchmark][2]:

| benchmark										| iter		| time/iter	| bytes alloc | allocs			 |
|------------------------------|---------|------------|-------------|--------------|
| BenchmarkJsonMarshal-8			 | 500000	| 3714 ns/op | 1232 B/op	 | 10 allocs/op |
| BenchmarkJsonUnmarshal-8		 | 500000	| 4125 ns/op | 416 B/op		| 7 allocs/op	|
| BenchmarkProtobufMarshal-8	 | 1000000 | 1554 ns/op | 200 B/op		| 7 allocs/op	|
| BenchmarkProtobufUnmarshal-8 | 1000000 | 1055 ns/op | 192 B/op		| 10 allocs/op |

Ok, now let's set up the environment.

## Setup

First we'll need to get the *protobuffer compiler* binaries from here:
https://github.com/google/protobuf/releases/tag/v3.0.0-beta-3
Unpack them somewhere in your **PATH**.

The next step is to get the golang plugin. Make sure that **GOPATH**/bin is in your **PATH**.
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Writing .proto files

Now it's time to define our structure we'll use. I'll create mine in my project root. I'll call it *clientStructure.proto*.

First we need to define the version of *protobuffers* we will use. Here we will use the newest - **proto3**. We'll also define the package of the file. This will also be our go package name of the generated file.

```proto
syntax = "proto3";
package main;
```

Ok, now we'll define our main structure in the file. The *Client* structure:

```proto
message Client {

}
```

Now it's time to define the available fields. Fields are refered to by id, so for each field we define the type, name and id like this:

```proto
type name = id;
```

We'll start with our first 4 fields of our client:

```proto
message Client {
		int32 id = 1;
		string name = 2;
		string email = 3;
		string country = 4;
}
```

We will also define an inner structure *Mail*:

```proto
		string country = 4;

		message Mail {
				string remoteEmail = 1;
				string body = 2;
		}
```

and finally define the inbox field. It's an array of mails, which we create using the *repeated* keyword:

```proto
		message Mail {
				string remoteEmail = 1;
				string body = 2;
		}

		repeated Mail inbox = 5;
}
```

***

Now let's compile it!
Open the directory with the protofile, and launch:
```
protoc --go_out=. clientStructure.proto
```

This will create a generated **GO** file with our *Client* structure.

My file looks like this:
```go
// Code generated by protoc-gen-go.
// source: clientStructure.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	clientStructure.proto

It has these top-level messages:
	Client
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Client struct {
	Id			int32					`protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name		string				 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email	 string				 `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Country string				 `protobuf:"bytes,4,opt,name=country" json:"country,omitempty"`
	Inbox	 []*Client_Mail `protobuf:"bytes,5,rep,name=inbox" json:"inbox,omitempty"`
}

func (m *Client) Reset()										{ *m = Client{} }
func (m *Client) String() string						{ return proto.CompactTextString(m) }
func (*Client) ProtoMessage()							 {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Client) GetInbox() []*Client_Mail {
	if m != nil {
		return m.Inbox
	}
	return nil
}

type Client_Mail struct {
	RemoteEmail string `protobuf:"bytes,1,opt,name=remoteEmail" json:"remoteEmail,omitempty"`
	Body				string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Client_Mail) Reset()										{ *m = Client_Mail{} }
func (m *Client_Mail) String() string						{ return proto.CompactTextString(m) }
func (*Client_Mail) ProtoMessage()							 {}
func (*Client_Mail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func init() {
	proto.RegisterType((*Client)(nil), "main.Client")
	proto.RegisterType((*Client_Mail)(nil), "main.Client.Mail")
}

var fileDescriptor0 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0x09, 0x2e, 0x29, 0x2a, 0x4d, 0x2e, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x3a, 0xcc, 0xc8, 0xc5, 0xe6, 0x0c, 0x96, 0x17,
	0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb2, 0x84, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x80, 0x22, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x08,
	0x17, 0x6b, 0x2a, 0x50, 0x5f, 0x8e, 0x04, 0x33, 0x58, 0x10, 0xc2, 0x11, 0x92, 0xe0, 0x62, 0x4f,
	0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x94, 0x60, 0x01, 0x8b, 0xc3, 0xb8, 0x42, 0xea, 0x5c, 0xac,
	0x99, 0x79, 0x49, 0xf9, 0x15, 0x12, 0xac, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x82, 0x7a, 0x20, 0x4b,
	0xf5, 0x20, 0x16, 0xea, 0xf9, 0x02, 0xf5, 0x06, 0x41, 0xe4, 0xa5, 0x6c, 0xb8, 0x58, 0x40, 0x5c,
	0x21, 0x05, 0x2e, 0xee, 0xa2, 0xd4, 0xdc, 0xfc, 0x92, 0x54, 0x57, 0xb0, 0x35, 0x8c, 0x60, 0xe3,
	0x90, 0x85, 0x40, 0xce, 0x4a, 0xca, 0x4f, 0xa9, 0x84, 0x39, 0x0b, 0xc4, 0x4e, 0x62, 0x03, 0x7b,
	0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x42, 0x16, 0x7b, 0xeb, 0x00, 0x00, 0x00,
}

```

## Using Protocol Buffers in Go

### The Server

Let's first create the server. We will just receive a protobuf in a **POST** body, and print the contents.

First the basic structure and the imports:

```go
package main

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	http.ListenAndServe(":3000", nil)
}
```

So, let's start with a *Client* structure to fill, and read the *body*.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myClient := Client{}

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println(err)
		}
	})
```

We'll *unmarshall* the data, passing in a *reference* to the *Client* to fill in, and check for errors.

```go
		if err != nil {
			fmt.Println(err)
		}

		if err := proto.Unmarshal(data, &myClient); err != nil {
			fmt.Println(err)
		}
```

and finally print it all

```go
		if err := proto.Unmarshal(data, &myClient); err != nil {
			fmt.Println(err)
		}

		println(myClient.Id, ":", myClient.Name, ":", myClient.Email, ":", myClient.Country)

		for _, mail := range myClient.Inbox {
			fmt.Println(mail.RemoteEmail, ":", mail.Body)
		}
	})
```

Now let's start with...

## The Client

The Client will just fill in the *Client* structure, and send it to the server.

The structure:
```go
package main

import (
	"github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"bytes"
)

func main() {
}
```

We create the *Client* structure and fill it in. For the purposes of this article we'll use, the oh so creatively named, John Doe.

```go
func main() {
	myClient := Client{Id: 526, Name: "John Doe", Email: "johndoe@example.com", Country: "US"}
	clientInbox := make([]*Client_Mail, 0, 20)
	clientInbox = append(clientInbox,
		&Client_Mail{RemoteEmail: "jannetdoe@example.com", Body: "Hello. Greetings. Bye."},
		&Client_Mail{RemoteEmail: "WilburDoe@example.com", Body: "Bye, Greetings, hello."})

	myClient.Inbox = clientInbox
}
```

We'll *marshall* the John (the *Client*) into raw data and finally send him to the server.

```go
	myClient.Inbox = clientInbox

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}
```

Note that we used bytes.NewBuffer, so our raw data satisfies the *Reader* requirement for the request body.

## Conclusion

As you can see, protobuffs are really easy to use and provide an actual speed boost in your application. Hope you'll try to use them instead of JSON or other forms of transport in your next project. You can get more information about the more advanced functionalities here: https://developers.google.com/protocol-buffers/docs/gotutorial

Happy coding!




[1]: https://github.com/google/protobuf
[2]: https://github.com/alecthomas/go_serialization_benchmarks