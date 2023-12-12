import time
import math

f = open('a.txt','r')
text = f.read()
lines = text.split("\n")

def parse_coordinate(line):
  src, dsts = line.split(" = ")
  dsts = dsts.replace("(", "")
  dsts = dsts.replace(")", "")
  left, right = dsts.split(", ")
  return (src, (left, right))

instructions = lines[0]
coordinates = dict([parse_coordinate(x) for x in lines[1:] if x != ''])

def find_steps(node, instructions, coordinates):
  cur = node
  steps = 0
  while not cur.endswith('Z'):
    for i in list(instructions):
      if i == 'R':
        cur = coordinates[cur][1]
      else:
        cur = coordinates[cur][0]

      steps += 1
      if cur.endswith('Z'):
        break
  return steps

print("Part 1: ", find_steps('AAA', instructions, coordinates))

starting_points = [x for x in coordinates.keys() if x.endswith('A') ]
steps = [find_steps(x, instructions, coordinates) for x in starting_points]

print("Part 2: ", math.lcm(*steps))