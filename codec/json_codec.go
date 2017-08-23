// Copyright 2015-2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"encoding/json"
	"io"
)

func init() {
	Reg(new(JsonCodec))
}

// JsonCodec json codec
type JsonCodec struct{}

// Name returns codec name
func (j *JsonCodec) Name() string {
	return "json"
}

//
// Id returns codec id
func (j *JsonCodec) Id() byte {
	return 'j'
}

// NewEncoder returns a new json encoder that writes to writer.
func (*JsonCodec) NewEncoder(writer io.Writer) Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder returns a new json decoder that reads from limit reader.
//
// The decoder introduces its own buffering and may
// read data from limitReader beyond the JSON values requested.
func (*JsonCodec) NewDecoder(limitReader io.Reader) Decoder {
	return json.NewDecoder(limitReader)
}