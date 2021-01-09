# dfa-minimization-algorithm
### El algoritmo Hopcroft para minimizar automatas finitos deterministicos es el mas eficiente propuesto hasta el momento, siendo O(n * k * log n)<sup>[1](#aclaracion)</sup>. Otro algoritmo conocido es el de Moore, siendo O(n^2).
### Ambos se basan  en la idea de refinamiento de particiones en clases de equivalencias de modo que al final del proceso cada una represente un estado de comportamiento unico en el automata mínimo.
Hay detalles de implementación a corregir según el paper de Hopcroft.



## Como utilizar

Para construir el automata con las estructuras es muy simple:
```
// DFA represents a finite automata
type DFA struct {
	States       Partition 			
	Alphabet     []int				
	InitialState State				
	FinalStates  Partition		
	Delta        map[State]map[int]State 
}
```

- States: representa los estados y se representan con valores int.
- Alphabet: representa el conjunto de posibles inputs, también se representa como un listado de int.
- InitialState: representa el único estado inicial.
- FinalStates: representa el conjunto de estados finales
- Delta: representa la función de transición en la forma: Delta[w] = { \< i : q \> | i in Alphabet ^ existe una arista dirigida de w a q mediante i }.

Una vez construido, llamar a la función:

```
  min := HopcroftDFAMin(M)
```
Siendo min el automata minimo. No hay relación entre el nombre de los estados de M y min.

## Observaciones
- No debe haber estados inalcanzables


## Papers:
* https://drive.google.com/file/d/16OnG54Z9QD2O4KD9Ca1CP-ClOiKibXRd/view?usp=sharing
* https://pdfs.semanticscholar.org/513b/7905cd8fe6a3b78c768d2b8cf7be2e4be66a.pdf
* https://ecommons.cornell.edu/bitstream/handle/1813/6002/72-151.pdf;jsessionid=2526590ADEC7619EFF8990BFA00FF155?sequence=1

<a name="aclaracion">1</a>: n es la cantidad de estados. k es el tamaño del alfabeto, asumimos en este caso finito.
