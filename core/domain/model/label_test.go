package model

import (
	"github.com/reecerussell/tw-management-system/core/domain/datamodel"
	"github.com/reecerussell/tw-management-system/core/domain/dto"
	"testing"
)

func TestNewLabel(t *testing.T) {
	testName := "My Label"
	data := &dto.CreateLabel{Name: testName}
	lbl := NewLabel(data)
	if lbl == nil {
		t.Errorf("exoected an instance of Label, but got nil")
	}

	if lbl.name != testName {
		t.Errorf("label's name expected to be '%s' but was '%s'", testName, lbl.name)
	}
}

func TestLabel_Update(t *testing.T) {
	l := &Label{name: "My Label"}
	d := &dto.Label{Name: "My Label 2"}

	l.Update(d)
	if l.name != d.Name {
		t.Errorf("label's name expected to be '%s' but was '%s'", d.Name, l.name)
	}
}

func TestLabel_DataModel(t *testing.T) {
	l := &Label{
		id: "148",
		name: "My Label",
	}

	dm := l.DataModel()
	if l.id != dm.ID {
		t.Errorf("the datamodel's id property was expected to be '%s' but was '%s'", l.id, dm.ID)
	}

	if l.name != dm.Name {
		t.Errorf("the datamodel's name property was expected to be '%s' but was '%s'", l.name, dm.Name)
	}
}

func TestLabel_DTO(t *testing.T) {
	l := &Label{
		id: "148",
		name: "My Label",
	}

	d := l.DTO()
	if l.id != d.ID {
		t.Errorf("the dto's id property was expected to be '%s' but was '%s'", l.id, d.ID)
	}

	if l.name != d.Name {
		t.Errorf("the dto's name property was expected to be '%s' but was '%s'", l.name, d.Name)
	}
}

func TestLabelFromDataModel(t *testing.T) {
	dm := &datamodel.Label{
		ID: "148",
		Name: "My Label",
	}

	l := LabelFromDataModel(dm)
	if l.id != dm.ID {
		t.Errorf("the label's id property was expected to be '%s' but was '%s'", dm.ID, l.id)
	}

	if l.name != dm.Name {
		t.Errorf("the label's name property was expected to be '%s' but was '%s'", dm.Name, l.name)
	}
}

func TestLabel_Name(t *testing.T) {
	testName := "MyLabel"
	l := &Label{name: testName}

	if l.Name() != testName {
		t.Errorf("the label's name should be '%s' not '%s", testName, l.Name())
	}
}