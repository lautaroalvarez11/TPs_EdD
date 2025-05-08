package main

/* el nodo raiz es el unico que no tiene padre
.el arbol binario es aquel que tiene solamente como maximo dos hijos
(las reglas
que utilizamos para ordenar en el arbo,
es a la izquierda los menores y a la derecha los mayores,
es por convencion pero se debe respetar siempre esto
)

.el arbol AVL adquiere mas complejidad que el binario (detallaremos despues)


.el arbol binario de busqueda es aquel donde tendras un nodo raiz y a partir de ahi
tendra un hijo izquierdo y un hijo derecho, y asi sucesivamente menor-mayor
o mayor-menor

la busqueda en el arbol binario es igual a la insercion.
Borrar NODO HOJA ES SENCILLO pero si quiero borrar una rama es mas complicado
ya que debemos reordenar el arbol
debo reemplazarlo por el mayor del subarbol izquierdo(una vez hacia la izquierda y
todas las que pueda hacia la derecha ) o
el menor del subarbol derecho (una vez hacia la derecha y
 todas las que pueda hacia la izquierda)


 .supongamos que tenemos un arbol (1,2,3,4,5,6,7)
 es una lista ordenada (cumple con el arbol binario)
 el problema que tiene esto es que pueden ser arboles desbalanceado


 .el arbol AVL es un arbol binario de busqueda que se encuentra balanceado
 ya que la rama izquierda tiene que tener el mismo nivel que la rama derecha
 Lo mismo aplica al revez.

 El factor de equilibrio en un árbol AVL es la diferencia
  entre las alturas de los subárboles izquierdo y derecho de un nodo.
  Se calcula como altura(subárbol izquierdo) - altura(subárbol derecho).
   En un árbol AVL equilibrado,
  el factor de equilibrio de cada nodo debe ser -1, 0 o 1.

  El factor equilibrio :
  se calcula con la altura ( DESDE EL NODO QUE ESTOY ESTUDIANDO HACIA EL
  CAMINO MAS LARGO QUE PUEDO TOMAR HACIA UNA HOJA,ES DECIR
  HACIA EL FINAL DEL NODO)


  desbalances se pueden resolver con 4 tipos de rotaciones :
  . rotacion simple a la derecha
  . rotacion simple a la izquierda
  . rotacion doble a la derecha
  . rotacion doble a la izquierda

  signos distintos indica que tenes que hacer una rotacion doble si o si
  para balancear el arbol. (mismo signo(rotacion simple)
  signos distintos(rotacion doble)
  en el caso de la rotacion doble dependera el sentido a
  donde tendremos que rotar
  dependiendo de cual es positivo y cual es el negativo


  la estructura de los arboles son RECURSIVAS


  cuando hacemos una rotacion simple, el hijo pasa a ser padre y el padre a hijo,
  y los hijos del anterior padre pasan a ser los hijos del nuevo hijo
 )

 si reemplazamos : mayor subarbol izquierdo o menor del subarbol derecho
 que el arbol quede balanceado tomando y reemplazando por el mayor
 subarbol izquierdo NO ASEGURA que el arbol quede balanceado tomando
 el menor subarbol derecho  (viceversa tambien aplica )
 ESTAMOS OBLIGADOS A CALCULAR FACTORES DE EQUILIBRIO ANTE CUALQUIER MODIFICACION




*/
