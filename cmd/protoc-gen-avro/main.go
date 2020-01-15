package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/syucream/avro-protobuf/pkg/schema"
)

func main() {
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var req plugin.CodeGeneratorRequest
	if err := proto.Unmarshal(buf, &req); err != nil {
		log.Fatal(err)
	}

	records := []map[string]interface{}{}
	for _, f := range req.GetProtoFile() {
		for _, m := range f.GetMessageType() {
			ns := schema.GetNamespace(proto.MessageName(m))
			record, err := schema.GetRecordSchemaFromDescriptor(f, m, ns)
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)
		}
	}

	resp := plugin.CodeGeneratorResponse{}
	for _, record := range records {
		recordName, ok := record["name"].(string)
		if !ok {
			log.Fatal("missing record name")
		}
		name := recordName + ".avsc"

		data, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		content := string(data)

		resp.File = append(resp.File, &plugin.CodeGeneratorResponse_File{
			Name:    &name,
			Content: &content,
		})
	}

	buf, err = proto.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(buf); err != nil {
		log.Fatal(err)
	}
}
