#### Утилита dirtree

Выводит дерево каталогов и файлов (если указана опция -f) на подобии утилиты `tree`

```
├───Makefile (64b)
├───dockerfile (79b)
├───go.mod (51b)
├───main.go (1777b)
├───main_test.go (1865b)
├───readme.md (1594b)
└───testdata
        ├───project
        │       ├───file.txt (19b)
        │       └───gopher.png (70372b)
        ├───static
        │       ├───a_lorem
        │       │       ├───dolor.txt (empty)
        │       │       ├───gopher.png (70372b)
        │       │       └───ipsum
        │       │               └───gopher.png (70372b)
        │       ├───css
        │       │       └───body.css (28b)
        │       ├───empty.txt (empty)
        │       ├───html
        │       │       └───index.html (57b)
        │       ├───js
        │       │       └───site.js (10b)
        │       └───z_lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        ├───zline
        │       ├───empty.txt (empty)
        │       └───lorem
        │               ├───dolor.txt (empty)
        │               ├───gopher.png (70372b)
        │               └───ipsum
        │                       └───gopher.png (70372b)
        └───zzfile.txt (empty)
```
##### Использование

- `make build`
- `./dirtree -f .`

##### Тесты

- `make test`
