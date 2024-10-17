package animal

type Animal interface {
	Speak() string
	GetName() string
}

//func makeVoiceOfAnimal(a animal) { //Без выше описанных методов обобщенная функция не будет работать
//	fmt.Printf("Животное по имени: %s, сказало: %s\n", a.getName(), a.speak())
//}
