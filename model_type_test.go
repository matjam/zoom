// Copyright 2014 Alex Browne.  All rights reserved.
// Use of this source code is governed by the MIT
// license, which can be found in the LICENSE file.

// File model_type_test.go contains unit tests for the
// code in model_type.go

package zoom

import (
	"reflect"
	"testing"
)

type regTestModel struct {
	DefaultData
	Int    int
	Bool   bool
	String string
}

func TestRegister(t *testing.T) {
	regTestModels, err := Register(&regTestModel{})
	if err != nil {
		t.Fatalf("Unexpected error in Register: %s", err.Error())
	}
	expectedName := "regTestModel"
	expectedType := reflect.TypeOf(&regTestModel{})
	testRegisteredModelType(t, regTestModels, expectedName, expectedType)

	// Effectively unregister the type by removing it from the map
	delete(modelNameToSpec, regTestModels.Name())
	delete(modelTypeToSpec, regTestModels.spec.typ)
}

func TestRegisterName(t *testing.T) {
	expectedName := "customName"
	regTestModels, err := RegisterName(expectedName, &regTestModel{})
	if err != nil {
		t.Fatalf("Unexpected error in Register: %s", err.Error())
	}
	expectedType := reflect.TypeOf(&regTestModel{})
	testRegisteredModelType(t, regTestModels, expectedName, expectedType)

	// Effectively unregister the type by removing it from the map
	delete(modelNameToSpec, regTestModels.Name())
	delete(modelTypeToSpec, regTestModels.spec.typ)
}

func testRegisteredModelType(t *testing.T, modelType *ModelType, expectedName string, expectedType reflect.Type) {
	// Check that the name and type are correct
	if modelType.Name() != expectedName {
		t.Errorf("Registered name was incorrect. Expected %s but got %s", expectedName, modelType.Name())
	}
	if modelType.spec.typ == nil {
		t.Fatalf("Registered model spec had nil type")
	}
	if modelType.spec.typ != expectedType {
		t.Errorf("Registered type was incorrect. Expected %s but got %s", expectedType.String(), modelType.spec.typ.String())
	}

	// Check that the model type was added to the appropriate maps
	if _, found := modelNameToSpec[expectedName]; !found {
		t.Error("Registered spec was not added to the modelNameToSpec map")
	}
	if _, found := modelTypeToSpec[expectedType]; !found {
		t.Error("Registered spec was not added to the modelTypeToSpec map")
	}

	// Check the underlying spec
	spec := modelType.spec
	if len(spec.fields) != 3 {
		t.Errorf("Expected spec to have 3 fields but got %d", len(spec.fields))
	}
	expectedFields := map[string]*fieldSpec{
		"Int": &fieldSpec{
			kind:      primativeField,
			name:      "Int",
			redisName: "Int",
			fieldType: reflect.TypeOf(1),
			indexKind: noIndex,
		},
		"Bool": &fieldSpec{
			kind:      primativeField,
			name:      "Bool",
			redisName: "Bool",
			fieldType: reflect.TypeOf(true),
			indexKind: noIndex,
		},
		"String": &fieldSpec{
			kind:      primativeField,
			name:      "String",
			redisName: "String",
			fieldType: reflect.TypeOf(""),
			indexKind: noIndex,
		},
	}
	for _, expectedField := range expectedFields {
		gotField, found := spec.fieldsByName[expectedField.name]
		if !found {
			t.Errorf("Expected field with name %s but it was not in spec", expectedField.name)
		}
		if !reflect.DeepEqual(expectedField, gotField) {
			t.Errorf("Field with name %s was incorrect. Expected %+v but got %+v", expectedField.name, expectedField, gotField)
		}
	}
}

func TestSave(t *testing.T) {
	testingSetUp()
	defer testingTearDown()

	// Create and save a test model
	model := createTestModels(1)[0]
	if err := testModels.Save(model); err != nil {
		t.Errorf("Unexpected error in testModels.Save: %s", err.Error())
	}

	// Make sure the model was saved correctly
	if model.Id == "" {
		t.Fatalf("model.Id is empty. Cannot continue.")
	}
	key, _ := testModels.KeyForModel(model)
	expectFieldEquals(t, key, "Int", model.Int)
	expectFieldEquals(t, key, "String", model.String)
	expectFieldEquals(t, key, "Bool", model.Bool)
	expectSetContains(t, testModels.KeyForAll(), model.Id)
}

func TestMSave(t *testing.T) {
	testingSetUp()
	defer testingTearDown()

	// Create and some test models
	models := createTestModels(5)
	if err := testModels.MSave(Models(models)); err != nil {
		t.Errorf("Unexpected error in testModels.MSave: %s", err.Error())
	}

	// Make sure each model was saved correctly
	for i, model := range models {
		if model.Id == "" {
			t.Fatalf("models[%d].Id is empty. Cannot continue.", i)
		}
		key, _ := testModels.KeyForModel(model)
		expectFieldEquals(t, key, "Int", model.Int)
		expectFieldEquals(t, key, "String", model.String)
		expectFieldEquals(t, key, "Bool", model.Bool)
		expectSetContains(t, testModels.KeyForAll(), model.Id)
	}
}

func TestFind(t *testing.T) {
	testingSetUp()
	defer testingTearDown()

	// Create and save a test model
	models, err := createAndSaveTestModels(1)
	if err != nil {
		t.Errorf("Unexpected error saving test models: %s", err.Error())
	}
	model := models[0]

	// Find the model in the database and store it in modelCopy
	modelCopy := &testModel{}
	if err := testModels.Find(model.Id, modelCopy); err != nil {
		t.Errorf("Unexpected error in testModels.Find: %s", err.Error())
	}
	if !reflect.DeepEqual(model, modelCopy) {
		t.Errorf("Found model was incorrect.\n\tExpected: %+v\n\tBut got:  %+v", model, modelCopy)
	}
}

func TestMFind(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}

func TestFindAll(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}

func TestCount(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}

func TestDelete(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}

func TestMDelete(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}

func TestDeleteAll(t *testing.T) {
	testingSetUp()
	defer testingTearDown()
}
