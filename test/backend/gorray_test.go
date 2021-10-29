package backend

import (
	"strings"
	"testing"

	"github.com/bezaeel/gorray/internal/pkg/repository"
	"github.com/bezaeel/gorray/test"
)
var records = test.GetMatrix()
func TestTranspose(t *testing.T) {
	var expected = "1,4,7\n2,5,8\n3,6,9\n"
	repo := repository.GetGorrayRepository()
	actual := repo.Transpose(records);
	if  expected != actual {
		t.Fatalf("Expected no error, got %v", actual)
	}
}

func TestFlatten(t *testing.T) {
	var expected = "1,2,3,4,5,6,7,8,9"
	repo := repository.GetGorrayRepository()
	actual := repo.Flatten(records);
	res := strings.Join(actual, ",")
	if  expected != res {
		t.Fatalf("Expected no error, got %v", actual)
	}
}

func TestAddition(t *testing.T) {
	var expected = 45
	repo := repository.GetGorrayRepository()
	actual := repo.Sum(records);
	if  expected != actual {
		t.Fatalf("Expected no error, got %v", actual)
	}
}

func TestMultiplication(t *testing.T) {
	var expected = 362880
	repo := repository.GetGorrayRepository()
	actual := repo.Multiply(records);
	if  expected != actual {
		t.Fatalf("Expected no error, got %v", actual)
	}
}