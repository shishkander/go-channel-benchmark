import tabulate
import sys

def process(lines):
    A = {}
    Order = ['UNB', '001', '010', '100', '150', '01K', '10K']
    Headers = [''] + Order[1:-1] + ['']
    MAP = dict((k, i) for i, k in enumerate(Order))
    for l in lines:
        # print l.split()
        # ['BencOrdermarkChanUNBx001-4', '20000', '79303', 'ns/op', '0.05', 'MB/s']
        l = l.split()
        size, tasks = l[0].split('-')[0][-7:].split('x')
        A.setdefault(MAP[size], [size] + ['']*len(Order))[1 + MAP[tasks]] = l[-2]

    table = [A[i] for i in xrange(len(Headers)) if i in A]
    return tabulate.tabulate(table, headers=Headers, tablefmt="grid")

if __name__ == "__main__":
    M = {}
    verbatim = list(sys.stdin)
    with open("output.txt", "w") as f:
        f.writelines(verbatim)
    for l in verbatim:
        if not l.startswith("BenchmarkChan"):
            continue
        part = l.split()[0]
        if "-" in part:
            key = int(part.split("-")[1])
        else:
            key = 1
        M.setdefault(key, []).append(l)
    prefix = ' ' * 8
    for i, k in enumerate(sorted(M)):
        out = str(process(M[k]))
        print "%d. NumCPU = %s" % (i, k)
        print
        print prefix + (("\n" + prefix).join(out.splitlines()))
        print
