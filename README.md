# Golang

Una vez instalado Go por defecto en mac queda en la ruta /usr/local/go
se puede validar la instalacion en la consola escribiendo el comando go version
esto imprime la version descarga.

Ya instalado crearemos la carpeta go, preferiblemente en $HOME, donde viviran nuestros proyectos y librerias
de terceros. dentro de esta carpeta crearemos 3 carpetas mas con los sigientes nombres:
pkg, bin, y src.

Entraremos a la carpeta src y crearemos dentro de ella una carpeta github.com.
dentro de esta crearemos una carpeta mas que sera el nombre de usuario de github.
en esta carpeta estaran todos nuestros proyectos.

Por ultimo configuraremos las variables de entorno, esto dependera de el tipo de shell 
usada en mi caso bash y son agregadas en la siguiente ruta ~/.bash_profile.

export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN




