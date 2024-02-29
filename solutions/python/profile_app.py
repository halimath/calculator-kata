from cProfile import run
from io import StringIO
from calculator.evaluator import eval_expr

if __name__ == '__main__':
    with open('../../testdata/10m', encoding='UTF-8') as f:
        run('eval_expr(f)')