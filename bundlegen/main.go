package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type imageArg struct {
	DigestPath string `json:"digestPath"`
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type image struct {
	Digest string `json:"contentDigest"`
	Name   string `json:"image"`
	Type   string `json:"imageType"`
}

type bundle struct {
	InvocationImages []image          `json:"invocationImages"`
	Images           map[string]image `json:"images"`
}

func mustConvertJSONToImage(raw string) image {
	var img imageArg
	err := json.Unmarshal([]byte(raw), &img)
	handle(err)

	digest, err := ioutil.ReadFile(img.DigestPath)
	handle(err)

	imageName := fmt.Sprintf("%s/%s:%s", img.Registry, img.Repository, img.Tag)
	return image{string(digest), imageName, "docker"}
}

func main() {
	indent := flag.Bool("indent", true, "indent bundle.json")
	invocationImages := flag.String("invocation-images", "", "bundle.json invocation images")
	images := flag.String("images", "", "bundle.json images")
	bundlePath := flag.String("bundle-path", "", "bundle.json output path")
	flag.Parse()

	bun := bundle{
		InvocationImages: []image{},
		Images:           map[string]image{},
	}
	for _, raw := range strings.Split(*invocationImages, "\n") {
		if raw == "" {
			continue
		}
		bun.InvocationImages = append(bun.InvocationImages, mustConvertJSONToImage(raw))
	}
	for _, raw := range strings.Split(*images, "\n") {
		if raw == "" {
			continue
		}
		elems := strings.Split(raw, "=")
		bun.Images[elems[0]] = mustConvertJSONToImage(elems[1])
	}

	if *indent {
		out, err := json.MarshalIndent(bun, "", "    ")
		handle(err)
		handle(ioutil.WriteFile(*bundlePath, out, os.ModePerm))
		return
	}

	out, err := json.Marshal(bun)
	handle(err)
	handle(ioutil.WriteFile(*bundlePath, out, os.ModePerm))
}

func handle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
