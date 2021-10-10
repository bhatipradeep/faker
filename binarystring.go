package faker

//BinaryString is the faker struct for BinaryStrict
type BinaryString struct {
	faker *Faker
}

// BinaryStringOfLength returns a random binarystring of given input length
func (bn BinaryString) BinaryStringOfLength(length int) string {
	var bs string
	for i := 0; i < length; i++ {
		bs += bn.faker.RandomStringElement([]string{"0", "1"})
	}
	return bs
}
