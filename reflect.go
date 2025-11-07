package main

import (
	"fmt"
	"reflect"
)

// Struct contoh untuk demonstrasi
type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email"`
	private string // field private tidak akan terlihat
}

func main() {
	fmt.Println("=== TUTORIAL PACKAGE REFLECT GOLANG ===\n")

	// 1. DASAR-DASAR REFLECT
	fmt.Println("1. MENDAPATKAN TYPE DAN VALUE")
	name := "Budi"
	age := 25

	// Mendapatkan tipe data
	typeOfName := reflect.TypeOf(name)
	typeOfAge := reflect.TypeOf(age)
	fmt.Printf("Tipe dari 'name': %v\n", typeOfName)
	fmt.Printf("Tipe dari 'age': %v\n", typeOfAge)

	// Mendapatkan value
	valueOfName := reflect.ValueOf(name)
	fmt.Printf("Value dari 'name': %v\n", valueOfName)
	fmt.Printf("Kind dari 'name': %v\n", valueOfName.Kind())
	fmt.Println()

	// 2. BEKERJA DENGAN STRUCT
	fmt.Println("2. INSPEKSI STRUCT")
	person := Person{
		Name:    "Siti",
		Age:     30,
		Email:   "siti@example.com",
		private: "rahasia",
	}

	typeOfPerson := reflect.TypeOf(person)
	valueOfPerson := reflect.ValueOf(person)

	fmt.Printf("Tipe: %v\n", typeOfPerson)
	fmt.Printf("Jumlah field: %d\n", typeOfPerson.NumField())
	fmt.Println()

	// Iterasi semua field
	fmt.Println("Field-field dalam struct:")
	for i := 0; i < typeOfPerson.NumField(); i++ {
		field := typeOfPerson.Field(i)
		value := valueOfPerson.Field(i)

		fmt.Printf("  Field %d: %s\n", i, field.Name)
		fmt.Printf("    Tipe: %v\n", field.Type)
		fmt.Printf("    Value: %v\n", value)
		fmt.Printf("    Tag JSON: %v\n", field.Tag.Get("json"))
		fmt.Printf("    Tag Validate: %v\n", field.Tag.Get("validate"))
		fmt.Println()
	}

	// 3. MENGUBAH VALUE (HARUS POINTER)
	fmt.Println("3. MENGUBAH VALUE DENGAN REFLECT")
	score := 100
	fmt.Printf("Score awal: %d\n", score)

	// Untuk mengubah value, harus menggunakan pointer
	valueOfScore := reflect.ValueOf(&score).Elem()
	if valueOfScore.CanSet() {
		valueOfScore.SetInt(200)
		fmt.Printf("Score setelah diubah: %d\n", score)
	}
	fmt.Println()

	// 4. MENGUBAH FIELD STRUCT
	fmt.Println("4. MENGUBAH FIELD STRUCT")
	personPtr := &person
	valueOfPersonPtr := reflect.ValueOf(personPtr).Elem()

	// Mengubah field Name
	nameField := valueOfPersonPtr.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Ahmad")
		fmt.Printf("Name diubah menjadi: %s\n", person.Name)
	}

	// Mengubah field Age
	ageField := valueOfPersonPtr.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(35)
		fmt.Printf("Age diubah menjadi: %d\n", person.Age)
	}
	fmt.Println()

	// 5. CEK TIPE DATA DENGAN KIND
	fmt.Println("5. MENGECEK KIND (JENIS TIPE)")
	checkKind("Hello")
	checkKind(123)
	checkKind(3.14)
	checkKind([]int{1, 2, 3})
	checkKind(map[string]int{"a": 1})
	checkKind(person)
	fmt.Println()

	// 6. BEKERJA DENGAN SLICE
	fmt.Println("6. BEKERJA DENGAN SLICE")
	numbers := []int{10, 20, 30}
	valueOfNumbers := reflect.ValueOf(numbers)

	fmt.Printf("Panjang slice: %d\n", valueOfNumbers.Len())
	fmt.Println("Isi slice:")
	for i := 0; i < valueOfNumbers.Len(); i++ {
		fmt.Printf("  Index %d: %v\n", i, valueOfNumbers.Index(i))
	}
	fmt.Println()

	// 7. BEKERJA DENGAN MAP
	fmt.Println("7. BEKERJA DENGAN MAP")
	studentScores := map[string]int{
		"Budi": 85,
		"Ani":  90,
		"Caca": 88,
	}

	valueOfMap := reflect.ValueOf(studentScores)
	fmt.Println("Isi map:")
	for _, key := range valueOfMap.MapKeys() {
		value := valueOfMap.MapIndex(key)
		fmt.Printf("  %v: %v\n", key, value)
	}
	fmt.Println()

	// 8. MEMBUAT VALUE BARU
	fmt.Println("8. MEMBUAT VALUE BARU DENGAN REFLECT")
	// Membuat slice baru
	sliceType := reflect.TypeOf([]int{})
	newSlice := reflect.MakeSlice(sliceType, 0, 5)
	newSlice = reflect.Append(newSlice, reflect.ValueOf(100))
	newSlice = reflect.Append(newSlice, reflect.ValueOf(200))
	fmt.Printf("Slice baru: %v\n", newSlice)

	// Membuat map baru
	mapType := reflect.TypeOf(map[string]string{})
	newMap := reflect.MakeMap(mapType)
	newMap.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf("value1"))
	fmt.Printf("Map baru: %v\n", newMap)
	fmt.Println()

	// 9. CONTOH PRAKTIS: VALIDASI STRUCT
	fmt.Println("9. CONTOH PRAKTIS: SIMPLE VALIDATOR")
	validateStruct(person)
}

// Helper function untuk cek kind
func checkKind(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Printf("Value: %v, Kind: %v\n", i, v.Kind())
}

// Contoh validator sederhana
func validateStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		fmt.Println("Bukan struct!")
		return
	}

	fmt.Println("Validasi struct:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "required" {
			if value.IsZero() {
				fmt.Printf("  ❌ Field '%s' wajib diisi!\n", field.Name)
			} else {
				fmt.Printf("  ✓ Field '%s' valid\n", field.Name)
			}
		}
	}
}