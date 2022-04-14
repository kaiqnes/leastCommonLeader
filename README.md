# Least common leader

## Description
Receive a structure that represents an employee and their direct manager. Ex: ```(Alan, James)```

**Given info:** 
```(Alan, James)
(Alex, Tiffany)
(Ray, Tiffany)
(James, Daniel)
(Tiffany, Daniel)
```

Find the least common leader for given two people.
```
Exp 1:
INPUT Alex, Ray
OUTPUT Tiffany
```
```
Exp 2:
INPUT Tiffany, James
OUTPUT Daniel
```
```
Exp 3:
INPUT Ray, James
OUTPUT Daniel
```