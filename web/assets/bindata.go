package assets

import (
	"fmt"
	"strings"
)
var _web_static_navigation_bar_css = []byte(``)

func web_static_navigation_bar_css() ([]byte, error) {
	return _web_static_navigation_bar_css, nil
}

var _web_static_style_css = []byte(`h1, h2, h3, h4, h5 {
    text-align: center;
}

ul li {
    list-style:none;
}

`)

func web_static_style_css() ([]byte, error) {
	return _web_static_style_css, nil
}

var _web_static_third_view_css = []byte(`.result-box {
    margin-top: 50px;
}

.result-underlined {
    text-decoration: underline;
}`)

func web_static_third_view_css() ([]byte, error) {
	return _web_static_third_view_css, nil
}

var _web_static_todolist_css = []byte(`/* Include the padding and border in an element's total width and height */
* {
  box-sizing: border-box;
}

/* Remove margins and padding from the list */
ul {
  margin: 0;
  padding: 0;
}

/* Style the list items */
ul li {
  cursor: pointer;
  position: relative;
  padding: 12px 8px 12px 40px;
  background: #eee;
  font-size: 18px;
  transition: 0.2s;

  /* make the list items unselectable */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

/* Set all odd list items to a different color (zebra-stripes) */
ul li:nth-child(odd) {
  background: #f9f9f9;
}

/* Darker background-color on hover */
ul li:hover {
  background: #ddd;
}

/* When clicked on, add a background color and strike out text */
ul li.checked {
  background: #888;
  color: #fff;
  text-decoration: line-through;
}

/* Add a "checked" mark when clicked on */
ul li.checked::before {
  content: '';
  position: absolute;
  border-color: #fff;
  border-style: solid;
  border-width: 0 2px 2px 0;
  top: 10px;
  left: 16px;
  transform: rotate(45deg);
  height: 15px;
  width: 7px;
}


.nametask {
  list-style:none;
}

/* Style the close button */
.close {
  position: absolute;
  right: 0;
  top: 0;
  padding: 12px 16px 12px 16px;
}

.close:hover {
  background-color: #f44336;
  color: white;
}

/* Style the header */
.header {
  background-color: #47309C;
  padding: 30px 40px;
  color: white;
  text-align: center;
}

/* Clear floats after the header */
.header:after {
  content: "";
  display: table;
  clear: both;
}

/* Style the input */
input {
  margin: 0;
  border: none;
  border-radius: 0;
  width: 75%;
  padding: 10px;
  float: left;
  font-size: 16px;
}

/* Style the "Add" button */
.addBtn {
  padding: 10px;
  width: 25%;
  background: #d9d9d9;
  color: #555;
  float: left;
  text-align: center;
  font-size: 16px;
  cursor: pointer;
  transition: 0.3s;
  border-radius: 0;
}

.addBtn:hover {
  background-color: #bbb;
}`)

func web_static_todolist_css() ([]byte, error) {
	return _web_static_todolist_css, nil
}

var _web_templates_index_html = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>ToDoList | index2</title>

        
        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/web/static/style.css">
        <link rel="stylesheet" href="/web/static/navigation_bar.css">
        <link rel="stylesheet" href="/web/static/todolist.css">
    </head>


    <body>
        <div class="container">
            {{.NavigationBar}}
        </div>

        



        <div class="container">
            <div id="myDIV" class="header">
                <h2>My To Do List</h2>
                <input type="text" id="myInput" placeholder="Write your task...">
                <span onclick="newElement()" class="addBtn">Add</span>
              </div>
          
              <ul id="myUL">
                <li class="nametask">Hit the gym</li>
                <li  class="nametask" class="checked">Pay bills</li>
                <li class="checked">Pay bills2</li>
              </ul>
        </div>

        
    </body>



    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
    <script>
     
        // Create a "close" button and append it to each list item
        var myNodelist = document.getElementsByClassName("nametask");
        var i;
        for (i = 0; i < myNodelist.length; i++) {
          var span = document.createElement("SPAN");
          var txt = document.createTextNode("\u00D7");
          span.className = "close";
          span.appendChild(txt);
          myNodelist[i].appendChild(span);
        }
        
        // Click on a close button to hide the current list item
        var close = document.getElementsByClassName("close");
        var i;
        for (i = 0; i < close.length; i++) {
          close[i].onclick = function() {
            var div = this.parentElement;
            div.style.display = "none";

            deleteinbackend(this.parentElement.id);  //se llama a la función para borrar el elemento en el backend
          }
        }
        
        // Add a "checked" symbol when clicking on a list item
        var list = document.querySelector('ul');
        list.addEventListener('click', function(ev) {
          if (ev.target.tagName === 'li') {
            ev.target.classList.toggle('checked');
          }
        }, false);
        
       
        
        /////////////////////////////////////////////////////////////
        /////////////////////////////////////////////////////////////
        
        const url = "http://localhost:8000";
        window.onload = handleGETtasksButton();

        function handleGETtasksButton() {   //muestra todas las tareas al cargar la pagina
            
            // solicitud a la API REST vía método GET
            fetch(url + "/tasks")
            // promesas encadenadas
            .then(res => res.json())
            // procesamiento de la respuesta
            .then(json => {
                // depuración en consola del navegador
                console.log(json);
                
                
                for(i=0; i<json.length; i++){
                    
                    var li = document.createElement("li");
                    var nombre = json[i].name;
                    
                    var t = document.createTextNode(nombre);
                    li.appendChild(t);
                    li.id = json[i].ID; //le asigna a cada etiqueta HTML <li> el ID de la task
                      
                    var span = document.createElement("SPAN");
                    var txt = document.createTextNode("\u00D7");
                    span.className = "close";
                    span.appendChild(txt);
                    li.appendChild(span);

                    document.getElementById("myUL").appendChild(li);
                    
                }	
                for (i = 0; i < close.length; i++) {
                  close[i].onclick = function() {
                    var div = this.parentElement;
                    div.style.display = "none";
                    deleteinbackend(this.parentElement.id);
                  }
                }
            
            })
            
        }	 
        
        ///////////////////////////////////////
        ///////////////////////////////////////
        

        // Create a new list item when clicking on the "Add" button
        function newElement() {
          var li = document.createElement("li");
          var taskname = document.getElementById("myInput").value;
          var t = document.createTextNode(taskname);
          li.appendChild(t);

          
          if (taskname === '') {
            alert("You must write something!");
          } else {
            document.getElementById("myUL").appendChild(li);


            // solicitud a la API REST vía método POST
            fetch(url + "/tasks", {
            method: 'POST',
            
            // se incluyen en el cuerpo del mensaje los datos del nuevo libro a registrar 
            body: JSON.stringify({            
                name: taskname
            })
            })
            // promesas encadenadas
            .then(res => res.text())
            // procesamiento de la respuesta
            .then(text => {
                // depuración en consola del navegador
                console.log(text);
                li.id=text; //le asigno el ID de la task devuelto por el backend a la etiqueta HTML <li>

                // conversión de respuesta de texto (html) a objeto json
                 //json = JSON.parse(text);
                 //volcado a pantalla de la respuesta
                 //movieId.innerHTML = json.id;
            })
            // procesamiento de errores
            .catch(err => {
                // depuración en consola del navegador
                console.error(err);
            });



          }
          document.getElementById("myInput").value = "";  //se vuelve a poner en blanco el textbox
        
          var span = document.createElement("SPAN");
          var txt = document.createTextNode("\u00D7");
          span.className = "close";
          span.appendChild(txt);
          li.appendChild(span);
        
          for (i = 0; i < close.length; i++) {
            close[i].onclick = function() {
              var div = this.parentElement;
              div.style.display = "none";
              deleteinbackend(this.parentElement.id);
            }
          }

        }


        /////////////////////////////
        ////////////////////////////

        function deleteinbackend(id){   //para borrar la task en el backend
            
            // solicitud a la API REST vía método POST
            fetch(url + "/tasks/"+id, {
            method: 'DELETE',
            
            })
            // promesas encadenadas
            .then(res => res.text())
            // procesamiento de la respuesta
            .then(text => {
                // depuración en consola del navegador
                console.log(text);
                
            })
            // procesamiento de errores
            .catch(err => {
                // depuración en consola del navegador
                console.error(err);
            });

        }
        </script>
        
    </html>`)

func web_templates_index_html() ([]byte, error) {
	return _web_templates_index_html, nil
}

var _web_templates_navigation_bar_html = []byte(`<!-- Navigation Bar -->
<nav class="navbar navbar-light bg-light">
    <div class="container">
        <span class="navbar-brand">My ToDoList</span>

        <ul class="nav justify-content-end">
            <li class="nav-item">
                <a id="homeNav" class="nav-link" href="/">Home</a>
            </li>
            <!-- <li class="nav-item">
                <a id="secondNav" class="nav-link" href="/second">Second</a>
            </li>
            <li class="nav-item">
                <a id="thirdNav" class="nav-link" href="/third/1">Third</a>
            </li> -->
        </ul>
    </div>
</nav>`)

func web_templates_navigation_bar_html() ([]byte, error) {
	return _web_templates_navigation_bar_html, nil
}

var _web_templates_second_view_html = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>Golang HTML Server</title>

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/static/style.css">
        <link rel="stylesheet" href="/static/navigation_bar.css">
    </head>
    <body>
        <div class="container">
            {{.NavigationBar}}

            <h1>Another View</h1>
            <h2>- Content Goes Here -</h2>
        </div>
    </body>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</html>`)

func web_templates_second_view_html() ([]byte, error) {
	return _web_templates_second_view_html, nil
}

var _web_templates_third_view_html = []byte(`<!DOCTYPE html>
<html lang="en">
    <head>
        <!-- Required meta tags -->
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <title>Golang HTML Server</title>

        <!-- Bootstrap CSS -->
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">

        <link rel="stylesheet" href="/static/style.css">
        <link rel="stylesheet" href="/static/navigation_bar.css">
        <link rel="stylesheet" href="/static/third_view.css">
    </head>
    <body>
        <div class="container">
            {{.NavigationBar}}

            <h1>Rendering Data</h1>
            <h4>This page takes the integer passed in and determines if it is odd or even</h4>
            <div class="result-box">
                {{if .StringQuery}}
                    <h2 class="result-underlined">You didn't enter an integer. This is what was entered:</h2>
                    <h3>{{.StringQuery}}</h3>
                {{else}}
                    <h2 class="result-underlined">The number entered is:</h2>
                    <h3>{{.Number}}</h3>
                    <h2 class="result-underlined">This number is:</h2>
                    <h3>{{.Number | formatOddOrEven}}</h3>
                {{end}}
            </div>
        </div>
    </body>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</html>`)

func web_templates_third_view_html() ([]byte, error) {
	return _web_templates_third_view_html, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"web/static/navigation_bar.css": web_static_navigation_bar_css,
	"web/static/style.css": web_static_style_css,
	"web/static/third_view.css": web_static_third_view_css,
	"web/static/todolist.css": web_static_todolist_css,
	"web/templates/index.html": web_templates_index_html,
	"web/templates/navigation_bar.html": web_templates_navigation_bar_html,
	"web/templates/second_view.html": web_templates_second_view_html,
	"web/templates/third_view.html": web_templates_third_view_html,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"web/static/navigation_bar.css": &_bintree_t{web_static_navigation_bar_css, map[string]*_bintree_t{
	}},
	"web/static/style.css": &_bintree_t{web_static_style_css, map[string]*_bintree_t{
	}},
	"web/static/third_view.css": &_bintree_t{web_static_third_view_css, map[string]*_bintree_t{
	}},
	"web/static/todolist.css": &_bintree_t{web_static_todolist_css, map[string]*_bintree_t{
	}},
	"web/templates/index.html": &_bintree_t{web_templates_index_html, map[string]*_bintree_t{
	}},
	"web/templates/navigation_bar.html": &_bintree_t{web_templates_navigation_bar_html, map[string]*_bintree_t{
	}},
	"web/templates/second_view.html": &_bintree_t{web_templates_second_view_html, map[string]*_bintree_t{
	}},
	"web/templates/third_view.html": &_bintree_t{web_templates_third_view_html, map[string]*_bintree_t{
	}},
}}
