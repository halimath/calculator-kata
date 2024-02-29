
from sys import stdin, argv
from .evaluator import eval_expr

result = None

if len(argv) > 1:
    with open(argv[1], encoding='UTF-8') as f:
        result = eval_expr(f)
else:
    result = eval_expr(stdin)

print(f"{result:.4f}")