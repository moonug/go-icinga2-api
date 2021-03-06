package iapi

import "testing"

func TestGetValidHost(t *testing.T) {

	hostname := "c1-mysql-1"

	_, err := Icinga2_Server.GetHost(hostname)

	if err != nil {
		t.Error(err)
	}
}

func TestGetInvalidHost(t *testing.T) {

	hostname := "c2-mysql-1"
	_, err := Icinga2_Server.GetHost(hostname)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateSimpleHost(t *testing.T) {

	hostname := "go-icinga2-api-1"
	IPAddress := "127.0.0.2"
	CheckCommand := "CheckItRealGood"
	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, CheckCommand, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateHostWithVariables(t *testing.T) {

	hostname := "go-icinga2-api-2"
	IPAddress := "127.0.0.3"
	CheckCommand := "CheckItRealGood"
	variables := make(map[string]string)

	variables["vars.os"] = "Linux"
	variables["vars.creator"] = "Terraform"

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, CheckCommand, variables)
	if err != nil {
		t.Error(err)
	}

	// Delete host after creating it.
	deleteErr := Icinga2_Server.DeleteHost(hostname)
	if deleteErr != nil {
		t.Error(err)
	}
}

func TestDeleteHost(t *testing.T) {

	hostname := "go-icinga2-api-1"

	err := Icinga2_Server.DeleteHost(hostname)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHostDNE(t *testing.T) {
	hostname := "go-icinga2-api-1"
	err := Icinga2_Server.DeleteHost(hostname)
	if err.Error() != "No objects found." {
		t.Error(err)
	}
}
