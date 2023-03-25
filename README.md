# laas
An API that has database of open source license. It stores the basic information about the license such as short name, long name and url.

## Prerequisite

Please [install and set-up Golang](https://go.dev/doc/install) as well as [install and set-up Postman](https://www.postman.com/) on your system in advance.

## How to run this project?

1. Clone this Project and Navigate to the folder.

``` bash
https://github.com/k-avy/laas.git
cd laas
```

2. Build the project using following command.

```bash
go build ./cmd/laas
```

3. Run the executable in your vscode terminal.

```bash
./laas
```

4. You can directly run it by the following command.

```bash
go run ./cmd/laas
```

5. You can see this work on Postman on your system\
or\
you can use run it on your browser using link:

```bash
localhost:8080/licenses
```

## Features 

1. You can get the data of all licenses in the database using the following link in browser:

```bash
localhost:8080/licenses
```

2. You can post the data of the license in the database using following command in your terminal:

```bash
 curl -X POST localhost:8080/licenses -H "Content-Type: application/json" -d '{ "shortname":"LGPL-3.0-or-later","longname":"GNU Lesser General Public License v3.0 or later","url": "https:\/\/www.gnu.org\/licenses\/lgpl-3.0-standalone.html"}'
```

3. You can also get information of a single license using following link in your browser:

```bash
localhost:8080/license?shortname="name_of_license_to_search"
```
