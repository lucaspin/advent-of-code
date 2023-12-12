import time

lines = open('input.txt', 'r').readlines()

result = 0

def go_down(ns, part_one):
  cur = ns[len(ns)-1]

  if not all(x == 0 for x in cur):
    new = []
    for i in range(len(cur)-1):
      if part_one:
        v = cur[i+1]-cur[i]
      else:
        v = cur[i]-cur[i+1]
      new.append(v)
    ns.append(new)
    return go_down(ns, part_one)

  return ns

def go_up(ns, level, part_one):
  if level < 0:
    return ns

  below = ns[level+1]
  diff = below[len(below)-1]
  l = ns[level]
  s = l[len(l)-1] + diff if part_one else l[len(l)-1] - diff
  l.append(s)
  ns[level] = l
  return go_up(ns, level-1, part_one)

# Part 1
for line in lines:
  ns = [[int(x) for x in line.split(" ")]]
  ns = go_down(ns, True)
  ns = go_up(ns, len(ns)-2, True)
  t = ns[0]
  result += t[len(t)-1]
print("Part 1: ", result)

# Part 2
# Same thing, but we reverse line
result = 0
for line in lines:
  l = [int(x) for x in line.split(" ")]
  l.reverse()
  ns = [l]
  ns = go_down(ns, True)
  ns = go_up(ns, len(ns)-2, True)
  t = ns[0]
  result += t[len(t)-1]
print("Part 2: ", result)