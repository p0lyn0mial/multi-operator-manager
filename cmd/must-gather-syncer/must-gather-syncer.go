package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/yaml"
)

func main() {
	mustGatherDir := "/Users/lszaszki/go/src/github.com/openshift/multi-operator-manager/must-gather"
	sampleOperatorOutputDir := "/Users/lszaszki/go/src/github.com/openshift/multi-operator-manager/sample-operator-output"

	sourceConfigMapToCreatePath := "UserWorkload/Create/namespaces/openshift-authentication/core/configmaps/001-body-foo.yaml"
	sourceConfigMapToUpdatePath := "UserWorkload/Update/namespaces/openshift-authentication/core/configmaps/001-body-foo.yaml"

	destinationConfigMapPath := "namespaces/openshift-authentication/core/configmaps.yaml"

	configMapList, err := readListFile(os.DirFS(mustGatherDir), destinationConfigMapPath)
	handleErr(err)

	demoConfigMapName := "foo"
	demoConfigMapFound := false
	newConfigMapList := []unstructured.Unstructured{}
	for _, obj := range configMapList.Items {
		if obj.GetName() != demoConfigMapName {
			newConfigMapList = append(newConfigMapList, obj)
			continue
		}

		demoConfigMap, err := readIndividualFile(os.DirFS(sampleOperatorOutputDir), sourceConfigMapToUpdatePath)
		handleErr(err)
		demoConfigMapFound = true
		newConfigMapList = append(newConfigMapList, *demoConfigMap)
	}
	if !demoConfigMapFound {
		demoConfigMap, err := readIndividualFile(os.DirFS(sampleOperatorOutputDir), sourceConfigMapToCreatePath)
		handleErr(err)

		newConfigMapList = append(newConfigMapList, *demoConfigMap)
	}

	configMapList.Items = newConfigMapList
	err = writeListFile(mustGatherDir, destinationConfigMapPath, configMapList)
	handleErr(err)
	fmt.Println("done syncing")
}

var localScheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(localScheme)

func writeListFile(destinationDir string, filePath string, obj runtime.Object) error {
	objYAMLBytes, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	destinationFilePath := filepath.Join(destinationDir, filePath)
	return os.WriteFile(destinationFilePath, objYAMLBytes, 0644)
}

func readIndividualFile(sourceFS fs.FS, path string) (*unstructured.Unstructured, error) {
	content, err := fs.ReadFile(sourceFS, path)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q: %w", path, err)
	}

	return decodeIndividualObj(content)
}

func readListFile(sourceFS fs.FS, path string) (*unstructured.UnstructuredList, error) {
	content, err := fs.ReadFile(sourceFS, path)
	if err != nil {
		return nil, fmt.Errorf("unable to read %q: %w", path, err)
	}

	return decodeListObj(content)
}

func decodeListObj(content []byte) (*unstructured.UnstructuredList, error) {
	obj, _, err := codecs.UniversalDecoder().Decode(content, nil, &unstructured.UnstructuredList{})
	if err != nil {
		return nil, fmt.Errorf("unable to decode: %w", err)
	}
	return obj.(*unstructured.UnstructuredList), nil
}

func decodeIndividualObj(content []byte) (*unstructured.Unstructured, error) {
	obj, _, err := codecs.UniversalDecoder().Decode(content, nil, &unstructured.Unstructured{})
	if err != nil {
		return nil, fmt.Errorf("unable to decode: %w", err)
	}
	return obj.(*unstructured.Unstructured), nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("got err: %v", err))
		os.Exit(1)
	}
}
