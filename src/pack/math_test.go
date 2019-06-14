package pack

import(
	"testing"

)

func TestCanAddNumbers(t *testing.T)  {

	result := Add(1,2)

	if result != 3{
		t.Log("Failed to add")
		t.Fail()
	}


	result = Add(1, 2, 3 ,4)

	if result != 10{

		t.Error("failed to add more than two")

	}




}