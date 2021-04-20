## literal
1. number
2. bool
3. string
4. null
5. name
6. array
7. object(json)

## expression
1. group: (expression)
2. member access: name.member
3. function call: expr()
4. math multi: expr * expr, expr / expr
5. math add: expr + expr, expr - expr
6. assignment: name = expr
7. function declaration

## statement
1. var decl: var name = expr
2. function declaration: function [name] ([arguments]) { [statements] }
3. expr statement: function call expr, assignment expr
4. return statement: return expr (must be in function body)
5. if statement: if (expression) {}
6. for statement: for (statement; expr; statement) {}
7. 