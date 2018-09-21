/*
 Copyright (c) 2018, SFEIR OSPO <ospo@sfeir.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package image

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"

	"github.com/nfnt/resize"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

const (
	scope = storage.DevstorageFullControlScope
)

func getStorageData(bucketName string, objectName string) ([]byte, string, error) {
	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		return nil, "", err
	}

	service, err := storage.New(client)
	if err != nil {
		return nil, "", err
	}

	obj, err := service.Objects.Get(bucketName, objectName).Do()
	if err != nil {
		return nil, "", err
	}

	response, err := client.Get(obj.MediaLink)
	if err != nil {
		return nil, "", err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	return contents, obj.ContentType, nil
}

func ResizeStorageImage(bucketName string, objectName string, width int) ([]byte, error) {
	var err error
	data, typ, err := getStorageData(bucketName, objectName)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	var img image.Image
	if typ == "image/png" {
		img, err = png.Decode(reader)
	} else if typ == "image/jpeg" || typ == "image/jpg" {
		img, err = jpeg.Decode(reader)
	} else if typ == "image/gif" {
		img, err = gif.Decode(reader)
	} else {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	resized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if typ == "image/png" {
		err = png.Encode(buf, resized)
	} else if typ == "image/jpeg" || typ == "image/jpg" {
		err = jpeg.Encode(buf, resized, nil)
	} else if typ == "image/gif" {
		err = gif.Encode(buf, resized, nil)
	}
	if err != nil {
		return nil, err
	}
	contents := buf.Bytes()

	return contents, nil
}
