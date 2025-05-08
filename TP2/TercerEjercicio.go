package main

/*
Arreglo: [3, 8, 15, 19, 25, 31, 42, 57, 63, 79]
Buscar: 31

- Paso a paso:
Primera llamada:
inicio = 0, fin = 9

medio = (0 + 9) / 2 = 4

arreglo[4] = 25

Como 31 > 25, se descarta la mitad izquierda [3, 8, 15, 19, 25]
----------------------------------------------------------------
Segunda llamada:
inicio = 5, fin = 9

medio = (5 + 9) / 2 = 7

arreglo[7] = 57

Como 31 < 57, se descarta la mitad derecha [57, 63, 79]
----------------------------------------------------------------
Tercera llamada:
inicio = 5, fin = 6

medio = (5 + 6) / 2 = 5

arreglo[5] = 31

- ¿Qué pasa si no está?:
Si en alguna llamada inicio > fin, quiere decir que el número no está presente y la función retorna -1 o un mensaje como "no encontrado".

- ¿Qué hace este algoritmo?:
1. Divide el arreglo en mitades.
2. Descarta la mitad que no contiene el número.
3. Repite hasta encontrarlo o agotar las opciones.
*/
