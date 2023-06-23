# modelbasierte-swe-projekt-sose23-01

## Aufgabenstellung

**Translating methods, interfaces and structural subtyping**

Recall the “run-time method look up” (RT) and “dictionary-translation” (DT) approach covered here

1. Compare the run-time performance of RT and DT (e.g. call “sumArea” in a loop and measure which version runs faster)

2. Apply the RT and DT approach to one further example of your own choice.

3. Extend RT and DT to deal with type assertions. See example below.

4. Extend RT and DT to deal with type bounds. See example below.

5. Summarize your findings in a short document (you could set up a github repo).


## Compare the run-time performance of RT and DT (e.g. call “sumArea” in a loop and measure which version runs faster)
Zunächst haben wir das Go Beispiel aus der Vorlesung [siehe Folien](https://sulzmann.github.io/ModelBasedSW/lec-go-2-types-methods-interfaces.html#(7))
so angepasst, dass wir einen Testcase für die `lookup` und eine für die `dict` Methode haben. Diese haben wir dann jeweils
`1000000000` ausgeführt, und mit dem go test tool die Laufzeit gemessen. Dabei kamen wir zu folgendem Ergebnis:


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

## Apply the RT and DT approach to one further example of your own choice.
Für diese Aufgabe haben wir uns entschieden, eine Volumenfunktion für Körper zu implementieren. Diese haben wir auf
Würfel und Zylinder angewandt.

```go
=== RUN   TestDict
--- PASS: TestDict (5.14s)
=== RUN   TestLookup
--- PASS: TestLookup (3.41s)
PASS
ok      modelbasierte-swe-projekt-sose23-01/volume      8.557s
```

Nach Ausführen der Tests fällt auf, dass die Laufzeit der `Dict` Methode nahezu unverändert ist, während die Laufzeit der
`Lookup` bei `1.000.000.000` Ausführungen etwa 1 Sekunde länger benötigt. 

## Extend RT and DT to deal with type assertions. See example below.


## Extend RT and DT to deal with type bounds. See example below.


## Summarize your findings in a short document (you could set up a github repo).




