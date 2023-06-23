# modelbasierte-swe-projekt-sose23-01

## Aufgabenstellung

**Translating methods, interfaces and structural subtyping**

Recall the “run-time method look up” (RT) and “dictionary-translation” (DT) approach covered here

1. Compare the run-time performance of RT and DT (e.g. call “sumArea” in a loop and measure which version runs faster)

2. Apply the RT and DT approach to one further example of your own choice.

3. Extend RT and DT to deal with type assertions. See example below.

4. Extend RT and DT to deal with type bounds. See example below.

5. Summarize your findings in a short document (you could set up a github repo).


## Compare the run-time performance of RT and DT
Zunächst haben wir das Go Beispiel aus der Vorlesung [siehe Folien](https://sulzmann.github.io/ModelBasedSW/lec-go-2-types-methods-interfaces.html#(7))
so angepasst, dass wir einen Testcase für die `lookup` und eine für die `dict` Methode haben. Diese haben wir dann jeweils
`1.000.000.000` ausgeführt, und mit dem go test tool die Laufzeit gemessen. Die Implementierung befindet sich in
[./area/area.go](./area/area.go).

Dabei kamen wir zu folgendem Ergebnis:


```go
=== RUN   TestDict
--- PASS: TestDict (5.18s)
=== RUN   TestLookup
--- PASS: TestLookup (2.45s)
PASS
ok      modelbasierte-swe-projekt-sose23-01     7.625s
```

Das Ergebnis unserer Evaluierung zeigt, dass die `lookup` Methode etwa doppelt so schnell ausgeführt wird wie die `dict` 
Methode. Festellen konnten wir außerdem, dass die Implementierung der `lookup` Methode weniger Zeilen Code benötigt 
und zunächst intuituver wirkt.

## Apply the RT and DT approach to one further example
Für diese Aufgabe haben wir uns entschieden, eine Volumenfunktion für Körper zu implementieren. Diese haben wir auf
Würfel und Zylinder angewandt. 
Die Implementierung befindet sich in [./volume/volume.go](./volume/volume.go).

```go
=== RUN   TestDict
--- PASS: TestDict (5.14s)
=== RUN   TestLookup
--- PASS: TestLookup (3.41s)
PASS
ok      modelbasierte-swe-projekt-sose23-01/volume      8.557s
```

Nach Ausführen der Tests fällt auf, dass die Laufzeit der `dict` Methode nahezu unverändert ist, während die Laufzeit der
`lookup` bei `1.000.000.000` Ausführungen etwa 1 Sekunde länger benötigt. 

## Extend RT and DT to deal with type assertions
Die `lookup` Methode implementiert die Type Assertion berets durch das Switch-Case
(siehe Zeile 41 [./areaTypeAssertion/area.go](./areaTypeAssertion/area.go)).

Für die `dict` Methode haben wir in die `sumArea_Dict` Methode zusätzlich Type Assertion Prüfungen eingebaut (siehe Zeile 62
[./areaTypeAssertion/area.go](./areaTypeAssertion/area.go)).


Auffällig war beim Testen, dass die `dict` Methode für den gleichen TestCase nun mit `16.96s` wesentlich länger benötigt.

```go 
=== RUN   TestDict
--- PASS: TestDict (16.96s)
=== RUN   TestLookup
--- PASS: TestLookup (3.11s)
PASS
ok      modelbasierte-swe-projekt-sose23-01/areaTypeAssertion   20.071s
```


## Extend RT and DT to deal with type bounds

### Dict
Wir implementieren Type Bounds für die `dict` Methode (siehe Zeile 55 [./areaTypeBounds/area.go](./areaTypeBounds/area.go)).
> Achtung: Da die beiden Generischen Typen zwar beide Subtyp von Shape sind, in der konkreten Implementierung allerdings
> andere Typen (Shape, Rectangle) annehmen können, benötigen wir zwei Typ Parameter T und R.

Führen wir wieder den Laufzeit Test aus, so erhalten wir folgendes Ergebnis:

```go 
=== RUN   TestDict
--- PASS: TestDict (3.67s)
=== RUN   TestLookup
--- PASS: TestLookup (2.93s)
PASS
ok      modelbasierte-swe-projekt-sose23-01/areaTypeBounds      6.607s
```

Es fällt auf, dass die `Dict` Methode mit dieser Implementierung wesentlich schneller ist als die vergleichbare Type 
Assertion Implementierung.

### Lookup
Für die Implementierung von `lookup` spezifizieren wir Type Bounds für die `area_Lookup` und `sumArea_Lookup` Methoden
(siehe Zeile 36 [./areaTypeBounds/area.go](./areaTypeBounds/area.go)).

Dadurch, dass der Typ jetzt direkt bekannt ist, können wir die `area()` Methode einfach mit dem Receiver Type von `x`
aufrufen: `x.area()`.

Der Laufzeit Test ergibt:

```go 
=== RUN   TestDict
--- PASS: TestDict (3.70s)
=== RUN   TestLookup
--- PASS: TestLookup (3.72s)
PASS
ok      modelbasierte-swe-projekt-sose23-01/areaTypeBounds      7.424s
```

Interessant ist, dass der Lookup TestCase etwa 1 Sekunde länger benötigt, als die ursprüngliche Implementation ohne Type
Bounds. 


## Summary

Übersicht der gemessenen Laufzeiten:

|            | area | volume  | areaTypeAssertion | areaTypeBounds 
|------------|------|---------| --------| --------
| **lookup** | 5.18s | 5.14s  | 16.96s | 3.70s
| **dict**   | 2.45s | 3.41s  | 3.11s | 3.72s

