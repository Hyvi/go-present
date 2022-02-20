package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	pb "github.com/Hyvi/go-present/protobuf/out"
	"google.golang.org/protobuf/proto"
)

func promptForAddress(r io.Reader) (*pb.Person, error) {
	// A protocol buffer can be created like any struct.
	p := &pb.Person{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID number: ")
	// An int32 field in the .proto file is represented as an int32 field
	// in the generated Go struct.
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name: ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	// A string field in the .proto file results in a string field in Go.
	// We trim the whitespace because rd.ReadString includes the trailing
	// newline character in its output.
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}
		// The PhoneNumber message type is nested within the Person
		// message in the .proto file.  This results in a Go struct
		// named using the name of the parent prefixed to the name of
		// the nested message.  Just as with pb.Person, it can be
		// created like any other struct.
		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}

		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		ptype = strings.TrimSpace(ptype)

		// A proto enum results in a Go constant for each enum value.
		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q.  Using default.\n", ptype)
		}

		// A repeated proto field maps to a slice field in Go.  We can
		// append to it like any other slice.
		p.Phones = append(p.Phones, pn)
	}

	return p, nil
}

// Main reads the entire address book from a file, adds one person based on
// user input, then writes it back out to the same file.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found.  Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	book := &pb.AddressBook{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	// Add an address.
	addr, err := promptForAddress(os.Stdin)
	if err != nil {
		log.Fatalln("Error with address:", err)
	}
	book.People = append(book.People, addr)
	// [END_EXCLUDE]

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	// 通过命令行解析 hex 之后的 buf 文件
	// echo 0a220a04656c6f6e107b1a0e31323340676c696573652e636f6d22080a063132333435360a250a05656c6f6e32107c1a0e32333440676c696573652e636f6d220a0a0632333435363710010a2a0a05656c6f6e3210d9021a0b33343540646a692e636f6d22060a043132333522090a0533333333331002 | xxd -r -p | protoc --decode tutorial.AddressBook addressbook.proto
	// fmt.Println(base64.StdEncoding.EncodeToString(out))
	fmt.Println(hex.EncodeToString(out))
	// 通过命令行解析 base64 之后的protobuf 文件
	//  echo CiIKBGVsb24QexoOMTIzQGdsaWVzZS5jb20iCAoGMTIzNDU2CiUKBWVsb24yEHwaDjIzNEBnbGllc2UuY29tIgoKBjIzNDU2NxABCioKBWVsb24yENkCGgszNDVAZGppLmNvbSIGCgQxMjM1IgkKBTMzMzMzEAI= | base64 --decode  |   protoc --decode tutorial.AddressBook addressbook.proto
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	// [END marshal_proto]
}
