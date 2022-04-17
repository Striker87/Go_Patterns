package main

import (
	"fmt"
	"patterns/snapshot/pkg"
)

/*
Снимок (Snapshot) — это поведенческий паттерн проектирования, который позволяет сохранять и восстанавливать прошлые состояния объектов, не раскрывая подробностей их реализации.
Предположим, что вы пишете программу текстового редактора.
Помимо обычного редактирования, ваш редактор позволяет менять форматирование текста, вставлять картинки и прочее.
Чтобы сделать копию состояния объекта, достаточно скопировать значение его полей.
Таким образом, если вы сделали редактор достаточно открытым, то любой другой объект сможет заглянуть внутрь, чтобы скопировать его состояние.

Возможность применения паттерна:
Когда вам нужно сохранять мгновенные снимки состояния объекта (или его части), чтобы впоследствии объект можно было восстановить в том же состоянии.
Паттерн Снимок позволяет создавать любое количество снимков объекта и хранить их, независимо от объекта, с которого делают снимок.
Снимки часто используют не только для реализации операции отмены, но и для транзакций, когда состояние объекта нужно «откатить», если операция не удалась.

Когда прямое получение состояния объекта раскрывает приватные детали его реализации, нарушая инкапсуляцию.
Паттерн Снимок - предлагает изготовить снимок самому исходному объекту, поскольку ему доступны все поля, даже приватные.

Преимущества и недостатки:
+ Не нарушает инкапсуляции исходного объекта.
+ Упрощает структуру исходного объекта. Ему не нужно хранить историю версий своего состояния.

- Требует много памяти, если клиенты слишком часто создают снимки.
- Может повлечь дополнительные издержки памяти, если объекты, хранящие историю, не освобождают ресурсы, занятые устаревшими снимками.
- В некоторых языках (например, PHP, Python, JavaScript) сложно гарантировать, чтобы только исходный объект имел доступ к состоянию снимка.

Паттерн Снимок позволяет нам сохранять «снимки» состояния объекта. В дальнейшем эти снимки можно использовать, чтобы вернуть его в прежнее состояние. Это весьма удобно в случаях, когда для объекта нужно реализовать операции «отменить/повторить».
*/

func main() {
	storage := &pkg.Keeper{
		Items: make([]*pkg.Snapshot, 0),
	}
	creator := &pkg.Creator{
		State: "A",
	}

	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("B")
	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("C")
	fmt.Printf("Creator current State: %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(1))
	fmt.Printf("Restored to State: %s\n", creator.GetState())

	creator.Restore(storage.Get(0))
	fmt.Printf("Restored to State: %s\n", creator.GetState())
}