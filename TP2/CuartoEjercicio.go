package main

/*
Arreglo original: [38, 27, 43, 3, 9, 82, 10]

- Paso a paso:
1. División:
[38, 27, 43, 3, 9, 82, 10] → Se divide en: [38, 27, 43] y [3, 9, 82, 10]
-------------------------------------------------------------------------
2. Seguir dividiendo:
- Subarreglo [38, 27, 43] → [38], [27, 43] → [27], [43]

- Subarreglo [3, 9, 82, 10] → [3, 9], [82, 10] → [3], [9] y [82], [10]
-------------------------------------------------------------------------
3. Combinar ordenadamente:
Empezar por los subarreglos de 1 solo elemento:
[38] y [27, 43] → se combinan en [27, 38, 43]

[3, 9, 82, 10] se arma así:
- [3], [9] → [3, 9]
- [82], [10] → [10, 82]
- [3, 9], [10, 82] → [3, 9, 10, 82]
-------------------------------------------------------------------------
4. Fusionar los dos grandes subarreglos:
- [27, 38, 43] y [3, 9, 10, 82]
- Se comparan de menor a mayor y se insertan en orden
- Resultado final: [3, 9, 10, 27, 38, 43, 82]

- ¿Qué hace MergeSort?:
1. Divide el arreglo hasta que cada parte tenga 1 elemento.
2. Luego fusiona cada par comparando valores, ordenándolos al unir.
3. Es estable, y su tiempo de ejecución es O(n log n).
*/
