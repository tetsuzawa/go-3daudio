package sofalib

import (
	"gonum.org/v1/hdf5"
	"log"
)

type Data struct {

}

type HDR5SOFA struct {
	C hdf5.Dataset
	Delay hdf5.Dataset
	IR hdf5.Dataset

}

type GeneralMeta struct {
	Conventions            string
	Title                  string
	Version                string
	DataType               string
	SOFAConventions        string
	SOFAConventionsVersion string
	ApplicationName        string
	ApplicationVersion     string
	AuthorContact          string
	License                string
	Organization           string
	History                string
	Comment                string
	RoomType               string
	DatabaseName           string
	DateCreated            string
	DateModified           string
	ListenerShortName      string
	APIName                string
	APIVersion             string
}

func NewSOFA(sofaFileName string) *SOFA {
	var sofa = new(SOFA)
	if sofaFileName != "" {
		sofa = fromFile(sofaFileName)
	} else {
		sofa = &SOFA{}
	}
	return sofa
}

func fromFile(sofaFileName string) *SOFA {
	// TODO
	//if not isinstance(sofa_file, tbls.file.File):
	//assert isinstance(sofa_file, str)
	//tblsfile = tbls.open_file(sofa_file)
	tblsFile, err := hdf5.OpenFile(sofaFileName, hdf5.F_ACC_RDONLY)
	if err != nil {
		log.Fatalln(err)
	}
	//else:
	//tblsfile = sofa_file

	// List of the Attributes that are required by the SOFA specivication
	var req_attributes = []string{
		"Conventions", "Version", "SOFAConventions",
		"SOFAConventionsVersion", "DataType",
		"RoomType", "Title", "DateCreated",
		"DateModified", "APIName", "APIVersion",
		"AuthorContact", "Organization", "License",
	}

	for param := range req_attributes {
		addAttributes(t)
	}

}

func addAttributes(object SOFA, tblsfile hdf5.File, node string, attribute string, isrequired bool, name string) {
	hdf5.OpenFile()

}
